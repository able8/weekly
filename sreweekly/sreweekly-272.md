## [SRE Weekly Issue #272](https://sreweekly.com/sre-weekly-issue-272/) - May 30, 2021
### Articles

1. [[Salesforce] Multi-Instance Service Disruption on May 11-12, 2021](https://help.salesforce.com/articleView?id=000358392&type=1&mode=1)

    Salesforce has posted a ton of information about their major outage two weeks ago.
It involved a change to their DNS system that combined with an issue in BIND daemon shutdown that prevented it from starting back up.The analysis goes into great detail on the fact that an engineer used the Emergency Break-Fix (EBF) process to rush out the DNS configuration change.Thanks to an anonymous reader for pointing this out to me.Salesforce
1. [That Salesforce outage: Global DNS downfall started by one engineer trying a quick fix](https://www.theregister.com/AMP/2021/05/19/salesforce_root_cause/?__twitter_impression=true)

    This article calls out the heavily blame-ridden language in the above incident analysis and the briefing given by Salesforce’s Chief Reliability Officer.I’m dismayed to see such language from someone who is at the C-level for reliability.Richard Speed — The Register
1. [@ReinH on Twitter Re: Salesforce Outage](https://mobile.twitter.com/ReinH/status/1395906200210837510)

    …and the Twittersphere agrees with me.@ReinH on Twitter
1. [Subverting the process](https://surfingcomplexity.blog/2021/05/25/subverting-the-process/)

    Another really great take on the Salesforce outage followup.Lorin Hochstein
1. [Building an SRE Team? How to Hire, Assess, & Manage SREs](https://www.blameless.com/blog/sre-team)

    I like how this article covers the different roles that SREs play.Emily Arnott — Blameless
1. [The Advanced Principles of Chaos Engineering](https://www.verica.io/blog/the-advanced-principles-of-chaos-engineering/)

    The principles covered in this article are:Casey Rosenthal — Verica
1. [Why do config changes keep coming up in major incidents?](https://surfingcomplexity.blog/2021/05/29/why-do-config-changes-keep-coming-up-in-major-incidents/)

    This post is full of thought-provoking questions on the nature of configuration changes and incidents.Lorin Hochstein
### Outages

1. [IBM Cloud](https://www.datacenterdynamics.com/en/news/ibm-cloud-suffers-second-outage-in-five-days/)
1. [Klarna](https://www.klarna.com/us/blog/written-statement-on-app-bug/)
    Klarna showed users information related to other users, as detailed in this followup post.

### [ << Prev ](sreweekly-271.md) ------------- [ Next >> ](sreweekly-273.md)