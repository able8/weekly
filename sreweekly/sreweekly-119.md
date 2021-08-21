## [SRE Weekly Issue #119](https://sreweekly.com/sre-weekly-issue-119/) - April 29, 2018
### Articles

1. [The STELLA report — Food Fight Podcast](http://foodfightshow.org//2018/04/stella-report.html)

    If you missed the STELLA Report, released last fall during Velocity NYC by John Allspaw, Richard Cook, and David Woods, this podcast is a great intro. And even if you did catch it, it’s well worth a listen. The Food Fight folks interview John Allspaw and there’s some really stellar (see what I did there) back-and-forth discussion.Alan Kraft and Nathen Harvey
1. [Why I usually run ‘w’ first when troubleshooting unknown machines](https://rachelbythebay.com/w/2018/03/26/w/)

    Great idea. This reminds me of a couple jobs back where I rigged up our infrastructure to log every command entered at the shell into a Slack channel.Rachel Kroll
1. [Google: A Collection of Best Practices for Production Services](http://feedproxy.google.com/~r/HighScalability/~3/ZOabXSr-iFI/google-a-collection-of-best-practices-for-production-service.html)

    This excerpt from the Google SRE book is worth reading if only for this nifty idea for graceful degradation:
1. [Walk, talk and git commit: SRE onboarding (2/2)](https://blog.hostedgraphite.com/2018/04/18/walk-talk-and-git-commit-sre-onboarding-2-2/)

    In part two of this story, the author causes their first incident (oops) and subsequently significantly improves the performance of the system in question (cool!).Evan Smith — Hosted Graphite
1. [Blue/Green Deployment: What It Is and How it Reduces Your Risk](https://rollout.io/blog/blue-green-deployment/)

    An introduction to blue/green deployments including the risks it helps to alleviate.Mark Henke — Rollout.io
1. [The Mon-ifesto Part 3: Alert Response and Post-Mortem](https://medium.com/capital-one-developers/the-mon-ifesto-part-3-alert-response-and-post-mortem-cd227c684ac0)

    Peter Christian Fraedrich — Capital One
1. [How to get a core dump for a segfault on Linux](https://jvns.ca/blog/2018/04/28/debugging-a-segfault-on-linux/)

    Especially in Ubuntu, it’s harder than it used to be to get a core dump, thanks to apport and the like.Julia Evans
1. [Disaster recovery sites of exchanges under focus as NCDEX volumes fall](http://www.business-standard.com/article/markets/disaster-recovery-sites-of-exchanges-under-focus-as-ncdex-volumes-fall-118042800285_1.html)

    NCDEX, a stock exchange in Mumbai, India, has been operating out of its disaster recovery site for two weeks. Unfortunately, it looks like performance is not on par with the standard site.Rajesh Bhayani — Business Standard
1. [Southwest 1380 (engine failure 4/17/2018) ENTIRE EVENT: actual multi-sector ATC audio](https://www.youtube.com/watch?v=FkVTdvcghHc)

    You may have heard that a Southwest flight suffered a catastrophic engine failure that left one passenger dead. The day after my family flew a Southwest flight to Orlando. Yikes.The air traffic control audio recording is incredible to listen to. The pilot that was on the radio was cool and calm as she responded to the incident and arranged for landing and emergency ground crews.
### Outages

1. [IRS (US tax system)](https://dzone.com/articles/another-taxing-us-tax-day)
    The IRS had to extend the deadline for Americans to file their taxes as a result of an overload and outage in their electronic tax filing system.
1. [TSB (bank)](http://www.itpro.co.uk/it-infrastructure/30990/major-tsb-outage-affecting-19m-customers-enters-fifth-day)
1. [Herokuthis one](https://status.heroku.com/incidents/1533)
    Also this one.
1. [Google Cloud Pub/Sub](https://status.cloud.google.com/incident/cloud-pubsub/18002#5655608640405504)
1. [Woolworth’s (grocery store chain)](http://www.news.com.au/finance/business/technology/big-business-failure-puts-us-at-risk/news-story/ebe10258a00b22760f47d2e0d65e930e)
1. [Discord](https://discord.statuspage.io/incidents/782x5rl48964)
1. [Fortnite (game)](https://www.thesun.co.uk/tech/6036096/pornhub-fortnite-tweet-server-crash/)
    I normally don’t include games, but this outage is amusing because downtime on Fortnite apparently causes a surge in traffic to a popular adult site, threatening their availability.
1. [Telegram](http://www.ibtimes.com/telegram-down-messaging-app-website-stopped-working-worldwide-back-again-2672876)
1. [Twitter](https://sputniknews.com/science/201804171063650816-twitter-down-social-media/)
1. [TSX (Montreal stock exchange)](https://www.cp24.com/news/tsx-montreal-exchange-down-due-to-technical-problems-1.3905487)

### [ << Prev ](sreweekly-118.md) ------------- [ Next >> ](sreweekly-120.md)