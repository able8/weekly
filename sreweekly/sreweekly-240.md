## [SRE Weekly Issue #240](https://sreweekly.com/sre-weekly-issue-240/) - October 18, 2020
### Articles

1. [Google Cloud Issue Summary — Google Chat — 2020-09-17](https://static.googleusercontent.com/media/www.google.com/en//appsstatus/ir/ti76kvgai1yqajx.pdf)

    This interesting post-incident analysis is marked as “Google Customer Confidential – Not for publication or distribution”, but Google linked it directly from their public status page. I normally would not include a seemingly “leaked” incident report like this, but in this case I think the “confidential” label is erroneous.Google
1. [40 milliseconds of latency that just would not go away](https://rachelbythebay.com/w/2020/10/14/lag/)

    I keep re-learning and re-forgetting about TCP_NODELAY.Rachel By the Bay
1. [“Manual” and “Automated” are just words](http://michaelnygard.com/blog/2020/10/manual-and-automated-are-just-words/)

    The distinction between the two is a lot more nuanced than it may seem. What are we really trying to say wit those words?Michael Nygard
1. [Heroku incident #2110 follow-up](https://status.heroku.com/incidents/2110)

    This incident from the week before last involved a Let’s Encrypt API rate limit.
1. [Fixing Linux filesystem performance regressions](https://engineering.linkedin.com/blog/2020/fixing-linux-filesystem-performance-regressions)

    Don’t you hate when you’re minding your own business upgrading your OS, and you run smack into a kernel bug in the ext4fs code?Ryan Underwood — LinkedIn
1. [Identifying and protecting against the largest DDoS attacks](https://cloud.google.com/blog/products/identity-security/identifying-and-protecting-against-the-largest-ddos-attacks/)

    Google discusses DDoS attacks and how they deal with them, including a 2.5Tbps attack in 2017.Damian Menscher — Google
1. [How I Broke `git push heroku main`](https://blog.heroku.com/how-i-broke-git-push-heroku-main)

    I love these first-hand incident stories. This one is from an engineer at Heroku who was a contributing factor in an incident last month.Damien Mathieu — Heroku (Salesforce)
### Outages

1. [BitBay](https://www.financemagnates.com/cryptocurrency/exchange/crypto-exchange-bitbay-suffers-service-outage/)
1. [Twittertaken down purposefully to protect a US presidential election candidate](https://www.androidcentral.com/as-if-2020-hadnt-ruined-enough-twitter-down-again)
    It definitely was not taken down purposefully to protect a US presidential election candidate.
1. [TikTok](https://metro.co.uk/2020/10/12/tiktok-down-across-the-uk-australia-and-parts-of-europe-13408958/)
1. [Crunchyroll](https://www.distractify.com/p/why-is-crunchyroll-not-working)
1. [Instagram](https://todaynewstalk.com/2020/10/12/instagram-down-trending-on-twitter-users-said-the-app-experiencing-service-outage/)
1. [Barnes and Noble](https://www.fastcompany.com/90564581/barnes-noble-apologizes-for-serious-network-issue-after-days-long-nook-disruptions)
    Nook e-readers have experienced a days’-long service disruption.
1. [keepthescorediscussion](https://keepthescore.co/blog/posts/deleting_the_production_database/)
    Linked is their blog post, “We deleted the production database by accident”.
Be sure to check out the HackerNews discussion about this article, too.
Caspar — Keepthescore
1. [FanDuel](https://twitter.com/FanDuel_Support/status/1315325128138461185)
    This incident seems to be ongoing, October 12 to present.

### [ << Prev ](sreweekly-239.md) ------------- [ Next >> ](sreweekly-241.md)