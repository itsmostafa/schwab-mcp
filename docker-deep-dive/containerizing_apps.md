# Containerizing an App

To get an application into an image, we use a **Dockerfile**

## The bigger picture

A Dockerfile is:
- A list of instructions on how to build an image with our app code inside

To build an image with a Dockerfile:
- `docker image build`

### Dockerfile notes
- Instructions for building images
- CAPITALIZE the instructions
- FROM is always the first instruction
- Good practice to list the maintainer
- RUN = execute command and create a new layer
- COPY = copy code into image as a new layer
- Some instructions add metadata instead of layers
- ENTRYPOINT = default app for the image/container

### Digging deeper

Build context
- Location of your source code
- Can be a remote git repo

When running `docker image build`, whatever is in there gets sent to the daemon and gets processed in the build.