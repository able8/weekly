## [SRE Weekly Issue #113](https://sreweekly.com/sre-weekly-issue-113/) - March 11, 2018
### Articles

1. [The New Best Engineer](https://honeycomb.io/blog/2018/03/the-new-best-engineer/)

    The best kind of engineer is one that understands not only their specialty, but at least something about the fields adjacent to theirs. The empathy this confers allows one to work incredibly effectively across the company. For SREs, this is even more important.Charity Majors — Honeycomb
1. [GitLabbers share how to recognize burnout (and how to prevent it)](https://about.gitlab.com/2018/03/08/preventing-burnout/)

    GitLab held a session about recognizing and preventing burnout at their recent employee summit. They share the best tips in this article, and true to their radically open culture, they also added what they learned to their employee handbook, which is publicly available.Clement Ho — GitLab
1. [Travis CI Status – Container-based Linux Precise infrastructure emergency maintenance](https://www.traviscistatus.com/incidents/sxrh0l46czqn)

    Here’s a post-analysis for a Travis CI incident early last year. Despite a couple of easy targets that could have been labelled as “root cause”, they instead skillfully laid out all of the contributing factors and left it at that.Travis CI
1. [Serverless Ops: What do we do when the server goes away?](https://www.serverlessops.io/blog/serverless-ops-what-do-we-do-when-the-server-goes-away)

    What indeed? The same thing, just organized differently. There’s a lot of great analysis here about how ops roles can adapt to a serverless infrastructure, and how teams can best make use of ops folks.Tom McLaughlin — ServerlessOps
1. [On-Call Rotations: How Best to Wake Devs Up in the Middle of the Night](https://thenewstack.io/call-rotations-best-wake-devs-middle-night/)

    Charity Majors wants you to look forward to on-call. This superb write-up of her recent conference talk explains why folks should think of on-call as an enjoyable privilege and how to shape your on-call to get there.Jennifer Riggins
1. [Canary Analysis Service](http://queue.acm.org/detail.cfm?id=3194655)

    The Canary Analysis Service is Google’s internal tool that automatically analyzes canary runs and decides whether performance has been negatively impacted. My favorite section is the Lessons Learned.Štěpán Davidovič with Betsy Beyer — ACM Queue
### Outages

1. [Snapchat](http://www.newsweek.com/snapchat-wont-refresh-servers-down-not-working-833094)
1. [123 Reg (hosting provider)](https://www.bleepingcomputer.com/news/business/123-reg-backup-snafu-causes-clients-to-lose-files-since-august-2017/)
    Customers lost files added since 123 Reg’s last valid backup from August, 2017.
1. [partypoker](https://www.highstakesdb.com/8561-partypoker-adds-nearly-2000000-in-guarantees-after-server-crash.aspx)
1. [eBay](https://www.ecommercebytes.com/C/blog/blog.pl?/comments/2018/3/1520289779.html)
1. [Signal and Telegram (messenger apps)](https://thenextweb.com/apps/2018/03/05/signal-messenger-many-users-globally/)
1. [Alexa](http://www.techtimes.com/articles/222284/20180304/amazon-alexa-is-having-some-major-issues-voice-commands-for-most-people-aren-t-working.htm)
    I missed this one last week — it was apparently due to the AWS outage I reported on.
1. [TD Bank](https://www.bizjournals.com/philadelphia/news/2018/03/08/td-bank-online-outage-march-braca-gartner.html)
1. [Oculus Rift](https://www.rollingstone.com/glixel/news/oculus-issues-credit-for-outage-details-vr-headset-repair-w517590)
    A code-signing certificate expired, rendering some existing VR headsets non-functional. Oculus is issuing a $15 store credit to affected customers.
Because of the particulars of what expired and how it happened, the company wasn’t able to simply push an update out to users because the expired certificate was blocking Oculus’ standard software update system.

### [ << Prev ](sreweekly-112.md) ------------- [ Next >> ](sreweekly-114.md)