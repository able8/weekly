## [Kubelist Issue #59 for 2019-04-11](https://kubelist.com/issue/59)

#### Kubernetes Resources Part 4 - Horizontal Pod Autoscaling : External & Custom Metrics

> Hello and welcome to the Horizontal Pod Autoscaling issue again!

1. [Introducing Flagger](https://docs.flagger.app/)

    Canary Analysis: What is it even? We will let you read more here but the tl;dr is Does the canary pass your core KPIs or not? (In keeping with the theme of this issue Flagger does consume HPA objects so don’t worry)
1. [Horizontal Pod Autoscaling based on custom Istio metrics](https://banzaicloud.com/blog/k8s-hpa-prom-istio/)

    Horizontal Pod Autoscaling on Istio Metrics the Banzai way!
1. [How to build a Kubernetes Horizontal Pod Autoscaler using custom metrics](https://sysdig.com/blog/kubernetes-scaler/)

    Applications use resources based on their workload (mostly) - Why not scale on the workload? That's where custom metrics come in to play. Included is an awesome example of how to implement a provider for the custom metrics API!
1. [Ensure High Availability and Uptime With Kubernetes Horizontal Pod Autoscaler and Prometheus](https://www.weave.works/blog/kubernetes-horizontal-pod-autoscaler-and-prometheus)

    How to connect your prometheus to your Horizontal Pod Autoscaler in a nutshell. Thanks @weaveworks.
1. [Kubernetes HPA Autoscaling with Custom and External Metrics - Using GKE and Stackdriver Metrics](https://medium.com/uptime-99/kubernetes-hpa-autoscaling-with-custom-and-external-metrics-da7f41ff7846)

    Maybe you use Stackdriver for HPA? Maybe you want an example of how to use external metrics for HPA? Checkout @JessicaGreben's slick article on how.
1. [Kubernetes Operations: Prioritize Workload in Overcommitted Cluster](https://medium.com/microsoftazure/kubernetes-operations-prioritize-workload-in-overcommitted-clusters-8b9a60de98ec)

    Included here because not a lot of people talk about the PriorityClass resource which is important when your over committing clusters (especially when you have horizontal autoscaling in play).
1. [Tweet of the Week](https://twitter.com/rskiwah/status/1115913244378578944)

    The aha moment of HPA - love it!

### [ << Prev ](kubelist-58.md) ------------- [ Next >> ](kubelist-60.md)