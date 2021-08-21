## [SRE Weekly Issue #92](https://sreweekly.com/sre-weekly-issue-92/) - October 8, 2017
### Articles

1. [The Stella Report](http://stella.report)

    This is the blockbuster PDF dropped by the SNAFUcatchers during their keynote on day two of Velocity. Even just the 15-minute summary by Richard Cook and David Woods had me on the edge of my seat. In this report, they summarize the lessons gleaned from presentations of “SNAFUs” by several companies during winter storm Stella.SNAFUs are anomalous situations that would have turned into outages were it not for the actions taken by incident responders. Woods et al. introduced a couple of concepts that are new to me: “dark debt” and “blameless versus sanctionless”. I love these ideas and can’t wait to read more.
1. [IT incident response ditches root cause analysis process](http://searchitoperations.techtarget.com/news/450427746/IT-incident-response-ditches-root-cause-analysis-process)

    
1. [Chaos engineering unearths IT deployments’ dark debt](http://searchitoperations.techtarget.com/news/450427774/Chaos-engineering-unearths-IT-deployments-dark-debt)

    These two articles provide a pretty good round-up of the ideas shared at Velocity this past week.
1. [The Coming Software Apocalypse](https://www.theatlantic.com/technology/archive/2017/09/saving-the-world-from-code/540393/)

    This one starts with a 6-hour 911 (emergency services) outage in 2014 and the Toyota unintended acceleration incidents, and then vaults off into really awesome territory. Research is being done into new paradigms of software development that leave the programming to computers, focusing instead on describing behavior using a declarative language. The goal: provably correct systems. Long read, but well worth it.
1. [The Value of Optimizing for Resilience](https://dzone.com/articles/the-value-of-optimising-for-resilience)

    Drawing from Woods, Allspaw, Snowden, and others, this article explains how and why to improve the resilience of a system. There’s a great hypothetical example of graceful degradation that really clarified it for me.
1. [Nines don’t matter T-Shirt](https://www.zazzle.com/nines_dont_matter_t_shirt-235118578582589495)

    In a recent talk, Charity Majors made waves by saying, “Nines don’t matter when users aren’t happy.” Look, you can have that in t-shirt and mug format!
1. [Beta Testing in Production Like a Pro](https://dzone.com/articles/beta-testing-with-feature-toggles-testing-in-produ)

    A summary of how six big-name companies test new functionality by gradually rolling it out in production.
1. [How New, Resilient Networks Change Data Center Design](http://www.datacenterknowledge.com/design/will-data-center-operators-shed-tier-favor-distributed-resiliency)

    This article jumps off from Azure’s announcement of availability zones to discuss a growing trend in datacenters. We’re moving away from highly reliable “tier 4” datacenters and pushing more of the responsibility for reliability to software and networks.
1. [Ever wanted to know how Xero does incident management?](https://itbrief.co.nz/story/ever-wanted-know-how-xero-does-incident-management/)

    Of course I do, and I don’t even know who Xero is! They use chat, chatops, and Incident Command, like a lot of other shops. I find it interesting that incident response starts off with someone filling out a form.
### Outages

1. [PagerDuty](https://status.pagerduty.com/incidents/v2vrgccbtgxn)
    PagerDuty posted a lengthy followup report on their outage on September 19-21. TL;DR: Cassandra. It was the worst kind of incident, in which they had to spin up an entirely new cluster and develop, test, and enact a novel cut-over procedure. Ouch.
1. [Heroku#1298#1301
](https://status.heroku.com/incidents/1290)
    Heroku suffered a few significant outages. The one linked above includes a followup that describes a memory leak in their request routing layer. These two don’t yet have followups: #1298, #1301
Full disclosure: Heroku is my employer.
1. [Azure](https://azure.microsoft.com/en-us/status/history/)
    On September 29, Azure suffered a 7-hour outage in Northern Europe. They’ve released a preliminary followup that describes an accidental release of fire suppression agent and the resulting carnage. Microsoft promises more detail by October 13.
Unfortunately can’t deep-link to this followup, so just scroll down to 9/29.
1. [New Relic](https://status.newrelic.com/incidents/rd5gn27v3tb7)
1. [Blackboard (education web platform)](http://thelinc.co.uk/2017/10/blackboard-offline-again-just-days-after-it-claim-technical-issues-were-resolved/)

### [ << Prev ](sreweekly-91.md) ------------- [ Next >> ](sreweekly-93.md)