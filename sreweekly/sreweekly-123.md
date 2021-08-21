## [SRE Weekly Issue #123](https://sreweekly.com/sre-weekly-issue-123/) - May 27, 2018
### Articles

1. [LinkedOut: A Request-Level Failure Injection Framework](https://engineering.linkedin.com/blog/2018/05/linkedout--a-request-level-failure-injection-framework)

    The system is highly configurable, allowing fine-grained A/B testing of failures at all levels of the microservice call tree.
1. [Ephemeral port exhaustion and how to avoid it](https://making.pusher.com/ephemeral-port-exhaustion-and-how-to-avoid-it/)

    Ephemeral port exhaustion can really ruin your day. Read this to learn how to deal with it, how to detect it before you have problems, and why you might run out of ports sooner than you expect.Will Sewell — Pusher
1. [Incident Management at Spotify](https://labs.spotify.com/2013/06/04/incident-management-at-spotify/)

    This incident report from 2013 is a great read. It’s really two inches in one, including an analysis of why a remediation task from the first wasn’t completed in time to prevent the second.David Poblador i Garcia — Spotify
1. [Charity Majors on Observability and Understanding the Operational Ramifications of a System](https://www.infoq.com/articles/charity-majors-observability-failure)

    There are a few nice tidbits in this interview, including this one:Daniel Bryant – InfoQ
1. [Canary Deployment: What Is It and How Can I Use It?](https://rollout.io/blog/canary-deployment/)

    This article has introduction to implementing canary deployment and also includes a discussion of the potential downsides.Erik [surname not given] — Rollout.io
1. [Open-sourcing Katran, a scalable network load balancer](https://code.facebook.com/posts/1906146702752923/open-sourcing-katran-a-scalable-network-load-balancer/)

    Lots of great detail in this announcement, including an analysis of how (and why) they designed their load balancer to function entirely in userspace without a kernel bypass mechanism.Nikita Shirokov and Ranjeeth Dasineni — Facebook
1. [Building low-overhead metrics collection for high-performance systems](https://blog.wallaroolabs.com/2018/02/building-low-overhead-metrics-collection-for-high-performance-systems/)

    Metrics are great, right? Except sometimes they’re not, when the metric collection itself adds enough load to impair the system.Jonathan Brown — Wallaroo
### Outages

1. [Google BigQuery](https://status.cloud.google.com/incident/bigquery/18036#18036003)
    Click through for the full incident report.
Configuration changes being rolled out on the evening of the incident were not applied in the intended order.
1. [GCP Networking in us-east4](https://status.cloud.google.com/incident/cloud-networking/18009#18009005)
    Here’s some detail on the BGP issue that took down us-east4 last week.
1. [Google StackDriver](https://status.cloud.google.com/incident/google-stackdriver/18007#18007005)
    It’s a hat trick of three GCP incident followup reports. Happy reading!
1. [Slack](https://status.slack.com/2018-05/42acc9f7860c0416)
1. [Bank of New Zealand](https://www.tvnz.co.nz/one-news/new-zealand/bnz-cards-and-atms-now-working-they-should-after-major-outage-leaves-customers-annoyed)
1. [Twitter](https://www.thesun.co.uk/news/6386866/twitter-down-for-users-around-the-world-as-thousands-report-social-network-crashing/)
1. [National Australia Bank](https://www.perthnow.com.au/technology/telstra/telstra-cctv-blunder-outage-caused-serious-privacy-breach-ng-b88847167z)
    This outage is particularly notable because the bank has stated their intention to compensate customers for their losses, such as estimated lost revenues from inability to complete sales transactions.

### [ << Prev ](sreweekly-122.md) ------------- [ Next >> ](sreweekly-124.md)