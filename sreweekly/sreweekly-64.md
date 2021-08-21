## [SRE Weekly Issue #64](https://sreweekly.com/sre-weekly-issue-64/) - March 19, 2017
### Articles

1. [SREcon17 Day 1 Highlights](http://www.sitereliability.engineer/2017/03/13/srecon17-day1/)

    I wasn’t able to make it to SRECon17 Americas this year, but it sounds like it was a great time. (day two summary)
1. [So you want to be a wizard](https://jvns.ca/blog/so-you-want-to-be-a-wizard/)

    My heroine, Julia Evans, gave the plenary session at SRECon17 Americas, all about how to learn how to be an excellent engineer (or really anything!). She proved herself once again not just as an excellent student, but also an inspiring teacher. The best part is that she posted the abstract, slides, and a transcript of her talk shortly after giving it! This is a really excellent resource for folks like me that weren’t there, and I hope more talk-givers will follow her example.
1. [Counterfactual Thinking, Rules, and The Knight Capital Accident](http://www.kitchensoap.com/2013/10/29/counterfactuals-knight-capital/)

    This article is long, but I wish I’d carved out time for it long ago, because it’s really incredible and well worth the read. John Allspaw uses the SEC analysis of the Knight Capital incident as a starting point to introduce and discuss the problems with Counterfactual Thinking (“if the engineer had just done ___, this wouldn’t have happened”).
1. [You Can’t Have a Rollback Button](https://blog.skyliner.io/you-cant-have-a-rollback-button-83e914f420d9#.n5yc446h5)

    Rolling back a flawed code release can have significant risk. It doesn’t always fix the problem because the erroneous code may have had effects on other parts of the system. Sometimes, as in the Knight Capital incident, a rollback exacerbates the problem.
1. [The Production Environment at Google](https://medium.com/@jerub/the-production-environment-at-google-8a1aaece3767?source=rss-3ea3ed468e7c------2)

    This is part two of an annotation of the google SRE book by Stephen Thorne, a Google SRE. Part Three is available too.
1. [Measuring Technical Debt With Incident Management Data](https://www.pagerduty.com/blog/technical-debt/)

    Here’s an interesting idea: using metadata about incidents as a proxy for measuring technical debt. PagerDuty goes over the definition of technical debt before diving into measuring it.
1. [How is team-member-1 doing?](https://about.gitlab.com/2017/03/17/how-is-team-member-1-doing/)

    GitLab posted an update on “team-member-1”, the engineer that entered the commands that caused their production DB to be erased. I love that they posted this, because I for one was worried about “team-member-1” as a second victim.
1. [U mad bro? Disaster planning for on-call – VictorOps](https://victorops.com/blog/u-mad-bro-disaster-planning/)

    During an incident, emotions can run strong. How can we set up incident response so as to provide the best environment for our responders?This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
### Outages

1. [AWS Route 53](http://status.aws.amazon.com/)
    Route 53 had a control plane outage, though actual query responses were unaffected.
1. [Square](https://medium.com/square-corner-blog/incident-summary-2017-03-16-2f65be39297#.fcxt7qyhs)
    Square suffered a 2-hour outage, and if this postmortem is any indication, they learned a lot from it. This bit is interesting in light of the article above about rollbacks:
We rolled back all software changes that happened leading up to the incident. This is a non-negotiable response to any customer-impacting event; our engineers are trained to undo any change that happened before an incident regardless of how plausible it is that the change caused the issue.
1. [StatusPage.io](http://metastatuspage.com/incidents/802s2lmr476r)
    This happened during Square’s outage and impacted their ability to communicate.
1. [CBS](http://ftw.usatoday.com/2017/03/ncaa-tournament-2017-brackets-march-madness-cbs-sports-website-down-twitter-reaction-freaking-out)
    CBS’s site was down, so people couldn’t fill out their fantasy sportsball brackets 1 hour before the game started.

### [ << Prev ](sreweekly-63.md) ------------- [ Next >> ](sreweekly-65.md)