## [DevOps'ish 200](https://devopsish.com/200) - Sun Jan 10, 2021

The first full work week of the year has already been filled with news. But, Monday saw a Slack outage, Wednesday saw an insurrection in the US, and there is a new twist in the <a href="https://devopsish.com/solarwinds-supply-chain-compromise/">Solarwinds supply chain compromise</a>. We’ll discuss two of these topics and more.

Note: I’m looking for an intern this summer to help with <a href="https://OpenShift.tv">OpenShift.tv</a> (live streaming). If you know anyone that may be interested, please ask them to apply. If they have questions, feel free to send them my way: <a href="https://us-redhat.icims.com/jobs/83032/openshift.tv-associate-producer-internship/job">have them apply</a>.

### People

1. [Linus Torvalds tears into Intel, favors AMD](https://www.zdnet.com/article/linus-torvalds-tears-into-intel-favors-amd/)

    After Spectre/Meltdown, I can’t blame anyone for being mad at Intel.
1. [What Buddhism can do for AI ethics](https://www.technologyreview.com/2021/01/06/1015779/what-buddhism-can-do-ai-ethics/)

    “Buddhism teaches us to focus our energy on eliminating suffering in the world.”
1. [Jobs report shows 140,000 jobs were lost in December. All of them were held by women](https://www.cnn.com/2021/01/08/economy/women-job-losses-pandemic/index.html)

    “Digging deeper into the data also reveals a shocking gender gap: Women accounted for all the job losses, losing 156,000 jobs, while men gained 16,000.”
1. [Working from home: 30% say they would quit if forced back to office](https://www.usatoday.com/story/money/2021/01/05/jobs-home-29-professionals-would-quit-if-forced-go-back-office/4142830001/)

    If I was working 9 to 5 in an office then was working from home and it was working well for me, I wouldn’t want to go back either.
1. [Cybersecurity Is Not (Just) a Tech Problem](https://hbr.org/2021/01/cybersecurity-is-not-just-a-tech-problem)

    “Remote work during the pandemic has meant that organizations had to quickly ramp up their cybersecurity efforts. But securing remote work isn’t just the job of the IT team: Ultimately companies need to make security part of every job description. And the key ingredient to make that happen is trust.”
### Process

1. [Widely Used Software Company May Be Entry Point for Huge U.S. Hacking](https://www.nytimes.com/2021/01/06/us/politics/russia-cyber-hack.html)

    “Russian hackers may have piggybacked on a tool developed by JetBrains, which is based in the Czech Republic, to gain access to federal government and private sector systems in the United States.” They don’t call them Advanced Persistent Threats for no reason.
1. [Download today: Kubernetes security ebook - tips, tricks, best practicesDownload this ebook](https://security.stackrox.com/kubernetes-security-ebook-tips-tricks-best-practices.html?Source=DevOpsIsh&LSource=DevOpsIsh)

    The rapid adoption of Kubernetes to manage containerized workloads is driving great efficiencies in application development, deployment, and scalability. However, when security becomes an afterthought, you risk diminishing the greatest gain of containerization - agility.  to learn how to (1) build secure images and prevent untrusted/vulnerable code, (2) configure RBAC, network policies, and runtime privileges, (3) detect unauthorized runtime activity, and (4) secure your Kubernetes infrastructure components such as the API server. SPONSORED
1. [Cyberattacks spike globally as criminals target under-pressure hospitals](https://www.timesofisrael.com/cyberattacks-spike-globally-as-criminals-target-under-pressure-hospitals/)

    “Check Point Software cybersecurity researchers say cybercrime targeting healthcare organizations has surged 45% globally since November, 25% in Israel” Attacking hospitals?!? Come on… Either the healthcare industry needs an awakening or an organization of some sort is formed to fortify it.
1. [How Red Hat is extending Kubernetes-native security across the open hybrid cloud with StackRox](https://www.redhat.com/en/blog/how-red-hat-extending-kubernetes-native-security-across-open-hybrid-cloud-stackrox)

    “By bringing StackRox’s powerful Kubernetes-native security capabilities to OpenShift, we are reinforcing our commitment to deliver a holistic open hybrid cloud platform. We want to enable users to build, deploy and more securely run applications across every IT footprint.” I’m excited about this latest Red Hat acquisition.
1. [Internet traffic disruption caused by the Christmas Day bombing in Nashville](https://blog.cloudflare.com/internet-traffic-disruption-caused-by-the-christmas-day-bombing-in-nashville/)

    A terrifying reminder how fragile our infrastructure is.
### Tools

1. [Slack Outage](https://status.slack.com/2021-01/9ecc1bc75347b6d1)

    Bad routing strikes again.
1. []()

    We need your voice!
1. [Honeycomb.io](https://www.honeycomb.io/?&utm_source=devopsish&utm_medium=newsletter&utm_campaign=ad&utm_content=honeycomb-homepage-devopish)

    In partnership with the team at ClearPath Strategies,  is collecting insights for changes in software development and operation practices across our industry. How do you see the world and what your team is doing?
1. [Take the survey](https://clearpathstrategies.sjc1.qualtrics.com/jfe/form/SV_cMAECZ6jv5wmjrL?&utm_source=devopsish&utm_medium=newsletter&utm_campaign=ad&utm_keyword=&utm_content=software-production-excellence-survey-clearpath-devopsish&utm_adgroup=)

    for a chance to win $500 from Apple, HelloFresh, or Fender. SPONSORED
1. [The Level Up Hour (S1E20): Kubernetes and Docker Deprecation](https://www.twitch.tv/videos/863349828)

    Langdon White, Scott McCarty, and I sit down to about the dockershim deprecation in Kubernetes.
1. [Nvidia Releases Drivers to Fix Newly Discovered GPU Security Flaws](https://www.ign.com/articles/nvidia-releases-drivers-to-fix-newly-discovered-gpu-security-flaws)

    CVSS score of 8.4, if you’re running Nvidia, update now.
1. [Migrating Your Open Source Builds Off Of Travis CI](https://blog.earthly.dev/migrating-from-travis/)

    “Starting in early December, a mad dash has been underway to migrate open-source projects off of Travis CI. What happened and where should you move your project to?”
1. [Using Podman and Docker Compose](https://www.redhat.com/sysadmin/podman-docker-compose)

    “Podman 3.0 now supports Docker Compose to orchestrate containers.”
1. [Introduction to WebAssembly](https://www.edx.org/course/introduction-to-webassembly-runtime)

    “Get a solid foundation on the WebAssembly runtime and its capabilities, and learn how and why WebAssembly has succeeded in bringing new programming languages to the web, when others have failed.”
1. [URGENT: SECURITY: New maintainer is probably malicious · Issue #1263 · greatsuspender/thegreatsuspender](https://github.com/greatsuspender/thegreatsuspender/issues/1263)

    “TLDR: The old maintainer appears to have sold the extension to parties unknown, who have malicious intent to exploit the users of this extension in advertising fraud, tracking, and more. In v7.1.8 of the extension (published to the web store but NOT to GitHub), arbitrary code was executed from a remote server, which appeared to be used to commit a variety of tracking and fraud actions. After Microsoft removed it from Edge for malware, v7.1.9 was created without this code: however, it can be added back at any time, and the original maintainer remains in control.”
1. [lyft/clutch](https://github.com/lyft/clutch)

    Extensible platform for infrastructure management
1. [jpillora/chisel](https://github.com/jpillora/chisel)

    A fast TCP/UDP tunnel over HTTP

### [ << Prev ](sreweekly-199.md) ------------- [ Next >> ](sreweekly-201.md)