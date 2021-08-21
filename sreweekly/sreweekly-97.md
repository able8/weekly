## [SRE Weekly Issue #97](https://sreweekly.com/sre-weekly-issue-97/) - November 12, 2017
### Articles

1. [SRE@Xero: Managing Incidents Part I](https://devblog.xero.com/sre-xero-managing-incidents-part-i-7d02d650a71c)

    Last month, I linked to an article on Xero’s incident response process, and I said:This article goes into detail on how the form works, why they have it, and the actual questions on the form! Then they go on to explain their “on-call configuration as code” setup, which is really nifty. I can’t wait to see part II and beyond.
1. [Stretching Spokes](https://githubengineering.com/stretching-spokes/)

    Spokes is GitHub’s system for storing distributed replicas of git repositories. This article explains how they can do this over long distances in a reasonable amount of time (and why that’s hard). I especially love the “Spokes checksum” concept.
1. [Fly the airplane: Three practices for effective incident response](https://www.itproportal.com/features/fly-the-airplane-three-practices-for-effective-incident-response/)

    From the CEO of NS1, a piece on the value of checklists in incident response.
1. [The Ultimate Guide to Secondary DNS](https://dzone.com/articles/the-ultimate-guide-to-secondary-dns)

    Here’s another great guide on the hows and whys of secondary DNS, including options on dealing with nonstandard record types that aren’t compatible with AXFR.
1. [Availability has a new meaning. And it doesn’t include planned downtime.](https://www.itproportal.com/features/availability-has-a-new-meaning-and-it-doesnt-include-planned-downtime/)

    From a customer’s perspective, “planned downtime” and “outage” often mean the same thing.
1. [Risks of a “serverless” future: dissolving valuable infrastructure](https://siliconangle.com/blog/2017/11/08/risks-serverless-future-dissolving-valuable-infrastructure-serverlessconf/)

    “serverless” != “NoOps”
1. [Root Cause Analysis as Storytelling – Wide Awake Developers](http://www.michaelnygard.com/blog/2017/11/root-cause-analysis-as-storytelling/)

    When we use root cause analysis, says Michael Nygard, we narrow our focus into counter-factuals that get in the way of finding out what really happened.CW: hypothetical violent imagery
### Outages

1. [Heroku](https://status.heroku.com/incidents/1334)
    Heroku posted a public followup for incident #1334, with a pretty interesting cause. At the end of the month, load on an internal API increased because the number of apps that ran out of monthly free quota hit a peak.
Full disclosure: Heroku is my employer.
1. [How a Tiny Error Shut Off the Internet for Parts of the USroute leak](https://www.wired.com/story/how-a-tiny-error-shut-off-the-internet-for-parts-of-the-us/)
    I normally don’t include ISP failures, but this one was widespread across the US and had an interesting cause. Level 3 accidentally created a route leak that broke traffic for many Comcast customers (including me).
1. [Google App Engine Memcache Service](http://status.cloud.google.com/incident/appengine/17007#5750790484393984)
    Linked is Google’s followup analysis, which suggests that the outage was due to a scaling issue in a configuration database.
1. [OVH to Disassemble Container Data Centers after Epic Outage in Europe](http://www.datacenterknowledge.com/uptime/ovh-disassemble-container-data-centers-after-epic-outage-europe)
1. [Snapchat](https://www.dailystar.co.uk/tech/news/658998/Snapchat-DOWN-AGAIN-Notifications-not-working-as-social-network-glitches-return)
1. [Instagram](http://indianexpress.com/article/technology/social/instagram-outage-app-goes-down-for-users-globally-4930942/)
1. [E-Trade](http://fortune.com/2017/11/09/e-trade-is-down/)
1. [Grindr](https://heavy.com/news/2017/11/grindr-outage-down-status-dating-app/)
1. [Netflix](http://www.express.co.uk/life-style/science-technology/877198/Netflix-down-not-working-connection-login)
1. [Yahoo Mail](http://www.express.co.uk/life-style/science-technology/876620/Yahoo-Mail-down-not-working-login-e-mails-send-read)

### [ << Prev ](sreweekly-96.md) ------------- [ Next >> ](sreweekly-98.md)