## [SRE Weekly Issue #127](https://sreweekly.com/sre-weekly-issue-127/) - June 24, 2018
### Articles

1. [Letter from Visa regarding service disruption, 15 June 2018](https://www.parliament.uk/documents/commons-committees/treasury/Correspondence/2017-19/visa-response-150618.pdf)

    Visa wrote a letter to the Chair of the Treasury Committee of the UK House of Commons, explaining their outage from a few weeks ago and answering the questions they posed. The good bits are in the first few pages, and the question answers mostly reiterate them. The last question about steps to prevent recurrence has some additional detail.Visa
1. [Introducing the Internet Intelligence Map](https://internetintel.oracle.com/blog-single.html?id=Introducing%20the%20Internet%20Intelligence%20Map)

    This is really nifty!
1. [Emily Nakashima on Twitter](https://mobile.twitter.com/eanakashima/status/1003751760622641156)

    Such an awesome idea:Emily Nakashima
1. [Communicating with twits: How to minimize friction between Dev and SRE](https://blog.hostedgraphite.com/2018/06/19/communicating-with-twits-how-to-minimize-friction-between-dev-and-sre/)

    How (and why) should an SRE team communicate with Dev and the rest of the organization? I especially enjoy the section on how communicating outwardly helps SRE.HostedGraphite
1. [Observability Show & Tell: show us your failure!](https://o11ycon.io/cff/)

    o11ycon has posted a Call for Failures:o11ycon
1. [MySQL High Availability at GitHub](https://githubengineering.com/mysql-high-availability-at-github/)

    There’s a great description of their current setup, but what really makes this article awesome is the explanation of what was wrong with their old system and why they replaced it.Shlomi Noach — GitHub
1. [Automated Database Deployments Iteration Zero](https://octopus.com/blog/automated-database-deployments-iteration-zero)

    Hilights of this article:Hen Peretz — BlazeMeter
1. [Just Culture & High Reliability: Steps to a More Reliable Organization](https://www.jems.com/ems-insider/articles/2018/june/just-culture-high-reliability-steps-to-a-more-reliable-organization.html)

    Rather than firing the driver that caused a rear-end collision, this company looked deeper and found an underlying flaw in their procedures. Larry Boxman and Paul LeSage — Journal of Emergency Medical Services
### Outages

1. [NPM (nodeJS package manager)Twitter bought an anti-harassment startup and immediately shut it downthis tweet](https://status.npmjs.org/incidents/l3wnv66v6rf1)
    This status posting is minimal, but there’s a deeper story at play here. There’s this article:
Twitter bought an anti-harassment startup and immediately shut it down
And this tweet by Laurie Voss (npmjs COO):
@seldo: A vendor notified us of their acquisition at 6am this morning and shut down their APIs 30 minutes later, creating a production outage for npm (package publishes and user registrations). The sheer unprofessionalism of this is blowing my mind.
Ouch.
1. [Datadog](https://status.datadoghq.com/incidents/6nh5v7h88t5q)
    These delays may result in “no data” alert conditions for Metric Monitors, to avoid spurious alerts we’ve temporarily disabled these alert types.
1. [DIRECTV NOW](http://www.telecompetitor.com/as-one-ott-service-suffers-outage-att-launches-new-watchtv-15-ott-service/)
    In the midst of suffering a major outage to their DIRECTV NOW OTT service, AT&T announced the official launch of AT&T WatchTV […]
1. [AlgeriaInternet Intelligence project](https://www.cnn.com/2018/06/21/africa/algeria-turns-off-internet-intl/index.html)
    Algeria switched off its internet on Wednesday in an attempt to prevent cheating on exams.
Algeria’s blackout can be seen in Oracle’s Internet Intelligence project, which maps web access globally.
Rory Smith — CNN
1. [Atlassian Statuspage (statuspage.io)](https://metastatuspage.com/incidents/mjxvtrscggbt)
    We have identified the issue as errant traffic from a single customer and have taken action to mitigate the issue, which appears to only affect status pages. The Management Portal is working as normal.
1. [New Relic](https://status.newrelic.com/incidents/lfwhprjxb7kz)
1. [GCP Networking in us-east1](http://status.cloud.google.com/incident/compute/18006#18006003)
1. [Azure North Europe region](https://azure.microsoft.com/en-us/status/history/)
    An environment control system failure caused a huge rise in humidity, taking down some equipment. Huge shout-out to the Microsoft employee who reached out to me to let me know that they saw my call for help last week and forwarded it on to the folks responsible for the status page!

### [ << Prev ](sreweekly-126.md) ------------- [ Next >> ](sreweekly-128.md)