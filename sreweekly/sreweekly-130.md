## [SRE Weekly Issue #130](https://sreweekly.com/sre-weekly-issue-130/) - July 15, 2018
### Articles

1. [Goodbye Microservices: From 100s of problem children to 1 superstar](https://segment.com/blog/goodbye-microservices/)

    Segment discovered the hard way that their move to a microservice architecture had brought far more problems than benefits. Here’s why they transitioned back and how they pulled it off. Awesome article!Alexandra Noonan — Segment
1. [Establishing Resilience: The Challenges and Opportunities of Complexity](https://blog.newrelic.com/2018/06/18/reliability-and-resilience/)

    Drawing on the work of Dr. David Woods and the rest of the SNAFU Catchers, this article discusses the concepts behind resiliency and how to measure and achieve it.Beth Long — New Relic
1. [Solving for serverless: How do you manage something that’s not there?](https://www.networkworld.com/article/3287648/lan-wan/solving-for-serverless-how-do-you-manage-something-that-s-not-there.html)

    Serverless is not the magical gateway to the land of NoOps. You still have to operate your system even if you’re not directly running the servers. This article does a great job of explaining why. Bhanu Singh — Network World
1. [How I use Wireshark](https://jvns.ca/blog/2018/06/19/what-i-use-wireshark-for/)

    New to me: Wireshark’s statistics view and how it can be useful.Julia Evans
1. [Health and availability in computer systems](https://medium.com/observability/health-and-availability-b2f32ce28716)

    How do you define whether your system is available and healthy? This article uses an anology to medical health.José Carlos Chávez — Typeform
1. [On the AWS Application Load Balancer HTTP/2 Support](https://medium.com/@ptforsberg/on-the-aws-application-load-balancer-http-2-support-fad4bc67b21a)

    These folks are experiencing mysterious latency with HTTP/2 traffic to their ALB and published this report on their investigation. There’s no happy ending here — ultimately they disabled HTTP/2 support. I hope they update if they do discover the culprit.Peter Forsberg — ShopGun
1. [relp 100% cpu – rsyslog stop after start · Issue #13 · rsyslog/librelp · GitHub](https://github.com/rsyslog/librelp/issues/13)

    I had some fun this week unearthing the cause for the chronic lockups in Rsyslog that we’ve experienced at work. I found the cause (overeager retries of socket writes) and put together a bug report and a pull request.Full disclosure: Fastly, my employer, is mentioned.
1. [Building Grab’s Experimentation Platform](https://engineering.grab.com/building-grab-s-experimentation-platform)

    I love science! Grab wrote a nifty tool to help them select cohorts of users and perform experiments on them.Abeesh Thomas and Roman Atachiants — Grab

1. [Auto Scaling Production Services on Titus – Netflix TechBlog – Medium](https://medium.com/netflix-techblog/auto-scaling-production-services-on-titus-1f3cd49f5cd7?source=rss----2615bd06b42e---4)

    Titus is the container orchestration system that Netflix created and open sourced. Rather than building a new auto-scaling feature for Titus, they instead got Amazon to productize EC2’s auto-scaling engine as a generalized auto-scaling tool, which Netflix then integrated with Titus. Neat!See Amazon’s Application Auto Scaling announcement, published this past week.Andrew Leung, Amit Joshi, and the rest of the Titus team — Netflix
### Outages

1. [Gmail](https://www.inverse.com/article/46822-gmail-outage-map-shows-massive-blackouts-affecting-new-york-la-and-sf)
1. [Google Docs, Sheets, et al.](https://www.google.com/appsstatus#hl=en&v=issue&sid=5&iid=df974d012febcd7ad8ef758e5ae79327)
1. [YouTube TV](https://twitter.com/YouTubeTV/status/1017121393857556485)
    During the World Cup match.
1. [Discord](https://discord.statuspage.io/incidents/0ddh383sfb1r)
    Discord had a couple of outages this week.
1. [Instagram](https://mashable.com/2018/07/13/instagram-down-worldwide/)
1. [Mastercard](https://www.pymnts.com/mastercard/2018/outage-declined-credit-card-payments-europe/)
1. [Facebook Messenger](http://www.dailymail.co.uk/sciencetech/article-5949911/Facebook-Messenger-Popular-messaging-service-goes-offline-users-worldwide.html)
1. [Snapchat](https://beebom.com/massive-snapchat-outage-hits-uk-and-europe-users-also-complain-about-lost-snapstreaks/)
1. [99acres (real estate site)](https://www.medianama.com/2018/07/223-99-acres-down/)
1. [Heroku](http://feedproxy.google.com/~r/HerokuStatus/~3/L3pKUuqvZd8/1583)
1. [Disney blames 4-hour tech woes on network maintenance](http://www.orlandosentinel.com/business/tourism/os-bz-disney-technology-follow-20180709-story.html)
    Here’s an update on the Disney system outage I linked to last week.
Gabrielle Russon — Orlando Sentinel

### [ << Prev ](sreweekly-129.md) ------------- [ Next >> ](sreweekly-131.md)