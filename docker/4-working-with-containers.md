# Working with Containers

- Containers are:
  - running instances of images
  - the smallest unit of work in Docker
  - a thin, writable layer that latched on top of an image's read only layer
  - an execution environment for an app
  - all about the applications
  - treated as ephemeral and immutable
    - Immutable meaning you won't be logging into them and poking around. Instead of changes need to be done, we build a new image instead and switch out the old container with a new one.
- There is a one to many relationship between an image and containers
- Linux containers can *only* work on a Linux kernal
- Windows containers can *only* work on a Windows kernal
- Each container generally runs a single process and has a single job and have them all connected via apis
- Each container can be started, stopped, restarted, and deleted

**Copy-on-write**
- A strategy of sharing and copying files for maximum efficiency. If a file or directory exists in a lower layer within the image, and another layer (including the writable layer) needs read access to it, it just uses the existing file.

## Diving Deeper

To log in to a container:
    - `docker container run -it {IMAGE_NAME} sh`
- To stop a container:
    - `docker container stop {PART_OF_THE_CONTAINER_ID}`
- List all containers:
    - `docker container ls -a`
- Run commands within container:
    - `docker container exec {PART OF CONTAINER ID} {COMMAND}`

## Logging
- Daemon logs
    - systemd:
        - `journalctl -u docker.service`
    - Non-systemd:
        - `cd /var/log/messages`

- Container logs
    - STDOUT
        - command for normal outputs
    - STDERR
        - command to output error messages

One can plugin to a logging solution using the `daemon.json` logging driver.

Override per-container with `--log-driver --log-opts`

Inspect logs with `docker logs <container>`
    - Doesn't work with all drivers


## Recap
- Containers are run-time cousins of images
- Containers look and feel like a standard OS
- Containers don't contain a kernal, must communicate with kernal of the host.
- Containers are Read/Write layers. Images are Read-Only layers.
- Once containers are deployed, leave them alone, if need to be fixed, replace with a new container. No need to login and modify.
- Ability to plugin a logging system