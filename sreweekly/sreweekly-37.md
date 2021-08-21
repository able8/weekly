## [SRE Weekly Issue #37](https://sreweekly.com/sre-weekly-issue-37/) - August 28, 2016
### Articles

1. [The “network partitions are rare” fallacy](http://kellabyte.com/2013/11/04/the-network-partitions-are-rare-fallacy/)

    Sometimes I follow chains of references from article to article until I find a new author to follow, and this time it’s Kelly Sommers.  In this gem, she debunks the rarity of network partitions by recasting them as availability partitions.  If half of your nodes aren’t responding because their CPUs are pegged, you still have a network partition.

1. [How DIGIT Created High Availability on the Public Cloud to Keep Its Games Running](https://www.linux.com/news/how-digit-created-high-availability-public-cloud-keep-its-games-running)

    Two engineers from MMO company DIGIT gave this short, nicely detailed interview in which they outline how they achieve HA on AWS.
1. [DevOps & SRE AMA Video](http://conferences.oreilly.com/velocity/devops-web-performance-ny/public/content/devops-sre-ama-video)

    Here’s a recording of the DevOps/SRE AMA from a couple weeks back, in case you missed it.
1. [No Way Out But Through](https://blog.skyliner.io/no-way-out-but-through-1db41c648697#.xt0e1ds7h)

    A blog post by Skyline, who is launching their new deployment-as-a-service offering.  The intro is pretty great, outlining the inherent risks in changing code and releasing new code into production.
1. [gh-ost: GitHub’s online schema migration tool for MySQL](http://githubengineering.com/gh-ost-github-s-online-migration-tool-for-mysql/)

    Other online schema-change tools I’m familiar with (e.g. pt-online-schema-change) use triggers to keep a new table in sync with changes while copying old rows over.  Instead, gh-ost monitors changes by hooking on as a replication slave.  Very clever!  This article goes into several reasons why this is a much better approach.
### Outages

1. [Google App Engine](http://status.cloud.google.com/incident/appengine/16008#5673385510043648)
    The outage occurred on August 11, but they posted a postmortem this week.
1. [Buildkite](https://building.buildkite.com/outage-post-mortem-for-august-23rd-82b619a3679b#.2lh0fawla)
    Includes an extremely detailed postmortem starting with paging failure and running through 6 lessons learned.  #ThereIsNoOneRootCause
1. [Slack](https://status.slack.com/2016-08/b0d3f4837a8ae62b)
1. [Second Life](https://community.secondlife.com/t5/Tools-and-Technology/A-Play-by-Play-Retelling-of-Yesterday-s-Downtime/ba-p/3060170)
    Another awesome postmortem by April Linden.
1. [Travis CI](https://www.traviscistatus.com/incidents/4mvp857qx8bw)
1. [Facebook](http://timesofindia.indiatimes.com/tech/tech-news/Facebook-suffers-momentary-outage-back-online-now/articleshow/53885130.cms)
1. [eBay](http://www.ecommercebytes.com/C/abblog/blog.pl?/pl/2016/8/1472223471.html)
1. [PlayStation Network](http://www.pushsquare.com/news/2016/08/psn_down_just_in_time_for_the_weekend)
1. [iiNet (ISP)](http://www.itwire.com/business-technology/74504-iinet-services-in-victoria-hit-by-four-hour-outage.html)
1. [Google Compute Engine](http://status.cloud.google.com/incident/compute/16017#5659645909663744)

### [ << Prev ](sreweekly-36.md) ------------- [ Next >> ](sreweekly-38.md)