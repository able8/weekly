## [SRE Weekly Issue #143](https://sreweekly.com/sre-weekly-issue-143/) - October 14, 2018
### Articles

1. [ How Etsy Handles Peeking in A/B Testing](https://codeascraft.com/2018/10/03/how-etsy-handles-peeking-in-a-b-testing/)

    There’s some great statistics theory in here. The challenge is: how can you have accurate, useful A/B tests without having to wait a long time to get a big enough sample size? Can you bail out early if you know the test has already failed? Can you refine the new feature mid-test?Callie McRee and Kelly Shen — Etsy
1. [SRE: The Biggest Lie Since Kanban | the agile admin](https://theagileadmin.com/2018/10/02/sre-the-biggest-lie-since-kanban/)

    Don’t just rename your Ops team to “SRE” and expect anything different, says this author.Ernest Mueller — The Agile Admin
1. [We can do better than percentile latencies | theburningmonk.com](https://theburningmonk.com/2018/10/we-can-do-better-than-percentile-latencies/)

    Great idea:Yan Cui
1. [Dropbox traffic infrastructure: Edge network](https://blogs.dropbox.com/tech/2018/10/dropbox-traffic-infrastructure-edge-network/#.W7438mlAiGc.twitter)

    There’s a ton of detail here, and it’s a great read. Lots of juicy tidbits about PoP selection, load balancing, and performance monitoring.Oleg Guba and Alexey Ivanov — DropboxFull disclosure: Fastly, my employer, is mentioned.
1. [Preliminary Report Pipeline: Over-pressure of a Columbia Gas of Massachusetts Low-pressure Natural Gas Distribution System](https://www.ntsb.gov/investigations/AccidentReports/Pages/PLD18MR003-preliminary-report.aspx)

    Even as a preliminary report there’s a lot to digest here about what caused the series of gas explosions last month in Massachusetts (US). I feel like I’ve been involved in incidents with similar contributing factors.US National Transportation Safety Board (NTSB)
1. [What I learned by bringing down LinkedIn.com – VentureBeat](https://venturebeat.com/2018/10/13/what-i-learned-by-bringing-down-linkedin-com/)

    This isn’t just a recap of a bad day, although the outage description is worth reading by itself. Readers also gain insight into the evolution of this engineer’s career and mindset, from entry-level to Senior SRE.Katie Shannon — LinkedIn
1. [https://about.gitlab.com/2018/10/11/gitlab-com-stability-post-gcp-migration/](https://about.gitlab.com/2018/10/11/gitlab-com-stability-post-gcp-migration/)

    GitLab, in their trademark radically open style, goes into detail on the reasons behind the recent increase in the reliability of their service. Andrew Newdigate — GitLab
1. [Getting to 99.999% Availability with Twilio’s Tyler Wells ](https://blameless.com/uncategorized/getting-to-99-999-availability-with-twilios-tyler-wells/)

    Five nines are key when you consider that Twilio’s service uptime can literally mean life and death. Click through to find out why.Charlie Taylor — Blameless
### Outages

1. [Travis CI](https://www.traviscistatus.com/incidents/t9ls519frgww)
1. [Google Compute Engine us-central1-c](https://status.cloud.google.com/incident/cloud-networking/18016#18016002)
    I can’t really summarize this incident report one well, but I highly recommend reading it.
1. [Azure](https://azure.microsoft.com/en-us/status/)
    Duplicated here since I can’t deep-link:
Summary of impact: Between 01:22 and 05:50 UTC on 13 Oct 2018, a subset of customers using Storage in East US may have experienced intermittent difficulties connecting to resources hosted in this region. Other services leveraging Storage in the region may have also experienced impact related to this incident.
1. [Instagram](http://infosurhoy.com/cocoon/saii/xhtml/en_GB/health/instagram-website-and-app-go-down-with-users-unable-to-log-in/)
1. [Heroku](https://status.heroku.com/incidents/1635?utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A+HerokuStatus+%28Heroku+%7C+Status%29)
    This one’s notable for the duration: about 10 days of diminished routing performance due to a bad instance.
1. [Microsoft Outlook](https://www.newsweek.com/outlook-email-down-not-working-1157587)

### [ << Prev ](sreweekly-142.md) ------------- [ Next >> ](sreweekly-144.md)