## [SRE Weekly Issue #117](https://sreweekly.com/sre-weekly-issue-117/) - April 8, 2018
### Articles

1. [No, seriously. Root Cause is a Fallacy. – ](http://willgallego.com/2018/04/02/no-seriously-root-cause-is-a-fallacy/)

    Brilliant, just brilliant. This isn’t just another “there isn’t just one root cause” article to skip over. The author takes time to explain the concept with cogent examples and useful metaphors. This one really caught my eye:Will Gallego
1. [Incident Management – Food Fight Podcast](http://foodfightshow.org/2018/03/incident-management.html)

    This episode of Food Fight is an hour-long interview with guests Rob Schnepp, Ron Vidal, and Chris Hawley, the 3 firefighters behind Blackrock 3 Partners. It’s a great intro to the Incident Management System, and well worth a listen.Shout-out to Maple Player, an android audio player with a really high-quality tempo increase feature. I was able to listen at 1.5x speed and still understand everything; otherwise, I wouldn’t have had time this week.Nell Shamrell-Harrington and Nathen Harvey
1. [Billing Incident Post-Mortem](https://www.twilio.com/blog/2013/07/billing-incident-post-mortem-breakdown-analysis-and-root-cause.html)

    Here’s one from the archives, an incident report from 2013. After a temporary network partition in a redis cluster, the replicas all tried to resynchronize at once, overloading the master. One of the results was that some customers got repeatedly charged for the same thing.Twilio
1. [It’s about what broke, not who broke it](https://rachelbythebay.com/w/2018/03/27/whowhat/)

    Rachel Kroll
1. [Consistent Hashing: Algorithmic Tradeoffs](https://medium.com/@dgryski/consistent-hashing-algorithmic-tradeoffs-ef6b8e2fcae8)

    I thought consistent hashing was largely solved. I was wrong! There are some good solutions out there, but you have to evaluate their relative trade-offs and pick the right one for your use case.Damian Gryski
Full disclosure: Damian Gryski is my coworker at Fastly.
1. [Computer science faces an ethics crisis. The Cambridge Analytica scandal proves it.](http://www.bostonglobe.com/ideas/2018/03/22/computer-science-faces-ethics-crisis-the-cambridge-analytica-scandal-proves/IzaXxl2BsYBtwM4nxezgcP/story.html)

    As you read this article, consider the ethical imperative of system reliability, when system reliability can literally mean life and death in some cases. That’s only going to be more common in the coming years. Yonatan Zunger 
1. [LogicMonitor Uses Terraform, Packer & Consul for Disaster Recovery Environments](http://www.hashicorp.com/blog/logic-monitor-uses-terraform-packer-and-consul-for)

    Randall Thomson — LogicMonitor
1. [The Travis CI Blog: Incident Post-Mortem and Security Advisory: Data Exposure After travis-ci.com Outage](https://blog.travis-ci.com/2018-04-03-incident-post-mortem)

    Oof. Sorry, Travis folks, but a sincere thanks for sharing your experience with us.Konstantin Haase — Travis CI
1. [Preliminary Analysis of the Site Reliability Engineer Survey](http://blog.catchpoint.com/2018/03/15/preliminary-analysis-site-reliability-engineer-survey/)

    I like these “preliminary results” better than the kinds of aggregate statistics you normally get from a survey report. There are real quotes from free-form survey answers, including a couple of real gems. There’s a link to download the actual survey report if you’re into that, too.Dawn Parzych — Catchpoint
### Outages

1. [Statuspage.io](https://metastatuspage.com/incidents/z23rb6wx7szf)
1. [Mindbody Online (fitness studio booking service vendor)](https://www.wellandgood.com/good-sweat/mindbody-online-crash-boutique-fitness/)
1. [Sling TV](https://www.cordcuttersnews.com/sling-tv-has-restored-service-after-overnight-outage/)
1. [Tinder](http://www.financialexpress.com/industry/technology/tinder-suffers-outage-after-facebooks-privacy-fixes/1122515/)
    The outage seemingly stemmed from privacy fixes Facebook put in place, resulting in a broken OAuth flow.
1. [Microsoft Office 365](https://www.computing.co.uk/ctg/news/3029760/microsoft-investigates-causes-of-europe-wide-office-365-outage)
1. [Twitter](https://www.gloucestershirelive.co.uk/news/gloucester-news/thousands-users-across-uk-hit-1428790)
1. [Multiple Indian Government Websites](https://timesofindia.indiatimes.com/india/defence-ministry-website-allegedly-hacked-displays-chinese-character/articleshow/63643681.cms)
1. [Grab](https://coconuts.co/kl/news/days-uber-exits-region-grab-faces-major-service-outage-across-southeast-asia/)
1. [YouTube](https://joltjournal.com/2018/04/02/youtube-channel-pages-are-experiencing-downtime/)

### [ << Prev ](sreweekly-116.md) ------------- [ Next >> ](sreweekly-118.md)