## [SRE Weekly Issue #9](https://sreweekly.com/sre-weekly-issue-9/) - February 7, 2016
### Articles

1. [January 28th Incident Report · GitHub](https://github.com/blog/2106-january-28th-incident-report)

    I spoke too soon in the last issue! Github has posted an extremely thorough postmortem that answers any questions one might have had about last week’s outage.  I like the standard they’re holding themselves to for timely communication:

1. [Monitoring Business Metrics](https://www.pagerduty.com/blog/monitoring-business-metrics/)

    Just monitoring servers isn’t enough to detect an outage.  Sometimes even detailed service monitoring can miss an overall performance degradation that involves multiple services in an infrastructure.  In this blog post, PagerDuty suggests also monitoring key business metrics (logins, purchase rate, etc).
1. [What happened yesterday and what we are doing about it](http://blog.mailgun.com/what-happened-yesterday-and-what-we-are-doing-about-it/)

    In this case, “yesterday” is on 2013, but this is an excellent postmortem from Mailgun that can serve as an example for all of us.
1. [Handling an Outage](http://qq.is/meta/2013/02/12/outage-handling.html)

    A customer’s perspective on a datacenter outage, with emphasis on the need for early, frequent, and thorough communication from service providers.
1. [Production Postmortem: the Razor Suicide](https://dzone.com/articles/production-postmortem-the-razor-suicide)

    A nicely detailed outage postmortem, including the gorey details of the train of thought the engineers followed on the way to a solution.  They hint at an important technique that’s not discussed nearly enough, in my opinion: judicious application of bandaid solutions to resolve the outage and allow engineers to continue their interrupted personal time.  It’s not necessary to fix a problem the “right” way in the moment, and carefully-applied bandaids help reduce on-call burnout.
1. [The Verification of a Distributed System](http://queue.acm.org/detail.cfm?id=2889274)

    How can we be sure (or at least sort of confident) that distributed systems won’t fail? They can be incredibly complex, and their failures can be even more complex.  Catie McCaffrey gives us this ACM Queue article about methods for formal and informal verification.

1. [Public Accountability — Postmortems — Medium](https://medium.com/postmortems/public-accountability-5935ed4d5e5a#.r89fqop0a)

    Medium has announced a commitment to publishing postmortems for all outages.  I’d love to see more companies making a commitment like this.  Thanks to reader Pete Shima for this link.
### Outages

1. [Healthplanfinder (WA, US)](http://m.heraldcourier.com/news/healthplanfinder-system-has-unscheduled-outage/article_5fcd46b6-8cf9-5ed5-b940-df6cf66f83cc.html?mode=jqm)
    The system went down right before the deadline for users to enroll in plans.
1. [Grindr](http://instinctmagazine.com/post/grindr-outage-spawns-hilarious-tweet-avalanche)
    The tweets during this outage were hilarious.
1. [Shaw (ISP)](http://calgaryherald.com/news/local-news/shaw-outage-hits-some-home-phone-lines-sunday)
1. [PlayStation Network](http://www.mirror.co.uk/news/technology-science/technology/psn-down-gamers-furious-sonys-7288525)
    The third outage this year for PSN.
1. [Virgin Australia](http://www.itnews.com.au/news/optus-outage-downs-virgin-australias-check-in-system-414501)
    Another airline grounded by an outage.
1. [British Telecom (UK ISP)](http://www.weupit.com/bt-network-suffers-major-internet-outage-across-uk/24367/)
1. [Amazon.com](http://www.ecommercebytes.com/C/blog/blog.pl?/pl/2016/2/1454530772.html)
1. [Google App Engine](http://status.cloud.google.com/incident/appengine/16001#5738886445662208)
1. [Delta](http://mobile.reuters.com/article/idUSKCN0VB242)
    What is it with airlines lately?
1. [Google Compute Engine](http://status.cloud.google.com/incident/compute/16002#5705718560718848)
    This one has a nice postmortem.
1. [IRS E-File (US tax system)](http://www.zdnet.com/article/irs-systems-hit-with-outage-as-tax-season-starts/)

### [ << Prev ](sreweekly-8.md) ------------- [ Next >> ](sreweekly-10.md)