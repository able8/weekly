## [SRE Weekly Issue #125](https://sreweekly.com/sre-weekly-issue-125/) - June 10, 2018
### Articles

1. [Don’t use Go’s default HTTP client (in production)](https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779)

    Go’s HTTP client defaults to no timeout. Making HTTP requests with no timeout is rarely a good idea and has been at the heart of many incidents I’ve been involved in.Nathan Smith
1. [What DBAs need to know about Cloud Spanner, part 1: Keys and indexes](https://cloudplatform.googleblog.com/2018/06/What-DBAs-need-to-know-about-Cloud-Spanner-part-1-Keys-and-indexes.html)

    A few times now, I’ve made offhand comments about how Spanner promises a lot and I’d like to know what the catches are. Here they are! In all fairness, they’re pretty reasonable constraints to work with.Niel Markwick and Robert Saxby — Google
1. [A Postmortem Template](http://michael-kehoe.io/post/postmortem-template/)

    I’d refer to this as more of a retrospective template, but in any case, it’s pretty nifty!Michael Kehoe
1. [What happens when online games go down for maintenance](https://www.pcgamer.com/what-happens-when-online-games-go-down-for-maintenance/)

    This is a news report rather than a technical deep-dive. It’s got some pretty interesting (and amusing) stories from various MMOs. Alex Wiltshire — PC Gamer
1. [Lessons from Building Observability Tools at Netflix](https://medium.com/netflix-techblog/lessons-from-building-observability-tools-at-netflix-7cfafed6ab17?source=rss----2615bd06b42e---4)

    Here’s how Netflix does observability.Kevin Lew and Sangeeta Narayanan — Netflix
1. [Heroku Status](https://status.heroku.com/)

    Looks like I’ve missed a few incident followup posts from Heroku in the past couple months:#1548: Increased errors in starting dynos
#1535: Post-incident Dyno Restarts
#1459: Scheduled API Maintenance on Monday March 26 at 23:00 UTC (4:00 PM PT)’
#1413: Dyno Availability
#1414: Heroku Connect Sync Delays
#1395: Heroku Connect Availability
#1393: Heroku Connect unavailable
#1379: Dyno boot issues
### Outages

1. [Walt Disney World Website and My Disney Experience Mobile App](https://wdwnt.com/2018/06/walt-disney-world-website-and-my-disney-experience-mobile-app-completely-unavailable-all-fastpass-dining-and-resort-reservations-affected-2/)
    Having just been to Disney World in April, I can attest to the severity of this kind of outage and the importance of the app.
1. [Paytm (digital wallet service)](https://www.techgenyz.com/2018/06/06/paytm-app-down-users-complains/)
1. [ASX (Australian Stock Exchange)](https://www.afr.com/business/banking-and-finance/financial-services/asx-back-online-after-fire-alarm-bungle-20180604-h10ytt)
    Inadvertent release of fire suppression gas damaged some equipment and halted trading.
1. [Instagram](https://www.express.co.uk/life-style/science-technology/969534/instagram-down-is-instagram-still-down-today-when-will-it-be-back)
1. [LSE (London Stock Exchange)](https://www.telegraph.co.uk/business/2018/06/07/london-stock-exchange-fails-open-rare-disruption/)
1. [Twitter](https://www.hindustantimes.com/tech/twitter-down-for-many-users-says-it-s-over-capacity/story-fKe5a2LTEh3g1gfSpKpliN.html)
1. [Today we mitigated 1.1.1.1](https://blog.cloudflare.com/today-we-mitigated-1-1-1-1/)
    On May 31, 2018 we had a 17 minute outage on our 1.1.1.1 resolver service; this was our doing and not the result of an attack.
Cloudflare shares some detail on what went wrong in this comprehensive incident analysis.
 Marek Majkowski — Cloudflare

### [ << Prev ](sreweekly-124.md) ------------- [ Next >> ](sreweekly-126.md)