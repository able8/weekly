## [Kubelist Issue #14 for 2018-05-17](https://kubelist.com/issue/14)

#### The "Kubernetes has many parts to it and they are all fun to learn about" edition

> Kubernetes is exciting, there is a lot going on in the community and a lot of components to why it rocks: CRDs, sweet networking, isolation, and more more more. Join us this week to look at why folks might think it&#39;s complicated and how we might embrace that complexity and get on with our kubelife.

1. [Is K8s Too Complicated?](http://jmoiron.net/blog/is-k8s-too-complicated/)

    A hopeful (well somewhat... maybe realistic or mature?) take on what really makes Kubernetes complicated. The idea that people are growing up on Kubernetes is a pretty neat one that is probably what is going to make Cloud Native Applications and Infrastructure work.
1. [Exploring container security: Isolation at different layers of the Kubernetes stack](https://cloudplatform.googleblog.com/2018/05/Exploring-container-security-Isolation-at-different-layers-of-the-Kubernetes-stack.html?m=1)

    The last post in google's recent container security series. A pretty interesting read and way of thinking about the different levels of boundaries in a Kubernetes cluster in terms of "security" and "trust". It would have been cool if they provided more information about how RBAC policies fit into their way of thinking about isolation.
1. [The growing need for Kubernetes Configuration Management](https://leebriggs.co.uk/blog/2018/05/08/kubernetes-config-mgmt.html)

    A neat discussion continues on reddit including whether you should ever apply or delete things with kubectl manually or whether that should be managed by CI.
1. [Litmus: Release a chaos monkey on your Kubernetes Stateful Workloads!](https://blog.openebs.io/litmus-release-a-chaos-monkey-on-your-kubernetes-stateful-workloads-6345e01b637d?gi=3773b456db3)

    Probably not a chaos monkey, but end-to-end testing for Kubernetes operations, now that is cool!
1. [Analyzing value of Operator Framework for Kubernetes community](https://itnext.io/analyzing-value-of-operator-framework-for-kubernetes-community-5a65abc259ec)

    If you are managing or implementing CRDs and controllers for them, you will want to check this out. The article ends on an interesting point, that Kubernetes applications are typically going to be composed of lots of CRDs and knowing how they work and interact with each other and with your cloud native applications is still a big challenge.
1. [Tweet of the week](https://twitter.com/julian_dunn/status/995448964513067008)

    "Kubernetes 1890s edition"

### [ << Prev ](kubelist-13.md) ------------- [ Next >> ](kubelist-15.md)