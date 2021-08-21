## [SRE Weekly Issue #235](https://sreweekly.com/sre-weekly-issue-235/) - September 13, 2020
### Articles

1. [Alerting on SLOs](https://mads-hartmann.com/sre/2020/09/08/alerting-on-slos.html)

    This isn’t just another boring article about SLOs. There’s a ton of good stuff in here about why they moved to SLO-based alerts, too.Mads Hartmann
1. [A nudge in the right direction](https://surfingcomplexity.blog/2020/08/09/a-nudge-in-the-right-direction/)

    Often, serendipity gets us out of an incident or makes it less severe.Lorin Hochstein
1. [Seamlessly Swapping the API backend of the Netflix Android app](https://netflixtechblog.com/seamlessly-swapping-the-api-backend-of-the-netflix-android-app-3d4317155187?source=rss----2615bd06b42e---4)

    It’s your classic “replace the engines on a jet while flying it” story. My favorite part is how they recorded real traffic and played it at the old and new backend API to compare the JSON responses.Rohan Dhruva and Ed Ballot — Netflix
1. [Using feature flags during incident management](https://launchdarkly.com/blog/using-feature-flags-during-incident-management/)

    Feature flags can help with load shedding and throttling, and feature flag activity can even be useful data that points to contributing factors. Dawn Parzych — LaunchDarkly
1. [Unimog – Cloudflare’s edge load balancer](https://blog.cloudflare.com/unimog-cloudflares-edge-load-balancer/)

    Unimog uses a lot of really interesting techniques to balance layer 4 traffic, about which this article goes into in great detail.David Wragg — Cloudflare
1. [Production testing with dark canaries](https://engineering.linkedin.com/blog/2020/production-testing-with-dark-canaries)

    I like this idea: it’s like a normal canary, except that you only send it a copy of traffic and discard the result, so as to avoid impacting users.David Hoa — LinkedIn
### Outages

1. [Google Drive](https://www.google.com/appsstatus#hl=en&v=issue&sid=4&iid=8a9eec54443fadd24be8227a6dee1220)
1. [GitHub](https://www.githubstatus.com/incidents/pbz6fh7mz86w)
1. [Facebook](https://finance.yahoo.com/news/facebook-reportedly-suffers-outage-parts-211847260.html)
1. [Snapchat](https://www.express.co.uk/life-style/science-technology/1266407/Snapchat-Down-iPhone-Android-Not-Working-Message-Failed-To-Send-Server-Status)
1. [Discord](https://discord.statuspage.io/incidents/ww1403twnxdw)

### [ << Prev ](sreweekly-234.md) ------------- [ Next >> ](sreweekly-236.md)