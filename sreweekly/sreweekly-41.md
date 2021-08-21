## [SRE Weekly Issue #41](https://sreweekly.com/sre-weekly-issue-41/) - September 25, 2016
### Articles

1. [Trestus by canonical-ols](https://canonical-ols.github.io/trestus/)

    Trestus is a new tool to generate a status page from a Trello board. Neat idea!
1. [Writing Your First Postmortem](https://medium.com/production-ready/writing-your-first-postmortem-8053c678b90f#.mw65xj6of)

    An excellent intro to writing post-incident analysis documents is the subject of this issue of Production Ready by Mathias Lafeldt. I can’t wait for the sequel in which he’ll address root causes.
1. [The Morning Paper on Operability](https://blog.acolyer.org/2016/09/21/the-morning-paper-on-operability/)

    Adrian Colyer of The Morning Paper gave a talk at Operability.IO with a round-up of his favorite write-ups of operations-related papers. I really love the fascinating trend of “I have no idea what I’m doing” — tools that help us infer interconnections, causality, and root causes in our increasingly complex infrastructures. Rather than try (and in my experience, usually fail) to document our massively complicated infrastructures in the face of increasing employee turnover rates, let’s just accept that this is impossible and write tools to help us understand our systems.
1. [Tweets](https://twitter.com/sreweekly)

    And for fun, a couple of amusing tweets I came across this week:— Alice Goldfuss (@alicegoldfuss) (check out her awesome talk at SRECon16 about the Incident Command System)— Honest Status Page (@honest_update)— Senior Oops Engineer (@ReinH)
1. [Zuul 2: The Netflix Journey to Asynchronous, Non-Blocking Systems](http://techblog.netflix.com/2016/09/zuul-2-netflix-journey-to-asynchronous.html)

    Netflix documents the new version of their frontend gateway system, Zuul 2. They moved from blocking IO to async, which allows them to handle persistent connections from clients and better withstand retry storms and other spikes.
1. [Uber hits bumps in the road with microservices challenges](http://searchitoperations.techtarget.com/news/450304823/Uber-hits-bumps-in-the-road-with-microservices-challenges)

    In last week’s issue, I linked to a chapter from Susan Fowler’s upcoming book on microservices. Here’s an article summarizing her recent talk at Velocity about the same subject: how to make microservices operable. She should know: Uber runs over 1300 microservices. Also summarized is her fellow SRE Tom Croucher’s keynote talk about outages at Uber.
1. [Introducing the GitHub Load Balancer](http://githubengineering.com/introducing-glb/)

    In this first of a series, GitHub lays out the design of their new load balancing solution. It’s pretty interesting due to a key constraint: git clones of huge repositories can’t resume if the connection is dropped, so they need to avoid losing connections whenever possible.
1. [Book Review: Site Reliability Engineering – How Google Runs Production Systems](https://www.infoq.com/articles/site-reliability-engineering)

    I’m embarrassed to say that I haven’t yet found the time to take my copy of the SRE book from its resting place on my shelf, but here’s another review with a good amount of detail on the highlights of the book.
1. [TCP connection repair](https://lwn.net/Articles/495304/)

    Live migration of VMs while maintaining TCP connections makes sense — the guest’s kernel holds all the connection state. But how about live migrating  containers? The answer is a Linux feature called TCP connection repair.
1. [SSP accused of making ‘wrong call’ over decision not to use secondary data centre after outage](http://www.postonline.co.uk/post/news/2469616/ssp-accused-of-making-wrong-call-after-decision-not-to-use-secondary-data-centre-after-outage)

    The SSP story (linked here  two issues ago) is getting even more interesting. They apparently decided not to switch to their secondary datacenter in order to avoid losing up to fifteen minutes’ worth of data, instead taking a week+ outage.
1. [Learning From UCLA](http://cen.acs.org/articles/87/i31/Learning-UCLA.html)

    While, in SRE, we generally don’t have to worry about our deploys literally blowing up in our faces and killing us, I find it valuable to look to other fields to learn from how they manage risk. Here’s an article about a tragic accident at UCLA in which a chemistry graduate student was severely injured and later died. A PhD chemist I know mentioned to me that the culture of safety in academia is much less rigorous than in the industry, perhaps due in part to a differing regulatory environment.
### Outages

1. [Destiny (game)](http://www.idigitaltimes.com/destiny-servers-down-error-code-tapir-stifles-rise-iron-release-date-excitement-557515)
1. [Pokemon GOpoke Mongo](http://www.itechpost.com/articles/31694/20160921/pokemon-go-pokemon-go-is-dead-niantic-pokemon-go-niantic-niantic-remains-mum-pokestops-gyms-wont-load-pokemon-go-pokestops-pokemon-go-gyms.htm)
    Not to be confused with poke Mongo.
1. [Pingdom](http://status.pingdom.com/incidents/ly1r4kyht0yw)
    An outage in both the admin/API and actual monitoring.
1. [ASX (Australia Stock Exchange)](http://www.investordaily.com.au/markets/40090-asx-investigates-service-outage)
1. [Global Switch (London datacenter)](http://www.datacenterdynamics.com/content-tracks/security-risk/details-emerge-of-global-switch-outage-in-london/96970.article)
1. [Phoenix (civil service pay system)](http://www.cbc.ca/news/politics/phoenix-down-september-1.3770156)
    The system used by Canada to pay its civil servants went down the day before payday.
1. [T-Mobile US](http://www.investopedia.com/news/tmobile-lte-outages-have-been-fixed-tmus/)
1. [Fonality](http://www.dailycloud.info/fonality-suffers-four-hour-outage/)

### [ << Prev ](sreweekly-40.md) ------------- [ Next >> ](sreweekly-42.md)