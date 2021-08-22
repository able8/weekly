## [Kubelist Issue #76 for 2019-08-15](https://kubelist.com/issue/76)

#### Kubernetes Horror Stories 👻

> This week we’re looking into Kubernetes failures and incidents. As you might imagine, these issues can have a nasty impact for users, and where recovery plays the most critical role. 

1. [First off, we have Spotify and the Accidental Kube Cluster Delete](https://www.youtube.com/watch?v=ix0Tw8uinWs)

    David Xia, Infrastructure Engineer at Spotify goes in depth with his story on how Spotify accidentally deleted all its kube clusters with no user impact. Since the postmortem, Spotify has increased security and recovery tactics.
1. [The Story of Grafana and Pod Priority](https://grafana.com/blog/2019/07/24/how-a-production-outage-was-caused-using-kubernetes-pod-priorities/)

    Back in July 2019, Grafana’s customers experienced a production outage, caused by Grafana’s pod priority system. With their recovery around 30 minutes long, it still had an impact for large kube pods. 
1. [Make Room for Memory!](https://www.bluematador.com/blog/post-mortem-kubernetes-node-oom)

    Blue Matador had recently written a blog post outlining some issues with nodes in Kubernetes clusters running out of memory space. This case study covers the incident, the fix, and the take-away points. 
1. [When Pods and Nodes are Crashing](https://updates.moonlightwork.com/outage-post-mortem-87370)

    Moonlight’s website dealt with multiple issues including kernel panics on nodes and CPU resources used at 100% by pods. Kernel panics are interesting problems to deal with in Kubernetes. 
1. [The Breakage and Recovery of a K8s Cluster](https://medium.com/civis-analytics/https-medium-com-civis-analytics-breaking-kubernetes-how-we-broke-and-fixed-our-k8s-cluster-adfa6fbade61)

    This article explores how stability issues led to a script overloaded cluster. Monitoring comes into a strong recovery point for this particular incident. 
1. [Moving to Kubernetes: Where it Could go Wrong](https://engineering.saltside.se/our-failure-migrating-to-kubernetes-25c28e6dd604)

    This story had trouble migrating one application tier to K8s. Read more on this above!
1. [Tweet of the Week](https://twitter.com/ruggerotonelli/status/1159844538405597184)

    Interested in reading on more Kubernetes failures? Check out this tweet. 

### [ << Prev ](kubelist-75.md) ------------- [ Next >> ](kubelist-77.md)