## [SRE Weekly Issue #141](https://sreweekly.com/sre-weekly-issue-141/) - September 30, 2018
### Articles

1. [Rethinking Netflix’s Edge Load Balancing](https://medium.com/netflix-techblog/netflix-edge-load-balancing-695308b5548c?source=rss----2615bd06b42e---4)

    An outline of the design of Netflix’s new load balancer, with special emphasis on dealing with faltering backends. Great idea: servers report their utilization level in a response header. Tricky pitfall: the LB is so good at moving requests off of ailing backends that backend failure rate alerts don’t fire.Mike Smith — Netflix
1. [NewSQL database systems are failing to guarantee consistency, and I blame Spanner](http://dbmsmusings.blogspot.com/2018/09/newsql-database-systems-are-failing-to.html)

    This article begins by explaining consistency versus availability in distributed data stores and argues that the trade-off is less significant than one might think. Then it describes a pitfall seen in some new data stores. I’ve pondered aloud here in the past on how Spanner can make the claims it does, and this article explains that nicely.Daniel Abadi
1. [The redux of the fallacies of distributed computing](https://ayende.com/blog/184545-A/the-redux-of-the-fallacies-of-distributed-computing?Key=ca5e02d1-378a-4189-b231-85b509071c44)

    And here’s a refutation of part of the previous article by the creator of RavenDB.Ayende Rahien
1. [Getting The Airlines Back On Their Feet After A Disaster](https://www.informationsecuritybuzz.com/articles/getting-the-airlines-back-on-their-feet-after-a-disaster/)

    Dr. Sandra Bell
1. [Upgrading GitHub from Rails 3.2 to 5.2](https://githubengineering.com/upgrading-github-from-rails-3-2-to-5-2/)

    GitHub used an innovative technique to avoid holding open a long-running code branch while upgrading their application to support rails 5.2.Eileen Uchitelle — GitHub
1. [Travis CI: Build VMs boot failure on the sudo-enabled infrastructure: incident postmortem](http://blog.travis-ci.com/2018-09-27-gce-worker-postmortem)

    Worker node errors led to cascading failure when they hit Google Compute Engine quotas.Bogdana Vereha — Travis CI
1. [Secret IBM script could have prevented 11-hour US tax day outage](https://www.theregister.co.uk/2018/09/25/secret_ibm_script_could_have_prevented_11_hour_us_tax_day_outage/)

    This week, the US Internal Revenue Service (IRS) issued a report analyzing the tax-day outage that occurred this past April. Linked is a nice summary by the Register.Thanks to reader Michael Fischer for a tip on the report.Chris Mellor — The Register
### Outages

1. [Facebook](https://www.rappler.com/technology/news/213055-facebook-suffers-apparent-downtime-forced-logouts-september-28-2018)
1. [Amazon Alexa](https://internetofbusiness.com/iot-security-why-alexas-international-outage-is-a-blow-for-amazon/)
1. [Delta Airlines](http://www.mysuncoast.com/news/national/service-resumed-after-delta-experienced-network-outage/article_87f8cfb3-d842-5be5-bcd4-d36bd82511e3.html)
1. [Honeywell (smart thermostat manufacturer)](https://www.digitaltrends.com/home/honeywell-server-outage-left-users-unable-control-devices/)
1. [Zoho](https://www.zdnet.com/article/domain-registrar-oversteps-taking-down-zoho-domain-impacts-over-30mil-users/)
    SaaS provider Zoho’s domain registration was revoked by its registrar after a run-of-the-mill phishing complaint, affecting 30 million users.
1. [Steemit](https://xbt.net/blog/steemit-platform-fully-operational-following-platform-outage/)

### [ << Prev ](sreweekly-140.md) ------------- [ Next >> ](sreweekly-142.md)