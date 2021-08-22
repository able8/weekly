## [DevOps'ish 023](https://devopsish.com/023) - Sun May 14, 2017

I have to admit, I am really, really angry this week. For the first time in a long time a worm was unleashed on the web. A little back story. Years ago, the NSA found a vulnerability in Windows. Instead of disclosing the vulnerability responsibility to Microsoft, the NSA decided to keep the vulnerability a secret. Years pass and NSA is happily using this zero-day to exploit the United States’ various enemies. Then, one day, several exploits are found on a server somewhere by <a href="https://en.wikipedia.org/wiki/The_Shadow_Brokers?ref=devopsish">The Shadow Brokers</a>. The Shadow Brokers then released the exploits they discovered.

Fast forward to this week and <a href="https://www.us-cert.gov/ncas/alerts/TA17-132A?ref=devopsish">the WannaCry exploit</a> is unleashed on the web. <a href="https://www.theguardian.com/society/live/2017/may/12/england-hospitals-cyber-attack-nhs-live-updates?ref=devopsish">The UK’s NHS was the first major victim</a>. Rapidly the <a href="http://money.cnn.com/2017/05/12/technology/ransomware-attack-nsa-microsoft/index.html?ref=devopsish">WannaCry tidal wave was washing over 99 countries</a>. Meanwhile, back in the UK, <a href="https://www.malwaretech.com/2017/05/how-to-accidentally-stop-a-global-cyber-attacks.html">a researcher discovers that there is a kill switch in the exploit</a>. Apparently, the NSA put a kill switch is in place in case the worm accidentally went public (WHICH IT DID). The kill switch, was a check to see if a specific domain existed; not responding with a 200, not having a specific payload or string, nothing! Just whether or not a domain was REGISTERED controlled the worm! The NSA didn’t think to spend the $11 to kill the worm. But, a 22-year-old in the UK saved billions of dollars and probably lives with $11 while <em>the NSA maintained its horrific negligence</em>. Unconscionable!

<a href="https://securitynewsletter.co/?utm_source=devopsish&amp;utm_medium=email&amp;utm_campaign=devopsish1"><strong>Security Newsletter: weekly digest of security news</strong></a><br/>Weekly e-mail that condenses last week’s security news into about 10 items worth knowing about. <em>SPONSORED</em>

<a href="http://www.brendangregg.com/blog/2017-05-09/cpu-utilization-is-wrong.html">CPU Utilization is Wrong</a> (Brendan Gregg)

<a href="https://github.com/hobby-kube/guide?ref=devopsish">hobby-kube</a>: A guide to Kubernetes clusters for the hobbyist

<a href="https://www.digitalocean.com/community/tutorials/how-to-host-a-website-with-caddy-on-ubuntu-16-04?ref=devopsish">How To Host a Website with Caddy on Ubuntu 16.04</a> (Digital Ocean)

Want a print copy of Daniel Stenberg’s <em>Everything curl</em>? <a href="https://daniel.haxx.se/blog/2017/05/10/everything-curl-printed/?ref=devopsish">Express your interest here</a>!

One of the first steps in DevOps is to understand the need for some failures and learn from them. <a href="http://www.informationweek.com/devops/why-failure-is-critical-to-devops-culture/a/d-id/1328830?ref=devopsish">Why Failure is Critical to DevOps Culture</a> by Jason Hand.

<a href="https://www.pagerduty.com/blog/modern-incident-resolution-lifecycle/?ref=devopsish">Announcing the Modern Incident Resolution Lifecycle</a> (PagerDuty)

<a href="https://ostif.org/the-openvpn-2-4-0-audit-by-ostif-and-quarkslab-results/?ref=devopsish">OpenVPN has been thoroughly audited and the discovered vulnerabilities have been patched</a>. It is good to see audits not finding horrific issues. It’s even more reassuring when it’s a project as relied upon as OpenVPN.

<a href="https://www.nytimes.com/2017/05/09/world/europe/hackers-came-but-the-french-were-prepared.html?ref=devopsish">Hackers Came, but the French Were Prepared</a> (New York Times)

This is too good not to share. Troy Hunt’s “<a href="https://www.troyhunt.com/reckon-youve-seen-some-stupid-security-things-here-hold-my-beer/">Reckon you’ve seen some stupid security things? Here, hold my beer…</a>” is a crying laughing/sad series of security missteps. What’s the worst security faux pas you’ve witnessed?

<a href="https://www.eff.org/deeplinks/2017/05/intels-management-engine-security-hazard-and-users-need-way-disable-it?ref=devopsish">Intel’s Management Engine is a security hazard, and users need a way to disable it</a> (EFF)

<a href="https://github.com/mrxinu/gosolar?ref=devopsish">GoSolar</a> is a SolarWinds client library written in Go. It allows you to submit queries to the SolarWinds Information Service (SWIS) and do various other things.

<a href="https://github.com/yandex/gixy?ref=devopsish">Gixy</a> is a tool to analyze Nginx configuration. The main goal of Gixy is to prevent security misconfiguration and automate flaw detection.

<a href="https://k6.io/">k6 is an open-source load testing tool for testing the performance of your systems</a>.

<a href="http://venturebeat.com/2017/05/10/microsoft-introduces-azure-cosmos-db-a-globally-distributed-database-with-5-consistency-choices/">Microsoft’s Build 2017 developer conference was this week and there were several DevOps worthy nuggets</a>:

Did everyone have a conference this week?!? OSCON, OpenStack Summit Boston, Microsoft Build, and Dell EMC World all took place this week. I saw the most activity around OSCON on Twitter. Surprisingly, Microsoft Build kept popping up in my Twitter stream frequently. As far as buzz, OSCON won the week though. <a href="https://www.oreilly.com/ideas/highlights-from-oscon-austin-2017?ref=devopsish">Highlights from OSCON are available</a>.

Coming up this week, I am co-hosting the <a href="https://tridevops.com/">Triangle DevOps</a> Meetup this Wednesday, May 17th. Nirmal Mehta will be presenting, “I Got 99 Problems But Technology Ain’t One” at 7 PM. <a href="https://www.meetup.com/Triangle-DevOps/events/238883192/?ref=devopsish">RSVP today if you are interested in going</a>.

I have been invited to speak at <a href="https://www.meetup.com/Open-Source-South-Carolina/events/239747095/?ref=devopsish">Open Source South Carolina</a> on Tuesday, May 23rd at 6:00 PM. I will be doing an extended DevOps 101 complete with deployment demonstration. If you are in or around Columbia, SC come join me!

<a href="https://www.devopsdays.org/events/2017-raleigh/welcome/?ref=devopsish">DevOpsDays Raleigh</a> is September 7th and 8th. The CFP closes May 15th!!! <a href="https://www.devopsdays.org/events/2017-raleigh/propose/?ref=devopsish">Propose your talk ASAP</a>! We are looking forward to seeing y’all this fall.

The other day my Vice President asked me a question, “<a href="https://chrisshort.net/sysadmin-to-devops-six-months/?ref=devopsish">What would it take to turn a good SysAdmin into a DevOps engineer?</a>” He followed it up with, “Don’t spend more than ten minutes thinking about it.” I spent WAY more than ten minutes on it.

The next volume in the <a href="https://opensource.com/open-organization">Open Organization</a> book series, “Guide to IT Culture Change: Open principles and practices for a more innovative IT department” comes out June 2nd. <a href="https://opensource.com/open-organization/resources/book-series?ref=devopsish">Register now</a> to be notified when it’s available.

GitHub has created a <a href="https://help.github.com/articles/adding-a-code-of-conduct-to-your-project/?ref=devopsish">Code of Conduct tool</a>. Adding a Code of Conduct to your project has never been easier.

<a href="https://charity.wtf/2017/05/11/the-engineer-manager-pendulum/?ref=devopsish">The Engineer/Manager Pendulum</a> (Charity Majors)

<a href="http://www.businessinsider.com/mathias-dopfner-tim-berners-lee-world-wide-web-interview-2017-5?ref=devopsish">The inventor of the web Tim Berners-Lee on the future of the internet, ‘fake news,’ and why net neutrality is so important</a> (Business Insider)

<a href="https://www.theguardian.com/technology/2017/may/09/what-is-net-neutrality-fcc-vote-why-it-matters?ref=devopsish">Net neutrality: why the next 10 days are so important in the fight for fair internet</a> (The Guardian)

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [CPU Utilization is Wrong](http://www.brendangregg.com/blog/2017-05-09/cpu-utilization-is-wrong.html)

    (Brendan Gregg)
1. [hobby-kube](https://github.com/hobby-kube/guide?ref=devopsish)

     A guide to Kubernetes clusters for the hobbyist
1. [How To Host a Website with Caddy on Ubuntu 16.04](https://www.digitalocean.com/community/tutorials/how-to-host-a-website-with-caddy-on-ubuntu-16-04?ref=devopsish)

    (Digital Ocean)
1. [Express your interest here](https://daniel.haxx.se/blog/2017/05/10/everything-curl-printed/?ref=devopsish)

    Want a print copy of Daniel Stenberg’s Everything curl? !
1. [Why Failure is Critical to DevOps Culture](http://www.informationweek.com/devops/why-failure-is-critical-to-devops-culture/a/d-id/1328830?ref=devopsish)

    One of the first steps in DevOps is to understand the need for some failures and learn from them.  by Jason Hand.
1. [Announcing the Modern Incident Resolution Lifecycle](https://www.pagerduty.com/blog/modern-incident-resolution-lifecycle/?ref=devopsish)

    (PagerDuty)
### Department of Data Defense

1. [OpenVPN has been thoroughly audited and the discovered vulnerabilities have been patched](https://ostif.org/the-openvpn-2-4-0-audit-by-ostif-and-quarkslab-results/?ref=devopsish)

    . It is good to see audits not finding horrific issues. It’s even more reassuring when it’s a project as relied upon as OpenVPN.
1. [Hackers Came, but the French Were Prepared](https://www.nytimes.com/2017/05/09/world/europe/hackers-came-but-the-french-were-prepared.html?ref=devopsish)

    (New York Times)
1. [Reckon you’ve seen some stupid security things? Here, hold my beer…](https://www.troyhunt.com/reckon-youve-seen-some-stupid-security-things-here-hold-my-beer/)

    This is too good not to share. Troy Hunt’s “” is a crying laughing/sad series of security missteps. What’s the worst security faux pas you’ve witnessed?
1. [Intel’s Management Engine is a security hazard, and users need a way to disable it](https://www.eff.org/deeplinks/2017/05/intels-management-engine-security-hazard-and-users-need-way-disable-it?ref=devopsish)

    (EFF)
### Department of Refreshment and Refurbishment

1. [GoSolar](https://github.com/mrxinu/gosolar?ref=devopsish)

    is a SolarWinds client library written in Go. It allows you to submit queries to the SolarWinds Information Service (SWIS) and do various other things.
1. [Gixy](https://github.com/yandex/gixy?ref=devopsish)

    is a tool to analyze Nginx configuration. The main goal of Gixy is to prevent security misconfiguration and automate flaw detection.
1. [k6 is an open-source load testing tool for testing the performance of your systems](https://k6.io/)

    .
1. [Microsoft’s Build 2017 developer conference was this week and there were several DevOps worthy nuggets](http://venturebeat.com/2017/05/10/microsoft-introduces-azure-cosmos-db-a-globally-distributed-database-with-5-consistency-choices/)

    
### Department of Discussion

1. [Highlights from OSCON are available](https://www.oreilly.com/ideas/highlights-from-oscon-austin-2017?ref=devopsish)

    Did everyone have a conference this week?!? OSCON, OpenStack Summit Boston, Microsoft Build, and Dell EMC World all took place this week. I saw the most activity around OSCON on Twitter. Surprisingly, Microsoft Build kept popping up in my Twitter stream frequently. As far as buzz, OSCON won the week though. .
1. [Triangle DevOpsRSVP today if you are interested in going](https://tridevops.com/)

    Coming up this week, I am co-hosting the  Meetup this Wednesday, May 17th. Nirmal Mehta will be presenting, “I Got 99 Problems But Technology Ain’t One” at 7 PM. .
1. [Open Source South Carolina](https://www.meetup.com/Open-Source-South-Carolina/events/239747095/?ref=devopsish)

    I have been invited to speak at  on Tuesday, May 23rd at 6:00 PM. I will be doing an extended DevOps 101 complete with deployment demonstration. If you are in or around Columbia, SC come join me!
1. [DevOpsDays RaleighPropose your talk ASAP](https://www.devopsdays.org/events/2017-raleigh/welcome/?ref=devopsish)

    is September 7th and 8th. The CFP closes May 15th!!! ! We are looking forward to seeing y’all this fall.
### Department of Interior

1. [What would it take to turn a good SysAdmin into a DevOps engineer?](https://chrisshort.net/sysadmin-to-devops-six-months/?ref=devopsish)

    The other day my Vice President asked me a question, “” He followed it up with, “Don’t spend more than ten minutes thinking about it.” I spent WAY more than ten minutes on it.
1. [Open OrganizationRegister now](https://opensource.com/open-organization)

    The next volume in the  book series, “Guide to IT Culture Change: Open principles and practices for a more innovative IT department” comes out June 2nd.  to be notified when it’s available.
### Department of Sane Workplaces

1. [Code of Conduct tool](https://help.github.com/articles/adding-a-code-of-conduct-to-your-project/?ref=devopsish)

    GitHub has created a . Adding a Code of Conduct to your project has never been easier.
1. [The Engineer/Manager Pendulum](https://charity.wtf/2017/05/11/the-engineer-manager-pendulum/?ref=devopsish)

    (Charity Majors)
### Not DevOps But Still Cool

1. [The inventor of the web Tim Berners-Lee on the future of the internet, ‘fake news,’ and why net neutrality is so important](http://www.businessinsider.com/mathias-dopfner-tim-berners-lee-world-wide-web-interview-2017-5?ref=devopsish)

    (Business Insider)
1. [Net neutrality: why the next 10 days are so important in the fight for fair internet](https://www.theguardian.com/technology/2017/may/09/what-is-net-neutrality-fcc-vote-why-it-matters?ref=devopsish)

    (The Guardian)

### [ << Prev ](sreweekly-22.md) ------------- [ Next >> ](sreweekly-24.md)