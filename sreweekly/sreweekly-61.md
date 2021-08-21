## [SRE Weekly Issue #61](https://sreweekly.com/sre-weekly-issue-61/) - February 26, 2017
### Articles

1. [7 Steps to Avoiding Downtime](https://www.pagerduty.com/blog/7-steps-avoiding-downtime/)

    Every week, there’s an article with a title like this (just like with “costs of downtime”). Almost every week, they’re total crap, but this one from PagerDuty is a bit better than the rest. The bit that interests me is the assertion that a microservice-based architecture “makes maintenance much easier” and “makes your app more resilient”. Sure it can, but it can also just mean that you trade one problem for 1300 problems.
1. [Optimizing Your Alert Management Process](https://www.pagerduty.com/blog/optimizing-alert-management/)

    
1. [Scaling @ HelloFresh: API Gateway](http://feedproxy.google.com/~r/HighScalability/~3/z4Syb9_rNzM/scaling-hellofresh-api-gateway.html)

    HelloFresh live-migrated their infrastructure to an API gateway to facilitate a transition to microservices. They kindly wrote up their experience, which is especially educational because their first roll-out attempt didn’t go as planned.
1. [The Pros and Cons of Eating Your Own Dog Food — Production Ready](https://medium.com/production-ready/the-pros-and-cons-of-eating-your-own-dog-food-f76f0ebadbcc)

    In this issue, Mathias shows us the benefits of “dogfooding” and cases where it can break down. I like the way the feedback loop is shortened, so that developers feel a painful user experience and have incentive to quickly fix it. It reminds me a lot of the feedback loop you get when developers go on call for the services they write.
1. [Making Sense of the Application Monitoring Landscape](https://blog.netsil.com/making-sense-of-the-application-monitoring-landscape-bd9ae2e60233#.2cnthryd9)

    A breakdown of four categories of monitoring tools using the “2×2” framework. I like the mapping of “personas” (engineering roles) to the monitoring typesa they tend to find most useful.
### Outages

1. [Cloudflare: “Cloudbleed”PagerDutyMonzoTechDirtMaxMindlist of sites using cloudflare](https://blog.cloudflare.com/incident-report-on-memory-leak-caused-by-cloudflare-parser-bug/)
    Cloudflare experienced a minor outage due to mitigating a major leak of private information. They posted this (incredibly!) detailed analysis of the bug and their response to it. Other vendors such as PagerDuty, Monzo, TechDirt, and MaxMind posted responses to the outage. There’s also this handy list of sites using cloudflare.
1. [mailgunlinked to](http://status.mailgun.com/incidents/p9nxxql8g9rh)
    Here’s a really interesting postmortem for a Mailgun outage I linked to in January. What apparently started off as a relatively minor outage was significantly exacerbated “due to human error”. The intriguing bit: the “corrective actions” section makes no mention at all of process improvements to make the system more resilient to this kind of error.
1. [All Circuits are Busy Now: The 1990 AT&T Long Distance Network Collapse](http://users.csc.calpoly.edu/~jdalbey/SWE/Papers/att_collapse.html)
    In 1990, the entire AT&T phone network experienced a catastrophic failure, and 50% of all calls failed. The analysis is pretty interesting and shows us that a simple bug can break even an incredibly resilient distributed system.
the Jan. 1990 incident showed the possibility for all of the modules to go “crazy” at once, how bugs in self-healing software can bring down healthy systems, and the difficulty of detecting obscure load- and time-dependent defects in software.
1. [vzaar](https://vzaar.com/blog/2017/02/22/mondays-outage-what-happened-how-well-prevent-it-from-happening-again/)
    They usually fork a release branch off of master, test it, and push that out to production. This time, they accidentally pushed master to production. How do I know that? Because they published this excellent post-analysis of the incident just two days after it happened.
1. [U.S. Dept. of Homeland Security](http://www.themarshalltown.com/u-s-homeland-security-employees-get-locked-out-of-computer-networks-officials-blame-presidents-day-report/23057)
    This article has some vague mention of an expired certificate.
1. [YouTube](https://9to5google.com/2017/02/21/its-not-just-you-youtube-is-down-for-many/)
1. [CD Baby](https://www.facebook.com/cdbaby/posts/10154989848889330)
1. [Facebook](http://www.dailymail.co.uk/news/article-4248540/Users-Britain-report-problems-Facebook.html)

### [ << Prev ](sreweekly-60.md) ------------- [ Next >> ](sreweekly-62.md)