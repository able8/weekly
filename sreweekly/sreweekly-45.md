## [SRE Weekly Issue #45](https://sreweekly.com/sre-weekly-issue-45/) - October 23, 2016
### Articles

1. [DNS devastation: Top websites whacked offline as Dyn dies again](http://www.theregister.co.uk/2016/10/21/dns_devastation_as_dyn_dies_under_denialofservice_attack/)

    The Register has a good overview of the attacks.
1. [Dyn Statement on 10/21/2016 DDoS Attack](http://hub.dyn.com/static/hub.dyn.com/dyn-blog/dyn-statement-on-10-21-2016-ddos-attack.html)

    Dyn released this statement on Saturday with more information on the outage. Anecdotally (based on my direct experience), I’m not sure their timeline is quite right, as it seems that I and others I know saw impact later than Dyn’s stated resolution time of 1pm US Eastern time. Dyn’s status post indicates resolution after 6pm Eastern, which matches more closely with what I saw.
1. [AWS](http://status.aws.amazon.com/)

    Among many other sites and services, AWS experienced an outage. They posted an unprecedented amount of detail during the incident in a banner on their status site. Their status page history doesn’t include the banner text, so I’ll quote it here:Nice job with the detailed information, AWS!
1. [DDoS on Dyn Impacts Twitter, Spotify, Reddit](https://krebsonsecurity.com/2016/10/ddos-on-dyn-impacts-twitter-spotify-reddit/)

    Krebs on Security notes that the attack came hours after a talk on DDoS attacks at NANOG. Krebs was a previous target of a massive DDoS, apparently in retaliation for his publishing of research on DDoS attacks.
1. [PagerDuty](https://www.pagerduty.com/blog/initial-outage-report/)

    Paging through PagerDuty was down or badly broken throughout Friday. Many pages didn’t come through, and those that did sometimes couldn’t be acknowledged. Some pages got stuck and PagerDuty’s system would repeatedly call engineers and leave the line silent. [source: personal experience] Linked is PagerDuty’s “Initial Outage Report”. It includes a preliminary explanation of what went wrong, an apology, and a pledge to publish two more posts: a detailed timeline and a root cause analysis with remediation items.
1. [How We Survived the Dyn DNS Outage](https://www.sumologic.com/blog-devops/survived-dyn-dns-outage/)

    Sumo Logic implemented dual nameserver providers after a previous DNS incident. This allowed them to survive the Dyn outage relatively unscathed.
1. [WikiLeaks supporters claim credit for massive U.S. cyberattack, but researchers skeptical](http://www.politico.com/story/2016/10/websites-down-possible-cyber-attack-230145)

    Anonymous and New World Hackers have claimed responsibility for Friday’s attack, but that claim may be suspect. Their supposed reasons were to “retaliate for the revocation of Julian Assange’s Internet access” and to “test their strength”.
1. [WikiLeaks urges supporters to ‘stop taking down the US internet’ following DDoS cyberattack](http://www.ibtimes.co.uk/wikileaks-urges-supporters-stop-taking-down-us-internet-following-massive-ddos-attack-1587718)

    WikiLeaks bought their claim and asked supporters to cut it out.
1. [Recursive and Iterative Queries](https://technet.microsoft.com/en-us/library/cc961401.aspx)

    A great basic overview of how recursive DNS queries work, published by Microsoft.
1. [Chapter 2 DNS Concepts, from Pro DNS and Bind by Ron Aitchison](http://www.zytrax.com/books/dns/ch2/)

    A much more detailed explanation if how DNS works. Recursive and iterative DNS queries are covered in more depth, along with AXFR/IXFR/NOTIFY which can be used to set up a redundant secondary DNS provider for your domain.
1. [Heroku Status Site Incident 965: Issues with DNS resolution](https://status.heroku.com/incidents/965)

    Heroku was among many sites impacted by the attack. Heroku’s status site was also impacted, prompting them to create a temporary mirror. This is painfully reminiscent of an article linked two issues ago about making sure your DNS is redundant.Full disclosure: Heroku is my employer.
1. [DOS Attacks and DNS: How to Stay Up If Your DNS Provider goes DOWN](http://blog.easydns.org/2010/08/19/dos-attacks-and-dns-how-to-stay-up-if-your-dns-provider-goes-down/)

    EasyDNS has this guide on setting up redundant DNS providers to safeguard against an attack like this. However, I’m not sure about their concept of “fast fluxing” your nameservers, that is, changing your nameserver delegation with your registrar when your main DNS provider is under attack.Unless I’m mistaken, a change in nameservers for a domain can take up to 2 days (for .com) to be seen by all end-users, because their ISP’s recursive resolver could have your domain’s NS records cached, and the TTL given by the .com nameservers is 2 days.Am I missing something? Maybe recursive resolvers optimistically re-resolve your domain’s NS records if all of your nameservers time out or return SERVFAIL? If you know more about this, please write to me!
### Outages

1. [PSN](http://411mania.com/games/playstation-network-goes-down-after-call-of-duty-infinite-warfare-beta-release/)
1. [The New York Times](http://mashable.com/2016/10/19/new-york-times-website-down/)
1. [Google Compute Engine Load Balancers](http://status.cloud.google.com/incident/compute/16020#5676325415157760)
    GCE’s LB offering returned progressively more 502s over the course of 2 hours, with a peak of 45% of requests failing. Linked is Google’s postmortem of the incident, which seems to have been a bad deploy.
1. [Virgin Australia website](http://australianaviation.com.au/2016/10/virgin-australia-website-back-up-and-running-after-sabre-outage/)

### [ << Prev ](sreweekly-44.md) ------------- [ Next >> ](sreweekly-46.md)