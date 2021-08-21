## [SRE Weekly Issue #7](https://sreweekly.com/sre-weekly-issue-7/) - January 24, 2016
### Articles

1. [Clients are Jerks: aka How Halo 4 DoSed the Services at Launch & How We Survived](http://caitiem.com/2015/06/23/clients-are-jerks-aka-how-halo-4-dosed-the-services-at-launch-how-we-survived/)

    This article is gold.  CatieM explains why clients can’t be trusted, even when they’re written in-house.  She describes how her team avoided an outage during the Halo 4 launch by turning off non-essential functionality.  Had she trusted the clients, she might not have built in the kill switches that let her shed the excessive load caused by a buggy client.
1. [Under the hood: Broadcasting live video to millions](https://code.facebook.com/posts/1653074404941839/under-the-hood-broadcasting-live-video-to-millions/)

    Facebook recently released a live video streaming feature.  Because they’re Facebook, they’re dealing with a scale that existing solutions can’t even come close to supporting (think millions of viewers for celebrity live video broadcasts).  This article goes into detail about how they handle that level of concurrency for live streaming.  I especially like the bit about request coalescing.
1. [Uptime Funk](http://highscalability.com/blog/2016/1/10/uptime-funk-best-sysadmin-parody-video-ever.html)

    Best.  I pretty much only like the parodies of Uptown Funk.
1. [The Monitoring Death Spiral](https://blog.raintank.io/the-monitoring-death-spiral/)

    This is a really great little essay comparing running a large infrastructure with flying a plane by instruments.  Paying attention to just one or two instruments without understanding the big picture results in errors.Thanks to Devops Weekly for this one.
1. [Diagnosing Human Fail](https://blog.raintank.io/litmus-insights-diagnosing-human-fail/)

    An awesome incident response summary for an outage caused by domain name expiration.  The live Grafana charts are awesome, along with the dashboard snapshot.  It’s exciting to see how far that project has come!
1. [The Factors That Impact Availability, Visualized](https://www.vividcortex.com/blog/the-factors-that-impact-availability-visualized)

    Calculating availability is hard.  Really hard.  First, you have to define just what constitutes availability in your system.  Once you’ve decided how you calculate availability, you’ve defined the goalposts for improving it.  In this article, VividCortex presents a general, theoretical formula for availability and a corresponding 3D graph that shows that improving availability involves both increasing MTBF and reducing MTTR.
1. [Cloudy, with a chance of outage](http://www.techcentral.ie/cloudy-with-a-chance-of-outage/)

    TechCentral.ie gives us this opinion piece on the frequency of outages in major cloud providers.  The author argues that, though reported outages may seem major, they still rarely cause violation of SLAs, and service availability is still probably better than individual companies could manage on their own.Full disclosure: Heroku, my employer, is mentioned.
1. [JetBlue, Verizon data center downtime raises DR, UPS questions](http://searchdatacenter.techtarget.com/news/4500271203/JetBlue-Verizon-data-center-downtime-raises-DR-UPS-questions)

    An external post-hoc analysis of the recent outage at JetBlue, with speculation on the seeming lack of effective DR plans at JetBlue and Verizon.  The article also mentions the massive outage at 365 Main’s San Francisco datacenter in 2007, which is definitely worth a read if you missed that one.
1. [Why Things Were Less Than Optimal This Past Weekend](https://community.secondlife.com/t5/Tools-and-Technology/Why-Things-Were-Less-Than-Optimal-This-Past-Weekend/ba-p/2994441)

    Linden Lab Systems Engineer April wrote up a detailed postmortem of the multiple failures that went into a rough weekend for Second Life users.  I worked on recovery from at least a few failures in that central database in my several years at Linden, and it’s pretty tricky managing the thundering herd that floods through the gates when you reopen them.  Good luck folks, and thanks for the excellent write-up! 
1. [The Netflix Tech Blog: Automated Failure Testing](http://techblog.netflix.com/2016/01/automated-failure-testing.html?m=1)

    Netflix has taken the Chaos Monkey to the next level.  Now their automated system investigates the services a given request touches and injects artificial failures in various dependencies to see if they cause end-user errors.  It takes a lot of guts to decide that purposefully introducing user-facing failures is the best way to ultimately improve reliability.

### Outages

1. [Twitter](http://twitterstatus.tumblr.com/post/137610751178/service-issue)
    Twitter suffered a massive outage at least 2 hours long with sporadic availability for several hours after.  Hilariously, they posted status about the outage on Tumblr.
1. [Comcast (SF Bay area)](http://www.theregister.co.uk/2016/01/20/comcast_service_outage/)
1. [Africa](https://thestack.com/world/2016/01/22/civil-construction-wipes-out-internet-connectivity-across-africa/)
    This is the first time I’ve had an entire continent in this section. Most of Africa’s Internet was cut off from the rest of the world due to a pair of fiber cuts. South Africa was hit especially hard.

### [ << Prev ](sreweekly-6.md) ------------- [ Next >> ](sreweekly-8.md)