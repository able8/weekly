## [SRE Weekly Issue #52](https://sreweekly.com/sre-weekly-issue-52/) - December 26, 2016
### Articles

1. [Why Percentiles Don’t Work the Way you Think](https://www.vividcortex.com/blog/why-percentiles-dont-work-the-way-you-think)

    Percentiles are tricky beasts. Does that graph really mean what you think it means?Thanks to Devops Weekly for this one.
1. [Uber blames humans for self-driving car traffic offenses as California orders halt](https://www.theguardian.com/technology/2016/dec/14/uber-self-driving-cars-run-red-lights-san-francisco)

    There’s that magical “human error” again.
1. [Recent ChangeIP Outage](http://resources.changeip.com/recent-changeip-outage/)

    ChangeIP suffered a major outage two weeks ago and they posted this analysis of the incident. Thanks, folks! Does this sound familiar?
1. [Shuffle Sharding: massive and magical fault isolation](https://www.awsarchitectureblog.com/2014/04/shuffle-sharding.html)

    Shuffle sharding is a nifty technique for preventing impact from spreading to multiple users of your service. A great example is the way Route 53 assigns nameservers for hosted DNS zones:sreweekly.com. 172800 IN NS ns-442.awsdns-55.com.
sreweekly.com. 172800 IN NS ns-894.awsdns-47.net.
sreweekly.com. 172800 IN NS ns-1048.awsdns-03.org.
sreweekly.com. 172800 IN NS ns-1678.awsdns-17.co.uk.

1. [Building and scaling the Fastly network, part 2: balancing requests](https://www.fastly.com/blog/building-and-scaling-fastly-network-part-2-balancing-requests)

    Fastly has a brilliant, simple, and clever solution to load balancing and connection draining using a switch ignorant of layer 4.
1. [Heroku Incident 1001: EU Region Network Latency](https://status.heroku.com/incidents/1001)

    Heroku shared a post-analysis of their major outage on December 15.Full disclosure: Heroku is my employer.
### Outages

1. [NTP server poolmore details](https://news.ntppool.org/2016/12/load/)
    Load on the worldwide NTP server pool increased significantly due to a “buggy Snapchat app update”. What was Snapchat doing with NTP? (more details)
1. [Zappos](https://mobile.twitter.com/zappos/status/811417629697028096)
    Zappos had a cross-promotion with T-Mobile, and the traffic overloaded them.Thanks to Amanda Gilmore for this one.
1. [Slack](https://status.slack.com/2016-12/92ca538bc556e84f)
    Among other forms of impairment, /-commands were repeated numerous times. At $JOB, this meant that people accidentally paged their coworkers over and over until we disabled PagerDuty.
1. [Librato](https://status.librato.com/incidents/8mygk2dhtj9p)
    “What went well” is an important part of any post-analysis.
1. [Tumblr](http://www.dailymail.co.uk/sciencetech/article-4057780/Tumblr-hacked-Cyber-criminals-site-Europe-online-attack-just-fun.html)
1. [Southwest Airlines](http://kfoxtv.com/news/nation-world/southwest-airlines-website-goes-offline)

### [ << Prev ](sreweekly-51.md) ------------- [ Next >> ](sreweekly-53.md)