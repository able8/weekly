## [Kubelist Issue #60 for 2019-04-18](https://kubelist.com/issue/60)

#### Kubernetes Resources Part 5: Vertical Pod Autoscaling

> Today we are ending our series on Kubernetes Resources with a final part on Vertical Pod Autoscaling. We started the series looking at resources, what they were and how they were used and asked what autopilot for Kubernetes would look like? Are we there yet? How much effort will it be? 

1. [Rightsize Your Pods with Vertical Pod Autoscaling](https://www.youtube.com/watch?v=Y4vnYaqhS74)

    Still up to date overview of the beta Vertical Pod Autoscaler and the use case for it: the burden of right sizing.
1. [HPA and VPA: scale your K8s cluster on any metrics](https://skillsmatter.com/skillscasts/13439-hpa-and-vpa-scale-your-k8s-cluster-on-any-metrics-ecaterina-gamanji-conde-nast-international)

    VPA starting in at 11:17 => Excellent overview and demo of the vertical pod autoscaler. Also though not necessarily about how they work together there is some nice discussion about that in the Q&A.
1. [Introducing GKE Advanced— enhanced reliability, simplicity and scale for enterprise workloads](https://cloud.google.com/blog/products/containers-kubernetes/introducing-gke-advanced-enhanced-reliability-simplicity-and-scale-for-enterprise-workloads)

    Included here because there is an announcement for a Vertical Pod Autoscaler for GKE Advanced. Does anyone have the inside scoop on how this is different from the beta VPA available through vanilla GKE?
1. [Vertical pod autoscaler](https://banzaicloud.com/blog/vertical-pod-autoscaler/)

    A great step by step on setting up the VPA yourself, how to use the update policies ect. Plus its done for a tricky Java workload. Read more here.
1. [One Month of Kubernetes VPA](https://srcco.de/posts/one-month-of-kubernetes-vertical-pod-autoscaler-vpa.html)

    Really lovely visual case study on using the VPA, though it’s only one pod, you can really see what the operator's experience of using the VPA is (an where it fails).
1. [Is Helm used just for templating Kubernetes YAML files?](https://learnk8s.io/helm-templating-kubernetes-yaml/)

    An awesome overview of not just Helm, but also Kustomize, Helm3, Tiller and more. TIL: Why Helm is so popular, it’s not just templating but it’s release management and chart repository enablement is 👌
1. [Tweet of the Week](https://twitter.com/obeattie/status/973165748322492416)

    @tryexcept’s post (One Month of Kubernetes VPA) is a good toy example of this but still wondering if anyone has experience using VPA at scale.

### [ << Prev ](kubelist-59.md) ------------- [ Next >> ](kubelist-61.md)