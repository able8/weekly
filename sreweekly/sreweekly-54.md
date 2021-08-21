## [SRE Weekly Issue #54](https://sreweekly.com/sre-weekly-issue-54/) - January 8, 2017
### Articles

1. [Open-Sourcing Our Incident Response Documentation](https://www.pagerduty.com/blog/incident-response-documentation/)

    Wow! PagerDuty made waves this week by releasing their internal incident response documentation. This is really exciting, and I’d love it if more companies did this. Their incident response procedures are detailed and obviously the result of hard-won experience. The hierarchical, almost militaristic command and control structure is intriguing and makes me wonder what problems they’re solving.
1. [Load Testing Kept New Relic Running During Cyber Monday](https://blog.newrelic.com/2017/01/05/load-testing-scalability/)

    Lots of detail on New Relic’s load testing strategy, along with an interesting tidbit:
1. [The Problem with Pre-aggregated Metrics: Part 3, the “metrics”](https://honeycomb.io/blog/2017/01/the-problem-with-pre-aggregated-metrics-part-3-the-metrics/)

    Last in the series, this article is an argument that metrics aren’t always enough. Sometimes you need to see the details of the actual events (requests, database operations, etc) that produced the high metric values, and traditional metrics solutions discard these in favor of just storing the numbers.
1. [Let’s Encrypt 2016 In Review](https://letsencrypt.org//2017/01/06/le-2016-in-review.html)

    Let’s Encrypt has gone through a year of intense growth in usage. Their Incidents page has some nicely detailed postmortems, if you’re in the mood.
1. [Blame. Language. Sharing.](http://fractio.nl/2015/10/30/blame-language-sharing/)

    An eloquent post on striving toward a learning culture in your organization, as opposed to a blaming one, when discussing adverse incidents.
1. [Who Called Git? An Unusual Debugging Story](https://www.schneems.com/2016/11/28/who-called-git-an-unusual-debugging-story/)

    I like to include the occasional debugging deep-dive article, because it’s always good to keep our skills fresh. Here’s one from my coworker on finding the source of an unexpected git error message.Full disclosure: Heroku, my employer, is mentioned.
### Outages

1. [Cloudflare leap second](https://blog.cloudflare.com/how-and-why-the-leap-second-affected-cloudflare-dns/)
    Another  leap second, another casualty.
Full disclosure: Heroku, my employer, is mentioned.
1. [U.S. Customs](http://www.usatoday.com/story/tech/news/2017/01/03/border-outage-not-caused-hackers-customs-and-border-patrol-airlines-lines-airport/96107764/)
1. [DomainMonster](http://www.theregister.co.uk/2017/01/04/domainmonster_email_outage/)

### [ << Prev ](sreweekly-53.md) ------------- [ Next >> ](sreweekly-55.md)