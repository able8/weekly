## [SRE Weekly Issue #48](https://sreweekly.com/sre-weekly-issue-48/) - November 13, 2016
### Articles

1. [When downtime is not an option](http://saudigazette.com.sa/business/downtime-not-option/)

    A detailed description of Disaster Recovery as a Service (DRaaS), including a discussion of the cost versus creating a DR site oneself. This is the part I always wonder about:
1. [The Prime Directive](http://retrospectives.com/pages/retroPrimeDirective.html)

    This one’s so short I can almost quote the whole thing here. I love its succinctness:
1. [The Netflix Tech Blog: Post-mortem of October 22,2012 AWS degradation](http://techblog.netflix.com/2012/10/post-mortem-of-october-222012-aws.html?m=1)

    Just over four years ago, Amazon had a major outage in Elastic Block Store (EBS). Did you see impact? I sure did. Here’s Netflix’s account of how they survived the outage mostly unscathed.
1. [Serverless promises and the persistent need for critical alerting](http://onpage.com/serverless-meets-security-and-critical-alerting/)

    I’m glad to see more people writing that Serverless != #NoOps. This article is well-argued even though it turns into an OnPage ad 3 paragraphs from the end.
1. [Episode 004: Charity Majors – Greater Than Code](https://www.greaterthancode.com/2016/10/21/episode-004-charity-majors/)

    What else can we expect from Greater Than Code + Charity Majors? This podcast is 50 minutes of awesome, and there’s a transcription, too! Listen/read for awesome phrases like “stamping out chaos”, find out why Charity says, “I personally hate [the term ‘SRE’] (but I hate a lot of things)”, and hear Conway’s law applied to microservices, #NoOps debunking, and a poignant ending about misogyny and equality.
1. [Microsoft Announces Azure DNS General Availability](https://www.infoq.com/news/2016/10/azure-dns-ga)

    Microsoft released its Route 53 competitor in late September. They say:
1. [Communication Breakdown Leads to Patient Burn](https://bwhsafetymatters.org/2016/08/12/communication-breakdown-leads-to-patient-burn/)

    This issue of BWH Safety Matters details an incident in which a communication issue between teams that don’t normally work together resulted in a patient injury. This is exactly the kind of pitfall that becomes more prevalent with the move toward microservices, as siloed teams sometimes come into contact only during an incident.
1. [New systems will fail: A site outage case study from Envato Market](https://envato.com/blog/new-systems-will-fail-unexpected-ways-case-study-envato-market/)

    A detailed postmortem from an outage last month. Lots of takeaways, including one that kept coming up: test your emergency tooling before you need to use it.
### Outages

1. [Canada’s immigration site](http://www.channel3000.com/news/politics/canada-immigration-site-goes-down/42429600)
    I’m sure this is indicative of something.
1. [Office 365](http://www.channelbiz.co.uk/2016/11/03/office-365-users-suffer-yet-another-network-outage/)
1. [Twitter](http://www.csoonline.com/article/3138934/security/bgp-errors-are-to-blame-for-monday-s-twitter-outage-not-ddos-attacks.html)
    Twitter stopped announcing their AS in BGP worldwide, resulting in a 30-minute outage on Monday.
1. [Google BigQuery](http://status.cloud.google.com/incident/bigquery/18022#5165166324875264)
    Google writes really great postmortems! Here’s one for a 4-hour outage in BigQuery on November 8, posted on November 11. Fast turnaround and an excellent analysis. Thanks, Google — we appreciate your hard work and transparency!
1. [Pingdom](http://status.pingdom.com/incidents/73tpmc4tfzk5)
    Normally I wouldn’t include such a minor outage, but I love the phrase “unintended human error” that they used. Much better than the intended kind.
1. [WikiLeaks](http://www.ibtimes.co.uk/was-twitter-outage-really-related-ddos-attack-wikileaks-microblogging-site-says-no-1590370)
1. [eBay](http://www.tweaktown.com/news/54840/ebay-experiencing-downtime-hours-twitter-outage/)

### [ << Prev ](sreweekly-47.md) ------------- [ Next >> ](sreweekly-49.md)