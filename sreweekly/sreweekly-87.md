## [SRE Weekly Issue #87](https://sreweekly.com/sre-weekly-issue-87/) - September 3, 2017
### Articles

1. [Multiple Perspectives On Technical Problems and Solutions](http://www.kitchensoap.com/2017/08/12/multiple-perspectives-on-technical-problems-and-solutions/)

    John Allspaw describes the Architecture Review Working Group at Etsy. I like the idea of an open discussion with peers before creating a novel system that will add significant operational burden.
1. [Post-Incident Reviews Part Two: The Demise of Root Cause Analysis](https://victorops.com/blog/post-incident-reviews-part-two-the-demise-of-root-cause-analysis/)

    Here’s part two of Jason Hand’s series of posts with key takeaways from his new eBook, “Post-Incident Reviews”. In the next three chapters, he shows why a traditional RCA process misses the mark.
1. [Observability: What’s In a Name?](https://honeycomb.io/blog/2017/08/observability-whats-in-a-name/)

    Honeycomb.io eschews plain monitoring in favor of “observability”, which they define as the ability to “ask any arbitrary question” about a system.
1. [Serverless computing: It’s all about functional stateless microservices](https://siliconangle.com/blog/2017/08/27/serverless-computing-functional-stateless-microservices/)

    Here’s another primer on microservices. It has a nice “caveats” section, which is exactly where operations and reliability come into the picture.
1. [Rapid release at massive scale](https://code.facebook.com/posts/270314900139291/rapid-release-at-massive-scale/)

    Facebook shared a lot of detail about how they evolved from 3 daily pushes to quasi-continuous releases. They’ve got a well-defined canary system, reminding me of Charity’s article on testing in production last week.
1. [10 Essential Skills of a Site Reliability Engineer](https://www.appdynamics.com/lp/10-essential-skills-of-a-site-reliability-engineer/)

    AppDynamics presents their list in shiny PDF form. You’ll have to fill in your spam-bucket address contact info to download it.
1. [Everything You Need to Know About Our Breakathon](https://www.pagerduty.com/blog/breakathon-details/)

    PagerDuty is hosting a “breakathon”: small teams will compete to resolve a series of infrastructure issues. Sounds like bunch of fun!
### Outages

1. [JapanBGPmon](https://bgpmon.net/bgp-leak-causing-internet-outages-in-japan-and-beyond/)
    Google accidentally announced some BGP prefixes it shouldn’t have, taking Japan offline for a couple of hours. Linked above is a really neat in-depth analysis from BGPmon, for all you BGP geeks out there.
Since Google essentially leaked a full table towards Verizon, we get to peek into what Google’s peering relationships look like and how their peers traffic engineer towards Google.
1. [Heroku](https://status.heroku.com/incidents/1273)
1. [AWS](http://status.aws.amazon.com/)
    EC2’s Ireland region suffered an outage in VPC peering on August 23. Their status site doesn’t allow for deep links, so here’s an excerpt:
11:32 AM PDT We are investigating network connectivity issues for some instances in the EU-WEST-1 Region.
11:55 AM PDT We have identified root cause of the network connectivity issues in the EU-WEST-1 Region. Connectivity between peered VPCs is affected by this issue. Connectivity between instances within a VPC or between instances and the Internet or AWS services is not affected. We continue to work towards full recovery.
12:51 PM PDT Between 10:32 AM and 12:44 PM PDT we experienced connectivity issues when using VPC peering in the EU-WEST-1 Region. Connectivity between instances in the same VPC and from instances to the Internet or AWS services was not affected. The issue has been resolved and the service is operating normally.
1. [Google Cloud](https://status.cloud.google.com/incident/cloud-networking/17002#5665117697998848)
    Google Cloud suffered a massive 30-hour worldwide outage in some cloud load-balancers. In their impressive style, they posted frequent updates during the incident and issued a followup analysis of the incident just 2 days after resolution.
In order to prevent the issue, Google engineers are working to enhance automated canary testing that simulates live-migration events, detection of load balancing packets loss, and enforce more restrictions on new configuration changes deployment for internal representation changes.
1. [WhatsApp](http://gadgets.ndtv.com/apps/news/whatsapp-is-down-for-many-users-1744639)
1. [Twitch (video streaming service)](http://www.esports-news.co.uk/2017/08/31/twitch-is-down/)

### [ << Prev ](sreweekly-86.md) ------------- [ Next >> ](sreweekly-88.md)