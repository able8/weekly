## [DevOps'ish 057](https://devopsish.com/057) - Sun Jan 7, 2018

Is anything going on in the InfoSec arena this week? I couldn’t tell. If you have been under a rock this week I have some news for you. There are two vulnerabilities in CPU designs that essentially put everything we thought we knew about computer security on its head. <a href="https://meltdownattack.com/">Meltdown and Spectre</a> are vulnerabilities affecting virtually all modern CPUs. I don’t think I could ever visualize a complete picture of all the vulnerable systems that are impacted. But, to put it in perspective, my first Windows PC was a used Dell something or other with a 486 DX2 66 MHz CPU I got back in 1995. If I still had it, I would not need to patch it. But, if you have a 120 MHz Intel Pentium CPU based system laying around, it’s impacted.

The blast radius of these vulnerabilities is massive. Considering that, I am making a new section of the newsletter this week dedicated to Meltdown and Spectre. The reason for this is twofold: 1) There’s so much information in this space it could be a newsletter edition all by itself. 2) You might have some fatigue from these vulnerabilities. I don’t want you to skip over other awesome things in the newsletter. Scroll down past Tools for the Meltdown and Spectre section. Stay vigilant, keep your eyes open for patches, and rest assured <a href="http://www.businessinsider.com/linus-torvalds-linux-inventor-is-furious-at-intel-2018-1"><strong>Linus is PISSED</strong></a>.

Two personal notes:

<a href="https://devopsish.us14.list-manage.com/track/click?u=631fcd11ad2a643d08035c221&amp;id=5a1471dfb5&amp;e=7cc492dc98"><strong>GoCD — Open Source Continuous Delivery Server</strong></a><br/>GoCD is a continuous delivery tool supporting modern infrastructure with elastic on-demand agents and cloud deployments. With GoCD, you can easily model, orchestrate and visualize complex workflows from end to end. <a href="https://devopsish.us14.list-manage.com/track/click?u=631fcd11ad2a643d08035c221&amp;id=3133731028&amp;e=7cc492dc98">It’s open source, free to use and download</a>. <em>SPONSORED</em>

### People

1. [Vincent Batts: An Open Source Career from KDE to OCI](https://thenewstack.io/vincent-batts-kde-oci/)

     Vincent is a friend, and an absolutely wonderful person. It’s awesome to see him get the recognition he deserves.
1. [20 years of the Open Source Initiative (OSI)](http://www.computerweekly.com/blog/Open-Source-Insider/20-years-of-Open-Source-Initiative-OSI)

     The ‘open source’ label itself was created at a strategy session held by members of the group that we now call the Open Source Initiative (OSI) on February 3rd, 1998 in Palo Alto, California USA.
1. [What Would Really Happen If Russia Attacked Undersea Internet Cables](https://www.wired.com/story/russia-undersea-internet-cables/)

    
1. [How and why we teach non-engineers to use GitHub at Thread](https://thread.engineering/teaching-non-engineers-how-to-contribute-code-2e85411ab464)

    
1. [Taking the Certified Kubernetes Administrator Exam](https://medium.com/@KevinHoffman/taking-the-certified-kubernetes-administrator-exam-eeab17d65476)

    
1. [Top 21 conferences for DevOps and sysadmins in 2018](https://www.hpe.com/us/en/insights/articles/2018/01/top-21-conferences-for-devops-and-sysadmins-in-2018.html)

    
1. [What I learned in 2017 Writing Go](https://www.commandercoriander.net/blog/2017/12/31/writing-go/)

    
1. [“Oh My God, This Is So F — -ed Up”: Inside Silicon Valley’s Secretive, Orgiastic Dark Side](https://www.vanityfair.com/news/2018/01/brotopia-silicon-valley-secretive-orgiastic-inner-sanctum)

    
1. [2017 SRE & DevOps Influencers](https://robhirschfeld.com/2018/01/01/2017-sre-devops-influencers/)

    
### Process

1. [Docker, Inc isn’t DeadDocker Inc is Dead](https://blog.iron.io/docker-inc-isnt-dead/)

     Dylan Stamat of iron.io responded to my  story. I’m not quite sure Dylan’s response is outright disproving anything I wrote (it might actually reinforce it) but, it’s interesting to see opposing opinions.
1. [The evolution of Fastly’s Open Source and Nonprofit Program: supporting an ethical and open internet](https://www.fastly.com/blog/evolution-fastlys-open-source-and-nonprofit-program-supporting-an-ethical-and-open-internet/)

    
1. [The Limitations of Chaos Engineering](https://medium.com/production-ready/the-limitations-of-chaos-engineering-2a74816c0df3)

     It’s evident that Chaos Engineering has become a technology trend, with more and more companies adopting it.
1. [Creative ways to encourage the integration of DevOps processes](http://www.theserverside.com/feature/Creative-ways-to-encourage-the-integration-of-DevOps-processes)

    
1. [Selecting a Cloud Provider](https://codeascraft.com/2018/01/04/selecting-a-cloud-provider/)

    by Etsy
1. [The future of DevOps is mastery of multi-cloud environments](https://opensource.com/article/18/1/future-devops)

    
1. [It’s 2018 and your Docker containers need to be secure](https://blog.cloudpassage.com/2018/01/02/2018-docker-containers-need-secure/)

    
### Tools

1. [The DevOps Glossary](https://caylent.com/devops-glossary/)

     Whether you’re new to the world of DevOps or a seasoned guru looking to brush up on pesky terminology, look no further. This glossary covers some of the core definitions you and your team need to know.
1. [Staging endpoint for ACME v2](https://community.letsencrypt.org/t/staging-endpoint-for-acme-v2/49605)

     The Let’s Encrypt wildcard certs are coming.
1. [7 systems engineering and operations trends to watch in 2018](https://www.oreilly.com/ideas/7-systems-engineering-and-operations-trends-to-watch-in-2018?cmp=tw-webops-confpro-info-vlca18_2018_trends)

    
1. [9 New Programming Languages To Learn In 2018](https://www.rankred.com/new-programming-languages-to-learn/)

    
1. [Top 5: Best of 2017, the future of DevOps, and more](https://opensource.com/article/18/1/top-5-january-5)

    
1. [Get Started with Spinnaker on Kubernetes](https://thenewstack.io/getting-started-spinnaker-kubernetes/)

     A walkthrough on how to run Spinnaker on Minikube.
1. [A Brief History of sed](https://blog.sourcerer.io/a-brief-history-of-sed-6eaf00302ed)

     Their story is interesting, not least because it can’t be told without mentioning many acknowledged giants of computer science. It’s especially interesting when you interpret it in the context of all the other emerging parts of the nascent UNIX ecosystem that were also in motion at the time.
1. [Kubernetes, OpenShift build hybrid cloud outside Silicon Valley](https://siliconangle.com/blog/2018/01/02/kubernetes-openshift-build-hybrid-cloud-outside-silicon-valley-kubecon/)

    
1. [These Kubernetes developments make the platform ripe to explode in 2018](https://www.techrepublic.com/article/these-kubernetes-developments-make-the-platform-ripe-to-explode-in-2018/)

    
1. [kubernetes-incubator/kube-arbitrator](https://github.com/kubernetes-incubator/kube-arbitrator)

     kube-arbitrator provides policy based resource sharing for a Kubernetes cluster.
1. [samoshkin/docker-letsencrypt-certgen](https://github.com/samoshkin/docker-letsencrypt-certgen)

     Docker image allowing to generate, renew, revoke RSA and/or ECDSA SSL certificates from LetsEncrypt CA using certbot and acme.sh clients in automated fashion.
1. [khenidak/dysk](https://github.com/khenidak/dysk)

     Dysk mounts Azure disks as Linux block devices directly on VMs without dependency on the host. Dysks can be used within Azure VMs or on-prem machines.
1. [alexellis/mine-with-docker](https://github.com/alexellis/mine-with-docker)

     This repository contains Docker images that lets you get from zero to mining in around 5 minutes on any Linux host anywhere.
### Meltdown and Spectre

1. [Intel’s CEO reportedly sold shares after the company already knew about massive security flaws](https://www.cnbc.com/2018/01/04/intel-ceo-reportedly-sold-shares-after-the-company-already-knew-about-massive-security-flaws.html)

    
1. [Nearly Every Computer Made Since 1995 Is Dangerously Flawed. Here’s What You Need to Know.I tech reviewed this article](http://nymag.com/selectall/2018/01/intel-chip-security-flaw-meltdown-spectre-what-to-know-explainer.html)

    ( before it was published)
1. [“Meltdown” and “Spectre”: Every modern processor has unfixable security flaws](https://arstechnica.com/gadgets/2018/01/meltdown-and-spectre-every-modern-processor-has-unfixable-security-flaws/)

    
1. [Apple Says All Macs, iPhones and iPads Exposed to Chip Security Flaws](https://www.bloomberg.com/news/articles/2018-01-05/apple-says-all-macs-iphones-ipads-exposed-to-chip-flaw)

    
1. [An Update on AMD Processor Security](https://www.amd.com/en/corporate/speculative-execution)

    
1. [Arm Processor Security Update](https://developer.arm.com/support/security-update)

    
1. [Processor Speculative Execution Research Disclosure](https://aws.amazon.com/security/security-bulletins/AWS-2018-013/)

    via AWS
1. [A collection of Meltdown/Spectre postings](https://lwn.net/Articles/742999/)

    via LWN.net
1. [Addressing Meltdown and Spectre in the kernel](https://lwn.net/SubscriberLink/743265/df1eea5a556de4d4/)

    via LWN.net
1. [Guide to Meltdown / Spectre CPU Vulnerabilities](http://help.packet.net/technical/infrastructure/guide-to-meltdown-spectre-cpu-vulnerabilities)

    via Packet
1. [Why Raspberry Pi isn’t vulnerable to Spectre or Meltdown](https://www.raspberrypi.org/blog/why-raspberry-pi-isnt-vulnerable-to-spectre-or-meltdown/)

     “Both vulnerabilities exploit performance features (caching and speculative execution) common to many modern processors to leak data via a so-called side-channel attack. Happily, the Raspberry Pi isn’t susceptible to these vulnerabilities, because of the particular ARM cores that we use.”
1. [How a researcher hacked his own computer and found ‘worst’ chip flaw](http://www.reuters.com/article/us-cyber-intel-researcher/how-a-researcher-hacked-his-own-computer-and-found-worst-chip-flaw-idUSKBN1ET1ZR)

    
1. [Intel Issues Updates to Protect Systems from Security Exploits](https://newsroom.intel.com/news-releases/intel-issues-updates-protect-systems-security-exploits/)

    
1. [Mitigations landing for new class of timing attack](https://blog.mozilla.org/security/2018/01/03/mitigations-landing-new-class-timing-attack/)

    via Mozilla
1. [Initial Benchmarks Of The Performance Impact Resulting From Linux’s x86 Security Changes](https://www.phoronix.com/scan.php?page=article&item=linux-415-x86pti&num=1)

    
1. [Intel facing multiple class action suits over chip security flaw](https://www.theverge.com/2018/1/5/16853732/intel-meltdown-spectre-cpu-vulnerability-class-action-suits)

     As you can imagine, Linus is not the only one pissed about Meltdown and Spectre.
1. [Why Intel x86 must die: Our cloud-centric future depends on open source chips](http://www.zdnet.com/article/why-intel-x86-must-die-our-cloud-centric-future-depends-on-open-source-chips-meltdown/)

    
1. [Speculative Execution Exploit Performance Impacts — Describing the performance impacts to security patches for CVE-2017–5754 CVE-2017–5753 and CVE-2017–5715](https://access.redhat.com/articles/3307751)

    
1. []()

    dig +short txt istheinternetonfire.com

### [ << Prev ](sreweekly-56.md) ------------- [ Next >> ](sreweekly-58.md)