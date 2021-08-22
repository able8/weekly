## [Kubelist Issue #58 for 2019-03-28](https://kubelist.com/issue/58)

#### Kubernetes Resources Part 3: Horizontal Pod Autoscaling + More

> Hello and welcome to the Horizontal Pod Autoscaling issue + Kubectl productivity!

1. [Boosting your kubectl productivity](https://learnk8s.io/blog/kubectl-productivity/)

    Wow, this is much more than just a simple article of how to type less - its really the most comprehensive overview on best practices of kubectlrey and a deep dive on how kubectl works. Highly recommended.
1. [AUTOSCALERS THE HARD WAY](https://www.engineyard.com/blog/autoscalers-the-hard-way)

    Control Theory applied to Kubernetes as it should be. Unlike other “controllers” the autoscalers are much more like classic feedback controllers in that their inputs and outputs are continuous values rather than heuristics-based decisions. Read (and watch) this to find out more.
1. [How we scale our kid-safe technology using Kubernetes](https://blog.superawesome.com/2019/02/26/how-we-scale-our-kid-safe-technology-using-auto-scaling-on-kubernetes/)

    “moderately in the morning, a little bit during the day, and a lot in the evening” – A classic use-case for horizontal scaling. This is a great recent article on a typical autoscaling journey we see from CPU HPA -> how does it work with cluster autoscaling -> using custom metrics.
1. [Kubernetes autoscaling with Istio metrics](https://medium.com/google-cloud/kubernetes-autoscaling-with-istio-metrics-76442253a45a)

    Now you can schedule based on metrics collected by Istio! If nothing else, this is a wonderful use case and example for the kube-metrics-adapter from Zalando. 
1. [Horizontal Pod Autoscaling by memory](https://koudingspawn.de/kubernetes-autoscaling/)

    What’s fantastic about this article is the fact that its horizontal pod autoscaling on memory with java so you get a little taste of what it looks like to autoscale on memory while using the JVM.
1. [Horizontal Pod Autoscaler Kubernetes Operator](https://banzaicloud.com/blog/k8s-hpa-operator/)

    A horizontal autoscaler operator based on annotations on the deployment. It’s so neat and elegant. Marking autoscaling config on a deployment like this seems to make a bunch of sense.
1. [Tweet of the Week](https://twitter.com/danielepolencic/status/1110853917166854145)

    Ask and you shall receive.

### [ << Prev ](kubelist-57.md) ------------- [ Next >> ](kubelist-59.md)