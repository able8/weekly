## [SRE Weekly Issue #88](https://sreweekly.com/sre-weekly-issue-88/) - September 10, 2017
### Articles

1. [Resources for Getting Started with Distributed Systems](https://caitiem.com/2017/09/07/getting-started-with-distributed-systems/)

    From Catie McCaffrey:
1. [Finding out if/why a server is dropping packets](https://jvns.ca/blog/2017/09/05/finding-out-where-packets-are-being-dropped/)

    Julia Evans just blew my mind (once again). In this article, among other things, she links to a tool that tells you which function in the kernel dropped a packet. I’ve been wishing for such a tool for years!
1. [3 lessons learned from an Elasticsearch game day](https://www.datadoghq.com/blog/elasticsearch-game-day/?__s=bwykwk1kcceogszq8abt)

    I love that companies are starting to publish lessons learned from game days and other chaos experiments. Just like a post-incident followup, there’s so much we can learn by following along.
1. [The missing link in your disaster recovery strategy](https://itbrief.co.nz/story/missing-link-your-disaster-recovery-strategy/)

    
1. [Post-Incident Reviews Part Four: A Post-Incident Review Guide and Next Best Steps](https://victorops.com/blog/post-incident-reviews-part-four-a-post-incident-review-guide-and-next-best-steps/)

    Here’s the last installment of Jason Hand’s digest version of his new eBook, Post-Incident Reviews.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [What If Your Colo Fails: Preventing Disaster](http://www.datacenterjournal.com/happens-colo-fails-protect-disaster/)

    How can you prevent a colo failure? Obviously, colo customers can’t, but we can at least prepare. This article has advice for understanding a provider’s history, policies, and procedures related to outages.
1. [Who Owns My Availability?](https://www.whoownsmyavailability.com/)

    Just click through.
1. [Fear of Landing – When Both Your Mind and Your Instruments Are Lying](https://fearoflanding.com/accidents/accident-reports/when-both-your-mind-and-your-instruments-are-lying/)

    In this analysis of the factors leading to a plane crash, we see another example of the critical role that human/computer interfaces play in allowing (or preventing) humans to recover from a system failure.
1. [Hurricane Harvey: A Test-Bed For Internet Resilience](http://www.xconomy.com/san-francisco/2017/09/05/hurricane-harvey-a-test-bed-for-internet-resilience/)

    Move over, backhoes: water is the other natural enemy of the fiber optic network.
1. [Publishing with Apache Kafka at The New York Times](https://www.confluent.io/blog/publishing-apache-kafka-new-york-times/)

    The New York Times has a Kafka installation containing everything they’ve published in their entire history, and it powers the front page, search, suggestions, and everything else.
### Outages

1. [AbeBooks.com](http://www.thebookseller.com/news/abe-books-outage-leaves-independent-booksellers-loss-631006)
    AbeBooks is the place to go for out-of-print books and old editions. The site going down meant that many used booksellers lost a major sales outlet.
1. [Gmail](https://www.iol.co.za/business-report/technology/google-apologises-for-gmail-outage-908653)
1. [Apple developer portal](https://www.cnet.com/news/apple-blames-software-bug-for-developer-portal-gaffe/)
1. [Google Drive](http://heavy.com/tech/2017/09/google-drive-down-outages-problems-fixed/)
1. [iCloud Mail](http://9to5mac.com/2017/09/07/icloud-mail-seemingly-experiencing-widespread-downtime/)
1. [HerokuIncidents 1251 and 1254Incident 1257Incident 1270](https://status.heroku.com/)
    Heroku posted a pile of public followups this past week:

Incidents 1251 and 1254 – In both of these incidents, applications failed due to missing debian packages normally provided by the Heroku platform.
Incident 1257 – For a few minutes, 10% of requests to Heroku applications hosted in Europe failed.
Incident 1270 – Applications last deployed over 3 years ago spontaneously stopped working.

Full disclosure: Heroku is my employer.Incidents 1251 and 1254 – In both of these incidents, applications failed due to missing debian packages normally provided by the Heroku platform.Incident 1257 – For a few minutes, 10% of requests to Heroku applications hosted in Europe failed.Incident 1270 – Applications last deployed over 3 years ago spontaneously stopped working.

### [ << Prev ](sreweekly-87.md) ------------- [ Next >> ](sreweekly-89.md)