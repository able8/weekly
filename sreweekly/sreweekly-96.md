## [SRE Weekly Issue #96](https://sreweekly.com/sre-weekly-issue-96/) - November 5, 2017
### Articles

1. [The Phone Book Is On Fire: Lessons From the Dyn DNS DDoS — Velocity NYC 2017](https://youtu.be/Y1frKCKL8xA)

    Here’s the recording of my Velocity 2017 talk, posted on YouTube with permission from O’Reilly (thanks!). Want to learn about some gnarly DNS details?
1. [Log20: Fully automated optimal placement of log printing statements under specified overhead threshold](https://blog.acolyer.org/2017/11/03/log20-fully-automated-optimal-placement-of-log-printing-statements-under-specified-overhead-threshold/)

    I fell in love with this after reading just the title, and it only got better from there. Why add debug statements haphazardly when an algorithm can automatically figure out where they’ll be most effective? I especially love the analysis of commit histories to build stats on when debug statements were added to various open source projects.
1. [Operating a Kubernetes network](https://jvns.ca/blog/2017/10/10/operating-a-kubernetes-network/)

    Julia Evans is back with another article about Kubernetes. Along with explaining how it all fits together, she describes a few things that can go wrong and how to fix them.
1. [How can we apply the principles of chaos engineering to AWS Lambda?](https://hackernoon.com/how-can-we-apply-the-principles-of-chaos-engineering-to-aws-lambda-80f87e3237e2?__s=gcxkayouhzyr45m1hboa)

    In this introductory post of a four-part series, we learn why chaos testing a lambda-based infrastructure is especially challenging.
1. [Google Vizier: A service for black-box optimization](https://blog.acolyer.org/2017/10/02/google-vizier-a-service-for-black-box-optimization/)

    I love the idea of a service that automatically optimizes things even without knowing anything about their internals. Mmm, cookies.
1. [Lyft’s Envoy dashboards – mattklein123 – Medium](https://medium.com/@mattklein123/lyfts-envoy-dashboards-5c91738816b1?__s=bwykwk1kcceogszq8abt)

    
1. [Microsoft has built a secret network emulator it says can prevent most cloud outages](https://siliconangle.com/blog/2017/11/01/microsoft-built-secret-network-emulator-says-can-prevent-cloud-outages/)

    It’s not a secret since they published a paper about it. This is an intriguing idea, but I’m wondering whether it’s really more effective than staging environments tend to be in practice.
1. [The Rise of Site Reliability Engineers](https://blog.newrelic.com/2017/10/30/site-reliability-engineer-sre/)

    A history of the SRE profession and a description of how New Relic does SRE.Full disclosure: Heroku, my employer, is mentioned.
### Outages

1. [Collision with buffer stops at King’s Cross station, London, 15 August 2017
](https://www.gov.uk/government/publications/safety-digest-152017-kings-cross/collision-with-buffer-stops-at-kings-cross-station-london-15-august-2017)
    This is the Rail Accident Investigation Branch’s report on a minor accident involving a driver that suffered a “microsleep” due to fatigue.
1. [LearnVest](https://www.financial-planning.com/news/learnvest-went-dark-leaving-users-in-the-lurch)
1. [Slack](https://techcrunch.com/2017/10/31/slack-is-still-down-and-its-past-5-oclock-so-go-home/?ncid=mobilenavtrend)

### [ << Prev ](sreweekly-95.md) ------------- [ Next >> ](sreweekly-97.md)