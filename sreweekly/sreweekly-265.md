## [SRE Weekly Issue #265](https://sreweekly.com/sre-weekly-issue-265/) - April 11, 2021
### Articles

1. [Insights into a Product SRE team at LinkedIn](https://www.linkedin.com/pulse/insights-product-sre-team-linkedin-zaina-afoulki/)

    Here’s a great look into how LinkedIn’s embedded SREs work.Zaina Afoulki and Lakshmi Namboori — LinkedIn
1. [DNS propagation does not exist](https://www.nslookup.io/blog/dns-propagation-does-not-exist/)

    It’s all just other people’s caches.Ruurtjan Pul
1. [Advice for someone moving from SRE to backend engineering](https://shoreline.io/advice-for-someone-moving-from-sre-to-backend-engineering/)

    Charles Cary — Shoreline
1. [The Mightiest Monolith](https://flyingbarron.medium.com/the-mightiest-monolith-7aa67ab8d84f)

    This is the first in a series about lessons SREs can learn from the space shuttle program. The author likens earlier spacecraft to microservices and the Shuttle to a monolith.Robert Barron
1. [The 5 characteristics of high reliability organizations](https://www.ems1.com/safety/articles/the-5-characteristics-of-high-reliability-organizations-sn0WB7bjD4WTN6cq/)

    This article is ostensibly about Emergency Medical Services (EMS), but as is so often the case, it’s directly applicable to SRE. The 5 characteristics are enlightening, and so is the fictitious anecdote about an EMT rattled from a previous incident.Ems1
1. [How we scaled the GitHub API with a sharded, replicated rate limiter in Redis](https://github.blog/2021-04-05-how-we-scaled-github-api-sharded-replicated-rate-limiter-redis/)

    Simple solution meets reality. I like how we get to see what they did when things didn’t quite work out as they were hoping.Robert Mosolgo — GitHub
1. [GitHub Availability Report: March 2021](https://github.blog/2021-04-07-github-availability-report-march-2021/)

    They did the work to convert a database column to a 64-bit integer before it was too late. Unfortunately, one of their library dependencies didn’t use 64-bit integers.Keith Ballinger — GitHub
1. [Learning from incidents: getting Sidekiq ready to serve a billion jobs](https://tech.scribd.com/blog/2020/sidekiq-incident-learnings.html)

    Nakul Pathak — Scribd
### Outages

1. [Let’s Encrypt](https://status.io/pages/incident/55957a99e800baa4470002da/606a943d0188ca052e1abf9a)
1. [Uber](https://www.thesun.co.uk/news/14552047/uber-app-down-users-something-went-wrong/)
1. [Multiple Airlines’ Online Booking Sites](https://simpleflying.com/google-software-airline-booking-sites-downtime/)
    An error in Google’s flight information service caused problems at multiple sites that consume it.
1. [Tinder](https://www.dailymail.co.uk/sciencetech/article-9441371/Tinder-Dating-app-crashes-frustrated-singletons-world.html)
1. [BBC Website](https://sputniknews.com/world/202104081082573857-bbc-website-goes-down-with-fatal-error/)
1. [Facebook, Instagram, and WhatsApp](https://www.rt.com/news/520528-facebook-instagram-down-error/)
1. [Stellar.org (cryptocurrency)](https://status.stellar.org/incidents/7h0czw6mnjzs)
1. [WazirX (cryptocurrency exchange)](https://support.wazirx.com/hc/en-us/articles/900005490846-April-4-2021-Downtime-Analysis-Report)
1. [Microsoft Azure and other services](https://status.azure.com/en-us/status/history/#:~:text=GVY5-TZZ)
    Azure DNS servers experienced an anomalous surge in DNS queries from across the globe targeting a set of domains hosted on Azure.

### [ << Prev ](sreweekly-264.md) ------------- [ Next >> ](sreweekly-266.md)