## [SRE Weekly Issue #11](https://sreweekly.com/sre-weekly-issue-11/) - February 21, 2016
### Articles

1. [A Skeleton Key of Unknown Strength](http://dankaminsky.com/2016/02/20/skeleton/)

    The big scary news this week was the buffer overflow vulnerability in glibc’s getaddrinfo function.  Aside from the obvious impact of intrusions on availability, this bug requires us to roll virtually every production service in our fleets.  If we don’t have good procedures in place, now is when we find out with downtime.
1. [Unikernels are unfit for production](https://www.joyent.com/blog/unikernels-are-unfit-for-production)

    Bryan Cantrill, CTO of Joyent, really crystallized the gross feelings that have been rolling around in my mind with regard to unikernels.  I would point colleagues to this article if they suggested we should deploy unikernel applications.  He makes lots of good points, especially this one:

1. [Inside Atlassian: how our site reliability engineers do incident management](https://blogs.atlassian.com/2016/02/inside-atlassian-site-reliability-engineers-incident-management/)

    Atlassian dissects their response to a recent outage and in the process shares a lot of excellent detail on their incident response and SRE process.  I love that they’re using the Incident Commander system (though under a different name).  This could have (and probably has) come out of my mouth:

1. [Fatal German train crash caused by human error, prosecutor says](http://mobile.reuters.com/article/idUSKCN0VP23L)

    My heart goes out to those passengers hurt and killed and to their families,  but also to the controller that made the error.  There’s a lot to investigate here about how a single human was in such a position that a single error could caus such devastation.  Hopefully there are ways in which the system can be remediated to prevent such catastrophes in the future.
1. [Steam Gauges Are Safer](http://www.avweb.com/news/features/Steam-Gauges-Are-Safer-225682-1.html)

    Like medicine, we can learn a lot about how to prevent and deal with errors from the hard lessons learned in aviation.

1. [Telstra customer chews through 425GB of mobile data during free data Sunday](http://www.smh.com.au/digital-life/mobiles/telstra-customer-chews-through-425gb-of-mobile-data-during-free-data-sunday-20160215-gmuz72.html)

    When I saw Telstra offer a day of free data to its customers to make up for last week’s outage, I cringed.  I’m impressed that they survived last Sunday as Australia used 1.8 petabtytes of data.
1. [The Power Of Paranoia](https://dzone.com/articles/the-power-of-paranoia)

    In this article, the author describes discovering that a service he previously ignored and assumed saw very little traffic actually served a million requests per day.

1. [Chaos Engineering 101 — Production Ready — Medium](https://medium.com/production-ready/chaos-engineering-101-1103059fae44#.eu2zq4glt)

    This is a great intro to Chaos Engineering, which is the field I didn’t know existed that was born out of Netflix’s Chaos Monkey.  This is the first article in what the author promises will be a biweekly series.Thanks to Devops Weekly for this one.
### Outages

1. [MTN (South Africa telecom co.)](http://www.techcentral.co.za/mtn-to-reimburse-customers-after-outage/63269/)
1. [Comcast](http://www.wmur.com/money/comcast-services-restored-after-second-major-outage-this-week/38041622)
    Two sets of sporadic outages for Comcast customers.  This one took out my home internet access a couple of times.
1. [Commuter Rail (Boston-area transit system)](http://www.bostonglobe.com/metro/2016/02/19/computer-issues-blamed-for-commuter-rail-problems/63TqSZ8FwgrzBT8CpvAPBJ/story.html)
1. [Xbox Live](http://www.express.co.uk/entertainment/gaming/644113/Xbox-Live-down-Xbox-One-not-working-Xbox-status-Microsoft)
1. [Street Fighter 5](http://www.ibtimes.co.uk/street-fighter-5-suffers-launch-day-server-problems-ps4-pc-1544112)
1. [Adobe Creative Cloud](http://cloudtweaks.com/2016/02/disruptions-adobe-creative-cloud/)

### [ << Prev ](sreweekly-10.md) ------------- [ Next >> ](sreweekly-12.md)