## [SRE Weekly Issue #155](https://sreweekly.com/sre-weekly-issue-155/) - January 13, 2019
### Articles

1. [Developer On Call](https://henrikwarne.com/2018/12/03/developer-on-call/)

    A developer’s perspective on why being on call is important and how to structure it fairly (hint: compensation).Henrik Warne
1. [Interpreting Kafka’s Exactly-Once Semantics](https://dzone.com/articles/interpreting-kafkas-exactly-once-semantics)

    The Conclusion section sums it up nicely: Rahul Agarwal — DZone
1. [DevOps Discussions: Postmortem Chat – Part 1 (YouTube)](https://www.youtube.com/watch?v=TkA6u8qEugc&feature=youtu.be)

    This is a riveting discussion about retrospective analysis of incidents, hosted by Microsoft. Throughout the discussion, there’s an emphasis on learning from incidents as opposed to simply coming up with action items.Note: one of the panelists is my fellow employee at Fastly.Jessica DeVita — Microsoft, with Duck Lawn (Pushpay), Tom Griffin (Pushpay), Sue Allspaw Pomeroy (Fastly), John Allspaw (Adaptive Capatacity Labs) and Dr. Richard Cook (Adaptive Capacity Labs)
1. [An Agile SRE Meeting Plan](https://tech.mangot.com/blog/2019/01/09/an-agile-sre-meeting-plan/)

    If you’re looking for a blueprint of how to structure your SRE organization’s meetings, this is a great resource. Dave Mangot
1. [Designing resilient systems: Circuit Breakers or Retries? (Part 2)](https://engineering.grab.com/designing-resilient-systems-part-2)

    This article is really thorough and includes a section on combining retries with circuit breakers.Corey Scott — Grab
1. [Towards Successful Resilient Software Design](https://www.infoq.com/articles/towards-resilient-software-design)

    Uwe Friedrichsen — InfoQ
1. [Courier: Dropbox migration to gRPC](https://blogs.dropbox.com/tech/2019/01/courier-dropbox-migration-to-grpc/)

    This really stood out to me:Ruslan Nigmatullin and Alexey Ivanov — Dropbox
### Outages

1. [Fastlyrepeat](https://status.fastly.com/incidents/w9f6c4krv5yt)
    Fastly had the above issue in its MDW PoP and also a repeat.
Full disclosure: Fastly is my employer.
1. [Zoom](https://status.zoom.us/incidents/7mcj28xfdhh1)
1. [Slack](https://status.slack.com/2019-01/65dd2378d9565469)
    Android notifications were busted.
1. [Gov Availabilityexpired TLS certificates](https://govavailability.info/)
    This site shows a live-updated availability percentage for the US Government. As of now, the “This Year” percentage is stuck at infinite zeroes (due to our current government shutdown). On a less tongue-in-cheek note, lots of US Government sites have expired TLS certificates because employees aren’t there to renew them.
1. [Duo Security](https://status.duo.com/incidents/mbc689bdjmy8)
1. [Azure Storage (UK south region)](https://www.theregister.co.uk/2019/01/10/microsoft_uk_south_wobble/)
1. [GitHub](https://www.githubstatus.com/incidents/6rbyhd6vrc0k)
1. [Reddit](https://reddit.statuspage.io/incidents/zw6khxbl5qk2)
1. [YouTube](https://9to5google.com/2019/01/08/psa-youtube-down/)
1. [Tinder](https://www.cnet.com/news/tinder-down-for-some-users-looking-to-swipe-this-morning/)
1. [Google Cloud Platform (various API functions)](http://status.cloud.google.com/incident/appengine/19001#19001008)
    Google engineers began rolling out a new feature designed to improve the fault-tolerance of the metadata store.
Ironically, that rollout took down the metadata store.

### [ << Prev ](sreweekly-154.md) ------------- [ Next >> ](sreweekly-156.md)