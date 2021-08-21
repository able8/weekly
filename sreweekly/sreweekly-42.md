## [SRE Weekly Issue #42](https://sreweekly.com/sre-weekly-issue-42/) - October 2, 2016
### Articles

1. [Making the Netflix API More Resilient](http://techblog.netflix.com/2011/12/making-netflix-api-more-resilient.html)

    Netflix’s API has an advanced circuit-breaker system including a defined automated fallback plan for every dependency.
1. [Just Culture](http://sidneydekker.com/just-culture/)

    This is Sydney Dekker’s course on Just Culture, including a full explanation of Restorative Just Culture. I especially like the concept of Second Victims of incidents: the practitioner (e.g. engineer) that was directly involved in the incident.
1. [TCP Puzzlers](https://www.joyent.com/blog/tcp-puzzlers)

    Think you know how TCP works? There are sneaky edge-cases that can cause an outage if you don’t know about them. Example: a MySQL replicating slave will happily report “0 seconds behind master” indefinitely while waiting on a connection to the master that’s long-since silently failed.
1. [API First Transformation at Etsy – Operations](https://codeascraft.com/2016/09/26/api-first-transformation-at-etsy-operations/)

    Etsy shares the operational issues they encountered as they moved toward an API/microservice architecture. I especially like the detail about limiting concurrent in-flight sub-requests per root request across the entire request tree.
1. [Site Availability is for Everybody](http://confreaks.tv/videos/railsconf2016-site-availability-is-for-everybody)

    My co-worker at Heroku, Stella Cotton, gave this rockin’ keynote at RailsConf 2016. She covers load testing and performance bottleneck diagnosis, and most of what she says applies not just to Rails.
1. [How Uber Manages a Million Writes Per Second Using Mesos and Cassandra Across Multiple Datacenters](http://feedproxy.google.com/~r/HighScalability/~3/WmFXI6yf9TY/how-uber-manages-a-million-writes-per-second-using-mesos-and.html)

    Here’s a summary of a talk about Uber’s system that stores live location data of riders and drivers. They run Cassandra in containers managed by Mesos.
1. [Building a Scalable Minimum Viable Product](https://www.infoq.com/news/2016/09/scalable-MVP)

    With an MVP, you’re just trying to get into the market and test the waters as quickly as possible, so there’s a temptation to leave considerations like scalability for later. But what if your MVP is unexpectedly successful?
1. [Systems We Love](http://dtrace.org/blogs/bmc/2016/09/30/submitting-to-systems-we-love/)

    Systems We Love is a new conference modeled after the popular Papers We Love. It looks really interesting, and they’re saying they already have a lot of great proposals.
1. [Travis CI: The day we deleted our VM images](http://blog.travis-ci.com/2016-09-30-the-day-we-deleted-our-vm-images/)

    Travis CI shares more about a major outage last month.
1. [So You’ve Been Paged: A Guide to Incident Response (For Those Who Hate Being Paged)](http://blog.scalyr.com/2016/09/so-youve-been-paged/)

    A nice incident response primer from Scalyr.
### Outages

1. [Vodafone](http://www.ibtimes.com.au/vodafone-offers-extra-2gb-data-customers-after-vodafail-network-outage-1529812)
    They’re offering a 2GB data credit.
1. [Google Cloud Pub/Sub](http://status.cloud.google.com/incident/cloud-pubsub/16002#5717271485874176)
1. [Yahoo Mail](http://www.ibtimes.co.uk/yahoo-mail-goes-down-parts-europe-sending-twitter-into-frenzy-1583657)
1. [Netflix](http://www.upi.com/Entertainment_News/2016/10/01/Netflix-outage-leaves-couch-sitters-stranded-on-Saturday/2211475359902/)
1. [Newsweek](http://www.theregister.co.uk/2016/09/30/criticizing_donald_trump_will_get_you_ddosed_off_the_internet/)
    The outage occurred after they posted an article critical of trump, perhaps in retaliation, possibly by Russia. Allegedly.

### [ << Prev ](sreweekly-41.md) ------------- [ Next >> ](sreweekly-43.md)