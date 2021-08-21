## [SRE Weekly Issue #147](https://sreweekly.com/sre-weekly-issue-147/) - November 11, 2018
### Articles

1. [Honeycomb’s Charity Majors: Go Ahead, Test in Production](https://thenewstack.io/honeycombs-charity-majors-go-ahead-test-in-production/)

    This is an excellent summary of a talk on testing in production last month. Joab Jackson — The New Stack
1. [DBMS Musings: Distributed consistency at scale: Spanner vs. Calvin](http://dbmsmusings.blogspot.com/2017/04/distributed-consistency-at-scale.html)

    The Pros and Cons of Calvin and Spanner, two data-store papers published in 2012. According to the author, Calvin largely stands out as the favorite.Daniel Abadi
1. [RobinHood: tail latency aware caching – dynamic reallocation from cache-rich to cache-poor](https://blog.acolyer.org/2018/10/26/robinhood-tail-latency-aware-caching-dynamic-reallocation-from-cache-rich-to-cache-poor/)

    What a cool concept!Adrian Colyer — The Morning Paper (summary)
Berger et al. (original paper)
1. [Cross shard transactions at 10 million requests per second](https://blogs.dropbox.com/tech/2018/11/cross-shard-transactions-at-10-million-requests-per-second/)

    With thousands(!) of MySQL shards, Dropbox needed a way to have transactions span multiple shards while maintaining consistency.Daniel Tahara — Dropbox
1. [Heatmaps Make Ops Better](https://www.honeycomb.io/blog/heatmaps-make-ops-better/)

    This is an excellent introduction to heatmaps with some hints on how to interpret a couple common patterns.Danyel Fisher — Honeycomb
1. [How Automatic Root Cause Analysis Works](https://www.instana.com/blog/how-automatic-root-cause-analysis-works/)

    This is a neat idea. By modelling the relationships between the components in your infrastructure, you can figure out which one might be to blame when everything starts alerting at once. Note: this article is heavily geared toward Instana.Steve Waterworth — Instana
1. [Getafix: How Facebook tools learn to fix bugs automatically](https://code.fb.com/developer-tools/getafix-how-facebook-tools-learn-to-fix-bugs-automatically/)

    Automated bug fixing seems to be all the rage lately. I wonder, is it practical for companies that aren’t the size of Facebook or Google?Johannes Bader, Satish Chandra, Eric Lippert, and Andrew Scott — Facebook

### Outages

1. [Slack in Europe](https://status.slack.com/2018-11-07)
1. [Netflix](https://www.express.co.uk/life-style/science-technology/1043276/Netflix-down-server-status-service-unavailable-not-working)
1. [Instagram](https://www.thesun.co.uk/tech/7693216/instagram-down-news-feed-not-working/)
1. [Microsoft’s Windows license activation service](https://www.itnews.com.au/news/microsoft-server-outage-deactivates-windows-licenses-515335)
    Microsoft has acknowledged a problem affecting its Windows license activation servers in multiple countries that has resulted in users being told their Windows 10 Pro and Enterprise installations are invalid.
1. [Lloyds Bank](https://www.thesun.co.uk/money/7680668/lloyds-internet-banking-down/)
1. [GPS ankle bracelets in Australia](https://www.dailymail.co.uk/news/article-6350809/Prison-parolee-run-Telstra-outage-knocks-800-electronic-trackers.html)
    A violent parolee is on the run after part of the GPS tracking system broke down due to Telstra’s network issues on Friday and Saturday.

### [ << Prev ](sreweekly-146.md) ------------- [ Next >> ](sreweekly-148.md)