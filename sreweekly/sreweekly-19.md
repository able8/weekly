## [SRE Weekly Issue #19](https://sreweekly.com/sre-weekly-issue-19/) - April 17, 2016
### Articles

1. [A mystery with memory leaks and a magic number](https://rachelbythebay.com/w/2016/02/21/malloc/)

    
1. [Growing Up with MySQL: How we scaled our primary datastore by over 20x in 3 weeks](https://www.nylas.com/blog/growing-up-with-mysql/)

    
1. [Complacency: The Enemy of Resilience](https://medium.com/production-ready/complacency-the-enemy-of-resilience-8123bc4210f9#.8bcgcwmkf)

    
1. [timesketch](http://www.timesketch.org/)

    
1. [SRE: An incomplete guide to cultural Narnia](http://anthonycaiafa.com/2016/04/10/sre-cultural-narnia/)

    
1. [Introducing DGit](http://githubengineering.com/introducing-dgit/)

    
1. [Book review: Site Reliability Engineering](https://hashbang.ca/2016/04/07/book-review-site-reliability-engineering)

    
1. [The importance of 'dogfooding' in the cloud](http://www.infoworld.com/article/3051017/cloud-computing/the-importance-of-dogfooding-in-the-cloud.html)

    Thanks to Charity for this one.
1. [Man accidentally ‘deletes his entire company’ with one line of bad code](http://www.independent.co.uk/life-style/gadgets-and-tech/news/man-accidentally-deletes-his-entire-company-with-one-line-of-bad-code-a6984256.html)

    Note that there’s been some question on hangops #sre on whether this is a hoax.  Either way I could totally see it happening.
1. [On Writing Well When You’re In A Damn Hurry](http://blog.statuspage.io/on-writing-well-when-you-re-in-a-damn-hurry)

    
### Outages

1. [Yahoo Mail](http://www.techtimes.com/articles/148931/20160411/yahoo-mail-goes-down-for-many-users-in-europe-and-us.htm)
1. [Business Wire](http://www.bloomberg.com/news/articles/2016-04-11/buffett-s-business-wire-unavailable-as-news-releases-halted)
1. [Google Compute Engine](https://status.cloud.google.com/incident/compute/16007?post-mortem)
    GCE suffered a severe network outage.  It started as increased latency and at worst became a full outage of internet connectivity.  Two days after the incident, Google released the best postmortem I’ve seen in a very long time.  Full transparency, a terrible juxtaposition of two nasty bugs, a heartfelt apology, fourteen(!) remediation items… it’s clear their incident response was solid and they immediately did a very thorough retrospective.
1. [North Korea](http://www.networkworld.com/article/3055994/internet/something-strange-just-happened-with-north-koreas-internet.html)
    North Korea had a series of internet outages, each of the same length at the same time on consecutive days.  It’s interesting how people are trying to learn things about the reclusive country just from this pattern of outages.
1. [Blizzard's Battle.net](http://www.cloudpro.co.uk/leadership/risks/5940/blizzards-battlenet-taken-offline-by-hackers)
1. [Twitter](http://www.techtimes.com/articles/150224/20160414/twitter-experiences-brief-service-disruption-across-the-globe.htm)
1. [Misco](http://www.channelnomics.eu/channelnomics-eu/news/2454643/misco-apologises-for-seven-hour-website-outage)
1. [Two Alt-Coin exchanges (Shapeshift and Poloniex)](https://www.cryptocoinsnews.com/uncertainty-alt-coin-markets-popular-exchanges-suffer-downtime/)
1. [Home Depot](http://abcnews.go.com/US/wireStory/home-depot-electronic-outage-slowed-card-purchases-38277063)

### [ << Prev ](sreweekly-18.md) ------------- [ Next >> ](sreweekly-20.md)