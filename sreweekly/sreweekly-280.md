## [SRE Weekly Issue #280](https://sreweekly.com/sre-weekly-issue-280/) - July 25, 2021
### Articles

1. [The Harmful Consequences of the Robustness Principle ](https://datatracker.ietf.org/doc/draft-iab-protocol-maintenance/)

    The Robustness Principle (“be conservative in what you send, and liberal in what you accept”) has its uses, but it may not be best for the development of mature protocols, according to this IETF draft.Martin Thomson
1. [No, we don’t use Kubernetes](https://ably.com/blog/no-we-dont-use-kubernetes)

    Docker without Kubernetes, does it make sense? These folks have a well-reasoned argument explaining why Kubernetes is not for them.Maik Zumstrull — Ably
1. [Personal data breach reporting for service outages (such as when your CDN is down)](https://decoded.legal/blog/2021/07/personal-data-breach-reporting-for-service-outages-such-as-when-your-cdn-is-down)

    Can a service outage unrelated to security count as a “personal data breach” in terms of GDPR and other regulations? If you agree with this article’s logic, then maybe it can.Neil Brown
1. [When You Do DevSecOps, Don’t Forget the SREs](https://rootly.io/blog/when-you-do-devsecops-don-t-forget-the-sres)

    The interactions between security and reliability incidents can be complex and hard to navigate. The example scenarios in this article really made me think.Quentin Rousseau — Rootly
1. [Solving the Three Stooges Problem](https://www.reddit.com/r/RedditEng/comments/obqtfm/solving_the_three_stooges_problem/)

    To deal with thundering herds, reddit implements caching in front of each of its microservices.Raj Shah — reddit
1. [What’s allowed to count as a cause?](https://surfingcomplexity.blog/2021/07/18/whats-allowed-to-count-as-a-cause/)

    Incident causes are a social construct, and it may be that your organizational structure prevents something from being counted as a cause.Lorin Hochstein
1. [IC1 Reliability Engineer – Dropbox Engineering Career Framework](https://dropbox.github.io/dbx-career-framework/ic1_reliability_engineer.html)

    Check it out, Dropbox publicly released their SRE career ladder.Dropbox
1. [Incidents, Response, and the People With Tim Nicholas](https://www.pageittothelimit.com/incidents-response-and-the-people-tim-nicholas/)

    There’s a moment halfway through this episode of Page It to the Limit where they talk about blamelessness. If you just tell people to “do blameless postmortems”, but you don’t tell them how, then they’ll be afraid to talk about anything people did, and that will hamper learning.Julie Gunderson, with guestTim Nicholas — Page It to the Limit
1. [Migrating Facebook to MySQL 8.0](https://engineering.fb.com/2021/07/22/data-infrastructure/mysql/)

    This was a monumental task, considering the 1000+(!!) internal code patches they had to port from MySQL 5.6 to 8.0.Herman Lee, Pradeep Nayak — Facebook
### Outages

1. [Akamai](https://edgedns.status.akamai.com/incidents/n5zl6dythvfv)
    Akamai had what they’re calling an “Edge DNS Service Incident”. It made headlines this week because it took down many of their customers, similar to the Akamai incident last month.
1. [Let’s Encrypt](https://status.io/pages/incident/55957a99e800baa4470002da/60f5b56d1c82f805369a7d98)
1. [Disney park-related apps](https://insidethemagic.net/2021/07/disney-park-apps-outage-ba1/)
1. [Heroku](https://status.heroku.com/incidents/2303)

### [ << Prev ](sreweekly-279.md) ------------- [ Next >> ](sreweekly-281.md)