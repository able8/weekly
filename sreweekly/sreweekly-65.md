## [SRE Weekly Issue #65](https://sreweekly.com/sre-weekly-issue-65/) - March 26, 2017
### Articles

1. [Monitoring Weekly Newsletter](http://weekly.monitoring.love)

    Look, a new newsletter about monitoring! I’m really excited to see what they have to offer.
1. [Last Week in AWS](https://lastweekinaws.com/)

    And another new newsletter! Like Monitoring Weekly, I’m betting this one will have a lot of articles of interest to SREs.
1. [Sandstorm or Significant: The evolving role of context in incident management](https://devops.com/webinar-sandstorm-significant-evolving-role-context-incident-management/)

    VictorOps held a webinar last Thursday to present and discuss the concept of context in incident management. Just paging in a responder isn’t enough: we need to get them up to speed on the incident as soon as possible. Ideally, the page itself would include snapshots of relevant graphs, links to playbooks, etc. But if we’re not careful and add too much information, the responder is overloaded by a “sandstorm” of irrelevant data. “time to learn” — post incident learning careful of info overload in presenting context with pagesThis webinar was created by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Embracing Risk](https://medium.com/@jerub/embracing-risk-74bd876a64da?source=rss-3ea3ed468e7c------2)

    Here’s the next in Stephen Thorne’s series of commentary on chapters of the SRE book. I like that Google makes an effort not to be too reliable for fear of setting expectations too high, and they’re also realistic in their availability goals: no end-user will notice a 20-second outage.
1. [Infinity Is a Bad Timeout](https://dzone.com/articles/infinity-is-a-bad-timeout)

    Writing an API, system, server or really anything people might make use of? Don’t make the default timeout be infinite.
1. [5 Incident Management Tools You Need During a Firefight](https://www.pagerduty.com/blog/5-incident-management-tools/)

    PagerDuty really has been churning out excellent articles in the past couple of weeks. [Spoiler Alert] The five things are: internal communication, monitoring, a public status site, a support ticket system, and a defined incident response procedure.
1. [Avoiding Incident Response Bottlenecks](https://www.pagerduty.com/blog/avoiding-incident-response-bottlenecks/)

    Keep on rockin’ it, PagerDuty. This time they identify common problems that hinder incident response and give suggestions on how to fix them.
1. [SREcon17: Brave new world of site reliability engineering](https://blog.netsil.com/srecon17-brave-new-world-of-site-reliability-engineering-bf69cc6cd2f7)

    The author reviews their experience at SRECon17 Americas, including interesting bits on Julia Evans, training, recruiting, and diversity.
1. [Human Error? No Problem](https://www.cannabisindustryjournal.com/column/human-error-no-problem/)

    I love that the ideas we’re talking about regarding human error apply even to commercial cannabis growing.
1. [Sometimes Boring is Better — Production Ready](https://medium.com/production-ready/sometimes-boring-is-better-d16d38214186)

    The newer and shinier the technology, the bigger the production risk.
### Outages

1. [Kings College London storage system outage and data loss](https://regmedia.co.uk/2017/02/23/kcl_external_review.pdf)
    Kings College London’s HP storage system suffered a routine failure that, due to a firmware bug, resulted in loss of the entire array. Linked is an incredibly detailed PDF including multiple contributing factors and many remediations. Example: primary backups were to another folder on the same storage system, and secondary tape backups were purposefully incomplete.
1. [Ryanair](http://www.dailystar.co.uk/travel-news-cheap-uk-holidays-luxury-breaks-more-daily-star/599011/Ryanair-down-website-crashed-not-working-flight-sale)
    This one’s interesting to me because it seems to have been self-inflicted due to a flash sale.
1. [Apple Store](https://9to5mac.com/2017/03/24/apple-store-goes-down-red-iphone-and-329-ipad-on-sale-today/)
    Another (possibly) self-inflicted outage due to a sale.
1. [Microsoft Azure](http://www.silicon.co.uk/cloud/microsoft-azure-outage-207503)
1. [Discord Status – Connectivity Issuesthundering herd](https://status.discordapp.com/incidents/dj3l6lw926kl)
    Finally, my search alert for “thundering herd” paid off! I hadn’t heard of Discord before now, but they sure do write a great postmortem. Did you know that the thundering herd is a sports team?

### [ << Prev ](sreweekly-64.md) ------------- [ Next >> ](sreweekly-66.md)