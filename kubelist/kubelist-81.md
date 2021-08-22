## [Kubelist Issue #81 for 2020-04-22](https://kubelist.com/issue/81)

#### Managing Costs 💸

> There are a lot of guides that talk about managing, reporting on, and lowering the costs of running Kubernetes clusters. This week, we share a set of links that all will have you thinking about optimizing the amount you are spending to run Kubernetes.

1. [The Definitive Guide to Lowering Cloud Costs of Kubernetes Clusters](https://managedkube.com/kubernetes/cloud/costs/2019/01/11/lowering-kubernetes-costs.html)

    Before investing in a lot of new tools, this article shows how to make sure you are taking advantage of the low hanging fruit — specifically what’s built into Kubernetes and should be easy to implement. Optimizing for cluster autoscaling sounds fun, but setting appropriately-sized requests and limits on pods is where you need to start or you’ll just be scaling away your dollars.
1. [The True Cost of Running Kubernetes in the Cloud](https://convox.com/blog/cost-of-running-k8s)

    Here is a detailed comparison of various managed Kubernetes offerings, with a focus on the costs of running each. If you are still in the position of being able to pick where to run K8s, read this before making any decisions.
1. [Kubernetes cost monitoring: approaches & best practices](https://medium.com/kubecost/kubernetes-cost-monitoring-approaches-best-practices-846e8e81ae69)

    This is a good introductory explanation of the various approaches of managing Kubernetes costs (and cloud spend in general). Whether you are running a small cluster for your startup, or a multi-tenant cluster for a large organization, this post will help define some terms (i.e. showback vs chargeback) that will be useful on your cost management journey.
1. [Saving Cloud Costs with Kubernetes on AWS](https://srcco.de/posts/saving-cloud-costs-kubernetes-aws.html)

    Henning Jacobs wrote a post that contains a lot of good recommendations and pointers to (free) tools that can help manage costs. This post and the utilities mentioned have a focus on cleaning up unused resources and minimizing idle nodes — both really important steps in managing costs.
1. [Kubernetes infrastructure cost optimization](https://platform9.com/blog/kubernetes-infrastructure-cost-optimization/)

    A well informed post that describes various techniques to manage Kubernetes costs, including using Artificial Intelligence and Machine Learning to get more predictive. Bringing in ML and AI shouldn’t be the first step to cost optimization, but once you’ve covered the basics, it could be useful to advance to the tools mentioned here.
1. [Cost optimization for Kubernetes on AWS](https://aws.amazon.com/blogs/containers/cost-optimization-for-kubernetes-on-aws/)

    This post is a bit AWS-specific (it is on their blog), but many of us run EKS clusters and have a large bill each month. There are some good tips in this post about autoscaling and spot instances. If nothing else, read this one to start speaking the language of cost savings on AWS.
1. [Tweet of the week](https://twitter.com/stephenaugustus/status/1252413369702744064)

    Some changes proposed to the Kubernetes 1.19 release process:

### [ << Prev ](kubelist-80.md) ------------- [ Next >> ](kubelist-82.md)