## [SRE Weekly Issue #86](https://sreweekly.com/sre-weekly-issue-86/) - August 27, 2017
### Articles

1. [Testing in production: Yes, you can (and should)](https://opensource.com/article/17/8/testing-production)

    Charity Majors knocks one out of the park with this article on the importance of testing (safely) in production.
1. [Stepping Up to the Plate: A Story About Being On-Call – PagerDuty](https://www.pagerduty.com/blog/intern-insights-on-call-experience/)

    And speaking of baseball metaphors, here’s a PagerDuty engineer’s first-person account of shadowing on-call during an incident and the lessons she learned.
1. [Post-Incident Reviews Survey](https://www.surveymonkey.com/r/devops-retros)

    If you have time, please consider filling out this short survey on post-incident reviews (a.k.a. “retrospectives”) as part of a master’s thesis.
1. [A Primer on Automating Chaos – Gremlin, Inc.](https://blog.gremlininc.com/a-primer-on-automating-chaos-84ff4b053be0)

    Mathias Lafeldt of Gremlin Inc. gives us this tutorial on moving from hand-run chaos experiments to a fully automated chaos system.
1. [Post-Incident Reviews: Learning from Failure for Improved Incident Response – VictorOps](https://victorops.com/blog/post-incident-reviews-learning-from-failure-for-improved-incident-response/)

    Recently, Jason Hand’s new ebook, Post-Incident Reviews, was published. Here’s his summary of the key points in the first three chapters.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Operational Metrics and Alerts for Distributed Software Systems](http://www.corriganjc.net/2017/07/designing-operational-metrics-for.html?m=1)

    This article describes metrics in three main categories and explains how (and whether) to set up alerts for each kind.
1. [Cutting Alert Fatigue in Modern Ops – PagerDuty](https://www.pagerduty.com/blog/cutting-alert-fatigue-modern-ops/)

    Like the previous article, Ilan Rabinovitch of Datadog advocates for symptom-based monitoring and alerting. I like his concept of the improved “durability” of symptom-based alerting (as opposed to cause-based):
1. [Impermanence: The Single Root Cause – Production Ready](https://medium.com/production-ready/impermanence-the-single-root-cause-bd9ebadf1e8e)

    Our systems are always in flux, and this sometimes leads to failure. Mathias expands on this line of thinking to urge seeking to understand the many conditions that led to a failure, rather than a particular root cause.
1. [Per-metric rate limiting: How we protect our backend](https://blog.hostedgraphite.com/2017/08/15/per-metric-rate-limiting-how-we-protect-our-backend/)

    Hosted Graphite had a gnarly problem to solve: how to get information about overload conditions from the backend to the front end where throttling could be enacted.
### Outages

1. [Honeycomb](https://honeycomb.io/blog/2017/08/our-first-outage/)
    Honeycomb suffered their first major outage this week.  I’m impressed by how quickly they were able to diagnose and fix the problem, owing at least in part to their use of their own service during troubleshooting.
1. [PagerDuty](https://status.pagerduty.com/incidents/t1vwfnb8bkjj)
    Here’s a followup from PagerDuty on an incident in May caused by “unanticipated side-effects of a system-wide load test”.
1. [Botched Firmware Update Bricks Hundreds of Smart Door Locks](https://www.bleepingcomputer.com/news/hardware/botched-firmware-update-bricks-hundreds-of-smart-door-locks/)
1. [RCA for SYNQ dashboard login and registration outage on August 11th, 2017](https://medium.com/synqfm-ops/rca-for-synq-dashboard-login-and-registration-outage-on-august-11th-2017-d86d6dc1ed70?source=rss----c5d5bfbdd67b---4)
1. [DreamHost](https://www.dreamhoststatus.com/pages/history/575f0f606826303142000510)
    DreamHost suffered a couple of DDoS attacks this week.Thanks to an anonymous SRE Weekly reader for this one.
1. [Facebook](https://www.geo.tv/latest/155414-thousands-unable-to-access-facebook-after-unexplained-outage)
    Facebook had a couple of outages this week.

### [ << Prev ](sreweekly-85.md) ------------- [ Next >> ](sreweekly-87.md)