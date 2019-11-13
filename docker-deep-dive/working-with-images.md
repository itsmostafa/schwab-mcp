# Working with Images

- ### An image is a read-only template for creating application containers. Inside of it is all the code and supporting files to run an application.
- ### Images are build-time constructs and containers are their run-time siblings.
- ### An image contains:
  - OS files & objects
  - Application files
  - Manifest

- ### Images are stored in a registry
- ### Images have a one to many relationship with containers
- ### All writes and updates happen with each container. One writable layer per container

## Images in Detail

- An image consists of multiple, independant layers that are very loosely connected by a **manifest file**
- Manifest file:
  - Also sometimes called a config file
  - Describes the Image including ID, tags, date created, etc
  - Includes the list of layers that get stacked and how to stack them

Here's an example of pulling a Redis image

1. Run the command `docker image pull redis`
2. The command generates an API request to the docker registry API, presumably the registry: Docker Hub.
3. Get Manifest
4. Get Layers


## Registries

- Images live in registries
- Docker defaults to Docker Hub
- can pull from an on premise registry and other unofficial registries

You can pull a specific image by specifying:
- Registry / Repo / Image : (Tag)
- For example:
  - `docker image pull docker.io/redis:4.0.1`

Push an image to a registry
- `docker image push`

Pull an image from a registry
- `docker image pull`

View image config including layer data
- `docker image inspect`

Remove old images
- `docker image rm`
