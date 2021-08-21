## [SRE Weekly Issue #100](https://sreweekly.com/sre-weekly-issue-100/) - December 3, 2017
### Articles

1. [ void *: Incidents as Untyped Pointers](https://www.snafucatchers.com/single-post/2017/11/14/void-Incidents-as-Untyped-Pointers)

    Richard Cook wrote this short, incredibly insightful essay on how we can use incidents to improve our mental model of the system.
1. [a short example of why dimensions are suuuuuper valuable](https://medium.com/@mhat/a-short-example-of-why-dimensions-are-suuuuuper-valuable-67e880055eb0)

    A nifty trip through a debugging session that shows the importance of being able to dig into high-cardinality fields in your monitoring system.
1. [How to Monitor the SRE Golden Signals](https://medium.com/devopslinks/how-to-monitor-the-sre-golden-signals-1391cadc7524?source=rss-e781fc30beac------2)

    Various sources list a couple of key metrics to keep an eye on, including request rate, error rate, latency, and others. This 6-part series defines the golden signals and shows how to monitor them in several popular systems.
1. [Thrift on Steroids: A Tale of Scale and Abstraction](http://bravenewgeek.com/thrift-on-steroids-a-tale-of-scale-and-abstraction/)

    This article explains some downsides of Thrift and introduces the author’s solution: Frugal, a Thrift wrapper.
1. [re:Invent 2017 | New Products & Services](https://aws.amazon.com/new/reinvent/)

    re:Invent 2017 is over (whew) and now we have a raft of new products and features to play with. I’m going to leave the detailed analysis for Last Week in AWS and just point out a few bits of special interest to SREs:
1. [ How Etsy caches: hashing, Ketama, and cache smearing](https://codeascraft.com/2017/11/30/how-etsy-caches/)

    Etsy details their caching setup and explains the importance of consistent hashing in cache cluster design. I haven’t heard of their practice of “cache smearing” before, and I like it.
1. [The role of software in spacecraft accidents](https://blog.acolyer.org/2017/11/30/the-role-of-software-in-spacecraft-accidents/)

    
1. [Won’t Get Fooled Again](https://blog.gremlin.com/wont-get-fooled-again-95a3f2bb0a2e?source=rss----94cbb73a7983---4)

    Gremlin had an incident that was caused by filled disks. Because they’re Gremlin, they now purposefully fill a disk on a random server every day just to make sure their systems deal with it gracefully, a practice they call “continuous chaos”.
1. [Fearless shared postmortems — CRE life lessons](https://cloudplatform.googleblog.com/2017/11/fearless-shared-postmortems-CRE-life-lessons.html?m=1)

    Google’s CRE team (Customer Reliability Engineering) discusses when to post public followups and how to write them. I love their idea of investigating where they got lucky during an incident, catching cases where things could have been much worse if not for serendipity. I’m going to start using that.
### Outages

1. [WhatsApp](https://www.techworm.net/2017/11/not-just-whatsapp-everyone.html)
1. [Spotify](https://www.christiantoday.com/article/spotify-update-music-streaming-up-and-running-after-downtime-last-tuesday/120083.htm)
1. [J Crew](http://www.ibtimes.com/j-crew-website-down-cyber-monday-check-out-not-working-2620128)
1. [Southwest Airlines](http://www.ibtimes.com/southwest-airlines-website-down-booking-problems-monday-2620070)
1. [Yahoo Groups](https://www.bleepingcomputer.com/news/technology/yahoo-groups-plagued-by-downtime-technical-issues-for-almost-a-week/)

### [ << Prev ](sreweekly-99.md) ------------- [ Next >> ](sreweekly-101.md)