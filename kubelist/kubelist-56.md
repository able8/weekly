## [Kubelist Issue #56 for 2019-03-14](https://kubelist.com/issue/56)

#### New Series: K8s Resources: Part 1 - Resource Management

> We are starting a new series at Kubelist this week: Kubernetes Resources and Autoscaling. We are going to be taking a little break from aggregating new content and curate some lists based on things we think are important to know for Kubernetes. (If you&#39;re really hungry for the newest stuff check out: <a href="https://kube.news/">https://kube.news/</a> or <a href="https://twitter.com/kubeweekly?lang=en">@kubeweekly</a>).

1. [Container resource consumption—too important to ignore](https://hackernoon.com/container-resource-consumption-too-important-to-ignore-7484609a3bb7)

    A classic case made for the use of a dynamic resource allocator within Kubernetes and the complications therein. Why should you have to set container resources manually? Checkout the reference implementation at: https://github.com/openshift-demos/resorcerer 
1. [Everything You Ever Wanted To Know About Resource Scheduling](https://www.youtube.com/watch?v=nWGkvrIPqJ4)

    Isolation, Right sizing, Utilization: The core components of kubernetes resource management. A stellar talk on the dangers of full utilization, stability over performance and the insights brought from Borg.
1. [How Full is My Cluster? Capacity Management and Monitoring on OpenShift](https://blog.openshift.com/full-cluster-capacity-management-monitoring-openshift/)

    Technically it’s OpenShift but it mostly applies to Vanilla Kubernetes. This series is really stellar because it goes into a lot of depth about how resources are set and managed by Kubernetes. The series could be called: autoscaling from first principles.
1. [Understanding Kubernetes Resources](https://www.noqcks.io/notes/2018/02/03/understanding-kubernetes-resources/)

    ALWAYS set limits and cpu, got it? Readable intro to how kubernetes manages resources - how apps get scheduled and how they get killed.
1. [How Kubernetes Resource Classes Promise to Change the Landscape for New Workloads](https://thenewstack.io/kubernetes-resource-classes-promise-change-landscape-new-workloads/)

    BYOR - Why resource classes will(do) enable bring your own resources and what the community is doing to make it happen.
1. [5 Ways to Manage Your Kubernetes Resource Usage](https://www.replex.io/blog/5-ways-to-manage-your-kubernetes-resource-usage)

    Some common theme in Kubernetes resource management: Cost creeps in, Over provisioning creeps in, Performance degradation creeps in. Learn how to avoid the creep with resource management primitives. 
1. [Kubernetes Container Resource Requirements — Part 1: Memory and Part 2: CPU](https://medium.com/hotels-com-technology/kubernetes-container-resource-requirements-part-1-memory-a9fbe02c8a5f)

    If you want to know how the kubernetes layer of resource management translates down to the docker layer which translates down to the kernel read here. 
1. [Optimizing Kubernetes resource allocation in production](https://opensource.com/article/18/12/optimizing-kubernetes-resource-allocation-production?sc_cid=70160000001273HAAQ)

    We end this issue with testing the k8s resources “limits” with load tests and telling us the results.
1. [Tweet of the Week](https://twitter.com/learnk8s/status/1104774251570184192)

    Wanna know how much resources you are using? Even at the cost level? Checkout out Purser!

### [ << Prev ](kubelist-55.md) ------------- [ Next >> ](kubelist-57.md)