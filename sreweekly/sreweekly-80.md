## [SRE Weekly Issue #80](https://sreweekly.com/sre-weekly-issue-80/) - July 9, 2017
### Articles

1. [Linux tracing systems & how they fit together](https://jvns.ca/blog/2017/07/05/linux-tracing-systems/)

    I had no idea there were so many tracing systems in Linux! Fortunately Julia Evans did, and she learned all about them so that she could explain them to us.
1. [So you want to be an SRE?](https://hackernoon.com/so-you-want-to-be-an-sre-34e832357a8c)

    What do you get when a high school teacher switches careers, goes to boot camp, and becomes an SRE? In this case, we get Krishelle Hardson-Hurley, who wrote this really great intro to the SRE field. She also included a set of links to other SRE materials. Thanks for the link to SRE Weekly, Krishelle!
1. [Embracing Failure in a Container World – Production Ready](https://medium.com/production-ready/embracing-failure-in-a-container-world-217a3cc414c1)

    This issue of Production Ready is a transcript (with slides) of Mathias’s talk at ContainerDays on doing chaos engineering in a container-based infrastructure. I really like the idea of attaching a side-car container to inject latency using tc.
1. [Why is Redfin running its site from a single data center without a backup facility?](https://www.geekwire.com/2017/redfin-running-site-single-data-center-without-backup-facility/)

    Here’s an interesting side-effect from an IPO: Redfin was obliged to mention the fact that its website runs out of a single datacenter.
1. [Event Foo: Designing for Results](https://honeycomb.io/blog/2017/07/event-foo-designing-for-results/)

    This article, part of a series from Honeycomb.io on structured event logging, contains some tips on structuring your events well to get the most out of your logs.
1. [The Peculiarities of High-Availability Data Center Design on a Cruise Ship](http://www.datacenterknowledge.com/archives/2017/07/06/peculiarities-high-availability-data-center-design-cruise-ship/)

    I’d never thought about what IT systems must exist on a cruise ship before. This article left me wanting to know more, so I found this ZDNet article with pictures and descriptions of another cruise ship datacenter layout.
### Outages

1. [Summary of the December 24, 2012 Amazon ELB Service Event in the US-East Region](https://aws.amazon.com/message/680587/)
    Here’s a classic postmortem from Amazon, in which a developer inadvertently deleted the production ELB state information.
1. [Slack: This was not normal. Really.post-mortems repo in github](https://slackhq.com/this-was-not-normal-really-230c2fd23bdc)
    I’d forgotten about this superlative example of an incident followup posting from Slack after a pair of outages in 2014. What reminded me was a commit to Dan Luu’s post-mortems repo in github that mentioned it.
1. [Heroku Incident 372: HTTP Routing Errors](https://status.heroku.com/incidents/372)
    Here’s another classic incident followup posting. Heroku spills the details on a major outage that cut off access to all applications for 30 minutes in 2012.Full disclosure: Heroku is my employer.

### [ << Prev ](sreweekly-79.md) ------------- [ Next >> ](sreweekly-81.md)