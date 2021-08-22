## [DevOps'ish 133](https://devopsish.com/133) - Sun Jun 23, 2019

<a href="https://turbonomic.com/state-of-multicloud/?utm_campaign=7012o000001oRz6AAE"><strong>2019 State of Multicloud</strong></a><br/>A Report on the Underlying Dynamics Fueling Multicloud Strategies. <a href="https://turbonomic.com/state-of-multicloud/?utm_campaign=7012o000001oRz6AAE">Download Today!</a> <em>SPONSORED</em>

<a href="https://try.datree.io/trace-commits-to-tickets"><strong>[Webinar] Every commit should have an issue ticket number</strong></a><br/>Referencing an issue ticket in each commit is a development best practice. It improves code reviews, creates audit trails, and keeps you compliant. Learn how to implement this scalably with CircleCI + Datree.io. <em>SPONSORED</em>

### DevOps’ish Last Week’s Top Five

### People

1. [The worst morale boosting gesture I’ve experienced](https://shkspr.mobi/blog/2019/06/the-worst-morale-boosting-gesture-ive-experienced/)

     This might be the single dumbest morale booster I’ve ever heard (and I’ve seen some really bad ones).
1. [When Your Job Is Your Identity, Professional Failure Hurts More](https://hbr.org/2019/06/how-we-confuse-our-roles-with-our-self)

     “…when you take professional kicks personally you compromise your own ability to recover and see the bigger picture…” Perhaps you are the one doing the kicking and your organization is reacting accordingly.
1. [Top 5 job markets for sysadmins, 7 different ways](https://www.redhat.com/sysadmin/top-job-markets-sysadmins-2019)

     ”Looking for work or charting out where you need to be for your career? Here’s a breakdown of where sysadmins are working and getting hired as their role evolves.”
1. [Command Line Heroes](https://www.redhat.com/en/command-line-heroes)

     I am in Season 3 of Command Line Heroes! Add it to your favorite podcast player and wait patiently for it to drop on June 25th!
1. [How my distributed team communicates so no context is left behind](https://circleci.com/blog/how-my-distributed-team-communicates-so-no-context-is-left-behind/)

     If you are on a remote friendly team or manage remote employees this is a REALLY important read.
### Process

1. [TCP SACK PANICNetflixLWN](https://access.redhat.com/security/vulnerabilities/tcpsack)

     Originally discovered by , these TCP selective acknowledgment vulnerabilities impact Linux and FreeBSD kernels. “Multiple TCP-based remote denial of service vulnerabilities” (four CVEs in total) basically creating a new ping of death. Woohoo! Y’all exhausted from these mega vulns yet?  has a great overview.
1. [DevOps’ish Summer 2019 Surveytakes the survey today](https://devopsi.sh/survey)

    is coming to a close. It’s vital feedback that I will use to pivot, fork, or modify the newsletter, if needed. Therefore it’s important that everyone !
1. [Forget monoliths vs. microservices. Cognitive load is what matters](https://techbeacon.com/app-dev-testing/forget-monoliths-vs-microservices-cognitive-load-what-matters)

     “Excessive cognitive load works against effective team ownership and supportability of software. Here’s why, and how to approach the problem.” FINALLY! 🤯🤯🤯
1. [3 Strategies for implementing a microservices architecture](https://about.gitlab.com/2019/06/17/strategies-microservices-architecture/)

     Three ways to attack that weighty monolith.
1. [Explaining CVE-2019-10164 + PostgreSQL Security Best Practices](https://info.crunchydata.com/blog/explaining-cve-2019-10164-with-postgresql-security-best-practices)

     I love how security issues in the PostgreSQL community are solved emphatically by upgrading to the latest version. I know that’s not as cut and dry as it sounds but, think about that recommendation for MySQL or Oracle 😱🙀😱🙀
1. [Understanding Public Key Infrastructure and X.509 Certificates](https://www.linuxjournal.com/content/understanding-public-key-infrastructure-and-x509-certificates)

     PKI basics are vital for almost everyone reading this. This will get you some CLI time with openssl.
1. [IBM to win unconditional EU okay for $34 billion Red Hat dealthis is the last quarter Red Hat will be reporting earningsDevOps’ish IBM Red Hat Acquisition Index](https://www.reuters.com/article/us-red-hat-m-a-ibm-eu/ibm-to-win-unconditional-eu-okay-for-34-billion-red-hat-deal-sources-idUSKCN1TK28K)

     Coupled with reports that , it would appear the end of an independent Red Hat is here. I’m optimistic about what the future brings! Note: both of these articles have been added to the .
1. [VMware Eyeing Containers](https://seekingalpha.com/article/4270923-vmware-eyeing-containers)

     A Wall Street take on the goings on of VMware which is starting to look like a great destination if the future of Kubernetes is something you want to work on.
1. [Cisco and IBM Cloud Announce Hybrid Cloud Partnership](https://blogs.cisco.com/datacenter/cisco-and-ibm-cloud-announce-hybrid-cloud-partnership)

     Enterprise vendors; holding hands and skipping through the hybrid clouds together.
1. [Nines are not enough: meaningful metrics for clouds](https://blog.acolyer.org/2019/06/19/nines-are-not-enough/)

     ”When it comes to SLOs, the interests of the customer and the cloud provider are at odds, and so we end up with SLAs (Service Level Agreements) that tie SLOs to contractual agreements.”
1. [What is DevOps](https://devopsish.com/what-is-devops/)

     “DevOps is the professional practice of frequent, continued, and iterative improvements through measurable changes, the goal of which is to become a high-velocity organization thus improving business outcomes.”
### Tools

1. [Future of CRDs: Structural Schemas](https://kubernetes.io/blog/2019/06/20/crd-structural-schema/)

     “The core of a structural schema is an OpenAPI v3 schema made out of properties, items, additionalProperties, type, nullable, title, and descriptions. In addition, all types must be non-empty, and in each sub-schema only one of properties, additionalProperties or items may be used.”
1. [The future of how Ansible content is handled](https://youtu.be/CGjDarrqBjg)

     Collections are coming and the Ansible team needs your feedback. Kick the tires and let us know!
1. [Kubernetes 1.15: Extensibility and Continuous Improvement](https://kubernetes.io/blog/2019/06/19/kubernetes-1-15-release-announcement/)

     ”25 enhancements: 2 moving to stable, 13 in beta, and 10 in alpha”
1. [All You Need to Know About Red Hat Universal Base Imagedisclaimercontainer for Alibaba Cloud’s ossutilmicrodnf](http://crunchtools.com/all-you-need-to-know-about-red-hat-universal-base-image/)

     If my social media is any indication, there appears to be significant interest in Universal Base Image (UBI). I’m biased (Red Hat employee; see ) but, I think it’s pretty slick. I made a very lean and secure  to use in CI in no time. The trick is using  (free login required).
1. [Introducing Volume Cloning Alpha for Kubernetes](https://kubernetes.io/blog/2019/06/21/introducing-volume-cloning-alpha-for-kubernetes/)

     ”This feature allows you to create new volumes using the contents of existing volumes in the user’s namespace using the Kubernetes API.”
1. [Follow logs from multiple K8s Pods in a Deployment, ReplicaSet, etc.](https://www.jeffgeerling.com/blog/2019/follow-logs-multiple-k8s-pods-deployment-replicaset-etc)

    
1. [Automating Highly Available Kubernetes and external ETCD cluster setup with terraform and kubeadm on AWS](https://ifritltd.com/2019/06/16/automating-highly-available-kubernetes-cluster-and-external-etcd-setup-with-terraform-and-kubeadm-on-aws/)

     ”[A] set of terraform and bash scripts which should be sufficient enough for you to literally just run terraform plan/apply to get your HA etcd and k8s cluster up and running without any hassle…” Bold statement.
1. [Faster Docker builds with pipenv, poetry, or pip-tools](https://pythonspeed.com/articles/pipenv-docker/)

     If you’re building python containers, these are some great suggestions.
1. [New C5 instance sizes and bare metal instances](https://aws.amazon.com/blogs/aws/now-available-new-c5-instance-sizes-and-bare-metal-instances/)

     More compute for the compute gods.
1. [Introducing time.cloudflare.comNTP Pool Project](https://blog.cloudflare.com/secure-time/)

     I’m a huge NTP nerd. The fact Cloudflare is offering this service is good. Quality time sources on the internet are drying up. The fact they took the extra step to provide a secure NTP service is good too. But, I would still use the .
1. [You can now download the new Open Source Windows Terminal](https://www.hanselman.com/blog/YouCanNowDownloadTheNewOpenSourceWindowsTerminal.aspx)

    
1. [Getting wildcard SSL certificate in Kubernetes with cert-manager](https://medium.com/@Amet13/wildcard-k8s-4998173b16c8)

     ”[H]ow to get an SSL certificate with HTTP01 validation and a wildcard certificate with DNS01 validation on AWS”
1. [bashfulrobot/bashfulrobot-ansible](https://github.com/bashfulrobot/bashfulrobot-ansible)

     Ansible Repo that utilizes an ansible-pull command to configure my workstations
1. [containerenv/containerenv](https://github.com/containerenv/containerenv)

     shippable linux user environments

### [ << Prev ](devopsweekly-132.md) ------------- [ Next >> ](devopsweekly-134.md)