## [SRE Weekly Issue #89](https://sreweekly.com/sre-weekly-issue-89/) - September 17, 2017
### Articles

1. [Cachet: The Open Source Status Page System](https://cachethq.io/)

    Cachet looks like a pretty good contender to incumbents like StatusPage.
1. [Adventures in fault tolerant alerting with Python](https://blog.hostedgraphite.com/2017/05/05/adventures-in-fault-tolerant-alerting-with-python/)

    Hosted Graphite used PySyncObj to create a fault-tolerant threshold alerting feature.
1. [On-Call Horror Story Number One: This Wins the Most Embarrassing Award](https://victorops.com/blog/call-horror-story-number-one-wins-embarrassing-award/)

    Talk about a high-pressure incident! When a teleconferencing provider’s wires got crossed, hilarity (and embarassment) ensued.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [5 Things I Didn’t Expect to See While Shadowing On-Call](https://www.pagerduty.com/blog/5-things-shadow-on-call/)

    This story is from a PagerDuty engineer. What’d you learn while shadowing on-call? I’d love to hear your story!
1. [Build a Modern Operations Process with StatusPage](https://medium.com/synqfm-ops/build-a-modern-operations-process-with-statuspage-d7244b8da18?source=rss----c5d5bfbdd67b---4)

    Here’s how SYNQ set their status page up. They’re the folks that committed to publishing all of their incident followups publicly a month or two back. Transparency FTW!
1. [Handling 1 Million Requests per Minute with Golang](https://medium.com/smsjunk/handling-1-million-requests-per-minute-with-golang-f70ac505fcaa)

    I’ll save you the math: that’s ~17k req/sec. I really like that this article takes us through their learning process and their first failed attempts.
1. [Our First Engineering Game Day – Quid Technology](http://technology.quid.com/2017/09/our-first-engineering-game-day/)

    Quid wrote up this explanation of how they set up their game day and what they learned. I really like the structure they used, and I may draw heavily on it for my own game days.
1. [Monitoring Isn’t Observability](https://www.vividcortex.com/blog/monitoring-isnt-observability)

    “Observability” as a term is making the rounds like “DevOps” did (and still does…). Here’s Baron Schwartz’s take on it.
### Outages

1. [Google Services](https://www.moroccoworldnews.com/2017/09/228614/google-services-experience-unexplained-outages-throughout-early-september/)
    As two astute readers pointed out (thanks!), the Gmail outage I included in the last issue was from 2009(!). Oops. However, Google has been experiencing a series of outages and degradations this month, so I’m just going to pretend I knew that rather than that I forgot to check the date on the article.
1. [s3 outageHeroku’s](https://s)
    S3 had an outage in us-east-1 on September 14th. This one showed up as yellow on their status site, with the text below. Companies that depend on S3 probably saw impact as well, but I couldn’t find any status posts other than Heroku’s.
11:58 AM PDT We are investigating increased error rates for Amazon S3 requests in the US-EAST-1 Region.
12:20 PM PDT We can confirm that some customers are receiving throttling errors accessing S3. We are currently investigating the root cause.
12:38 PM PDT We continue to work towards resolving the increased throttling errors for Amazon S3 requests in the US-EAST-1 Region. We have identified the subsystem responsible for the errors, identified root cause and are now working to resolve the issue.
12:49 PM PDT We are now seeing recovery in the throttle error rates accessing Amazon S3. We have identified the root cause and have taken actions to prevent recurrence.
1:05 PM PDT Between 11:40 AM and 12:56 PM PDT we experienced throttling errors accessing Amazon S3 in the US-EAST-1 Region. The issue is resolved and the service is operating normally.
Full disclosure: Heroku is my employer.
1. [IBM](https://www.theregister.co.uk/2017/09/08/ibms_global_load_balancer_and_reverse_dns_degraded_by_domain_transfer_mess/)
    IBM had a mishap when transferring control of some of its domains to a different registrar. Some of their services including their Global Load Balancer went down.

### [ << Prev ](sreweekly-88.md) ------------- [ Next >> ](sreweekly-90.md)