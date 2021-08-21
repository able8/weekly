## [SRE Weekly Issue #36](https://sreweekly.com/sre-weekly-issue-36/) - August 21, 2016
### Articles

1. [Multi data center redundancy – sysadmin considerations](https://blog.serverdensity.com/multi-data-center-redundancy-sysadmin-considerations/)

    This is the second half of Server Density’s series on the lessons they learned as they transitioned to a multi-datacenter architecture.  There are lots of interesting tidbits in here, such as an explanation of how they handle failover to the secondary DC and what they do if that goes wrong.Full disclosure: Heroku, my employer, is mentioned.
1. [How Complex Web Systems Fail ](https://medium.com/production-ready/how-complex-web-systems-fail-part-2-1b1afb4fa7be#.d0un51ndx)

    Here’s the second half of Mathias Lafeldt’s series that seeks to apply  Richard Cook’s How Complex Systems Fail to web systems.  The article is great, but the really awesome part is the thoughtful responses by Cook himself to both parts one and two, linked at the end of this article.
1. [Why Reddit was down on Aug 11](https://www.reddit.com/r/announcements/comments/4y0m56/why_reddit_was_down_on_aug_11/)

    Here’s a postmortem for last week’s outage that involved a migration gone awry. Thanks to Jonathan Rudenberg for this one.
1. [US Patent Office sued after it declared a power outage a ‘national holiday’](http://www.theregister.co.uk/2016/08/18/uspto_sued_for_early_xmas/)

    A patent holding firm is alleging that the USPTO overstepped its authority in declaring a system outage (reported in issue #4) to be treated as a national holiday for purposes of deadlines, and that this led to the plaintiff being sued.  
1. [Know Anyone With This High-Burnout Job?](https://victorops.com/blog/burnout-monitorama/)

    Burnout is a crucially important consideration in a field with on-call work.  VictorOps has a few tips for alleviating burnout gleaned from this year’s Monitorama.
1. [Staging Servers Are Dead!](https://dzone.com/articles/staging-servers-are-dead)

    Edith Harbaugh says that staging servers present a reliability risk that doesn’t outweigh their benefit.  This article is an update to her original article, which I also recommend reading.
1. [Context aware MySQL pools via HAProxy](http://githubengineering.com/context-aware-mysql-pools-via-haproxy/)

    Github uses HAProxy to balance across is read-only MySQL replicas, which is a common method.  Their technique for excluding lagging nodes while avoiding entirely emptying the pool if all nodes are lagging is pretty neat.Thanks to Devops Weekly for this one.
1. [Serverless Architectures](http://martinfowler.com/articles/serverless.html)

    A highly detailed deep-dive on Serverless — what it means, benefits, and drawbacks.  I especially enjoyed the #NoOps section:
 #ServerlessIsMadeOfServersFull disclosure: Heroku, my employer, is mentioned.
### Outages

1. [Slack](https://status.slack.com/2016-08/42ad7156695998fc)
    A relatively minor issue, but it impacted me, so I logged it here while awaiting resolution.
1. [MTN (mobile telecom)](https://www.techcentral.co.za/gremlins-again-hit-mtn-network/67764/)
1. [Google Cloud Status Dashboard](http://status.cloud.google.com/incident/cloud-dataflow/16001#5668600916475904)
    Postmortem included, with an interesting cause:
During mitigation of a lower impact performance issue, Google engineers made a configuration change to pipeline orchestration.  An error in this configuration caused validation within the orchestration component to reject all requests.
1. [Tesla Vehicles](http://www.datacenterdynamics.com/content-tracks/core-edge/mobile-network-outage-shuts-down-tesla-connectivity/96792.fullarticle)
1. [Xbox Live](http://www.dailystar.co.uk/tech/gaming/538685/Xbox-Live-down-Xbox-One-not-working-Xbox-status-Microsoft-Core-services-signing-in)
1. [Sky (ISP)](http://www.cbronline.com/news/telecoms/network/sky-broadband-customers-hit-by-outage-for-second-time-this-month-4981418)
1. [Facebook](https://www.rt.com/news/356184-facebook-down-users-say/)
1. [Apple’s App Store](http://www.telegraph.co.uk/technology/2016/08/16/app-store-goes-down-leaving-iphone-users-unable-to-download-new/)
1. [Twitter](http://indianexpress.com/article/technology/social/twitter-was-down-worldwide-for-a-few-hours-2978068/)
1. [Cisco Jasper](http://www.m2mcafe.com/news/cisco-jasper-iot-network-faces-outage-in-us-3021.html)
1. [Optus](http://www.smh.com.au/sport/soccer/optus-out-epl-fans-rage-after-football-streaming-service-hits-technical-faults-20160814-gqsen8.html)
1. [AT&T](http://www.bizjournals.com/dayton/news/2016/08/13/att-hit-by-massive-wireless-network-outage.html)
1. [NSA](http://www.scmagazine.com/nsa-blames-storm-for-website-outage/article/516909/)

### [ << Prev ](sreweekly-35.md) ------------- [ Next >> ](sreweekly-37.md)