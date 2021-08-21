## [SRE Weekly Issue #47](https://sreweekly.com/sre-weekly-issue-47/) - November 6, 2016
### Articles

1. [SRECon17 Americas Call for Participation](https://www.usenix.org/conference/srecon17americas/call-for-participation)

    Next year, SRECon is expanding to three events: Americas, EMEA, and Asia. The Americas event is also moving from Santa Clara to San Francisco, which I, for one, am especially grateful for. The CFP for SRECon17 Americas just opened up, and proposals are due November 30th, so get cracking! I can’t wait to see what all of you have to share!
1. [Introducing anomaly detection in Datadog](http://www.tmcnet.com/usubmit/-datadog-announces-machine-learning-based-anomaly-detection-cloud-/2016/10/27/8443520.htm)

    I have a somewhat dim view of automated anomaly detection in metrics based on my experience with a few tools, but if Datadog’s algorithms live up to their description, they might really have something worthwhile.
1. [From Zero to Staging and Back](https://medium.com/production-ready/from-zero-to-staging-and-back-cede8e29f5c#.dpl7b1iqa)

    This issue of Production Ready chronicles Mathias Lafeldt’s effort to create a staging environment. I like the emphasis on using an entirely separate AWS account for staging. This is increasingly becoming a best practice.
1. [Honeycomb :: Nylas Guest Post: Ghosts in the WSGI Machine](https://honeycomb.io/blog/2016/11/nylas-guest-post-ghosts-in-the-wsgi-machine/)

    What’s causing all that API request latency? Here’s an interesting debug run using Honeycomb. Negative HTTP status codes? Sure, that’s totally a thing, right?
1. [The Irreproducibility of Bugs in Large-Scale Production Systems](http://www.susanjfowler.com/blog/2016/11/2/the-irreproducibility-of-bugs-in-large-scale-production-systems)

    I love this idea: Susan Fowler notes that large, complex systems are constantly changing, and this makes reproducing bugs difficult or impossible. Her suggestion is to log enough that you can logically reconstruct the state of the system at the time the bug occurred. This is the same kind of thing the Honeycomb folks are saying: throw a lot of information into your logs, just in case you might need it to debug something unforeseen.
### Outages

1. [Instagram](http://www.itechpost.com/articles/48958/20161102/instagram-down-november-1-app-website-break.htm)
1. [Level 3](http://arstechnica.com/information-technology/2016/11/level-3-drops-its-packets-for-hours-causing-internet-traffic-jam/)
    Another big Level 3 outage.
1. [Battlefield 1 (game)](http://www.gamenguide.com/articles/62561/20161103/battlefield-1-news-updates-origin-access-servers-all-down-yesterday-whats-the-status-now.htm)
1. [AT&T](http://www.newschannel10.com/story/33565740/att-has-3rd-service-outage-in-6-days)
    Three separate outages.
1. [FIFA 17 (game)](http://www.product-reviews.net/2016/11/04/fifa-17-ea-server-maintenance-shock-on-nov-4/)
1. [Unprecedented cyber attack takes Liberia’s entire internet down](http://www.telegraph.co.uk/technology/2016/11/04/unprecedented-cyber-attack-takes-liberias-entire-internet-down/)
    The attackers used the Mirai botnet, the same one used to attack Dyn.

### [ << Prev ](sreweekly-46.md) ------------- [ Next >> ](sreweekly-48.md)