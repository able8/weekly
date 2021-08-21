## [SRE Weekly Issue #115](https://sreweekly.com/sre-weekly-issue-115/) - March 25, 2018
### Articles

1. [Moving Past Shallow Incident Data](http://www.adaptivecapacitylabs.com/blog/2018/03/23/moving-past-shallow-incident-data/)

    Metrics like Mean Time to Detection (MTTD), Resolution (MTTR), and the like pave over all of the incredibly valuable details of the individual incidents. If you place a lot of emphasis on aggregate incident response metrics, this article may cause you to rethink your methods.John Allspaw — Adaptive Capacity Labs
1. [Look for the duct tape](https://rachelbythebay.com/w/2018/03/23/ducttape/)

    Duct tape: you know, all the little shell scripts you have in your ~/bin directory that you wrote because your system’s tooling got in your way or didn’t do what you needed? Find that, according to this article, and you’ll find interesting things to work on to make the system better. I’d add that these rough edges are often also the kinds of things that contribute to incidents.Rachel Kroll
1. [Incident review: API and Dashboard outage on 10 October 2017](https://gocardless.com/blog/incident-review-api-and-dashboard-outage-on-10th-october/)

    A thoughtful and detailed incident post-analysis, including an in-depth discussion of the weeks-long investigation to determine the contributing factors. The outage involved the interaction of Pacemaker and Postgres.Chris Sinjakli , Harry Panayiotou , Lawrence Jones , Norberto Lopes and Raúl Naveiras — GoCardless
1. [How Chaos Engineering Can Bring Stability to Your Distributed Systems](https://thenewstack.io/chaos-engineering-can-give-distributed-systems-stability/)

    Here’s a nice overview of chaos engineering, including a mention of a tool I wasn’t aware of for applying chaos to Docker containers.Jennifer Riggins — The New Stack
1. [Pull doesn’t scale – or does it?](https://prometheus.io/blog/2016/07/23/pull-does-not-scale-or-does-it/)

    The question in the title refers to the gathering of metrics from many systems in an infrastructure. Do they push their metrics in, or should the system pull metrics from each host instead? This Prometheus author explains why they pull and how it scales.Julius Volz — Prometheus
1. [Zero downtime deployments with containers](http://blog.kontena.io/zero-dowtime-deployments-with-containers/)

    A primer on achieving seamless deployments with Docker, including examples.Jussi Nummelin — Kontena
1. [observability – Food Fight](http://foodfightshow.org/2018/03/observability.html)

    I had some extra time for reviewing content this week, and I took the opportunity to listen to this episode of the Food Fight podcast, with a focus on observability. The discussion is really excellent, and there are some really thought-provoking moments.Nell Shamrell-Harrington, with Nathen Harvey, Charity Majors, and Jamie Osler
1. [Enable your Devs to do Ops](https://blog.buildo.io/enable-your-devs-to-do-ops-9a0a870baa1)

    How? By writing runbooks. This article takes you through how, why, and what tools to use as you develop runbooks for your systems.Francesco Negri — Buildo
1. [How Threat Stack Does DevOps (Part IV): Making Engineers Accountable](https://www.threatstack.com/blog/how-threat-stack-does-devops-part-iv-making-engineers-accountable/)

    As a security-focused company, it only makes sense that Threat Stack would focus on safety when giving developers access to operate their software production.Pete Cheslock — Threat Stack
### Outages

1. [Data Action](https://www.itnews.com.au/news/dozens-of-aussie-banks-go-down-after-telstra-outage-487252)
    Data Action is a dependency of many Australian banks.
1. [Travis CI](https://www.traviscistatus.com/incidents/jww6stq6pk4s)
1. [S3VPC Endpoints](http://status.aws.amazon.com/)
    Amazon S3 had a pair of outages for connections through VPC Endpoints. The Travis CI, Datadog, and New Relic outages were around the same time, but I can’t tell conclusively whether they were related.
1. [Datadog](https://status.datadoghq.com/incidents/8vpg5zvlnsht)
1. [New Relic](https://status.newrelic.com/incidents/63jvb51dyj6q)

### [ << Prev ](sreweekly-114.md) ------------- [ Next >> ](sreweekly-116.md)