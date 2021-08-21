## [SRE Weekly Issue #70](https://sreweekly.com/sre-weekly-issue-70/) - April 30, 2017
### Articles

1. [Enabling DNS split authority with OctoDNS](https://githubengineering.com/enabling-split-authority-dns-with-octodns/)

    GitHub has released OctoDNS, their tool for synchronizing DNS across multiple providers. Shortly after the Dyn outage last fall (covered here), they still only had one DNS provider (source: direct observation). I suspected that this may have had to do with complication in keeping records synched across two providers – perhaps that’s why they created OctoDNS.
1. [Introducing Bolt: On Instance Diagnostic and Remediation Platform](http://techblog.netflix.com/2017/04/introducing-bolt-on-instance-diagnostic.html)

    Bolt is Netflix’s “event driven diagnostic and remediation platform”, although it actually seems like a full-blown remote execution system for large fleets of servers.
1. [Incident management at Google — adventures in SRE-land](https://cloudplatform.googleblog.com/2017/02/Incident-management-at-Google-adventures-in-SRE-land.html?m=1)

    A Google SRE takes us through their first on-call shift including running incident command for a production incident. I like the emphasis on a blameless postmortem.
1. [Onboarding, On-Call and Learning](https://medium.com/@petey5000/onboarding-on-call-and-learning-46ff297a2587)

    Pete Shima received some questions about onboarding SREs, and lucky us, he decided to answer them publicly. My favorite section is the one about connecting a new SRE to people across the company. I find that solid connections to folks in various positions are vital to getting my job done well. Thanks to Pete for the SRE Weekly mention!
1. [Take A Moment To Refocus – Salesforce + Open Source = ❤](https://medium.com/salesforce-open-source/take-a-moment-to-refocus-86b6546c90c)

    Salesforce has a humongous infrastructure, and they needed a tool to help visualize data from lots of monitoring systems. They created Refocus to serve that need, and they open sourced it. They had three goals: gather data from all of the monitoring systems, on-board new services quickly, and visualize data in a way that makes sense for each service.Full disclosure: Salesforce (parent company of my employer, Heroku), is mentioned.
1. [New zine: let’s learn tcpdump!](https://jvns.ca/blog/2017/04/29/new-zine--let-s-learn-tcpdump/?source=sreweekly)

    Tcpdump is a critical tool for debugging thorny network issues. Julia Evans created a new zine to help you learn the basics, although if her other zines are any indication, even a pro may learn a new trick or two. The zine is $10 now and will be available for free at some point in the future.
1. [Sharks Want to Bite Google’s Undersea Cables](https://www.wired.com/2014/08/shark_cable/)

    Turns out that sharks are a reliability risk. And not just those WFLB.
1. [Why Code Gets Released too Early (and how to fix It)](https://about.gitlab.com/2017/04/27/why-code-is-released-too-early/)

    From their Global Developer Survey, GitLab learned that it’s common for developers to release code before it’s production-ready in response to organizational pressures.
1. [Four nines & failure rates – will the cloud ever cut it for transactional banking?](https://www.wandisco.com/news-events/blog/four-nines-failure-rates-will-the-cloud-ever-cut-it-for-transactional-banking-)

    Here’s a pretty excellent analysis of why adopting the cloud can be difficult for banks. Just skip past the bit with the incorrect uptime calculation, since four nines of uptime actually equates to about 53 minutes’ downtime per year, not 9 hours.
### Outages

1. [London Marathon Donations](http://www.thecsuite.co.uk/cio/strategy-cio/not-all-websites-are-created-equal-learning-from-the-london-marathon-let-down/)
    Ebay and Virgin Money Giving both went down under the load as many flocked to place donations before the London Marathon.
1. [CARLI](https://www.carli.illinois.edu/network-outage-today-april-28)
    CARLI is the Consortium of Academic Research Libraries in Illinois. I included this outage because of the short but sweetly personal postmortem from their network engineer.
1. [Instagram](https://techcrunch.com/2017/04/24/instagram-down/)
1. [Reddit](http://www.redditstatus.com/incidents/x0svvzdrsn5m)
    Sorry for the extended outage there. We failed back the maintenance performed earlier tonight. We’ll provide a post-mortem at a later date.

### [ << Prev ](sreweekly-69.md) ------------- [ Next >> ](sreweekly-71.md)