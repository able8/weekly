## [SRE Weekly Issue #191](https://sreweekly.com/sre-weekly-issue-191/) - October 27, 2019
### Articles

1. [The Post-Incident Review Issue 1: Autumn/Winter 2019](https://zine.incidentlabs.io)

    Check it out! A new zine dedicated to post-incident reviews. This first issue includes a reprint of 4 real gems from the past month plus one original article about disseminating lessons learned from incidents.Emil Stolarsky and Jaime Woo
1. [New – Amazon CloudWatch Anomaly Detection](https://aws.amazon.com/blogs/aws/new-amazon-cloudwatch-anomaly-detection/)

    I swear, it’s like they heard me talking about anomaly detection last week. Anyone used this thing? I’d love to hear your experience. Better still, perhaps you’d like to write a blog post or article?
1. [CPDoS: Cache Poisoned Denial of Service](https://cpdos.org)

    I know this isn’t Security Weekly, but this vulnerability has the potential to cause reliability issues, and it’s dreadfully simple to understand and exploit.Hoai Viet Nguyen and Luigi Lo Iacono
1. [Behind the Scenes of a long EVE Online downtime [2015]](http://community.eveonline.com/news/dev-blogs/behind-the-scenes-of-a-long-eve-online-downtime/)

    In this incident followup from the archives, read the saga of a deploy gone horribly wrong. It took them hours and several experiments to figure out how to right the ship.CCP Goliath — EVE Online
1. [Nine Experimentation Best Practices](https://launchdarkly.com/blog/nine-experimentation-best-practices/)

    The best practices:Dawn Parzych — LaunchDarkly
1. [Open Sourcing Mantis: A Platform For Building Cost-Effective, Realtime, Operations-Focused Applications](https://medium.com/netflix-techblog/open-sourcing-mantis-a-platform-for-building-cost-effective-realtime-operations-focused-5b8ff387813a?source=rss----2615bd06b42e---4)

    Mantis uses an interesting stream processing / subscriber model for observability tooling.Cody Rioux, Daniel Jacobson, Jeff Chao, Neeraj Joshi, Nick Mahilani, Piyush Goyal, Prashanth Ramdas, and Zhenzhong Xu — Netflix
1. [Deploy on Fridays, or Don’t.](https://hackernoon.com/deploy-on-fridays-or-dont-qg2y32jk)

    We can’t ever be sure deploy will be safe, but we can be sure that folks have plans for their weekend.David Mangot — Mangoteque
### Outages

1. [Amazon Route 53](https://status.aws.amazon.com)
    Route 53 had significant DNS resolution impairment.
Their status site still doesn’t allow deep linking or browsing the archive in any kind of manageable way, so here’s the full text of their followup post:
On October 22, 2019, we detected and then mitigated a DDoS (Distributed Denial of Service) attack against Route 53. Due to the way that DNS queries are processed, this attack was first experienced by many other DNS server operators as the queries made their way through DNS resolvers on the internet to Route 53. The attack targeted specific DNS names and paths, notably those used to access the global names for S3 buckets. Because this attack was widely distributed, a small number of ISPs operating affected DNS resolvers implemented mitigation strategies of their own in an attempt to control the traffic. This is causing DNS lookups through these resolvers for a small number of AWS names to fail. We are doing our best to identify and contact these operators, as quickly as possible, and working with them to enhance their mitigations so that they do not cause impact to valid requests. If you are experiencing issues, please contact us so we can work with your operator to help resolve.
1. [Heroku](https://status.heroku.com/incidents/1916)
    I’m guessing this stemmed from the Route 53 incident.
Our infrastructure provider is currently reporting intermittent DNS resolution errors. This may result in issues resolving domains to our services.
1. [Twitter](https://www.express.co.uk/life-style/science-technology/1194100/Twitter-Down-Offline)
1. [Yahoo Mail](https://www.express.co.uk/life-style/science-technology/1194153/Yahoo-Mail-down-not-working-desktop-Android-iPhone-server-status-updates)
1. [Hosted Graphite](https://status.hostedgraphite.com/incidents/r8dgxtgsjf7p)
1. [Discord](https://discord.statuspage.io/incidents/7sq5g4fnhtdz)
1. [Google Cloud Platformpossible Google search outage](http://status.cloud.google.com/incident/cloud-networking/19020#19020005)
    Many GCP services were affected. There was also a possible Google search outage, though I wasn’t able to corroborate this report.

### [ << Prev ](sreweekly-190.md) ------------- [ Next >> ](sreweekly-192.md)