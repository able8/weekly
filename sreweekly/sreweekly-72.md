## [SRE Weekly Issue #72](https://sreweekly.com/sre-weekly-issue-72/) - May 14, 2017
### Articles

1. [Designing robust and predictable APIs with idempotency](https://stripe.com/blog/idempotency)

    Idempotence is a critically important tool in building a reliable system. Stripe explains the concept and shows how they wrap theoretically non-idempotent actions like charging a credit card into safely idempotent API calls.
1. [What’s not Actionable & Business Critical Shouldn’t Ring: Building the Right Alerting System](https://thoughts.t37.net/whats-not-actionable-business-critical-shouldn-t-ring-building-the-right-alerting-system-e8f4b085a2cb?__s=bwykwk1kcceogszq8abt)

    Here’s an account of an effort to move from server-based paging (this server is down) to functional-based alerting (this user action isn’t working), with a resulting impressive reduction in out-of-hours paging.
1. [CPU Utilization is Wrong](http://www.brendangregg.com/blog/2017-05-09/cpu-utilization-is-wrong.html)

    It pays to study up and deeply understand what a simple metric like “cpu utilization” really means.
1. [AWS Service Health Dashboard](http://status.aws.amazon.com/)

    Why am I linking to AWS’s status site? Look closely, and you’ll see that the “green checkmark i” symbol has been replaced with a far more noticeable blue circle with a white diamond. Check out the old icon here for comparison. End of an era, or just another way of presenting the same information?
1. [Circuit breaker and monitoring of a gRPC service in Ruby (Part 1)](https://medium.com/@shiladitya16/circuit-breaker-and-monitoring-of-a-grpc-service-in-ruby-part-1-7509c7e1356a)

    The author introduces a new Ruby gem, grpc-commons that makes it easy to add circuit breaker and statsd support to a grpc client.
1. [Introducing distributed tracing in your Python application via Zipkin](http://t.dripemail2.com/c/eyJhY2NvdW50X2lkIjoiMjY1Njc0MyIsImRlbGl2ZXJ5X2lkIjoiODE3NjA4NzUzIiwidXJsIjoiaHR0cDovL2VjaG9yYW5kLm1lL2ludHJvZHVjaW5nLWRpc3RyaWJ1dGVkLXRyYWNpbmctaW4teW91ci1weXRob24tYXBwbGljYXRpb24tdmlhLXppcGtpbi5odG1sP19fcz1id3lrd2sxa2NjZW9nc3pxOGFidCJ9)

    Along with being a tutorial on setting up Zipkin with Python, this article also explains some basic Zipkin concepts.
1. [Announcing the Modern Incident Resolution Lifecycle](http://t.dripemail2.com/c/eyJhY2NvdW50X2lkIjoiMjY1Njc0MyIsImRlbGl2ZXJ5X2lkIjoiODE3NjA4NzUzIiwidXJsIjoiaHR0cHM6Ly93d3cucGFnZXJkdXR5LmNvbS9ibG9nL21vZGVybi1pbmNpZGVudC1yZXNvbHV0aW9uLWxpZmVjeWNsZS8_X19zPWJ3eWt3azFrY2Nlb2dzenE4YWJ0In0)

    PagerDuty is apparently trying to position itself as more than just a paging service, with a few new features around the entire incident lifecycle. I’m especially interested in checking out the new postmortem tooling.
1. [How we Upgraded a 22TB MySQL Cluster from 5.6 to 5.7 (in 9 months)](https://thoughts.t37.net/how-we-upgraded-a-22tb-mysql-cluster-from-5-6-to-5-7-in-9-months-cc41b391895d)

    I included this article last week, but my link was outdated and returned a 404. Here’s the corrected link — sorry about that!
1. [A first look at Elastic’s new Machine Learning Technology](https://www.linkedin.com/pulse/first-look-elastics-new-machine-learning-technology-robert-cowart)

    I put a call out for a review of Elastic’s new beta anomaly detection feature last week, and here one is! Thanks to an Elastic employee for forwarding this link to me.
1. [IT Outages, Who’s Really at Fault?](http://www.informationweek.com/strategic-cio/executive-insights-and-innovation/it-outages-whos-really-at-fault/a/d-id/1328869?_mc=RSS_IWK_EDT)

    This article cautions one to be careful to look past an obvious root cause, because a deeper systemic or policy problem may be lurking behind it.
1. [Watch out for serverless computing’s blind spot](http://www.infoworld.com/article/3196133/cloud-computing/watch-out-for-serverless-computings-blind-spot.html)

    Serverless / FaaS abstract away traditional provisioning, and they make it really easy to ignore planning for resource usage.
1. [Safety Moment – Are Accidents a Failure of Imagination? | PreAccident Investigation Podcast](http://preaccidentpodcast.podbean.com/e/safety-moment-are-accidents-a-failure-of-imagination/)

    Wow, what a concept:This 2.5-minute podcast from Todd Conklin has a really great question: to achieve reliability, do we have to try to imagine in advance all of the possible ways our systems could fail?
1. [“The Scariest Moment of My Life” – BWH Safety Matters](https://bwhsafetymatters.org/2017/05/09/the-scariest-moment-of-my-life/)

    A patient was given an incorrect syringe resulting in a 5x insulin overdose. Brigham and Women’s Hospital reports on the accident and what they’re doing to prevent mistakes of this sort in the future.
1. [PagerDuty’s 2017 State of Digital Operations Report](https://www.pagerduty.com/resources/reports/digital-operations/)

    
### Outages

1. [Our First Kubernetes Outage – Saltside Engineering](https://engineering.saltside.se/our-first-kubernetes-outage-c6b9249cfd3a)
    Kudos to the Saltside folks for sharing a public postmortem for an internal, non-customer-impacting outage!
This is public postmortem for an a complete shutdown of our internal Kubernetes cluster. It’s shared with you all so everyone may learn.
1. [“Re-experience the fun of customizing your Place Page!” A Tale of Oops from Ops](https://community.secondlife.com/blogs/entry/2167-re-experience-the-fun-of-customizing-your-place-page-a-tale-of-oops-from-ops/)
    Ouch. Linden Lab’s ops team discovered the hard way that they didn’t have a working backup copy of some customer data. The best part of this article is the discussion of the “Shrek Ears” tradition at Linden. It’s one of the things I remember most fondly from my time there, and having worn the ears a few times in my day, I can attest to the fact that it’s a great way to handle the psychological impact of making a mistake.
1. [Chase (bank)](https://www.americanbanker.com/news/chase-outage-knocks-out-payments-services-for-customers-nationwide)
1. [Facebook](https://techcrunch.com/2017/05/08/facebook-is-down-in-asia-pacific-and-now-parts-of-north-america-too/)

### [ << Prev ](sreweekly-71.md) ------------- [ Next >> ](sreweekly-73.md)