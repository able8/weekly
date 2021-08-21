## [SRE Weekly Issue #114](https://sreweekly.com/sre-weekly-issue-114/) - March 18, 2018
### Articles

1. [Level 3 technician’s misstep causes largest outage ever reported](https://www.fiercetelecom.com/telecom/fcc-finally-specifies-cause-2016-level-3-network-outage)

    The FCC has released a report on the major Level 3 outage in October of 2016. This summary article serves as a good TL;DR summary on what went wrong and includes a link to the full report.Brian Santo — Fierce Telecom
1. [Migrating edge network providers](https://webuild.envato.com/blog/migrating-edge-providers/)

    They had an awesome approach: use RSpec to create a test suite of HTTP requests and run it continuously during the deployment to ensure that nothing changed from the end-user’s perspective. Bonus points for generating tests automatically. Jacob Bednarz — Envato
1. [Project Nimble: Region Evacuation Reimagined](https://medium.com/netflix-techblog/project-nimble-region-evacuation-reimagined-d0d0568254d4)

    Netflix reduced the time it takes to evacuate a failed AWS region from 50 minutes to just 8.Luke Kosewski, Amjith Ramanujam, Niosha Behnam, Aaron Blohowiak, and Katharina Probst — Netflix
1. [Tonight We Monitor, For Tomorrow, We Test in Production!](http://blog.launchdarkly.com/tonight-we-monitor-for-tomorrow-we-test-in-production/)

    I don’t usually link to talks, but this talk transcript reads almost like an article, and it’s a good one. The premise: if you’re not monitoring well, then you can’t safely test in production. Scalyr found a few ways in which their monitoring showed cracks, and now they’re sharing it with us.Steven Czerwinski — Scalyr
1. [Oopsy DDoSy: Accidental DDoS Attacks Causing Major Grief](http://www.businesscomputingworld.co.uk/oopsy-ddosy-accidental-ddos-attacks-causing-major-grief/)

    Design carefully, especially around retries, lest you create a thundering herd that makes it much harder to recover from an outage. That lesson and more, in this article on shooting yourself in the foot at web scale.Benjamin Campbell — Business Computing World
1. [How our production team runs the weekly on-call handover](https://about.gitlab.com/2018/03/14/the-on-call-handover-at-gitlab/)

    Have I mentioned how much I love GitLab’s openness? Here’s how they handle on-call shift transitions in their remote-only organization.John Jarvis — GitLab
1. [Twitter: Charity Majors on distributed systems, complexity, and microservices](https://twitter.com/mipsytipsy/status/973806469136121856)

    What is the definition of a distributed system, and why are they difficult? I really love the definition in the second tweet.Charity Majors
1. [Troubleshooting IPv6 badness to certain hosts in a rack](https://rachelbythebay.com/w/2018/03/16/slowroad/)

    I sure love a good troubleshooting story. This one has a pretty excellent failure mode, A+ investigative technique, and an emphasis on following something through until you find an answer.Rachel Kroll
1. [The Makeup of Successful Geographically-Distributed SRE Teams: Part 1 | LinkedIn Engineering](https://engineering.linkedin.com/blog/2018/03/the-makeup-of-successful-geographically-distributed-sre-teams--p)

    This discussion of how and why to create a globally-distributed SRE team may only apply to bigger companies, but it’s got a lot of useful bits in it. I just have to stop laughing at the acronym “GD”…Akhil Ahuja — LinkedIn
### Outages

1. [DoubleClick (ad provider)](http://blog.catchpoint.com/2018/03/14/doubleclick-outage-another-lesson-third-party-optimization/)
    DoubleClick went down, and it took a lot of sites with it. Click through for Catchpoint’s excellent analysis.
Kameerath Kareem — Catchpoint
1. [Travis CI](https://www.traviscistatus.com/incidents/z2b3lz2kwcfp)
1. [SmartThings (IoT platform)](https://www.cnet.com/news/smartthings-outage-hits-north-american-users/)
1. [Air Canada](https://www.wingsjournal.com/air-canada-suffers-major-outage-worst-time)
1. [Netflix](https://www.thesun.co.uk/news/5813045/netflix-down-uk-what-nw-2-5-error-streaming-service-working-again/)

### [ << Prev ](sreweekly-113.md) ------------- [ Next >> ](sreweekly-115.md)