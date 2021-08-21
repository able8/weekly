## [SRE Weekly Issue #95](https://sreweekly.com/sre-weekly-issue-95/) - October 29, 2017
### Articles

1. [Abstracting the Geniuses Away from Failure Testing](http://queue.acm.org/detail.cfm?id=3155114)

    Chaos Engineering and Jepsen-style testing is still in its infancy. As this ACM Queue article explains, figuring out what kind of failure to test is still a manual process involving building a mental model of the system. Can we automate it?
1. [Scaling the GitLab database](https://about.gitlab.com/2017/10/02/scaling-the-gitlab-database/)

    GitLab shares the story of how they implemented connection pooling and load balancing with read-only replicas in PostgreSQL.
1. [Moving Half a Million Database Tables to AWS Aurora (Part 1)](https://pressbooks.org/blog/2017/10/19/moving-half-a-million-database-tables-to-aws-aurora-part-1/?utm_source=last_week_in_AWS&utm_medium=newsletter&__s=gcxkayouhzyr45m1hboa)

    When you have 600,000(!!) tables in one MySQL Database, traditional migration tools like mysqldump or AWS’s Database Migration Service show cracks. The folks at PressBooks used a different tool instead: mydumper.
1. [Serverless availability zones are the missing level of resiliency for AWS](https://read.acloud.guru/serverless-availability-zones-are-the-missing-level-of-resiliency-for-aws-f5067ab1b688?__s=gcxkayouhzyr45m1hboa)

    AWS Lambda spans multiple availability zones in each region. This author wonders whether it would it be more reliable to have separate installations of Lambda running in each availability zone, to protect against failure in Lambda itself.
1. [ Metrics: not the observability droids you’re looking for ](https://honeycomb.io/blog/2017/10/metrics-not-the-observability-droids-youre-looking-for/)

    High-cardinality fields are where all the interesting data exist, says Charity Majors of Honeycomb. But that’s exactly where most monitoring systems break down, leaving you to throw together hacks to work around their limitations.
1. [Google Cloud Platform Blog: Building good SLOs – CRE life lessons](https://cloudplatform.googleblog.com/2017/10/building-good-SLOs-CRE-life-lessons.html?utm_source=monitoringweekly&utm_medium=email&utm_campaign=Monitoring+Weekly+-+Issue+%2523032&__s=bwykwk1kcceogszq8abt&m=1)

    Google shares some best practices for building Service Level Objectives.
1. [Collaboration > evaluation: Why we pay SRE candidates to interview all-day](https://blog.hostedgraphite.com/2017/10/25/collaboration-evaluation-why-we-pay-sre-candidates-to-interview-all-day/)

    Hosted Graphite brings candidates in to work with them for a day and pays them for their time.
1. [On-Call Horror Story Number Three: This Wins the Most Grueling Award](https://victorops.com/blog/call-horror-story-number-three-wins-grueling-award/)

    Grueling is right: their entire team came to the office over the weekend to work on the outage. Lesson learned:This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Engineering a culture of psychological safety – Inside Intercom](https://blog.intercom.com/psychological-safety/)

    Google’s Project Aristotle discovered that the number one predictor of successful teams is psychological safety. The anecdotes in this piece show how psychological safety is also critical in analyzing incidents.
### Outages

1. [Power outage, coupled with Murphy’s law, leads to raw sewage spill](http://www.themountaineer.com/news/power-outage-coupled-with-murphy-s-law-leads-to-raw/article_f7ff31e2-b8cf-11e7-bff1-1723e32146e9.html)
    The power failed, and then the backup generator failed. Sound familiar? I’m glad datacenters don’t flood with sewage when this happens…
1. [eBay](http://www.ecommercebytes.com/C/blog/blog.pl?/pl/2017/10/1509021696.html)
1. [Texas State Fair](http://www.fox4news.com/news/network-outage-forces-state-fair-to-admit-tens-of-thousands-for-free)
    Their ticket system went down, so they had to admit fair-goers for free.

### [ << Prev ](sreweekly-94.md) ------------- [ Next >> ](sreweekly-96.md)