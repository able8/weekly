## [SRE Weekly Issue #202](https://sreweekly.com/sre-weekly-issue-202/) - January 12, 2020
### Articles

1. [Conveying confusion without confusing the reader](https://lorinhochstein.wordpress.com/2020/01/11/conveying-confusion-without-confusing-the-reader/)

    When writing about an incident, it’s important to skillfully show the reader how the participants’ understanding of the situation evolved.Lorin Hochstein
1. [The Morning Paper: Ironies of automation](https://blog.acolyer.org/2020/01/08/ironies-of-automation/)

    This is a summary of Bainbridge’s seminal paper, and I really love where Adrian Colyer goes with it.Adrian Colyer — The Morning Paper (summary)Bainbridge — Automatica (original paper)
1. [Sharing SQLite databases across containers is surprisingly brilliant](https://medium.com/@rbranson/sharing-sqlite-databases-across-containers-is-surprisingly-brilliant-bacb8d753054)

    I have to admit, it is brilliant. Why add the risk (and latency) of a centralized configuration repository service when a local DB on each host will do?Rick Branson — Segment
1. [Managing Failure Modes in Microservice Architectures ](https://www.infoq.com/presentations/microservices-failure-modes)

    This one covers a lot. My favorite parts:Adrian Cockcroft
1. [Understanding Observability](https://sdarchitect.blog/2020/01/08/understanding-observability/)

    The author guides you through the moment they began to truly understand what observability is all about. Worth reading even if you’re already quite familiar with the concept.Sanjeev Sharma
1. [Intelligent DNS based load balancing at Dropbox](https://blogs.dropbox.com/tech/2020/01/intelligent-dns-based-load-balancing-at-dropbox/)

    
1. [How We Prevented App Performance Degradation From Sudden Ride Demand Spikes](https://engineering.grab.com/preventing-app-performance-degradation-due-to-sudden-ride-demand-spikes)

    Grab uses bulkheading to prevent localized demand spikes from affecting the service for customers elsewhere. The notable part is that they shed load they can’t satisfy anyway, due to a limited supply of available vehicles.Corey Scott — Grab
### Outages

1. [Dyn](https://www.dynstatus.com/incidents/p7nywmk83gl7)
    Dyn had a delay in DNS resolution in London.
1. [Google Cloud Platform (update on December 18 outage)](https://status.cloud.google.com/incident/cloud-networking/20001#20001001)
    On Wednesday, 18 December, 2019, a part of Google’s production network experienced a temporary reduction in capacity, due to multiple fiber cuts in optical links interconnecting Sofia, Bulgaria with other points-of-presence.
1. [Travelex](https://www.bankingdive.com/news/travelex-hack-halts-currency-exchange-hsbc-barclays/570230/)
1. [Twitter](https://digistatement.com/twitter-down-not-working-in-usa/)
1. [Airbnb](https://mybroadband.co.za/news/internet/334732-airbnb-hit-with-major-outage.html)
1. [Thingiverse](https://digistatement.com/thingiverse-website-down-not-working-users-are-getting-a-502-bad-gateway-error/)
1. [Southwest Airlines website](https://digistatement.com/official-update-southwest-airlines-website-down-not-working-for-many-users/)
1. [Yahoo Mail](https://inews.co.uk/news/uk/yahoo-mail-down-email-outage-reported-by-uk-users-as-others-complain-of-aol-problems-1358414)
1. [Disney+](https://insidethemagic.net/2020/01/disney-plus-down-across-the-world-tm1/)
1. [QuickBooks](https://digistatement.com/official-update-quickbooks-site-down-not-working-users-getting-quickbooks-is-offline-and-were-working-on-it-while-when-attempting-to-log-in/)
1. [Trello](https://www.gamenguides.com/articles/trello-down-and-outage-server-connection-website-and-apps-issues-reported-58971/)
1. [Reddit](https://reddit.statuspage.io/incidents/lqybk5mrnmjk)

### [ << Prev ](sreweekly-201.md) ------------- [ Next >> ](sreweekly-203.md)