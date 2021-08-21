## [SRE Weekly Issue #107](https://sreweekly.com/sre-weekly-issue-107/) - January 28, 2018
### Articles

1. [Google Cloud Platform Blog: An example escalation policy — CRE life lessons](http://feedproxy.google.com/~r/ClPlBl/~3/r97ylq1HV0I/an-example-escalation-policy-CRE-life-lessons.html)

    Here, “escalation policy” refers to ongoing work by SRE to get a service back into its SLO, rather than an escalation policy definition in PagerDuty (for example). This article describes the tactics a hypothetical Google SRE team has at their disposal to deal with an ailing service. It’s especially striking to me how this policy comes across as almost punitive in nature.
1. [Now You See Me, Now You Don’t: LinkedIn’s Real-Time Presence Platform](https://engineering.linkedin.com/blog/2018/01/now-you-see-me--now-you-dont--linkedins-real-time-presence-platf)

    
1. [If You’re Going to Fail, Fail Safely](http://blog.launchdarkly.com/if-youre-going-to-fail-fail-safely/)

    This article from LaunchDarkly is about assuming failure and mitigating harm, through the lens of feature-flag-based deployment.
1. [What Tools Do Site Reliability Engineers Use?](https://blog.newrelic.com/2018/01/23/best-sre-tools/?)

    New Relic shares this list of the categories of tools that SREs use to standardize the systems they support.
1. [Building a Distributed Log from Scratch, Part 5: Sketching a New System](https://bravenewgeek.com/building-a-distributed-log-from-scratch-part-5-sketching-a-new-system/)

    In the final article of this series, Tyler Treat lays out a design for a new distributed log based on NSQ.
1. [Observations on the Enterprise of Hiring ](https://honeycomb.io/blog/2018/01/observations-on-the-enterprise-of-hiring/)

    While perhaps not strictly SRE-related, hiring is still critically important for SRE teams. I really love Honeycomb’s approach to hiring as laid out in this blog post.
1. [Why is random testing effective for partition tolerance bugs?](https://blog.acolyer.org/2018/01/23/why-is-random-testing-effective-for-partition-tolerance-bugs/)

    Why indeed? This issue of The Morning Paper discusses a paper on the effectiveness of random testing in distributed systems. More specifically, it goes over the mathematics behind why randomized testing in Jepsen is actually useful, despite classical theories that it ought not be.
### Outages

1. [Pinterest](http://www.ibtimes.com/pinterest-down-not-working-popular-website-unavailable-many-users-2644790)
1. [Google Cloud Storage](https://status.cloud.google.com/incident/storage/18001)
    This one’s worth a read. Google’s original status posting stated 100% impact to cloud storage in its US region, but their followup post retroactively reduced that to 2.0% average and 3.6% peak.
1. [Netflix](https://help.netflix.com/en/is-netflix-down)
    This one happened seemingly at the same time as the Google Cloud Storage outage, but that may be a spurious correlation. This is the first time that I learned that Netflix does have a status page of sorts: it’s an article in their help center entitled “Is Netflix Down?” and they update it live. Who knew?
1. [Facebook/Instagram](https://tribune.com.pk/story/1618454/9-facebook-several-parts-world/)
1. [National Health Service (UK)](http://www.computerweekly.com/news/252433769/NHS-Wales-IT-outage-What-went-wrong-with-its-datacentres)

### [ << Prev ](sreweekly-106.md) ------------- [ Next >> ](sreweekly-108.md)