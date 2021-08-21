## [SRE Weekly Issue #59](https://sreweekly.com/sre-weekly-issue-59/) - February 12, 2017
### Articles

1. [Significant Threats to Patient Safety from Healthcare Provider Overload and Burnout Should Not Be Overlooked](http://www.prweb.com/releases/2017/02/prweb14061034.htm)

    Here’s a great article about burnout in the healthcare sector. There’s mention of second victims (see also Sydney Dekker) and a vicious circle: burnout leads to mistakes, which lead to adverse patient outcomes, which lead to guilt and frustration, which leads to burnout.
1. [What Does Downtime Cost Your Business?](https://www.pagerduty.com/blog/cost-downtime/)

    Every week, I find and ignore at least one bland article about the “huge cost of downtime”. They almost never have anything interesting or new to say. This article by PagerDuty takes a different approach that I find refreshing, starting off by defining “downtime” itself.
1. [Honest status reporting and AWS service status “truth” in a post-truth world](https://blog.ably.io/honest-status-reporting-and-aws-service-status-truth-in-a-post-truth-world-8b9a31c8cc90#.vmxj5nuep)

    A frustrated CEO speaks out against AWS’s infamously sanguine approach to posting on their status site.
1. [Postmortem of database outage of January 31 | GitLab](https://about.gitlab.com/2017/02/10/postmortem-of-database-outage-of-january-31/)

    As mentioned last week, here’s the final, published version of GitLab’s postmortem for their incident at the end of last month.
1. [Jepsen: MongoDB 3.4.0-rc3](https://jepsen.io/analyses/mongodb-3-4-0-rc3)

    MongoDB contracted Jepsen to test their new replication protocol. Jepsen found some issues, which are fixed, and now MongoDB gets a clean bill of health. Pretty impressive! Even cooler is that the Mongo folks have integrated Jepsen’s tests into their CI.
### Outages

1. [Instapaper](http://blog.instapaper.com/post/157027537441)
    Instapaper hit a performance cliff with their database, and the only path forward was to dump all data and load it into a new, more powerful DB instance.
1. [Google Cloud Status Dashboard](https://status.cloud.google.com/incident/compute/17003#5660850647990272)
    Google released a postmortem for a network outage at the end of January.
1. [OWASA (Orange County, FL, USA water authority)](http://www.indyweek.com/news/archives/2017/02/10/human-error-caused-owasa-fluoride-overdose-owasa-very-sorry-about-that)
    OWASA had to cut off the municipal water supply for 3 days after an accidental overfeed of fluoride into the drinking supply. They engaged in an impressive post-analysis and released a detailed root cause analysis document. It was a pretty interesting read, and I highly recommend clicking through to the PDF and reading it.  There you’ll see that “human error” was a proximal but by no means root cause of the outage, especially since the human in question corrected their error after just 12 seconds.

### [ << Prev ](sreweekly-58.md) ------------- [ Next >> ](sreweekly-60.md)