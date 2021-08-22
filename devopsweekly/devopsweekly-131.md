## [DevOps'ish 131](https://devopsish.com/131) - Sun Jun 9, 2019

This week’s image was something I whipped up when talking to a friend. When I explained that running Kubernetes on a Raspberry Pi cluster or Digital Ocean didn’t mean they had to be an expert at running Kubernetes, it tilted their world off axis a little. I blew away a cluster and rebuilt it from a handful of files in less than thirty minutes a couple of weeks ago. This should be a rare occurrence (and there should be backups). But, if there isn’t time to troubleshoot (and you maintain proper configuration management), burn it down and start over if you make a mistake. It points to the practice Rob Hirschfeld discovered of <a href="https://thenewstack.io/the-optimal-kubernetes-cluster-size-lets-look-at-the-data/">keeping Kubernetes clusters small</a> to reduce blast radius and snowflakes.

<a href="https://turbonomic.com/state-of-multicloud/?utm_campaign=7012o000001oRz6AAE"><strong>2019 State of Multicloud</strong></a><br/>A Report on the Underlying Dynamics Fueling Multicloud Strategies. <a href="https://turbonomic.com/state-of-multicloud/?utm_campaign=7012o000001oRz6AAE">Download Today!</a> <em>SPONSORED</em>

<a href="https://devopsi.sh/survey"><strong>DevOps’ish Summer 2019 Survey</strong></a><br/>Please take this year’s survey. It is a handful of questions, will provide actionable feedback, and be greatly appreciated. The intent is to improve the overall experience and increase the value this newsletter provides to its readers. <a href="https://devopsi.sh/survey">Take the survey today!</a>

<a href="https://www.oreilly.com/pub/cpc/224549"><strong>How companies adopt and apply cloud native infrastructure–from O’Reilly</strong></a><br/>Survey results reveal the path organizations face as they integrate cloud native infrastructure and harness the full power of the cloud. <em>SPONSORED</em>

### DevOps’ish Last Week’s Top Five

### People

1. [IBM layoffs affect more than 1,000 employees](https://www.cnbc.com/2019/06/06/ibm-layoffs-affect-more-than-1000-employees.html)

     It was interesting to learn this week some folks at Red Hat received retention bonuses while others did not. This is a reminder of why the IBM acquisition of Red Hat makes people concerned.
1. [Experience from Six Months of Remote Work](https://www.justingarrison.com/blog/2019-05-29-six-months-remote/)

     ”While I’ve had lots of experience working in open source and tons of experience with communicating asynchronously doing full time work remotely is different. I wanted to share some things I’ve learned so far.”
1. [The guy who made a tool to track women in porn videos is sorry](https://www.technologyreview.com/s/613607/facial-recognition-porn-database-privacy-gdpr-data-collection-policy/)

     What could possibly go wrong?!? ¯\_(ツ)_/¯
1. [New Evidence Suggests Satoshi Nakamoto Is Paul Solotshi, The Creator of Encryption Software E4M and TrueCryptChloe Condon](https://www.investinblockchain.com/new-evidence-suggests-satoshi-nakamoto-is-paul-solotshi-the-creator-of-encryption-software-e4m-and-truecrypt/)

     Everyboy knowns that  created Bitcoin (and probably blockchain).
1. [How I came back from burnout by learning to lean on my team - CircleCI](https://circleci.com/blog/how-i-came-back-from-burnout-by-learning-to-lean-on-my-team/)

    
### Process

1. [The problem with Amazon and Open Source isn’t Amazon](https://anonymoushash.vmbrasseur.com/2019/06/07/the-problem-with-amazon-and-open-source-isnt-amazon/)

     Jen Wike Huger and I played a part in starting this shit storm. But, in general, there is nothing wrong with open source licenses. There are a lot of things wrong with trying to build an entire business on 100% open source software with little to no value add. Yet, I do believe a reemergence of shareware could occur in a big way. Something along the lines of this software is for personal use only. Contact us for a license is happening already. I see a potential future where a competitor to OSI might form around these form of open source. Enough lawyers are poking at this and Linux Foundation is standing up new foundations every time Linus Torvalds blinks. There is too much smoke here for something not to be on fire.
1. [Why We’re Relicensing CockroachDB](https://www.cockroachlabs.com/blog/oss-relicensing-cockroachdb/)

     This is the only license I think is being honest while legal about commercial usage. But, make a better business.
1. [Google Cloud Networking Incident #19009](https://status.cloud.google.com/incident/cloud-networking/19009)

     Network configuration change resulting in degraded service preventing connections to network devices to rollback changes is like the worst nightmare I think we’ve all been in at least once. The distinction in traffic Google can then manage and prioritize is fascinating.
1. [What is DevOps](https://devopsish.com/what-is-devops/)

     ”DevOps is the professional practice of frequent, continued, and iterative improvements through measurable changes, the goal of which is to become a high-velocity organization thus improving business outcomes.”
1. [Empowering Users through Site Reliability Engineering](http://jasonhand.com/blog/empowering_users_through_site_reliability_engineering/)

     ”How does the service we are building change the way they perceive themselves? Does it make them better at something?”
1. [Sustainable Operations in Complex Systems with Production Excellence](https://www.infoq.com/articles/production-excellence-sustainable-operations-complex-systems/)

     ”Cultural and process changes, rather than changes in tooling alone, are necessary for teams to sustainably manage services… Services run more smoothly when quantitative risk analysis allows teams to prioritize fixes.”
1. [This week’s dead Google product is Google Trips, may it rest in peace](https://arstechnica.com/gadgets/2019/06/this-weeks-dead-google-product-is-google-trips-may-it-rest-in-peace/)

     Hard to break this image Google has when products keep getting knocked off left and right.
1. [The Linux Foundation Fires All Staff and Editors at Linux.com. Future Uncertain.](http://techrights.org/2019/06/05/linux-com-future-uncertain/)

     This is not positive.
1. [Rapid Kubernetes Growth Creates Some Pain](https://containerjournal.com/2019/06/04/rapid-kubernetes-growth-creates-some-pain/)

     ”A survey of 500 attendees at the recent KubeCon + CloudNativeCon 2019 Europe conference finds that nearly half the respondents (47%) are now running an instance of Kubernetes in a production environment.”
### Tools

1. [Bash coprocess](https://medium.com/@copyconstruct/bash-coprocess-2092a93ad912)

     “Bash versions 4.0 and above offer another way to start asynchronous processes in a subshell. This is done with the help of the coproc keyword.”
1. [CNCF Tools Overview: Are You Cloud-Native?](https://epsagon.com/blog/cncf-tools-overview-are-you-cloud-native/)

    
1. [Why I no longer use Terraform for Templating Kubernetes](https://medium.com/@stobiewankenobi/why-i-no-longer-use-terraform-for-templating-kubernetes-9aef37741447)

     Helm is the apparent solution but Helm leaves a lot to be desired.
1. [Service Mesh Interface with Lachlan Evenson](https://softwareengineeringdaily.com/2019/06/06/service-mesh-interface-with-lachlan-evenson/)

     ”Whichever container deployment system you choose, your application and its multiple servers need a way to route traffic, measure telemetry, and configure security policy. A service mesh abstraction can help serve these use cases.”
1. [Generate new repositories with repository templates](https://github.blog/2019-06-06-generate-new-repositories-with-repository-templates/)

     ”[R]epository templates to make boilerplate code management and distribution a first-class citizen on GitHub.”
1. [Why Red Hat is investing in CRI-O and Podman](https://www.redhat.com/en/blog/why-red-hat-investing-cri-o-and-podman)

    
1. [12 Common Tools for Your DevOps Team](https://www.tripwire.com/state-of-security/devops/common-tools-for-your-devops-team/)

     Ansible is mentioned with venerable brands and projects.
1. [Why does macOS Catalina use Zsh instead of Bash? Licensing](https://thenextweb.com/dd/2019/06/04/why-does-macos-catalina-use-zsh-instead-of-bash-licensing/)

     Why the sudden change? In a word: licensing.
1. [Kubernetes basics: Learn how to drive first](https://opensource.com/article/19/6/kubernetes-basics)

     ”Quit focusing on new projects and get focused on getting your Kubernetes dump truck commercial driver’s license.”
1. [59 Linux Networking commands and scripts](https://haydenjames.io/linux-networking-commands-scripts/)

    
1. [bitfield/script](https://github.com/bitfield/script)

     Making it easy to write shell-like scripts in Go
1. [majkinetor/au](https://github.com/majkinetor/au)

     Chocolatey Automatic Package Updater Module
1. [CoolerVoid/HiddenWall](https://github.com/CoolerVoid/HiddenWall)

     Tool to generate a Linux kernel module for custom rules with Netfilter hooking. (block ports, Hidden mode, rootkit functions etc)

### [ << Prev ](devopsweekly-130.md) ------------- [ Next >> ](devopsweekly-132.md)