## [SRE Weekly Issue #167](https://sreweekly.com/sre-weekly-issue-167/) - April 7, 2019
### Articles

1. [Conference Report: SRECon Americas 2019](https://noidea.dog/blog/srecon-americas-2019)

    This is an awesome write-up of SRECon, but the part I really love is the intro. The author gives voice to a growing tension I’ve seen in our field, as we try to adopt the tenets of Safety II which can seem to be at odds with traditional SRE practices. There’s a lot here that we SREs need to work out as our profession matures, and I’m really enjoying the process.Tanya Reilly
1. [Control Is an Illusion](https://ryanfrantz.com/posts/control-is-an-illusion.html)

    Experts recommend trying to keep the concepts of blame, root cause, and hindsight bias out of our retrospective investigations. This insightful article explains that they all stem from the illusion that we are in full control of our systems.Thanks to Will Gallego for this one.Ryan Frantz
1. [What We Learned from the Recent Mandrill Outage](https://mailchimp.com/what-we-learned-from-the-recent-mandrill-outage/)

    Here’s a top-notch followup analysis from Mailchimp on the Mandrill outage last month. Their Postgresql DB ran out of transaction IDs (a common failure mode), causing a painful outage. Tons of great stuff here including a mention of rotating ICs every 3 hours to prevent exhaustion and allow them to sleep.Mailchimp
1. [Ethiopian crash: Boeing 737 Max pilots followed expected procedures, aviation officials say](https://www.cnn.com/2019/04/04/us/ethiopian-airlines-737-max-crash-preliminary-report/index.html)

    And here’s where things get really interesting. Incidents are never as simple as they seem from the outside, and the 737 MAX situation is no exception. I anxiously await the full report, in which we’ll hear more about the confluence of contributing factors that must have been involved here.Thom Patterson — CNN
1. [The true story behind the deadliest air disaster of all time](https://www.telegraph.co.uk/travel/comment/tenerife-airport-disaster/)

    There’s a lot in this, and I don’t feel comfortable summarizing it with a little blurb about lessons learned. Chilling though it is, I’m glad I read it.Thanks to Sri Ray for this one.Patrick Smith — The Telegraph
1. [ Production ready code is much more than error handling – Ayende @ Rahien ](https://ayende.com/blog/186849-A/production-ready-code-is-much-more-than-error-handling)

    Ayende Rahien
### Outages

1. [Travis CI](https://www.traviscistatus.com/incidents/my5wm56npf7q)
1. [Slackthis one](https://status.slack.com/2019-03/146118a2b0b98ec6)
    And this one.
1. [Google Cloud Platform (us-central1)](https://status.cloud.google.com/incident/cloud-networking/19007)
1. [Heroku](http://feedproxy.google.com/~r/HerokuStatus/~3/5fypaFAqtIo/1762)
1. [Instagram](http://www.appocalypse.co/technology/instagram-down-again/)
1. [Squarespace](https://status.squarespace.com/incidents/w8tvp6s2x5fr)
    Click for another A+ followup analysis from Squarespace. Thanks, folks!

### [ << Prev ](sreweekly-166.md) ------------- [ Next >> ](sreweekly-168.md)