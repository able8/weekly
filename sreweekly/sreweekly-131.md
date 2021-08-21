## [SRE Weekly Issue #131](https://sreweekly.com/sre-weekly-issue-131/) - July 22, 2018
### Articles

1. [Twitter: @alicegoldfuss on New Ops](https://mobile.twitter.com/alicegoldfuss/status/1020002647196102656)

    I love the idea of using hobbies as a gauge for your overload level at work. Also, serious kudos to Alice for the firm stance against alcohol at work and especially in Ops.Alice Goldfuss
1. [Open sourcing oomd, a new approach to handling OOMs](https://code.fb.com/production-engineering/open-sourcing-oomd-a-new-approach-to-handling-ooms/)

    If the Linux OOM killer gets involved, you’ve already lost. Facebook reckons they can do better.Daniel Xu — Facebook

1. [Debug a Real Honeycomb Outage with Honeycomb](https://play.honeycomb.io/)

    This is radical transparency: Honeycomb has set up a sandbox copy of their app for you to play with and loaded it with data from a real outage on their platform! Tinker away. It’s super fun.Honeycomb
1. [Good housekeeping for error budgets – part two – CRE life lessons](http://feedproxy.google.com/~r/ClPlBl/~3/kLed_L-HTWU/cre-life-lessons-good-housekeeping-for-error-budgets.html)

    It may not actually make sense to halt feature development if your team has exhausted the error budget. What do you do instead?Adrian Hilton, Alec Warner and Alex Bramley — Google
1. [Introducing Centrifuge](https://segment.com/blog/introducing-centrifuge/)

    The parallels to the Plaid article a few weeks ago (scaling 9000+ heterogeneous bank integrations) are intriguing.Calvin French-Owen — Segment
1. [SLOs & You: A Guide To Service Level Objectives](https://www.circonus.com/2018/07/a-guide-to-service-level-objectives/)

    A solid definition of SLIs, SLOs, and SLAs (from someone other than Google!). Includes some interesting tidbits on defining and measuring availability, choosing a useful time quantum, etc.Kevin Kamel — Circonus
1. [Rolling the Heroku Redis Fleet](https://blog.heroku.com/rolling-redis-fleet)

    Read about how Heroku deployed a security fix to their fleet of customer Redis instances. This is awesome:Camille Baldock — Heroku
1. [Exploring Multi-level Weaknesses using Automated Chaos Experiments](https://medium.com/chaosiq/exploring-multi-level-weaknesses-using-automated-chaos-experiments-aa30f0605ce?source=rss----812e5d368303---4)

    Russ Miles — ChaosIQ
1. [Load Testing Round Up: 8 tools you can use to strengthen your stack](https://buttercms.com/blog/load-testing-round-up-8-tools-you-can-use-to-strengthen-your-stack)

    A comparison of 2 free and 6 paid tools for load testing, along with advice on how to use them.Noah Heinrich — ButterCMS
1. [Why Having a Feature Flag Microservice Is a Bad Idea](https://rollout.io/blog/feature-flag-microservice/)

    One could even call this article, “Why having a single microservice that every other microservice depends on is a bad idea”.Mark Henke — Rollout.io
### Outages

1. [Google Cloud Platform](https://status.cloud.google.com/incident/cloud-networking/18012#18012005)
    Perhaps you noticed that a ton of sites fell over this past Tuesday? Or maybe you were on the front lines dealing with it yourself. Google’s Global Load Balancer fleet suffered a major outage, and they posted this detailed analysis/apology the next day.
1. [Amazon’s Prime Day](https://money.cnn.com/2018/07/16/technology/amazon-website-down-prime-day/index.html?utm_source=CNN%2BFive%2BThings&utm_campaign=0409980018-EMAIL_CAMPAIGN_2018_07_16_11_47&utm_medium=email&utm_term=0_6da287d761-0409980018-83152125)
    Seems like a tradition at this point…
1. [Azure](https://azure.microsoft.com/en-us/status/history/)
    A BGP announcement error caused global instability for VM instances trying to reach Azure endpoints.
1. [PagerDuty](https://status.pagerduty.com/incidents/dmvgsgs7vjc3)
1. [Slack](https://status.slack.com/2018-07/c3d474b41bb85298)
1. [Atlassian Statuspage](https://metastatuspage.com/incidents/v0ctbw20x441)
1. [British Airways](http://www.itpro.co.uk/business-operations/31540/british-airways-it-outage-leaves-travellers-stranded)
1. [Twitter](https://www.newsweek.com/twitter-outage-not-working-down-today-1033391)
1. [Fortnite: Playground LTM Postmortem](https://www.epicgames.com/fortnite/en-US/news/playground-ltm-postmortem)
    This is a really juicy incident analysis! Epic Games tried to release a new game mode for Fortnite and quickly discovered a major scaling issue in their system, which they explain in great detail.
The process of getting Playground stable and in the hands of our players was tougher than we would have liked, but was a solid reminder that complex distributed systems fail in unpredictable ways. We were forced to make significant emergency upgrades to our Matchmaking Service, but these changes will serve the game well as we continue to grow and expand our player base into the future.
The Fortnite Team — Epic Games
1. [Snapchat](http://uk.businessinsider.com/snapchat-down-for-some-users-second-time-two-weeks-2018-7)
1. [Facebook](https://news.finance.co.uk/facebook-amid-major-international-outage/)
1. [reddit](https://reddit.statuspage.io/incidents/jyg5kv7p29fd)

### [ << Prev ](sreweekly-130.md) ------------- [ Next >> ](sreweekly-132.md)