## [SRE Weekly Issue #4](https://sreweekly.com/sre-weekly-issue-4-2/) - January 2, 2016
### Articles

1. [MoonGen](https://github.com/emmericp/MoonGen)

    A nifty-looking packet generator with packets crafted by Lua scripts. If this thing lives up to the hype in its documentation, it’d be pretty awesome! Thanks to Chris Maynard for the link and for the sleepless days and nights we spent mucking with trafgen’s source.
1. [sysadvent: Day 24 – It’s not Production without an audit trail](http://feedproxy.google.com/~r/sysadvent/~3/loqKJj-7QfA/day-24-its-not-production-without-audit.html)

    Just as we design systems to be monitored, this article suggests that we should design systems to be audited. Doing the work up front and incrementally rather than as an afterthought can take the pain out of auditing.
1. [Structured Logging](http://kartar.net/2015/12/structured-logging/)

    A nice intro to structured logging. I’m a big fan of ELK, and especially using Logstash to alert on events that might be difficult to catch otherwise.
1. [Back to Black Friday: Performance Testing Lessons From Target](https://blazemeter.com/blog/back-black-friday-performance-testing-lessons-target)

    I looked at a few “lessons learned from black Friday 2015” articles, but they’re all low on good technical detail. My consolation prize is this article that seems eerily appropriate, given Target’s outage on Cyber Monday.The strategy of turning away only some requesters to avoid a full site outage is interesting, but I could see it causing a thundering herd problem if not done carefully, where folks just repeatedly hit reload and cause more traffic.
1. [Top five predictions for load testing in 2016](http://www.information-age.com/technology/mobile-and-networking/123460715/top-five-predictions-load-testing-2016)

    These “predictions” (suggestions, really) about load testing may be review to some, but this article caught my interest because it was the first time I’d heard the term Performance Engineering. Definitely a field worth paying attention to as it becomes more prevalent due to its overlap with SRE.
1. [The Nature of Human Error: Implications for Surgical Practice](http://www.ncbi.nlm.nih.gov/pmc/articles/PMC1856596/#!po=0.458716)

    Modern medicine has been working through very similar issues to SRE, related to controlling the impact of human error through process design and analysis of human factors. We stand to learn a lot from articles such as this one. For example, they’ve been doing the “blameless retrospective” for a long time:

1. [[Steam] Update on Christmas Issues](http://store.steampowered.com/news/19852/?snr=1_550_552&utm_source=twitterfeed&utm_medium=twitter)

    A speedy and detailed postmortem from Valve on the Steam issue on Christmas.
### Outages

1. [EA Games](http://gearnuke.com/ea-servers-go-company-issues-apology/)
1. [New FAA drone registration site](https://www.yahoo.com/tech/faa-45-000-drones-registered-104503132.html)
1. [PlayStation Network](http://www.express.co.uk/entertainment/gaming/629365/PSN-down-PlayStation-Network-PS4-Why-is-PSN-down-DDOS-attack-Phantom-Squad)
1. [Roku](http://www.theverge.com/2015/12/25/10665846/christmas-2015-online-service-outages)
1. [Valve’s Steam Store](http://www.biztekmojo.com/001812/steam-store-server-back-online-after-major-security-glitch-stemming-caching-issues)
1. [Servers Of Electronic Arts Hit By Short Outage](http://empowerednews.net/servers-of-electronic-arts-hit-by-short-outage/1865079/)
1. [Time Warner Cable (US ISP)](http://www.ubergizmo.com/2015/12/time-warner-cable-suffered-national-outage-yesterday/)
1. [iiNet (AU ISP)](http://www.theregister.co.uk/2015/12/30/iinet_xmas_outage/)
    A 5-day outage for some customers of Australian ISP iiNet.
1. [BBC](http://www.techweekeurope.co.uk/e-marketing/bbc-iplayer-outage-ddos-attack-182934)
    In a hilarious twist, BBC’S official statement of “technical issues” was contradicted by a BBC story citing “sources inside BBC” and blaming the outage on a DDoS.
1. [WhatsApp](http://www.dailymail.co.uk/sciencetech/article-3380408/WhatsApp-goes-Users-Europe-report-problems-connecting-chats-messaging-app.html)
1. [USPTO (US Patent & Trademark Office)](http://www.uspto.gov/about-us/news-updates/statement-uspto-acting-chief-communications-officer-patrick-ross)
    The USPTO suffered a severe power incident that destroyed their power conditioning units and had to close early for the holidays.

### [ << Prev ](sreweekly-3.md) ------------- [ Next >> ](sreweekly-5.md)