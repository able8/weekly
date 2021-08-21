## [SRE Weekly Issue #43](https://sreweekly.com/sre-weekly-issue-43/) - October 9, 2016
### Articles

1. [Honeycomb and the Five Why’s](https://honeycomb.io/blog/2016/10/honeycomb-and-the-five-whys-summary-post/)

    A theme here in the past few issues has been the insane growth in complexity in our infrastructures. Honeycomb is a new tool-as-a-service to help you make sense of that complexity through event-based introspection. Think ELK or Splunk, but opinionated and way faster. The goal is to give you the ability to reach a state of flow in asking and answering questions about your infrastructure, so you can understand it more deeply, find problems you didn’t know you had, and discover new questions to ask. Here’s where I started getting really interested:
1. [On Finding Root Causes — Production Ready](https://medium.com/production-ready/on-finding-root-causes-c0ce524bf98b#.arr4i0lmz)

    Mathias Lafeldt rocks it again, this time with a great essay on finding root causes for an incident. I love the idea of using the term “Contributing Conditions” instead. And the Retrospective Prime Directive is so on-point I’ve gotta re-quote it here:
1. [Simple testing can prevent most critical failures](https://blog.acolyer.org/2016/10/06/simple-testing-can-prevent-most-critical-failures/)

    This paper review by The Morning Paper reminds us of the importance of checking return codes and properly handling errors. Best part: solid statistical evidence.
1. [1213486160 has a friend: 1195725856](https://rachelbythebay.com/w/2016/10/07/magic/)

    A followup note on Rachel Kroll’s hilarious and awesome story about 1213486160 (a.k.a. “HTTP”). Basically, if you see a weird number showing up in your logs, it might be a good idea to try interpreting it as a string!
1. [Netflix details chaos engineering](http://sdtimes.com/netflix-details-chaos-engineering/)

    A solid basic primer on Netflix’s chaos engineering tools, with some info about the history and motivation behind them. I love the bit about how they ran into issues when Chaos Monkey terminated itself. Oops.
1. [How to Handle an Outage Like a Pro](http://www.circleid.com/posts/20161006_how_to_handle_an_outage_like_a_pro/)

    This article should really be titled, Make Sure Your DNS Is Reliable! It’s easy to forget that all the HA in the world won’t help your infrastructure if the traffic never reaches it due to a DNS failure. And here’s a really good corollary:
1. [Exploring Airline Outages: A Developer’s Perspective](https://dzone.com/articles/exploring-airline-outages-a-developers-perspective)

    We’ve had a couple of high-profile airline computer system failures this year. Here’s an analysis of the difficulty companies are having bolting new functionality onto systems from the 90s and earlier, even as those systems try to support higher volume due to airline mergers. You may want to skip the bits toward the end that read like an ad, though.
1. [The Accidental DBA](https://charity.wtf/2016/10/02/the-accidental-dba/)

    I don’t think I’ve ever been at a company with a dedicated DBA role. It’s becoming a thing of the past, and instead ops folks (and increasingly developers) are becoming the new DBAs. Charity Majors tells us that we need to apply proper operational principals to our datastores. One change at a time, proper deploy and rollback plans, etc.
1. [GitHub – kamalmarhubi/shell-workshop](https://github.com/kamalmarhubi/shell-workshop)

    I love this idea: it’s an exercise in building your own command-line shell. It’s important to have a good grounding in the fundamentals of how processes get spawned and IO works in POSIX systems. Occasionally that’s the only way you can get to the root cause(s) of a really thorny incident.
### Outages

1. [Anik F2 (TV/telecom satellite)](http://www.ctvnews.ca/sci-tech/telesat-restores-anik-f2-satellite-seeks-cause-of-highly-unusual-17-hour-outage-1.3099242)
1. [eBay](http://www.ecommercebytes.com/C/abblog/blog.pl?/pl/2016/10/1475521357.html)
1. [GitHub](https://status.github.com/)
1. [Twilio](https://status.twilio.com/incidents/13dbtpmqqr51)
1. [National Australia Bank](http://www.news.com.au/finance/business/breaking-news/nab-apologises-for-national-system-outage/news-story/b697b9a19741d669093dcc5af6eab38a)
1. [Destiny (game)](https://gamerant.com/destiny-server-down-tapir/)
1. [Three (UK telecom)](http://uknip.co.uk/2016/10/major-mobile-phone-outage-on-three-mobile-network-effects-hundreds-of-south-coast-customers/)
1. [Level 3 and many major US telecomsmentioned](http://www.dailydot.com/layer8/phone-carrier-east-coast-down/)
    Level 3 mentioned a “configuration error”.
1. [Outlook.com](http://www.theinquirer.net/inquirer/news/2473416/microsoft-slow-to-respond-to-outlookcom-outage)
1. [PlayStation Network](http://www.express.co.uk/entertainment/gaming/718565/PlayStation-Network-down-PSN-servers-PS4-status)
1. [Verizon](http://www.13newsnow.com/news/local/north-carolina/verizon-wireless-outages-reported-in-northeastern-north-carolina/329655869)
1. [iTunes, App Store, and Apple Music](https://9to5mac.com/2016/10/03/itunes-app-store-and-apple-music-suffering-from-widespread-outage/)
1. [Netflix](http://www.upi.com/Entertainment_News/2016/10/01/Netflix-outage-leaves-viewers-stranded-on-Saturday/2211475359902/)
1. [Facebook](http://www.independent.co.uk/life-style/gadgets-and-tech/news/facebook-down-not-working-wont-load-outage-broken-is-it-up-a7343246.html)

### [ << Prev ](sreweekly-42.md) ------------- [ Next >> ](sreweekly-44.md)