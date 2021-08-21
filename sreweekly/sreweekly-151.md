## [SRE Weekly Issue #151](https://sreweekly.com/sre-weekly-issue-151/) - December 9, 2018
### Articles

1. [A victim of its own popularity: Scaling our CloudWatch integration](https://blog.hostedgraphite.com/2018/12/03/a-victim-of-its-own-popularity-scaling-our-cloudwatch-integration/)

    They used feature flags to safely transition from a single-host service to a horizontally-scaled distributed system.Ciaran Egan and Cian Synnott — Hosted Graphite
1. [Working with AWS Limits](https://www.awsadvent.com/2018/12/07/working-with-aws-limits/)

    Limits and quotas can really ruin your day, and it can be very difficult to predict limit exhaustion before a change reaches production, as we learn in this incident story from RealSelf.Bakha Nurzhanov — RealSelf
1. [Defending Against Abuse at LinkedIn’s Scale](https://engineering.linkedin.com/blog/2018/12/defending-against-abuse-at-linkedins-scale)

    The challenge: you have to defend against abuse to keep your service running, but the abuse detection also must not adversely impact the user experience.Sahil Handa — LinkedIn
1. [Answer to the Ultimate Question of (On-Call) Life, the Universe, and Everything: 71](https://www.pagerduty.com/blog/answer-ultimate-question-on-call-life/)

    PagerDuty has developed a system for measuring on-call health, factoring in quantity of pages, time of each page, frequency, clustering of pages, etc. I love what they’re doing and I hope we see more of this in our industry.Lisa Yang — PagerDuty
1. [Spooky Tales of Testing In Production: A Recap and Lessons Learned](https://www.honeycomb.io/blog/spooky-tales-of-testing-in-production-a-recap-and-lessons-learned/)

    A summary of three outage stories from Honeycomb’s recent event. My favorite is the third:Alaina Valenzuela — Honeycomb
1. [Reasons to Scale Horizontally](https://blog.wallaroolabs.com/2018/11/horizontal-scaling-reasons/)

    Looking at that title, I thought to myself, “Uh, because it’s better?” It’s worth a read though, because it so eloquently explains horizontal versus vertical scaling, why you’d do one or the other, and why horizontal scaling is hard.Sean T. Allen — Wallaroo Labs
1. [Cache warming: Agility for a stateful service](https://medium.com/netflix-techblog/cache-warming-agility-for-a-stateful-service-2d3b1da82642?source=rss----2615bd06b42e---4)

    Netflix has some truly massive cache systems at a scale of hundreds of terabytes. Find out what they do to warm up new cache nodes before putting them in production.Deva Jayaraman, Shashi Madappa, Sridhar Enugula, and Ioannis Papapanagiotou — Netflix
1. [Software Sprawl, The Golden Path, and Scaling Teams With Agency](https://charity.wtf/2018/12/02/software-sprawl-the-golden-path-and-scaling-teams-with-agency/)

    This article lays out a promising plan for reducing the number of technologies your engineering department is using while still giving engineers the freedom to choose the right tool for the job.Charity Majors
### Outages

1. [Nest](https://www.engadget.com/2018/12/03/nest-outage/)
1. [GitHub](https://status.github.com/messages)
1. [O2 (UK) and SoftBank (Japan)](https://www.theregister.co.uk/2018/12/06/ericsson_o2_telefonica_uk_outage/)
    I normally don’t bother mentioning mobile phone service outages, but this one has an interesting cause: an expired TLS certificate in Ericsson’s systems.
1. [Google Allo and Duo](https://9to5google.com/2018/12/05/google-allo-down-outage/)
1. [Facebook](https://www.plymouthherald.co.uk/news/facebook-goes-down-again-users-2295108)

### [ << Prev ](sreweekly-150.md) ------------- [ Next >> ](sreweekly-152.md)