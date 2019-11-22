# Kubernetes Architecture

Kubernetes is an orchestrator for microservice apps

Kubernetes organizes various services so that they work together on the right networks with the right secrets. These services include:
    - Load balancer
    - Search
    - HTTPS
    - Auth
    - Database
    - Logs

A **Cluster** is made up of one or more Masters and multiple Nodes aka *Minions*.

The **Master** is is the *control plane*.

## Masters

- A Master is a Kubernetes Control Plane
- Run your apps on the Nodes and not in the Master
- The **API Server** is the front end to the control plane. It exposes the REST API and consumes JSON.
    - by default it uses port 443
- The **Cluster Store** is a persistent storage. It uses *etcd*
- **etcd** is a NoSQL database
- The *source of truth* for the cluster
- The **kube-controller-manager* controls controllers. This helps maintans the *desired state*. This includes
    - Node controller
    - Endpoint controller
    - Namespace controller
- The **kube-scheduler** watches the *apiserver* for new pods.
    - Assigns work to nodes
        - affinity/anti-affinity
        - constraints
        - resources

# Nodes

also known as *Minions*

Nodes contains:
    - **Kublet**
        - The main Kubernetes agent
        - Registers node with the cluster
        - Watches the *apiserver*
        - Instantiates *pods*
        - Reports back to *master*
        - Exposes endpoint on :10255
    - **Container Engine**
        - Does container management:
            - Pull images
            - Start/stop containers
        - Pluggable
            - Usually *Docker*
            - Can be *rkt*
    - **kube-proxy**
        - Kubernetes networking:
            - Pod  IP addresses
                - All Containers in a pod share a single IP
            - Load balances across all pods in a service

## Desired State & the Declarative Model

Kubernetes operates on a declarative model.
    - Give the apiserver the manifest files that describe how we want the cluster to look and feel.
    - We don't give it a long list of commands.

Kubernetes uses manifest files in *YAML* or *JSON* to describe the Desired State of the cluster.
    - If actual state does not match the desired state of the cluster, k8 will keep trying until it meets the desired state.

## Pods

Ring fenced environment to run containers inside.
    - Network stack
    - Kernel namespaces
    - All containers in a pod share the pod environment
        - ex: sharing the same IP
    - Tight Coupling of two containers that share volumes or memory should go into a single pod.
    - Loose coupling of two containers can be placed in seperate pods.

- Containers always run inside *pods*.
- You can run more than one container in one pod however this is an advance used case
- You scale elements in your application by adding or removing pod replicas.
- Pods are atomic
    - They are either all up or all down.
- Pods have a single lifecycle. Pods are never reused, only replaced with new ones.

Deploying Pods
    - usually via higher level objects
    - by giving the apiserver a pod manifest file
    - via a replication controller
        - deploying multiple replicas of a single pod defintion and that they are always running.
