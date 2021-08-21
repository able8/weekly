## [SRE Weekly Issue #85](https://sreweekly.com/sre-weekly-issue-85/) - August 13, 2017
### Articles

1. [Ops: It’s everyone’s job now](https://opensource.com/article/17/7/state-systems-administration)

    Here’s Charity Majors with another gem about how ops looks in the era of distributed systems.
1. [So, about this Googler’s manifesto.](https://medium.com/@yonatanzunger/so-about-this-googlers-manifesto-1e3773ed1788)

    I hope most of you have been reading up on the infamous “Googler manifesto”, and if so, maybe you’ve already seen this article. What caught my eye is the emphasis on people-oriented engineering, because these are the skills that have become increasingly important to me as an SRE.
1. [Beyond Observing Behavior](https://blog.turbinelabs.io/beyond-observing-behavior-a0887fe0d3ca)

    A key metric goes through the roof and pages you. Why? Answering that can be really easy if you can quickly see the changes deployed to your system around the same time. This article is about a specific product that solves this problem and is thus a bit advertisey, but it’s still a good read.
1. [Glitches are inevitable, and consumers know it: how anomaly detection can safely corral a glitch stampede](http://www.techshout.com/internet/2017/10/glitches-inevitable-consumers-know-anomaly-detection-can-safely-corral-glitch-stampede/)

    Here’s a good argument for anomaly detection. Great, but I still have yet to see anomaly detection that I trust! That said, this was still an interesting read due to the real-world story about a glitch Wal-Mart faced.
1. [Achieving Fault Tolerance With Resilience4j](https://dzone.com/articles/resilience4j-intro)

    For the Java crowd, here’s a primer on Resilience4j, a framework that makes it easier to write code that can recover from errors.
1. [Beyond Google SRE: What is Site Reliability Engineering like at Medium?](https://blog.netsil.com/beyond-google-sre-what-is-site-reliability-engineering-like-at-medium-71c65bd35f4e)

    I like the description of their “The Watch” pager rotation in which developers periodically serve.
1. [Migrating Existing Datastores](http://engineering.grab.com/migrating-existing-datastores)

    Grab engineers talk about migrating from Redis to ElastiCache veeeery carefully.
### Outages

1. [Paragon (game)](https://www.epicgames.com/paragon/forums/showthread.php?79443-ongoing-outage-technical-update)
    Epic Games released version 42 of Paragon, and the new version unexpectedly overloaded their servers. To get back to a good state, they were forced into developing novel code and upgrading a DB on the fly.
1. [FedEx](https://www.nytimes.com/aponline/2017/08/07/business/ap-us-fedex-outage.html)
1. [SYNQ](https://medium.com/synqfm-ops/rca-for-live-stream-playback-issue-on-june-20th-2017-767bc1d01fcc?source=rss----c5d5bfbdd67b---4)
    As mentioned here previously, SYNQ has dedicated to posting their incident RCAs publicly. In this one, they identified a need for better regression testing.

### [ << Prev ](sreweekly-84.md) ------------- [ Next >> ](sreweekly-86.md)