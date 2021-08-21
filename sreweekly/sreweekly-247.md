## [SRE Weekly Issue #247](https://sreweekly.com/sre-weekly-issue-247/) - December 6, 2020
### Articles

1. [2020 09 25 Incident: Infrastructure connectivity issue impacting multiple systems](https://www.datadoghq.com/blog/2020-09-25-infrastructure-connectivity-issue/)

    This incident report from a September Datadog outage has an interesting tidbit aboiut scaling external incident response in tandem with internal.Alexis Lê-Quôc — Datadog
1. [Google Cloud Issue Summary — Google Drive — 2020-11-16](https://static.googleusercontent.com/media/www.google.com/en//appsstatus/ir/0qduwf33xnjkxpl.pdf)

    This is Google’s write-up for an interesting issue that involved repeated re-sending of invitations to edit a Google Drive document.Google
1. [What I Wish I Knew About Incident Management](https://ronaknathani.com/blog/2020/11/what-i-wish-i-knew-about-incident-management/)

    I basically want to immediately absorb any article with this title, unless it’s just clickbait spam. This one definitely isn’t.Ronak Nathani
1. [Scaling Datastores at Slack with Vitess](https://slack.engineering/scaling-datastores-at-slack-with-vitess/)

    Lots of juicy details in this one about the difficulty Slack has had in scaling their DB layer and how Vitess solved their problems.Arka Ganguli, Guido Iaquinti, Maggie Zhou, and Rafael Chacón — Slack
1. [Mitigate Connection Leaks in Production via Proxies](https://reliability.substack.com/p/mitigate-connection-leaks-in-production)

    Hitting file descriptor limits is such an annoying kind of outage. Some good tips here, clearly coming from hard-won experience.Utsav Shah
1. [Improving the Resiliency of Our Infrastructure DNS Zone](https://blog.cloudflare.com/improving-the-resiliency-of-our-infrastructure-dns-zone/)

    They used two providers synced with OctoDNS.Ryan Timken and Kiran Naidoo — Cloudflare
1. [Root Cause Analysis For Reliability: A Case Study](https://medium.com/last9/root-cause-analysis-for-reliability-a-case-study-8a987ed3a31c)

    This is all about understanding the whole system (people and technology) and building learning, rather than finding a superficial “root cause”.Piyush Verma — Last9
### Outages

1. [Solana](https://cryptobriefing.com/solana-turns-back-on-after-six-hour-outage/)
1. [Poloniex](https://sunriseread.com/crypto-exchange-poloniex-faces-outage-due-to-ddos-attack/124460/)
1. [New Zealand Reserve Bank](https://www.stuff.co.nz/business/123550767/retail-banking-interruptions-as-reserve-bank-payment-system-suffers-threehour-failure)
1. [OneDayOnly](https://www.itweb.co.za/content/mYZRXv9alpEvOgA8/zlP3gQ2qGRMnRD1W)
    Local e-commerce site OneDayOnly is running Black Friday discount deals again today, after the shopping site was down for a few hours last Friday.
1. [Infura](https://sunriseread.com/ethereum-withdrawals-disabled-as-infura-experiences-downtime/121429/)
1. [MobileCause](https://twitter.com/MobileCause/status/1333808645738426368?ref_src=twsrc%5Etfw%7Ctwcamp%5Etweetembed%7Ctwterm%5E1333817587726168065%7Ctwgr%5E%7Ctwcon%5Es2_&ref_url=https%3A%2F%2Fwww.prnewsonline.com%2Fgivingtuesday-outage-mobilecause%2F)
    This outage occurred on Giving Tuesday, a very important day for nonprofits to raise funds.

### [ << Prev ](sreweekly-246.md) ------------- [ Next >> ](sreweekly-248.md)