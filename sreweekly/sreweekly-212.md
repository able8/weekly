## [SRE Weekly Issue #212](https://sreweekly.com/sre-weekly-issue-212/) - March 22, 2020
### Articles

1. [Meaningful availability](https://blog.acolyer.org/2020/02/26/meaningful-availability/)

    Hauer et al. — NSDI’20 (original paper)Adrian Colyer — The Morning Paper (summary)
1. [Our Top 5 On-Call Practices – Blameless: Better Reliability Through SRE](https://www.blameless.com/our-top-5-on-call-practices/)

    Their top 5 are:Emily Arnott — Blameless
1. [NTP: Building a more accurate time service at Facebook scale](https://engineering.fb.com/production-engineering/ntp-service/)

    Synchronizing clocks can be critical in an HA system, and Facebook went to great lengths to ensure clock accuracy.Zoe Talamantes and Oleg Obleukhov — Facebook
1. [The Fallacy of Move Fast and Break Things](https://launchdarkly.com/blog/the-fallacy-of-move-fast-and-break-things/)

    You might end up just breaking things.Dawn Parzych — LaunchDarkly
1. [InSearch: LinkedIn’s new message search platform](https://engineering.linkedin.com/blog/2020/introducing-insearch)

    LinkedIn’s message search system takes advantage of the fact that relatively few users actually search their message. It only builds a search index the first time a user performs a search. Suruchi Shah and Hari Shankar — LinkedIn
1. [Destiny 2 Outage and Rollback](https://www.bungie.net/en/News/Article/48723)

    This followup post from Bungie covers two related incidents in February that caused loss of user data.Bungie
1. [Involving Engineers in Incident Management: QCon London Q&A ](https://www.infoq.com/news/2020/03/engineers-incident-management/)

    An interview about how one company got their developers to join the on-call rotation. It covers how they trained them to help them build confidence and what benefits they got by joining.Ben Linders — InfoQ
### Outages

1. [Statuspage.iounrelated outage](https://metastatuspage.com/incidents/2pl8snrgv60x)
    The text of this incident originally mentioned Heroku, and it lines up with the Heroku outage below.They also had this unrelated outage.
1. [HerokuIncident #1961Incident #1968](https://status.heroku.com/incidents/1974)
    Heroku suffered two short bouts of 85% request failure to applications hosted on their platform.Separately, they recently posted a couple of followup reports for previous incidents:

Incident #1961: logging outage
Incident #1968: EU application errors

Incident #1961: logging outageIncident #1968: EU application errors
1. [Zoom](https://status.zoom.us/incidents/xgbq4dghxc1h)
1. [MacStadium](https://status.macstadium.com/incidents/l464rrkdg4yn)
1. [Hulu](https://popculture.com/streaming/2020/03/20/coronavirus-hulu-down-subscriber-reactions/)
1. [Bumble](https://www.republicworld.com/technology-news/apps/is-bumble-down-right-now.html)
1. [Microsoft Teams and Office 365](https://mspoweruser.com/microsoft-cloud-outages-extends-to-office-365-with-some-services-inaccessible/)
1. [Discord](https://discord.statuspage.io/incidents/62gt9cgjwdgf)
    Discord posted this gem of a followup analysis just a few days after their outage last week.
1. [GoToMeeting](https://status.gotomeeting.com/incidents/kqrc9yd3zhrz)
1. [Google Nest](https://9to5google.com/2020/03/18/nest-outage-march-2020/)
1. [DoorDash](https://twitter.com/DoorDash_Help/status/1241041482145914880)

### [ << Prev ](sreweekly-211.md) ------------- [ Next >> ](sreweekly-213.md)