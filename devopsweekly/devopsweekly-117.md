## [DevOps'ish 117](https://devopsish.com/117) - Sat Mar 2, 2019

<a href="https://www.theregister.co.uk/2019/02/25/docker_cfo_layoffs_ceo_interview/">Docker confirmed they laid off folks</a>. A source stated Docker closed their APAC office, laid off some developers, and punted some sales folks in Europe too. The source thought some managers on the development teams would be leaving soon as well. I have absolutely no real way of confirming that. But, all is well at Docker if you ask them. Things are so well in fact, that when a reporter reached out to Docker to ask about layoffs (based on a tweet from a random technologist), their PR team called the reporter to give them an “exclusive” interview with the Docker CEO to bury the layoffs. That’s totally normal behavior for a company that’s doing just peachy I guess. ¯\_(ツ)_/¯ Regardless, if you’re impacted and want help, let me know.

In personal news, it appears <a href="https://twitter.com/ChrisShort/status/1101550232490381312">my eight-year long appeal with the Veteran’s Administration</a> has come to an end. As I mentioned in the <a href="https://twitter.com/ChrisShort/status/1101550232490381312">Twitter thread</a>, there’s a lot to unpack here. It’s been an overwhelming week and finding out life-changing news at the end of it is wild but on brand for 2019 so far.

<strong>New Microsoft Azure Elastic Agent Plugin for GoCD</strong><br/>With GoCD’s new Azure plugin, run your CI/CD pipelines on Azure virtual machines, and let GoCD scale up on-demand agents based on your need. Remove bottlenecks and reduce the cost of your agent infrastructure.

Learn more about GoCD: <a href="https://www.gocd.org">www.gocd.org</a><br/>Learn more about GoCD’s Azure plugin: <a href="https://extensions-docs.gocd.org/azure/current/">https://extensions-docs.gocd.org/azure/current/</a> <em>SPONSORED</em>

### DevOps’ish Top Five from Last Week

### People

1. [Linus Torvalds pulls pin, tosses in grenade: x86 won, forget about Arm in server CPUs, says Linux kernel supremo](https://www.theregister.co.uk/2019/02/23/linus_torvalds_arm_x86_servers/)

     Linus is wrong.
1. [How SRE Creates a Blameless Culture](https://devops.com/how-sre-creates-a-blameless-culture/)

     SRE as a way to manage blamelessness.
1. [Common Assumptions to Avoid When Starting with Kubernetes](https://blog.flux7.com/common-assumptions-to-avoid-when-starting-with-kubernetes)

     Assumptions are the devil (they’re why this e-mail has had delivery issues lately).
1. [Silence Is Violence](https://www.xaprb.com/blog/silence-is-violence/)

     “If I don’t object, I’m floating in the swift, powerful river of male supremacy. If I don’t object, I’m his ally.”
### Process

1. [CVE-2019-8912](https://nvd.nist.gov/vuln/detail/CVE-2019-8912#vulnCurrentDescriptionTitle)

     In the Linux kernel through 4.20.11, af_alg_release() in crypto/af_alg.c neglects to set a NULL value for a certain structure member, which leads to a use-after-free in sockfs_setattr.
1. [OpenSSL Security Advisory [26 February 2019]](https://www.openssl.org/news/secadv/20190226.txt)

     Severity: Moderate
1. [Is it time to be afraid of IBM again?](https://redmonk.com/jgovernor/2019/02/28/is-it-time-to-be-afraid-of-ibm-again/)

     Maybe… Maybe.
1. [Killing Kubernetes](https://blog.freetrade.io/killing-kubernetes-7f8b61c701aa)

     Kubernetes is a way to get to what folks want. But, there could be better ways to meet your desired outcomes.
1. [How I passed the CKA (Certified Kubernetes Administrator) Exam](https://medium.com/@krystiannowaczyk/how-i-passed-the-cka-certified-kubernetes-administrator-exam-f94b11566528)

     How one person passed the CKA.
1. [Who needs malware? IBM says most hackers just PowerShell through boxes now, leaving little in the way of footprints](https://www.theregister.co.uk/2019/02/26/malware_ibm_powershell/)

     “A company’s internal network, once compromised, is now more likely to be ransacked by automated scripts than a piece of malware.”
1. [So, you think your CI/CD environment is secure?](https://snyk.io/blog/so-you-think-your-ci-cd-environment-is-secure/)

     Secure pipelines for GitOps
1. [Top ten most popular docker images each contain at least 30 vulnerabilities](https://snyk.io/blog/top-ten-most-popular-docker-images-each-contain-at-least-30-vulnerabilities/)

    
1. [Opensource and Go. What license?](https://medium.com/@henvic/opensource-and-go-what-license-f6b36c201854)

     “87% of Go code in GitHub is licensed using a permissive license such as MIT, Apache, and BSD 3-clause.”
1. [Experts beg the Supreme Court to overturn Oracle/Google ruling](https://www.businessinsider.com/oracle-google-supreme-court-java-android-2019-2)

    
1. [We Need Chrome No More](https://redalemeden.com/blog/2019/we-need-chrome-no-more)

     “Ten years ago, we needed Google Chrome to break the Web free from corporate hegemony, and we managed to do so for a brief period. Today, its dominance is stifling the very platform it once saved from the clutches of Microsoft. And no one, beside Google, needs that.”
1. [xkcd: Modern OSI Model](https://xkcd.com/2105/)

    
### Tools

1. [OperatorHub](https://www.operatorhub.io/)

     ”OperatorHub.io is a new home for the Kubernetes community to share Operators. Find an existing Operator or list your own today.”
1. [A quickstart guide to Ansible](https://opensource.com/article/19/2/quickstart-guide-ansible)

     A month ago, someone on my team (pretty sure my awesome boss) mentioned Ansible’s birthday was coming up (February 23rd-ish). Jason Hibbets and I huddled up with the Opensource.com team and built a nice looking book for y’all. Happy 7th birthday, Ansible!
1. [Ansible Operators](https://podcasts.apple.com/in/podcast/ansible-operators/id1270983443?i=1000430756448)

     Brian talks with Fabian von Feilitzsch (Sr. Software Engineer at RedHat) and Shawn Hurley (Sr. Software Engineer at Red Hat) about Ansible Operators, how they work with Ansible Playbook, on-platform and off-platform usage, and examples to help people learn the new Kubernetes technology.
1. [Scaling PostgreSQL with Kubernetes Operators](https://opensource.com/article/19/2/scaling-postgresql-kubernetes-operators)

     Operators let users create standardized interfaces for managing stateful applications, like PostgreSQL, across Kubernetes-enabled cloud environments.
1. [How to Install and Use Gremlin Locally with Docker For MacGremlin Free](https://www.gremlin.com/community/tutorials/how-to-install-and-use-gremlin-locally-with-docker-for-mac/)

     Oh… So you want to unleash your own Gremlin army? Now you can!  is a thing; I can’t wait to mess with this.
1. [Git series 1⁄3: Understanding git for real by exploring the .git directory](https://dev.to/daolf/git-series-13-understanding-git-for-real-by-exploring-the-git-director--5bd0)

     After I learned Git a sysadmin I was working with at the time told me they learned Git in a few hours by digging around the .git directory. YMMV.
1. [Cloudflare’s RPKI Toolkit](https://blog.cloudflare.com/cloudflares-rpki-toolkit/)

     How Cloudflare is preventing route hijacks.
1. [Introducing “Ask NGINX…”](https://www.nginx.com/blog/introducing-ask-nginx/)

     A great idea.
1. [Use less YAML](https://blog.chernand.io/2019/02/27/use-less-yaml/)

     “When working with kubernetes, you will run into lots of YAMLs that you will be copying and pasting. You can save yourself some typing if use the kubectl to create these resources for you!”
1. [sk8s – Simple Kubernetes (k8s) Virtual Appliance](https://www.virtuallyghetto.com/2019/02/sk8s-simple-kubernetes-k8s-virtual-appliance.html)

     vSphere and “simple”… Hmm. Okay.
1. [What is Dark?](https://medium.com/darklang/the-design-of-dark-59f5d38e52d2)

     I’m going to launch Short, a Programming Language the premise of which is to write as little as possible. You know what, don’t write that code. Take some painkillers and lay on the couch. In other news, Dark is interesting.
1. [Lenovo introduces the 14e Chromebook, built for the frontline workforce](https://cloud.google.com/blog/products/chrome-enterprise/lenovo-introduces-the-14e-chromebook-built-for-the-frontline-workforce)

     I might use something like this traveling to China.
1. [New VMware Kubernetes product comes courtesy of Heptio acquisition](https://techcrunch.com/2019/02/26/latest-vmware-kubernetes-product-comes-courtesy-of-heptio-acquisition/)

     There is something to be said for pure-play Kubernetes.
1. [Google Kubernetes Engine; Explain Like I’m Five!](https://medium.com/devopslinks/google-kubernetes-engine-explain-like-im-five-1890e550c099)

     “Creating your first managed Kubernetes cluster on Google Kubernetes Engine using Terraform”
1. [Kubernetes Quick Setup with Prometheus, Grafana & Jaeger](https://blog.nobugware.com/post/2019/kubernetes_quick_development_setup_minikube_prometheus_grafana/)

     “Here is a quick setup (mostly notes) to create a dev environment using Minikube and the workflow I’m using with it.”
1. [Let’s deploy to Openshift / Habr](https://habr.com/en/post/441906/)

     “I’d like to share my story about migration an application to Openshift.”
1. [A serverless approach to Weather Underground – IBM Developer](https://developer.ibm.com/articles/a-serverless-approach-to-weather-underground/)

     Using Cloud Functions and the composer programming model for Intellicast radar maps
1. [The art of writing eBPF programs: a primer](https://sysdig.com/blog/the-art-of-writing-ebpf-programs-a-primer/)

     Sysdig is doing some pretty neat stuff in the eBPF space.
1. [Go 1.12 Release Notes - The Go Programming Language](https://golang.org/doc/go1.12)

    
1. [k3s - Lightweight Kubernetes](https://k3s.io/)

     Easy to install. A binary of less than 40 MB. Only 512 MB of RAM required to run.
1. [Single-Node Kubernetes on Raspberry Pi with MicroK8s and Ubuntu](https://blog.ubuntu.com/2019/02/27/single-node-kubernetes-on-raspberry-pi-microk8s-ubuntu)

     It’s Kubernetes is The Genie in Aladdin explaining his life inside the bottle, “Phenomenal Cosmic Powers! Itty Bitty Living Space!”
1. [gyselroth/kube-icinga](https://github.com/gyselroth/kube-icinga)

     Monitor kubernetes services / resources using icinga2 (including autodiscovery support)
1. [lnsp/k8s-mattermost-informer](https://github.com/lnsp/k8s-mattermost-informer)

     Informs a Mattermost channel if an annotated pod goes into crash loop

### [ << Prev ](sreweekly-116.md) ------------- [ Next >> ](sreweekly-118.md)