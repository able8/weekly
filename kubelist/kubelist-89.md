## [Kubelist Issue #89 for 2020-07-01](https://kubelist.com/issue/89)

#### Chaos 👻

> No Kubernetes migration is complete without a chaos engineering story. To successfully run a production system on Kubernetes without unexpected outages, everything should be horizontally scalable and resilient. In other words, expect restarts and other “problems”, code and design around them, and you’ll be successful. Chaos Engineering is the practice of running experiments on your cluster to see how well everything behaves. This week, we are looking at some blog posts and some tools to help you get started or to advance your Chaos Engineering practice.

1. [Introduction to Chaos Engineering](https://www.chaosmesh.com/chaos-engineering/)

    ChaosMesh has one of the best introductions to the concepts of chaos engineering we’ve found. If you haven’t even started reading about and understanding the practice of chaos engineering, start here for an explanation of the terms you will see in products and projects that implement cloud-native chaos engineering. The ChaosMesh product (from the same team who wrote the post) is capable of running some of the most sophisticated attacks possible, including complex experiments running against specific nodes, AWS accounts, and K8s pods, all at the same time.
1. [ChaosHub](https://hub.litmuschaos.io/)

    Litmus is a Chaos Engineering CNCF Sandbox project sponsored by MayaData (makers of OpenEBS). One of the compelling features of the Litmus project is ChaosHub, which is a place to find existing Chaos experiments to run. The product is very capable, and the centralized, public “hub” where you can find mature experiments to run against various software in Kubernetes might be a good reason to choose this platform as part of any chaos story. There’s not a lot in the ChaosHub yet, but it was just accepted as a CNCF project recently, so keep an eye on this one.
1. [Gremlin: The Chaos Engineering Platform for Kubernetes](https://www.gremlin.com/kubernetes/)

    Gremlin is one of the commercial platforms that is well known in the chaos engineering ecosystem. They’ve created a Kubernetes-native solution that retains all of the capabilities of their core platform, but can target Kubernetes specific objects also, such as pods and deployments. While not open source, this is a commercial offering worth checking out.
1. [KubeInvaders - Gamified Chaos Engineering Tool for Kubernetes](https://kubernetes.io/blog/2020/01/22/kubeinvaders-gamified-chaos-engineering-tool-for-kubernetes/)

    A fun introduction to the concepts of chaos engineering by playing a game. Your real chaos engineering story on Kubernetes will be automated, scheduled and much more elaborate, but if you want to experience the feeling of randomness and helplessness that you should be comfortable with, give this a try. Load the pods and destroy one, and see what breaks. Ok, don’t actually try this in production until you have some confidence that everything will reschedule and you won’t create an outage. 👾
1. [ChaosKube](https://github.com/linki/chaoskube)

    ChaosKube doesn’t let you schedule or plan any sophisticated attacks, it simply kills a pod every 10 minutes, all day long. We’ve included ChaosKube in the list because it’s local to your cluster, super easy to set up, and pods are already restarting in your cluster. This just increases the frequency of the restarts. If you don’t have anything else running chaos experiments in your cluster today, start here. 💀
1. [Chaos Mesh - Your Chaos Engineering Solution for System Resiliency on Kubernetes](https://pingcap.com/blog/chaos-mesh-your-chaos-engineering-solution-for-system-resiliency-on-kubernetes)

    Chaos Mesh is a Kubernetes Operator that implements a controller capabilities of sophisticated attacks on your Kubernetes cluster, not to be confused with the ChaosMesh product mentioned in the first link of this issue. While it can kill pods, Chaos Mesh can simulate I/O slowdowns and more. Sponsored by the developers of TiDB, this is another example of a chaos engineering tool created by a software vendor who normally focuses on persistent storage. Maybe storage is just hard and needs resiliency. ⚙
1. [Tweet of the week](https://twitter.com/devgerred/status/1277964998015254529)

    We agree. While sandbox projects might be pretty early and don’t have a lot of adoption, watch this space for trends and early insights into what’s going to be common in the future.

### [ << Prev ](kubelist-88.md) ------------- [ Next >> ](kubelist-90.md)