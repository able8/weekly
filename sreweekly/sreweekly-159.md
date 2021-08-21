## [SRE Weekly Issue #159](https://sreweekly.com/sre-weekly-issue-159/) - February 10, 2019
### Articles

1. [Ironies of Automation](https://www.ise.ncsu.edu/wp-content/uploads/2017/02/Bainbridge_1983_Automatica.pdf)

    My favorite bit of irony: presenting data to the user in the manner most readily understood results in lower likelihood of remembering the data, so perhaps the most easily grasped display is not actually the best!Lisanne Bainbridge
1. [Laziness Does Not Exist](https://medium.com/@devonprice/laziness-does-not-exist-3af27e312d01)

    Like malice and incompetence, laziness should be far off our radar when we investigate an incident. I hope that reading this article opens minds about the true scope of blamelessness.Devon Price
1. [The New Systems Engineer](https://mattouille.com/articles/2019-01/the-new-systems-engineer/)

    Whether or not you agree with this particular attempt at defining what a Systems Engineer (or SRE or anything related) is, it’s worth thinking about and discussing. Our field is evolving quickly, and titles are a moving target.Matt Ouille
1. [Behind the Lion Air Crash, a Trail of Decisions Kept Pilots in the Dark](https://www.nytimes.com/2019/02/03/world/asia/lion-air-plane-crash-pilots.html)

    Driven by a desire to update their 737 without causing airlines to have to retrain pilots, Boeing seemingly kept pilots in the dark about what may have been an important little detail of how the new 737 Max operates, with a tragic result.James Glanz, Julie Creswell, Thomas Kaplan and Zach Wichter — New York Times
1. [Questions for a new technology.](https://kellanem.com/notes/new-tech)

    An experienced SRE will develop an innate skepticism of new technologies, even if they don’t realize it. This article provides an excellent list of questions to help articulate that skepticism when evaluating a potential design.Kellan Elliott-McCrea
1. [When AWS Autoscale Doesn’t](https://segment.com/blog/when-aws-autoscale-doesn-t/)

    Auto-scaling isn’t all roses. Like any tool, you have to understand how it works in order to avoid the pitfalls. Read this article to learn what these folks learned the hard way.Tyson Mote — Segment
1. [Postmortems Part 2: How to Adopt a Learning Culture](https://www.pagerduty.com/blog/postmortems-cultural-change/)

    Transitioning to a blameless culture can be difficult, especially as folks might blame each other for forgetting to be blameless!Rachael Byrne — PagerDuty
1. [Logs vs Structured Events](https://charity.wtf/2019/02/05/logs-vs-structured-events/)

    Many of the old arguments for not instrumenting code (mostly about performance) no longer apply, and a host of new arguments push toward structured events.Charity Majors
### Outages

1. [QuadrigaCXnote](https://www.bloomberg.com/news/articles/2019-02-04/crypto-exchange-founder-dies-leaves-behind-200-million-problem?fbclid=IwAR1-vukflPz-dQjxU55vpCFBzjQw_j3AWtqqnk8WdMOOWZO4i50TXV68uSA)
    Bloomberg’s title for the above-linked article says it all:
Crypto CEO Dies Holding Only Passwords That Can Unlock Millions in Customer Coins
QuadrigaCX ceased trading and posted a note on their front page.
1. [Gmail](https://www.google.com/appsstatus#hl=en&v=issue&sid=1&iid=366f3a7bd51af8441fd97852989b300f)
1. [Mailchimp Mandrill](https://twitter.com/mandrillapp/status/1092824328255741952)
    A PostgreSQL transaction ID wraparound in a central database caused this prolonged outage on Superbowl Sunday.
1. [Wells Fargo (bank)](https://www.americanbanker.com/news/flaws-in-testing-may-be-real-source-of-wells-fargos-tech-failure)
1. [Crunchyroll](https://www.express.co.uk/entertainment/gaming/1083511/Crunchyroll-down-Server-status-latest-anime-service-error)
1. [Hosted Graphite](https://status.hostedgraphite.com/incidents/zvkxt1vct10g)
1. [reddit](https://reddit.statuspage.io/incidents/fn9mlv6knwgw)

### [ << Prev ](sreweekly-158.md) ------------- [ Next >> ](sreweekly-160.md)