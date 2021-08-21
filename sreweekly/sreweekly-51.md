## [SRE Weekly Issue #51](https://sreweekly.com/sre-weekly-issue-51/) - December 18, 2016
### Articles

1. [Etsy’s Debriefing Facilitation Guide for Blameless Postmortems](https://codeascraft.com/2016/11/17/debriefing-facilitation-guide/)

    This is a big moment for the SRE field. Etsy has distilled the internal training materials they use to teach employees how to facilitate retrospectives (“debriefings” in Etsy parlance). They’ve released a guide and posted this introduction that really stands firmly on its own. I love the real-world story they share.
1. [Debriefing Facilitation Guide [PDF]](https://extfiles.etsy.com/DebriefingFacilitationGuide.pdf)

    And here’s the guide itself. This is essential reading for any SRE interested in understanding incidents in their organization.
1. [Slicer: Auto-sharding for datacenter applications | the morning paper](https://blog.acolyer.org/2016/12/02/slicer-auto-sharding-for-datacenter-applications/)

    Click through to find out. It’ll be interesting to see what open source projects this paper inspires.
1. [The Problem with Pre-aggregated Metrics: Part 2, the “aggregated”](https://honeycomb.io/blog/2016/12/the-problem-with-pre-aggregated-metrics-part-2-the-aggregated/)

    The second in a series, this article delves into the pitfalls of aggregating metrics. Aggregation means you have to choose between bloating your time-series datastore or leaving out crucial stats that you may need during an investigation.
1. [sysadvent: Day 15 – Take That Vacation: Eliminate Alerts Dragging You Back to the Office](http://feedproxy.google.com/~r/sysadvent/~3/PzUZ-02cjCM/day-15-take-that-vacation-eliminate.html)

    I thought this was going to be primarily an argument for reducing burnout to improve reliability. That’s in there, but the bulk of this article is a bunch of tips and techniques for improving your monitoring and alerting to reduce the likelihood that you’ll be pulled away from your vacation.
1. [sysadvent: Day 17 – Write it down or suffer the consequences](http://feedproxy.google.com/~r/sysadvent/~3/isXGHmIvSE4/day-17-write-it-down-or-suffer.html)

    The title says it all. Losing the only person with the knowledge of how to keep your infrastructure running is a huge reliability risk. In this article, Heidi Waterhouse (who I coincidentally just met at LISA16!) makes it brilliantly clear why you need good documentation and how to get there.
1. [Why Now Is the Time to Implement Redundant DNS](https://dzone.com/articles/why-now-is-the-time-to-implement-redundant-dns)

    Here’s another overview of implementing a secondary DNS provider. I like that they cover the difficulties that can arise when you use a provider’s proprietary non-RFC DNS extensions such as weighted round-robin record sets.
### Outages

1. [EC2 (us-west-1), Herokuaffected](http://status.aws.amazon.com)
    EC2’s Dublin region had an outage in the DNS resolver provided to instances via DHCP. Heroku was affected as well.Full disclosure: Heroku is my employer.
1. [DirecTV Now](http://awfulannouncing.com/2016/directv-now-experiences-three-hour-outage-has-further-rsn-issues.html)
1. [ChangeIP (DNS provider)](http://www.networkworld.com/article/3149576/internet/dns-provider-changeip-cites-mysql-database-crash-for-days-long-outage.html)
    ChangeIP tweeted that they suffered a major MySQL failure.
1. [ATO (Australian Tax Office)](http://theconversation.com/server-down-what-caused-the-ato-systems-to-crash-70396)
    The ATO lost a petabyte of data from their HPE 3PAR StoreServe SAN.
1. [PSN](http://www.dailystar.co.uk/tech/gaming/569697/PSN-Status-DOWN-PlayStation-Network-not-working-PS4-PS3-PS-Vita-Hacked)
1. [Battlefield 1](https://www.hackread.com/you-are-not-alone-ea-servers-are-down/)

### [ << Prev ](sreweekly-50.md) ------------- [ Next >> ](sreweekly-52.md)