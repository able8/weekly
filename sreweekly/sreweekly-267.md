## [SRE Weekly Issue #267](https://sreweekly.com/sre-weekly-issue-267/) - April 25, 2021
### Articles

1. [SRE Case Study: Mysterious Traffic Imbalance](https://tech.ebayinc.com/engineering/sre-case-study-mysterious-traffic-imbalance/)

    Yet more proof that DNS behavior varies way more than is obvious at first glance. Who the heck thought longest common prefix matching was a good idea?Charles Li — eBay
1. [Fast and flexible observability with canonical log lines](https://stripe.com/blog/canonical-log-lines)

    The application may log multiple lines during the lifecycle of a request. Stripe has found it invaluable to also log one final line with a fully summary of the request.Brandur Leach — Stripe
1. [Google Incident Report — April 12, 2021](https://static.googleusercontent.com/media/www.google.com/en//appsstatus/ir/4663s0fuvyqu7fg.pdf)

    This is a followup with more detail on the G-Suite outage I reported here last week. A database issue caused two separate outages.Google
1. [The top 3 mistakes companies make with SLOs, SLAs, and SLIs](https://www.getcortexapp.com/post/the-top-3-mistakes-companies-make-with-slos-slas-and-slis)

    Really great advice about 3 common pitfalls in implementing SL*s.Cortex
1. [Going solid: a model of system dynamics and consequences for patient safety – Resilience Roundup](https://resilienceroundup.com/issues/going-solid-a-model-of-system-dynamics-and-consequences-for-patient-safety/)

    This research paper explores the marginal boundary, a set of conditions beyond which a system enters a different operating mode and an accident is much more likely. It discusses the concept of coupling between seemingly unrelated parts of the system and shows how economic incentives can push a system toward this boundary.Dr. Richard Cook and Jens Rasmussen (Original paper)Thai Wood — Resilience Roundup (summary)
1. [Vodafone Idea BGP Leak – Global Routing System Must Implement MANRS](https://blog.catchpoint.com/2021/04/19/vodafone-idea-bgp-leak-global-routing-system-must-implement-manrs/?utm_campaign=BGP_Monitoring&utm_medium=email&_hsenc=p2ANqtz-9wk0LlV1LPmkKDB6ZzTZsKSjXQo6GjN0SB9NiueXZukAzdzELZQ02Tg9X2hzAYp8RjCYXjCJE_mvrfKMoqNIowK-O6wQ&_hsmi=122704848&utm_content=122702980&utm_source=hs_email&hsCtaTracking=50b29aee-04e9-446f-8f27-bfd5feedd7f8%7Ca85fcd1c-ea8d-48a6-9f47-ab39f2f1a86a)

    This is an analysis of a recent BGP leak with a discussion about how the impact from such events can be mitigated through emerging best practices.Alessandro Improta and Luca Sani — Catchpoint
1. [How to Successfully Hand Over Systems](https://developers.soundcloud.com/blog/how-to-successfully-hand-over-systems)

    How do you hand over ownership of a system, transferring enough knowledge that the new owners can maintain its availability and reliability successfully?Aleksandra Gavrilovska — SoundCloud
1. [Resiliency Planning for High-Traffic Events](https://shopify.engineering/resiliency-planning-for-high-traffic-events)

    Shopify works toward Black Friday / Cyber Monday all year long, through a combination of load testing, failure mode analysis, game days, and incident analysis.Ryan McIlmoyl — Shopify
### Outages

1. [Microsoft Azure web portal](https://piunikaweb.com/2021/04/20/microsoft-azure-down-throwing-error-503-outage-is-known-and-under-investigation/)
1. [Microsoft 365](https://www.bleepingcomputer.com/news/security/exchange-online-down-microsoft-365-outage-affects-email-delivery/)
1. [Discord](https://discord.statuspage.io/incidents/gl4j21rcfq3d)
1. [google.com.ar](https://www.newser.com/story/305277/somebody-bought-argentinas-google-domain-for-under-3.html)
    This one’s interesting. A random person was able to buy the domain name google.com.ar, despite the fact that its registration had not expired.

### [ << Prev ](sreweekly-266.md) ------------- [ Next >> ](sreweekly-268.md)