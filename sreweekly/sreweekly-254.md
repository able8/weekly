## [SRE Weekly Issue #254](https://sreweekly.com/sre-weekly-issue-254/) - January 24, 2021
### Articles

1. [Coinbase Incident Post Mortem: January 6–7, 2021](https://blog.coinbase.com/brief-incident-post-mortem-january-6-7-2021-441f6224da93)

    This one’s juicy. At one point, the front-end was blocked up, so the back-end saw less traffic and scaled down. Then when the traffic came flooding back, the back-end was ill-prepared. We can all learn from this.Coinbase
1. [Soar: Simulation for Observability, reliAbility, and secuRity](https://blog.cloudflare.com/soar-simulation-for-observability-reliability-and-security/)

    Cloudflare has what amounts to a sophisticated staging environment for testing new code.Yan Zhai — Cloudflare
1. [Failing to make progress under excess request load](https://rachelbythebay.com/w/2021/01/16/load/)

    Sometimes rolling back doesn’t actually get you back to a good state, especially when there’s pent-up demand.Rachel By the Bay
1. [Google Cloud Issue Summary — Google Meet — 2021-01-08](https://static.googleusercontent.com/media/www.google.com/en//appsstatus/ir/3pc4s2k9hgsdcso.pdf)

    Here’s Google’s follow-up on a Google Meet outage earlier this month.Google
1. [The Next Gen Database Servers Powering Let’s Encrypt ](https://letsencrypt.org/2021/01/21/next-gen-database-servers.html)

    Those are some seriously big database servers.Josh Aas and James Renken — Let’s Encrypt
1. [Incident Management in 2021: from Basics to Best Practices](https://betteruptime.com/blog/incident-management-and-on-call/)

    A great general overview of all aspects of incident response, including definitions and best practices.Better Uptime
1. [Using GPT-3 for plain language incident root cause from logs](https://www.zebrium.com/blog/using-gpt-3-with-zebrium-for-plain-language-incident-root-cause-from-logs)

    Check out what happens when you unleash a generalized language model AI on some log messages related to an incident.Larry Lancaster — Zebrium
1. [Taming Operational Load with VMware CRE](https://tanzu.vmware.com/content/blog/taming-operational-load-vmware-cre)

    The CRE team at VMware undertook a project to find and reduce toil. Note that “with VMware CRE” does not mean “with some product named VMware CRE™”.Gustavo Franco — VMware
1. [Slack RCA for outage on January 4, 2021](https://devopsish.com/pdf/Slack-Incident-Jan-04-2021-RCA-Final.pdf)

    This is Slack’s RCA for their outage earlier this month. This is a great example of a complex incident with many contributing factors — certainly no single “root cause” here.Slack
### Outages

1. [Slack](https://status.slack.com//2021-01/a9b04615180de30a)
1. [Signal](https://www.cnet.com/news/signal-operational-again-after-daylong-outage/)
1. [Apple iCloud](https://telecom.economictimes.indiatimes.com/news/apple-icloud-sign-in-activation-suffer-32-hour-outage-fixed-now/79975310)
1. [Facebook](https://www.complex.com/life/2021/01/facebook-experiences-massive-outage-logs-out-users)
1. [CBS](https://www.sportingnews.com/us/nfl/news/why-is-cbs-down-outage-chiefs-browns/ycr8nlfesesr1a8w8bjf7vluk)

### [ << Prev ](sreweekly-253.md) ------------- [ Next >> ](sreweekly-255.md)