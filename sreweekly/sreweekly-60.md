## [SRE Weekly Issue #60](https://sreweekly.com/sre-weekly-issue-60/) - February 20, 2017
### Articles

1. [Reflecting on one very, very strange year at Uber](https://www.susanjfowler.com/blog/2017/2/19/reflecting-on-one-very-strange-year-at-uber)

    Susan Fowler’s articles have been featured here several times previously, and she’s one of my all-time favorite authors. Now it seems that while she was busy writing awesome articles and a book, she was also dealing with a terribly toxic and abhorrent environment of sexual harassment and discrimination at Uber. I can only be incredibly thankful that somehow, despite their apparent best efforts, Uber did not manage to drive Susan out of engineering as happens all to often in this kind of scenario.
1. [Your Team Could Be Just Like Uber, Especially If You’re Certain It’s Not.](https://medium.com/@aickin/your-team-could-be-just-like-uber-especially-if-youre-certain-it-s-not-7d76cf491d0d#.hstgza7dk)

    Even, and perhaps especially if we think we’re doing a good job preventing the kind of abusive environment Susan described, it’s quite possible we’re just not aware of the problems. Likely, even. This kind of situation is unfortunately incredibly common.
1. [GitLab.com / runbooks](https://gitlab.com/gitlab-com/runbooks)

    Wow, what a cool idea! GitLab open-sourced their runbooks. Not only are their runbooks well-structured and great as examples, some of them are general enough to apply to other companies.
1. [Ship Small Diffs](https://blog.skyliner.io/ship-small-diffs-741308bec0d1#.xeevfm77f)

    Full disclosure: Heroku, my employer, is mentioned.
Thanks to Devops Weekly for this one.
1. [Swapping, memory limits, and cgroups](https://jvns.ca/blog/2017/02/17/mystery-swap/)

    TIL: cgroup memory limits can cause a group of processes to use swap even when the system as a whole is not under memory pressure. Thanks again, Julia Evans!
1. [Firefighting is a Team Sport](https://victorops.com/blog/firefighting-team-sport/)

    This week from VictorOps is nifty primer on structuring your team’s on-call and incident response. I love when a new concept catches my eye like this one:This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [How We Improved Our Server Performance](https://dzone.com/articles/how-we-improved-our-server-performance-to-100m-mes)

    Open source IoT platform ThingsBoard’s authors share a detailed account of how they diagnosed and fixed reliability and throughput issues in their software so that it could handle 30k incoming events per second.
1. [Building a Scalable Online Game with Azure – Part 1](http://www.gamasutra.com/blogs/NickPruehs/20170214/291318/Building_a_Scalable_Online_Game_with_Azure__Part_1.php)

    There’s both theory and practice in this article, which opens with an architecture discussion and then continues into the steps to deploy a first verison in a testing Azure environment on your workstation.
1. [Load Balancers: Simplifying High Availability](https://www.digitalocean.com/company/blog/load-balancers-simplifying-high-availability/)

    I don’t often link to new product announcements, but DigitalOcean’s new Load Balancer product caught my attention. It looks to be squarely aimed at improving on Amazon’s ELB product.
1. [Cloud Spanner](https://cloud.google.com/spanner/)

    Okay, apparently I do link to product announcements often.  Google unveiled a new beta product this week for their Cloud Platform: Cloud Spanner. Based on their Spanner paper from 2012, they have some big claims.
### Outages

1. [US National Weather Service](http://www.usnews.com/news/us/articles/2017-02-14/us-weather-service-says-hit-by-first-ever-data-system-outage)
    The U.S. National Weather Service said on Tuesday it suffered its first-ever outage of its data system during Monday’s blizzard in New England, keeping the agency from sending out forecasts and warnings for more than two hours. [Reuters]
1. [The Travis CI Blog: Postmortem for 2017-02-04 container-based Infrastructure issues](http://blog.travis-ci.com/2017-02-13-feb4-container-based-issues-postmortem)
    A garden-variety bug in a newly-deployed version was exacerbated by a failed rollback, in a perfect example of a complex failure with a complex intersection of contributing factors.
1. [Instapaper Outage Cause & Recoveryincorrectly stated](https://medium.com/making-instapaper/instapaper-outage-cause-recovery-3c32a7e9cc5f#.5zhf8gkyd)
    Last week, I incorrectly stated that Instapaper’s database hit a performance cliff. In actuality, their RDS instance was, unbeknownst to them, running on an ext3 filesystem with its single-file limit of 2TB per file. Their only resolution path when they ran out of space was to mysqldump all their data and restore into a new DB running on ext4.
Even if we had executed perfectly, from the moment we diagnosed the issue to the moment we had a fully rebuilt database, the total downtime would have been at least 10 hours.

### [ << Prev ](sreweekly-59.md) ------------- [ Next >> ](sreweekly-61.md)