## [SRE Weekly Issue #93](https://sreweekly.com/sre-weekly-issue-93/) - October 15, 2017
### Articles

1. [Reasons Kubernetes is cool](https://jvns.ca/blog/2017/10/05/reasons-kubernetes-is-cool/)

    Julia Evans tells us why she likes Kubernetes, and along the way explains how its resilient architecture works.
1. [distsys-class/README.markdown at master · aphyr/distsys-class · GitHub](https://github.com/aphyr/distsys-class/blob/master/README.markdown)

    From the Jepsen folks, this outline is detailed enough to read by itself:
1. [When Optimising For Robustness Fails](http://www.alwaysagileconsulting.com/articles/when-optimising-for-robustness-fails/)

    
1. [What will programming look like in the future?](http://feedproxy.google.com/~r/HighScalability/~3/r60BASKWgrU/what-will-programming-look-like-in-the-future.html)

    This article seems like a direct reply to last week’s “The Coming Software Apocalypse“. I gave that one a good review, so I feel compelled to include this refutation, but I was left really wishing for more detail on the arguments put forward. Perhaps there’s more to come?
1. [Monitoring in the time of Cloud Native](https://medium.com/@copyconstruct/monitoring-in-the-time-of-cloud-native-c87c7a5bfa3e?__s=bwykwk1kcceogszq8abt)

    This is an article version of Cindy Sridharan’s Velocity 2017 talk. She covers a lot, including major monitoring methods, existing OSS tools, the pitfalls of each, and how to achieve observability in a cloud-based infrastructure.
1. [Mitigating replication lag and reducing read load with freno](https://githubengineering.com/mitigating-replication-lag-and-reducing-read-load-with-freno/)

    GitHub ensures low MySQL replication lag by rate-limiting expensive batch-processing queries based on replica lag. Before freno, this logic resided in each client, with multiple implementations in different languages. Freno (which is open source) centralizes the replica lag polling and query rate-limiting decisions into a queryable service.
1. [Open Sourcing Iris and Oncall](https://engineering.linkedin.com/blog/2017/06/open-sourcing-iris-and-oncall?__s=bwykwk1kcceogszq8abt)

    Earlier this year, LinkedIn open sourced their alerting system duo. Together, these tools provide functionality similar to vendor solutions like PagerDuty and VictorOps.
1. [NGINX Rate Limiting](https://dzone.com/articles/nginx-rate-limiting)

    Here’s a great guide to rate-limiting in NGINX including config snippets.
1. [Developer Experience Lessons Operating a Serverless-like Platform At Netflix](https://medium.com/netflix-techblog/developer-experience-lessons-operating-a-serverless-like-platform-at-netflix-part-ii-63a376c28228?source=rss----2615bd06b42e---4)

    Netflix has an in-house serverless environment on which they run “nano-services”. It has nifty features including automatic pre-warming, gradual roll-out scheduling, and canary deployments.
1. [Transit and Peering: How your requests reach GitHub](https://githubengineering.com/transit-and-peering-how-your-requests-reach-github/)

    GitHub details their Internet-facing network topology and explains how they use traffic engineering to ensure their connectivity is fast and reliable.
1. [The pitfalls of A/B testing in social networks](https://tech.okcupid.com/the-pitfalls-of-a-b-testing-in-social-networks/)

    What if two people try to interact, but only one of them is flagged into a new feature? OKCupid tells us why A/B testing is much harder than it seems, and then they explain how they developed useful test cohorts.
1. [Focus on Remediation: Leverage Runbooks to Reduce MTTR](https://victorops.com/blog/focus-on-remediation-leverage-runbooks-to-reduce-mttr/)

    A primer on runbooks, including a nice template you can use as a starting point in writing yours.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
### Outages

1. [Facebook/Instagram](http://thenewdaily.com.au/life/tech/2017/10/12/facebook-instagram-global-outage/)
1. [Pingdom](http://status.pingdom.com/incidents/9f1x0g2dbv0w)

### [ << Prev ](sreweekly-92.md) ------------- [ Next >> ](sreweekly-94.md)