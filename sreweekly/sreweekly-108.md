## [SRE Weekly Issue #108](https://sreweekly.com/sre-weekly-issue-108/) - February 4, 2018
### Articles

1. [Talking Technology: Nick Rockwell + Charity Majors](https://open.nytimes.com/talking-technology-nick-rockwell-charity-majors-2acad1690dcf)

    There’s so incredibly much awesome in this conversation, and I’ve already seen the internet alight with people quoting it. Charity says so many insightful things that I’m going to have to reread this a couple of times to absorb it all. It’s a must-read!
1. [SRE@Xero: Managing Incidents Part II](https://devblog.xero.com/sre-xero-managing-incidents-part-ii-224a6e06f426)

    Xero SRE is back, this time with an article about their incident response process and an overview of their chatbot, Multivac. The bot assists with paging and information tracking and, crucially, guides incident responders through a checklist of actions such as determining severity.
1. [Production Test Run The self flagellating server](https://ayende.com/blog/181441-C/production-test-run-the-self-flagellating-server)

    Here’s a fun little distributed system debugging story from the founder of RavenDB.
1. [Hawaii false missile alert ‘button pusher’ fired](http://www.cnn.com/2018/01/30/us/hawaii-false-alarm-investigation/index.html?utm_source=CNN+Five+Things&utm_campaign=4745781153-EMAIL_CAMPAIGN_2018_01_31&utm_medium=email&utm_term=0_6da287d761-4745781153-83152125)

    This CNN article goes into a little more detail about what happened. To my eye, there’s not enough in those details to warrant firing, so there must be more than has been shared publicly.
1. [Lessons Learned from LinkedIn’s Data Center Journey](https://engineering.linkedin.com/blog/2018/02/lessons-learned-from-linkedins-data-center-journey)

    LinkedIn’s growth from a single datacenter to multiple “hyperscale” locations was accompanied by a cultural shift. They transitioned from “‘Site-Up’ is priority #1” to “taking intelligent risks” as their overall reliability improved.
1. [Vanderbilt School of Engineering offers new master of risk, reliability, and resilience engineering](https://news.vanderbilt.edu/2018/01/31/vanderbilt-school-of-engineering-offers-new-master-of-risk-reliability-and-resilience-engineering/)

    The program is nominally aimed toward “a variety of industries, including the aerospace, automotive, maritime, manufacturing, oil, chemical, power transmission, medical device, infrastructure planning and extreme event response sectors”, though I can’t help but wonder if it might be applicable to IT.
1. [Stop Wasting Your Beer Money](https://blog.realkinetic.com/stop-wasting-your-beer-money-12c3fe5e4d54)

    This author pushes us to resist the urge to write something in-house and instead look for external services or software, when the tool is not key to delivering customer value.
1. [Feature Flags as a Service: The Only Way You Want Feature Flags](https://rollout.io/blog/feature-flags-as-a-service/)

    Here’s a very well-articulated argument for using a third-party feature-flag service rather than writing your own. I’ve seen every pitfall they mention and more. This article is by Rollout.io, a feature-flag service, but they notably don’t mention their product even once, and they don’t need to. Nicely done, folks.
1. [Using Postmortems to Understand Service Reliability](https://www.pagerduty.com/blog/postmortem-understand-service-reliability/)

    We should look beyond merely preventing the same kind of incident in the future and improving our incident response process, says this article from PagerDuty.
1. [Predicting Resource Exhaustion with Double Exponential Smoothing](https://signalfx.com/blog/predicting-resource-exhaustion-double-exponential-smoothing/)

    How many times have you been paged for a server at 95% disk usage, only to find that it’s still months away from full? This article by SignalFX is about a feature on their platform, but its concepts are generally applicable to other tools.
1. [Planning for Chaos with MongoDB Atlas: Using the “Test Failover” Button | MongoDB](https://www.mongodb.com/blog/post/planning-for-chaos-with-mongodb-atlas-using-the-test-failover-button?utm_medium=dzone-synd)

    A primer on testing failover in a MongoDB Atlas cluster.
1. [Meltdown Performance Impact on MongoDB: AWS, Azure](https://scalegrid.io/blog/meltdown-performance-impact-on-mongodb-aws-azure-digitalocean/?)

    Large numbers of SREs went scrambling last month when we realized that we may suddenly run out of resources on our NoSQL workloads. Here are some concrete numbers on how things actually turned out.
### Outages

1. [PolitiFact](http://www.cnn.com/2018/01/30/politics/politifact-crash-state-of-the-union-2018/index.html?utm_source=CNN+Five+Things&utm_campaign=4745781153-EMAIL_CAMPAIGN_2018_01_31&utm_medium=email&utm_term=0_6da287d761-4745781153-83152125)
    PolitiFact was down for a bit during President Trump’s yearly State of the Union address.
1. [Skype](https://www.theregister.co.uk/2018/01/29/skype_signin_issues/)
    It seems that folks with two-factor authentication were unable to log in for multiple days.
1. [The Travis CI Blog: Major build outage: a postmortem report](http://blog.travis-ci.com/2018-01-31-build-outage-postmortem)
    Linked is a highly detailed summary of their troubles with an overloaded RabbitMQ cluster.
1. [Netflix](https://www.express.co.uk/news/world/912252/Netflix-down-streaming-issues-iPhone-iPad-laptop-Twitter-reaction)

### [ << Prev ](sreweekly-107.md) ------------- [ Next >> ](sreweekly-109.md)