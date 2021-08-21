## [SRE Weekly Issue #106](https://sreweekly.com/sre-weekly-issue-106/) - January 21, 2018
### Articles

1. [The Limitations of Chaos Engineering – Production Ready](https://medium.com/production-ready/the-limitations-of-chaos-engineering-2a74816c0df3)

    Chaos engineering is extremely useful, and Mathias Lafeldt has written plenty about its virtues. But as with everything, it’s important to be aware of its pitfalls and shortcomings too.
1. [What Went Wrong In Hawaii, Human Error? Nope, Bad Design.](https://www.fastcodesign.com/90157153/don-norman-what-went-wrong-in-hawaii-human-error-nope-bad-design)

    There’s been a lot of talk of firing (or worse) the person whose actions led to the false alarm in Hawaii. That’s why I’m especially glad to see this excellent analysis by Don Norman (The Design of Everyday Things and others). Bonus content: another article along the same vein with some more interesting tidbits.
1. [In defence of swap: common misconceptions](https://chrisdown.name/2018/01/02/in-defence-of-swap.html)

    Think twice before you disable swap, says Chris Down, an author of the upcoming cgroup v2 in the Linux kernel.
1. [SRE Survey 2018](http://pages.catchpoint.com/SRE-Survey-2018.html)

    Catchpoint is running a survey of SREs and SRE-like folks, and I’d really appreciate it if you’d take a moment to fill it out. Not only will the resulting data be very interesting, but Catchpoint is donating $5 to charity for every survey completed. Let’s stuff that ballot box and get them to hit their cap of $3000!
1. [Building a Distributed Log from Scratch, Part 4: Trade-Offs and Lessons Learned](https://bravenewgeek.com/building-a-distributed-log-from-scratch-part-4-trade-offs-and-lessons-learned/)

    The awesome continues this week with a discussion of the importance of simplicity in the design of a reliable system.
1. [What Makes a Failure a Disaster?](http://blog.launchdarkly.com/what-makes-a-failure-a-disaster/)

    This article from Heidi Waterhouse at Launch Darkly starts off with a really interesting take on the Y2K bug and continues on to discuss risk management in operations.
1. [When letting the user put the system into an invalid state is a desirable property](https://ayende.com/blog/180993/when-letting-the-user-put-the-system-into-an-invalid-state-is-a-desirable-property)

    This short article has an extremely cogent point: design your system to be flexible enough to allow the user to do something seemingly incorrect, because they might need to while responding to an incident!
1. [Project STAR*: Streamlining Our On-Call Process](https://engineering.linkedin.com/blog/2018/01/project-star-streamlining-our-on-call-process)

    LinkedIn had a problem: their on-call system was so dysfunctional that they had to scramble to find coverage for an engineer that had been scheduled to be on call when they were on vacation. They explain how they identified the problem, came up with a solution, and implemented it, including automation and cultural fixes.
1. [Monitoring in a DevOps World](http://queue.acm.org/detail.cfm?id=3178371)

    If the phrase “a DevOps World” makes you feel ill, don’t dismiss this article from ACM Queue out of hand. It’s got some great points about designing effective monitoring, and I like the introduction of the “Real Systems Monitoring” concept (akin to “Real User Monitoring” or RUM).
### Outages

1. [Heroku](https://status.heroku.com/incidents/1372)
    Heroku had a 29-hour impairment to their application log routing platform.

### [ << Prev ](sreweekly-105.md) ------------- [ Next >> ](sreweekly-107.md)