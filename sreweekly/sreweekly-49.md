## [SRE Weekly Issue #49](https://sreweekly.com/sre-weekly-issue-49/) - November 27, 2016
### Articles

1. [Price-Checking for Prescriptions Results in Dangerous Combination of Medications](https://bwhsafetymatters.org/2016/07/20/price-checking-for-prescriptions-results-in-dangerous-combination-of-medications/)

    Another issue of BWH’s Safety Matters, this time about a prescribing accident. The system seems to have been almost designed to cause this kind of error, so it’s good to hear that a fix was already about to be deployed.
1. [Do You Make This Critical Root Cause Analysis (RCA) Mistake?](http://www.meddeviceonline.com/doc/do-you-make-this-critical-root-cause-analysis-rca-mistake-0001)

    This is a great article on identifying the true root cause(s) of an incident, as opposed to stopping with just a “direct cause”. I only wish it were titled, Use These Five Weird Tricks to Run Your RCA!
1. [How Etsy Uses Code “Slush” to Manage Development During the Holidays](https://codeascraft.com/2016/11/16/code-slush-holidays/)

    Etsy describes how they do change management during the holidays:
1. [Production Ready: Always Leave the Campground Cleaner Than You Found It](http://mail01.tinyletterapp.com/production-ready/always-leave-the-campground-cleaner-than-you-found-it/6751777-medium.com/production-ready/always-leave-the-campground-cleaner-than-you-found-it-b176a91bc0b5?c=c89e58e4-6eb7-4315-b12e-c8f3dae7ed01)

    This issue of Production Ready is about battling code rot through incrementally refactoring an area of a codebase while you’re doing development work that touches it.
1. [5 ways to hone your production incident postmortems](http://tech.shutterstock.com/2016/11/11/5-ways-to-hone-your-production-incident-postmortems/)

    Shutterstock shares some tips they’ve learned from writing postmortems. My favorite part is about recording a timeline of events in an incident. I’ve found that reading an entire chat transcript for an incident can be tedious, so it can be useful to tag items of interest using a chat-bot command or a search keyword so that you can easily find them later.
1. [OUTAGE! AMA on-demand video](http://pages.catchpoint.com/OUTAGE-AMA.html)

    The “Outage!” AMA happened while I was on vacation, and I still haven’t had a chance to listen to it. Here’s a link in case you’d like to.
1. [10 DevOps Interview Questions to Gauge a Candidate’s Real Knowledge](https://dzone.com/articles/10-devops-interview-questions-to-gauge-a-candidate)

    My favorite:
1. [Weaver: Ill-Behaved Microservice Emulator](https://dzone.com/articles/weaver-ill-behaved-microservice-emulator)

    Weaver is a tool to help you identify problems in your microservice consumers by doing “bad” things like responding slowly to a fraction of requests.
1. [How Barclays Avoids Downtime Chaos](http://www.techweekeurope.co.uk/data-storage/how-barclays-avoids-downtime-chaos-200812)

    Barclays reduced load on their mainframe by adding MongoDB as a caching layer to handle read requests. What the heck does “mainframe” mean in this decade, anyway?
1. [SOASTA Report: Online Holiday Shoppers Will Only Wait for Two Seconds](http://www.prweb.com/releases/2016/11/prweb13858053.htm)

    We’d do well to remember during this holiday season that several seconds of latency in web requests is tantamount to an outage.
1. [You Are Not Paid to Write Code](http://bravenewgeek.com/you-are-not-paid-to-write-code/)

    Tyler Treat gives us an eloquently presented argument for avoiding writing code as much as possible, for the sake of stability.
### Outages

1. [Everest (datacenter)](http://everestdc.com/documents/rfo/2016-11-15_IP_Network_Incident.pdf)
    Everest suffered a cringe-worthy network outage subsequent to a power failure. Power came off and on a couple of times, prompting their stacked Juniper routers to assume they’d failed in booting and go into failure recovery mode. Unfortunately, the secondary OS partitions on the two devices contained different JunOS versions, so they wouldn’t stack properly.
I’d really like to read the RFO on the power outage itself, but I can’t find it. If anyone has a link, could you please send it my way?
1. [Argos](http://www.ibtimes.co.uk/argos-website-down-outage-hits-just-before-launch-black-friday-deals-1592047)

### [ << Prev ](sreweekly-48.md) ------------- [ Next >> ](sreweekly-50.md)