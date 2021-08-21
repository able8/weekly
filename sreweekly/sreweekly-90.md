## [SRE Weekly Issue #90](https://sreweekly.com/sre-weekly-issue-90/) - September 24, 2017
### Articles

1. [On the design of distributed programming models](https://blog.acolyer.org/2017/08/17/on-the-design-of-distributed-programming-models/)

    We’re all becoming distributed systems engineers, and this stuff sure isn’t easy.
1. [Continuous self-testing at Hosted Graphite — why we send external canaries, every second](https://blog.hostedgraphite.com/2017/07/06/continuous-self-testing-at-hosted-graphite-why-we-send-external-canaries-every-second/)

    Every-second canarying is a pretty awesome concept. Not only that, but they even post the results on their status page. Impressive!
1. [Mea culpa: Lessons learned from firefighting](https://theconversation.com/mea-culpa-lessons-learned-from-firefighting-79994)

    So many lessons! My favorite is to make sure you test the “sad path”, as opposed to just the “happy path”. If a customer screws up their input and then continues on correctly from there on, does everything still work?
1. [SRECon EMEA 17 Notes](https://docs.google.com/document/d/1CTn0GO7dNnl06Fgb_lbKjQ6txYQbOSJX2Ud9RaaGYkI/edit#heading=h.52gt707t262o)

    Extensive notes taken during 19 talks at SRECon 17 EMEA. I’m blown away by the level of detail. Thanks, Aaron!
1. [CPU Performance Checklist](https://www.linkedin.com/pulse/cpu-performance-checklist-sredevopsperformance-ramnath-jk/)

    A cheat sheet and tool list for diagnosing CPU-related issues. There’s also one on network troubleshooting by the same author. Note: LinkedIn login required to view.
1. [Antifragility 101 – Production Ready](https://medium.com/production-ready/antifragility-101-13f52f385d67)

    Antifragility is an interesting concept that I was previously unaware of. I’m not really sure how to apply it practically in an infrastructure design, but I’m going to keep my eye out for antifragile patterns.
1. [Eight reasons why you should conduct a DNS audit](https://www.computerworld.com.au/article/627665/eight-reasons-why-should-conduct-dns-audit/)

    It’s easy to overlook your DNS, but a failure can take your otherwise perfectly running infrastructure down — at least from the perspective of your customers.
1. [An Often Overlooked Tool In Workplace Safety Prevention: The Near-Miss Investigation](http://www.jdsupra.com/legalnews/an-often-overlooked-tool-in-workplace-85725/)

    Do you run a retrospective on near misses? The screws they tightened in this story could just as easily be databases quietly running at max capacity.
1. [Introducing Azure Availability Zones for resiliency and high availability](https://azure.microsoft.com/en-us/blog/introducing-azure-availability-zones-for-resiliency-and-high-availability/)

    Oh look, Azure has AZs now.
1. [What’s in a transport layer?](https://www.oreilly.com/ideas/whats-in-a-transport-layer?cmp=tw-webops-na-article-vleu17_transport_layer_vleu_kb)

    The transport layer in question is gRPC, and this article discusses using it to connect a microservice-based infrastructure. If you’ve been looking for an intro to gRPC, check this out.
1. [Beyond Lights-Out: Future Data Centers Will Be Human-Free](http://www.datacenterknowledge.com/design/beyond-lights-out-future-data-centers-will-be-human-free)

    How do you prevent human error? Remove the humans. Yeah, I’m not sure I believe it either, but this was still an interesting read just to learn about the current state of lights-out datacenters.
1. [Documenting your architecture: Wireshark, PlantUML and a REPL to glue them all.](http://danlebrero.com/2017/04/06/documenting-your-architecture-wireshark-plantuml-and-a-repl/)

    This is a really neat idea: generate an interaction diagram automatically using a packet capture and a UML tool.Thanks to Devops Weekly for this one.
### Outages

1. [.iolast fall](http://nic.io)
    The .io TLD went down again, in exactly the same way as last fall.
1. [PagerDutythe hangops slack](https://status.pagerduty.com/incidents/1jj1tzw836rp)
    PagerDuty suffered a major outage lasting over 12 hours this past thursday. Customers scrambled to come up with other alerting methods.
Some really excellent discussion around this incident happened on the hangops slack in the #incident_response channel. I and others requested more details on the actual paging latency and PagerDuty delivered them on their status site. Way to go, folks!
1. [StatusPage.io](http://metastatuspage.com/incidents/gpnmrq14wgnh)
    I noticed this minor incident after getting a 500 reloading PagerDuty’s status page.
1. [The Travis CI Blog: Sept 6 – 11 macOS outage postmortem](http://blog.travis-ci.com/2017-09-20-sept-6-11-macos-outage-postmortem)
    This week, Travis posted this followup describing the SAN performance issues that impacted their system.
1. [Outlook and Hotmail](http://www.manchestereveningnews.co.uk/news/uk-news/microsoft-outlook-hotmail-down-problems-13639969)

### [ << Prev ](sreweekly-89.md) ------------- [ Next >> ](sreweekly-91.md)