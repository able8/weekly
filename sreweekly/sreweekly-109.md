## [SRE Weekly Issue #109](https://sreweekly.com/sre-weekly-issue-109/) - February 11, 2018
### Articles

1. [Per-IP rate limiting with iptables](https://making.pusher.com/per-ip-rate-limiting-with-iptables/)

    Pusher had a problem: their service was being bombarded by connections from rogue clients, and they needed to enforce limits. This article is highly polished, with beautiful diagrams and well-constructed explanations.
1. [Structured Logging and Your Team](https://honeycomb.io/blog/)

    Structured logging can bring a lot of uniformity to your infrastructure, as lovingly explained in this article. Snyk explains how that uniformity allows for a standardized troubleshooting methodology that helps them get to the bottom of most problems in minutes.
1. [Your Feature Flag Management Needs to Include Retirement](https://rollout.io/blog/feature-flag-retirement/)

    Feature flags are awesome! But there’s a downside: adding lots of conditional handling to your code can significantly increase code complexity, which can in turn decrease maintainability and increase risk.
1. [Charity Majors on Twitter](https://twitter.com/mipsytipsy/status/957761131216449536)

    Following up on her appearance in the New York Times last week, Charity Majors posted this excellent Twitter thread about the importance of vendor relationship management and generating business value, as any kind of engineer. I’d argue especially as an SRE.
1. [Google Cloud Platform Blog: Applying the Escalation Policy](http://feedproxy.google.com/~r/ClPlBl/~3/X9sgkf7XSWk/applying-the-escalation-policy-CRE-life-lessons.html)

    Here’s the latest in Google’s CRE Life Lessons series. Previously, they explained how to build an Escalation Policy, and in this article, they analyze how it would be applied to several fictitious scenarios.
1. [Dynamometer: Scale Testing HDFS on Minimal Hardware with Maximum Fidelity](https://engineering.linkedin.com/blog/2018/02/dynamometer--scale-testing-hdfs-on-minimal-hardware-with-maximum)

    LinkedIn needed a way to test their HDFS cluster against real-world traffic patterns. The existing solutions didn’t meet their needs (for reasons they explain toward the end), so they created Dynamometer.
1. [Humanize Your Digital Operations](https://www.pagerduty.com/blog/operations-health/)

    PagerDuty released a report this week entitled, “The State of IT Work-Life Balance”, which contains the results of their recent survey. This article is an overview, along with some related tidbits about alert fatigue.
1. [Schrodinger’s Outage](https://www.xaprb.com/blog/schrodingers-outage/)

    Through an anecdote, Baron Schwartz cautions against the use of counter-factuals (“you should have…”) in analyzing the decisions leading up to an outage.
1. [8 Things to Monitor During a Software Deployment](https://stackify.com/monitor-software-deployment/)

    What it says on the tin. This article would make for a great checklist for deploys.
### Outages

1. [Singpass (Singapore ID system)](http://www.todayonline.com/singapore/singpass-corppass-services-suffer-2nd-day-intermittent-outage)
1. [Uber](https://www.cnet.com/news/uber-outage-drivers-and-riders-complain-on-twitter/)
1. [Fortnite](https://www.epicgames.com/fortnite/en-US/news/postmortem-of-service-outage-at-3-4m-ccu)
    Fortnite hit a new peak of 3.4 million concurrent players last Sunday… and that didn’t come without issues!
They suffered 6 different outages over two days, and they posted this highly-detailed incident analysis just 5 days later. Normally I tend not to include outages for MMO games because they have so many and rarely post in-depth analyses, but this one is worth a read.
1. [Binance (cryptocurrency exchange)](http://bitcoinist.com/crypto-exchange-binance-goes-dark-12-hour-outage-maintenance-system-upgrades/)
1. [Google App Engine](http://status.cloud.google.com/incident/appengine/18001#5757907245203456)
1. [US stock brokerages](https://www.cnbc.com/2018/02/06/fidelitys-website-reports-temporary-outage-during-wild-trading-in-us-markets.html)
    The US stock market had a rough week, and so did several brokerage websites as they dealt with the high trading volume.
1. [Super Bowl Advertisers](http://www.netimperative.com/2018/02/avoid-first-lessons-learn-super-bowl-lii/)
    Several companies that purchased expensive commercial slots during the SuperBowl (an american sportsball thing, for you folks outside the US) were unable to handle the web traffic they brought in.
1. [Super Bowl](http://boston.cbslocal.com/2018/02/04/super-bowl-2018-blackout-commercial-nbc/)
    NBC had a 45-second blackout in their broadcast of the Superb Owl.

### [ << Prev ](sreweekly-108.md) ------------- [ Next >> ](sreweekly-110.md)