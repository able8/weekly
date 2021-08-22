## [Kubelist Issue #74 for 2019-08-01](https://kubelist.com/issue/74)

#### Dealing with Networking Issues in Kubernetes

> Today, we’re opening common networking issues within Kubernetes. Let’s take a look at what Kubernetes Networking is dealing with, to specific issues within K8s. And solutions too.

1. [Kubernetes Networking: DNS, Communicating, and Networking Model](https://www.aquasec.com/wiki/display/containers/Kubernetes+Networking+101)

    First and foremost, it’s important to understand Kubernetes networking. This article dives into DNS, Communication, and Networking Model, a great start to understanding K8s networking in depth. 
1. [4 Big Networking Issues within Kubernetes](https://kubernetes.io/docs/concepts/cluster-administration/networking/)

    This link will take you to explore the four areas of Networking that Kubernetes struggles to resolve. It also introduces the Kubernetes Network Model and how to implement it in your own clusters. 
1. [Pod to Pod Networking: What is the Problem?](https://cloudnativelabs.github.io/post/2017-04-18-kubernetes-networking/)

    The main issue with pod to pod networking is that it isn’t supported with CNI plug-ins. This makes it somewhat difficult for pod to pod communication, but there are a few solutions that can help with this.
1. [Containers to Containers to Communication](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/)

    Normally, containers can only communicate with other containers if they are on the same machine, but this makes it difficult to communicate across nodes. The solution? Dynamically allocated ports. 
1. [External to Services: How these Communicate](https://dzone.com/articles/kubernetes-networking)

    An Ingress Network may be what you’re looking for with connecting externally to services. It supports external to service connection in a powerful way. 
1. [MORE Containers to Containers Communication](https://www.mirantis.com/blog/multi-container-pods-and-container-communication-in-kubernetes/)

    Included here are even more container to container communication. It also takes a deeper look into shared volumes in a pod and inter-process communication.
1. [Tweet of the Week](https://twitter.com/b0rk/status/1108164989012135936)

    @b0rk created a visual comic on Kubernetes Networking problems that also exist. She gives a few problems listed and how to recognize these failures. 

### [ << Prev ](kubelist-73.md) ------------- [ Next >> ](kubelist-75.md)