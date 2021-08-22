## [Kubelist Issue #85 for 2020-05-20](https://kubelist.com/issue/85)

#### Lessons Learned 📚

> As organizations of all sizes transition to Kubernetes, a lot has been written and shared about the journey. There are plenty of success stories and also many posts about mistakes and lessons learned. Whether you have already moved to Kubernetes or just starting out, there’s a lot to learn from these stories. This week’s issue is a collection of links that share some lessons that have been learned the hard way so that you can avoid making the same mistakes.

1. [10 most common mistakes using Kubernetes](https://blog.pipetail.io/posts/2020-05-04-most-common-mistakes-k8s/)

    Ask any early Kubernetes adopter to list 10 lessons they had to learn the hard way, and it’s going to look a lot like this collection. If you are just starting to move to Kubernetes, save this link and build your internal processes to avoid these common mistakes. Some of these fall into “well known”, such as setting resource limits, while others are more subtle and too often learned during an outage.🙅
1. [There and Back Again — Scaling Multi-Tenant Kubernetes Cluster(s)](https://medium.com/usa-today-network/there-and-back-again-scaling-multi-tenant-kubernetes-cluster-s-67afb437716c)

    This is one of the best “big company adopted Kubernetes” stories that’s been posted in a while. The team at USA Today didn’t simply write another “we moved to Kubernetes” article, but did a great job describing the whole journey. There’s a lot here, including the decision to have one cluster or many and evaluating gitops or other (legacy) deployment processes. You might not have 40 different teams trying to use your cluster(s), but this is a good look at the platform-team approach to delivering Kubernetes to an org.
1. [What Can We Learn from Twitter's Move to Kubernetes](https://www.alibabacloud.com/blog/what-can-we-learn-from-twitters-move-to-kubernetes_595156)

    Before Kubernetes, Twitter created Mesos and used it to power their entire platform. This was the same time that a few other large companies were building their own internal, modern platforms (Google’s Borg, Facebook’s Tupperware, Uber’s Peloton, etc). When Twitter moved from their own product to Kubernetes, they’ve sent a clear signal that Kubernetes has staying power and is a safe bet in this ecosystem. 🐦
1. [Building a Kubernetes platform at Pinterest](https://medium.com/pinterest-engineering/building-a-kubernetes-platform-at-pinterest-fb3d9571c948)

    An interesting writeup describing how Pinterest has created a set of internal CustomResourceDefinitions to handle the complexity of their infrastructure. The system they’ve developed removes the general Kubernetes expertise needed to interop with their cluster, and decided to make it more specialized. Developers don’t have to write the same boilerplate YAML this way. And because the Kubernetes objects are generated at runtime, this is a way to ensure that everything conforms to the requirements. It would be fun to compare this approach to using more common tooling like OPA Gatekeeper and AdmissionControllers. 📍
1. [Migrating to Kubernetes (Sensu)](https://blog.sensu.io/migrating-to-kubernetes)

    Sort of a “moving to Kubernetes 101” crash course, designed for the beginner. This post starts at containerization and proceeds to explain the core fundamentals of Kubernetes, with the goal of explaining these to someone considering adopting the platform. If you are new to Kubernetes and want a very high level view of the path to Kubernetes-based immutable infrastructure, this is a good place to start. 🗺
1. [Develop Hundreds of Kubernetes Services at Scale with Airbnb](https://www.infoq.com/presentations/airbnb-kubernetes-services/)

    We decided to include this one for the custom kubectl plugins section. Airbnb enables their engineering team to perform cluster management with internal kubectl plugins instead of remembering a lot of commands and running debug containers and more. Building these commands and distributing them to the engineering team is an interesting way to distribute the knowledge of debugging. We should dedicate an issue to debugging Kubernetes clusters soon.
1. [Tweet of the week](https://twitter.com/jimmangel/status/1260751239211597824?s=21)

    Life as an early adopter does not come easy.

### [ << Prev ](kubelist-84.md) ------------- [ Next >> ](kubelist-86.md)