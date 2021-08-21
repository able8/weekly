## [SRE Weekly Issue #140](https://sreweekly.com/sre-weekly-issue-140/) - September 23, 2018
### Articles

1. [Errata: The Servers Are Burning](https://logicmag.io/05-the-servers-are-burning/)

    My sincerest apologies to Dale Markowitz, the author of this article who I mispronouned in last week’s issue. I’m kicking myself, because I totally didn’t need to use a pronoun at all.Dale Markowitz — LOGIC Magazine
1. [Linux 4.19-rc4 released, an apology, and a maintainership note](https://lore.kernel.org/lkml/CA+55aFy+Hv9O5citAawS+mVZO+ywCKd9NQ2wxUmGsz9ZJzqgJQ@mail.gmail.com/)

    Linus Torvalds made waves this week with an email apologizing for his unprofessional behavior and committing to improving.Linus Torvalds
1. [Designing for Failure to Avoid Disaster](https://launchdarkly.com/blog/designing-for-failure-to-avoid-disaster/)

    A pretty detailed article on how LaunchDarkly designed their system for reliability. The streaming vs. polling section is especially interesting.Adam Zimman — LaunchDarklyFull disclosure: Fastly, my employer, is mentioned.
1. [LogDevice: a distributed data store for logs – Facebook Code](https://code.fb.com/core-data/logdevice-a-distributed-data-store-for-logs/)

    Lots of details about how they achieve their reliability goals. I’d love to see a followup with more detail on why writing a solution in-house made sense versus adopting something like Kafka.Mark Marchukov — Facebook
1. [13 Reasons a Staging Environment Is Failing in Your Organization – DZone DevOps](https://dzone.com/articles/13-reasons-why-staging-environment-is-failing-for-1)

    Harshit Paul — DZone
1. [Mockers – overcoming testing challenges at Grab](https://engineering.grab.com/mockers)

    The challenges in question involve testing a microservice’s interactions with other microservices. Read about their system for distributing and running mock servers for each microservice.Mayank Gupta, K.Vineet Nair, Shivkumar Krishnan, Thuy Nguyen, and Vishal Prakash — Grab
1. [BP is to blame for Deepwater Horizon, but its mistake was actually years of small mistakes.](http://www.slate.com/articles/health_and_science/science/2016/09/bp_is_to_blame_for_deepwater_horizon_but_its_mistake_was_actually_years.html)

    My partner suggested I look into the Deepwater Horizon incident, and I’m glad I did. My two key takeaways were normalization of deviance and this gem:James B. Meigs — Slate
### Outages

1. [Cashplus (bank)](https://www.compelo.com/cashplus-ceo-apology-letter/)
1. [Heroku Incident #1619 (August 30) followup posting](https://status.heroku.com/incidents/1619)
1. [SunTrust (bank)](http://www.heraldtribune.com/news/20180917/suntrust-digital-banking-down-for-second-day)
1. [Ryanair](https://www.paddleyourownkanoo.com/2018/09/17/uh-oh-the-ryanair-crashes-due-to-server-issues-days-before-lockdown-agm-gets-underway/)

### [ << Prev ](sreweekly-139.md) ------------- [ Next >> ](sreweekly-141.md)