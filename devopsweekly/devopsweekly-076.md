## [DevOps'ish 076](https://devopsish.com/076) - Sun May 20, 2018

This has been a whirlwind week. On Sunday, we celebrated Mother’s Day. It was a good day where I tried very hard to give Julie the day off from the rigors of Max management. I had to bail a little before Max’s bedtime though to head to Cleveland for an <a href="http://sjtechcorp.com/">SJ Technologies</a> client visit. It was a great visit but it was, in military lingo, a turn and burn. I was back in Michigan by Monday night. I was working on a couple of tickets this week before heading off to <a href="https://chefconf.chef.io/">ChefConf</a>. One issue revealed a lovely <a href="https://access.redhat.com/solutions/3391931">little bug in Kubernetes</a> (login required). I was attempting to mount a ConfigMap as a file in the root directory of a container and it was not working no matter what. It was awful and the workaround was thankfully easy enough due to the fact it was a web app and <a href="https://caddyserver.com/">Caddy</a> was available. I created an HTTP redirect and I was off and running. I submitted some additional bug reports where appropriate and was off to the next thing. Tripping over bugs is annoying. But, there’s a great Kubernetes community right here in Michigan: <a href="https://www.meetup.com/orchestructure/">Orchestructure</a>. Thanks to y’all for helping me out this week.

<a href="https://www.gocd.org/2018/04/25/five-considerations-continuous-delivery-microservices/?utm_campaign=cd_microservices&amp;utm_medium=newsletter_ad&amp;utm_source=devopsish&amp;utm_content=cd_microservices_blog&amp;utm_term="><strong>5 considerations for continuous delivery of microservices</strong></a><br/>Continuous Delivery is an essential component of any software delivery practice, including microservices. This new blog series shares five considerations we recommend keeping in mind when designing a CD workflow on microservices architectures. <em>SPONSORED</em>

<a href="https://chefconf.chef.io/"><strong>ChefConf 2018</strong></a><br/>Dates: 2018-05-22 through 2018-05-25<br/>Join an awesome community of #DevOps and automation professionals at <a href="https://chefconf.chef.io/">ChefConf</a> in Chicago. I’ll be presenting, <em>DevOps is Not a War</em>. Save 10% with discount code <strong>Hugs4Chef</strong>.

<a href="https://www.devopsdays.org/events/2018-toronto/welcome/"><strong>DevOpsDays Toronto 2018</strong></a><br/>Dates: 2018-05-30 through 2018-05-31<br/>I’ll admit it, I’ve never been to Canada. But, I’m definitely going to <a href="https://www.devopsdays.org/events/2018-toronto/welcome/">DevOpsDays Toronto</a> this year to present <em>What the Military Taught Me about DevOps</em>.

<a href="https://devnationfederal.org/"><strong>DevNation Federal</strong></a><br/>Date: 2018-06-05<br/>Join us to learn about the revolutions happening in communities around containers, data, and application modernization. This is an opportunity for you to hear how visionary teams in the federal government are innovating with open source, and hear from leaders in the private sector doing the same.

<a href="https://chaosconf.splashthat.com/"><strong>Chaos Conf</strong></a><br/>Date: 2018-09-28<br/>Chaos Conf looks super awesome. Opening the event will be Adrian Cockcroft, VP AWS, who called 2018 “The year of #chaosengineering”. Closing out the night will be Jessie Frazelle, one of the top #containers experts on the planet currently at Microsoft.

<a href="https://www.usenix.org/conference/lisa18"><strong>LISA18</strong></a><br/>Dates: 2018-10-29 through 2018-10-31<br/>Have something to say on the present &amp; future of #ops? The LISA18 CFP closes May 24. <a href="https://www.usenix.org/blog/usenix-lisa18-cfp-nashville">Submit your ideas</a>!

### Events

1. [ChefConf 2018ChefConf](https://chefconf.chef.io/)

    Dates: 2018-05-22 through 2018-05-25Join an awesome community of #DevOps and automation professionals at  in Chicago. I’ll be presenting, DevOps is Not a War. Save 10% with discount code Hugs4Chef.
1. [DevOpsDays Toronto 2018DevOpsDays Toronto](https://www.devopsdays.org/events/2018-toronto/welcome/)

    Dates: 2018-05-30 through 2018-05-31I’ll admit it, I’ve never been to Canada. But, I’m definitely going to  this year to present What the Military Taught Me about DevOps.
1. [DevNation Federal](https://devnationfederal.org/)

    Date: 2018-06-05Join us to learn about the revolutions happening in communities around containers, data, and application modernization. This is an opportunity for you to hear how visionary teams in the federal government are innovating with open source, and hear from leaders in the private sector doing the same.
1. [Chaos Conf](https://chaosconf.splashthat.com/)

    Date: 2018-09-28Chaos Conf looks super awesome. Opening the event will be Adrian Cockcroft, VP AWS, who called 2018 “The year of #chaosengineering”. Closing out the night will be Jessie Frazelle, one of the top #containers experts on the planet currently at Microsoft.
1. [LISA18Submit your ideas](https://www.usenix.org/conference/lisa18)

    Dates: 2018-10-29 through 2018-10-31Have something to say on the present & future of #ops? The LISA18 CFP closes May 24. !
### People

1. [The ultimate DevOps hiring guide](https://opensource.com/article/18/4/ultimate-devops-hiring-guide)

     Trying to round up DevOps talent? We’re a unique lot that requires some tweaks to your hiring practices.
1. [DevOps hiring strategies to attract top talent](https://opensource.com/article/18/5/devops-hiring-strategies-attract-top-talent)

     Top DevOps recruiter Ken Middleton offers insight on how to attract and hire the best candidates.
1. [#127 SRE vs Devops with Liz Fong-Jones and Seth Vargo](https://www.gcppodcast.com/post/episode-127-sre-vs-devops-with-liz-fong-jones-and-seth-vargo/)

     “Liz Fong-Jones and Seth Vargo battle it out on which is better: SRE or Devops (hint - everyone wins!).”
1. [Goodbye, Google (an open micro-letter)](https://medium.com/@sadams.codes/goodbye-google-b249cd513102)

     This made me very sad. I miss Sarah’s energy in our community.
1. [DevOps Culture through the Prism of Maslow’s Hierarchy of Needs](https://devops.com/devops-culture-through-the-prism-of-maslows-hierarchy-of-needs/)

    
1. [27 things I learned about hiring in tech from looking for a new engineering management role](http://wunder.schoenaberselten.com/2018/05/12/27-things-i-learned-about-hiring-in-tech-from-looking-for-a-new-engineering-management-role/)

    
### Process

1. [Not So Pretty: What You Need to Know About E-Fail and the PGP Flaw](https://www.eff.org/deeplinks/2018/05/not-so-pretty-what-you-need-know-about-e-fail-and-pgp-flaw-0)

     “Don’t panic! But you should stop using PGP for encrypted email and switch to a different secure communications method for now.”
1. [The DevOps Security Checklist](https://www.sqreen.io/checklists/devops-security-checklist.html)

     “This security checklist aims to give DevOps professionals a list of DevOps security best practices they can follow to implement DevSecOps.”
1. [This Week on The New Stack: Is a Kubernetes Backlash Afoot?](https://thenewstack.io/this-week-on-the-new-stack-kubernetes-needs-developers/)

    
1. [Red Hat and CoreOS put Kubernetes on autopilot](https://siliconangle.com/blog/2018/05/18/red-hat-and-coreos-put-kubernetes-on-autopilot-rhsummit/)

    
1. [How to start a Go project in 2018](https://boyter.org/posts/how-to-start-go-project-2018/)

     “Getting started with a Go project in 2018 is frankly a little more painful then getting anything else started IMHO.”
1. [Rackspace Launches Comprehensive Kubernetes-as-a-Service Solution with Fully Managed Operations](https://globenewswire.com/news-release/2018/05/16/1507403/0/en/Rackspace-Launches-Comprehensive-Kubernetes-as-a-Service-Solution-with-Fully-Managed-Operations.html)

    
1. [Tech giants offer startups free patents in bid to foil lawsuits](http://www.dailyherald.com/business/20180519/tech-giants-offer-startups-free-patents-in-bid-to-foil-lawsuits)

     “Red Hat Inc. and Lenovo Group Ltd. are giving away free patents to any start-up that joins a group of more than 200 companies devoted to keeping its members and their patents out of court.”
1. [Kubernetes Chaos Engineering: Lessons Learned](https://learnk8s.io/blog/kubernetes-chaos-engineering-lessons-learned)

    
1. [North Carolina, Apple negotiating deal on Triangle campus](https://www.wral.com/north-carolina-apple-negotiating-deal-on-triangle-campus-second-site-in-cary/17558985/)

     “Apple is close to announcing a deal that would bring as many as 10,000 jobs to North Carolina, including a major investment in the Research Triangle Park, according to multiple sources with knowledge of the deal.”
### Tools

1. [Edge computing and the importance of open infrastructure](https://opensource.com/article/18/5/edge-computing)

     The “edge” is diverse, dispersed, often independently owned and operated, and comes with a set of constraints not addressed in the average data center.
1. [Making Kubernetes work for the average engineer—via PaaS](https://www.infoworld.com/article/3273104/containers/making-kubernetes-work-for-the-average-engineer-via-paas.html)

     Despite being the hottest thing since, well, Docker, Kubernetes remains a dark art for most mainstream enterprises
1. [Citrix snuffs Xen and NetScaler brands](https://www.theregister.co.uk/2018/05/14/citrix_rebranding/)

     Arise, ‘Citrix Hypervisor’ and ‘Citrix SD-WAN’
1. [How HTTPS Works: Why Do We Need HTTPS?](https://howhttps.works/why-do-we-need-https/)

    
1. [Julia EvansBite size linux!](https://jvns.ca/)

    made another fantastic zine:
1. [The headers we don’t want](https://www.fastly.com/blog/headers-we-dont-want)

    by Fastly
1. [Windows Package Management with Ansible](https://www.ansible.com/blog/windows-package-management)

    
1. [Introducing Git protocol version 2](https://opensource.googleblog.com/2018/05/introducing-git-protocol-version-2.html)

     “This update removes one of the most inefficient parts of the Git protocol and fixes an extensibility bottleneck, unblocking the path to more wire protocol improvements in the future.”
1. [google/nomulus](https://github.com/google/nomulus)

     Top-level domain name registry service on Google App Engine
1. [Mayeu/awesome-open-source-organizations](https://github.com/Mayeu/awesome-open-source-organizations)

     A list of organizations that have open sourced everything they do
1. [gocolly/colly](https://github.com/gocolly/colly/releases/tag/v1.0.0)

     Congrats to Colly on going 1.0!
1. [goreleaser/nfpm](https://github.com/goreleaser/nfpm)

     NFPM is Not FPM - a simple deb and rpm packager written in Go
### Jobs

1. [SJ Technologieslet me knowKaren LawtonJohn WillisBarbara BouldinME](http://sjtechcorp.com/)

    is looking to bring someone on board to work in our DevOps and Digital Transformation practice. If you want to help some big-time companies eager to implement change . Plus, you get to work for great team of folks including , , , and .
1. [DZone](http://careers.dzone.com/apply/DBWe0hiNCN/Site-Reliability-Engineer)

    is looking for a Site Reliability Engineer in Cary, NC. Based off my experiences with DZone, this is a cool place to work. To quote a decision maker at DZone, “We need strong cloud, automation, security [experience].” Check them out.

### [ << Prev ](devopsweekly-075.md) ------------- [ Next >> ](devopsweekly-077.md)