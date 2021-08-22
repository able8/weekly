## [Kubelist Issue #10 for 2018-04-19](https://kubelist.com/issue/10)

#### 🗹 Operating the Kube 🗹

> Welcome to the operating Kubernetes edition of kubelist. If you thought operating the Kube was fun wait till you read up on these operator goodies.

1. [Introducing kaniko: Build container images in Kubernetes and Google Container Builder without privileges](https://cloudplatform.googleblog.com/2018/04/introducing-kaniko-Build-container-images-in-Kubernetes-and-Google-Container-Builder-even-without-root-access.html)

    Typically, you build your docker containers with root privileges. How can we mitigate these risks especially building containers on the 🗹? GCP thinks kainko can save the day but as @jessfraz points out kainko’s design might not provide you with what your looking for. Checkout a more indepth conversation on what it might take to do secure builds on k8s here
1. [Building a Kubernetes Ingress Controller](https://eyskens.me/building-a-kubernetes-ingress-controller/)

    There is a lot going on in networking the 🗹. Ingresses, services, network policies, affinities and on and on. One of the biggest things that the kubelist editors struggled with was how Ingresses route to services and the DNS sugar on top. Check out Maartje’s blog post on figuring out how these work by building your own Ingress controller, sweet!
1. [Deploying Hyperledger Fabric on Kubernetes](https://opensource.com/article/18/4/deploying-hyperledger-fabric-kubernetes)

    The kubelist editors LOVE reading about how folks set up their 🗹. We're not huge into blockchain but this article gives a great overview of how any organization can take their infra to the next level with Kubernetes. ☸️ P.S. there is a little Helm sprinkled in there too! ☸️
1. [Log aggregation with ElasticSearch, Fluentd and Kibana stack on ARM64 Kubernetes cluster](https://medium.com/@carlosedp/log-aggregation-with-elasticsearch-fluentd-and-kibana-stack-on-arm64-kubernetes-cluster-516fb64025f9)

    Deploying and maintaining ElasticSearch is surprisingly hard stuff. It's a heavy JVM app, there is tons of networking and storage config going on there, and its topology means you need to ensure the worker/master distribution is just right 👌. Kudos to Carlos for his sweet setup of his ?EKF? Stack on the 🗹
1. [Exploring container security: Protecting and defending your Kubernetes Engine network](https://cloudplatform.googleblog.com/2018/04/exploring-container-security-protecting-and-Defending-your-Kubernetes-Engine-network.html)

    Most of this article contains tactics folks are already implementing on their own (VPC segmentation, HTTPS, etc). We suppose this post by Google serves as a good review of what you should do when you set up your infra when not using GCP? In all seriousness though, the stuff on Encryption in Transit and Attack Mitigation is dope.
1. [Tweet of the week](https://twitter.com/ipedrazas/status/985895482223362049)

    A little bit about the kubelist editors: One of us had no idea Pods refered to a pod of whales until last week. They always thought it had to do with peas?

### [ << Prev ](kubelist-9.md) ------------- [ Next >> ](kubelist-11.md)