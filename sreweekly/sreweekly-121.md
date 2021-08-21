## [SRE Weekly Issue #121](https://sreweekly.com/sre-weekly-issue-121/) - May 13, 2018
### Articles

1. [Defining SLOs for services with dependencies](https://cloudplatform.googleblog.com/2018/05/Defining-SLOs-for-services-with-dependencies-CRE-life-lessons.html)

    This latest in the CRE Life Lessons series takes on dependencies and how they impact a service’s SLO in obvious and subtle ways.Robert van Gent — Google
1. [How a SaaS provider made microservices deployment safely chaotic](https://searchmicroservices.techtarget.com/feature/How-a-SaaS-provider-made-microservices-deployment-safely-chaotic)

    This company discovered that the benefits of microservices came with some significant downsides. Here’s how they turned to chaos testing to improve reliability.Meredith Courtemanche — TechTaret
1. [The benefits of chaos engineering-as-a-service](https://jaxenter.com/chaos-engineering-service-144113.html)

    Keeping in mind that this is written by the CTO of Gremlin, it contains some good points about buying versus building your chaos engineering system. It would apply to other chaos engineering services too — if there were any.Matt Fornaciari — Gremlin, Inc.
1. [Zero Downtime Updates with HashiCorp Terraform](https://www.hashicorp.com/blog/zero-downtime-updates-with-terraform)

    Even as an experienced Terraform user, I learned about some Terraform features I hadn’t been aware of.Nic Jackson — Hashicorp
1. [How Your Systems Keep Running Day After Day](https://itrevolution.com/john-allspaw-how-your-systems-keep-running-day-after-day/)

    In issue #98, I linked to a recording of John Allspaw’s DOES17 talk. In case you didn’t have time to listen, here’s a transcript. If you didn’t have time to read the Stella Report, I highly recommend reading this as an intro to the major concepts therein.John Allspaw
### Outages

1. [Fastly](https://status.fastly.com/incidents/68rr588xcqrr?utm_source=slack)
    Full disclosure: Fastly is my employer.
1. [Travis CI](https://www.traviscistatus.com/incidents/0ch1w22m3lwc)
1. [Python Package Index (PyPI)](https://status.python.org/incidents/zjxwp5lpyn8s)
1. [Honeycomb](https://honeycomb.io/blog/2018/05/rds-performance-degradation---postmortem/)
    Wow, I just love Honeycomb’s post-incident analyses, and this one is no exception. Highly recommend.
Rule of thumb as a developer: it’s probably not the database, it’s probably your code.
Turns out that it was, in this case!
Andy Isaacson
1. [Hulu](https://www.dailydot.com/debug/hulu-down-outage/)
1. [Instagram](https://nationalwire.com.ng/instagram-is-down-across-the-world/)
1. [MasterCard](http://www.kctv5.com/story/38162026/mastercard-outage-angers-customers-nationwide)

### [ << Prev ](sreweekly-120.md) ------------- [ Next >> ](sreweekly-122.md)