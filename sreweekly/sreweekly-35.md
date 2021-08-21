## [SRE Weekly Issue #35](https://sreweekly.com/sre-weekly-issue-35/) - August 14, 2016
### Articles

1. [AWS Networking, Environments and You](http://charity.wtf/2016/03/23/aws-networking-environments-and-you/)

    Whoops, here’s one that got lost in my review queue.  Charity Majors (one of the usual suspects here at SRE Weekly) wrote one of her characteristically detailed and experience-filled posts on how to isolate your production, staging, and development environments in AWS.
1. [Paradigm Check Point: Prefacing Debriefings](http://www.kitchensoap.com/2014/03/19/paradigm-check-point-prefacing-debriefings/)

    I can’t quite tell how much of this is John Allspaw’s writing and how much is written by the US Forestry Service, but I love it all.  Here’s a bulleted list of points driving home the fact that we constantly strike a balance between risk and safety.
1. [Multi data center redundancy – application considerations](https://blog.serverdensity.com/multi-data-center-redundancy-application-considerations/)

    Server Density added multi-datacenter redundancy to their infrastructure in 2013, and they were kind enough to document what they learned.  In this first of two articles, they outline different kinds of multi-datacenter setups and go over the kinds of things you’ll have to think about as you retrofit your application.
1. [Making a point with SLAs](https://blog.serverdensity.com/making-a-point-with-slas/)

    This short opinion piece raises an excellent idea: SLAs aren’t for recouping the cost you incurred due to an outage, they are for making a point to a service provider about the outage.
1. [Cost of Southwest’s tech outage climbs to at least $54 million](http://www.dallasnews.com/business/airline-industry/20160810-cost-of-southwest-s-tech-outage-climbs-to-at-least-54-million.ece)

    Southwest has released some numbers on the impact of last month’s outage that resulted in thousands of cancelled flights.
1. [Netflix and Fill](http://techblog.netflix.com/2016/08/netflix-and-fill.html)

    Netflix gives us a rundown of how they prepare a title for release by pre-filling caches in their in-house CDN.  I like the part about timing pre-filling during off-peak hours to avoid impacting the service.
1. [Delta Datacenter Crash: Do the Math on Disaster Recovery ROI](http://www.nextplatform.com/2016/08/09/delta-datacenter-crash-math-disaster-recovery-roi/)

    How much is your company willing to invest for a truly effective DR solution?  This article asks that question and along the way digs into what an effective DR solution looks like and why it costs so much.
### Outages

1. [Syria](http://www.ibtimes.co.uk/syrias-mysterious-internet-outages-may-have-finally-been-solved-1575750)
    The Syrian government shut internet access down to prevent cheating on school exams.
1. [Mailgun](http://blog.mailgun.com/mailgun-api-outage-post-mortem-august-2016/)
    Linked, find a really interesting postmortem: Mailgun experienced an outage when their domain registrar placed their domain on hold abruptly.  The registrar was subsequently largely uncommunicative, hampering incident resolution.  Lesson learned: make sure you can trust your registrar, because they have the power to ruin your day.
1. [Belnet](http://www.datacenterdynamics.com/content-tracks/security-risk/optical-card-failure-brings-down-belgiums-belnet/96758.fullarticle)
    The linked article has some intriguing detail about a network equipment failure that caused a routing loop.
1. [Australia’s census website](http://www.crn.com.au/news/ibm-breaks-silence-on-census-fail-433502)
    This caught my eye:
Revolution IT simulated an average sustained peak of up to 350 submissions per second, but only expected up to 250 submission per second.
Load testing only 40% above expected peak demand?  That seems like a big red flag to me.
1. [Reddit](http://www.thebitbag.com/reddit-reason-site-offline/171592)
1. [Etisalat (UAE ISP)](http://www.thenational.ae/uae/technology/etisalat-internet-outages-reported-across-uae-for-second-day)
1. [Vodafone](http://www.nltimes.nl/2016/08/10/nationwide-vodafone-4g-outage-resolved-hours-offline/)
1. [Google Drive](http://www.androidpolice.com/2016/08/08/psa-google-drive-many-today-google-working/)
1. [AT&T](http://www.kansascity.com/news/local/article95516492.html)
1. [Delta Airline](http://bgr.com/2016/08/14/delta-finally-explained-how-one-power-outage-grounded-an-entire-airline/)
    A datacenter power system failure resulted in cancelled flights worldwide.

### [ << Prev ](sreweekly-34.md) ------------- [ Next >> ](sreweekly-36.md)