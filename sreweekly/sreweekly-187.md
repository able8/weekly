## [SRE Weekly Issue #187](https://sreweekly.com/sre-weekly-issue-187/) - September 29, 2019
### Articles

1. [Atlassian Incident Management Handbook](https://www.atlassian.com/incident-management/get-the-handbook)

    I love it when companies publish their incident management documentation! Atlassian’s offering is high-quality — both in content and production value. The Major Incident Manager Cheatsheet at the end is worth distributing to your team.Atlassian
1. [Evolving Regional Evacuation](https://medium.com/netflix-techblog/evolving-regional-evacuation-69e6cc1d24c6?source=rss----2615bd06b42e---4)

    Netflix shares more about their N+1 AWS region redundancy design, and it all revolves around accurately modeling demand.Niosha Behnam — Netflix
1. [Simple is Complex](https://avdi.codes/simple-is-complex/)

    Interactions between simple microservices can lead to unexpected emergent behaviors.Avdi Grimm
1. [What Really Brought Down the Boeing 737 Max?](https://www.nytimes.com/2019/09/18/magazine/boeing-737-max-crashes.html)

    While I don’t necessarily agree with the blame-laden language of this article, it provides some interesting new details. It strikes me that, while MCAS may not be directly responsible for the crashes, it made it significantly harder to recover from contemporaneous pilot errors.William Langewiesche — The New York Times
1. [Observability — A 3-Year Retrospective](https://thenewstack.io/observability-a-3-year-retrospective/)

    My favorite part is the role-playing scenarios of debugging a problem with observability tooling and traditional tools.Charity Majors
1. [TCP: Out of Memory — Consider Tuning TCP_Mem](https://dzone.com/articles/tcp-out-of-memory-consider-tuning-tcp-mem)

    Tuning your TCP stack is important on busy servers.Ram Lakshmanan
### Outages

1. [Google Cloud Platform](https://status.cloud.google.com/incident/google-stackdriver/19007#19007003)
    This incident primarily affected the control plane of many GCP services. It stemmed from a cascading failure in an important key-value store used by all of them.
1. [Facebook and Instagram](https://economictimes.indiatimes.com/magazines/panache/a-friday-blackout-for-facebook-social-network-down-in-the-uk-instagram-suffers-outage-too/articleshow/71346312.cms)
1. [Google Maps](https://kymkemp.com/2019/09/28/google-maps-server-is-down/)
1. [GoDaddy](https://babblesports.com/tech/godaddy-down-most-of-the-users-from-united-states-hit-by-the-outage/)
1. [Target (retailer)](https://footwearnews.com/2019/business/retail/target-stores-website-app-down-technical-issues-outage-1202845470/)
1. [Discord](https://discord.statuspage.io/incidents/gtvzqy8b2plz)
1. [Fastlytwoothers](https://status.fastly.com/incidents/91rq8b4ncwbx)
    Plus two others.
1. [Squarespace](https://status.squarespace.com/incidents/40r1b6f2zks5)
1. [GitHub](https://www.githubstatus.com/incidents/b54zv9nz8m01)

### [ << Prev ](sreweekly-186.md) ------------- [ Next >> ](sreweekly-188.md)