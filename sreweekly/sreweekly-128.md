## [SRE Weekly Issue #128](https://sreweekly.com/sre-weekly-issue-128/) - July 1, 2018
### Articles

1. [The Saddest Moment](http://scholar.harvard.edu/files/mickens/files/thesaddestmoment.pdf)

    Humor for SREs! This is the most hilarious thing I’ve read all week.James Mickens — USENIX ;login:logout
1. [A broad overview of how modern Linux systems boot](https://utcc.utoronto.ca/~cks/space/blog/linux/LinuxBootOverview)

    This focuses on various ways that Linux systems can fail to boot.Chris Siebenmann
1. [DevOps Chat: SRE w/ Stig Sorensen of Bloomberg](https://devops.com/devops-chat-sre-w-stig-sorensen-of-bloomberg/)

    A (raw) transcript of a chat about Bloomberg’s adoption of SRE practices. It might be worth dropping it in a text editor and removing all occurrences of the phrase “sort of”. The real meat is in the discussion of what Bloomberg has learned (text search: “lessons learned”) and how to sell SRE as necessary in a company (text search: “ROI”).Alan Shimel — devops.com
1. [How Pusher Channels has delivered 10,000,000,000,000 messages](https://making.pusher.com/how-pusher-channels-has-delivered-10000000000000-messages/)

    Jim Fisher — Pusher
1. [How we implemented consistent hashing efficiently](https://blog.ably.io/how-to-implement-consistent-hashing-efficiently-fe038d59fff2)

    An in-depth explanation of how consistent hashing works. Love the hand-drawn diagrams!Srushtika Neelakantam — Ably
1. [Colm MacCárthaigh on Twitter: random number generation](https://twitter.com/colmmacc/status/1012719876706840578)

    Not strictly SRE-related, but then again it’s by Colm MacCárthaigh, who is SRE-related.Colm MacCárthaigh
1. [Understanding error budget overspend – part one – CRE life lessons](http://feedproxy.google.com/~r/ClPlBl/~3/HuV7dYkNxJk/understanding-error-budget-overspend-cre-life-lessons.html)

    What should you do if you blow your error budget? Depends on whether you leaked it like a dripping faucet or splurged it all on big outages. Either way, you’ll need to investigate and make a plan.Adrian Hilton, Alec Warner and Alex Bramley — Google
1. [Migrating Messenger storage to optimize performance](https://code.fb.com/core-data/migrating-messenger-storage-to-optimize-performance/)

    I love the two-method approach: a simple migration path for users that aren’t active all the time, and a more careful (and more complex) path for very busy users.Xiang Li and Thomas Georgiou — Facebook
1. [Sri Harsha Kalavala on Twitter: alert on support page views](https://twitter.com/harshaunplugged/status/1012765568892485632)

    Click through for the graph. Monitor status and support page views… do we actually need any other monitoring? Only half-kidding.Sri Harsha Kalavala
### Outages

1. [Google BigQuery](https://status.cloud.google.com/incident/bigquery/18037#18037005)
    Google posted a followup analysis of the BigQuery outage on June 22.
A new release of the BigQuery API introduced a software defect that caused the API component to return larger-than-normal responses to the BigQuery router server.
1. [Fastly](https://status.fastly.com/incidents/gkjfzv6vmgsn)
    Full disclosure: Fastly is my employer.
1. [G Suite Status Dashboard](https://www.google.com/appsstatus#hl=en&v=issue&sid=22&iid=5472df2c903714bd8138fccd76e150dd)
1. [Slack](https://status.slack.com/2018-06/142edcb9e52c7663)
    This week, Slack had a ~3-hour, near-total outage. Click through for their followup post.
The network problems were caused by a bug included in an offline batch process of data, which resulted in unexpected network spikes and led all of our customers to become disconnected and unable to reconnect.
1. [Google Home and Chromecast](https://www.slashgear.com/google-home-and-chromecast-outage-leaves-users-livid-27535829/)

### [ << Prev ](sreweekly-127.md) ------------- [ Next >> ](sreweekly-129.md)