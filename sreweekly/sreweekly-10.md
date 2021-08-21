## [SRE Weekly Issue #10](https://sreweekly.com/sre-weekly-issue-10/) - February 14, 2016
### Articles

1. [How Complex Systems Fail](http://blog.acolyer.org/2016/02/10/how-complex-systems-fail/)

    So much about what modern medicine has learned about system failures applies directly to SRE, usually without any adaptation required.  In this edition of The Morning Paper, Adrian Colyer gives us his take on an excellent short paper by an MD.  My favorite quotes:

1. [Survey on Release Processes](https://sealuzh.wordpress.com/2016/02/12/survey-on-modern-release-processes-2/)

    The Software Evolution & Architecture Lab of the University of Zurich is doing a study on modern continuous deployment practices.  I’m really interested to see the results, especially where CD meets reliability, so if you have a moment, please hop on over and add your answers.  Thanks to Gerald Schermann at UZH for reaching out to me for this.
1. [What’s behind most data center outages?](https://gcn.com/articles/2016/02/09/data-center-outages.aspx?m=1)

    I’ve been debating with myself whether or not to link to Emerson Network Power’s survey of datacenter outage costs and causes.  The report itself is mostly just uninteresting numbers and it’s behind a signup-wall.  However, this article is a good summary of the report and links in other interesting stats.
1. [Under the hood: Shipping #friendsday Videos](https://code.facebook.com/posts/1705739416339373/under-the-hood-shipping-friendsday-videos/)

    Facebook algorithmically generated hundreds of millions of custom-tailored video montages for its birthday celebration.  How they did it without dedicating specific hardware to the task and without impacting production is a pretty interesting read.
1. [Scaling Elasticsearch: Sharding and Availability for Millions of Documents](https://dzone.com/articles/scaling-elasticsearch-sharding-and-availability-fo)

    Administering ElasticSearch can be just as complicated and demanding as MySQL.  This article has an interesting description of SignalFX’s method for resharding without downtime.
1. [Vision Solutions 2016 State of Resilience Report](http://visionsolutions.com/Resources/Resource/Details/2016-state-of-resilience-report?return_url=http%3a%2f%2fvisionsolutions.com%2fresources%2fwhite-papers#)

    This is a pretty interesting report that I’d never heard of before.  It’s long (60 pages), but worth the read for a few choice tidbits.  For example, I’ve seen this over and over in my career:
 Also, I was surprised that even now, over 70% of respondents said they still use “Tape Backup / Off-site Storage”.  I wonder if people are lumping S3 into that.
1. [PagerDuty Hacks: Alert by Foam Dart Bombardment](https://www.pagerduty.com/blog/pagerduty-hacks-alert-foam-dart-bombardment/)

    Never miss an ack or you’ll be in even worse trouble.
1. [IRS officials testify on outage](http://fedscoop.com/irs-officials-testify-in-oversight-hearing)

    More on last week’s outage.  I have to figure “voltage regular” means power supply.  Everyone hates simultaneous failure.
1. [Completing the Netflix Cloud Migration](https://media.netflix.com/en/company-blog/completing-the-netflix-cloud-migration)

    A full seven years after they started migration, Netflix announced this week that their streaming service is now entirely run out of AWS.  That may seem like a long time until you realize that Netflix took a comprehensive approach to the migration:

### Outages

1. [Telstra](http://www.breakingnews.com/item/2016/02/09/telstra-says-embarrassing-human-error-caused-nat/)
1. [Visual Studio Online](http://www.theregister.co.uk/2016/02/09/microsoft_sql_server_2014_bug/)
    Caused by a memory-hogging bug in MS SQL Server’s query planner.
1. [TNReady](http://tn.chalkbeat.org/2016/02/08/tnready-online-platform-crashes-causing-testing-delays-statewide/)
    Tennessee (US state) saw an outage of the new online version of its school system’s standardized tests.
1. [CBS Sports App](http://www.theverge.com/2016/2/7/10933844/super-bowl-2016-cbs-sports-app-outage-apple-tv)
    During the Super Bowl is a terrible time to fail, but of course it’s more likely due to the peak in demand.
1. [TPG Submarine Fiber Optic Cable](http://www.zdnet.com/article/tpg-submarine-cable-outage-expected-to-last-a-month/)
    This one has some really interesting discussion about how the fiber industry handles failures.
1. [Apple Pay](http://www.techweekeurope.co.uk/e-marketing/apple-pay-suffers-major-outage-185742)

### [ << Prev ](sreweekly-9.md) ------------- [ Next >> ](sreweekly-11.md)