## [SRE Weekly Issue #209](https://sreweekly.com/sre-weekly-issue-209/) - March 2, 2020
### Articles

1. [Gandalf: an intelligent, end-to-end analytics service for safe deployment in cloud-scale infrastructure](https://blog.acolyer.org/2020/02/28/microsoft-gandalf/)

    Azure developed this tool to sniff out production problems caused by deploys and guess which deploy might have been the culprit. Its accuracy is impressive.Adrian Colyer — The Morning Paper (summary)Li et al. — NSDI’20 (original paper)
1. [fork() can fail: this is important](https://rachelbythebay.com/w/2014/08/19/fork/)

    This one made me laugh out loud.  Better check those system call return codes, people.rachelbythebay
1. [Managing the Hidden Costs of Coordination](https://queue.acm.org/detail.cfm?id=3380779)

    This caught my eye:Laura M.D. Maguire — ACM Queue Volume 17, Issue 6
1. [How much money do SREs make?](https://www.gremlin.com/site-reliability-engineering/how-much-money-do-sres-make/)

    A guide on salary expectations for various levels of SRE, especially useful if you’re changing jobs.Gremlin
1. [3 microservices resiliency patterns for better reliability](https://searchapparchitecture.techtarget.com/tip/3-microservices-resiliency-patterns-for-better-reliability)

    Joydip Kanjilal
1. [It’s time for smart home devices to have local failover options during cloud outages](https://staceyoniot.com/smart-home-devices-cloud-outage-vs-local/)

    There have been several recent failures of consumer devices based on a cloud service outage, and this author argues for change.Kevin C. Tofel — Stacey on IoT
1. [Human error, miscommunication and lack of training behind false alarm at Pickering nuclear station](https://www.durhamradionews.com/archives/124716)

    This sounds familiar…Durham Radio News
1. [Friday deploys: comfort, not pressure](https://blog.leftclick.com.au/2020/02/21/_20200221-friday-deploys-comfort-not-pressure/)

    Ben New
### Outages

1. [Fidelity](https://www.marketwatch.com/story/investors-unable-to-access-brokerage-accounts-as-stocks-fall-what-you-should-know-for-your-retirement-accounts-2020-02-24)
    This one was especially problematic because it happened on Monday, a day of huge losses for the US stock market.
1. [GitHubThis oneshort note](https://www.githubstatus.com/incidents/xp2qc958g4wt)
    This one too. GitHub posted a short note on the recent outages.
1. [TechCrunchthe certificate](https://twitter.com/TechCrunch/status/1232662882615091202?s=20)
    TechCrunch was serving an expired TLS certificate. The strange thing is that the certificate had only been valid for 12 hours.
1. [Petnet pet feeders](https://boingboing.net/2020/02/27/internet-of-starving-pets-ani.html)
1. [Google Nest](https://9to5google.com/2020/02/24/nest-cam-down/)

### [ << Prev ](sreweekly-208.md) ------------- [ Next >> ](sreweekly-210.md)