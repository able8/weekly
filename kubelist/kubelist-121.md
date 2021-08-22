## [Kubelist Issue #121 for 2021-03-31](https://kubelist.com/issue/121)

#### New Sandbox Projects 🏜

> Recently, the CNCF TOC (Technical Oversight Committee) changed the requirements and application process in order to make it easier to get projects accepted into the Sandbox. This has been a welcome change as we’ve seen more projects than ever included earlier in their lifecycle. In this week’s issue, we are pointing to some of the new projects included in the Sandbox so far in 2021. This isn’t an exhaustive list of what’s been added, so check out the <a href="https://www.cncf.io/sandbox-projects/">CNCF Site</a> for the latest list (and subscribe to the <a href="https://kubelist.com/podcast/">Kubelist Podcast</a> 😉).

1. [k8gb](https://www.k8gb.io/docs/)

    A Kubernetes Global Balancer (added in March 2021) that can route traffic to pods in various clusters. When we think about going to multi-cloud or another highly available method of deploying a service, we are often limited to a single Kubernetes cluster. So when we want to go multi-cluster, we have to think about how to bring external load balancers above the cluster to route traffic to a LoadBalancer or Ingress. This is a neat idea to make it all Kubernetes-native. ⚖
1. [Ingraind](https://ingraind.org)

    Also added in March: InGRAINd is an eBPF-powered monitoring tool (built in Rust) to detect anomalies. The ingraind agent can monitor activity on the system (or container for most of us), and either report that an event occurred or take an action. We love seeing open source security tools, and ingraind looks great. This is a new project, and we are curious to see how it compares to other tools in the CNCF and ecosystem. Let’s get someone from ingraind on the podcast soon to learn more! 👓
1. [Kuberhealthy](https://github.com/Comcast/kuberhealthy)

    Also added in March 2021, this is a synthetic monitoring application built for Kubernetes, originally developed at Comcast. If it’s a new term, synthetic monitoring is an active process that attempts to run a specific action on a regular basis, and report the result. Maybe you want to check that an application-specific functionality exists (such as completing a checkout) or some action that’s not frequent enough to monitor, or easy to end-to-end monitor. Add a khcheck (Kuberhealth check) to the spec and deploy it, and the Kuberhealth operator will take over and report. ⛑
1. [Trickster](https://github.com/tricksterproxy/trickster/)

    Added in March, Trickster defines themselves as “an HTTP reverse proxy/cache for http applications and a dashboard query accelerator for time series databases.” This looks incredibly useful for running time series databases (Prometheus, InfluxDB, and more) in production at scale. As we see more and more production clusters running stateful services inside the cluster, Trickster could be a key project to keep it available. 🃏
1. [Curifense](https://www.curiefense.io)

    A cloud-native, Envoy-based web application firewall. Security tools! We are pretty impressed with the maturity of this project, given the early stage of the project in the landscape. We’ve been playing around with Curifense lately, and are planning on a podcast episode soon to dig into this project. 🦏
1. [Athenz](https://www.athenz.io)

    More security tools! Included in the last batch, this is a cool project that enables strong identity in a cluster based on Zero Trust concepts. Every cluster should be using strong identity and mTLS inside the cluster, and there are a few great options available to get started. Athenz is a welcome addition to add more options, and we can’t wait to learn more about the differences between Athenz and other projects, and how it integrates into service meshes. 🏛
1. [Tweet of the week](https://twitter.com/chronicleflask/status/1375848046953701383)

    The obvious choice here was one of the many (oh so many) memes comparing the ship stuck in the Suez Canal with the complexity of running Kubernetes and containers. But we wanted to give you something just light-hearted and funny to watch this week.

### [ << Prev ](kubelist-120.md) ------------- [ Next >> ](kubelist-122.md)