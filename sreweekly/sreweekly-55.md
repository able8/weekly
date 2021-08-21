## [SRE Weekly Issue #55](https://sreweekly.com/sre-weekly-issue-55/) - January 15, 2017
### Articles

1. [Continuous MySQL backup validation: Restoring backups](https://code.facebook.com/posts/1007323976059780/continuous-mysql-backup-validation-restoring-backups/)

    Nothing is worse than finding out that your confidence in your backup strategy was ill-founded (the hard way). Facebook prevents this with what is, in retrospect, a blatantly obvious idea that I never thought of: continuously, automatically testing your backups by trying to restore them.
1. [Running Multiple HTTP Endpoints as a Highly Available Health Proxy](https://www.awsarchitectureblog.com/2014/06/health-proxy.html)

    Route 53 can do failover based on health checks, but it doesn’t know how to check if a database is healthy. This article discusses using an HTTP endpoint that checks the status of the DB and returns status 200 or 500 depending on whether the DB is up. There’s also a discussion of how to handle failure of the HTTP endpoint itself.
1. [Using Chaos Monkey whenever you feel like it](https://medium.com/production-ready/using-chaos-monkey-whenever-you-feel-like-it-e5fe31257a07#.kmxp7sefb)

    Chaos Monkey was designed with the idea of having it run all the time on a schedule, but as Mathias Lafeldt shares, you can also (or even exclusively) trigger failures through an API. He even wrote a CLI for the API.
1. [Four reasons developers should write their own load tests](https://engineering.klarna.com/four-reasons-developers-should-write-their-own-load-tests-fac74c1be9f1)

    Here’s a link shared with me by its author. If you write something you think other SREs will like, please don’t hesitate to send it my way! I love this article, because load testing is yet another aspect of the growing trend toward developers owning the operation of their code.
1. [The First Four Things You Measure](https://honeycomb.io/blog/2017/01/instrumentation-the-first-four-things-you-measure/)

    This article is short and sweet. There are four rock-bottom metrics that you really need to figure out if something is wrong with your service. They had me at “Downstreamistan”.
1. [Chaos Engineering](https://www.infoq.com/articles/chaos-engineering)

    This description of Chaos Engineering is more rigorous than casual articles, making for a pretty interesting read even if you already know all about it.
1. [A Chilling PBS Documentary Shows How Mistakes Are Made](https://www.nytimes.com/2017/01/04/arts/television/a-chilling-pbs-documentary-shows-how-mistakes-are-made.html?_r=0)

    I haven’t had a chance to watch this yet, but the description is riveting even by itself. Click through for a link to play the documentary directly.
### Outages

1. [Second Life](https://community.secondlife.com/t5/Tools-and-Technology/Someone-had-a-case-of-the-Mondays/ba-p/3094252)
    One transit provider failed and automatic failover didn’t work. Once they were back up, the subsequent thundering herd of logins threatened to take them back down. Click through for a detailed post-analysis.
1. [S3, EC2 APIHerokuPackageCloud.io](http://status.aws.amazon.com/)
    On January 10, S3 had issues processing DELETE requests (though you wouldn’t know it from looking at the history section of their status page). Various (presumably) dependent services such as Heroku and PackageCloud.io had simultaneous outages.
Full disclosure: Heroku is my employer.
1. [Lloyds Bank](http://www.theregister.co.uk/2017/01/11/lloyds_bank_group_sites_kicked_offline_by_tech_gremlins/)
1. [Mailgun](http://status.mailgun.com/incidents/p9nxxql8g9rh)
1. [Battlefield 1](http://www.itechpost.com/articles/74113/20170113/battlefield-1-battlefield-1-servers-battlefield-1-servers-down-ps4-xbox-one-battlefield-1-ps4-battlefield-1-xbox-one.htm)
1. [Facebook](http://www.express.co.uk/life-style/science-technology/753940/Facebook-down-social-media-site-not-working)

### [ << Prev ](sreweekly-54.md) ------------- [ Next >> ](sreweekly-56.md)