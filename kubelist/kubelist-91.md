## [Kubelist Issue #91 for 2020-07-16](https://kubelist.com/issue/91)

#### Kubernetes Dashboards and kubectl Alternatives

> The Kubernetes Dashboard and kubectl are both powerful tools, but are not the only ways to view and interact with your cluster. There are some powerful alternatives to these components that provide unique views and functionality. Dashboards in this category aren’t intended as a replacement for full and proper monitoring, but are incredibly useful to get a high level view of what’s running in the cluster. Kubectl can drive a cluster with its composable commands, but some other CLIs have more interactive views that can save time with some tasks. This week, we are sharing some of the kubectl alternatives and lightweight Kubernetes dashboards that you can install.

1. [K9s](https://k9scli.io/)

    Starting out our list this week is k9s. This is an interactive alternative for some of the kubectl functionality. K9s provides a tree view of cluster resources to help you track down a pod ownership and the logs viewer is a lot easier than writing your own label selectors and passing them to kubectl. The top resource view is powerful when you want to look at where all of your node resources are allocated without leaving the command line. ⌨️
1. [Kubelive](https://github.com/ameerthehacker/kubelive)

    Kubelive is an interactive kubectl. Run “kubelive get pods” and select a namespace using your arrow keys instead of typing the namespace on the command line. Not every command from kubectl is implemented, notably the ability to get logs or exec into a pod. But the most common commands are implemented and kubelive could save you time when managing pods in multiple namespaces on a cluster.
1. [Kube Ops View](https://github.com/hjacobs/kube-ops-view)

    A useful utility from Henning Jacobs, kube-ops-view shows a read only, basic operational view of your cluster. It shows each node using color coded boxes, and gives you an easy way to see the number of pods along with the cpu/memory utilization per node. This high level view of where resources are being used could make it easy to track down unexpected resource utilization when debugging a problem. 🟥🟨🟦
1. [Kubernetes Resource Report](https://github.com/hjacobs/kube-resource-report)

    Instead of looking at how the cluster is running, this dashboard specialized in tracking down opportunities to save cost. 💸 The Kubernetes Resource Report looks for pod specs where the resource requests exceed the actual resource utilization (slack), estimating the cost savings and showing this to you. If your cluster is nearing capacity and you have CPU and Memory requests set on all pods (hopefully you do), the Kubernetes Resource Report could help you find opportunities to lower requests on some pods.
1. [Octant](https://octant.dev/docs/master/index.html)

    Octant is a local, web-based dashboard to help you understand what’s in your cluster. Octant has some features not found in other dashboards, such as the ability to add a local port forward into a pod. This project uses a plug-in model so that other developers and projects can add their own tabs and UI to the dashboard. Octant is a good tool for someone to get a better understanding of what is running in the cluster. 🖱️
1. [Weave Scope](https://github.com/weaveworks/scope)

    Scope is an open source web dashboard from Weave. Scope is a visual view showing you all of the workloads running in the cluster, allowing you to view logs, stop and restart containers, and even get a shell into a pod from the browser. Scope is a good choice for an SRE team who understands what’s running in the cluster, but needs to be able to easily troubleshoot an outage.
1. [Tweet of the week](https://twitter.com/wm/status/1283102961820856321)

    Seems pretty straightforward! 🤷

### [ << Prev ](kubelist-90.md) ------------- [ Next >> ](kubelist-92.md)