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

