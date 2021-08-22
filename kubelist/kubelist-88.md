## [Kubelist Issue #88 for 2020-06-24](https://kubelist.com/issue/88)

#### Let’s Talk Storage 💾

> This week, we’ve been thinking a lot about storage in Kubernetes. Storage is an extremely important, but also a very difficult capability to provide in a cluster. Doing so often means relying on external resources to create a reliable solution. There are some projects that let you run your cluster on EKS without EBS volumes, or to run a bare metal cluster that doesn’t come with any managed storage. This is a topic that we should check back in on more often, and we’d love to hear about other in-cluster storage solutions not listed here!

1. [Why Is Storage On Kubernetes So Hard](https://softwareengineeringdaily.com/2019/01/11/why-is-storage-on-kubernetes-is-so-hard/)

    Let’s start by looking at why storage is such a hard problem. This article very succinctly explains the cause: “Persistent storage cannot be bound to the rules of being dynamically created and destroyed”. Applying the Kubernetes concepts to in-cluster storage is a non-trivial problem, but it can be done. There are several open source (and some proprietary) projects and products to help here, and the rest of this Kubelist issue looks at some of the most popular.
1. [Rook](https://rook.io/docs/rook/v1.3/)

    Rook is an Incubation-level CNCF project, which makes it the only storage provisioner to get out of the Sandbox to date. Rook is a Kubernetes Operator to operate Ceph and other storage systems. The list of Alpha storage systems (non-Ceph) supported by Rook is extensive and worth watching, especially if you aren’t excited about running Ceph in production. 🏰
1. [Longhorn](https://longhorn.io/docs/1.0.0/what-is-longhorn/)

    Longhorn recently released v1.0.0 and is an appealing choice when you have a greenfield environment running completely in Kubernetes and you cannot (or choose not to) rely on cloud provider storage. Longhorn doesn’t have the benefit of building on top of Ceph or another mature project, but this means Longhorn doesn’t come with the baggage of pre-Kubernetes projects. If simplicity and support for Kubernetes-only environments are big goals, Longhorn is worth checking out. 🐮
1. [OpenEBS](https://docs.openebs.io/docs/next/casengines.html)

    OpenEBS brings multiple Storage Engines, and even allows you to enable multiple engines at the same time. This is powerful if you want to use different storage classes for different workloads in your StatefulSets and Persistent Volumes. By enabling cStor and LocalPV both, you can allow a workflow for workloads to specify (via storage classes) whether their data needs to be replicated across nodes or not. This is useful when you want to run Postgres and Elasticsearch both on the same cluster, using the same storage system, but having different QoS for each. LocalPV is a good choice when the application manages its own replication and cStore is useful when the application relies on the underlying storage provider. 
1. [ChubaoFS](https://chubao.io/)

    ChubaoFS provides both a block store and object store while providing a focus on scaling the filesystem metadata service. For some workloads, querying a filesystem backed by distributed storage represents significant overhead and can create bottlenecks and other performance problems. Now that ChubaoFS has a CSI plugin for Kubernetes, it’s relatively easy to deploy into an existing cluster without requiring anything to be added to each node. 🗃
1. [Portworx Essentials](https://docs.portworx.com/concepts/portworx-essentials/)

    Portworx is known for delivering a commercial storage product, but now has Portworx Essentials, which is free for 5 TB and 5 Nodes. Essentials is an interesting choice to get an upgrade path to a commercially supported, highly scalable storage backend. Just keep an eye on the license because there are some restrictions including the requirement to have telemetry enabled to use the product. 
1. [Tweet of the week](https://twitter.com/kubernetesio/status/1273666581827063809)

    New docs. Send some feedback on what’s good or what you don’t like.

### [ << Prev ](kubelist-87.md) ------------- [ Next >> ](kubelist-89.md)