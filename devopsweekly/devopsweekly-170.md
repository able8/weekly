## [DevOps'ish 170](https://devopsish.com/170) - Sun Mar 8, 2020

Everything is canceled, postponed, or going virtual. Coronavirus (COVID-19) has already made drastic impacts across the planet. Particularly this week, when several companies restricted travel forcing conferences to do something if they hadn’t already. KubeCon EU is getting pushed back to July or August. I canceled my plans earlier in the week because, with everything up in the air, there’s no telling I’ll be available to participate whenever it does eventually get rescheduled. Since a lot of us are probably stuck at home, I’d like to remind you, I’m always looking for opportunities to sit down and talk to folks from my office. I’ve got the setup for it and have a few upgrades in mind. If you have an idea for a conversation instead of a Meetup, panel, or talk, feel free to reach out. I’m looking for collaborative opportunities.

### People

1. [Don’t Hide Bad News in Times of Crisis](https://hbr.org/2020/03/dont-hide-bad-news-in-times-of-crisis)

    Dumbest thing you can do in a crisis is double down on existing decisions. The next dumbest thing is to double down without being transparent about what work is in progress to reassess the situation.
1. [Coronavirus, US: Naming victims would violate HIPAA, be dangerous](https://www.usatoday.com/story/news/nation/2020/03/06/coronavirus-us-naming-victims-would-violate-hipaa-dangerous/4964498002/)

    Imagine how appalled I was when a coworker told me they had their town’s first coronavirus victim identified on the local news. It was with consent as part of a public health announcement but, that has to be a pretty dire situation (like coronavirus).
1. [The Diffblue Developer Survey](https://www.diffblue.com/Education/research_papers/2019-diffblue-developer-survey/?utm_source=newsletter&utm_medium=email&utm_campaign=DevOpsIsh)

    What’s wrong with software speed, quality, and cost? Oxford-based AI company Diffblue surveyed 300 developers to find out what the people who actually work with code have to say about testing, their pain points, and what’s holding back DevOps in their organizations. SPONSORED
1. [Women in Ops](https://www.redhat.com/sysadmin/women-ops)

    With the push towards diversity in tech, why are women under-represented in operations?
1. [GitHub hires top DevOps researcher who founded Google Cloud-owned DORA](https://www.businessinsider.com/github-hires-nicole-forsgren-devops-researcher-google-cloud-dora-2020-3)

    I think having Business Insider covering you being poached is a new career goal (that is a very high bar). Congratulations to Nicole Forsgren on her new role!
1. [Should open source be ethical?](https://www.infoworld.com/article/3530300/should-open-source-be-ethical.html)

    “At present, software licenses that prohibit socially harmful or unethical uses cannot be considered open source. Should we change that?”
### Process

1. [Google GKE to introduce management fees by June 6th](https://www.reddit.com/r/kubernetes/comments/fdgblk/google_gke_to_introduce_management_fees_by_june/)

    “If the main value of GKE over <other> is $73, you should totally use <other>.” But, your GKE bill is going up.
1. [Unfixable vulnerability in Intel chipsets lets attackers decrypt data](https://www.scmagazineuk.com/unfixable-vulnerability-intel-chipsets-lets-attackers-decrypt-data/article/1676112)

    “A vulnerability in Intel CSME - CVE-2019-0090 - can enable a local attacker to extract the chipset key stored on the PCH microchip and obtain access to data encrypted with the key. Discoverers Positive Technologies say that it is impossible to detect such a key breach.”
1. [Most of Linux distros affected by a critical RCE in PPP Daemon flaw](https://securityaffairs.co/wordpress/99043/hacking/linux-rce-ppp-daemon-flaw.html)

    “The flaw, tracked as CVE-2020-8597, was discovered by the expert Ilja Van Sprundel from IOActive, it is a stack buffer overflow issue that is caused by a logical error in the Extensible Authentication Protocol (EAP) packet parser of the pppd software.” I feel like Obi Won Kenobi on this one. EAP? That’s a protocol I haven’t heard in a long time.
1. [Join us for the inaugural O’Reilly Infrastructure & Ops ConferenceInfrastructure & Ops ConferenceRegister today](https://conferences.oreilly.com/infrastructure-ops/io-ca?utm_medium=paid+media&utm_source=devopsish&utm_campaign=ioca20&utm_content=paid+devopsish+march+9)

    From legacy to leading edge, learn how to modernize your systems, train with industry experts, and network with peers at the first  in Santa Clara, June 15-18. Save up to $700 during Best Price, which ends on March 20. . SPONSORED
1. [2019 CNCF Survey results are here](https://www.cncf.io/blog/2020/03/04/2019-cncf-survey-results-are-here-deployments-are-growing-in-size-and-speed-as-cloud-native-adoption-becomes-mainstream/)

    Deployments are growing in size and speed as cloud native adoption becomes mainstream
1. [SLO Pitfalls](https://www.linkedin.com/pulse/slo-pitfalls-femi-agbabiaka/)

    “[S]etting and forgetting is not as simple as it may seem. I wanted to walk through a few common pitfalls I’ve seen and discuss how your organizations can avoid them.”
1. [A Survey of Istio’s Network Security Features](https://research.nccgroup.com/2020/03/04/a-survey-of-istios-network-security-features/)

    “In this blog post we used lab-based examples to illustrate common misconceptions and pitfalls encountered when using Istio to limit workloads’ network traffic”
### Tools

1. [Kubernetes operators: Embedding operational expertise side by side with containerized applications](https://www.redhat.com/sysadmin/kubernetes-operators)

    Kubernetes isn’t complex, your business problem is. Learn how operators make it easy to run complex software at scale. “I often describe the Operator pattern as deploying a robot sysadmin next to the containerized application.”
1. [Istio in 2020 - Following the Trade Winds](https://istio.io/blog/2020/tradewinds-2020/)

    The who, what, where, and how of Istio plans for 2020.
1. [Anycast routing](https://blog.baturin.org/anycast-routing.html)

    A very quick overview of one of my favorite pieces of technology: Anycast.
1. [How OpenEBS Brings Container Attached Storage to Kubernetes](https://thenewstack.io/how-openebs-brings-container-attached-storage-to-kubernetes/)

    “Container Attached Storage enables Kubernetes users to treat storage entities as microservices. CAS has two elements — the control plane and the data plane. The control plane is deployed as a set of Custom Resource Definitions (CRD) that deal with the low-level storage entities. The data plane runs as a collection of Pods close to the workload.”
1. [Future Kubernetes Will Mimic What Facebook Already Does](https://www.nextplatform.com/2019/06/10/future-kubernetes-will-mimic-what-facebook-already-does/)

    “The pity is that Facebook is not going to create an open source version of the Tupperware cluster and container controller, or the new Delos storage service that is underpinning the current iteration of the control plane in Tupperware, both of which were discussed at Facebook’s System @scale event late last week.”
1. [Investigate and Troubleshoot Linux Issues Faster](https://www.datadoghq.com/dg/monitor/linux-monitoring/?utm_source=Advertisement&utm_medium=Advertisement&utm_campaign=DevOpsish-Newsletter03&utm_content=Linux)

    Start monitoring all your servers, both on-prem and in the cloud, within minutes with Datadog. Create actionable alerts to improve Linux performance and seamlessly navigate between logs, infrastructure metrics, and application traces in one place for faster troubleshooting. Try monitoring Linux with Datadog for free. SPONSORED
1. [7 best practices: Building applications for containers and Kubernetes](https://enterprisersproject.com/article/2020/3/kubernetes-best-practices-building-applications-containers)

    Let’s examine key considerations for building new applications specifically for containers and Kubernetes, according to cloud-native experts
1. [Migrating applications to containers and Kubernetes: 5 best practices](https://enterprisersproject.com/article/2020/3/kubernetes-migrating-applications-containers-5-best-practices)

    Let’s examine key considerations for migrating existing applications to containers and Kubernetes, according to experts
1. [Kubernetes Namespaces Explained in 15 mins (YouTube)](https://youtu.be/K3jNo4z5Jx8)

    Introduction to Kubernetes Namespaces: What are Kubernetes Namespaces? How do Namespaces help you manage your Kubernetes resources and how to use them?

### [ << Prev ](devopsweekly-169.md) ------------- [ Next >> ](devopsweekly-171.md)