## [Kubelist Issue #80 for 2020-04-15](https://kubelist.com/issue/80)

#### Kubelet replacements

> This week we take a look at three alternate Kubelets you can run or even build, and share some recent announcements.

1. [Introducing Krustlet, the WebAssembly Kubelet](https://deislabs.io/posts/introducing-krustlet/)

    Easily deploy WebAssembly workloads on Kubernetes using this alternative Kubelet. Bonus points for writing it in rust and creatively naming it Krustlet. This project is early (how many of us run WebAssembly AND Kubernetes in production?), but is a great example of how the Kubernetes API is becoming a standard.
1. [Virtual Kubelet](https://github.com/virtual-kubelet/virtual-kubelet#virtual-kubelet)

    Having an idea for another Kubelet alternative? This CNCF Sandbox project contains the building blocks to get started. Who knows, maybe in the future, custom Kubelets will be as popular as Operators are today! 😱
1. [Running Kubernetes without nodes](https://jpetazzo.github.io/2019/02/13/running-kubernetes-without-nodes-with-kiyot/)

    Instead of having idle capacity, what would happen if the Kubelet was configured with cloud provider credentials and created instances for each and every pod? Nodeless did exactly this by creating a virtual kubelet, and it's ready to try today. Nodeless is a pretty clever way to run somewhat isolated and dynamic workloads in Kubernetes.
1. [Argo accepted as a CNCF Incubator project](https://www.cncf.io/blog/2020/04/07/toc-welcomes-argo-into-the-cncf-incubator/)

    While we still wait to hear how Argo and Flux are going to work together, it's pretty interesting to see Argo move from sandbox up to incubating status. That's a big step, congratulations! 👏 Excited to see what comes next from Argo.
1. [Open Source Service Mesh Hub from Solo.io](https://www.solo.io/blog/open-source-service-mesh-hub-technical-overview/)

    One of the most common complaints about running a service mesh is the additional complexity. Solo.io has released an open source service mesh hub to help manage multiple service meshes in a cluster. Now, in addition to VirtualServices, we can have VirtualMeshes. The future is meshes all the way down.
1. [Portworx Essentials (free!)](https://portworx.com/announcing-portworx-essentials/)

    Another option in the free & open source storage providers for Kubernetes. We know about Rook (and ceph), and we know about OpenEBS. These are already some popular options, and it's now great to see Portworx come in with a free version, even if it's limited to 5 nodes and 5 TB of storage. 
1. [Tweet of the week](https://twitter.com/kelseyhightower/status/1249906545917186054)

    Kelsey asks a simple question and the replies are great.

### [ << Prev ](kubelist-79.md) ------------- [ Next >> ](kubelist-81.md)