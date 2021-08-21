## [SRE Weekly Issue #136](https://sreweekly.com/sre-weekly-issue-136/) - August 26, 2018
### Articles

1. [The 18 ghosts in your infrastructure stack that can cause failure (and how to avoid them)](https://blog.ably.io/the-18-ghosts-in-your-infrastructure-stack-that-can-cause-failure-and-how-to-avoid-them-1373daebfe1a)

    This infographic shows how Ably’s client library and backend infrastructure is designed to work around many common failure modes. My favorite: they have redundant TLS certificates from distinct issuers.Matthew O’Riordan — Ably
1. [QA Instability Implies Production Instability](http://michaelnygard.com/blog/2016/07/qa-instability-implies-production-instability/)

    This article argues that spending a little time to fix staging can make production significantly more stable.Michael Nygard
1. [Through a Dashboard Darkly – Brain of Buildchimp – Curated Selections from the Voices in My Head](https://buildchimp.me/Through-a-Dashboard-Darkly/)

    […]John Casey
1. [Simple/hard metrics that help reduce MTTR when looking for a root cause](https://blog.okmeter.io/simple-hard-metrics-that-help-reduce-mttr-when-looking-for-a-root-cause-637eaf5b3518)

    A story of a somewhat rare failure case (a datacenter heat buildup event) and how to monitor for such a thing without contributing to metrics overload.Pavel Trukhanov — okmeter
1. [Shipping Software Should Not Be Scary – charity.wtf](https://charity.wtf/2018/08/19/shipping-software-should-not-be-scary/)

    Great advice, useful for managers and individual contributors.Charity Majors
### Outages

1. [Apple CloudKit](https://9to5mac.com/2018/08/24/cloudkit-outage-data-loss/)
    There appears to be some prolonged issues with Apple’s CloudKit service today, which Apple offers to developers as a way to store user data and sync across devices. Several developers have reported to us that they have seen data for their apps temporarily wiped in the last 24 hours as the CloudKit service experiences some form of outage.
1. [Heroku](https://status.heroku.com/incidents/1617?utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A+HerokuStatus+%28Heroku+%7C+Status%29)
1. [Commonwealth Bank (AU)](https://www.heraldsun.com.au/news/victoria/commonwealth-bank-online-and-visa-transactions-down/news-story/38c5bb1da1414ffeb2e7b22a6ff21637)
1. [Coles (Supermarket chain)](https://www.itnews.com.au/news/coles-suffers-nationwide-it-outage-500469)
1. [Sydney, AU train system](https://www.dailytelegraph.com.au/news/nsw/commuter-chaos-as-technical-issue-causes-major-delays-on-sydney-trains/news-story/3ec4c50ff3d1a1eb816a93d816318b82)
1. [redditanother one](https://reddit.statuspage.io/incidents/3rfmfn4vxxrr)
    And another one.

### [ << Prev ](sreweekly-135.md) ------------- [ Next >> ](sreweekly-137.md)