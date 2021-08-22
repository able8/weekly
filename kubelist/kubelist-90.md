## [Kubelist Issue #90 for 2020-07-08](https://kubelist.com/issue/90)

#### Quality of Service, Limit Ranges and more 

> Kubernetes makes it easier to run different services in a single, shared cluster. This creates an opportunity for any service to disrupt and create outages in any other service or the entire cluster. When you have improperly configured services in a cluster, they will be subject to the various scheduling guarantees that Kubernetes offers. But Kubernetes does offer a solution here with built-in QoS Classes (Quality of Service) and LimitRanges. These features should be enabled and enforced in every cluster to keep your cluster healthy. This week, we’ve shared a few blog posts to help you understand these Kubernetes features and how to use them.

1. [What are Quality of Service (QoS) Classes in Kubernetes](https://medium.com/google-cloud/quality-of-service-class-qos-in-kubernetes-bb76a89eb2c6)

    An introduction to the scheduling (placement and termination) guarantees in Kubernetes implements with various pod configurations. Once you understand the QoS classes identified in the bottom part of this article, you’ll be able create pod definitions that will work as expected in your cluster. Even if you know about resource requests and limits, check out how these values map to QoS classes. ⏳
1. [Everything you Need to Know about Kubernetes Quality of Service (QoS) Classes](https://www.replex.io/blog/everything-you-need-to-know-about-kubernetes-quality-of-service-qos-classes)

    Another post explaining the various QoS classes of pods, but written as a FAQ. Now that you know how request and resource definitions translate into QoS classes, this post will help you understand what to expect for each class. In this post, the same set of questions is answered for each of the three QoS classes to help provide a comparison. One takeaway from this article is that any pod where all containers have CPU/memory requests AND limits defined will be the last evicted by the scheduler. By setting a request but not a limit, you can guarantee placement and a minimum allocatable but allow bursting to capacity of the node. The cost for this bursting capability is earlier eviction when the node needs resources.
1. [Configure Quality of Service for Pods](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/)

    The official Kubernetes documentation has a good hands-on tutorial to learn the various QoS classes for pods. The documentation here walks through creating several different pods with different configurations. It then shows how you can verify the QoS that was assigned to the pod, based on the configuration supplied. 📦
1. [Kubernetes Resources Management - QoS, Quota, and LimitRange](https://darumatic.com/blog/k8s_resources_manage_quota)

    Another post that maps memory/cpu requests/limits to QoS and explains some scheduling decisions made by Kubernetes. We wanted to include this post because of the very practical tips, such as the statement that Kubernetes will throttle a pod exceeding it’s CPU limit, but terminate a pod that exceeds its memory limit. This post also starts to discuss how LimitRanges can be used to ensure or even set the QoS on a pod in some cases. ⚡
1. [Namespace LimitRange](https://sdbrett.com/BrettsITBlog/2019/11/kubernetes-namespace-limitrange/)

    LimitRanges are a built-in way to set and enforce both cpu and memory requests and limits at the namespace level. Min and max values work as expected: if a pod configuration includes a request or limit that is out of the min/max bounds configured in the LimitRange, scheduling will be disabled. Interestingly, LimitRanges also include a default and defaultRequest, which will apply limit and request values if not otherwise set. You’ll need to move to MutatingWebhooks or Gatekeeper to modify pod config and change set values though. 🗜
1. [Reserve Compute Resources for System Daemons](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/)

    If you are running your own Kubernetes cluster, there are some useful configuration flags and suggestions here. You can adjust the eviction policies, change the reserved limits for the API server, and more. If you are seeing the Kubernetes API server starved of resources, consider changing the parameters of the kubelet and adjusting the scheduling parameters a little. 🚩
1. [Tweet of the week](https://twitter.com/jrrickard/status/1280519718458163206)

    Our answer: it doesn’t matter if you’ve set resource requests and limits on all containers properly.

### [ << Prev ](kubelist-89.md) ------------- [ Next >> ](kubelist-91.md)