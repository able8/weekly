## [SRE Weekly Issue #38](https://sreweekly.com/sre-weekly-issue-38/) - September 4, 2016
### Articles

1. [The Fire Service and the Aviation Industry – Firefighter Safety – Crew Resource Management](http://www.firehouse.com/article/12249720/the-fire-service-and-the-aviation-industry-firefighter-safety-crew-resource-management)

    What can the fire service learn about safety from the aviation industry?  A 29-year veteran in the fire service answers that question in detail.  We could in turn apply all of those lessons to operating complex infrastructures.
1. [Wikipedia: High Reliability Organization](https://en.wikipedia.org/wiki/High_reliability_organization)

    I’m surprised that I haven’t come across the term “High Reliability Organization” before reading the previous article.  Here’s Wikipedia’s article on HROs.

1. [Tracking Every Release](https://codeascraft.com/2010/12/08/track-every-release/)

    Etsy instruments their deployment system to strike a vertical line on their graphite graphs for every deploy.  This helps them quickly figure out whether a deploy happened right before a key metric took a turn for the worse.
1. [Undersea cables keep global enterprise networks afloat](http://searchnetworking.techtarget.com/feature/Undersea-cables-keep-global-enterprise-networks-afloat)

    A really interesting dive into the world of subsea network cables and the impact that cuts can have on businesses worldwide.
1. [Testing in production comes out of the shadows](http://sdtimes.com/testing-production-comes-shadows/)

    How closely can you really mimic production in your testing environments?  In a way we’re all testing in production, and this article talks about getting that fact out in the open.
1. [Two Suggestions to Help SL Scale](https://www.lexneva.name/blog/2008/06/11/two-suggestions-to-help-sl-scale/)

    I wrote this article on my terrible little blog back in 2008 — forgive the horrid theme and apparently broken unicode support.  This was well before I worked in Linden Lab’s Ops team, back when I was making a living as a user selling content in Second Life.  What’s interesting to me in reading this article 8 years later is the user perspective on the impact of the string of bad outages, and especially Linden’s poor communication during outages.
1. [Delta says it lost $100 million in revenue due to big outage](http://www.heraldcourier.com/workittricities/business_wire/delta-says-it-lost-million-in-revenue-due-to-big/article_99b10454-36ee-5b00-862d-ef6e2712072d.html)

    More on the impact of Delta Airline’s major outage last month.
1. [We’re learning the wrong lessons from airline IT outages](http://www.networkworld.com/article/3113660/data-center/were-learning-the-wrong-lessons-from-airline-it-outages.html)

    
### Outages

1. [Salesforce](http://trust.salesforce.com)
    Salesforce.com was down or impared for several hours.
Full disclosure: Salesforce is the parent company of my employer, Heroku.
1. [dynamodb](http://status.aws.amazon.com)
1. [Telstra Mail](http://www.zdnet.com/article/telstra-repairs-issues-with-telstra-mail-service/)
1. [Google Cloud Platform](http://status.cloud.google.com/incident/compute/16017#5688737870643200)
    Normally I don’t include single-zone failures in EC2 or GCP, but this one has an extremely interesting and detailed postmortem.
1. [EA (FIFA 16 and Battlefield 1 Beta)](http://www.psu.com/news/31000/ea-servers-down-fifa-16-battlefield-1-servers-down-ps4)
1. [Vodafone (Ireland)](https://www.siliconrepublic.com/comms/vodafone-network-outage)
1. [Interpublic Group (Hollywood PR agency)](http://www.hollywoodreporter.com/news/daylong-email-outage-hits-hollywood-924899)
1. [Veskabout page](http://www.theregister.co.uk/2016/08/31/vesk_outage/)
    The Register noted that Vesk bragged about their 100% uptime even after the outage — including for all of 2016.  From Vesk’s recently-changed about page:
VESK has hit 100% uptime for all 2012, 2013, 2014, 2015 and 2016.”
1. [PlayStation Network](http://www.express.co.uk/entertainment/gaming/706471/PlayStation-Network-Down-PSN-under-maintenance-PS4-PSN-Down-today)
1. [PagerDuty](https://status.pagerduty.com/incidents/c4lsczzsdlkw)
    PagerDuty is currently unable to process some inbound events. We are investigating the cause.
1. [Telkom (South Africa telecom)](http://mybroadband.co.za/news/telecoms/177597-massive-telkom-outage-after-sabotage.html)
    The company cited suspected sabotage and offered a monetary reward.
1. [Washington, DC 911 system](http://www.wusa9.com/news/local/dc/human-error-knocks-out-dc-911-system/310860978)
    Emergency services were knocked out for 90 minutes after a contract worker mistakenly hit the emergency shut-off button.  The phrase “human error” is being tossed about.

### [ << Prev ](sreweekly-37.md) ------------- [ Next >> ](sreweekly-39.md)