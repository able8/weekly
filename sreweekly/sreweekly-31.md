## [SRE Weekly Issue #31](https://sreweekly.com/sre-weekly-issue-31/) - July 17, 2016
### Articles

1. [Improve your Sleep while On-Call](https://blog.serverdensity.com/on-call-sleep-monitoring/)

    Opzzz is a new app that graphs sleep data (from a Fitbit) against pager alerts (from PagerDuty or Server Density).  I love this idea!

1. [Opsweekly: Measuring on-call experience with alert classification](https://codeascraft.com/2014/06/19/opsweekly-measuring-on-call-experience-with-alert-classification/)

    Speaking of measuring sleep data against pages, Etsy is doing that too with their open source on-call analysis tool Opsweekly.  Engineers also classify their events based on whether they were actionable.

1. [Dealing with Anxiety in Operations](http://www.slideshare.net/jlintz/dealing-with-anxiety-in-operations-velocity-2016)

    Slides from a talk on a really important topic.  There are some great resource links included.
1. [How to Work an On Call Job and Keep Your Sanity](http://lifehacker.com/5983847/how-to-work-an-on-call-job-and-keep-your-sanity)

    I’m a firm believer in work/life balance, especially as it pertains to on-call.  I have a reputation for rigidly defending my personal time and that of my co-workers.  I strongly feel that this is the best thing I can do for my company because exhaustion and burnout are huge reliability risks.  Read this article if you’re trying to figure out how to improve your on-call experience and aren’t sure how to start.
1. [Making Facebook self-healing: Automating proactive rack maintenance](https://code.facebook.com/posts/629906427171799/making-facebook-self-healing-automating-proactive-rack-maintenance/)

    FBAR, Facebook’s Auto-Remediation system, was mentioned here last month.  This week, they posted an update explaining AMH, their system for safely handing maintenance of blocks of servers.
1. [[Pingdom] Post-mortem for recent incidents](http://royal.pingdom.com/2016/07/12/post-mortem-recent-incidents/)

    Pingdom released this set of short postmortems for last week’s series of outages.
1. [From idea to reality: containers in production at GoCardless](https://gocardless.com/blog/from-idea-to-reality-containers-in-production-at-gocardless/)

    A really detailed article about how one company got Docker into production safely and reliably.  I especially love the parts about nginx cutover (when deploying new container versions) and supervising running containers.  With the common refrain that Docker isn’t ready for production, it’s nice to see how GoCardless did it — but it’s also interesting to see how much tooling they felt compelled to write in-house.
1. [The true meaning of availability](https://datacenternews.asia/story/expert-insights-true-meaning-availability/)

    What good is an arbitrary number of nines from a cloud service provider if their transit links go down?  Or if vast swathes of end-users can’t reach your site due to a major internet disruption?  ServiceNow’s vice president argues that cloud providers must pay attention to “real availability” and partner with their customers to deal with external threats to availability.
1. [Bitfinex Outages Raise Questions of Reliability and Regulation](http://bitcoinist.net/bitfinex-outages-reliability-regulation/)

    Last month, Bitfinex (a bitcoin exchange) experienced multiple outages, and the subsequent bitcoin sell-off caused the price of the bitcoin to drop 7.5%.  Bitcoin’s lack of regulation is a blessing, but is it also a curse?
1. [Learning from Failure at Etsy](http://www.kitchensoap.com/2013/09/30/learning-from-failure-at-etsy/)

    How can I even intro a gem like this? John Allspaw’s essay on blameless and just culture at Etsy is a classic, and it’s a great read even if you’re well-versed in the topic.  I especially liked the concept of the “Second Story”.
### Outages

1. [Fasthosts](http://www.theregister.co.uk/2016/07/11/major_fasthosts_outage/)
1. [Comcast Phone](http://www.cruiserschoice.com/news/comcast-phone-outage-knock-business-customers-offline-nationwide-r1237/)
    Comcast is a cable ISP in the US that also offers VoIP land line phone service.
1. [Pokémon Go](http://heavy.com/games/2016/07/pokemon-go-server-servers-status-are-the-down-right-now-how-to-check-outage-map/)
    Multiple outages this week.
1. [NOAA (US weather service)](https://www.washingtonpost.com/news/capital-weather-gang/wp/2016/07/13/weather-service-suffers-major-network-issue-warning-system-interrupted/)
    NOAA suffered a system outage and also published a bogus flood alert.
1. [AT&T (telecom)](https://historiccity.com/2016/staugustine/news/florida/how-do-you-reach-out-without-internet-cellular-or-phone-59449)
1. [Charter (ISP)](http://www.polkio.com/news/2016/jul/13/charter-outage-puts-businesses-tough-spot/)
1. [PlayStation Network](http://thetechportal.com/2016/07/13/playstation-network-seems-many-worldwide/)
1. [Lloyds](http://www.theregister.co.uk/2016/07/14/llloyds_online_banking_hit_with_intermitent_outage/)
1. [SGX (Singapore exchange)](http://www.todayonline.com/business/sgx-downtime-thursday-triggered-hardware-issue)
1. [Iraq](http://www.techweekeurope.co.uk/cloud/iraq-shuts-down-internet-195163)
    Iraq’s government shut down internet access in response to protests.
1. [Vodacom](http://mybroadband.co.za/news/cellular/172057-massive-vodacom-data-outage-in-south-africa.html)
1. [EA Games](http://www.dailystar.co.uk/tech/gaming/530616/EA-Games-Down-Server-Crash-FIFA-Battlefront-Battlefield-crash-PS4-Xbox-One-PCEA)

### [ << Prev ](sreweekly-30.md) ------------- [ Next >> ](sreweekly-32.md)