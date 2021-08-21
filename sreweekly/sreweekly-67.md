## [SRE Weekly Issue #67](https://sreweekly.com/sre-weekly-issue-67/) - April 9, 2017
### Articles

1. [Risky Business Requires Active Operators](https://blog.skyliner.io/risky-business-requires-active-operators-9debbb082995#.a31cxs9ul)

    This article is about the risks of automation. While automation can reduce risk by making errors less likely, it also disengages human operators from what’s actually happening, meaning that they’re less likely to catch and correct problems.
1. [Things I Learned Managing Site Reliability for Some of the World’s Busiest Gambling Sites](https://zwischenzugs.wordpress.com/2017/04/04/things-i-learned-managing-site-reliability-for-some-of-the-worlds-busiest-gambling-sites/)

    The author spent seven months sifting through, categorizing, and documenting over 1700 production incidents. The result was impressive: a massive improvement in the SRE team’s incident response process and documentation. It’s got me wondering if we can do something similar at $JOB.Thanks to Steven Farlie for this one.
1. [Fault Tolerance in a High Volume, Distributed System](http://techblog.netflix.com/2012/02/fault-tolerance-in-high-volume.html)

    A danger of a microservice architecture is that one failing service can affect those that depend on it, even indirectly. The Netflix API handles over 10000 requests per second, and it was carefully designed to avoid the case where a slow dependency breaks unrelated requests.
1. [Human Factors at The Fringe: Nuclear Family](https://humanisticsystems.com/2016/09/14/human-factors-at-the-fringe-nuclear-family/)

    Nuclear Family is an interactive play in which the audience is presented with critical decisions as the characters move inexorably toward a nuclear plant disaster. The goal is to demonstrate local rationality, the principal that people make the best decision they can with the information they have at hand — even if in retrospect that decision led to an adverse outcome.
1. [Owning Your Code is Better](https://www.pagerduty.com/blog/developers-own-code/)

    Last year, PagerDuty moved toward giving developers operational responsibilty for the systems they create. The really cool thing about their transition is that they have hard stats on reduction of incidents, decrease in MTTR, and increase in changes deployed to production.
1. [Making On-Call as Painless as Possible – PagerDuty](https://www.pagerduty.com/blog/developer-enhancements/)

    This post is primarily a new feature announcement, but the intro section is just awesome. I love the idea of designing a system with empathy for your future self that will be on call for it.
1. [Graceful degradation](http://sethgodin.typepad.com/seths_blog/2016/08/graceful-degradation.html)

    A short but enlightening blog post on designing systems to degrade gracefully.
### Outages

1. [Razer](http://vrzone.com/articles/razers-servers-go-second-time-month-resetting-mice-keyboards-globally/124761.html)
    Notably, this outage reset the careful customizations that people had made to their peripherals.
Thanks to Steven Farlie for this one.
1. [Heroku](https://status.heroku.com/incidents/1091)
    Heroku had a 2-day long disruption that spanned 3 status site posts.
Full disclosure: Heroku is my employer.
1. [DigitalOcean](https://www.digitalocean.com/company/blog/update-on-the-april-5th-2017-outage/)
    DigitalOcean accidentally deleted their primary database, resulting in a ~5-hour outage.
A process performing automated testing was misconfigured using production credentials.

### [ << Prev ](sreweekly-66.md) ------------- [ Next >> ](sreweekly-68.md)