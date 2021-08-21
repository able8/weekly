## [SRE Weekly Issue #223](https://sreweekly.com/sre-weekly-issue-223/) - June 14, 2020
### Articles

1. [Prevent application and network instability by serving stale content](https://www.fastly.com/blog/prevent-application-network-instability-serve-stale-content)

    I’ve used this technique in the past with a single-page app and a highly-cacheable API, to ensure stability even when the backend goes down.Patrick HamannFull disclosure: Fastly is my employer.
1. [The Impending Doom of Expiring Root CAs and Legacy Clients](https://scotthelme.co.uk/impending-doom-root-ca-expiring-legacy-clients/)

    Here’s a deep dive into how your CA’s certificate can affect your application’s reliability — at least in the eyes of your customers.Scott Helme
1. [[Coinbase] Incident Post Mortem: June 1, 2020](https://blog.coinbase.com/incident-post-mortem-june-1-2020-1cff2e51fa64)

    Here’s Coinbase’s followup from their outage last week.Michael de Hoog — Coinbase
1. [Who’s afraid of serializability?](https://surfingcomplexity.blog/2020/06/13/whos-afraid-of-serializability/)

    Lorin Hochstein
1. [Achieving FMEA goals faster with Chaos Engineering](https://www.gremlin.com//blog/achieving-fmea-goals-faster-with-chaos-engineering/)

    If you’ve been tasked with applying FMEA in your SRE work, this article will get you started.Matthew Helmke
### Outages

1. [IBM Softlayerbit of detail from IBM](https://techcrunch.com/2020/06/09/ibm-cloud-suffers-prolonged-outage/)
    This is the big one this week, with downstream effects on lots of sites and services hosted in Softlayer.
There’s a bit of detail from IBM that seems to indicate that a BGP error by a third party flooded IBM with misrouted traffic.
1. [Reddit](https://reddit.statuspage.io/incidents/y3zj6vp41m1f)
1. [Squarespace](https://status.squarespace.com/incidents/smtsd1khg9l4)
1. [Snapchat](https://twitter.com/snapchatsupport/status/1270479837484859393)
1. [Microsoft Teams](https://twitter.com/MSFT365Status/status/1271088190540124166)

### [ << Prev ](sreweekly-222.md) ------------- [ Next >> ](sreweekly-224.md)