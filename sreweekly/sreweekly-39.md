## [SRE Weekly Issue #39](https://sreweekly.com/sre-weekly-issue-39/) - September 11, 2016
### Articles

1. [Who’s On Call?](http://www.susanjfowler.com/blog/2016/9/6/whos-on-call)

    A+ article! Susan Fowler has been a developer, an ops person, and now an SRE. That means she’s well-qualified to give an opinion on who should be on call, and she says that the answer is developers (in most cases). Bonus content includes “What does SRE become if developers are on call?”Thanks to Devops Weekly for this one.
1. [New zine: Linux debugging tools you’ll love](http://jvns.ca/blog/2016/09/07/new-zine-linux-debugging-tools-youll-love/)

    I figured this new zine from Julia Evans would be mostly review for me. WRONG. I’d never heard of dstat, opensnoop, or execsnoop, or perf before, but I sure will be using them now. As far as I can tell, Julia wants to learn literally everything, and better yet, she wants to teach us what she learned and how she learned it. Hats off to her.
1. [The ‘Change One Thing’ Rule](https://dzone.com/articles/the-change-one-thing-rule)

    This article argues that we should never do Y. If something goes wrong, we won’t know whether to roll back X or Y, and it’ll take twice as long to figure out which one is to blame.
1. [Systems blindness and how we deal with it](https://medium.com/production-ready/systems-blindness-and-how-we-deal-with-it-d601fa63b7f4#.7o87kxdu6)

    This week, Mathias introduces “system blindness”, the flawed understanding of how a system works and the lack of knowledge of how incomplete our understanding of it is. Whether we realize it or not, we struggle to mentally model the intricate interconnections in the increasingly complex systems we’re building.
1. [Building resilience in Spokes](http://githubengineering.com/building-resilience-in-spokes/)

    I’ve mentioned Spokes (formerly DGit) here previously. This time, GitHub shares the details on how they designed Spokes for high durability and availability.
1. [Ruby Is Dead! – You Need to Take Care of Its Memory Issues](https://dzone.com/articles/ruby-is-dead-you-need-to-take-care-of-its-memory-i)

    TIL: Ruby can suffer from Java-style stop-the-world garbage collection freezes.
1. [Facebook Engineers Crash Data Centers in Real-World Stress Test](http://spectrum.ieee.org/view-from-the-valley/computing/it/facebook-engineers-crash-data-centers-in-realworld-stress-test)

    Here’s recap of a talk about Facebook’s “Protect Storm”, given by VP Jay Parikh at @Scale. Project Storm involved retrofitting Facebook’s infrastructure time handle the failure of entire datacenters.
1. [Failure is Always An Option: How a Blameless Culture Leads to Better Results](https://victorops.com/blog/blameless-culture/)

    Here’s an interview with Jason Hand of VictorOps about the importance of a blameless culture. He mentions the idea that “Why?” is an inherently blameful kind of question (hat tip to John Allspaw’s Infinite “How?”s). I have to say that I’m not sure I agree with Jason’s other point that we shouldn’t bother attempting incident prevention, though. Just look at the work the aviation industry has done toward accident prevention.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [SCALE 15x Call for Proposals](https://www.socallinuxexpo.org/scale/15x/cfp)

    SCALE has opened their CFP, and one of the chairs told me that they’d “love to get SRE focused sessions on open-source.”
### Outages

1. [British Airways](http://www.bloomberg.com/news/articles/2016-09-06/iag-s-british-airways-says-computer-outage-is-delaying-flights)
1. [FLOW (Jamaica telecom)](http://go-jamaica.com/pressrelease/item.php?id=6916)
1. [SSP](http://www.theregister.co.uk/2016/09/07/ssp_decommissions_outage_hit_data_centre/)
    SSP provides a SaaS for insurance companies to run their business on. They’re dealing with a ten-plus-day outage initially caused by some kind of power issue that fried their SAN. As a result, they’re going to decommission the datacenter in question.
1. [Heroku](https://status.heroku.com/incidents/941)
    Full disclosure: Heroku is my employer.
1. [Azure](http://www.geekwire.com/2016/microsoft-azure-hit-widening-outages-europe-india/)
    Two EU regions went down simultaneously.
1. [Overwatch (game)](http://www.mobipicker.com/overwatch-free-weekend-news/)
1. [Asana](https://blog.asana.com/2016/09/yesterdays-outage/)
    Linked is a postmortem with an interesting set of root causes. A release went out that increased CPU usage, but it didn’t cause issues until peak traffic the next day. Asana is brave for enabling comments on their postmortem — not sure I’d have the stomach for that.Thanks to an anonymous contributor for this one.
1. [ESPN’s fantasy football](https://techcrunch.com/2016/09/11/espns-fantasy-football-site-isnt-working-on-the-most-important-day-of-the-season/)
    Unfortunate timing, being down on opening day.

### [ << Prev ](sreweekly-38.md) ------------- [ Next >> ](sreweekly-40.md)