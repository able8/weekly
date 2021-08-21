## [SRE Weekly Issue #44](https://sreweekly.com/sre-weekly-issue-44/) - October 16, 2016
### Articles

1. [The Ops Identity Crisis](http://www.susanjfowler.com/blog/2016/10/13/the-ops-identity-crisis)

    With all the “NoOps” and “Serverless” stuff floating around, do we need ops? Susan Fowler says not necessarily, but that we do need ops skills.
1. [VictorOps State of On-Call Survey](https://victorops.az1.qualtrics.com/SE/?SID=SV_3mfwsMp1IuzyiCp&source=sreweekly)

    VictorOps is gathering data for the 2016 edition of their yearly State of On-Call Report (2015’s if you missed it). Please click the link above and take the survey if you have a moment! The report provides some pretty awesome stats that we can all use to improve the on-call experience at our organizations.This survey is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Irreversible Failures: Lessons from the DynamoDB Outage](http://blog.scalyr.com/2015/09/irreversible-failures-lessons-from-the-dynamodb-outage/)

    Scalyr writes about cascading failure scenarios, using the DynamoDB outage of September 20th, 2015 (no, not this year’s September DynamoDB outage) as a case study.
1. [cron.weekly newsletter](https://www.cronweekly.com)

    Check it out! Apparently this newsletter started around the same time that SRE Weekly did. Content includes a lot of really nifty stuff about Linux system administration.
1. [Deployer API Outage Postmortem](https://gist.github.com/mlafeldt/01953b48b0be5fea34c11a8a47d1e7f4)

    I previously linked to a two–part series by Mathias Lafeldt on writing postmortems. At my request, Jimdo graciously agreed to release their (previously) internal postmortem about the incident that prompted him to write the articles. Thanks so much, Mathias!
1. [Human Factors at The Fringe: My Eyes Went Dark](https://humanisticsystems.com/2016/09/09/human-factors-at-the-fringe-my-eyes-went-dark/)

    A review of what sounds like a really interesting play about just culture, blameless retrospectives, and restorative justice in aviation, based on real events.Thanks to Mathias Lafeldt for this one.
1. [Building one of the highest-capacity subsea cables in the Pacific](https://code.facebook.com/posts/1184407111619685/building-one-of-the-highest-capacity-subsea-cables-in-the-pacific/)

    When you’re big like Facebook, sometimes reliability means essentially building your own Internet.
1. [Lessons Learned from Scaling Uber to 2000 Engineers, 1000 Services, and 8000 Git repositories](http://highscalability.com/blog/2016/10/12/lessons-learned-from-scaling-uber-to-2000-engineers-1000-ser.html)

    If you haven’t had time to watch Matt Ranney’s talk on Scaling Uber to 1000 Microservices, check out this detailed summary. Growing your engineering force 10x over a year while still keeping the service reliable is a pretty impressive feat.
1. [6 Essential Steps to Reducing Incident Resolution Time](https://www.pagerduty.com/blog/6-steps-reducing-incident-resolution-time/)

    PagerDuty shares some tips for lowering your MTTR, but first they ask the important question: how are you measuring MTTR, and is lowering it meaningful?
1. [Put Down the Mouse, Step Away From the Keyboard](https://www.linkedin.com/pulse/put-down-mouse-step-away-from-keyboard-david-christensen)

    David Christensen riffs on Charity Majors’s concept of “3 Types of Code”: “no code” (SaaS, PaaS, etc), “someone else’s code”, and “your code”. Try to spend as much development time as possible writing code that supports what makes your business unique (your key differentiator).
1. [Operations for software developers for beginners](https://jvns.ca/blog/2016/10/15/operations-for-software-developers-for-beginners/)

    Julia Evans is back with a write-up of the lessons she’s learned as she’s begun to gain an understanding of operations. My favorite bit:
1. [SysAdvent 2016 Author/Editor Signup](https://goo.gl/forms/NX3P1mokKRjVmKGf2)

    SysAdvent is happening again this year! Click the link above if you’d like to propose an article or volunteer to be an editor.
### Outages

1. [United Airlines](http://www.wrcbtv.com/story/33390000/united-airlines-systems-outage-causes-delays-globally)
1. [Yahoo mail](http://www.heraldcourier.com/news/yahoo-restores-automatic-email-forwarding-after-brief-outage/article_6599ffe4-90d7-5278-8c89-2b1536b81d44.html)
1. [Google Cloud](http://status.cloud.google.com/incident/compute/16020#5640471028170752)
1. [FNB (South Africa bank)](http://mybroadband.co.za/news/banking/182418-fnb-online-banking-and-app-outage.html)
1. [GlobalSign (SSL certificate authority)postmortem](https://www.globalsign.com/en/customer-revocation-error/)
    GlobalSign had a major problem in their PKI that resulted in all of their certificates being treated as revoked. They’ve posted a detailed postmortem that’s pretty heavy on deep SSL details, but the basic story is that their OCSP service misinterpreted a routine action as a request to revoke their intermediate CA certificate. Yikes.I love this quote and the mental image of a panicked party with streamers and ribbon-cutting that it conjures up:
Our AlphaSSL and CloudSSL customers had to wait a few hours more while an emergency key ceremony was held to create alternatives.

### [ << Prev ](sreweekly-43.md) ------------- [ Next >> ](sreweekly-45.md)