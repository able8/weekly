## [SRE Weekly Issue #40](https://sreweekly.com/sre-weekly-issue-40/) - September 18, 2016
### Articles

1. [On designing and deploying internet-scale services](https://blog.acolyer.org/2016/09/12/on-designing-and-deploying-internet-scale-services/)

    Adrian Colyer summarizes James Hamilton’s 2007 paper in this edition of The Morning Paper. There’s a lot of excellent advice here — some I knew explicitly, some I mostly implement without thinking about it, and some I’d never thought about. The paper is great, but even if you don’t have time to read it, Colyer’s digest version is well worth a browse.
1. [Embracing Failure](http://www.susanjfowler.com/blog/2016/8/24/embracing-failure)

    Susan Fowler (featured here a couple weeks ago) has a philosophy of failure in her life that I find really appealing as an SRE:
1. [Second Chapter of Production-Ready Microservices by Susan Fowler](http://www.susanjfowler.com/blog/2016/9/17/my-ebook-is-available-for-free-download)

    And while we’re on the subject of Susan Fowler, she’s got a book coming soon about writing reliable microservices. In the linked ebook-version of the second chapter, she goes over the requirements for a production-ready microservice: stability, reliability, scalability, fault-tolerance, catastrophe-preparedness, performance, monitoring, and documentation.
1. [Sharding Pinterest: How we scaled our MySQL fleet](https://engineering.pinterest.com/blog/sharding-pinterest-how-we-scaled-our-mysql-fleet)

    Pinterest explains how they broke their datastore up into 4096(!) shards on 4 pairs of MySQL servers (later 8192 on 8 pairs). It’s an interesting approach, although in essence it treats MySQL as a glorified key-value store for JSON documents.
1. [Scalable and secure access with SSH](https://code.facebook.com/posts/365787980419535/scalable-and-secure-access-with-ssh/)

    Do you use Kerberos or similar to authenticate your SSH users? What happens if there’s an incident that’s bad enough to take down your auth infrastructure? I hadn’t realized that openSSH supports CAs, but Facebook shows us that PKI support is easy and feature-rich.
1. [DHCPLB: An open source load balancer](https://code.facebook.com/posts/1734309626831603/dhcplb-an-open-source-load-balancer/)

    Another project from Facebook: a load balancer for DHCP. Facebook found that anycast was not distributing requests evenly across DHCP servers, so they wrote a loadbalancer in Go.
1. [Safety Moment – Fundamental Attribution Error](http://preaccidentpodcast.podbean.com/e/safety-moment-fundamental-attribution-error/)

    In incident post-analysis, a fundamental attribution error is a tendency to see flaws in others as a cause if they were involved in an incident, but to blame the system if we were the one involved. This 4-minute segment from the Pre-Accident Podcast explains fundamental attribution error in more detail.
1. [Introducing 411: A new open source framework for handling alerting](https://codeascraft.com/2016/09/15/introducing-411-a-new-open-source-framework-for-handling-alerting/)

    411 is Etsy’s new tool that runs scheduled queries against Elasticsearch and alerts on the result.
### Outages

1. [ING Bankhearing damageshouting at hard drives](http://www.ibtimes.co.uk/loud-noise-shut-down-ing-banks-main-data-centre-bucharest-10-hours-1580799)
    Here’s a terribly interesting root cause: during a test, the fire response system emitted an incredibly loud sound while dumping an inert gas into the datacenter — probably loud enough to cause hearing damage. This caused failure in multiple key spinning hard drives. Remember shouting at hard drives?
1. [Heroku Status](https://status.heroku.com/incidents/941)
    Heroku released a followup with details on last week’s outage.
Full disclosure: Heroku is my employer.
1. [Gmail for Work](http://www.cbronline.com/news/cloud/public/cloud-outages-hit-microsoft-and-google-5007386)
1. [Microsoft Azurereport](http://www.zdnet.com/article/global-dns-outage-hits-microsoft-azure-customers/)
    Major outage involving most DNS queries for Azure resources failing. Microsoft posted a report including a root cause analysis.

### [ << Prev ](sreweekly-39.md) ------------- [ Next >> ](sreweekly-41.md)