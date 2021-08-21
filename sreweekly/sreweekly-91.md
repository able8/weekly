## [SRE Weekly Issue #91](https://sreweekly.com/sre-weekly-issue-91/) - October 1, 2017
### Articles

1. [No brogrammers: Practical tips for writing inclusive job ads](https://blog.hostedgraphite.com/2016/04/26/no-brogrammers-practical-tips-for-writing-inclusive-job-ads/)

    Here’s how Hosted Graphite made their job ad for an SRE-like role (Ops Automation Engineer) more inclusive. The article is filled with specific before/after language snippets, each with a detailed explanation of why they made the change.
1. [Advanced Secondary DNS For The Technically Inclined](https://dyn.com/blog/advanced-secondary-dns-for-the-technically-inclined/)

    A couple weeks after their major outage last October, Dyn published this article explaining secondary DNS. It’s a great primer and digs into what to do if you use advanced non-standard functionality like ALIAS records and traffic balancing.
1. [Predicting Resource Exhaustion with Double Exponential Smoothing](http://t.dripemail2.com/c/eyJhY2NvdW50X2lkIjoiMjY1Njc0MyIsImRlbGl2ZXJ5X2lkIjoiMTM0OTM0MDg1MSIsInVybCI6Imh0dHBzOi8vc2lnbmFsZnguY29tL2Jsb2cvcHJlZGljdGluZy1yZXNvdXJjZS1leGhhdXN0aW9uLWRvdWJsZS1leHBvbmVudGlhbC1zbW9vdGhpbmcvP19fcz1id3lrd2sxa2NjZW9nc3pxOGFidCJ9)

    SignalFx goes into deep detail on their feature for predicting future metric values. We get an explanation of why prediction is difficult and a discussion of the math involved in their solution.
1. [Handling System Failures during Payment Communication](https://blogs.dropbox.com/tech/2017/09/handling-system-failures-during-payment-communication/)

    Payments: we really have to get them right. Here’s DropBox’s Jessica Fisher with a discussion of how they reconcile failed payments.
1. [Circuit Breaker Implementation in Resilience4j](https://dzone.com/articles/circuit-breaker-implementation-in-resilience4j)

    A couple of weeks ago, I linked to a story about Resilience4j, a fault tolerance library for Java. This week is the second installment that shows you how to use it to implement circuit breakers. There’s also an interesting discussion of one of the implementation details.
1. [ntpd won’t save you from one particular rogue bit](https://rachelbythebay.com/w/2017/09/27/2153/)

    Here’s a cute little debugging story. Turns out ntpd has a bit of a blind spot!
1. [Adcash – 1 Trillion HTTP Requests Per Month](http://feedproxy.google.com/~r/HighScalability/~3/iRsnmCaun6w/adcash-1-trillion-http-requests-per-month.html)

    Adcash CTO Arnaud Granal gives us a rare glimpse into the multiple iterations of their infrastructure. Hear what worked well and what didn’t as they scaled to handle 500k requests per second at peak.
### Outages

1. [OpenSRS (DNS provider)OpenSRS](https://www.theregister.co.uk/2017/09/29/major_domain_name_server_goes_down_taking_websites_with_it/)
    OpenSRS (registrar and DNS provider, among other services) had a major outage in their DNS service.
At 1AM UTC we were the target of a sophisticated DNS attack that was followed by an unrelated double failure of core network equipment at our main Canadian data center, caused by an undocumented software limitation.
Yikes.
1. [Amadeus (airline booking system)Amadeus](https://www.ttgmedia.com/news/technology/amadeus-suffers-major-it-network-issue-11742)
    Amadeus provides the technical underpinnings of many airlines around the world. They had issues this past week, taking a lot of airlines with them.
1. [SourceForge](https://www.theregister.co.uk/2017/09/27/faulty_data_center_takes_out_sourceforge/)
    Our [data center] hosting provider has been having issues with a power distribution unit.
1. [Facebook](http://www.express.co.uk/life-style/science-technology/858892/Facebook-down-not-working-server-status)

### [ << Prev ](sreweekly-90.md) ------------- [ Next >> ](sreweekly-92.md)