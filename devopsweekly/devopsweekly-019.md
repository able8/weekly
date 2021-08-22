## [DevOps'ish 019](https://devopsish.com/019) - Sun Apr 16, 2017

Happy Easter! Hope you are enjoying your <a href="https://devopsish.com/"><strong>DevOps</strong></a>‘ing selves. I have family in from the northern reaches of the US. Trying to optimize time this week so no “monologue”; let’s get it!

<a href="http://www.techrepublic.com/article/ex-facebook-engineers-launch-honeycomb-a-new-tool-for-your-debugging-nightmares/">Ex-Facebook engineers launch Honeycomb, a new tool for your debugging nightmares</a>: Friend of <a href="https://devopsish.com/">DevOps’ish</a>, Charity Majors, and her fellow team members have officially launched <a href="https://honeycomb.io/">Honeycomb.io</a>. Honeycomb.io is an impressive product. I ingested a week’s worth of nginx logs in seconds and had a beautiful dashboard up in running in as much time as it took to write this sentence. Dope!

<a href="https://www.redhat.com/en/about/press-releases/red-hat-delivers-advanced-network-automation-latest-version-ansible?sc_cid=7016000000127NJAAY">Ansible 2.3 has been released</a>! This release appears to focus on bringing more advocates into the Ansible fold as it enhances Windows and networking capabilities. I welcome the additional features as a fully unified management platform for all components in an enterprise has never really panned out.

<a href="https://blog.openebs.io/openebs-sprinting-ahead-0-2-released-28f5001deeaa">OpenEBS 0.2 was released this week</a>

Kelsey Hightower has assembled <a href="https://github.com/kelseyhightower/google-cloud-functions-go">an unofficial collection of tutorials and hacks for using Go with Google Cloud Functions</a>.

Kubernetes tooling firm, <a href="https://deis.com/blog/2017/deis-to-join-microsoft/">Deis, has been acquired by Microsoft</a>. I was not sure how to feel about this (Microsoft does not do open source well, in my experience). CNBC says Microsoft, “<a href="http://www.cnbc.com/2017/04/10/microsoft-acquires-deis-in-latest-bid-to-catch-aws.html">is racing to grab a bigger share of the market as companies rapidly move workloads out of their own data centers and into the cloud.</a>” I hope this works out for Deis.

<a href="https://motherboard.vice.com/en_us/article/your-governments-hacking-tools-are-not-safe">Your Government’s Hacking Tools Are Not Safe</a>: When governments say they want a backdoor remind them the CIA and NSA had their tools compromised for MONTHS without telling a soul (including the software vendors).

<a href="https://tools.ietf.org/html/rfc6844">DNS CAA (Certificate Authority Authorization) records</a> are something you need to familiarize yourself with. A CAA record allows domain owners to declare which certificate authorities are allowed to issue a certificate for a domain. <a href="https://ma.ttias.be/caa-checking-becomes-mandatory-ssltls-certificates/?utm_source=cronweekly.com">Mattias Geniar states, “As of September 2017, Certificate Authorities are obligated to check for CAA records and honor those preferences.”</a>

<a href="http://www.techrepublic.com/article/5-tips-for-securing-your-docker-containers/">5 tips for securing your Docker containers</a>

When running your own Kubernetes cluster on Raspberry Pis wouldn’t it be cool to have <a href="https://youtu.be/a7MX6ED2zVM">a visual indicator</a>?

<a href="https://blog.codedellemc.com/2017/04/12/vagrant-docker-mesos-kubernetes-persistent-storage/">Vagrant Up Docker, Mesos and Kubernetes with Persistent Storage</a>

<a href="https://www.brianchristner.io/test-driving-docker-function-as-a-service-faas/">Test Driving Docker Function As A Service (FaaS)</a>

Susan Fowler has landed at Stripe and is doing something very awesome. <a href="https://increment.com/">Increment</a> looks to be, “a software engineering magazine dedicated to providing <em>practical and useful insight into what effective teams are doing so that the rest of us can learn from them more quickly</em>.”

<a href="http://philippe.bourgau.net/a-seamless-way-to-keep-track-of-technical-debt-in-your-source-code/">A Seamless Way to Keep Track of Technical Debt in Your Source Code</a>

<a href="https://dave.cheney.net/2017/04/11/why-slack-is-inappropriate-for-open-source-communications">Dave Chaney went full on H.A.M. on Slack and how it is inappropriate for open source communications</a>. I agree with the sentiment but the problem is the barrier to entry with IRC is just slightly higher than Slack. Slack was released 25 years after IRC came on the scene and it is barely more user friendly. This is a problem that still needs solving in my opinion.

As <a href="../018/">mentioned last week</a>, I have been tinkering with macOS monitoring. Not much success in that arena until I looked at one of the Solarwinds’ families of products: <a href="https://www.librato.com/">Librato</a>. It’s a pretty awesome tool that I like and would use regardless of the company affiliation. To get an idea of how awesome Librato is check out their <a href="https://metrics.librato.com/s/public/r623c3itn?duration=21600&amp;end_time=1489696705&amp;intercom=true">Bay Area DMV wait times dashboard</a>.

<a href="https://blog.golang.org/developer-experience">Go has created the Developer eXperience Working Group (DXWG)</a>.

<a href="https://aphyr.com/">Kyle Kingsbury</a> did an <a href="https://usesthis.com/interviews/kyle.kingsbury/">interview with The Setup</a> that I found fascinating. His desktop is, “a ridiculous 48-way Xeon (2x E5–2697v2), with 128GB of ECC DDR3 and 11 TB of miscellaneous SSDs &amp; spinning rust.”

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Refreshment and Refurbishment

1. [Ex-Facebook engineers launch Honeycomb, a new tool for your debugging nightmaresDevOps’ishHoneycomb.io](http://www.techrepublic.com/article/ex-facebook-engineers-launch-honeycomb-a-new-tool-for-your-debugging-nightmares/)

     Friend of , Charity Majors, and her fellow team members have officially launched . Honeycomb.io is an impressive product. I ingested a week’s worth of nginx logs in seconds and had a beautiful dashboard up in running in as much time as it took to write this sentence. Dope!
1. [Ansible 2.3 has been released](https://www.redhat.com/en/about/press-releases/red-hat-delivers-advanced-network-automation-latest-version-ansible?sc_cid=7016000000127NJAAY)

    ! This release appears to focus on bringing more advocates into the Ansible fold as it enhances Windows and networking capabilities. I welcome the additional features as a fully unified management platform for all components in an enterprise has never really panned out.
1. [OpenEBS 0.2 was released this week](https://blog.openebs.io/openebs-sprinting-ahead-0-2-released-28f5001deeaa)

    
1. [an unofficial collection of tutorials and hacks for using Go with Google Cloud Functions](https://github.com/kelseyhightower/google-cloud-functions-go)

    Kelsey Hightower has assembled .
### Department of Assemblage Obtainment

1. [Deis, has been acquired by Microsoftis racing to grab a bigger share of the market as companies rapidly move workloads out of their own data centers and into the cloud.](https://deis.com/blog/2017/deis-to-join-microsoft/)

    Kubernetes tooling firm, . I was not sure how to feel about this (Microsoft does not do open source well, in my experience). CNBC says Microsoft, “” I hope this works out for Deis.
### Department of Data Defense

1. [Your Government’s Hacking Tools Are Not Safe](https://motherboard.vice.com/en_us/article/your-governments-hacking-tools-are-not-safe)

     When governments say they want a backdoor remind them the CIA and NSA had their tools compromised for MONTHS without telling a soul (including the software vendors).
1. [DNS CAA (Certificate Authority Authorization) recordsMattias Geniar states, “As of September 2017, Certificate Authorities are obligated to check for CAA records and honor those preferences.”](https://tools.ietf.org/html/rfc6844)

    are something you need to familiarize yourself with. A CAA record allows domain owners to declare which certificate authorities are allowed to issue a certificate for a domain.
1. [5 tips for securing your Docker containers](http://www.techrepublic.com/article/5-tips-for-securing-your-docker-containers/)

    
### Department of Happy Little Clouds

1. [a visual indicator](https://youtu.be/a7MX6ED2zVM)

    When running your own Kubernetes cluster on Raspberry Pis wouldn’t it be cool to have ?
1. [Vagrant Up Docker, Mesos and Kubernetes with Persistent Storage](https://blog.codedellemc.com/2017/04/12/vagrant-docker-mesos-kubernetes-persistent-storage/)

    
1. [Test Driving Docker Function As A Service (FaaS)](https://www.brianchristner.io/test-driving-docker-function-as-a-service-faas/)

    
### Department of Sane Workplaces

1. [Increment](https://increment.com/)

    Susan Fowler has landed at Stripe and is doing something very awesome.  looks to be, “a software engineering magazine dedicated to providing practical and useful insight into what effective teams are doing so that the rest of us can learn from them more quickly.”
1. [A Seamless Way to Keep Track of Technical Debt in Your Source Code](http://philippe.bourgau.net/a-seamless-way-to-keep-track-of-technical-debt-in-your-source-code/)

    
1. [Dave Chaney went full on H.A.M. on Slack and how it is inappropriate for open source communications](https://dave.cheney.net/2017/04/11/why-slack-is-inappropriate-for-open-source-communications)

    . I agree with the sentiment but the problem is the barrier to entry with IRC is just slightly higher than Slack. Slack was released 25 years after IRC came on the scene and it is barely more user friendly. This is a problem that still needs solving in my opinion.
### Department of Interior

1. [mentioned last weekLibratoBay Area DMV wait times dashboard](../018/)

    As , I have been tinkering with macOS monitoring. Not much success in that arena until I looked at one of the Solarwinds’ families of products: . It’s a pretty awesome tool that I like and would use regardless of the company affiliation. To get an idea of how awesome Librato is check out their .
### Not DevOps But Still Cool

1. [Go has created the Developer eXperience Working Group (DXWG)](https://blog.golang.org/developer-experience)

    .
1. [Kyle Kingsburyinterview with The Setup](https://aphyr.com/)

    did an  that I found fascinating. His desktop is, “a ridiculous 48-way Xeon (2x E5–2697v2), with 128GB of ECC DDR3 and 11 TB of miscellaneous SSDs & spinning rust.”
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-018.md) ------------- [ Next >> ](devopsweekly-020.md)