# =============================================================================
# Dockerfile — Multi-stage build for a Go microservice
# =============================================================================
#
# INTERVIEW CONTEXT:
# Multi-stage builds are one of the most important Docker concepts for Go.
# They let you compile in a full SDK image but ship only the static binary
# in a minimal runtime image. The result is a production image that is
# typically < 20 MB instead of > 800 MB.
#
# Key topics covered here:
#   1. Multi-stage builds (small, secure production images)
#   2. Layer caching strategy (fast rebuilds)
#   3. Distroless base images (minimal attack surface)
#   4. Non-root execution (security best practice)
#   5. ENTRYPOINT vs CMD
#   6. Health checks in container vs orchestrator
# =============================================================================


# ---------------------------------------------------------------------------
# Stage 1: BUILD
# ---------------------------------------------------------------------------
# We use the full Go SDK image with Alpine Linux for a smaller download.
# "AS builder" names this stage so we can reference it later with
# COPY --from=builder. Everything in this stage is thrown away in the final
# image — only explicitly copied artifacts survive.
# ---------------------------------------------------------------------------
FROM golang:1.23-alpine AS builder

# Set the working directory inside the build container.
# All subsequent commands (COPY, RUN, etc.) are relative to this path.
WORKDIR /app

# ---------------------------------------------------------------------------
# Layer Caching Optimization (INTERVIEW FAVORITE)
# ---------------------------------------------------------------------------
# Docker builds images layer by layer. Each instruction creates a layer.
# If a layer's inputs haven't changed, Docker reuses the cached layer and
# skips re-executing the instruction.
#
# Strategy: Copy dependency manifests FIRST, then download dependencies,
# and ONLY THEN copy the rest of the source code.
#
# Why this matters:
#   - go.mod and go.sum change infrequently (only when deps change).
#   - Source code changes on every commit.
#   - By separating these steps, `go mod download` is cached across builds
#     as long as go.mod/go.sum are unchanged — even if source code changed.
#   - This saves significant time because downloading all dependencies
#     can take minutes on large projects.
#
# If we did `COPY . .` followed by `go mod download`, the cache would be
# invalidated on EVERY source code change, forcing a full re-download.
# ---------------------------------------------------------------------------
COPY go.mod go.sum ./

# Download all Go module dependencies into the module cache.
# This layer is cached until go.mod or go.sum changes.
RUN go mod download

# Now copy the rest of the application source code.
# This layer is invalidated on every code change, but that's fine —
# the expensive dependency download above is already cached.
COPY . .

# ---------------------------------------------------------------------------
# Build the Go binary
# ---------------------------------------------------------------------------
# CGO_ENABLED=0:
#   Disables cgo, which means the binary won't link against any C libraries.
#   This produces a fully static binary that can run in a container with
#   NO libc (like scratch or distroless). Without this, the binary might
#   depend on glibc/musl and fail to start in minimal images.
#
# GOOS=linux:
#   Cross-compile target. Ensures we build a Linux binary even if the
#   developer is on macOS or Windows. Essential for CI/CD pipelines.
#
# -ldflags="-s -w":
#   -s  strips the symbol table (function names, etc.)
#   -w  strips DWARF debugging information
#   Together they reduce binary size by ~30% with no runtime impact.
#   Trade-off: stack traces will have less detail, and you can't use
#   delve (debugger) on the stripped binary. Fine for production.
#
# -o /app/server:
#   Output path for the compiled binary.
# ---------------------------------------------------------------------------
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/server


# ---------------------------------------------------------------------------
# Stage 2: RUNTIME
# ---------------------------------------------------------------------------
# This is the image that actually ships to production.
# We start FROM a completely different base — only the artifacts we
# explicitly COPY from the builder stage end up here.
#
# ---------------------------------------------------------------------------
# Base Image Choice — Distroless vs Scratch vs Alpine
# ---------------------------------------------------------------------------
#
# scratch:
#   - Literally empty. Zero bytes. The absolute smallest image.
#   - No shell, no libc, no CA certificates, no timezone data, no users.
#   - You must bundle CA certs and tz data yourself if your app needs them.
#   - Debugging is impossible (can't exec into the container).
#
# alpine (~5 MB):
#   - Minimal Linux distribution with musl libc, shell, and apk package
#     manager.
#   - Good for debugging (you can exec in and run shell commands).
#   - Slightly larger attack surface because it includes a package manager
#     and shell that an attacker could exploit.
#
# distroless (~2 MB for static variant):
#   - Google-maintained images containing ONLY the runtime (CA certs,
#     timezone data, /etc/passwd with a nonroot user).
#   - No shell, no package manager, no ls, no cat — nothing an attacker
#     can use for lateral movement.
#   - Smaller attack surface than alpine but more convenient than scratch
#     (includes CA certs and tz data out of the box).
#   - RECOMMENDED for production Go services.
#
# For a statically linked Go binary, use the "static" variant.
# For a dynamically linked binary (CGO_ENABLED=1), use the "base" variant
# which includes glibc.
# ---------------------------------------------------------------------------
FROM gcr.io/distroless/static-debian12

# ---------------------------------------------------------------------------
# Copy the compiled binary from the builder stage.
# --from=builder references the first stage by its AS name.
# This is the ONLY artifact that crosses the stage boundary.
# The entire Go SDK, source code, and module cache are left behind.
# ---------------------------------------------------------------------------
COPY --from=builder /app/server /server

# ---------------------------------------------------------------------------
# Copy configuration files if needed.
# In a real project you might copy migration files, templates, etc.
# Environment-specific config should come from env vars or mounted volumes,
# not baked into the image (12-factor app principle).
# ---------------------------------------------------------------------------
# COPY --from=builder /app/configs /configs

# ---------------------------------------------------------------------------
# Non-root Execution (SECURITY BEST PRACTICE)
# ---------------------------------------------------------------------------
# By default, containers run as root (UID 0). If an attacker exploits a
# vulnerability in your app, they have root access inside the container.
# Combined with a container escape exploit, this could mean root on the host.
#
# Distroless images ship with a built-in "nonroot" user (UID 65532).
# We switch to that user so the process runs with minimal privileges.
#
# In alpine/scratch images, you'd need to create the user manually:
#   RUN addgroup -S appgroup && adduser -S appuser -G appgroup
#   USER appuser
#
# The USER instruction sets the default user for RUN, CMD, and ENTRYPOINT
# instructions that follow it, and also for `docker exec` sessions.
# ---------------------------------------------------------------------------
USER nonroot:nonroot

# ---------------------------------------------------------------------------
# EXPOSE documents which port the container listens on.
# This is purely informational metadata — it does NOT actually publish the
# port. You still need -p 8080:8080 at runtime or a ports: mapping in
# docker-compose. Think of it as documentation for whoever reads the
# Dockerfile or inspects the image with `docker inspect`.
# ---------------------------------------------------------------------------
EXPOSE 8080

# ---------------------------------------------------------------------------
# HEALTHCHECK — Container-level Health Monitoring
# ---------------------------------------------------------------------------
# Docker supports a HEALTHCHECK instruction that periodically runs a command
# inside the container to verify the app is healthy. If the check fails
# repeatedly, the container status changes to "unhealthy".
#
# HOWEVER: Distroless images have NO shell and NO curl/wget.
# This means you CANNOT use typical health check commands like:
#   HEALTHCHECK CMD curl -f http://localhost:8080/healthz || exit 1
#
# Options for distroless:
#   1. Omit HEALTHCHECK and rely on the orchestrator (Kubernetes liveness/
#      readiness probes hit the endpoint over the network — no shell needed).
#   2. Build a tiny health-check binary into your Go app and use it:
#      HEALTHCHECK CMD ["/healthcheck"]  (exec form, no shell needed)
#   3. Use a multi-binary build where one binary is the health checker.
#
# For Kubernetes deployments, option 1 is standard practice. Kubernetes
# probes (httpGet, tcpSocket, exec) operate from outside the container
# and don't need any tooling inside it.
#
# Below is commented out because we're using distroless. Uncomment if
# you switch to alpine and have curl available, or if you add a health
# check binary.
# ---------------------------------------------------------------------------
# HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#   CMD ["/healthcheck"]

# ---------------------------------------------------------------------------
# ENTRYPOINT vs CMD (INTERVIEW QUESTION)
# ---------------------------------------------------------------------------
# ENTRYPOINT: Sets the main executable. The container essentially "is" this
#   process. When someone runs `docker run myimage`, this is what executes.
#   It cannot be easily overridden without --entrypoint flag.
#
# CMD: Provides default arguments to the ENTRYPOINT (if ENTRYPOINT is set),
#   or a default command (if ENTRYPOINT is not set). CMD is easily overridden
#   by appending arguments to `docker run myimage <args>`.
#
# Common patterns:
#   ENTRYPOINT ["/server"]                  — fixed binary, no default args
#   ENTRYPOINT ["/server"] + CMD ["--port", "8080"]  — fixed binary, default args overridable
#   CMD ["/server"]                         — fully overridable (useful for debugging)
#
# Exec form ["binary", "arg"] vs Shell form "binary arg":
#   - Exec form: runs the binary directly as PID 1. Receives signals (SIGTERM)
#     properly. ALWAYS USE THIS IN PRODUCTION.
#   - Shell form: runs via /bin/sh -c "binary arg". The shell is PID 1 and
#     may not forward signals to your app, causing ungraceful shutdowns.
#     Also impossible in distroless (no shell).
#
# We use ENTRYPOINT here because this image has one purpose: run the server.
# ---------------------------------------------------------------------------
ENTRYPOINT ["/server"]
