## [SRE Weekly Issue #110](https://sreweekly.com/sre-weekly-issue-110/) - February 18, 2018
### Articles

1. [How production engineers support global events on Facebook](https://code.facebook.com/posts/166966743929963/how-production-engineers-support-global-events-on-facebook/?utm_source=codedot_rss_feed&utm_medium=rss&utm_campaign=RSS+Feed)

    Facebook goes in-depth on their preparations for New Year’s Day 2018 in their live streaming infrastructure. They used forecasting based on last year and various kinds of load testing to develop the right kind of scaling strategy to meet demand.
1. [On-call doesn’t have to suck](http://t.dripemail2.com/c/eyJhY2NvdW50X2lkIjoiMjY1Njc0MyIsImRlbGl2ZXJ5X2lkIjoiMjExODAxMzgzMSIsInVybCI6Imh0dHBzOi8vbWVkaXVtLmNvbS9AY29weWNvbnN0cnVjdC9vbi1jYWxsLWIwYmQ4YzVlYTRlMD91dG1fc291cmNlPW1vbml0b3Jpbmd3ZWVrbHlcdTAwMjZ1dG1fY2FtcGFpZ249aXNzdWVfMDQ2XHUwMDI2X19zPWJ3eWt3azFrY2Nlb2dzenE4YWJ0XHUwMDI2dXRtX3NvdXJjZT1tb25pdG9yaW5nd2Vla2x5XHUwMDI2dXRtX21lZGl1bT1lbWFpbFx1MDAyNnV0bV9jYW1wYWlnbj1Nb25pdG9yaW5nK1dlZWtseSstK0lzc3VlKyUyMzA0NiJ9)

    Cindy Sridharan went and blew up the internet with an excellent and controversial tweet about on-call. She took to Medium to address all of the discussion that followed, and the result is a pretty excellent article about on-call and work/life balance.
1. [ Production Test Run: Overburdened and under provisioned ](https://ayende.com/blog/181537-B/production-test-run-overburdened-and-under-provisioned)

    A discussion about how RavenDB handles resource exhaustion, and just how resource exhaustion can be defined and detected.
1. [Development at Honeycomb: Crossing the Observability Bridge to Production](https://honeycomb.io/blog/2018/02/development-at-honeycomb-crossing-the-observability-bridge-to-production/)

    Honeycomb on using observability tooling to precisely analyze how a change actually affects your users. Did the new feature/bugfix have the effect you expected?
1. [Low latency, large working set, and GHC’s garbage collector: pick two of three](https://making.pusher.com/latency-working-set-ghc-gc-pick-two/)

    Pusher is obsessed with low latency, and for good reason. When they saw high long-tail latency, they discovered that Haskell’s garbage collector is optimized for throughput, rather than latency.
1. [Resilience Engineering at LinkedIn with Project Waterbear | LinkedIn Engineering](https://engineering.linkedin.com/blog/2017/11/resilience-engineering-at-linkedin-with-project-waterbear)

    Facebook’s Project Waterbear seeks to improve resiliency across many of their services through a combination of chaos engineering, cultural changes, and improvements to Rest.li, their common REST framework.
1. [Observability: the new wave or buzzword?](https://medium.com/@dlite/observability-the-new-wave-or-buzzword-fc23a68abf72)

    
1. [Everything You Need to Know About DynamoDB Global Tables](https://engineering.opsgenie.com/everything-you-need-to-know-about-dynamodb-global-tables-952d020d9834)

    OpsGenie analyzes AWS’s new DynamoDB Global Tables, a cross-region multi-master NoSQL datastore. They share the upsides and the pitfalls and include a discussion of how to transition to a global table.
1. [Why, as a Netflix infrastructure manager, am I on call?](https://medium.com/@awspyker/why-as-a-netflix-infrastructure-manager-am-i-on-call-bdc551ac01fe)

    A Netflix manager shares his reasons for still being on-call even though he’s a manager, and they’re pretty great. A lot of it has to do with keeping in tune with what it’s like being a developer on his team, especially with regard to on-call burden and operability.
### Outages

1. [Visual Studio Team Services (Microsoft)](https://blogs.msdn.microsoft.com/vsoservice/?p=15946)
    Microsoft posted an incredibly detailed analysis of an incident that occurred on February 7th. The interesting bit is that they still don’t know what went wrong, and they included a lot of detail on how they’ve tried to track it down so far. Lots to learn from here.
1. [TD Bank](https://delawarebusinessnow.com/2018/02/td-bank-computer-outages-reported-system-upgrade-continues/)
1. [Snapchat](https://www.slashgear.com/snapchat-down-for-many-heres-the-official-advice-15519776/)
1. [Vocus Communications (data center provider)](https://www.arnnet.com.au/article/633314/servers-australia-uproots-services-after-vocus-data-centre-outage/)
1. [Twitter](http://www.ibtimes.com/twitter-down-desktop-mobile-app-not-working-2652671)

### [ << Prev ](sreweekly-109.md) ------------- [ Next >> ](sreweekly-111.md)