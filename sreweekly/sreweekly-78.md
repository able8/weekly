## [SRE Weekly Issue #78](https://sreweekly.com/sre-weekly-issue-78/) - June 25, 2017
### Articles

1. [Language Bias in Accident Investigation](http://lup.lub.lu.se/student-papers/record/2971193)

    This Master’s thesis by Crista Vesel seeks to answer the question, “How does the language used in the U.S. Forest Service’s Serious Accident Investigation Guide bias accident investigation analysis?” It’s an awe-inspiring analysis, drawing on Dekker, Woods, Cook, and other authors I’ve linked here repeatedly.The most exciting part for me was the confirmation of some vague thoughts I’ve had around the use of passive versus active voice in retrospectives. By using passive voice, we can seek to reduce the kind of blaming that is inherent in active/agentive language.
1. [What can developers learn from being on call? – Julia Evans](https://jvns.ca/blog/2017/06/18/operate-your-software/?__s=bwykwk1kcceogszq8abt)

    It’s by Julia Evans. Just read it.
1. [Determining Alert Urgency](https://www.pagerduty.com/blog/determining-alert-urgency/)

    PagerDuty again draws on ITIL, this time to outline an example system for classifying incident impact and urgency in order to determine priority.
1. [ChaosCat: Automating Fault Injection at PagerDuty](https://www.pagerduty.com/blog/chaoscat-automating-fault-injection/)

    PagerDuty’s take on automating chaos includes a chat-bot that lets folks trigger one-off host failures, along with running periodically, of course.
1. [Reduce downtime with Azure Site Recovery service](http://searchwindowsserver.techtarget.com/tip/Reduce-downtime-with-Azure-Site-Recovery-service)

    This article is an overview of Microsoft’s DRaaS offering, Azure Site Recovery. Protip: you can just scroll past the signup-gate if you don’t feel like entering your email address.
1. [How We Scaled Our Cache and Got a Good Night’s Sleep](http://engineering.grab.com/how-we-scaled-our-cache-and-got-a-good-nights-sleep)

    Grab evaluated a couple of existing solutions but went with a simple custom sharding layer as a method to scale out their Redis usage.
### Outages

1. [Rollbar](http://status.rollbar.com/incidents/2mwkq3l71yx6)
1. [LinkedIn](http://www.express.co.uk/life-style/science-technology/819699/LinkedIn-down-not-working-UK-US)
1. [Skype](http://www.informationsecuritybuzz.com/expert-comments/microsoft-confirms-skype-outage-global-ddos-attack-suggested/)
    Suspected DDoS.
1. [ATO (Australian Tax Office)](http://www.sbs.com.au/news/article/2017/06/22/tax-office-website-goes-down-again)
1. [Dyn](https://www.dynstatus.com/incidents/1zb4n5qxvn51)
    Dyn suffered a long outage, and they posted an amazing 28 detailed updates to their status site before all was said and done. That’s something to aspire to.
1. [Heroku](https://status.heroku.com/incidents/1182)
    Heroku posted a followup for their series of incidents early this month. Sorry for missing posting those outages when they happened!Full disclosure: Heroku is my employer.

### [ << Prev ](sreweekly-77.md) ------------- [ Next >> ](sreweekly-79.md)