## [SRE Weekly Issue #175](https://sreweekly.com/sre-weekly-issue-175/) - June 30, 2019
### Articles

1. [Some Observations On the Messy Realities of Incident Reviews](https://www.adaptivecapacitylabs.com/blog/2019/06/17/some-observations-on-the-messy-realities-of-incident-reviews/)

    This and other enlightened reflections on incident reviews can be found in this article:Richard Cook — Adaptive Capacity Labs
1. [The secret life of DNS packets: investigating complex networks](https://stripe.com/blog/secret-life-of-dns)

    Jeff Jo — Stripe
1. [Multi-Cloud: You keep using that word…](https://discoposse.com/2019/06/19/multi-cloud-you-keep-using-that-word/)

    “Multi-cloud” never really lived up to its hype, did it? This article argues that a multi-cloud strategy is only useful in specific, constrained situations.Disco Posse
1. [How we migrated our database to Amazon Aurora with zero downtime](https://links.lastweekinaws.com/mpss/c/JQE/rDx0AA/t.2sq/mukNMXPXSnupOKK_7X9hqg/h10/u3g-2Bt1LsPzCnacHHvD6jGHWIstAbMzQxMBM8oqScuPV-2FDyI7MIc0Uyy1vlPH4JHt313-2BEU-2FzhNXUE2J2xnazRBCE43y-2BXRYAiDNtDyn-2BMC-2B38KXV4CylXKXW9-2BnB8008ehLrNHCkZPDewDKI5CztfCceHnTfJJsHgrF9Fro-2B1624-2Fu6mMzKnpqyXDrlfruJ7)

    I love how they used idempotency to avoid downtime and missed or repeated transactions during the cutover.Miguel Carranza — RevenueCat
1. [Ebay to hold ‘Crash Sale’ on July 15 in case Amazon’s site goes down](https://www.marketwatch.com/story/ebay-to-hold-crash-sale-on-july-15-in-case-amazons-site-goes-down-2019-06-26)

    This is either really clever or just unsporting.Tonya Garcia — MarketWatch
1. [How SRE teams are organized, and how to get started | Google Cloud Blog](https://cloud.google.com/blog/products/devops-sre/how-sre-teams-are-organized-and-how-to-get-started/)

    This article discusses six kinds of SRE team (“kitchen sink”, infrastructure, tools, product/application, embedded, and consulting) and the pros and cons of each.Gustavo Franco and Matt Brown — Google
1. [When does a reduction in injury numbers become statistically significant?](http://www.safetydifferently.com/when-does-a-reduction-in-injury-numbers-become-statistically-significant/)

    If you see half the incidents this quarter compared to last, does it actually mean anything, statistically speaking? The math in this article applies equally well to SRE, and casts a shadow on the idea of tracking “metrics” like MTTR.Marloes Nitert — Safety Differently
1. [What does debugging a program look like?](https://jvns.ca/blog/2019/06/23/a-few-debugging-resources/)

    This field guide to debugging is the synthesis of a bunch of contributions by folks on Twitter, forged into an article by the inimitable Julia Evans. Maybe a zine is in the works?Julia Evans
### Outages

1. [How Verizon and a BGP Optimizer Knocked Large Parts of the Internet Offline Today](https://blog.cloudflare.com/how-verizon-and-a-bgp-optimizer-knocked-large-parts-of-the-internet-offline-today/)
    The big outage this week happened when a small ISP accidentally told the Internet that it was the best place to send all their packets. Tom Strickx — Cloudflare
1. [Statuspage.io](https://metastatuspage.com/incidents/fnhs9y39zh5s)
1. [Slack](https://status.slack.com/2019-06/9f63d8e30ee85f46)
1. [Hulu](https://consumer.press/hulu-crashes-during-democratic-debate/)
    Hulu suffered an outage during their live stream of an important US political debate.

### [ << Prev ](sreweekly-174.md) ------------- [ Next >> ](sreweekly-176.md)