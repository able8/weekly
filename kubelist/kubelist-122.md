## [Kubelist Issue #122 for 2021-04-07](https://kubelist.com/issue/122)

#### Istio ⛵

> Let’s talk about <a href="https://istio.io">Istio</a>. Tune in to the <a href="https://kubelist.com/podcast">podcast episode</a> released today, where <a href="https://twitter.com/craigbox">Craig Box</a> from Google joined me to discuss the history and benefits of Istio. The conversation was really eye-opening, and helped me get a much better understanding of why Istio, and how Istio compares to other platforms. In this episode Craig also does a good job of demystifying Istio (if it’s a mystery to you). While Istio is not a CNCF project, I don’t feel like we are breaking our character of discussing projects in the CNCF ecosystem, because Istio works well with Kubernetes and is built on top of <a href="https://envoyproxy.io">Envoy</a>.

1. [What is Istio?](https://istio.io/latest/docs/concepts/what-is-istio/)

    Straight from the Istio docs, let’s start with a great writeup on what Istio is and what problems the software can help with. Remember, software isn’t going to single-handedly solve your operational problems. But if you identify with any of these problem statements, you should consider Istio.
1. [Kubernetes Service Mesh: A Comparison of Istio, Linkerd and Consul](https://platform9.com/blog/kubernetes-service-mesh-a-comparison-of-istio-linkerd-and-consul/#Comparison-of-Istio-Linkerd-and-Consul-Connect-for-Kubernetes-Service-Mesh)

    Sachin Manpathak at Platform9 wrote a detailed comparison of three service meshes: Istio, Linkerd, and Consul. This isn’t an opinionated post; instead it starts out by building a comparison table. We’ve talked about the patterns that a service mesh brings to a cluster, and after the comparison table, Sachin goes into a little more detail on these patterns. This is a great post for anyone just getting started with a service mesh today. 🍎🍊
1. [Mutual TLS Migration](https://istio.io/latest/docs/tasks/security/authentication/mtls-migration/)

    Let’s say you aren’t enforcing mTLS inside your cluster. How would you get started, and how can Istio help you migrate from your insecure setup to a secure one, without breaking production traffic? This doc shows how much effort the Istio team has put into thinking through “migrating to Istio” scenarios. 🚚🚦
1. [Better External Authorization](https://istio.io/latest/blog/2021/better-external-authz/)

    This is a description of a new feature, introduced in Istio 1.9. Some folks may have complex and possibly proprietary existing authentication systems. And when you have that, it’s unlikely that you’ll change your entire auth system in order to adopt a single product. Istio now supports delegating access control decisions out to external providers. This is a great feature that could make it possible to use Istio in environments where it was difficult to use before. 🔑
1. [Virtual Machines](https://istio.io/latest/docs/setup/install/virtual-machine/)

    In the podcast, Craig mentions that Istio isn’t limited to working inside a Kubernetes cluster. Here’s a walkthrough of connecting a VM to an Istio installation. Being able to connect to external (out of the cluster) VMs isn’t unique to Istio, but it’s a really powerful feature. I’m looking forward to the day when everything is in K8s, but maybe you have some older, legacy workloads that you’d like to secure. 🖥
1. [Continuous Blue-Green Deployments with Kubernetes](https://semaphoreci.com/blog/continuous-blue-green-deployments-with-kubernetes)

    A good writeup by Tomas Fernández that shows how to use Istio to manage blue-green deployments in Kubernetes. This post goes into a lot of detail, and assumes that you have Istio already installed. Hopefully you are deploying automatically via a CI/CD pipeline, and if you aren’t already able to to do blue-green deployments, this is a great read. 🔁
1. [Tweet of the week](https://twitter.com/cra/status/1379492808663371780)

    Apache Mesos is moved to the Apache Attic. Formally, projects move to the attic at the end of life. Just in case anyone out there wasn’t ready to commit to Kubernetes yet. ☠

### [ << Prev ](kubelist-121.md) ------------- [ Next >> ](kubelist-123.md)