## [SRE Weekly Issue #76](https://sreweekly.com/sre-weekly-issue-76/) - June 11, 2017
### Articles

1. [Uber Fires 20 Amid Investigation Into Workplace Culture](https://www.nytimes.com/2017/06/06/technology/uber-fired.html)

    In case you missed it, Uber kicked off this and another investigation in response to a blog post by Susan Fowler, an SRE whose writing I’ve featured here a number of times. I’m pleased at this first step by Uber and I’m looking forward to what comes next. It might be a leave of absence for Uber’s CEO, although no decision has been made yet.
1. [Jepsen: On the perils of network partitions](https://aphyr.com/posts/281-call-me-maybe-carly-rae-jepsen-and-the-perils-of-network-partitions)

    Here’s the 2013 article that started it all. If you’re unfamiliar with Jepsen, it’s an article series on testing various distributed data systems for partition tolerance, along with a companion tool set for inducing failures.
1. [Internet Routing and Traffic Engineering](https://www.awsarchitectureblog.com/2014/12/internet-routing.html)

    For those not completely “cloud native” (ugh) by this point, here’s a nifty primer on some of the BGP tricks you’ll need to know if you manage your own IP transit links.
1. [A Key Expired In Redis, You Won’t Believe What Happened Next](http://engineering.grab.com/a-key-expired-in-redis-you-wont-believe-what-happened-next)

    Redis has a pretty big gotcha regarding deletion of expired keys, as these engineers discovered. In fact, my experience with Redis was full of operational gotchas like this.
1. [Reddit – cscareerquestions – Accidentally destroyed production database on first day of a job, and was told to leave, […] how screwed am i?](https://amp.reddit.com/r/cscareerquestions/comments/6ez8ag/accidentally_destroyed_production_database_on/)

    This poor anonymous Reddit poster had a very bad day. The community rallied around them to explain that no, the anonymous poster is not to blame. One of the top commenters is Yorick Peterse, the engineer that inadvertently deleted GitLab.com’s main database earlier this year. Click through to see blamelessness in action.
1. [Top Skills for an Incident Commander](https://www.pagerduty.com/blog/top-skills-incident-commander/)

    PagerDuty is deeply invested in the Incident Management System, and most especially Incident Command. This article is a great overview, and if you want more, don’t forget that they also released their incident response documentation awhile back, including their Incident Commander training material.
1. [Four nines and beyond: A guide to high availability infrastructure](http://blog.statuspage.io/high-availability)

    The main theme in this article by StatusPage.io is the direct relationship between increasing complexity and difficulty in attaining high reliability. I like the mention of microservices as a trade-off and not a panacea.
1. [“Serverless and the the death of devops”. Can you not?](http://redmonk.com/jgovernor/2017/06/02/serverless-and-the-the-death-of-devops-can-you-not/)

    
### Outages

1. [Amazon product pages went down today in a rare outage](https://www.theverge.com/2017/6/7/15759046/amazon-product-pages-down-outage-offline-503)
    The linked story was for an outage on June 7th. There was at least one additional similar outage on June 9th (source: personal experience).
1. [Verelox](https://www.bleepingcomputer.com/news/security/ex-admin-deletes-all-customer-data-and-wipes-server-of-dutch-hosting-provider/)
    Dutch hosting provider Verelox is having a really rough time:
First of all, we want to offer our apologies for any inconvenience. Unfortunately, an ex administrator has deleted all customer data and wiped most servers.
Ouch. Good luck, folks.

### [ << Prev ](sreweekly-75.md) ------------- [ Next >> ](sreweekly-77.md)