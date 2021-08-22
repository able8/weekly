## [Kubelist Issue #114 for 2021-01-27](https://kubelist.com/issue/114)

#### Sysdig Falco 🦅

> On today’s episode of the <a href="https://kubelist.com/podcast">Kubelist Podcast</a> I had a conversation with <a href="https://twitter.com/danpopnyc">Dan (“Pop”)</a> from Sysdig. Pop is the Lead of Opensource Ecosystem and Community, specifically Falco, but has been at Sysdig in various roles before finding his current role. He’s extremely knowledgeable and passionate about helping everyone run a secure cluster. If you aren’t already subscribed to his podcast, subscribe to the <a href="https://popcast-d9f7b6dc.simplecast.com">POPCAST</a> and get ready for season 2.

1. [{Join POP Falco.org}](https://www.cncf.io/blog/2020/12/14/join-pop-falco-org/)

    A really good post written by Dan "POP" Panandrea to explain his mission and goals leading open source community and ecosystem at Sysdig. If you don't know Pop, you should definitely subscribe to the POPCast for one of the best shows about the cloud native ecosystem. The story in this post provides a lot of the framework, but then Pop shares how he's approaching his new role, and how you can help. If you run Falco, check this out for ways to be more involved. If you don't run Falco, well, you should.
1. [An Introduction to Kubernetes Security using Falco](https://falco.org/blog/intro-k8s-security-monitoring/)

    This is a quick overview and walkthrough of installing Falco to a Kubernetes cluster, and using the default rules to detect unusual behavior. In this example, we get to see how Falco detects and reports when a new terminal is spawned from inside a running Pod on a Kubernetes cluster. This attack vector should make sense, and knowing when a new shell is spawn unexpectedly is a great way to start monitoring the runtime of your cluster. 👀
1. [Introducing Falco: Open source behavioral security from Sysdig](https://sysdig.com/blog/sysdig-falco/)

    This one is from 2016, it’s the original blog post announcing the Falco project. While this post is almost 5 years old, it’s pretty ambitious and is a really good read to see how well-thought and planned the approach to Falco has been. The team at Sysdig saw a problem and started to build the solution a while ago, but hasn’t stopped building and iterating since. When you read this and think about the efforts put into the project to date, it shows how mature this solution is. 🎂
1. [Monitoring AS EKS audit logs with Falco](https://xebia.com/blog/monitoring-aws-eks-audit-logs-with-falco/)

    While this isn’t a classic example of detecting abnormal or unexpected behavior in the Linux kernel, here’s a Falco use case that shows the flexibility of the project. Falco is a rules engine, and it can detect other unexpected events, such as the event stream coming from the Kubernetes API server. This walkthrough shows how to look at the events coming out of an Amazon EKS cluster, and set up a pipeline to get them analyzed by Falco for irregularities and interesting events. 📑
1. [Falco Sidekick](https://github.com/falcosecurity/falcosidekick)

    This was a project mentioned several times in the Kubelist Podcast episode today. While Falco does a great job of analyzing the stream of events happening in the kernel, and applying rules to detect the most interesting, relevant, or unexpected events in the stream, Sidekick is a project that will help you connect the output of Falco into other integrated systems. Just a quick look at the list of supported outputs from the README shows how easy it will be to connect to a supported system in your organization. 🦸‍
1. [Falco + Kubeless = a Kubernetes Response Engine](https://falco.org/blog/falcosidekick-kubeless/)

    This is a basic walkthrough of connecting Falco to a “serverless” runtime to respond and react to events in a Kubernetes cluster. Once again, this example shows how to respond to a new shell being spawned, but all of the Falco events are supported here. What’s really unique here is that Kubeless functions are the output of Falco in a cluster. If Falco detects abnormal behavior, Falco Sidekick can run a remediation action in a serverless system (Kubeless in this case) to take appropriate action. The example deletes the pod that triggered the event, but once you get the Kubeless function called, anything is possible. 🚒
1. [Dancing with Kubernetes](https://techblog.fexcofts.com/2020/07/09/dancing-with-kubernetes-part-ii/)

    A fun article, walking through deploying Falco to detect runtime threats in a Kubernetes cluster. This is both well written and a good explanation of why the defense in depth approach requires something like Falco operating at runtime. You can take all of the preventative measures possible to protect the supply chain (and you should), but you’ll never be confident in your security posture until you also have runtime detection in place. 🕺
1. [Tweet of the week](https://twitter.com/cra/status/1352307918834642945)

    CNCF Member Acquisitions, as reported by @cra, CTO of the CNCF. This slide really shows the velocity and adoption of everything cloud-native.

### [ << Prev ](kubelist-113.md) ------------- [ Next >> ](kubelist-115.md)