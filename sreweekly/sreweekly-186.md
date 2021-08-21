## [SRE Weekly Issue #186](https://sreweekly.com/sre-weekly-issue-186/) - September 22, 2019
### Articles

1. [DBMS Musings: Introduction to Transaction Isolation Levels](http://dbmsmusings.blogspot.com/2019/05/introduction-to-transaction-isolation.html)

    This article is highly technical, while also not being overwhelmingly detailed.Daniel Abadi
1. [How to Avoid the 5 SRE Implementation Traps that Catch Even the Best Teams](https://www.blameless.com/how-to-avoid-the-5-sre-implementation-traps-that-catch-even-the-best-teams/)

    The traps are:Lyon Wong — Blameless
1. [Distributed Tracing: Impact on Engineering Organizations](https://medium.com/@dm03514/distributed-tracing-impact-on-engineering-organizations-d2f775e94aae)

    Need to argue the benefits of implementing distributed tracing in your organization? This article will help you get started.dm03514
1. [Love (and Alerting) in the Time of Cholera (and Observability)](https://charity.wtf/2019/09/20/love-and-alerting-in-the-time-of-cholera-and-observability/)

    Charity Majors
1. [Enhancing Bandaid load balancing at Dropbox by leveraging real-time backend server load information](https://blogs.dropbox.com/tech/2019/09/enhancing-bandaid-load-balancing-at-dropbox-by-leveraging-real-time-backend-server-load-information/)

    Round-robin load balancing often isn’t good enough; it’s necessary to intelligently route requests to nodes that aren’t overloaded. How do you get information about backend health to distributed load balancer nodes efficiently? A: add a response header.Haowei Yuan, Yi-Shu Tai, and Dmitry Kopytkov — Dropbox
1. [Splash the cache: how caching improved our reliability](https://making.pusher.com/splash-the-cache-how-caching-improved-our-reliability/)

    By adding in-memory caching with a mere 3-second TTL, these folks achieved a ~75% cache hit rate, allowing them to withstand request spikes without an outage.MINA GYIMAH — Pusher
### Outages

1. [Tokbox](https://status.tokbox.com/incidents/kp1v7g5yqmc2)
    Thanks to Aos Dabbagh for this one.
1. [Chef (system administration tool)removed their codefollowup](https://status.chef.io/incidents/lj7h89v2v7r6)
    Many of us experienced failures in our Chef runs after their former employee removed their code. Chef posted a followup explaining their position on the matter.
1. [Fastly](https://status.fastly.com/incidents/kjrxdc6xs31s)
1. [Reddit](https://reddit.statuspage.io/incidents/4q8mfp969hky?u=hxjfj9djsk48)
1. [Net4 (hosting provider)](https://www.outlookindia.com/newsscroll/indias-domain-hosting-provider-net4com-suffers-major-outage/1621007)
1. [Salesforce](https://marketingland.com/performance-degradation-affecting-salesforce-clients-267699)
1. [LinkedIn](https://piunikaweb.com/2019/09/19/linkedin-down-and-not-loading-content-reportedly/)
1. [Google Search](https://www.rt.com/news/468963-google-search-engine-down-reports/)
1. [Heroku](http://feedproxy.google.com/~r/HerokuStatus/~3/EocpT5hc_x4/1901)
1. [Squarespacethis one](https://status.squarespace.com/incidents/pjt4wpyrw9ws)
    Also this one.

### [ << Prev ](sreweekly-185.md) ------------- [ Next >> ](sreweekly-187.md)