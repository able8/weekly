## [SRE Weekly Issue #145](https://sreweekly.com/sre-weekly-issue-145/) - October 28, 2018
### Articles

1. [To Err Is Human](https://blueskiesmag.com/2018/09/26/to-err-is-human/)

    An article on looking past human error in investigating air sports (definition) accidents, drawing on the writing of Don Norman. Special emphasis on slips versus mistakes:Mara Schmid — Blue Skies Magazine
1. [Five Ways to Tackle Digital Transformation Without Downtime](https://www.informationweek.com/strategic-cio/it-strategy/five-ways-to-tackle-digital-transformation-without-downtime/a/d-id/1332895)

    An VP of NS1 explains how his company rewrote and deployed their core service without downtime.Shannon Weyric — NS1
1. [Status page updates: It’s all about timing](https://blog.hostedgraphite.com/2018/10/16/status-page-updates-its-all-about-timing/)

    This guide from Hosted Graphite has a ton of great advice and reads almost as if they’ve released their internal incident response guidelines. Bonus content: check out this exemplary post-incident followup from their status site.Fran Garcia — Hosted Graphite
1. [Atlassian Incident Management Handbook](https://www.atlassian.com/software/jira/ops/handbook)

    Check it out, Atlassian posted their incident management documentation publicly!
1. [Ten Platform Commandments](https://charity.wtf/2018/10/24/ten-platform-commandments/)

    Charity Majors
1. [Not All Bugs Are Worth Fixing (And That’s Okay)](https://blog.bugsnag.com/application-stability-monitoring/)

    Kristine Pinedo — Bugsnag
1. [Maelstrom: mitigating datacenter-level disasters by draining interdependent traffic safely and efficiently](https://blog.acolyer.org/2018/10/24/maelstrom-mitigating-datacenter-level-disasters-by-draining-interdependent-traffic-safely-and-efficiently/)

    Maelstrom is Facebook’s tool to assist engineers in safely moving traffic off of impaired infrastructure.Adrian Colyer — The Morning Paper (summary)
Veeraraghavan et al. — Facebook (original paper)
1. [Failure differently](http://www.safetydifferently.com/failure-differently/)

    Attempting to stamp out failure entirely can have the paradoxical effect of reducing resiliency to anomalous situations. Instead, we need to handle failure constructively.Daniel Hummerdal — Safety Differently
### Outages

1. [Postmortem: RDS Clogs & Cache-Refresh Crash Loops – Honeycomb](https://www.honeycomb.io/blog/postmortem-rds-clogs-cache-refresh-crash-loops/)
    I guess it’s probably mean of me, but I always get excited when Honeycomb has an outage, because I love reading their followup analyses. This one expertly deconstructs a messy incident with lots of contributing factors.
Rachel Fong — Honeycomb
1. [GitHubGitHub Engineering Adopts New Architecture for MySQL High Availability](https://blog.github.com/2018-10-21-october21-incident-report/)
    GitHub had a severe outage this week. Their brief summary (linked above) brings to mind the mention of the risk of data center isolation in this article from July:

GitHub Engineering Adopts New Architecture for MySQL High Availability

GitHub Engineering Adopts New Architecture for MySQL High Availability
1. [Travis CI](https://www.traviscistatus.com/incidents/z5qn5d4q1nzq)
    Caused by the GitHub outage.
1. [FastlyDegraded Performance in South AfricaElevated Errors in DCA/AshburnCDN Degraded Shield PerformanceElevated Errors in GIG/Rio de JaneiroElevated Errors in HHN/Frankfurt](https://status.fastly.com/incidents/zj8phqj1sz52)
    Fastly had a rough week:

Degraded Performance in South Africa
Elevated Errors in DCA/Ashburn
CDN Degraded Shield Performance
Elevated Errors in GIG/Rio de Janeiro
Elevated Errors in HHN/FrankfurtFull disclosure: Fastly is my employer.

Degraded Performance in South AfricaElevated Errors in DCA/AshburnCDN Degraded Shield PerformanceElevated Errors in GIG/Rio de JaneiroElevated Errors in HHN/FrankfurtFull disclosure: Fastly is my employer.
1. [PagerDuty](https://status.pagerduty.com/incidents/pvkbrkdf7877)
1. [YouTube](https://www.express.co.uk/life-style/science-technology/1036707/YouTube-down-video-streaming-site-not-working-playback-error-message-displayed)
1. [Herokuthis one](http://feedproxy.google.com/~r/HerokuStatus/~3/VJDSXnBZVxk/1651)
    Also this one and a few other minor ones.
1. [Snapchat](https://www.newsweek.com/snapchat-down-not-working-outage-could-not-refresh-1189519)
1. [BitBucketless severe incident](https://status.bitbucket.org/incidents/xr5y9ckb6sls)
    The above is a total outage for one hour. They also had a less severe incident the previous day.
1. [Reddit](https://www.dailystar.co.uk/news/latest-news/738357/reddit-down-not-working-down-detector-social-media-twitter-facebook-instagram)
1. [iCloud](https://www.digitaltrends.com/mobile/apples-icloud-is-experiencing-issues-across-multiple-service/)

### [ << Prev ](sreweekly-144.md) ------------- [ Next >> ](sreweekly-146.md)