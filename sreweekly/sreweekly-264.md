## [SRE Weekly Issue #264](https://sreweekly.com/sre-weekly-issue-264/) - April 4, 2021
### Articles

1. [Balancing act: the current limits of AWS network load balancers](https://ably.com/blog/limits-aws-network-load-balancers)

    This well-researched article caught me by surprise. It’s shocking that Ably received advice from AWS to stay under 400,000 simultaneous connections, despite Amazon’s own documentation stating support for “millions of connections per second”.Paddy Byers — Ably
1. [A Journey Into SRE](https://www.algolia.com/blog/a-journey-into-sre/)

    There’s a lot of detail about what their SREs do and how they communicate, with 3 projects as case studies.Sergio Galvan — Algolia
1. [March 2 incident update](https://deno.com/blog/02-03-2021-outage-post-mortem)

    This is an incident followup from an incident at Deno earlier this year. Their CDN saw their heavy use of .ts files (TypeScript, a JavaScript variant) and mistakenly assumed they were MPEG transport segments, a violation of the CDN’s ToS. Oops.Luca Casonato — Deno
1. [Kubernetes Supports Nine Pillars of SRE](https://containerjournal.com/topics/container-ecosystems/kubernetes-supports-nine-pillars-of-sre/)

    Wait, there are 9 now?Marc Hornbeek — Container Journal
1. [Frequently Asked Questions on Deviations](https://www.pharmtech.com/view/frequently-asked-questions-on-deviations)

    There’s a nice little discussion of why “human error” is not a good enough answer for why a deviation (from standard operating procedure) happened.Susan J. Schniepp and Steven J. Lynn — Pharmaceutical Technolog
1. [How To Get Fooled By Metrics](https://tech.trivago.com/2020/12/04/how-to-get-fooled-by-metrics/)

    They deployed an optimization that skipped sending some requests to the backend… and the backend metrics got worse. Why? Hint: aggregate metrics.Dominik Sandjaja — Trivago
### Outages

1. [Twitter](https://financialpost.com/pmn/business-pmn/twitter-says-service-fixed-after-disruption-affects-thousands)
1. [National Weather Service (US)](https://www.washingtonpost.com/weather/2021/03/30/nws-internet-infrastructure-outages/)
1. [reddit](https://reddit.statuspage.io/incidents/jskgtxd0bjqy)
1. [Squarespace.com](https://status.squarespace.com/incidents/69fcntchbl7s)
    Squarespace.com itself, but not user sites.
1. [Microsoft 365](https://twitter.com/MSFT365Status/status/1377771854237069316?ref_src=twsrc%5Etfw%7Ctwcamp%5Etweetembed%7Ctwterm%5E1377788226941845504%7Ctwgr%5E%7Ctwcon%5Es2_&ref_url=https%3A%2F%2Fwww.nxsttv.com%2Fnmw%2Fnews%2Fmicrosoft-users-report-problems-with-multiple-services%2F)

### [ << Prev ](sreweekly-263.md) ------------- [ Next >> ](sreweekly-265.md)