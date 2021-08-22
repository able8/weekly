## [DevOps'ish 113](https://devopsish.com/113) - Sun Feb 3, 2019

If you received this e-mail, it means we all made it through another <a href="https://www.imdb.com/title/tt0107048/">Groundhog Day</a> here in the US (<a href="https://en.wikipedia.org/wiki/Groundhog_Day">it’s a real “holiday”</a>). The movie, starring Bill Murray, is hands down, bar none one of the best movies ever made. A curmudgeonly news reporter is sent to cover a groundhog (<em>Marmota monax</em>, also called “woodchuck”) peeking out of its hole; if it sees its shadow, it means six more weeks of winter. But the reporter ends up repeating the same day over and over again until he gets his poop in a group. It ends happily, but the journey was truly formative. Again, this movie is up there with some of the best.

It is a lot like DevOps (or any kind of orchestrated change). You show up intending to do one thing and end up having to do a bunch of other stuff. Then you start measuring, optimizing, and you start getting to the stuff in the backlog. You’ll get through some of that backlog at a good clip once you have processes in place to handle day-to-day operations. Once you’re executing at a high level like this, things get easier to optimize. There’s a <a href="https://www.gartner.com/en/research/methodologies/gartner-hype-cycle">Trough of Disillusionment</a> along the way, but it does end happily most of the time. The journey might be more torture than learning at times, but you need to stick to it. The outcome is the goal, not the tools or how you got to it.

<a href="https://www.gocd.org/2019/01/14/cd-metrics-deployment-lead-time/"><strong>Use Lead Time Metric to Improve Your CI/CD Process</strong></a><br/>Check out GoCD’s latest blog in CD Metrics series. It talks about what lead time mean in CD context and guide you how to identify bottlenecks and improve your CI/CD process. <em>SPONSORED</em>

<a href="https://www.meetup.com/triangle-devops/events/257189603/"><strong>Triangle DevOps presents DevOps is Not War with Chris Short of Red Hat</strong></a><br/>Over the past 500 years, there have been 16 cases of a rising power threatening to displace a ruling power. 75% of those cases resulted in war. Although your organizational transformation probably won’t lead to war, it could be contentious. History can help prevent conflict when driving change. This talk will analyze human tendencies, historical data, and provide real-world examples of how to prevent conflict during your DevOps journey. <em>SPONSORED</em>

### DevOps’ish Top Five from Last Week

### People

1. [DevOps Team Topologies](https://web.devopstopologies.com/)

     “What Team Structure is Right for DevOps to Flourish?” This is solid gold right here. So many nails being hit on the head here.
1. [The Open Source Computer Science Degree](https://github.com/ForrestKnight/open-source-cs)

     ”This is a curated list of free courses from reputable universities like MIT, Stanford, and Princeton that satisfy the same requirements as an undergraduate Computer Science degree, minus general education.” 18 year old me would’ve killed for this.
1. [What we’ve learned about hiring engineering managers](https://circleci.com/blog/what-we-ve-learned-about-hiring-engineering-managers/)

     A thorough but not terrible hiring process. Tell me what you think.
1. [On Being A Principal Engineer](https://blog.dbsmasher.com/2019/01/28/on-being-a-principal-engineer.html)

     Silvia Botros’ being on being a Principal Engineer. I would go so far to say Silvia’s Principal definition applies to all fields.
1. [An Open Source Artificial Pancreas](https://lwn.net/SubscriberLink/777587/1427d9a6bda5d719/)

     “Getting diagnosed with a chronic disease is ‘like getting struck by lightning’, [Dana Lewis] said; there is no time to prepare and you know that everything will be different from that point forward.” Can confirm.
1. [The Importance of One-on-Ones](https://css-tricks.com/the-importance-of-one-on-ones/)

     Having some guide rails is important. Validating people is important. Make people better and they’ll make things better.
### Process

1. [FacebookGoogleFacebook employees demonstrated questionable ethics](https://www.businessinsider.com/facebook-employees-angry-after-apple-blocks-its-internal-ios-apps-2019-1)

    Apple finally took action against  and  for unsavory practices violating Apple developer terms. It’s an open secret that the companies have been playing a game of cat and mouse. Apple finally said a line had been crossed.  on dumpster fire app, Blind. Taking money for the privilege to spy on people should be against the law.
1. [7 pieces of contrarian DevOps advice](https://enterprisersproject.com/article/2019/1/devops-advice-7-contrarian-pieces)

     “Common wisdom sometimes falls flat for DevOps teams. Consider these DevOps tips learned the hard way.”
1. [The key differences in DevOps for small vs. large organizations](https://opensource.com/article/19/1/devops-small-medium-large-organizations)

     When embracing a DevOps mindset, does an organization’s size matter?
1. [What is GitOps and why you should know about it](https://venturebeat.com/2019/02/02/what-is-gitops-and-why-you-should-know-about-it/)

     GitOps is good stuff.
1. [Microsoft Azure data deleted because of DNS outage](https://nakedsecurity.sophos.com/2019/02/01/dns-outage-turns-tables-on-azure-database-users/)

     Not a typo.
1. [How to Respond to Any Crisis (IT or Otherwise) in an Organized Way](https://thenewstack.io/how-to-respond-to-any-crisis-it-or-otherwise-in-an-organized-way/)

     A talk about PagerDuty’s Incident Response process
1. [If You Think “Digital Transformation” is About Kubernetes, You’re Wrong](https://content.pivotal.io/pivotal-blog/digital-transformation-kubernetes)

     People, Process, Tools. They’re in that order for a reason.
1. [Linux Kernel Spectre Protection Changes to Boost App Performance](https://www.bleepingcomputer.com/news/linux/linux-kernel-spectre-protection-changes-to-boost-app-performance/)

     Welcome news from my choked to death eighth-gen Intel i5.
1. [It’s Time to Stop Treating R&D as a Discretionary Expenditure](https://hbr.org/2019/01/its-time-to-stop-treating-rd-as-a-discretionary-expenditure)

     A lot of what I do I would classify as research. Figuring out how systems of a certain age and size still requires a degree of understanding that can only be gained by research.
### Tools

1. [Reaching for the Stars with Ansible Operator](https://blog.openshift.com/reaching-for-the-stars-with-ansible-operator/)

     I’m excited af about Ansible Operators, y’all. “Hello Kubernetes, Meet Ansible” 👀
1. [Scaling to billions of requests on top of AWS EKS](https://medium.com/followanalytics/scaling-to-billions-of-requests-on-top-of-aws-eks-e692ec09e162)

     This is the kind of scale cloud providers are for. But, the density, elasticity, and scale you can get with Kubernetes and AWS (or Google Cloud, Azure, etc.) is staggering.
1. [Kubernetes as a surprisingly affordable platform for personal projects - a follow-up](https://blog.florentdelannoy.com/blog/2019/kubernetes-surprisingly-affordable-platform-followup/)

     Installing GKE on low cost, preemptible VMs is a compelling platform for personal projects.
1. [Everything you need to know about Kubernetes, the Google-created open source software so popular even Microsoft and Amazon had to adopt it](https://www.businessinsider.com/what-is-kubernetes-google-cloud-2019-1)

     Thanks for the REALLY long title, Business Insider. It’ll make CIOs happy for sure so share it with them.
1. [Lessons learned scaling PostgreSQL database to 1.2bn records/ monthaiven](https://medium.com/@gajus/lessons-learned-scaling-postgresql-database-to-1-2bn-records-month-edc5449b3067)

     Great googly moogly that’s a lot of data. But, the journey on how they got to that point is fascinating.I’d never heard of  before either.
1. [Building a hybrid x86–64 and ARM Kubernetes Cluster](https://medium.com/@carlosedp/building-a-hybrid-x86-64-and-arm-kubernetes-cluster-e7f94ff6e51d)

     Multiarch Kubernetes Clusters? Yes, please!
1. [Podman can now ease the transition to Kubernetes and CRI-O](https://developers.redhat.com/blog/2019/01/29/podman-kubernetes-yaml/)

     Out with Docker, in with Podman
1. [DNS Servers You Should Have Memorized](https://danielmiessler.com/blog/dns-servers-you-should-have-memorized/)

     I cannot emphasize how important it is to know a few good DNS servers.
1. [Git for Computer Scientists](http://eagain.net/articles/git-for-computer-scientists/)

     Quick introduction to git internals for people who are not scared by words like Directed Acyclic Graph.
1. [openfaas/openfaas-cloud](https://github.com/openfaas/openfaas-cloud)

     OpenFaaS Cloud: multi-user serverless functions managed with git
1. [Homebrew 2.0.0](https://brew.sh/2019/02/02/homebrew-2.0.0/)

     “The most significant changes since 1.9.0 are official support for Linux and Windows 10 (with Windows Subsystem for Linux)”
1. [cristim/serverless-failure-stories](https://github.com/cristim/serverless-failure-stories)

     Compilation of public failure/horror stories related to Serverless, inspired by hjacobs/kubernetes-failure-stories
1. [lorin/resilience-engineering](https://github.com/lorin/resilience-engineering)

     Resilience Engineering Notes

### [ << Prev ](devopsweekly-112.md) ------------- [ Next >> ](devopsweekly-114.md)