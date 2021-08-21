## [SRE Weekly Issue #8](https://sreweekly.com/sre-weekly-issue-8/) - January 31, 2016
### Articles

1. [Maslow’s Hierarchy of SRE Needs](https://plus.google.com/+lizthegrey/posts/MLAJFVyEb2f)

    Liz Fong-Jones, a Googler and co-chair of SRECon, describes a scale of activities SRE teams engage in, from the basics (keeping the service operating) to  having the freedom to improve the service.
1. [High-Availability at Massive Scale: Building Google’s Data Infrastructure for Ads](http://research.google.com/pubs/pub44686.html)

    This is a really awesome paper. Two Googlers describe in detail the pitfalls of failover-based systems and explain how they design multi-homed active/active services.  If Google has learned a lesson, we’d all do well to learn from it, too:

1. [Cloud outage audit update: The challenges with uptime](http://searchcloudcomputing.techtarget.com/news/4500271502/Cloud-outage-audit-update-The-challenges-with-uptime)

    A review of CloudHarmony’s numbers on various cloud providers’ availability in 2015 versus 2014, along with a discussion of how customers deal with outages.  I’m a little puzzled by this one:
  I’m pretty sure plenty of mission-critical stuff is running in EC2, for example.
1. [Interview: Building the Latest Campaign for David Guetta — Serverless Code](https://serverlesscode.com/post/david-guetta-online-recording-with-lambda/)

    
1. [Europa Water Siphon](http://what-if.xkcd.com/143/)

    Randall Monroe takes on an important question: is it possible to siphon water from a Europa to Earth? Okay, the only relation to SRE is that a team of Google SREs submitted the question, but I really love What If.
1. [The How and Why of Minimum Viable Runbooks](https://victorops.com/knowledge-drop/devops-docs/minimum-viable-runbooks/)

    VictorOps distilled their Minimum Viable Runbooks series (featured here previously) into a polished PDF in their usual high quality and style.
1. [Vodafone set to follow Spark in providing near live information on network faults](http://i.stuff.co.nz/business/industries/76400718/vodafone-set-to-follow-spark-in-providing-near-live-information-on-network-faults)

    During an outage this week, Vodafone admitted that they forgot to update their status site.  They are looking into an automated system to make updates during outages.
1. [On Call Compensation for a Solo Sys Admin – What’s your experience? – Spiceworks](https://community.spiceworks.com/topic/384773-on-call-compensation-for-a-solo-sys-admin-what-s-your-experience)

    I’ve worked mostly jobs without compensation for on-call, but one with.  Compensation is nice, but it was to offset a truly heinous level of pages, so it was small comfort.  If you have any good articles about the merits and pitfalls of on-call compensation, please send them my way.
### Outages

1. [The Patriots’ (sports team) Microsoft Surface tablets](http://www.npr.org/sections/thetwo-way/2016/01/25/464319146/oops-tablet-outage-in-afc-title-game-coincides-with-microsoft-ad)
    This one’s notable because immediately after, a Surface ad played, causing a Microsoft marketing outage.
1. [Microsoft Office 365 IMAP service](http://www.theregister.co.uk/2016/01/25/office_365_imap_outage/)
1. [WhatsApp](http://timesofindia.indiatimes.com/tech/tech-news/WhatsApp-suffers-brief-outage-globally/articleshow/50728488.cms)
1. [Kaiser Permanente](http://www.fiercehealthit.com/story/server-issues-lead-website-outages-kaiser-permanente/2016-01-25)
1. [PlayStation Network](http://www.express.co.uk/entertainment/gaming/638130/PSN-Down-PS4-PlayStation-Network-Problems-Not-Working-Server-Error-Sony)
    A new round of trouble for Sony.
1. [Safari (browser)](http://www.dailymail.co.uk/sciencetech/article-3419088/Safari-outage-affects-iPhone-iPad-users-worldwide.html)
    This one’s interesting: existing installations of Safari started crashing on iOS devices everywhere, due to a bug triggered by a change in some backend web service.
1. [Africa (again)](http://businesstech.co.za/news/internet/110455/why-seacom-has-suffered-another-outage/)
    The backhoe: natural enemy of the network.
1. [Vodafone internet in New Zealand](https://netguide.co.nz/story/vodafone-internet-down-around-new-zealand-customers-want-answers/)
1. [Google Drive](http://www.technewstoday.com/28289-google-drive-faces-outage-worldwide/)
1. [GitHub](https://github.com/blog/2101-update-on-1-28-service-outage)
    Github suffered a two-hour outage on Thursday that was caused by a power outage.  Hats off to them for releasing a postmortem the day after, although I think there’s a lot left unsaid as to why a power outage took them down at all.  It’s a shame, because those kinds of analyses can be especially educational and help us learn how to avoid similar problems.
1. [HSBC online banking](http://arstechnica.com/security/2016/01/hsbc-online-banking-suffers-major-outage-blames-ddos-attack/)
    Another DDoS.

### [ << Prev ](sreweekly-7.md) ------------- [ Next >> ](sreweekly-9.md)