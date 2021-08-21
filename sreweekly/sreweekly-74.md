## [SRE Weekly Issue #74](https://sreweekly.com/sre-weekly-issue-74/) - May 28, 2017
### Articles

1. [The Always On Architecture – Moving Beyond Legacy Disaster Recovery](http://highscalability.com/blog/2016/8/23/the-always-on-architecture-moving-beyond-legacy-disaster-rec.html)

    The holy grail of high availability is a multi-datacenter (or cloud) active/active architecture. This article goes into why, including examples of common pitfalls of traditional disaster recovery solutions.
1. [Documenting an outage for a post-mortem review](http://serverfault.com/questions/29188/documenting-an-outage-for-a-post-mortem-review)

    Neat idea: here’s a Stack Overflow question asking for critique of a proposed outline for a post-incident analysis. It’s a great start already, and the answers include some pretty top-notch suggestions.
1. [Multi-region S3 failover /w Route53](https://read.iopipe.com/multi-region-s3-failover-w-route53-64ff2357aa30?__s=gcxkayouhzyr45m1hboa)

    A tutorial on setting up multi-region failover for an S3-hosted website, written in response to February’s major S3 outage in us-east.
1. [DNS Resolution in Go and Cgo](http://engineering.grab.com/dns-resolution-in-go-and-cgo)

    Last week, I linked to an article about debugging an overloaded ELB node. This week we have the sequel, a deep dive into the intricate details behind the problem, complete with a trip into the glibc source code.
1. [How Data Science Helps Power Worldwide Delivery of Netflix Content](https://medium.com/netflix-techblog/how-data-science-helps-power-worldwide-delivery-of-netflix-content-bac55800f9a7?source=rss----2615bd06b42e---4)

    Netflix uses data science to figure out how to fill the limited space on their edge content delivery nodes with the videos that people will request, all while (hopefully) avoiding hot nodes.
1. [Shadowing Customer Support for a Day](https://www.pagerduty.com/blog/shadowing-customer-support/)

    Zayna Shahzad, a PagerDuty software engineer, did customer support for a day, and she learned a ton. As SREs, we have the customer experience directly in our sights, so this kind of thing sounds like a really great idea.
1. [Is SRE a Good Term?](https://m.youtube.com/watch?feature=youtu.be&v=t2gX2mc-GmU)

    Charity Majors does not want to be an SRE. Find out why by watching this 5-minute video interview between her and Rob Hirschfeld. I don’t often link to videos, because who has time to watch stuff? But this one is pretty intriguing.
1. [How we do HumanOps at Server Density](http://feedproxy.google.com/~r/serverdensity/~3/EjvK8zk26lo/)

    Server Density originated the term “humanops”, and now they share 12 parts of how they practice it.
1. [Modifications to the current on-call system?](http://www.malaysiakini.com/letters/382996)

    A Malaysian doctor writes about how to ensure that the national health system’s on-call policy is safe for doctors.
1. [Five Things Tech Companies Can Do Better](https://www.susanjfowler.com/blog/2017/5/20/five-things-tech-companies-can-do-better)

    Do what better? Prevent and end illegal and unethical actions like discrimination, harassment, and retaliation. This article is by Susan Fowler, featured here a bunch, and while it’s not directly related to SRE, it’s so important that I urge you to read it.
### Outages

1. [Monitorama 2017 PDXsummary](http://monitorama.com/#outage)
    Monitorama (and a swathe of Portland) suffered a power outage last week. The organizers created a status site post (linked) and quickly organized a disaster recovery site: an entirely separate conference venue. Seriously amazing work, and oddly appropriate given the conference subject matter.
If you didn’t make it to Monitorama, here’s a summary from LinkedIn SRE Michael Kehoe.
1. [Sacramento Airport (CA, USA)](http://www.sacbee.com/news/local/transportation/article152306877.html)
1. [British Airways](https://www.neowin.net/news/british-airways-experiences-global-system-outage-resulting-in-significant-delays-to-flights)

### [ << Prev ](sreweekly-73.md) ------------- [ Next >> ](sreweekly-75.md)