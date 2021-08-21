## [SRE Weekly Issue #81](https://sreweekly.com/sre-weekly-issue-81/) - July 16, 2017
### Articles

1. [Failure Fridays: Four Years On](https://www.pagerduty.com/blog/failure-fridays-four-years/)

    PagerDuty shared this timeline of their progress in adopting Chaos Engineering through their Failure Friday program. This is brilliant:
1. [How Platforms and SREs Change the DevOps Contract](https://developer.capitalone.com/blog-post/how-platforms-and-sres-change-the-devops-contract/)

    I’m a big proponent of having developers own their code in production. This article posits that SRE’s job is to provide a platform that enables developers to do that more easily. I like the idea that containers and serverless are ways of getting developers closer to operations.
1. [Interview with AWS’s Werner Vogels about response to Amazon outages](http://uk.businessinsider.com/interview-aws-werner-vogels-amazon-outages-2017-7)

    This reads less like an interview and more like a description of Amazon’s incident response procedure. I started paying close attention at step 3, “Learn from it”:
1. [From Scala Unified Logging to Full System Observability — Part 1 of 3: Our Original State of Logging](https://victorops.com/blog/from-scala-unified-logging-to-full-system-observabilitypart-1-our-original-state-of-logging/)

    This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Site Reliability Engineer: Don’t fall victim to the bias blind spot](http://sdtimes.com/site-reliability-engineer-dont-fall-victim-to-the-bias-blind-spot/)

    This article is about a different kind of human factor than articles I often link to: cognitive bias. The author presents a case for SREs as working to limit the effects of cognitive bias in making operational decisions.
### Outages

1. [OVHOVH](https://www.theregister.co.uk/2017/07/13/watercooling_leak_killed_vnx_array/)
    OVH suffered a major outage in a datacenter, taking down 50,000 websites that they host. The outage was caused by a leak in their custom water-cooling system and resulted in a painfully long 24-hour recovery from an offsite backup. The Register’s report (linked) is based on OVH’s incident log and is the most interesting datacenter outage description I’ve read this year.
1. [Google Cloud Storage](http://status.cloud.google.com/incident/storage/17002#5713144022302720)
    Google posted this followup for an outage that occurred on July 6th. As usual, it’s an excellent read filled with lots of juicy details. This caught my eye:
[…] attempts to mitigate the problem caused the error rate to increase to 97%.
Apparently this was caused by a “configuration issue” and was quickly reverted. It’s notable that they didn’t include anything about this error in the remediations section.
1. [Melbourne, AU’s Metro rail network](http://www.theage.com.au/victoria/metros-80m-backup-system-may-have-failed-at-moment-melbourne-needed-it-20170714-gxbjs6.html)
    A network outage stranded travelers, and switching to the DR site “wasn’t an option”.
1. [Somalia](http://www.businessghana.com/site/news/politics/148521/Somalia-internet-outage-is-)

### [ << Prev ](sreweekly-80.md) ------------- [ Next >> ](sreweekly-82.md)