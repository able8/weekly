## [SRE Weekly Issue #236](https://sreweekly.com/sre-weekly-issue-236/) - September 20, 2020
### Articles

1. [My first outage](https://mads-hartmann.com/2019/03/01/my-first-outage.html)

    A nice juicy post-incident report from the archives. Remember the first time you took down production?Mads Hartmann — Glitch
1. [Fault during testing of NordLink](https://www.statnett.no/en/about-statnett/news-and-press-releases/news-archive-2020/fault-during-testing-of-nordlink/)

    While testing a new power transmission link, it was accidentally overloaded by a factor of ~14x, with far-reaching but ultimately well-managed effects.Thanks to Jesper Lundkvist for this one.
1. [Throughput autoscaling: Dynamic sizing for Facebook.com](https://engineering.fb.com/networking-traffic/throughput-autoscaling/)

    As Facebook moved from a static to an auto-scaled web pool, they had to try to predict their expected demand as accurately as possible.Daniel Boeve, Kiryong Ha, and Anca Agape — Facebook
1. [Database migrations lessons learned](https://octopus.com/blog/database-migrations-lessons-learned)

    The key lesson involves ensuring that your migrations avoid using parts of the production code, which could cause their action to change down the line inadvertently.Frank Lin — Octopus Deploy
1. [Moobot vs. Gatebot: Cloudflare Automatically Blocks Botnet DDoS Attack Topping At 654 Gbps](https://blog.cloudflare.com/moobot-vs-gatebot-cloudflare-automatically-blocks-botnet-ddos-attack-topping-at-654-gbps/)

    Cloudflare uses an interesting multi-layered approach to mitigating attacks.Omer Yoachimik — Cloudflare
1. [Availability, Maintainability, Reliability: What’s the Difference?](https://www.blameless.com/blog/availability-maintainability-reliability-whats-the-difference)

    The availability/reliability distinction in this article is thought-provoking.Emily Arnott — Blameless
1. [Troubled Times: Episode 3](https://www.adaptivecapacitylabs.com/blog/2020/09/19/troubled-times-episode-3/)

    This article (not a video or podcast despite the name) also focuses on the increasing importance of learning from incidents.Dr. Richard Cook — Adaptice Capacity Labs
1. [Building and revising adaptive capacity sharing for technical incident response: A case of resilience engineering](https://www.sciencedirect.com/science/article/pii/S0003687020301903)

    This academic paper uses a case study to show how a company engineered the resilience of their system in response to a series of incidents.Richard I. Cook and Beth Adele Long — Applied Ergonomics
### Outages

1. [Google Drive](https://static.googleusercontent.com/media/www.google.com/en//appsstatus/ir/edgugm6kpj2ewx7.pdf)
    This is a post-analysis for two outages, one from this past week and the other from the week before.
1. [Instagram](https://www.thestar.com.my/tech/tech-news/2020/09/18/instagram-down-users-experience-brief-outages-technical-problems)
1. [Facebook](https://www.thedenverchannel.com/news/national/facebook-users-report-widespread-outages)
1. [Discord](https://discord.statuspage.io/incidents/gfqnsk0l036s)
1. [Fastly](https://status.fastly.com/incidents/jpxtngv8hxgs)
1. [Gandi](https://news.gandi.net/en/2020/09/postmortem-regarding-the-network-incident-from-september-15th-2020-on-iaas-and-paas-fr-sd3-fr-sd5-and-fr-sd6/)
    Postmortem regarding the Network Incident from September 15, 2020 on IAAS and PAAS FR-SD3, FR-SD5, and FR-SD6
A layer 2 network loop was accidentally introduced, on two separate occasions.
Sébastien Dupas — Gandi
1. [Azure](https://status.azure.com/en-us/status/history/)
    This was an outage on Sept. 14 in the UK South region.  A cooling system was shut off in error during a maintenance procedure.

### [ << Prev ](sreweekly-235.md) ------------- [ Next >> ](sreweekly-237.md)