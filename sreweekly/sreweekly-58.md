## [SRE Weekly Issue #58](https://sreweekly.com/sre-weekly-issue-58/) - February 5, 2017
### Articles

1. [GitLab.com Database Incident](https://about.gitlab.com/2017/02/01/gitlab-dot-com-database-incident/)

    I’m going to break my normal format and post this outage up here in the article section. Here’s why: GitLab was extremely open about this incident, their incident response process, and even the actual incident response itself.Linked is their blog post about the incident, with an analysis from 24 hours after the incident that runs circles around the postmortems released by many other companies days after an outage. They also linked to their raw incident response notes (a Google Doc).Here’s what really blows me away: they live streamed their incident response on youtube. They’re also working on their postmortem document publicly in a merge request and tracking remediations publicly in their issue tracker. Incredible.Their openness is an inspiration to all of us. Here are a couple of snippets from the email I sent them earlier this week that is (understandably) still awaiting a response:
1. [New zine: “Networking! ACK!”](https://jvns.ca/blog/networking-zine-launch/)

    Julia Evans is back this week with a brand new zine about networking. It’ll be posted publicly in a couple weeks, but until then, you can get your own shiny copy just by donating to the ACLU (who have been doing a ton of awesome work!). Great idea, Julia!
1. [Google SRE Book is Now Free](https://landing.google.com/sre/book.html)

    You can now read the Google SRE book online for free! Pretty nifty. Thanks Google.
1. [The Infrastructure Behind Twitter: Scale](https://blog.twitter.com/2017/the-infrastructure-behind-twitter-scale)

    An in-depth dive into how Twitter scales. I’m somewhat surprised that they only moved off third-party hosting as recently as 2010. Huge thanks to Twitter for being so open about their scaling challenges and solutions.
1. [Understanding Unikernels](https://bsdmag.org/understanding_unikernels/)

    Here’s a good intro to unikernels, if you’re unfamiliar with them. The part that caught my attention is under the heading, “How Do You Debug the Result?”. I’m skeptical of the offered solution, “just log everything you need to debug any problem”. If that worked, I’d never need to pull out strace and lsof, yet I find myself using them fairly often.
1. [Leeds Teaching pathology IT crash blamed on “human error”](http://www.digitalhealth.net/infrastructure/48444/leeds-teaching-pathology-it-crash-blamed-on-)

    This article reads a whole lot more like “process problems” than “human error”. Gotta love the flashy headline, though.
1. [Are you being fooled when it comes to resilience?](https://thestack.com/data-centre/2017/01/23/are-you-being-fooled-when-it-comes-to-resilience-in-the-data-centre/)

    Just what exactly does that “five nines” figure in that vendor’s marketing brochures mean, anyway?
1. [Twitter Killed The Call Center](https://www.pagerduty.com/blog/twitter-killed-call-center/)

    Your customers may well take to Twitter to tell you (and everyone else) about your outages. PagerDuty shares suggestions for how to handle it and potentially turn it to your advantage.
1. [The Dark Standup](https://18f.gsa.gov/2017/01/19/the-dark-standup/)

    This operations team all agreed to work a strict 9-to-5 and avoid checking email or slack after hours. They shared their experience every day in a “dark standup” on Slack: a text-based report of whether each engineer is getting behind and what they’ll do with the extra hours they would normally have worked. They shared their conclusions in this article, and it’s an excellent read.
1. [A very short primer on DIY, technical debt and DevOps alerting](https://www.onpage.com/primer-on-diy-technical-debt-and-devops-alerting/)

    And the result is operational technical debt.
1. [No Easy Answers For Scheduling Physician On-Call Coverage](http://histalk2.com/2017/02/01/readers-write-no-easy-answers-for-scheduling-physician-on-call-coverage/)

    It’s really interesting to me that paying physicians extra for on-call shifts seems to be an industry standard. Of all my jobs, only one provided special compensation for on-call. It made the rather rough on-call work much more palatable. Does your company provide compensation or Time Off In Lieu (TOIL)? I’d love it if you’d write an article about the reasons behind the policy and the pros and cons!
1. [On-Call Ways and Means: A Developer’s Guide](https://victorops.com/blog/on-call-ways-means-developers-guide/)

    This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
1. [Booted up in 1993, this server still runs — but not for much longer](http://www.computerworld.com/article/3162416/data-center/booted-up-in-1993-this-server-still-runs-but-not-for-much-longer.html)

    Kind of neat, especially in this era of infrastructures built on (mostly) commodity hardware.
1. [Don’t let serverless applications dodge performance monitoring](http://searchitoperations.techtarget.com/tip/Dont-let-serverless-applications-dodge-performance-monitoring)

    The Emperor has no clothes on, NoOps isn’t a thing, and you still have to monitor your serverless applications. Sorry about that.
### Outages

1. [Telstra](http://www.businessinsider.com.au/a-fire-in-a-telstra-exchange-is-causing-flight-delays-and-network-outages-2017-2)
    A fire at an exchange resulted in an outage and somehow also caused SMSes to be mis-delivered to random recipients.
1. [Heroku](https://status.heroku.com/incidents/1042)
    Full disclosure: Heroku is my employer.
1. [Google App Engine](http://status.cloud.google.com/incident/appengine/17001#5743114304094208)
1. [123-Reg](https://www.theregister.co.uk/2017/02/03/123reg_hit_by_email_outage/)

### [ << Prev ](sreweekly-57.md) ------------- [ Next >> ](sreweekly-59.md)