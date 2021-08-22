## [Kubelist Issue #100 for 2020-09-16](https://kubelist.com/issue/100)

#### A few of our favorite SIGs... 💯

> Hi again! This week, we decided to take a look at some of the lesser-known (at least we haven’t seen much discussion about them) special interest groups (SIGs) for Kubernetes. There are a lot of great SIGs, so instead of providing links to the popular repos, we focus on some that seem to be flying under the radar. If these are interesting to you, go digging around in kubernetes-sigs for more! We didn’t have enough room in this issue to share everything that we found!

1. [Node Feature Discovery](https://github.com/kubernetes-sigs/node-feature-discovery)

    The node-feature-discovery application is a Kubernetes Deployment and Daemonset that detect as much information as possible on the node, and add labels. I’m not sure why you’d need to schedule a pod on a node that specially has the X86 AESNI instruction set available; but it’s great to know that you can! Turns out it is really useful to be able to automatically label nodes based on their specific hardware capabilities, and then use that in Affinity and Anti-Affinity scheduling policies. 🏷
1. [External DNS](https://github.com/kubernetes-sigs/external-dns)

    This is a great idea. This SIG is an in-cluster utility to automatically update DNS on external DNS providers when an Ingress or service IP changes. There’s a huge list of supported DNS providers already, so chances are it will work for you out of the box. This project is a Deployment, not a CustomResourceDefinition, so it’s not going to need any special or elevated permissions to install.
1. [Descheduler](https://github.com/kubernetes-sigs/descheduler)

    In normal operations, Kubernetes schedules pods and then monitors them to reschedule if something goes wrong. But once a pod is scheduled and running, Kubernetes won’t normally reschedule it. This can sometimes create unbalanced workloads and nodes that have more pods than others if a service is removed from the cluster. The Descheduler SIG will watch for certain events and, well, “deschedule” pods. At this point, the scheduler will detect that it has work to do and schedule it again (hopefully somewhere different). 🗓
1. [Secrets Store CSI Driver](https://github.com/kubernetes-sigs/secrets-store-csi-driver)

    The Secret Store CSI Driver is an interesting approach to making swappable backend stores for secrets. While they currently only support Azure and Vault, it’s possible to write your own provider. This is a CSI driver that will mount secrets from enterprise-grade external secret stores as pod volumes, allowing the pod to run without knowledge of how to retrieve secrets from a specific store. 🗝
1. [Cluster Capacity](https://github.com/kubernetes-sigs/cluster-capacity)

    If you’re not already setting resource requests and limits on your pods, you may want to start now, as they’re a requirement to use the Cluster Capacity CLI. This nifty utility will look at a podspec and the pods already allocated to try to give you an estimate of how many replicas of a pod will fit on a node/cluster. Especially useful when pod resources are relative units and we have operators creating new resources constantly. 📏
1. [Clientgofix](https://github.com/kubernetes-sigs/clientgofix)

    If you have an operator or some client-go code and want to update to Kubernetes 1.18, there’s some breaking changes in the client-go library. This helpful project will update all of your code to be compatible with the new function signatures. It’s definitely better than making the changes all over your code!
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1303766370723618816)

    We’d be sad if we didn’t share a link to register for Virtual KubeCon in November!

### [ << Prev ](kubelist-099.md) ------------- [ Next >> ](kubelist-101.md)