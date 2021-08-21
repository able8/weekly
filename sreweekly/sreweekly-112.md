## [SRE Weekly Issue #112](https://sreweekly.com/sre-weekly-issue-112/) - March 4, 2018
### Articles

1. [Spooky action at a distance, how an AWS outage ate our load balancer](https://blog.hostedgraphite.com/2018/03/01/spooky-action-at-a-distance-how-an-aws-outage-ate-our-load-balancer/)

    I don’t think I even need to add anything to that to make you want to read this article.Fran Garcia — Hosted Graphite
1. [GitHub’s report on the Memcached-based DDoS](https://githubengineering.com/ddos-incident-report/)

    The big story this week is the memcached UDP amplification DDoS method, used to send 1.3 Tbps (!) toward our friends at GitHub. Their description is linked above.Sam Kottler — GitHub
The internet was alight with related discussions:
1. [Runbook Template](https://github.com/CaitieM20/Talks/blob/master/TacklingAlertFatigue/runbook.md)

    An excellent template that you can use as a basis for writing runbooks.Catie McCaffrey
1. [DevOps and SRE Contribution – The Lemur Book](https://lemurbook.com/devops-and-sre-contribution/)

    This author of an upcoming O’Reilly book is looking for small contributions for a crowd-sourced chapter:David Blank-Edelman
1. [Meet Bandaid, the Dropbox service proxy](https://blogs.dropbox.com/tech/2018/03/meet-bandaid-the-dropbox-service-proxy/)

    I’m intrigued by the way it handles its queue in last-in first-out order, on the theory that a request that’s been waiting for a long time is likely to be cancelled by its requester.Dmitry Kopytkov and Patrick Lee — Dropbox
1. [5 of the world’s biggest network outages](https://www.itproportal.com/features/5-of-the-worlds-biggest-network-outages/)

    This is an amusing recap of five major outages of the past few years. If you’ve been subscribed for awhile, it’ll be review, but I still enjoyed the reminder.Michael Rabinowitz
1. [Fail-slow at scale: When the cloud stops working](http://www.zdnet.com/article/how-clouds-fail-slow/)

    This article summarizes a new research paper on “fail-slow” hardware failures. When hardware only kind of fails, it can often have more disastrous consequences than when it stops working outright. Robin Harris — Storage Bits
1. [Launching An Entire Fireworks Display At Once](https://www.youtube.com/watch?v=dabnx8VSdkE&feature=youtu.be)

    This is an awe-inspiring way to make a point about designing systems to be resilient to human error.Tom Scott
1. [Managing Feature Flag Debt with Split](https://www.split.io/blog/managing-feature-flag-debt-split/)

    There are some really great ideas in this article around preventing and ameliorating the technical debt that can be inherent in the use of feature flags. Ostensibly this article is about using Split.io, but the ideas are broadly applicable.Adil Aijaz — Split
### Outages

1. [Slackhangops](https://status.slack.com/2018-03/07e49841fd4c74f8)
    Possibly due to the AWS outage.Thanks to marc on hangops #incident_response for this one.
1. [AWS us-east-1](https://www.geekwire.com/2018/widespread-outage-amazon-web-services-u-s-east-region-takes-alexa-atlassian-developer-tools/)
1. [Statuspage.io](https://metastatuspage.com/incidents/6fbn11c94j3n)
1. [Abebooks](https://tamebay.com/2018/03/abebooks-experienced-another-outage-this-week.html)
1. [LinkedIn](http://www.newsweek.com/linkedin-down-not-working-why-site-login-826916)
1. [AOL Email](https://www.express.co.uk/life-style/science-technology/926377/AOL-down-not-working-e-mail-AOL-Mail)
1. [Vero (social network)](https://www.inverse.com/article/41672-vero-down-increased-demand-new-users)
1. [CoinsMarkets (cryptocurrency exchange)](https://themerkle.com/coinsmarkets-exchange-resumes-withdrawals-after-major-outage/)

### [ << Prev ](sreweekly-111.md) ------------- [ Next >> ](sreweekly-113.md)