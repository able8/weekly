## [SRE Weekly Issue #34](https://sreweekly.com/sre-weekly-issue-34/) - August 7, 2016
### Articles

1. [Netflix: Introducing Winston – Event driven Diagnostic and Remediation Platform](http://techblog.netflix.com/2016/08/introducing-winston-event-driven.html)

    Winston is Netflix’s tool for runbook automation, based on the open source StackStorm.  Winston helps reduce pager burden by filtering out false-positive alerts, collecting information for human responders, and remediating some issues automatically.
1. [High Tempo, High Consequence](http://www.kitchensoap.com/2014/03/13/high-tempo-high-consequence/)

    Is it valid for those working on non-life-critical systems to try to draw on lessons learned in safety-critical fields like surgery and air traffic control?  John Allspaw, citing Dr. Richard Cook, answers with an emphatic yes.  
1. [Expired credit cards can shut down businesses](https://blog.serverdensity.com/how-an-expired-credit-card-can-shut-down-your-entire-business/)

    The best HA infrastructure design in the world won’t save you when your credit card on file expires.
1. [Why Uber Engineering Switched from Postgres to MySQL](https://eng.uber.com/mysql-migration/)

    There’s a huge amount of detail on both PostgreSQL and MySQL in this article, including some sneaky edge-case pitfalls that prompted Uber to look for a new database.
1. [How to Setup a Highly Available Multi-AZ Cassandra Cluster on AWS EC2](http://feedproxy.google.com/~r/HighScalability/~3/1-yg4hGo4Ow/how-to-setup-a-highly-available-multi-az-cassandra-cluster-o.html)

    This article goes into a good amount of depth on setting up a Caassandra cluster to survive a full AZ outage.
1. [Officials try to explain 911 outage that should ‘never happen’](http://www.cortezjournal.com/article/20160802/AP/308029754/Officials-try-to-explain-911-outage-that-should-'never-happen')

    When a Maryland, US county’s emergency services went offline for two hours, 100 calls were missed, possibly contributing to two deaths.  In the vein of last week’s theme of complex failures:

1. [Netflix Billing Migration to AWS – Part III](http://techblog.netflix.com/2016/08/netflix-billing-migration-to-aws-part.html)

    Here’s the third (final?) installment in this series.  This one has some fascinating details on a topic near and dear to my heart: live migration of a database.  Their use of DRBD and synchronous replication is especially intriguing.
1. [AMA DevOps & SRE](http://pages.catchpoint.com/DEVOPS-SRE-AMA-Registration.html)

    Ooh, this is gonna be fun.  Catchpoint and O’Reilly are hosting an AMA (Ask Me Anything) with DevOps and SRE folks, including Liz Fong-Jones and Charity Majors, both of whose articles have been featured here previously.  The questions posted so far look pretty great.
### Outages

1. [EE (UK telecom)](http://www.wired.co.uk/article/ee-roaming-network-down)
    EE users saw a 2-day outage when roaming.
1. [Vocus (AU telecom)](http://www.itnews.com.au/news/vocus-network-outage-downs-internet-voice-services-432164)
1. [PlayStation Network](http://www.express.co.uk/entertainment/gaming/696455/PSN-down-PlayStation-Network-PS4-Sony-offline)
1. [Battle.net](http://www.forbes.com/sites/erikkain/2016/08/02/overwatch-servers-still-down-as-blizzard-suffers-ongoing-ddos-attack/)
1. [123-Reg (UK web host)](http://www.theregister.co.uk/2016/08/02/123_reg_goes_titsup_again/)
1. [Airtel](http://timesofindia.indiatimes.com/business/india-business/Outage-in-Airtels-broadband-network/articleshow/53495432.cms)
1. [Commonwealth Bank](http://www.lifehacker.com.au/2016/08/commonwealth-bank-services-are-down-around-the-country/)
1. [OGS (Online Go Server)](https://forums.online-go.com/t/please-excuse-our-mess/8614)
1. [Neotel](http://mybroadband.co.za/news/wireless/174890-neotel-users-suffering-with-slow-data-speeds-network-downtime.html)

### [ << Prev ](sreweekly-33.md) ------------- [ Next >> ](sreweekly-35.md)