## [Kubelist Issue #92 for 2020-07-22](https://kubelist.com/issue/92)

#### Coordination Free Rollouts 🤹

> A modern application is complex, and starting may require some specific sequencing. Maybe the database needs to be running before the API, or you need to provision TLS certs before starting a web server. It’s possible (recommended even) to make code changes to remove these hard sequencing dependencies, and this isn’t always easy or possible. So this week we&#39;re looking at built-in Kubernetes features that can be used to start up a complex application or upgrade, without using external coordination utilities.

1. [Starting containers in order on Kubernetes with InitContainers](https://medium.com/@xcoulon/initializing-containers-in-order-with-kubernetes-18173b9cc222)

    This walkthrough from @xcoulon assumes that there’s an API that expects a Postgresql service to be available at startup. By default, this will become “eventually ready” by “crashlooping” until the dependencies are running, but that’s not an ideal solution. A cleaner deployment would be able to make guarantees that the pod restart count should be at 0 after a successful deployment. This post shows how to take an “eventually ready” application and use initContainers to make it start without crashlooping. 🔄
1. [Kubernetes: A Pod’s Life](https://www.openshift.com/blog/kubernetes-pods-life)

    A great overview of the lifecycle of a pod, written by @mhausenblas. If you are trying to build a Kubernetes application that doesn’t rely on external coordination, you’ll want to understand the pod lifecycle so you can identify where to check for dependencies and how to perform common initialization patterns. A pod contains more than an initContainer. Chances are you’ll learn something here, regardless of your Kubernetes experience. 🌱
1. [Kubernetes Patterns: The Init Container Pattern](https://www.magalix.com/blog/kubernetes-patterns-the-init-container-pattern)

    This post shows a couple of different use cases for initContainers including seeding a database and waiting for dependencies. Most common uses of initContainers is to perform any preliminary work and prepare a volume to be ready. The dependency check on init is an interesting option, just remember that the initContainer terminates, so this doesn’t replace liveness probes. Maybe Kuberntes pod readiness gates will support dependency readiness in the future? 
1. [PreStart and PostStop event hooks](https://github.com/kubernetes/kubernetes/issues/140)

    In issue #140 in the kubernetes/kubernetes repo, Brian Grant advocated for adding some additional lifecycle hooks into the pod lifecycle. While Kubernetes does support some lifecycle hooks today, we’ve included this original issue because the proposal contains some insights into the existing solutions that Kubernetes is replacing, and how they handle various issues related to rollout and lifecycle. Also note that the issue is still open to add PreStart and PostStop hooks because Kubernetes today supports PostStart and PreStop only. 🎣
1. [Pod readiness gates](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate)

    When building an application specifically for Kubernetes, startup sequencing can be extended with readinessGates. These are additional readiness checks that can be used to gate the overall readiness of the pod. As implemented, these readinessGates must be on the same pod, so they are useful for operators and other code that interacts with the Kubernetes API directly, to write fields into the Status subresource of the Kubernbetes object in question.
1. [Attach Handlers to Container Lifecycle Events](https://kubernetes.io/docs/tasks/configure-pod-container/attach-handler-lifecycle-event/)

    We’ll wrap it up today with a link to a walk through on the Kubernetes documentation. This is a pretty short tutorial, but it shows how to add postStart (and preStop) lifecycle hooks to a container in a pod. While postStart hooks aren’t a solution for all coordination challenges when deploying and starting an application, they are a simple way to solve some problems. 📑
1. [Tweet of the week](https://twitter.com/IstioMesh/status/1280864115335053314)

    Istio is in the Open Usage Commons while Linkerd and others are in the CNCF. Need to keep an eye on this.

### [ << Prev ](kubelist-91.md) ------------- [ Next >> ](kubelist-93.md)