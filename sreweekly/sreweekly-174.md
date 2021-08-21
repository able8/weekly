## [SRE Weekly Issue #174](https://sreweekly.com/sre-weekly-issue-174/) - June 23, 2019
### Articles

1. [What bugs cause cloud production incidents?](https://blog.acolyer.org/2019/06/21/what-bugs-cause-cloud-production-incidents/)

    Adrian Colyer (summary)Liu et al., HotOS’19 (original paper)
1. [Listen to a Recorded Incident Response Call](https://www.pagerduty.com/blog/incident-response-reenactment/)

    PagerDuty created this re-enactment of an incident response phone bridge. It’s obviously fairly heavily redacted and paraphrased, but it’s still quite educational. It includes interludes where terms such as Incident Commander are explained.George Miranda — PagerDuty
### Outages

1. [Google Calendar](https://www.google.com/appsstatus#hl=en&v=issue&sid=2&iid=cc21ebe3962430b2e4ae2b52e3dde98f)
1. [Netflix](https://comicbook.com/movies/2019/06/21/netflix-crash-freaks-out-subscribers-reactions/)
1. [Hulu](https://www.digitaltrends.com/news/hulu-xbox-live-down-outage/)
1. [Joyent May 27 2014 outage followup](https://www.joyent.com/blog/postmortem-for-outage-of-us-east-1-may-27-2014/)
    In this 2014 outage followup, we learn that a Joyent engineer accidentally rebooted an entire datacenter:
The command to reboot the select set of new systems that needed to be updated was mis-typed, and instead specified all servers in the data center.
1. [Salesforce May 17 outage followup](https://help.salesforce.com/articleView?id=000320234&mode=1&type=1)
    Click through to read about the massive Salesforce outage last month. A database edit script contained a bug that ran an UPDATE without its WHERE clause, granting elevated permissions to more users than intended. Salesforce shut down broad chunks of their service to prevent data leakage.
1. [Second Life mid-May outage followup](https://community.secondlife.com/blogs/entry/2550-the-road-to-downtime-was-paved-with-good-intentions/)
    Linden Lab posted about a network maintenance that went horribly wrong, resulting in a total outage.
Everything started out great. We got the first new core router in place and taking traffic without any impact at all to the grid. When we started working on the second core router, however, it all went wrong.
April Linden — Linden Lab
1. [Monzo May 30 outage followup](https://monzo.com/blog/2019/06/20/why-bank-transfers-failed-on-30th-may-2019/)
    Monzo posted this incredibly detailed followup for an outage from several weeks ago. Not only does it give us a lot of insight into their incident response process, I also got to learn about how UK bank transfers work.Thanks to an anonymous reader for this one.
Nicholas Robinson-Wall — Monzo
1. [Google Cloud Platform June 2 outage followup](https://status.cloud.google.com/incident/cloud-networking/19009)
    Along with the blog post I linked to last week, Google also posted this technical followup for their major June 2 outage. I’ve never seen one of their followups even close to this long or detailed, and that’s saying a lot.

### [ << Prev ](sreweekly-173.md) ------------- [ Next >> ](sreweekly-175.md)