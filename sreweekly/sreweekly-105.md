## [SRE Weekly Issue #105](https://sreweekly.com/sre-weekly-issue-105/) - January 14, 2018
### Articles

1. [Page It Forward!](https://www.pagerduty.com/blog/page-it-forward/)

    PagerDuty put a call out on Twitter, asking what folks are doing to improve the on-call experience at their companies.
1. [Building a Distributed Log from Scratch, Part 3: Scaling Message Delivery](https://bravenewgeek.com/building-a-distributed-log-from-scratch-part-3-scaling-message-delivery/)

    Here’s part three in the series. This one’s about sharding, horizontal scaling, and client versus server complexity.
1. [What are Azure Availability Zones and why should you use them](https://blogs.technet.microsoft.com/uktechnet/2018/01/08/what-are-azure-availability-zones-and-why-should-you-use-them/)

    Here’s how Azure’s new availability zones change the way highly available apps can be designed on Azure.
1. [Dealing with the Meltdown patch at Grab](http://engineering.grab.com/dealing-with-the-meltdown-patch-at-grab)

    The meltdown patch seems to be having a disproportionate impact on Redis performance. Here’s Grab’s story of how they figured out what was up and what they did to deal with it.
1. [Twitter: mipsytipsy on monitoring vs observability](https://mobile.twitter.com/mipsytipsy/status/951343352359763968)

    I don’t often do the Twitter thing, but this chain by Charity Majors is worth reading. Is that what they call it? a chain?
1. [ Google Cloud Platform Blog: Why you should pick strong consistency, whenever possible
](https://cloudplatform.googleblog.com/2018/01/why-you-should-pick-strong-consistency-whenever-possible.html)

    Google on the advantages of Cloud Spanner’s strong consistency and why to use it. I’m still looking out for an explanation of what the downside to Spanner is…
1. [Machine Learning Drives Changing Disaster Recovery At Facebook](https://www.nextplatform.com/2018/01/10/machine-learning-drives-changing-disaster-recovery-facebook/)

    Just to be clear, this is about how critical it is that Facebook keep their machine learning applications running, rather than using machine learning to design disaster recovery solutions.
1. [Planning Better for Failure: How Mainframe Error Messages Impact CX](https://compuware.com/mainframe-error-messages-impact-cx/)

    This article is about useful error messages, which are important both for the customer experience and for operations. I’m not sure what really qualifies as a “mainframe” these days, though….
1. [Automating Your Oncall: Open Sourcing Fossor and Ascii Etch](https://engineering.linkedin.com/blog/2017/12/open-sourcing-fossor-and-ascii-etch)

    LinkedIn is open-sourcing two tools that they use for troubleshooting during incidents. Fossor automates running data-gathering can and Ascii Etch displays graphs using ASCII art.
### Outages

1. [LastPass](https://www.lastpass.com/)
1. [Slack](https://status.slack.com/2018-01/d8cf1517de9ecfa8)
1. [Spotify](http://www.ibtimes.com/spotify-status-error-website-down-customer-service-contact-info-2639243)
1. [Bitbucket](https://www.theregister.co.uk/2018/01/10/bitbucket_outage/)
    Bitbucket has had severe performance problems due to a failure in their storage layer.
1. [Kraken (cryptocurrency exchange)](https://cointelegraph.com/news/kraken-down-nearly-48-hours-gives-engineers-time-to-rest-before-resuming-service)
    This appears to have been a scheduled upgrade that blew up in complexity, preventing Kraken from coming back up for two days. From the article:
Most astonishing of all, about 36 hours after the upgrade began, Kraken apparently sent their engineers home to take a nap!
Not that astonishing! Tired engineers make mistakes, after all.
1. [Missile threat alert for Hawaii a false alarm](http://edition.cnn.com/2018/01/13/politics/hawaii-missile-threat-false-alarm/index.html)
    There’s so much more to this story than we’ve been told, and I really wish I could be a fly on the wall during the retrospective.

### [ << Prev ](sreweekly-104.md) ------------- [ Next >> ](sreweekly-106.md)