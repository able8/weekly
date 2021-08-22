## [DevOps'ish 195](https://devopsish.com/195) - Sun Dec 6, 2020

A few themes to this week’s news are worth discussing here in the newsletter’s introduction to give folks a clearer picture of each topic. We’ll tackle them in the same People, Process, and Tools format DevOps’ish uses (which are the three core components of DevOps, in order of importance). Surprisingly, I have to explain the Tools section of the news the most because it involves one of the world’s most toxic companies, <strong>Docker</strong>.

I want to take this moment to remind folks about this site’s <a href="https://devopsish.com/terms#disclaimer">disclaimer</a>.

Google continues to display lousy behavior towards people who highlight what it’s doing wrong, potentially to the detriment of the entire planet. Google fired Timnit Gebru for an email (which isn’t unheard of), but it’s a fact she found flaws in large language models, which are a big part of Google’s operating model could hurt people. The concern is around the staggering impact on the environment and economics of large language models.

Building these models takes a significant amount of compute costing somewhere between $942,973-$3,201,722 and dumping 626,155 pounds (284,019 kilograms) of carbon into the environment. There’s also a <a href="https://googlewalkout.medium.com/standing-with-dr-timnit-gebru-isupporttimnit-believeblackwomen-6dadc300d382">Standing with Dr. Timnit Gebru</a> page that I’ve signed, and I’d encourage you to do so as well (if it safe for you to do so without risk of retaliation).

The most comprehensive coverage of Timnit Gebru’s firing is from MIT Technology Review’s, <a href="https://www.technologyreview.com/2020/12/04/1013294/google-ai-ethics-research-paper-forced-out-timnit-gebru/"><em>We read the paper that forced Timnit Gebru out of Google. Here’s what it says.</em></a>

The actual email Timnit Gebru sent that got her fired is available via <a href="https://www.platformer.news/p/the-withering-email-that-got-an-ethical">Platformer News</a>. It delves into the current climate at Google and how it’s not living up to its expectations. The truth hurts so much it’ll get you fired at Google.

All your SaaS are belong to Salesforce. At least that’s their goal it seems. The world was thinking this could be bigger than the IBM-Red Hat acquisition. It almost did, <a href="https://www.nytimes.com/2020/12/01/technology/salesforce-slack-deal.html">clocking in $27.7 Billion</a>. But, why did Slack want to be acquired? According to <a href="https://www.theverge.com/22150313/how-microsoft-crushed-slack-salesforce-acquisition">The Verge</a>, Slack grew to 12 million users, while since 2016, Microsoft Teams grew to 115 million thanks to its free inclusion in Office365. The acquisition will position Slack to better compete with Microsoft Teams but, this is a fine example of Microsoft’s monopoly building skills of the 1990s in full effect.

From the <a href="https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md">Kubernetes 1.20 changelog</a>:

Mirantis and Docker don’t have their eye on Kubernetes. I think we can infer that from this week’s debacle. That became apparent this week when <a href="https://twitter.com/iancoldwater/status/1334161889476157442">Ian Coldwater noted the dockershim deprecation</a> was coming in Kubernetes 1.20. In early October, a <a href="https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1985-remove-dockershim">Kubernetes Enhancement Proposal (KEP) was created to remove dockershim from kubelet</a> (because either a shim turns into something maintainable or it goes away).

For most people, <a href="https://kubernetes.io/blog/2020/12/02/dont-panic-kubernetes-and-docker/">this shouldn’t be a huge deal</a>. But the entire team of people it took, led by Kat Cosgrove, to explain how the technologies folks were using work deserve a huge thank you. They all took on the monumental task of informing run of the mill Docker users that what they are working on is <a href="https://twitter.com/mlbiam/status/1335238217168203776">legacy tooling</a> at best.

This usually isn’t a big deal, but Docker has built a toxic community full of people who think they know better. I knew this from <a href="https://chrisshort.net/docker-inc-is-dead/">my 2017 article, where I declared Docker dead</a>. Ian and Kat sadly experienced this first hand. Ian raised the alarm in the CNCF Ambassadors channel that they could use some help. This is where I screwed up. Sitting on the Kubernetes Upstream Marketing team, I should have alerted our crew immediately. But, in all fairness, this shouldn’t have been necessary.

<a href="https://queue.acm.org/detail.cfm?id=3380777">We all make assumptions about how things work. It’s a proven fact (I strongly encourage you to read more on this topic from Dr. Richard I. Cook).</a> When someone shows us differences in our assumptions, we should assume that we’re in for a learning experience. Never should we mount up for a fight, which is what a lot of Docker users did. That and being creepy or derogatory toward the people trying to explain the aforementioned KEP. On a medium not designed to have deep, technical conversations.

Item nine on the <a href="https://raw.githubusercontent.com/cncf/trailmap/master/CNCF_TrailMap_latest.pdf">Cloud Native Trail Map</a> are Container Registry and Runtime. You can run any <a href="https://opencontainers.org/">OCI compliant</a> container runtime thanks to the <a href="https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/">Container Runtime Interface (CRI) in Kubernetes</a>. The CNCF landscape is rather <a href="https://landscape.cncf.io/category=container-runtime&amp;format=card-mode&amp;grouping=category">specific about what container runtimes work with Kubernetes</a>. Containerd, donated to CNCF by Docker years ago, is sort of the default as it’s the only graduated runtime out there. Docker was hard to work with back in their heyday (still is apparently), so Red Hat created and donated CRI-O to CNCF, a lightweight alternative to Docker, which like containerd is based on <a href="https://github.com/opencontainers/runc">runc</a> from OCI. But, all these runtimes offer something that the other doesn’t. CRI-O answered Docker’s bloated implementations (Docker is more than one thing, and they have a considerable naming problem). Containerd is the OCI compliant runtime that Docker officially switched over to a while back. But, you could still run <a href="http://petty.company/products/container-runtime">docker-engine runtimes</a> in Kubernetes.

But, dockershim is on its way out of the kubelet. What should you do? Investigate what runtimes are in use in your environment and make the switch to an OCI compliant runtime. Trust but verify, folks.

Allow me to elaborate on the “Mirantis and Docker don’t have their eye on Kubernetes” point. The deprecation of dockershim has been a long time in the making. The KEP, which is an enormous amount of information to assemble and has many artifacts as a part of it, was published on 2020-10-06. A full two months ago. That’s when Mirantis and Docker should’ve done something. But, no, it took this massive Twitter shit storm, an <a href="https://kubernetes.io/blog/2020/12/02/dockershim-faq/">FAQ</a>, and a <a href="https://kubernetes.io/blog/2020/12/02/dont-panic-kubernetes-and-docker/">blog post</a> from the Kubernetes team before Mirantis (<a href="https://twitter.com/mlbiam/status/1335238217168203776">the legacy company in Marc Boorshtein’s example</a>) did anything to address this issue.

Mirantis did so in <a href="https://www.mirantis.com/blog/mirantis-to-take-over-support-of-kubernetes-dockershim-2/">a blog post stating they’d continue support for dockershim</a> and details on how to keep it as a runtime in Kubernetes would be coming. Who knows if Mirantis even looked at the code and all the reasons for its deprecation before they made this promise. But, they did it anyway because their business appears to be helping legacy workloads trod on. Regardless, <a href="https://twitter.com/dims/status/1335262856187695105">plans remain unchanged to remove support for dockershim from kubelet</a>. CRI is still moving towards production readiness in Kubernetes and docker-engine is still incompatible with CRI.

<strong>I’ve mentioned this before, but getting off Docker should be a strategic imperative in your organization.</strong> You must remove this <a href="https://twitter.com/chumboslice/status/1335576919434293248">technical debt</a> sooner rather than later. To quote fellow CNCF Ambassador <a href="https://twitter.com/Dixie3Flatline/">Kat Cosgrove</a>, who took on the nearly impossible task of explaining this change to folks, “Docker is no more than a user experience at this point.” <em>Note: Do not harass Kat or you’ll be biting off way more than you can chew.</em>

Alternative tools have come a long way since the advent of Docker. Consider those before continuing to march forward with what is all but legacy software. You’ll be better for it. Remember, something can die in tech without a replacement. It’s not always competition or time that ends a business or technology. They often implode under their own weight.

Still curious what this all means in more technical terms, check fellow CNCF Ambassador Gianluca Arbezzano’s blog: <a href="https://gianarb.it/blog/kubernetes-1-20-dockershim-in-practice">Kubernetes v1.20 the docker deprecation dilemma in practice</a>.

<strong>DevOps’ish is brought to you by</strong> <a href="https://www.accurics.com/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=195"><strong>Accurics</strong></a>

### Google’s firing of Timnit Gebru

1. []()

    Google continues to display lousy behavior towards people who highlight what it’s doing wrong, potentially to the detriment of the entire planet. Google fired Timnit Gebru for an email (which isn’t unheard of), but it’s a fact she found flaws in large language models, which are a big part of Google’s operating model could hurt people. The concern is around the staggering impact on the environment and economics of large language models.
1. [Standing with Dr. Timnit Gebru](https://googlewalkout.medium.com/standing-with-dr-timnit-gebru-isupporttimnit-believeblackwomen-6dadc300d382)

    Building these models takes a significant amount of compute costing somewhere between $942,973-$3,201,722 and dumping 626,155 pounds (284,019 kilograms) of carbon into the environment. There’s also a  page that I’ve signed, and I’d encourage you to do so as well (if it safe for you to do so without risk of retaliation).
1. [We read the paper that forced Timnit Gebru out of Google. Here’s what it says.](https://www.technologyreview.com/2020/12/04/1013294/google-ai-ethics-research-paper-forced-out-timnit-gebru/)

    The most comprehensive coverage of Timnit Gebru’s firing is from MIT Technology Review’s,
1. [Platformer News](https://www.platformer.news/p/the-withering-email-that-got-an-ethical)

    The actual email Timnit Gebru sent that got her fired is available via . It delves into the current climate at Google and how it’s not living up to its expectations. The truth hurts so much it’ll get you fired at Google.
### Slack acquired by Salesforce

1. [clocking in $27.7 BillionThe Verge](https://www.nytimes.com/2020/12/01/technology/salesforce-slack-deal.html)

    All your SaaS are belong to Salesforce. At least that’s their goal it seems. The world was thinking this could be bigger than the IBM-Red Hat acquisition. It almost did, . But, why did Slack want to be acquired? According to , Slack grew to 12 million users, while since 2016, Microsoft Teams grew to 115 million thanks to its free inclusion in Office365. The acquisition will position Slack to better compete with Microsoft Teams but, this is a fine example of Microsoft’s monopoly building skills of the 1990s in full effect.
### dockergate (again) 🤦‍♂️🤦‍♂️🤦‍♂️

1. [Kubernetes 1.20 changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md)

    From the :
1. [Ian Coldwater noted the dockershim deprecationKubernetes Enhancement Proposal (KEP) was created to remove dockershim from kubelet](https://twitter.com/iancoldwater/status/1334161889476157442)

    Mirantis and Docker don’t have their eye on Kubernetes. I think we can infer that from this week’s debacle. That became apparent this week when  was coming in Kubernetes 1.20. In early October, a  (because either a shim turns into something maintainable or it goes away).
1. [this shouldn’t be a huge deallegacy tooling](https://kubernetes.io/blog/2020/12/02/dont-panic-kubernetes-and-docker/)

    For most people, . But the entire team of people it took, led by Kat Cosgrove, to explain how the technologies folks were using work deserve a huge thank you. They all took on the monumental task of informing run of the mill Docker users that what they are working on is  at best.
1. [my 2017 article, where I declared Docker dead](https://chrisshort.net/docker-inc-is-dead/)

    This usually isn’t a big deal, but Docker has built a toxic community full of people who think they know better. I knew this from . Ian and Kat sadly experienced this first hand. Ian raised the alarm in the CNCF Ambassadors channel that they could use some help. This is where I screwed up. Sitting on the Kubernetes Upstream Marketing team, I should have alerted our crew immediately. But, in all fairness, this shouldn’t have been necessary.
1. [We all make assumptions about how things work. It’s a proven fact (I strongly encourage you to read more on this topic from Dr. Richard I. Cook).](https://queue.acm.org/detail.cfm?id=3380777)

    When someone shows us differences in our assumptions, we should assume that we’re in for a learning experience. Never should we mount up for a fight, which is what a lot of Docker users did. That and being creepy or derogatory toward the people trying to explain the aforementioned KEP. On a medium not designed to have deep, technical conversations.
1. [Cloud Native Trail MapOCI compliantContainer Runtime Interface (CRI) in Kubernetesspecific about what container runtimes work with Kubernetesruncdocker-engine runtimes](https://raw.githubusercontent.com/cncf/trailmap/master/CNCF_TrailMap_latest.pdf)

    Item nine on the  are Container Registry and Runtime. You can run any  container runtime thanks to the . The CNCF landscape is rather . Containerd, donated to CNCF by Docker years ago, is sort of the default as it’s the only graduated runtime out there. Docker was hard to work with back in their heyday (still is apparently), so Red Hat created and donated CRI-O to CNCF, a lightweight alternative to Docker, which like containerd is based on  from OCI. But, all these runtimes offer something that the other doesn’t. CRI-O answered Docker’s bloated implementations (Docker is more than one thing, and they have a considerable naming problem). Containerd is the OCI compliant runtime that Docker officially switched over to a while back. But, you could still run  in Kubernetes.
1. []()

    But, dockershim is on its way out of the kubelet. What should you do? Investigate what runtimes are in use in your environment and make the switch to an OCI compliant runtime. Trust but verify, folks.
1. [FAQblog postthe legacy company in Marc Boorshtein’s example](https://kubernetes.io/blog/2020/12/02/dockershim-faq/)

    Allow me to elaborate on the “Mirantis and Docker don’t have their eye on Kubernetes” point. The deprecation of dockershim has been a long time in the making. The KEP, which is an enormous amount of information to assemble and has many artifacts as a part of it, was published on 2020-10-06. A full two months ago. That’s when Mirantis and Docker should’ve done something. But, no, it took this massive Twitter shit storm, an , and a  from the Kubernetes team before Mirantis () did anything to address this issue.
1. [a blog post stating they’d continue support for dockershimplans remain unchanged to remove support for dockershim from kubelet](https://www.mirantis.com/blog/mirantis-to-take-over-support-of-kubernetes-dockershim-2/)

    Mirantis did so in  and details on how to keep it as a runtime in Kubernetes would be coming. Who knows if Mirantis even looked at the code and all the reasons for its deprecation before they made this promise. But, they did it anyway because their business appears to be helping legacy workloads trod on. Regardless, . CRI is still moving towards production readiness in Kubernetes and docker-engine is still incompatible with CRI.
1. [technical debtKat Cosgrove](https://twitter.com/chumboslice/status/1335576919434293248)

    I’ve mentioned this before, but getting off Docker should be a strategic imperative in your organization. You must remove this  sooner rather than later. To quote fellow CNCF Ambassador , who took on the nearly impossible task of explaining this change to folks, “Docker is no more than a user experience at this point.” Note: Do not harass Kat or you’ll be biting off way more than you can chew.
1. []()

    Alternative tools have come a long way since the advent of Docker. Consider those before continuing to march forward with what is all but legacy software. You’ll be better for it. Remember, something can die in tech without a replacement. It’s not always competition or time that ends a business or technology. They often implode under their own weight.
1. [Kubernetes v1.20 the docker deprecation dilemma in practice](https://gianarb.it/blog/kubernetes-1-20-dockershim-in-practice)

    Still curious what this all means in more technical terms, check fellow CNCF Ambassador Gianluca Arbezzano’s blog: .
1. [Accurics](https://www.accurics.com/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=195)

    DevOps’ish is brought to you by
### People

1. [How a Vibrating Smartwatch Could Be Used to Stop NightmaresNightWare](https://www.wired.com/story/how-a-vibrating-smartwatch-could-be-used-to-stop-nightmares/)

    Those of us with PTSD that haunts us at night now have a reason for hope of a medication-free relief to the struggles we face thanks to . I’ve already sent a message to my docs at the VA to see if I can get my hands on this. Let me know if you attempt to get a prescription for NightWare. I’d like to hear about who is on board with this and who isn’t.
1. [newsletter sponsor](https://devopsish.com/sponsor/)

    Want to see your ad in DevOps’ish? Check out the  page for all the details. SPONSORED
1. [KubeCon + CloudNativeCon North America 2020 - Virtual](https://www.youtube.com/playlist?list=PLj6h78yzYM2Pn8RxfLh2qrXBDftr6Qjut)

    CNCF has shared all the videos from KubeCon + CloudNativeCon North America 2020 - Virtual.
1. [Twitter expands hate speech rules](https://www.engadget.com/twitter-hate-speech-rules-race-ethnicity-194849330.html)

    FINALLY!!! “Under the latest rules, Twitter will require users to delete tweets with dehumanizing language targeting people based on their race, ethnicity or national origin.”
1. [Court Suspends ‘Copyright Troll’ Lawyer From Practicing Law](https://torrentfreak.com/court-suspends-copyright-troll-lawyer-from-practicing-law-201201/)

    “To protect the public from future missteps, the grievance committee of the Southern District of New York decided to suspend [Richard] Liebowitz until further order.” This is a big win.
1. [‘Father of the Indian IT industry’, Tata Consulting Services founder F. C. Kohli passes, aged 96](https://www.theregister.com/2020/11/30/tcs_founder_fc_kohli_in_memorian/)

    “Hailed as a visionary pioneer who changed both India and the software industry.” A legend leaves this plane for another.
1. [The Leadership and Artistry of Tony Hsieh](https://hbr.org/2020/11/the-leadership-and-artistry-of-tony-hsieh)

    “Tony Hsieh, the former CEO of Zappos and a legendary internet entrepreneur and management thinker, has died at age 46. His passing has inspired an outpouring of affection and remembrances, and there is an understandable rush to evaluate the impact of his leadership.”
### Process

1. [GitOps: It’s the cloud-native way](https://sdtimes.com/softwaredev/gitops-its-the-cloud-native-way/)

    I said this in 2018, “GitOps is also being declared as the ‘next big thing for DevOps’ because of their strong connection. According to Weaveworks’ [Cornelia] Davis, while DevOps doesn’t have a concrete set of practices, GitOps does provide a concrete way of doing DevOps.”
1. [Download today: container security ebook - going beyond image scanningDownload this eBook](https://security.stackrox.com/container-security-going-beyond-image-scanning.html?Source=DevOpsish&LSource=DevOpsish)

    Vulnerability scanning is often the first step to securing containers and Kubernetes, but it’s insufficient by itself. , from StackRox, to learn about the security checks, controls, and best practices DevOps should consider to secure container images and CI build pipeline, Kubernetes deployments, running workloads, and the underlying Kubernetes infrastructure. SPONSORED
1. [GitOps Decisions](https://zwischenzugs.com/2020/11/30/gitops-decisions/)

    “[C]onstructing your GitOps pipelines is far from trivial, and involves many big and small decisions that add up to a lot of work to implement as you potentially chop and change as you go.”
1. [The State of the Octoverse explores a year of change with new deep dives into developer productivity, security, and how we build communities on GitHub](https://octoverse.github.com/)

    “But as the global workplace shifted into its new reality, we also saw an increase in developer connection and camaraderie through open source. With this in mind, an important question to ask is how we make all that work sustainable.
1. []()

    Thanks to automation and collaboration, developers have been able to communicate more effectively and increase efficiency, carving out more time to do the work that matters most.”
1. [A Byzantine failure in the real world](https://blog.cloudflare.com/a-byzantine-failure-in-the-real-world/)

    “An analysis of the Cloudflare API availability incident on 2020-11-02”
1. [OpenShift/Kubernetes Failure Stories at Scale - Lessons Learned from Large and Dense Deployments](https://www.openshift.com/blog/openshift-failure-stories-at-scale-cluster-on-fire)

    Celebrating failures are more and more important than ever. Large distributed systems are hard to understand. When a failure occurs it creates an opportunity to learn and better the system.
1. [Why failure should be normalized and how to do it](https://opensource.com/article/20/11/normalize-failure)

    ”‘Everybody is perfect, except you.’ This toxic thought can creep in and ruin your confidence. Here’s how to normalize failure during the learning process and remember that everyone makes mistakes.”
1. [Threat Alert: Fileless Malware Executing in Containers](https://blog.aquasec.com/fileless-malware-container-security)

    “Our cyber research team detected a new type of attack that executes and runs malware straight from memory in containers, thus evading common defenses and static scanning. This malware is using a rootkit to hide its running processes, then hijacks resources by executing a crypto miner from memory — leaving a backdoor that enables attackers to do more damage. We found four container images in Docker Hub designed to execute fileless malware attacks.”
1. [HPE is relocating headquarters to Houston from California](https://www.cnbc.com/2020/12/01/hpe-is-relocating-headquarters-to-houston-from-california.html)

    “The coronavirus pandemic has given a number of tech companies, and prominent Silicon Valley figures, an excuse to exit California.”
1. [What an Anti-Racist Business Strategy Looks Like](https://hbr.org/2020/11/what-an-anti-racist-business-strategy-looks-like)

    “Confronted with protests against systemic racism in the United States, companies are now working to ensure that their workforces and communities are more diverse, equitable and inclusive. But what does anti-racist business strategy look like in practice? There is no one-size-fits all approach, but you can focus on four foundational and four functional elements to stand up successful initiatives. This includes figuring out the purpose of the action you are taking, including many perspectives and people, and measuring how you perform against goals. You will also want to rethink the way you do purchasing and philanthropy, review your policies, and reengage with the places in which you operate.”
### Tools

1. [Introducing another free CA as an alternative to Let’s EncryptACME servers](https://scotthelme.co.uk/introducing-another-free-ca-as-an-alternative-to-lets-encrypt/)

    ZeroSSL is offering free certs via ACME. Here’s a list of additional .
1. [OpenTelemetry Starter Kit at the Honeycomb booth](https://info.honeycomb.io/honeycomb-and-kubecon)

    🎺 Looking to up your hands-on skills at KubeCon? We are premiering our . Instrument a sample app or your own production data. There are prize-drawings for completed steps, but the real prize is launching with OTel and getting valuable insights right away. SPONSORED
1. [A sysadmin’s guide to basic Kubernetes components](https://www.redhat.com/sysadmin/kubernetes-components)

    “Kubernetes control plane nodes and worker nodes, their features, and how they interact.”
1. [Git stash doesn’t have to be scary](https://jemma.dev/blog/git-stash)

    “Stashes are more easily understood as a basic last in, first out stack.”
1. [Weaveworks Brings GitOps to Amazon EKS Distro](https://www.weave.works/blog/on-prem-kubernetes-gitops-eks-distro)

    “This is especially for those customers who need more than just a ‘day zero’ installer for one or two clusters.”
1. [Why Linkerd doesn’t use Envoy](https://linkerd.io/2020/12/03/why-linkerd-doesnt-use-envoy/)

    “But we chose not to build Linkerd on top of Envoy. Instead, we built a dedicated “micro-proxy”, called simply Linkerd2-proxy, which is optimized for the service mesh sidecar use case. In the increasingly crowded field of similar-sounding service mesh projects, Linkerd is unique in this regard. But why did we go this route?”
1. [A Broken Piece of Internet Backbone Might Finally Get Fixed](https://www.wired.com/story/bgp-routing-manrs-google-fix/)

    BGP in its current state is one of the most fragile things on the internet and it, for better or worse, controls the flow of traffic on the internet. “Efforts to secure the Border Gateway Protocol have picked up critical momentum, including a big assist from Google.”
1. [git wip: What the heck was I just doing?](https://carolynvanslyck.com/blog/2020/12/git-wip/)

    “Here’s a git alias, git wip, that displays your branches and when you last changed them. It’s really useful for remembering which one you were just working on, and which branches probably should be deleted.” This is going to save me a lot of time in the future.
1. [Malicious npm packages caught installing remote access trojans](https://www.zdnet.com/article/malicious-npm-packages-caught-installing-remote-access-trojans/)

    Typo squatting and name similarity can wreak all out havoc on your infrastructure. You have to keep an extra close eye on dependencies.
1. [Pulumi Import: Generate IaC for Existing Cloud Resources](https://www.pulumi.com/blog/pulumi-import-generate-iac-for-existing-cloud-resources/)

    “As of v2.12.0, Pulumi has introduced a pulumi import command. This command will import the cloud resource into the Pulumi state and generate the code for the user’s Pulumi program in the appropriate language.” I love environment import tools.
1. [Kubernetes YAML: Enforcing best practices and security policies in CI/CD and GitOps pipelines](https://thechief.io/c/editorial/kubernetes-yaml-enforcing-best-practices-and-security-policies-cicd-and-gitops-pipelines/)

    A wonderful list of tools to keep your YAML in check.
1. [wustho/epr](https://github.com/wustho/epr)

    CLI Epub Reader
1. [Release OpenZFS 2.0.0 · openzfs/zfs](https://github.com/openzfs/zfs/releases/tag/zfs-2.0.0)

    “The ZFS on Linux project has been renamed OpenZFS! Both Linux and FreeBSD are now supported from the same repository making all of the OpenZFS features available on both platforms”

### [ << Prev ](sreweekly-194.md) ------------- [ Next >> ](sreweekly-196.md)