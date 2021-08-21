## [SRE Weekly Issue #83](https://sreweekly.com/sre-weekly-issue-83/) - July 30, 2017
### Articles

1. [Do You Suffer From Decision Fatigue?](https://mobile.nytimes.com/2011/08/21/magazine/do-you-suffer-from-decision-fatigue.html?referer=)

    Decision fatigue is the diminishment of certain mental faculties after making many decisions. It can cause incidents, and just as importantly, it can make incident response more difficult. After reading this article, I’m wondering if I should be asking incident responders to stop and drink a glass of orange juice before making a tough call during an incident.
1. [The mystery of the hanging S3 downloads](https://www.snellman.net/blog/archive/2017-07-20-s3-mystery/?__s=gcxkayouhzyr45m1hboa)

    Here’s an interesting debugging session that plumbs some of the more obscure depths of TCP.
1. [Serverless computing has landed: How IT Ops can adapt](https://techbeacon.com/serverless-computing-has-landed-how-it-ops-can-adapt)

    What does DR look like if your system is serverless? How do you manage performance if you don’t control the thing that loads (and hopefully pre-caches) your code?
1. [Incident Management for Operations (book)](http://shop.oreilly.com/product/0636920036159.do)

    The new book on incident response from the folks at Blackrock3 has arrived! They draw on their years of fire incident response experience to teach us how to resolve outages. I had the privilege of attending one of Blackrock3’s 2-day training sessions last week and I highly recommend it.
1. [Support Driven Development: Listen now so you don’t hear it later](http://blog.scalyr.com/2017/07/support-driven-development/)

    I like the idea of focusing on reducing customer pain points, even if they’re not directly due to bugs. After all, reliability is all about the customer experience.
1. [ChAP: Chaos Automation Platform](https://medium.com/netflix-techblog/chap-chaos-automation-platform-53e6d528371f?source=rss----2615bd06b42e---4)

    Netflix’s ChAP tests a target microservice by creating experimental and control clusters and routing a small portion of traffic to them.
1. [Starting the Avalanche](https://medium.com/netflix-techblog/starting-the-avalanche-640e69b14a06?source=rss----2615bd06b42e---4)

    Microservice-based architecture is great, right? The problem is that the fan-out of backend requests can create an amplification vector for a DDoS attack. A small, carefully-constructed API call from an attacker can result in a massive number of requests to services in the backend, taking them down.
1. [Learning From Failure and Success – Production Ready](https://medium.com/production-ready/learning-from-failure-and-success-955a2ef1405)

    The latest from Mathias Lafeldt is this article about post-hoc learning. He draws on Zwiebeck and Cook, reminding us that both success and failure are normal circumstances in complex systems.
1. [GitHub – dastergon/awesome-chaos-engineering: A curated list of awesome Chaos Engineering resources](https://github.com/dastergon/awesome-chaos-engineering)

    Remember Awesome SRE? The same author, Pavlos Ratis, has pulled together a ton of links on Chaos Engineering.  Thanks, Pavlos!
1. [GitHub – dastergon/postmortem-templates: A collection of postmortem templates](https://github.com/dastergon/postmortem-templates)

    He’s also compiled this set of postmortem templates, drawn from various sources.  He’s unstoppable!
1. [Pingdom’s Live Map Shows You The State Of The Internet As It Happens](http://royal.pingdom.com/2017/07/25/live-map-the-state-of-the-internet/)

    What a great idea, and I wish I’d known about it earlier! Pingdom uses their aggregate monitoring data to create a live map of the internet. Might be useful for those big events like the Dyn DDoS or the S3 outage.
1. [Verizon points finger at Niantic for problems at Pokémon Go event](http://www.businessinsider.com/verizon-points-finger-niantic-problems-pokemon-go-event-2017-7)

    Last week, I reported on the disaster that was Niantic’s Pokémon Go live event. Verizon wants to assure us that it wasn’t a capacity issue on their part.
### Outages

1. [EC2 (us-east-1)HerokuRollbar](https://status.aws.amazon.com)
    Between 6:47 AM and 7:10 AM PDT we experienced increased launch failures for EC2 Instances, degraded EBS volume performance and connectivity issues for some instances in a single Availability Zone in the US-EAST-1 Region. The issue has been resolved and the service is operating normally.
This one seems to have affected several companies including Heroku and Rollbar.
1. [Marketo](https://www.itnews.com.au/news/marketo-suffers-major-outage-after-domain-renewal-fail-469460)
    Marketo failed to renew their domain name registration, reportedly due to a failure in their automated tooling.
1. [Instagram](http://www.dailymail.co.uk/sciencetech/article-4733034/Instagram-crashes-European-users.html)
1. [Report on July 7, 2017 incident | Gandi News](https://news.gandi.net/en/2017/07/report-on-july-7-2017-incident/)
    Here’s one I missed from earlier this month.
In all, 751 domains were affected by this incident, which involved an unauthorized modification of the name servers [NS] assigned to the affected domains that then forwarded traffic to a malicious site exploiting security flaws in several browsers.
Thanks to an anonymous reader for this one.
1. [Threat Stack Status – Config Audit Database Maintenance](https://threatstack.statuspage.io/incidents/0fyphzz1gh4p)
    Another one I missed. This one appears to be a maintenance that went wrong.Thanks to an anonymous reader for this one.

### [ << Prev ](sreweekly-82.md) ------------- [ Next >> ](sreweekly-84.md)