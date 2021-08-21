## [SRE Weekly Issue #261](https://sreweekly.com/sre-weekly-issue-261/) - March 14, 2021
### Articles

1. [What Do Fighter Pilots and Incident Management Have in Common?](https://www.transposit.com/blog/fighter-pilots-and-incident-management/)

    I find it really refreshing that fighter pilots have a retrospective about every single mission, successful or not. There’s always something to learn.Jessica Abelson — Transposit
1. [Incident Response at Heroku](https://blog.heroku.com/incident-response-at-heroku-2020)

    Heroku applies the Incident Management System, designating an Incident Commander who keeps the incident on track and oversees communications, both external and internal.Guillaume Winter — Heroku
1. [How Khan Academy Successfully Handled 2.5x Traffic in a Week](https://blog.khanacademy.org/how-khan-academy-successfully-handled-2-5x-traffic-in-a-week/)

    This story is becoming common: Khan had a sudden influx of traffic when pandemic lockdowns began. Their strategy involved the use of the cloud and a CDN.Marta Kosarchyn — Khan AcademyFull disclosure: Fastly, my employer, is mentioned.
1. [Under the Hood: Ensuring Site Reliability](https://engineering.squarespace.com/blog/2017/under-the-hood-ensuring-site-reliability)

    Here’s a great summary of how Squarespace does SRE.Franklin Angulo — Squarespace
1. [[Increment: Reliability] Reliability at scale](https://increment.com/reliability/reliability-at-scale/)

    The leaders each answer a series of questions about how their organization handles reliability, giving an interesting compare-and-contrast overview.IncrementFull disclosure: Fastly is my employer.
1. [[Increment: Reliability] Case study: Resilience as adaptability at Freshworks](https://increment.com/reliability/resilience-as-adaptability-freshworks/)

    Using a disaster plan created after a devastating hurricane, Freshworks survived and thrived during the pandemic, delivering a major new product by its pre-pandemic deadline.Ipsita Agarwal — Increment
1. [What Is a Canary Deployment?](https://launchdarkly.com/blog/what-is-a-canary-deployment/)

    This one explains what a canary deployment is, how it can help you, and how canary deployments differ from blue/green deployments.LaunchDarkly
1. [How to Build an SRE Team with a Growth Mindset](https://www.blameless.com/blog/how-to-build-an-sre-team-with-a-growth-mindset)

    This article explains the meaning of a growth mindset and shows how it applies to SRE.Emily Arnott — Blameless
### Outages

1. [Fastly](https://status.fastly.com/incidents/pk90fj9sj1pr)
    Full disclosure: Fastly is my employer.
1. [OVH CloudOVH Cloudaccording to its creators](https://twitter.com/olesovhcom/status/1369478732247932929)
    This week, there was a major fire at an OVH Cloud datacenter. As a result, Rust (an MMOG) permanently lost data, according to its creators.
1. [All domains containing “t.co” in Russia](https://twitter.com/DougMadory/status/1369665894494900224?s=19)
    It appears that Russia tried to impair access to Twitter’s URL-shortening domain t.co, but their pattern-matching was overzealous and affected any domain that contained “t.co” (think reddit.com, microsoft.com, and many others).
1. [DynHeroku](https://www.dynstatus.com/incidents/x8psymmjc5bw)
    Dyn had a DNS outage. I noted impact to Heroku, but I didn’t see any other related outage postings.
1. [Chef](https://status.chef.io/incidents/k1919203q8ky)
1. [GitHub](https://www.githubstatus.com/incidents/wdhkj70kynwk)

### [ << Prev ](sreweekly-260.md) ------------- [ Next >> ](sreweekly-262.md)