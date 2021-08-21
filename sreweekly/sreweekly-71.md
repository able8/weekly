## [SRE Weekly Issue #71](https://sreweekly.com/sre-weekly-issue-71/) - May 7, 2017
### Articles

1. [How we Upgraded a 22TB MySQL Cluster from 5.6 to 5.7 (in 9 months)](https://thoughts.t37.net/how-we-upgraded-a-22tb-mysql-cluster-from-5-6-to-5-7-in-9-months-b09211af57ba)

    The interesting bit in this story is that upgrading to 5.7 requires a full table rewrite (<tt>ALTER TABLE</tt>) for any table that has time-related columns. Their initial test-run took months and still hadn’t finished.
1. [Migrating from Heroku to AWS: Our Story](https://medium.com/adstage-engineering/migrating-from-heroku-to-aws-our-story-80084d31025e?__s=gcxkayouhzyr45m1hboa)

    AdStage made the move from Heroku to running their service directly on EC2, and in this article they explain why and how.Language Warning: contains the word “sexy” used to describe new or interesting technology.Full disclosure: Heroku, my employer, is mentioned.
1. [The Discipline of Chaos Engineering](https://blog.gremlininc.com/the-discipline-of-chaos-engineering-e39d2383c459)

    I’ve featured many articles from Mathias Lafeldt as part of his series, Production Ready. Now that he’s moved to Gremlin Inc (a SaaS helping customers run chaos experiments), Mathias reintroduces the history and theory of Chaos Engineering.
1. [Homegrown master-master replication for a NoSQL database](http://feedproxy.google.com/~r/HighScalability/~3/XrwZuZ560tk/homegrown-master-master-replication-for-a-nosql-database.html)

    The folks behind Mail.ru implemented their own master-master replication system on top of Tarantool, a DBMS I’d never heard of. Their implementation is based on some details of their use-case that may not apply more broadly, but the design discussion is interesting nonetheless.
1. [OnlineSchemaChange rebuilt in Python](https://code.facebook.com/posts/1290069194423954/onlineschemachange-rebuilt-in-python/)

    Facebook rewrote their tool, OnlineSchemaChange in Python (from the original PHP). OSC is a tool for doing DDL in MySQL without downtime.
1. [After the Disaster: How to Learn from Historical Incident Management Data](https://www.pagerduty.com/blog/learn-from-incident-management-data/)

    From PagerDuty, an article on the incident management data to gather, how to gather it, and how to analyze it.
1. [What is Structured Logging?](https://dzone.com/articles/what-is-structured-logging)

    A basic introduction to structured logging, including rationale on why you’d want to use it. With infrastructures growing more and more complicated, I find structured logging indispensable in keeping everything up and running and debugging difficult problems.
1. [Building Express Backbone: Facebook’s new long-haul network](https://code.facebook.com/posts/1782709872057497/building-express-backbone-facebook-s-new-long-haul-network/)

    For the network nerds, Facebook details their new inter-datacenter network topology.
1. [Introducing Machine Learning for the Elastic Stack](https://www.elastic.co/blog/introducing-machine-learning-for-the-elastic-stack)

    New in the latest version of Elastic Stack (think ElasticSearch, Logstash, Kibana, etc) is built-in anomaly detection using machine learning, based on technology from Prelert (acquired by Elastic in 2016). “Machine Learning” — they might as well say it’s powered by “Lasers™”. If you try this out and have any success, please write up your results and send me a link!
### Outages

1. [WhatsApp](https://motherboard.vice.com/en_us/article/whatsapp-is-down)
1. [Churchill Downs](http://www.wlky.com/article/network-outage-delays-racing-at-churchill-downs/9607493)
1. [Mumsnet](http://www.mirror.co.uk/tech/mumsnet-down-popular-parenting-website-10354233)
1. [Cloudflare Status – Network Performance Issues in multiple locations](https://www.cloudflarestatus.com/incidents/51q3xhq8w7t8)
1. [TeliaCloudFlarePingdomDiscordthis isn’t the first time](https://www.theregister.co.uk/2017/05/02/telia_hiccups_cloudflare_falls_over/)
    Telia, a major backbone internet provider, deployed a misconfiguration that caused routing issues across the globe. CloudFlare noticed, as did Pingdom and Discord. Think back to almost a year ago, and you may remember that this isn’t the first time that they’ve caused this kind of far-reaching problem.

### [ << Prev ](sreweekly-70.md) ------------- [ Next >> ](sreweekly-72.md)