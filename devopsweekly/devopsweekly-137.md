## [DevOps'ish 137](https://devopsish.com/137) - Sun Jul 21, 2019

It’s a big newsletter this week and I’ve made a lot of notes so let’s get to it!

<strong>Editor’s Note:</strong> Did you know you can reply to this e-mail? Yep! I get your out of offices. But, I can also get your feedback directly. If something doesn’t work for you or live up to your expectations, let me know.

<a href="https://www.bluematador.com/cloudwatch-guided-tour?utm_campaign=CloudWatch%20Guided%20Tour%20Webinar&amp;utm_source=devopsish&amp;utm_medium=newsletter"><strong>Join DevOps Expert, Matthew Barlocker, For A CloudWatch Guided Tour</strong></a><br/>The founder and CEO of Blue Matador, the alert automation service, will be hosting a CloudWatch Guided Tour Webinar. You’ll learn about CloudWatch concepts, alarms, metrics, best practices, and more. <a href="https://www.bluematador.com/cloudwatch-guided-tour?utm_campaign=CloudWatch%20Guided%20Tour%20Webinar&amp;utm_source=devopsish&amp;utm_medium=newsletter">Join either July 25th or July 31st</a>! <em>SPONSORED</em>

### DevOps’ish Last Week’s Top Five

### People

1. [A worrying change in Open Source perception](https://dev.to/codepo8/a-worrying-change-in-open-source-perception-30m7)

     ”Open Source should bring a positive experience.” If we as community members create a bar for participation that is too high, we all suffer.
1. [DiversityTickets: AnsibleFest 2019 Atlanta](https://diversitytickets.org/en/events/488)

     The Ansible team is using DiversityTickets again this year to make sure folks have an opportunity to attend AnsibleFest Atlanta. “This diversity program is aimed to help underrepresented groups in tech to attend and lowering the barriers for them. This includes but isn’t limited to: women, people of color, LGBTQIA+ people, and disabled people.”
1. [CloudSkills.fm 030: Cloud Native Ops with Ansible and Kubernetes](https://cloudskills.fm/030)

     I sat down with Mike Pfeiffer to talk about, “Ansible, Kubernetes, and the importance of building up your Linux skills as the industry transitions into a cloud native world.” This was a lot of fun.
1. [The cloud skills shortage and the unemployed army of the certified](https://itnext.io/the-cloud-skills-shortage-and-the-unemployed-army-of-the-certified-bd405784cef1)

     “Roles that have been separate for years — database admin, server admin, quality engineer, software developer — are being smashed together because the headcount is vanishing. These roles are asking you to do the work of several people but using your attraction to intellectual challenges to make it sound more appealing.” Entry-level certifications are going to get you architect-level jobs or any job for that matter. Experience matters a lot these days.
1. [Have you ever wondered what the hiring process was 20 years ago compared to today? Probably not, but I’ll tell you anyway](https://www.reddit.com/r/cscareerquestions/comments/brjexy/have_you_ever_wondered_what_the_hiring_process/)

     Here’s some interesting perspective.
1. [6 Causes of Burnout, and How to Avoid Them#7 and #8 open source projects in the world](https://hbr.org/2019/07/6-causes-of-burnout-and-how-to-avoid-them)

     Values mismatch is something I am struggling with right now in my own position. I highly value working in the Kubernetes community. One of my personal goals is to work full time on Kubernetes (or closer to the core of it). However, my job on the Ansible team doesn’t directly benefit from this work. It does feel like I’m being stretched incredibly thin by the . But, I also think that there are indirect benefits of this work. The problem is that it could lead to a perceived lack of fairness. I’m galavanting across the cloud native countryside some days while my co-workers are pushing the Ansible go to market strategy. Open source is weird.
### Process

1. [Preventing Attacks Using HTTP HeadersHTTP Security Headers - A Complete Guide](https://www.twilio.com/blog/preventing-attacks-using-http-headers)

     I recently re-platformed all my web sites so I could make better use of HTTP headers and more granular caching. HTTP security headers are becoming really important. If you are not familiar with them, read this article and .
1. [Facebook’s $5 billion FTC fine is an embarrassing joke](https://www.theverge.com/2019/7/12/20692524/facebook-five-billion-ftc-fine-embarrassing-joke)

     The fine is less than 1% of Facebook’s market cap. Endangering democracy is a lucrative business. The fine should have started with a T, as in trillion.
1. [MTTR is dead, long live CIRT](https://opensource.com/article/19/7/measure-operational-performance)

     ”By focusing on business-impacting incidents, CIRT is a more accurate way to gauge ops performance.” PagerDuty descends on opensource.com to bend the arc of performance measurement. Critical incident response time (CIRT); I like it! “CIRT and the percentage of incidents acknowledged and resolved form a valuable set of metrics that give you a much better idea of how your operations are performing.”
1. [Judge shoots down Oracle protest over $10B JEDI cloud contract, leaving Amazon and Microsoft as finalistsTrump Expressed Concern About Pentagon Cloud-Computing Contract](https://www.geekwire.com/2019/judge-shoots-oracle-protest-10b-jedi-cloud-contract-leaving-amazon-microsoft-finalists/)

     But, President Trump hates Jeff Bezos with a burning passion. . If you think you can afford to not be involved in politics, you are an incredibly privileged person.
1. [The Kubernetes Hierarchy of Needs](https://thenewstack.io/the-kubernetes-hierarchy-of-needs/)

     Build infrastructure, deploy applications, and operate. This falls in line with my Kubernetes maintainer/operator paradigm.
1. [Kubernetes catches up with operational reality](https://www.computerweekly.com/blog/Write-side-up-by-Freeform-Dynamics/Kubernetes-catches-up-with-operational-reality)

     Day 2 operations on Kubernetes might be the new buzz phrase for 2020.
1. [Post Mortem: Kubernetes Node OOMbelow the line](https://www.bluematador.com/blog/post-mortem-kubernetes-node-oom)

     Blue Matador (this week’s sponsor) shows us all the importance of resource limits. Especially for the  workloads you might not ever have time to fully debug.
1. [Incident Review: You Can’t Deploy Binaries That Don’t Exist](https://www.honeycomb.io/blog/incident-review-you-cant-deploy-binaries-that-dont-exist/)

    
1. [House Hearing on Facebook Libra: Why was the Rust language chosen?](https://www.c-span.org/video/?c4808083/rust-language-chosen)

     Huh.
### Tools

1. [What are the downsides of using OpenShift instead of plain vanilla Kubernetes?k/k](https://www.reddit.com/r/kubernetes/comments/cdztqg/what_are_the_downsides_of_using_openshift_instead/)

     A discussion on /r/kubernetes about the downsides of using OpenShift. Which sound a lot like the downsides of using . That’s encouraging I guess. Kubernetes at scale is no trivial task.
1. [Monitoring Linux Logs with Kibana and Rsyslog](http://devconnected.com/monitoring-linux-logs-with-kibana-and-rsyslog/)

     “This tutorial details how to build a monitoring pipeline to analyze Linux logs with ELK 7.2 and Rsyslog.” A super thorough explanation of how to get your logs in ELK.
1. [Continuous Integration for a Monolithic Ansible Repository](https://medium.com/@arslan70/continuous-integration-for-a-monolithic-ansible-repository-ce4ea09f1588)

     ”I want to focus this post on Why instead of How. There is a ton of information on implementation details. Some links in the references can help you get started as well. I hope after the end of this post you will be able to evaluate if your ansible repo requires a CI pipeline or not.” This is dope.
1. [Kubernetes Deployments: The Ultimate Guide by Jérôme Petazzoni](https://semaphoreci.com/blog/kubernetes-deployment)

     ”Ready to deploy that app you’ve just ‘dockerized’? Here’s all you need to know about Kubernetes deployments to deliver your containers to production.”
1. [Announcing Linkerd 2.4: Traffic Splitting and SMI](https://linkerd.io/2019/07/11/announcing-linkerd-2.4/)

     ”Linkerd’s new traffic splitting feature allows users to dynamically control the percentage of traffic destined for a service… Support for the Service Mesh Interface makes it easier for ecosystem tools to work with Linkerd.”
1. [Automate Progressive Deployments to Kubernetes with Flagger and Linkerd](https://www.weave.works/blog/automate-progressive-deployments-to-kubernetes-with-flagger-and-linkerd)

     ”Linkerd implements the Service Mesh Interface (SMI) Traffic Split API. This allows Flagger to control the traffic between two versions of the same application.” Weaveworks is absolutely killing it right now.
1. [IBM unveils trio of open source Kubernetes projects, and not a Red Hat-trick in sight](https://devclass.com/2019/07/16/ibm-unveils-trio-of-open-source-kubernetes-projects-and-not-a-red-hat-trick-in-sight/)

     ”[A]llow users to “architect, build, deploy, and manage the lifecycle of Kubernetes-based applications.” The project includes “pre-built deployments to Kubernetes and Knative (using Operators and Helm charts)…so, developers can spend more time developing scalable applications and less time understanding infrastructure.”
1. [Red Hat builds OpenShift Pipelines on Google-spawned Tekton](https://devclass.com/2019/07/12/red-hat-builds-openshift-pipelines-on-tekton/)

     I’m not quite sure how I feel about this. I’m sure it’s solid tech but, it seems like it’s missing a really interesting piece of software that Red Hat already owns.
1. [What is a TTY on Linux? (and How to Use the tty Command)](https://www.howtogeek.com/428174/what-is-a-tty-on-linux-and-how-to-use-the-tty-command/)

     ”What does the tty command do? It prints the name of the terminal you’re using. TTY stands for ‘teletypewriter.’ What’s the story behind the name of the command? That takes a bit more explaining.” TTY is old af. How do I know? I’ve used a physical TTY before.
1. [How to Teach Old Apps New Tricks with Ansible-based Operators (Parts 1 & 2)](https://chrisshort.net/how-to-teach-old-apps-new-tricks-with-ansible-based-operators-parts-1-2/)

     I presented two fifteen-minute talks on Ansible-based Kubernetes Operators at Red Hat Summit. I wanted to share the slides because conceptually, we’ve got k8s operators narrowed down pretty closely to Ansible contexts with a smidgen or k8s YAML. It lowers the barrier to entry.
1. [33 Kubernetes security tools](https://sysdig.com/blog/33-kubernetes-security-tools/)

     “Kubernetes security tools … there are so freaking many of them; with different purposes, scopes and licenses.”
1. [30 Top Tech and Dev Podcasts in 2019](https://rosnovsky.us/blog/tech-podcasts-in-2019/)

    
1. [intel/hyperscan](https://github.com/intel/hyperscan)

     High-performance regular expression matching library
1. [Qiskit/qiskit](https://github.com/Qiskit/qiskit)

     Qiskit is an open source framework for working with noisy quantum computers at the level of pulses, circuits, and algorithms. I spoke with two teams members from Qiskit this week. IBM has some brilliant people on this project. They could use help community building and scaling up. If you have some free cycles to help with a project, they would greatly appreciate it.
1. [digitalocean/DOKS](https://github.com/digitalocean/DOKS)

     Managed Kubernetes designed for simple and cost effective container orchestration.
1. [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper)

     Gatekeeper - Policy Controller for Kubernetes
1. [stakater/Reloader](https://github.com/stakater/Reloader)

     A Kubernetes controller to watch changes in ConfigMap and Secrets and then restart pods for Deployment, StatefulSet, DaemonSet and DeploymentConfig
1. [google/docsy-example](https://github.com/google/docsy-example)

     An example documentation site using the Docsy Hugo theme. I really like the Docsy project.
1. [xo/xo](https://github.com/xo/xo)

     Command line tool to generate idiomatic Go code for SQL databases supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server
1. [kubernetes-sigs/k8s-container-image-promoter](https://github.com/kubernetes-sigs/k8s-container-image-promoter)

     A tool to promote Docker images from one registry to another, based on a declarative YAML manifest
1. [deislabs/smi-spec](https://github.com/deislabs/smi-spec)

     Service Mesh Interface
1. [reneschoonrok/Kubernetes-3d-visualizer](https://github.com/reneschoonrok/Kubernetes-3d-visualizer)

     Interact with your kubernetes cluster rendered in 3d using threejs and css3d

### [ << Prev ](sreweekly-136.md) ------------- [ Next >> ](sreweekly-138.md)