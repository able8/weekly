## [SRE Weekly Issue #194](https://sreweekly.com/sre-weekly-issue-194/) - November 17, 2019
### Articles

1. [Errata for issue #193](http://square space)

    Last week, I mistakenly listed an outage as “Connectivity Issues”, when it should have been attributed to Squarespace. Sorry about that!
1. [Sleep, Interrupted: Niall Richard Murphy on Taking the Emergency Out of On-Call](https://www.getrevue.co/profile/incidentlabs/issues/sleep-interrupted-niall-richard-murphy-on-taking-the-emergency-out-of-on-call-210008)

    From the authors of the new Post-Incident Review Zine comes this summary of Niall Murphy’s SRECon talk. It’s way more than a talk recap, tying together related blog posts and talks from other authors.Jaime Woo and Emil Stolarsky
1. [The silence of the racks is deafening, production gear has gone dark – so which wire do we cut?](https://www.theregister.co.uk/2019/11/15/on_call/)

    They didn’t trust the datacenter’s backup power, so they added rack UPSes. Little did they realize that a single UPS failure could take down all of the rest.Richard Speed — The Register
1. [Taiji: managing global user traffic for large-scale Internet services at the edge](https://blog.acolyer.org/2019/11/15/facebook-taiji/)

    Taiji chooses which datacenter to route a Facebook user’s traffic to. It identifies clusters of users that have friended each other and routes them to the same place, on the theory that they’re likely to be interested in the same content.Adrian Colyer (summary)Xu et al., SOSP’19 (original paper)
1. [What Tracking Down Missing TCP Keepalives Taught Me About Docker, Golang, and GitLab](https://about.gitlab.com/blog/2019/11/15/tracking-down-missing-tcp-keepalives/)

    <3 detailed debugging stories. TIL: Google Compute Engine’s network drops connections from its state table after 10 minutes with no packets.Stan Hu — GitHub
1. [Monitoring server applications with Vortex](https://blogs.dropbox.com/tech/2019/11/monitoring-server-applications-with-vortex/)

    Vortex is DropBox’s custom-built metrics system, designed for horizontal scalability. Find out why they rolled their own and learn how it works in this article that includes shiny diagrams.Dave Zbarsky — DropBox
1. [Magic Numbers and second guessing SLOs – why is 96% better than 95%?](https://www.unixdaemon.net/sysadmin/magic-numbers-and-second-guessing-slos/)

    How do we come up with our SLOs, anyway? This one puts me in mind of Will Gallego’s post on error budgets.Dean Wilson (@unixdaemon)
1. [Snap: a microkernel approach to host networking](https://blog.acolyer.org/2019/11/11/snap-networking/)

    A network stack in userland as an alternative to TCP/IP? Yup, that seems like a pretty Google thing to do.Adrian Colyer (summary)Marty et al., SOSP’19 (original paper)
### Outages

1. [Disney+ (streaming service)](https://www.bloomberg.com/news/articles/2019-11-12/disney-ushers-in-new-era-as-netflix-rival-with-streaming-launch)
    Disney’s new streaming service suffered a few hiccups due to unexpectedly high demand.
1. [Codeanywhere](https://www.cbronline.com/news/codeanywhere-gcp-projects-missing)
1. [Google Nest](https://phandroid.com/2019/11/11/google-nest-services-are-down-customers-are-not-happy/)
1. [NFL Network (streaming service)](https://digistatement.com/nfl-network-not-working-nfl-network-down-streaming-not-working/)
1. [YouTube](https://www.independent.co.uk/life-style/gadgets-and-tech/news/youtube-down-not-working-loading-wrong-error-app-website-iphone-android-a9198121.html)
1. [Hulu](https://piunikaweb.com/2019/11/11/hulu-down-videos-reportedly-not-working-for-many/)
1. [Heroku: followup for incident #1922](https://status.heroku.com/incidents/1922)
1. [Heroku](https://status.heroku.com/incidents/1927)
1. [Transferwise](https://www.businesscloud.co.uk/news/transferwise-down-as-technical-issue-affects-app-and-website)
1. [Google Cloud Platform](https://status.cloud.google.com/incident/cloud-datastore/19006)
    A problem with KMS impacted multiple services in several regions. Google’s detailed followup analysis is linked above.

### [ << Prev ](sreweekly-193.md) ------------- [ Next >> ](sreweekly-195.md)