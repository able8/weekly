## [SRE Weekly Issue #94](https://sreweekly.com/sre-weekly-issue-94/) - October 22, 2017
### Articles

1. [High Reliability Health Care: Getting There from Here](https://www.jointcommission.org/assets/1/6/Chassin_and_Loeb_0913_final.pdf)

    This article by the Joint Commission opened my eyes to just how far medicine in the US is from being a High Reliability Organization (HRO). It’s long, but I’m really glad I read it.
1. [Center stage: Best practices for staging environments](https://increment.com/development/center-stage-best-practices-for-staging-environments/)

    Increment issue #3 is out this week, and Alice Goldfuss gives us this juicy article on staging environments. I love the section on potential pitfalls with staging environments.
1. [Microservice Usage at Honeycomb](https://devops.com/microservice-usage-honeycomb/)

    A Honeycomb engineer gives us a deep-dive into Honeycomb’s infrastructure and shows how they use their product itself (in a separate, isolated installation) to debug problems in their production service. Microservices are key to allowing them to diagnose and fix problems.
1. [Tail-Tolerance by Google](https://dzone.com/articles/tail-tolerance-by-google)

    This is a nice summary of a paper by Google employees entitled, “The Tail at Scale”. 99th percentile behavior can really bite you if you’re composing microservices. The paper has some suggestions for how to deal with this.
1. [Focus on Analysis: The End of Root Cause](https://victorops.com/blog/focus-on-analysis-the-end-of-root-cause/)

    This post by VictorOps recommends moving away from Root Cause Analysis (RCA) toward a Cynefin-based method.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Open-sourcing RacerD: Fast static race detection at scale](https://code.facebook.com/posts/293371094514305)

    I love the idea of detecting race conditions through static analysis. It sounds hard, but the key is that RacerD seeks only to avoid false-positives, not false-negatives.
1. [ButterCMS Architecture: a Mission-Critical API Serving Millions of Requests per Month](http://highscalability.com/blog/2017/10/16/buttercms-architecture-a-mission-critical-api-serving-millio.html)

    Full disclosure: Heroku, my employer, is mentioned.
### Outages

1. [Honeycomb](https://honeycomb.io/blog/2017/10/bitten-by-a-kafka-bug---postmortem/)
    Honeycomb had a partial outage on the 17th due to a Kafka bug, and they posted an analysis the next day (nice!). They chronicle their discovery of a Kafka split-brain scenario through snapshots of the investigation they did using their dogfood instance of Honeycomb.
1. [Visual Studio Team Services](https://blogs.msdn.microsoft.com/vsoservice/?p=15276)
    Linked is an absolutely top-notch post-incident analysis by Microsoft. The bug involved is fascinating and their description had me on the edge of my seat (yes, I’m an incident nerd).
1. [Herokulink](https://status.heroku.com/incidents/1325)
    Heroku posted a followup for an outage in their API. Faulty rate-limiting logic prevented the service from surviving a flood of requests. Earlier in the week, they posted a followup for incident #1297 (link).Full disclosure: Heroku is my employer.

### [ << Prev ](sreweekly-93.md) ------------- [ Next >> ](sreweekly-95.md)