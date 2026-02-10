# Docker

## Key Concepts

### Images, containers, and layers

- An **image** is a read-only template built from a series of layers. Each Dockerfile instruction (`FROM`, `COPY`, `RUN`) creates one layer.
- A **container** is a running instance of an image with a thin read-write layer on top. Containers share the host kernel — they are processes with isolated namespaces (pid, net, mnt, uts, ipc) and cgroups (CPU, memory limits), not virtual machines.
- Layers are content-addressable and shared across images. If two images use the same `golang:1.23-alpine` base, that layer is stored only once on disk.
- The **union filesystem** (overlay2 on Linux) stacks layers so the container sees a single merged view. Writes go to the container's own layer (copy-on-write).

### Dockerfile instructions reference

- `FROM` — sets the base image. Every Dockerfile starts here. Can be `scratch` (empty), `alpine`, `debian`, or a language-specific image.
- `WORKDIR` — sets the working directory for subsequent instructions. Creates the directory if it doesn't exist. Prefer over `RUN mkdir && cd`.
- `COPY` — copies files from build context into the image. Prefer over `ADD` unless you need URL fetching or automatic tar extraction.
- `RUN` — executes a command in a new layer. Combine related commands with `&&` to reduce layer count.
- `ENV` — sets environment variables persisted in the image. Available at both build and run time.
- `ARG` — build-time variables only (not available at run time). Useful for version pinning.
- `EXPOSE` — documents which port the container listens on. Does not publish the port — that requires `-p` at runtime.
- `ENTRYPOINT` — the executable that always runs. Prefer exec form `["binary", "arg"]` over shell form.
- `CMD` — default arguments to `ENTRYPOINT`, or the default command if `ENTRYPOINT` is not set. Overridden by `docker run <args>`.
- `HEALTHCHECK` — built-in liveness probe. Docker marks the container as `healthy`, `unhealthy`, or `starting` based on the check result.

### Multi-stage builds for Go

Multi-stage builds are the standard pattern for Go services. The idea: compile in a fat image with the full toolchain, then copy the static binary into a minimal runtime image.

```dockerfile
# ---- Build stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Cache module downloads (these layers invalidate only when deps change)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/server

# ---- Runtime stage ----
FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
```

Key decisions:

- **`CGO_ENABLED=0`**: produces a statically linked binary with no libc dependency. This is what makes `scratch` and distroless viable as runtime bases.
- **`-ldflags="-s -w"`**: strips debug symbols and DWARF info, reducing binary size by ~30%.
- **Build stage base** — `golang:1.23-alpine` or `golang:1.23-bookworm`. Alpine is smaller; Bookworm is needed if you require CGO or system libraries.
- **Runtime stage base** options:
  - `scratch` — truly empty. Smallest image (~5-15 MB total). No shell, no CA certs, no timezone data, no user database. You must bundle anything extra yourself.
  - `gcr.io/distroless/static` — Google's minimal image. Includes CA certificates, timezone data, `/etc/passwd` with a `nonroot` user. Slightly larger than scratch but more practical.
  - `alpine` — small (~7 MB base), includes shell and package manager. Useful when you need debugging tools or the binary needs CGO/musl.
- Interview note: if asked "why not just use `golang:*` as your final image?" — the Go toolchain image is ~800 MB. A distroless Go binary image is ~15 MB. Less surface area, faster pulls, smaller attack surface.

### Layer caching — instruction ordering for fast rebuilds

Docker caches each layer. If the instruction and its inputs haven't changed, the cached layer is reused. Once a layer is invalidated, all subsequent layers are rebuilt.

The optimization principle: **put things that change rarely at the top, things that change often at the bottom.**

For Go projects, the optimal ordering is:

1. `FROM` / `WORKDIR` — almost never changes.
2. `COPY go.mod go.sum` then `RUN go mod download` — dependencies change infrequently. This layer is cached across most builds.
3. `COPY . .` — source code changes on every build, so this goes last.
4. `RUN go build ...` — runs after source copy, rebuilt every time source changes.

Bad ordering (everything invalidates on any source change):

```dockerfile
# BAD: any file change busts the module download cache
COPY . .
RUN go mod download
RUN go build -o /app/server ./cmd/server
```

Good ordering (module cache survives source changes):

```dockerfile
# GOOD: module layer cached unless go.mod/go.sum change
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/server ./cmd/server
```

Additional caching tips:

- Use `.dockerignore` to exclude files that should not enter the build context (e.g., `.git/`, `README.md`, `docs/`, test fixtures). Smaller context = faster `COPY` and fewer cache busts.
- BuildKit mount caches (`--mount=type=cache,target=/go/pkg/mod`) persist the module cache across builds even when `go.mod` changes. This is faster than `go mod download` for incremental changes.

### Docker Compose — multi-service local environments

Docker Compose defines and runs multi-container applications with a single `docker-compose.yml` (or `compose.yaml`). For the Integration Sync Service, a typical setup:

```yaml
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://user:pass@postgres:5432/syncdb?sslmode=disable
      - JAEGER_ENDPOINT=http://jaeger:4318
    depends_on:
      postgres:
        condition: service_healthy
    restart: unless-stopped

  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: pass
      POSTGRES_DB: syncdb
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U user -d syncdb"]
      interval: 5s
      timeout: 3s
      retries: 5

  jaeger:
    image: jaegertracing/all-in-one:1.53
    ports:
      - "16686:16686"  # UI
      - "4318:4318"    # OTLP HTTP receiver
    environment:
      - COLLECTOR_OTLP_ENABLED=true

volumes:
  pgdata:
```

Important Compose concepts:

- **`depends_on` with `condition`**: `service_started` (default) only waits for the container to start. `service_healthy` waits for the healthcheck to pass — use this for databases.
- **Service name as hostname**: within the Compose network, `postgres` resolves to the Postgres container's IP. No hardcoded IPs needed.
- **`restart: unless-stopped`**: auto-restart on crash but not when you explicitly stop the container.
- **Named volumes vs bind mounts**: named volumes (`pgdata:`) are managed by Docker and persist across `docker-compose down`. Bind mounts (`.:/app`) map host directories into the container.
- **Profiles**: group optional services (e.g., debugging tools) so they only start when you pass `--profile debug`.
- `docker compose up --build` rebuilds images before starting. `docker compose down -v` removes containers and volumes (useful for clean resets).

### Volumes — data persistence and development mounts

Docker containers have an ephemeral filesystem. When a container is removed, its writable layer is deleted. Volumes solve this.

**Named volumes** — managed by Docker, stored at `/var/lib/docker/volumes/` on the host:

```yaml
volumes:
  pgdata:

services:
  postgres:
    volumes:
      - pgdata:/var/lib/postgresql/data
```

- Survive `docker-compose down` (but not `down -v`).
- Best for database storage, caches, and any data you want to persist between restarts.
- Docker handles permissions and lifecycle.

**Bind mounts** — map a host directory directly into the container:

```yaml
services:
  app:
    volumes:
      - ./src:/app/src    # host path : container path
```

- Useful for development: edit files on the host, see changes in the container immediately.
- Not portable (depends on the host filesystem path).
- Can cause permission issues when the host UID doesn't match the container UID.

**tmpfs mounts** — in-memory, never written to disk. Useful for secrets or scratch data that shouldn't persist:

```yaml
services:
  app:
    tmpfs:
      - /tmp
```

### Networking — service discovery and port mapping

Docker Compose creates a default bridge network for each project. All services join it automatically.

- **Service discovery**: containers address each other by service name. `app` connects to `postgres:5432` using the service name as the DNS hostname.
- **Port mapping** (`-p` or `ports:`): `"8080:8080"` maps host port 8080 to container port 8080. Format is `host:container`. Only needed for ports you want to access from outside the Docker network.
- **Inter-service communication** does not require port mapping. Services talk to each other on the internal network using container ports directly.
- **Custom networks** allow you to isolate groups of services:

```yaml
networks:
  frontend:
  backend:

services:
  app:
    networks: [frontend, backend]
  postgres:
    networks: [backend]       # not reachable from frontend
  nginx:
    networks: [frontend]
```

- **DNS resolution**: Docker's embedded DNS server resolves service names to container IPs. Round-robin for scaled services (`docker compose up --scale app=3`).
- Network modes: `bridge` (default, isolated), `host` (shares host network stack — no isolation but no NAT overhead), `none` (no networking).

### Health checks

Health checks let Docker (and orchestrators) know whether a container is actually ready to serve traffic, not just running.

```dockerfile
HEALTHCHECK --interval=15s --timeout=3s --start-period=10s --retries=3 \
  CMD ["/server", "-healthcheck"]
```

Or in Compose:

```yaml
healthcheck:
  test: ["CMD", "curl", "-f", "http://localhost:8080/healthz"]
  interval: 15s
  timeout: 3s
  start_period: 10s
  retries: 3
```

- `interval` — how often to run the check.
- `timeout` — how long to wait for the check command to return.
- `start_period` — grace period after container start during which failures don't count. Gives the app time to initialize.
- `retries` — consecutive failures needed before marking `unhealthy`.
- Container status transitions: `starting` → `healthy` or `unhealthy`.

For Go services, a common pattern is a `/healthz` endpoint that returns 200 when the service is ready (DB connected, dependencies reachable). For distroless images without `curl`, compile a small health-check binary or use the main binary with a `-healthcheck` flag that performs an HTTP GET and exits.

### Security best practices

- **Run as non-root**: add `USER nonroot` (distroless) or create a dedicated user in the Dockerfile. Running as root inside a container is a privilege escalation risk if there's a container escape vulnerability.

```dockerfile
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
```

- **Use specific image tags**: `golang:1.23-alpine`, not `golang:latest`. Pinning prevents silent base image changes from breaking builds.
- **Scan images**: tools like `docker scout`, `trivy`, or `grype` detect known CVEs in base images and dependencies. Integrate into CI.
- **Minimize the image**: fewer packages = fewer CVEs. Distroless and scratch images have almost no attack surface beyond your binary.
- **Don't store secrets in the image**: never `COPY .env` or put credentials in `ENV`. Use runtime environment variables, Docker secrets, or a secrets manager.
- **Use `.dockerignore`**: prevents `.git/`, `.env`, credentials files, and other sensitive or unnecessary files from entering the build context.

```
# .dockerignore
.git
.env
*.md
docs/
tmp/
```

- **Read-only filesystem**: run containers with `--read-only` and use tmpfs for directories that need writes. This limits the impact of a compromised process.

### Build context and .dockerignore

The build context is the set of files sent to the Docker daemon when you run `docker build`. By default, it's the directory you specify (usually `.`).

- A large build context slows down builds because everything is tarred and sent to the daemon before any instruction runs.
- `.dockerignore` works like `.gitignore`: patterns listed are excluded from the context.
- Essential excludes for Go projects:

```
.git
.github
*.md
docs/
vendor/  # if using module proxy
tmp/
.env
.env.*
```

### Multi-platform builds

Go's cross-compilation makes multi-platform Docker images straightforward:

```dockerfile
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder
ARG TARGETOS TARGETARCH

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /app/server ./cmd/server

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

- `$BUILDPLATFORM` — the platform of the machine running the build (keeps the builder stage native and fast).
- `$TARGETOS` / `$TARGETARCH` — the platform you're building for (e.g., `linux/arm64`).
- Build with: `docker buildx build --platform linux/amd64,linux/arm64 -t myapp:latest .`
- Interview relevance: deploying Go services to ARM-based cloud instances (AWS Graviton, GCP Tau) is increasingly common for cost savings.

### Docker in CI/CD pipelines

- **Build and push**: CI builds the Docker image, tags it (usually with git SHA or semver), pushes to a registry (Docker Hub, ECR, GCR, GHCR).
- **Layer caching in CI**: use `--cache-from` with a registry-backed cache or BuildKit's cache export/import to avoid rebuilding unchanged layers on every CI run.
- **Image scanning**: run `trivy image myapp:latest` as a CI step. Fail the pipeline on critical/high CVEs.
- **Tagging strategy**: use both a mutable tag (`latest`, `main`) and an immutable tag (git SHA: `myapp:abc1234`). Deploy by immutable tag to ensure reproducibility.
- **Least-privilege CI**: CI should only have push access to the image registry, not production deployment credentials.

## Interview Questions

1. What is the difference between an image and a container?
2. Explain how Docker layers work and why layer ordering in a Dockerfile matters.
3. Walk through a multi-stage Dockerfile for a Go service. Why use multiple stages?
4. What is the difference between `scratch`, `distroless`, and `alpine` as runtime base images? When would you choose each?
5. Why do we set `CGO_ENABLED=0` when building Go binaries for Docker?
6. How do you optimize Docker layer caching for a Go project? What happens if you `COPY . .` before `go mod download`?
7. Explain the difference between `ENTRYPOINT` and `CMD`.
8. How does Docker Compose networking work? How do services discover each other?
9. What is the difference between named volumes and bind mounts? When would you use each?
10. How do health checks work in Docker? What are the container health states?
11. Why should containers run as a non-root user?
12. What is a `.dockerignore` file and why is it important?
13. How does `depends_on` with `condition: service_healthy` differ from plain `depends_on`?
14. How would you build a multi-platform Docker image for a Go service?
15. How do you handle secrets in Docker without baking them into the image?
16. What is the difference between `COPY` and `ADD`?
17. Explain the difference between Docker bridge, host, and none network modes.
18. How do you persist database data across container restarts in a Compose setup?
19. What tools would you use to scan Docker images for vulnerabilities in CI?
20. What is BuildKit and how does `--mount=type=cache` improve Go build times?

Practice prompts:

- Write a multi-stage Dockerfile for the Integration Sync Service with optimized layer caching, non-root user, and a health check.
- Write a `docker-compose.yml` for the full local stack (Go app, Postgres, Jaeger) with health checks and `depends_on` conditions.
- Take an existing naive Dockerfile (single stage, `COPY . .` first) and optimize it for build speed and image size. Explain each change.

## Resources

- Dockerfile reference: https://docs.docker.com/reference/dockerfile/
- Docker Compose file reference: https://docs.docker.com/compose/compose-file/
- Multi-stage builds: https://docs.docker.com/build/building/multi-stage/
- Best practices for writing Dockerfiles: https://docs.docker.com/develop/develop-images/dockerfile_best-practices/
- Google distroless images: https://github.com/GoogleContainerTools/distroless
- Docker networking overview: https://docs.docker.com/network/
- Docker volumes: https://docs.docker.com/storage/volumes/
- Docker BuildKit: https://docs.docker.com/build/buildkit/
- Trivy container scanner: https://github.com/aquasecurity/trivy
- Official Go Docker images: https://hub.docker.com/_/golang
