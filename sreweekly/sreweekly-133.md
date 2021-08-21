## [SRE Weekly Issue #133](https://sreweekly.com/sre-weekly-issue-133/) - August 5, 2018
### Articles

1. [Errata: miscredited article in last week’s issue](http://sreweekly.com/sre-weekly-issue-132/)

    My sincerest apology to Ali Haider Zaveri, author of the article Location-Aware Distribution: Configuring servers at scale. I originally miscredited the article to two folks, claiming they were from Facebook when in fact they work at Google.
1. [Cloud infrastructure at Grubhub](https://bytes.grubhub.com/cloud-infrastructure-at-grubhub-94db998a898a)

    As Grubhub built out their service-oriented architecture, they first developed “base frameworks for building highly available, distributed services”.William Blackie — Grubhub
1. [How we scaled nginx and saved the world 54 years every day](https://blog.cloudflare.com/how-we-scaled-nginx-and-saved-the-world-54-years-every-day/)

    Cloudflare discusses an optimization that improves their p99 response time in the face of occasionally slow disk access. Today I learned: Linux does not allow for non-blocking disk reads. Ka-Hing Cheung — Cloudflare
1. [Google Online Security Blog: Chrome’s Plan to Distrust Symantec Certificates](https://security.googleblog.com/2017/09/chromes-plan-to-distrust-symantec.html)

    I include this article not just to warn you in case you depend on GeoTrust certificates, but also to highlight what’s involved in running a reliable and trustworthy CA.Devon O’Brien, Ryan Sleevi, and Andrew Whalley — Google
1. [How we built a globally distributed rate limiter that scales horizontally](https://blog.ably.io/how-adopting-a-distributed-rate-limiting-helps-scale-your-platform-1afdf3944b5a?source=rss----63d285882655--realtime_tech)

    They go over the 6 key constraints that influenced their design and describe the solution they came up with. Some of the constraints seem to involve preserving not just their own systems’ reliability, but that of their customers’ systems.Simon Woolf — Ably
1. [Repairing network hardware at scale with SRE principles](http://feedproxy.google.com/~r/ClPlBl/~3/SjVNWkWit9k/repairing-network-hardware-at-scale-with-sre-principles.html)

    James O’Keeffe — Google
1. [Understanding Azure Load Balancing Solutions – Azure Load Balancer, Azure Application Gateway and Azure Traffic Manager – Rahul Rajat Singh’s blog](http://rahulrajatsingh.com/2018/07/understanding-azure-load-balancing-solutions-azure-load-balancer-azure-application-gateway-and-azure-traffic-manager/)

    Rahul Rajat Singh
### Outages

1. [Google Cloud Networking europe-west2](http://status.cloud.google.com/incident/cloud-networking/18014#18014005)
1. [GitHub](https://status.github.com/messages)
1. [Facebook](https://www.thesun.co.uk/tech/6937553/facebook-down-not-working-crashed-website-broken/)
1. [FastMail](https://www.theregister.co.uk/2018/08/01/fastmail_takes_a_nap_in_network_blunder/)
1. [Chipotle](https://www.kansas.com/news/nation-world/national/article215836735.html)
    It appears that Chipotle may have DoSed themselves with an offer of free guacamole to folks that order online.
1. [MoviePass](https://deadline.com/2018/07/moviepass-manic-monday-ticket-service-outage-again-at-major-chains-1202436719/)

### [ << Prev ](sreweekly-132.md) ------------- [ Next >> ](sreweekly-134.md)