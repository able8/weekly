## [SRE Weekly Issue #126](https://sreweekly.com/sre-weekly-issue-126/) - June 17, 2018
### Articles

1. [Introducing Grab-Kit: Distributed Service Design at Grab](https://engineering.grab.com/introducing-grab-kit)

    Karen Kue and Michael Cartmell — Grab
1. [10 ways to avoid CDN issues at peak](https://blog.peer5.com/avoid-cdn-issues-at-peak/)

    Some tips on surviving peak traffic as we head into World Cup season. I like the discussion in #10 (load testing): accurately testing your CDN is all but impossible.Hadar Weiss — Peer 5 (CDN)
1. [A story of being on call](https://www.youtube.com/watch?v=p_paJ2PB4MY)

    This is a video recording of a talk by Charity Majors at Monkigras 2018. She has a lot of awesome stuff to say about making on-call enjoyable and owning your code, including this gem:Charity Majors — Honeycomb
1. [Heroku Incident #1561 Followup (Platform-wide Outages)](https://status.heroku.com/incidents/1561)

    Presumably this was the us-east-1 power issue I reported on in Issue 124.
1. [Future of Reliability Engineering (Part 1)](http://michael-kehoe.io/post/the-future-of-reliability-engineering-part1/)

    The first article in this new series is about the evolution of the Network Engineer into a Network Reliability Engineer. It’s part of the broader breakdown of silos with the goal of understanding holistic reliabilty.Michael Kehoe
1. [Protecting network availability for GDPR compliance](https://www.itproportal.com/features/protecting-network-availability-for-gdpr-compliance/)

    I hadn’t realized that GDPR has provisions related to site/service reliability.Theresa Abbamondi — Netscout
1. [Sample Your Traffic: But Keep The Good Stuff](https://www.youtube.com/watch?v=oRVy0Y0qcCo)

    To shamelessly steal a line from this recorded talk, it’s very rarely the right thing for your observability system’s scale to match that of the system it’s observing. To avoid that, you need to throw away some event data rather than storing and indexing everything. How do you do that while still achieving functioning observability?Ben Hartshorne — Honeycomb
1. [Automated Database Deployments Series Kick Off](https://octopus.com/blog/automated-database-deployments-series-kick-off)

    I’m looking forward to seeing where this article series goes. Database changes can be a huge reliability risk, and getting them right is critical.Bob Walker — Octopus Deploy
### Outages

1. [Azure south-central US region](https://azure.microsoft.com/en-us/status/history/)
    A load spike in a backend storage system caused impact across a range of Azure services, according to the RCA linked above.
Actually, I’ve linked to their generic “status history” page, since that seems to be as specific as I can get. Readers from Microsoft, perhaps you could ask the folks that run the Azure status page to create dedicated permalinks for each incident, or at least for each RCA? Even an anchor link in the status history page would be super-awesome!
1. [New Relic infrastructure alerting](https://status.newrelic.com/incidents/knxgjwsmm6wz)
1. [Travis CI](https://www.traviscistatus.com/incidents/l4gjpkbwj885)
1. [WhatsApp](https://www.rt.com/news/429708-whatsapp-down-worldwide-outages/)
1. [American Airlines](http://www.newsweek.com/american-airlines-system-website-down-not-working-checkin-tickets-975439)
1. [Instagram](https://studybreaks.com/news-politics/instagram-power-outage/)
1. [Google Compute Engine](https://status.cloud.google.com/incident/compute/18005#18005006)
    While instances were stopped (shut down), newly-launched instances were allowed to take their IPs. The stopped instances then failed on startup due to the IP conflicts. The situation lasted for around 20 hours.
1. [Optus Sport](https://www.9news.com.au/technology/2018/06/15/23/01/world-cup-optus-coverage-dropping-in-and-out-viewers-angry)
    World Cup fans had issues watching through Optus. World Cup streaming traffic is massive this time around.
1. [Apple Maps](https://thenextweb.com/apple/2018/06/15/apple-maps-outage-gave-people-an-excuse-to-stop-using-apple-maps/)
1. [Netflix](http://www.nme.com/news/tv/netflix-fans-panic-global-streaming-issues-2337571)
1. [.my TLD](http://www.freemalaysiatoday.com/category/nation/2018/06/16/internet-glitch-that-brought-down-my-sites-almost-solved/)

### [ << Prev ](sreweekly-125.md) ------------- [ Next >> ](sreweekly-127.md)