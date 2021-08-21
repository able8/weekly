## [SRE Weekly Issue #46](https://sreweekly.com/sre-weekly-issue-46/) - October 30, 2016
### Articles

1. [Being an Effective Ally to Women and Non-Binary People](https://codeascraft.com/2016/10/19/being-an-effective-ally-to-women-and-non-binary-people/)

    I’ve linked to several posts on Etsy’s Code as Craft blog in the past, and here’s another great one. Perhaps not the typical SRE article you might have been expecting me to link to, but this stuff is important in every tech field, including SRE. We can’t succeed unless every one of us has a fair chance at success.
1. [Hacker puts ‘full redundancy’ code-hosting firm out of business](http://www.pcworld.com/article/2365602/hacker-puts-full-redundancy-codehosting-firm-out-of-business.html)

    In 2014, CodeSpaces suffered a major security breach that abruptly ended their business. I’d say that’s a pretty serious reliability risk right there, showing that security and reliability are inextricably intertwined.
1. [OUTAGE! AMA](http://pages.catchpoint.com/OUTAGE-AMA.html)

    Check it out! Catchpoint is doing another Ask Me Anything, this time about incident response. Should be interesting!
1. [Complex System Failures and Blameless Retrospectives by Courtney Eckhardt](https://m.youtube.com/watch?v=Sj0sdbiyatk)

    My fellow Heroku SRE, Courtney Eckhardt, expanded on a section of our joint SRECon talk for this session at OSFeels. She had time for Q&A, and there were some really great questions!
1. [The Power of Less Code](https://medium.com/production-ready/the-power-of-less-code-56764e2cd534#.4xunt244j)

    Mathias rocks it, as usual, in this latest issue of Production Ready.
1. [The Netflix Tech Blog: Netflix Chaos Monkey Upgraded](http://techblog.netflix.com/2016/10/netflix-chaos-monkey-upgraded.html)

    Netflix has released a new version of Chaos Monkey with some interesting new features.
1. [The Myth of the Root Cause: How Complex Web Systems Fail](http://blog.scalyr.com/2016/10/the-myth-of-the-root-cause/)

    Scalyr worked with Mathias Lafeldt to turn his already-awesome pair of articles into this excellent essay. He brings in real-world examples of major outages and draws conclusions based on them. He also hits on a number of other topics he’s written about previously. Great work, folks!
1. [octocatalog-diff: GitHub’s Puppet development and testing tool – GitHub Engineering](http://githubengineering.com/octocatalog-diff-github-s-puppet-development-and-testing-tool/)

    How many times have you pushed out a puppet change that you tested very thoroughly, only to find that it did something unexpected on a host you didn’t even think it would apply to? Etsy’s solution to that is this tool that shows catalog changes for all host types in your fleet in a diff-style format.
1. [Pokemon Go: How the cloud saved the smash hit game from collapse](http://www.techrepublic.com/article/pokemon-go-how-the-cloud-saved-the-smash-hit-game-from-collapse/)

    It’s pretty impressive to me that Niantic managed to keep Pokemon Go afloat as well as they did. They worked very closely with Google to grow their infrastructure much faster than they had planned to.
1. [Spotify Engineering: Making Ops Human](http://feedproxy.google.com/~r/serverdensity/~3/u9DJxFV_7H8/)

    As Spotify has grown to 1400 microservices and 10,000 servers, they’ve moved toward a total ownership model, in which development teams are responsible for their code in production.
1. [How the Friday DDoS attack affected Pingdom](http://royal.pingdom.com/2016/10/24/ddos-attack-affects-pingdom/)

    Pingdom suffered a major outage during the Dyn DDoS, not only due to their own DNS-related issues, but also due to the massive number of alerts their system was trying to send out to notify customers that their services were failing.
1. [Dyn Analysis Summary Of Friday October 21 Attack](http://dyn.com/blog/dyn-analysis-summary-of-friday-october-21-attack/)

    Here’s Dyn’s write-up of the DDoS.
1. [Service Disruption Root Cause Analysis and Follow-up Actions from October 21st, 2016](https://www.pagerduty.com/blog/service-disruption-root-cause-analysis-follow-actions-october-21st-2016/)

    As they promised, here’s PagerDuty’s root cause analysis from the Dyn DDoS.
1. [Routing around single point of failure DNS issues](https://blog.ably.io/routing-around-single-point-of-failure-dns-issues-7c20a8757903#.9paqmyma9)

    This is a pretty great idea. Ably has written their client libraries to reach their service through a secondary domain if the primary one is having DNS issues. Interestingly, their domain, ably.io, was also impacted by the .io TLD DNS outage (detailed below) just days after they wrote this.
1. [Looking Back On The Largest DDoS In History](http://blog.statuspage.io/looking-back-on-the-largest-ddos-in-history)

    StatusPage.io gives us some really interesting numbers around the Dyn DDoS, based on status posts made by their customers.  I wonder, was it really the largest in history?
1. [Serverless Operations is Not a Solved Problem](https://www.infoq.com/news/2016/10/Serverless-Operations)

    Here’s a nice write-up of day one of ServerlessConf, with the theme, “NoOps isn’t a thing”.
1. [AWS Server Migration Service](https://aws.amazon.com/blogs/aws/new-aws-server-migration-service/)

    Not a magic bullet, but still pretty interesting.
1. [Reset router could have saved Census](http://www.technologydecisions.com.au/content/security/news/reset-router-could-have-saved-census-80336641)

    Earlier this year, Australia’s online census site suffered a major outage. Here’s a little more detail into what went wrong. TL;DR: a router dropped its configuration on reboot.
1. [How GOV.UK Reduced their Incidents and Alerts](http://feedproxy.google.com/~r/serverdensity/~3/6SWqyz9HePI/)

    Gov.uk has put in place a lot of best practices in incident response and on-call health.Unfortunately, I’m guessing one of those six types happened this week, as you can see in the Outages section below.
### Outages

1. [.io TLDnegative-caching TTL](https://news.ycombinator.com/item?id=12813065)
    The entire .io top-level domain went down, resulting in impact to a lot of trendy companies that use *.io domains. It doesn’t matter how many DNS providers you have for your domain if your TLD’s nameservers aren’t able to give out their IPs. Worse yet, .io‘s servers sometimes did respond, but with an incorrect NXDOMAIN for valid domains. .io‘s negative-caching TTL of 3600 seconds made this pretty nasty.
On the plus side, this outage provided the last piece in the puzzle in answering my question, “does ‘fast-fluxing’ your DNS providers really work?”. Answer: no. I’ll write up all of my research soon and post a link to it here.
1. [The Pirate Bay](http://www.christiantoday.com/article/the.pirate.bay.news.tpb.now.back.online.after.weekend.downtime/98713.htm)
1. [California DMV](https://www.dmv.ca.gov/portal/dmv/detail/about/field_office_outage_update?lang=en)
1. [AT&T](http://nbc4i.com/2016/10/28/att-users-report-major-phone-service-outage/)
1. [British Telecom](http://www.itproportal.com/news/bt-internet-outage-affects-thousands-across-the-uk/)
1. [gov.uk](http://metro.co.uk/2016/10/26/fears-of-cyberattack-on-uk-government-as-entire-gov-uk-website-goes-down-6216800/)
1. [PlayStation Network](https://www.vg247.com/2016/10/26/all-psn-services-are-down-sony-investigating/)

### [ << Prev ](sreweekly-45.md) ------------- [ Next >> ](sreweekly-47.md)