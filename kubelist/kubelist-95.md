## [Kubelist Issue #95 for 2020-08-12](https://kubelist.com/issue/95)

#### Open Service Mesh, Kubernetes Scheduler and more

> This week’s issue is full of technical blog posts and discussion of a couple of projects that are worth checking out. We can’t leave Open Service Mesh from Microsoft out, so we’ll start there. But we are really excited about some of the work happening to make the Kubernetes Scheduler more extensible and powerful. When reading these links about the scheduler, imagine how flexible a Kubernetes cluster of the future will be. 🔮

1. [Introducing Open Service Mesh](https://github.com/openservicemesh/osm/blob/main/DESIGN.md)

    Last week, Microsoft announced the Open Service Mesh project and committed to donating it to the CNCF. This new service mesh does implement the Service Mesh Interface (SMI) and is an interesting new addition to the ecosystem. 🕸
1. [Introducing PodTopologySpread](https://kubernetes.io/blog/2020/05/introducing-podtopologyspread/)

    PodTopologySpread is a new scheduler plugin, promoted to Beta in 1.18. If your PodSpecs are cluttered with Affinity and AntiAffinity rules to nudge the Kubernetes scheduler to selecting the right distribution of your pods across nodes, check out this new feature. It’s a first class plugin now, and allows more flexibility and power than Affinity and AntiAffinity. 📆
1. [Scheduling Profiles](https://kubernetes.io/docs/reference/scheduling/profiles/)

    Speaking of scheduler plugins, an alpha feature of 1.18 is Scheduling Profiles; allowing a cluster operator to configure which scheduler plugins are executed at each scheduler phase. The Kubernetes scheduler is getting quite powerful, thanks to the efforts over at #sig-scheduling.
1. [Anatomy of “kubernetes.default”](https://networkop.co.uk/post/2020-06-kubernetes-default/)

    Have you ever wondered what the kubernetes service was providing in your cluster? Also, why is this service recreated when you delete it? This article digs in and explains how it’s created and what the purpose of this built-in service is. 🧩
1. [Antrea: The Ubiquitous CNI](https://neonmirrors.net/post/2020-06/antrea-the-ubiquitous-cni/)

    Antrea is one of the newest additions to the Container Network Interface (CNI) ecosystem. Kubernetes clusters use a CNI provider to deliver overlay networking, network-level encryption, network policy enforcement and more. Antrea has a unique advantage being built on top of Open vSwitch, which was designed for high performance environments and is pre-installed on most modern Linux distributions these days. The promise of removing kube-proxy to leverage functionality from the CNI provider is appealing.
1. [Minimum Viable Kubernetes](https://eevans.co/blog/minimum-viable-kubernetes/)

    We are familiar with installing Kubernetes “the hard way” thanks to Kelsey Hightower. The hard way is a bit like learning how a computer works by buying all of the parts on PCPartPicker and assembling a computer. This is an incredibly useful exercise if you want to understand everything about how the various parts of Kubernetes work together. In the spirit of doing things the hard way for learning, but optimizing for time, MVK is useful to get started understanding what’s going on in your cluster. 🏗
1. [Tweet of the week](https://twitter.com/QuinnyPig/status/1291515938500616192)

    Some ancient Sysadmin wisdom from @QuinnyPig 🐂

### [ << Prev ](kubelist-94.md) ------------- [ Next >> ](kubelist-96.md)