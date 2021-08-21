## [SRE Weekly Issue #62](https://sreweekly.com/sre-weekly-issue-62/) - March 5, 2017
### Articles

1. [Think of Latency as a Pseudo-permanent Network Partition](http://highscalability.com/blog/2010/8/12/think-of-latency-as-a-pseudo-permanent-network-partition.html)

    When you do as the title suggests, you realize that network partitions go from the realm of theoretical to everyday.
1. [Ask 5 Whys to get to the root of any problem](https://blog.asana.com/2015/06/workstyle-ask-5-whys-to-get-to-the-root-of-any-problem/)

    Asana shares their “Five Whys” process, which they use not only for outages but even for missed deadlines. This caught my eye:
1. [Organizing Software Deployments to Match Failure Conditions](https://www.awsarchitectureblog.com/2014/05/organizing-software-deployments-to-match-failure-conditions.html)

    Using Route 53 as a case study, AWS engineers explain how they carefully designed their deploy process to reduce impact from failed deploys.
1. [Moving persistent data out of Redis](http://githubengineering.com/moving-persistent-data-out-of-redis/)

    GitHub used a data-driven approach when migrating a storage load from Redis to MySQL. It’s a good thing they did, because a straight one-for-one translation would have overloaded MySQL.
1. [Actionable Alerts](https://victorops.com/blog/actionable-alerts/)

    We’ve heard before that it’s important to make sure that your alerts are actionable. I like that this article goes into some detail on why we sometimes tend to create inactionable alerts before explaining how to improve your alerting.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [How to stop Ubuntu Xenial from randomly killing your big processes](https://blog.meteor.com/how-to-stop-ubuntu-xenial-from-randomly-killing-your-big-processes-4a3e2d09323f#.44zwekwch)

    Ubuntu backported a security fix into Xenial’s kernel last month, and unfortunately, they introduced a regression. Under certain circumstances, the kernel will give up way too easily when attempting to find memory to satisfy an allocation and will needlessly trigger the OOM killer. A fix was released on February 20th.
1. [Beating the CAP Theorem Checklist](http://ferd.ca/beating-the-cap-theorem-checklist.html)

    Need to tell someone their perpetual motion machine CAP-satisfying system won’t work? Low on time? Use this handy checklist to explain why their idea won’t work.
1. [Why we are not leaving the cloud](https://about.gitlab.com/2017/03/02/why-we-are-not-leaving-the-cloud/)

    GitLab seriously considered fleeing the cloud for a datacenter, and they asked the community for feedback. That feedback was very useful and was enough to change their minds. The common theme: “you are not an infrastructure company, so why try to be one?”
1. [Instrumenting High Volume Services: Part 2](https://honeycomb.io/blog/2017/03/instrumenting-high-volume-services-part-2/)

    If you’ve got a firehose of events going into your metrics/log aggregation system, you may need to reduce load on it by only sending in a portion of your events. Do you pick one out of every N? HoneyComb’s makers suggest an interesting alternative: tag each sampled event you send as representing N events from the source — and N is allowed to very between samples.
### Outages

1. [Amazon S3TrelloHeroku,SlackGitHubstatus page](https://aws.amazon.com/message/41926/)
    Amazon S3 in the us-east-1 region went down, taking many sites and services down with it, including Trello, Heroku, portions of Slack and GitHub, and tons more. Amazon’s status page had a note at the top but was otherwise green across the board for hours.  Meanwhile nearly 100% of S3 requests failed and many other AWS services burned as well.Their outage summary (linked above) indicated that the outage uncovered a dependency of their status site on S3. Oops. Once they got that fixed a few hours later, they posted something I’ve never seen before: actual red icons.Full disclosure: Heroku is my employer.
1. [Joyent: Postmortem for July 27 outage of the Manta service](https://www.joyent.com/blog/manta-postmortem-7-27-2015)
    Here’s a deeply technical post-analysis of a Postgresql outage that Joyent experienced in 2015. A normally benign automatic maintenance (an auto-vacuum) turned into total DB lockup due to their workload.
1. [PagerDuty](https://status.pagerduty.com/incidents/70m30bh7qfmx)
1. [GoDaddy](https://www.theregister.co.uk/2017/03/02/godaddy_dns_has_gone_diddy/)
    DDoS attack on their nameservers.

### [ << Prev ](sreweekly-61.md) ------------- [ Next >> ](sreweekly-63.md)