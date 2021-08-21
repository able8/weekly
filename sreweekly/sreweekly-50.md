## [SRE Weekly Issue #50](https://sreweekly.com/sre-weekly-issue-50/) - December 11, 2016
### Articles

1. [sysadvent: Day 1 – Why You Need a Postmortem Process](http://sysadvent.blogspot.com/2016/12/day-1-why-you-need-postmortem-process.html?m=1)

    Gabe Abinante has been featured here previously for his contributions to the Operations Incident Board: Postmortem Report Reviews project. To kick off this year’s sysadvent, here’s his excellent argument for having a defined postmortem process.
1. [sysadvent: Day 4 – Change Management: Keep it Simple, Stupid](http://sysadvent.blogspot.com/2016/12/day-4-change-management-keep-it-simple.html?m=1)

    Having a change management process is useful, even if it’s just a deploy/rollback plan. I knew all that, but this article had a really great idea that I hadn’t thought of before (but should have): your rollback plan should have a set of steps to verify that the rollback was successful.
1. [sysadvent: Day 6 – No More On-Call Martyrs](http://sysadvent.blogspot.com/2016/12/day-6-no-more-on-call-martyrs.html?m=1)

    Let’s be honest: being on-call is kind of an ego boost. It makes me feel important. But not getting paged is way better than getting paged, and we should remember that. #oncallselfie
1. [The State of On-Call 2016-2017 — Kicking off Results Season](https://victorops.com/blog/state-call-2016-2017-devops-webinar/)

    It’s that time of year again! In a long-standing (1-year-long) tradition here at SRE Weekly, I present you this year’s State of On-Call report from my kind sponsor, VictorOps.
1. [The Problem with Preaggregated Metrics: Part 1, the “Pre”](https://honeycomb.io/blog/2016/12/the-problem-with-preaggregated-metrics-part-1-the-pre/)

    Storing 99th and 95th percentile latency in your time-series DB is great, but what if you need a different percentile? Or if you need to see why those 1% of requests are taking forever? Perhaps they’re all to the same resource?
1. [Orchestrator at GitHub](http://githubengineering.com/orchestrator-github/)

    Orchestrator is a tool for managing a (possibly complex) tree of replicating MySQL servers. This includes master failure detection and automatic failover, akin to MHA4Mysql and other tools. GitHub has adopted Orchestrator and shares some details on how they use it.
1. [Black Friday and Cyber Monday Performance Report 2016](https://dzone.com/articles/black-friday-and-cyber-monday-performance-report-2)

    A few notable brands suffered impaired availability on and around Black Friday this year. Hats off to AppDynamics for giving us some hard numbers.
1. [Microsoft refuses to join the Zero Outage brigade, Google and AWS keep mum](http://www.theregister.co.uk/2016/12/02/microsoft_refuses_to_join_the_zero_outage_brigade_google_and_aws_keep_mum/)

    Looks like I missed this “Zero Outage Framework” announcement the first time around. I love the idea of information-sharing and it’ll be interesting to see what they come up with. We can all benefit from this, especially if the giants like Microsoft join up.
1. [HumanOps: Etsy on how unclear workplace expectations contribute to staff burnout](http://www.computerweekly.com/news/450403359/HumanOps-Etsy-web-ops-boss-explains-how-inflated-workplace-expectations-contribute-to-staff-burnout)

    All IT managers would do well to heed this advice. Remember, burnout very often directly and indirectly impacts reliability.
1. [Multiple DNS providers, the Perfect Gift this holiday Season](https://blog.signifai.io/multiple-dns-providers-devops-prevent-ddos-attack/)

    Signifai has this nice write-up about setting up redundant DNS providers. My favorite bit is how they polled major domains to see who had added a redundant provider since October 21, and they even shared the source for their polling tool!
1. [The Burden of Running Systems](https://medium.com/production-ready/the-burden-of-running-systems-80de1186f899#.l0igs4be4)

    I’ve featured a lot of articles lately about reducing the amount of code you write. But does that mean that it’s always better to contract with a SaaS provider? This week’s Production Ready delves into the tradeoffs.
### Outages

1. [WoW, Battle.net](http://www.ibtimes.com/are-world-warcraft-battlenet-down-blizzard-says-ddos-attack-affecting-its-games-2456519)
1. [Grubhub, Seamless](http://nypost.com/2016/12/10/seamless-grubhub-suffering-constant-site-crashes/)
1. [DirecTV Now](http://www.theverge.com/2016/12/9/13905004/att-directv-now-internet-tv-streaming-service-outage)
    AT&T’s new online version of DirecTV saw issues in its second week of operation.
1. [Zapier](https://status.zapier.com/incidents/n1v5c4kx1h86)

### [ << Prev ](sreweekly-49.md) ------------- [ Next >> ](sreweekly-51.md)