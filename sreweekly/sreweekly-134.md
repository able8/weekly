## [SRE Weekly Issue #134](https://sreweekly.com/sre-weekly-issue-134/) - August 12, 2018
### Articles

1. [What Do I Need To Know about “SegmentSmack”](https://isc.sans.edu/forums/diary/What Do I Need To Know about SegmentSmack/23964/)

    The big news this week is SegmentSmack, a denial of service vulnerability in the Linux kernel that allows an attacker to cause high CPU consumption. Linked is a SANS Technology Institute researcher’s summary of the attack. Other coverage:Johannes B. Ullrich, PhD — SAN Technology Institute
1. [How to manage changing requirements for a high availability service](https://bytes.grubhub.com/how-to-manage-changing-requirements-for-a-high-availability-service-4715750576bc)

    It’s rare that any system we create will remain static throughout its lifetime. How can you handle retrofitting it without sacrificing reliability?Yiwei Liu — Grubhub
1. [GLB: GitHub’s open source load balancer](https://githubengineering.com/glb-director-open-source-load-balancer/)

    Theo Julienne — GitHub
1. [The Secrets of Load-balancing Long Lived TCP Connections](https://blog.hostedgraphite.com/2018/08/07/the-secrets-of-load-balancing-long-lived-tcp-connections/)

    HostedGraphite had a load-balancing challenge: some connections carried 5 data points per second while others had 5000. Here’s how they solved it.Ciaran Gaffney — HostedGraphite
1. [How we designed the Quotas microservice to prevent resource abuse](https://engineering.grab.com/quotas-service)

    Here’s how Grab designed their global rate-limiting system, ensuring nearly instant local rate-limiting decisions controlled asynchronously by a global service.Jim Zhan and Gao Chao — Grab
1. [Envoy Service Mesh Case Study: Mitigating Cascading Failure at Lyft](https://www.infoq.com/articles/envoy-service-mesh-cascading-failure)

    Find out how Lyft avoids cascading failure in their microservice-based architecture, through the use of a client- and server-side rate-limiting proxy. Daniel Hochman and Jose Nino — Lyft
1. [Post-mortems to the rescue](https://increment.com/documentation/post-mortems-to-the-rescue/)

    Sweta Ackerman — Increment
1. [John Oliver viewers, not hackers, responsible for FCC system outage](https://www.foxbusiness.com/media/john-oliver-viewers-not-hackers-responsible-for-fcc-system-outage)

    The FCC blamed their outage this past May on a DDoS. Turns out it was just massively distributed requests for legitimate service.Thomas Barrabi — Fox Business
1. [You can’t debug systems with dashboards](https://read.acloud.guru/why-you-cant-effectively-debug-your-modern-systems-with-dashboards-57fe3ecd26bf)

    My favorite part of this interview with Charity Majors is the discussion of operations in a serverless infrastructure (toward the end).Forrest Brazeal — A Cloud Guru
### Outages

1. [Travis CI](https://www.traviscistatus.com/incidents/118bjpkgkg00)
1. [Google G Suite administrator console](https://www.google.com/appsstatus#hl=en&v=issue&sid=0&iid=1cb41110feccbe06b3596f7b58effa05)
1. [Datadog](https://status.datadoghq.com/incidents/g43bjgbxxs0t)
1. [Google Compute Engine](https://status.cloud.google.com/incident/cloud-networking/18013#18013006)
    This is a followup analysis of an outage that occurred on July 27.
The issue was caused by an unintended side effect of a configuration change […]

### [ << Prev ](sreweekly-133.md) ------------- [ Next >> ](sreweekly-135.md)