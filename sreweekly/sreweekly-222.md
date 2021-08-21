## [SRE Weekly Issue #222](https://sreweekly.com/sre-weekly-issue-222/) - June 7, 2020
### Articles

1. [Meaningful availability: How many nines do you actually need?](https://techbeacon.com/enterprise-it/meaningful-availability-how-many-nines-do-you-actually-need)

    This article in a nutshell:Kolton Andrus — Gremlin
1. [ Byzantine and non-Byzantine distributed systems](https://ayende.com/blog/191073-A/byzantine-and-non-byzantine-distributed-systems?Key=26f97b3c-f827-4408-832f-6ae183ec2698)

    I hadn’t heard of this distinction before. If you haven’t either, click through to find out more.Ayende Rahien — RavenDB
1. [Using SRE to meet reliability challenges](https://cloud.google.com/blog/products/management-tools/meeting-reliability-challenges-with-sre-principles/)

    Cheryl Kang — Google
1. [Faulty Equipment, Lapsed Training, Repeated Warnings: How a Preventable Disaster Killed Six Marines](https://www.propublica.org/article/marines-hornet-squadron-242-crash-pacific-resilard)

    ProPublica picks apart the incident in exhaustive detail, showing how multiple problems interwoven in the organization contributed to this tragedy. Robert Faturechi, Megan Rose and T. Christian Miller — ProPublica
1. [SRE, CSE, and the safety boundary](https://surfingcomplexity.blog/2020/05/25/sre-cse-and-the-safety-boundary/)

    There’s a great review of Rasmussen’s safety boundary model, which I wasn’t previously familiar with. A system moves between three boundaries:Lorin Hochstein
1. [The Tail at Scale Revisited](https://billduncan.org/the-tail-at-scale-revisited/)

    This one includes a really nifty graph showing how reliable your N backend microservices need to be in order to hit a given reliability target R.Bill Duncan
1. [Oncall and COVID-19 Survey Results](https://www.firehydrant.io/blog/oncall-and-covid-19-survey-results/)

    Here are the results of the survey I linked here a couple weeks ago. There are some interesting and surprising results, well worth a read.Rich Burroughs — FireHydrant
1. [The mystery of the expiring Sectigo web certificate](https://nakedsecurity.sophos.com/2020/06/02/the-mystery-of-the-expiring-sectigo-web-certificate/)

    A commonly-used CA’s Root certificate expired, causing some havoc. Even though Sectigo did everything right, some software didn’t handle the transition to the new root well.Paul Ducklin — Naked Security
### Outages

1. [PagerDuty](https://status.pagerduty.com/incidents/yw8s50fbm6s2)
1. [Coinbase](https://blog.coinbase.com/incident-post-mortem-june-1-2020-1cff2e51fa64)
    Coinbase had an outage on June 1. Click for their post-incident analysis.
1. [Robinhoodstatus page](https://androidgram.com/robinhood-down-both-app-website-are-having-server-issues-users-unable-to-transfer-funds/)
    Robinhood’s status page doesn’t show history, so I can’t verify this one.
1. [iCloud](https://telanganatoday.com/apple-fixes-outage-preventing-users-from-accessing-icloud-storage)
1. [Ebaystatus page](https://www.thesun.co.uk/tech/10729859/ebay-down-not-working-offline-outage-2020/)
    Ebay’s status page also doesn’t show history, so I can’t verify this one either.
1. [Lloyds and Halifax (bank)](https://www.thescottishsun.co.uk/money/5654832/lloyds-bank-halifax-app-website-down-customers-locked/)
1. [Adobe Cloud](https://www.theregister.co.uk/2020/05/27/adobe_cloud_outage/)
1. [Squarespace](https://status.squarespace.com/incidents/1wg7g5tc5jy2)
    Their followup post discusses the large-scale DDoS that contributed to the outage.
1. [HostedGraphite](https://status.hostedgraphite.com/incidents/f545qq2nck4x)
1. [Telegram](https://twitter.com/telegram/status/1269379445674127362)

### [ << Prev ](sreweekly-221.md) ------------- [ Next >> ](sreweekly-223.md)