## [SRE Weekly Issue #271](https://sreweekly.com/sre-weekly-issue-271/) - May 23, 2021
### Articles

1. [Naming names in incident writeups](https://surfingcomplexity.blog/2021/05/22/naming-names-in-incident-writeups/)

    Should you keep things anonymous (“an engineer”), or should you say exactly who did what? Here’s a solid argument for the latter.Lorin Hochstein
1. [How Systems Complexity Reduces Uptime](https://which-50.com/how-systems-complexity-reduces-uptime/)

    This article explores the downsides to a design composed of independent parts such as with microservices.Ephraim Baron
1. [User Simulation for Rapid Outage Mitigation ](https://www.infoq.com/presentations/uber-user-simulation/)

    Uber designed a tool they call Blackbox to perform simulated user requests and measure availability. I was struck by the candid discussion of complexity — no one person can understand how all of Uber’s microservices go together.Carissa Blossom — Uber
1. [Nobl9 Makes SLO Specification Open Source](https://devops.com/nobl9-makes-slo-specification-open-source/)

    They’ve made a YAML specification and validator for expressing SLOs in a machine-readable format.Mike Vizard — Devops.com
1. [What is Observability](http://www.brendangregg.com/blog/2021-05-23/what-is-observability.html)

    A new spin: this one makes the distinction between “experimental tools” that affect the state of the system, and “observability tools” that are read-only.Brendan Gregg
1. [The Incident Review: 4 Odd Incidents Caused by Animals](https://rootly.io/blog/the-incident-review-4-odd-incidents-caused-by-animals)

    “Contributing factors: moose and squirrel.”JJ Tang — Rootly
1. [When Debuggers Lie](https://tech.okcupid.com/when-debuggers-lie-700e40f11891)

    Every once in awhile, I need to pull out gdb. In times like those, it’s useful to have this kind of thing floating around in the back of my mind.Brendon Scheinman — okcupid
### Outages

1. [Slack](https://status.slack.com/2021-05/33aed870657e30a0)
1. [Colonial Pipeline](https://www.independent.co.uk/news/world/americas/colonial-pipeline-cyberattack-shutdown-b1849644.html)
    The same major US oil pipeline mentioned here last week is still having network issues.
1. [Binance and Coinbase](https://www.benzinga.com/markets/cryptocurrency/21/05/21198280/major-crypto-exchanges-binance-and-coinbase-down-as-crypto-selloff-intensifies-gas-fees-sk)
1. [YouTube](https://9to5google.com/2021/05/18/youtube-playback-error-outage-may18/)
1. [Sabre](https://www.11alive.com/article/travel/long-lines-at-atlanta-airport-after-server-issue-affects-check-ins-around-the-country/85-b502967d-8abe-4c04-aaa2-75fb7b2595cc)
    Sabre is a backend service provider used by a lot of airlines.
1. [Azure web portal](https://status.azure.com/en-us/status/history/)
    There’s an interesting followup post about a DNS issue.

### [ << Prev ](sreweekly-270.md) ------------- [ Next >> ](sreweekly-272.md)