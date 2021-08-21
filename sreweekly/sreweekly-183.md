## [SRE Weekly Issue #183](https://sreweekly.com/sre-weekly-issue-183/) - September 1, 2019
### Articles

1. [I test in prod – Increment: Testing](https://increment.com/testing/i-test-in-production/)

    Another issue of Increment, on a topic integral to SRE: testing.It doesn’t matter if you’ve already read everything Charity Majors has written; in this article she’s still managed to find new and even more compelling ways to argue that we should embrace testing in production.My two other favorite articles from this issue:Charity Majors — Honeycomb
1. [We Re-Launched The New York Times Paywall and No One Noticed](https://open.nytimes.com/we-re-launched-the-new-york-times-paywall-and-no-one-noticed-5cd1f795f76b)

    They rewrote this critical service and carefully deployed it to avoid user impact, using a technique I love: run the new code alongside the old for awhile to verify that it returns the same result.Jeremy Gayed, Said Ketchman, Oleksii Khliupin, Ronny Wang and Sergey Zheznyakovskiy — New York Times
1. [For CAs: What makes a Good Incident Response?](https://groups.google.com/forum/#!topic/mozilla.dev.security.policy/oP8XuNXrANw)

    This is aimed at Certification Authorities dealing with TLS certificate misissuance issues and the like, but it very much applies to any kind of incident.BONUS CONTENT: An incident report from LetsEncrypt just a few days later included this gem, exactly in line with what Ryan wrote:Ryan Sleevi
1. [Our incident postmortem template](https://www.hostedgraphite.com/blog/incident-postmortem-template)

    Whose? Hosted Graphite’s. Definitely worth a read.Fran Garcia — Hosted Graphite
1. [Understanding the risk profile of your technical debt ](https://content.pivotal.io/intersect/risk-profile-of-technical-debt)

    Dormain Drewitz — Pivotal
1. [Building a Culture of Continuous Improvement – Blameless: Better Reliability Through SRE](https://www.blameless.com/resources/workshops/building-a-culture-of-continuous-improvement/)

    Blameless is running a free workshop on writing post-incident reports.Blameless
### Outages

1. [Heroku Status](http://feedproxy.google.com/~r/HerokuStatus/~3/IADu42kfzgs/1892)
    Heroku experienced 8+ hours of instability. This status page posting is really worth a read because it has:

meticulously detailed customer impact
no sugar-coating
detailed workarounds when they were available

Hats off to you, folks.meticulously detailed customer impactno sugar-coatingdetailed workarounds when they were available
1. [Slack](https://status.slack.com/2019-08/a9840f3eddf76b8b)
1. [Reddit](https://reddit.statuspage.io/incidents/zx80r08k1w49)
1. [Sling TV](https://www.cordcuttersnews.com/sling-tv-is-down-for-many-subscribers-right-now/)
1. [Disney Plus](https://www.cnet.com/news/disney-plus-discount-demand-crashes-its-preorder-website/)
    Increased traffic from a sale caused instability.
1. [Fastly](https://status.fastly.com/incidents/hcs8jy0y5vpf)

### [ << Prev ](sreweekly-182.md) ------------- [ Next >> ](sreweekly-184.md)