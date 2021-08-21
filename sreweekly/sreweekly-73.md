## [SRE Weekly Issue #73](https://sreweekly.com/sre-weekly-issue-73/) - May 21, 2017
### Articles

1. [Troubleshooting Unusual AWS ELB 5XX Error](http://engineering.grab.com/troubleshooting-unusual-aws-elb-5xx-error)

    ELBs (Amazon’s Elastic Load Balancers) depend on clients properly respecting DNS round-robin record sets. This article follows a debugging session in excellent detail as they try to answer the question: why are our clients preferring (and overloading) just one ELB IP?
1. [Serverless: A replacement for servers?](https://duringthedrive.com/2017/05/15/serverless-a-replacement-for-servers/)

    Sarah Schieffer Riehl shares her take on ServerlessConf Austin 2017. She’s got a healthy dose of skepticism that I like, concluding that “serverful and serverless architectures don’t do the same things.” I like this bit:
1. [Premortems: The Art of Negative Visualization](https://medium.com/production-ready/premortems-the-art-of-negative-visualization-583081559ea1)

    Wow, this dovetails so well into the Todd Conklin’s “Safety Moment” from last week, on imagining all the possible things that could go wrong.  I’d love to hear more thoughts along these lines: is it possible to design a reliable system without envisioning the majority of things that could go wrong?
1. [Keep Critical Apps and Infrastructure Up and Running](https://www.pagerduty.com/blog/incident-lifecycle-management/)

    PagerDuty outlines an incident lifecycle management policy based on ITIL.
1. [Introducing Cape](https://blogs.dropbox.com/tech/2017/05/introducing-cape/)

    DropBox created Cape for “asynchronous processing of billions of events a day, powering many Dropbox features”. Example: you upload a text file, and a Cape job indexes it immediately for full-text searching. I’d love to hear more on why existing solutions didn’t fit the bill, although they do cover their requirements in depth.
1. [What a SaaS outage taught me about the meaning of partnerships](http://www.itproportal.com/features/what-a-saas-outage-taught-me-about-the-meaning-of-partnerships/)

    When I signed on for my first SRE position, I had no idea how huge a part vendor relations would play in ensuring reliability.
1. [Building the SRE Culture at LinkedIn](https://engineering.linkedin.com/blog/2017/05/building-the-sre-culture-at-linkedin)

    Initially, LinkedIn’s SRE team hired engineers only based on technical skill. As they’ve grown, they’ve discovered the importance of collaboration skills as well.
1. [Incident communication best practices](http://blog.statuspage.io/incident-communication-best-practices)

    StatusPage.io explains the reasons for having a solid incident communication policy and guides you through setting one up.
1. [The Calculus of Service Availability](http://queue.acm.org/detail.cfm?id=3096459)

    As the title suggest, this ACM Queue article goes into some depth on the kinds of calculations one might make when designing a reliable system. Specifically, they focus on service dependencies and introduce Google’s “rule of the extra 9”: a dependency should have one more nine of reliability than the thing that critically depends on it.
1. [It Takes More Than a Circuit Breaker to Create a Resilient Application](https://dzone.com/articles/it-takes-more-than-a-circuit-breaker-to-create-a-r)

    
### Outages

1. [Starbucks](https://patch.com/washington/seattle/starbucks-down-stores-across-country-affected-reported-outage)
    A server outage halted sales at many stores, and some gave out free drinks to mollify customers. Coincidentally, I also was unable to order at Wendy’s the other night due to a “server update”, and they offered me a free Frosty.
1. [Let’s EncryptOCSP outagewouldn’t actually cause any impact](https://letsencrypt.status.io/pages/incident/55957a99e800baa4470002da/591e962c4f9ef22239001819)
    Certificate issuance was impaired for about 17 hours. They also had an OCSP outage around the same time, but as far as I can tell, this wouldn’t actually cause any impact to end-users of Let’s Encrypt certificates.
1. [Twitter](http://www.dailymail.co.uk/sciencetech/article-4522504/Twitter-Glitch-panics-users-worldwide.html)
1. [Whatsapp](http://allafrica.com/stories/201705190233.html)
1. [AT&T Gets Light FCC Wrist Slap For Largest 911 Outage Ever](http://www.dslreports.com/shownews/ATT-Gets-Light-FCC-Wrist-Slap-For-Largest-911-Outage-Ever-139600)
    The FCC released a report on AT&T’s 911 outage last March. The cause was apparently a faulty whitelist update.
1. [Instagram](http://www.express.co.uk/life-style/science-technology/805454/Instagram-Down-Photo-Sharing-Service-Not-Working)

### [ << Prev ](sreweekly-72.md) ------------- [ Next >> ](sreweekly-74.md)