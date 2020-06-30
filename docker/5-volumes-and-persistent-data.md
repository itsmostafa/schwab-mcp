# Working with Volumes and Persistent Data

- Containers are Non-persistent and Immutable
- Volumes are Persistent and changeable
- Every container gets its own non-persistent storage managed by the storage driver / graph driver
- Persistent storage spaces in docker are called `volumes`
- Volumes are fully independant from containers

## Managing Volumes

- To attach a volume to a container add the `--mount` flag
    - example; to create a new volume called *ubervol* and attached it to an alpine container: `docker container run -dit --name voltest --mount source-ubervol target=/vol alpine:latest`
- volumes typically live in this directory within the container:
    - `ls -l /var/lib/docker/volumes/`
- you can not delete volumes as long as its in use with a container

## Recap
- Containers have local graph driver storage bound to it.
- Volumes are needed for persistent data with a lifecycle independant to a container. The volume can stick around even when attached containers are deleted
- `docker volume create | ls | inspect`
