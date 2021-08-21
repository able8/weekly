## [SRE Weekly Issue #66](https://sreweekly.com/sre-weekly-issue-66/) - April 2, 2017
### Articles

1. [Sockets in a Bind: Troubleshooting Port Exhaustion in Heroku’s Routing Layer](https://engineering.heroku.com/blogs/2017-03-30-sockets-in-a-bind/)

    I hope you’ll enjoy reading this debug session as much as I enjoyed co-writing it. My former co-worker and I did some serious digging to get to the bottom of an unexpected EADDRINUSE that caused a production incident.Full disclosure: Heroku is my employer.
1. [Redundancy does not imply fault tolerance: analysis of distributed storage reactions to single errors and corruptions](https://blog.acolyer.org/2017/03/08/redundancy-does-not-imply-fault-tolerance-analysis-of-distributed-storage-reactions-to-single-errors-and-corruptions/)

    Distributed filesystems provide high availability by duplicating data. In this research paper, the researchers created errorfs, a FUSE plugin that passes through a backing filesystem but introduces single-bit errors. Result: almost all major distributed filesystems missed the error, resulting in corruption.
1. [How we implement Disaster Recovery and High Availability with Postgres on Citus Cloud](https://www.citusdata.com/blog/2017/03/23/a-look-into-disaster-recovery-and-high-availability-and-how-they-work-with-postgres-on-citus-cloud/)

    The part I like most about this article is the emphasis on the difference between DR and HA.Full disclosure: Heroku, my employer, is mentioned.
1. [Breaking Things on Purpose](https://blog.gremlininc.com/breaking-things-on-purpose-a519c0f5698b)

    The S3 outage a month ago is a great reminder that chaos experiments are useful not just for taking down parts of our own infrastructure, but also simulating the failure of external dependencies.
1. [HumanOps: It’s time to make DevOps personal](https://jaxenter.com/humanops-time-make-devops-personal-132949.html)

    
1. [Now Available – Videos from SREcon17](http://www.sitereliability.engineer/2017/03/30/srecon17-videos/)

    Impressively quickly, USENIX has posted the videos from SRECon17 Americas! I’ve linked to a post by Woodland Hunter, whose review of SRECon I featured here two weeks ago, with links to the talks he reviewed and more.
1. [April Foolishness](https://en.wikipedia.org/wiki/April_Fools%27_Day)

    The first article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [How Incident Management Boosts Employee Morale](https://www.pagerduty.com/blog/incident-management-employee-morale/)

    PagerDuty theorizes that if developers don’t trust the incident response process, they’ll fear outages and thus be less productive. Proper incident management eases that fear so that they feel safer deploying code.
1. [Reducing Alert Noise: Going from 1000 Alerts to 10 Alerts Overnight](https://victorops.com/blog/reducing-alert-noise-going-1000-alerts-10-alerts-overnight/)

    This article could be titled, “Use these three wacky tricks to reduce your pages by 100x!” In all seriousness, the methods described are aggregation (group related alerts), routing (sort alerts by team), and classification (page-worthy alerts versus warnings).This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Slow down your internet with tc](https://jvns.ca/blog/2017/04/01/slow-down-your-internet-with-tc/)

    A nice primer on using tc to induce latency, which is really important when testing the resiliency of systems to network instability. Thanks, Julia!
1. [Risk Tolerance of Services](https://medium.com/@jerub/risk-tolerance-of-services-3bdc8a61a065?source=rss-3ea3ed468e7c------2)

    Here’s the second half of Stephen Thorne’s commentary on “Embracing Risk”, the third chapter in Google’s SRE book.
1. [Scaling Incident Management](https://www.pagerduty.com/blog/scaling-incident-management/)

    As your company grows in infrastructure size, number of employees, load, and other areas, how do you change your incident response to cope?
### Outages

1. [Azure status history](https://azure.microsoft.com/en-us/status/history/)
    While following up on an outage from a couple of weeks ago, I came upon this archive of Azure incidents, several with detailed postmortems. It’s a goldmine of interesting RCAs, but I wish they’d give each its own page for easy linking.

### [ << Prev ](sreweekly-65.md) ------------- [ Next >> ](sreweekly-67.md)