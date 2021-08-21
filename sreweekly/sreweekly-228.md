## [SRE Weekly Issue #228](https://sreweekly.com/sre-weekly-issue-228/) - July 19, 2020
### Articles

1. [Change Advisory Boards Don’t Work](https://octopus.com/blog/change-advisory-boards-dont-work)

    They don’t. They just don’t.Alex Yates — Octopus Deploy
1. [Google Cloud Issue Summary: Gmail 2020-06-30](https://static.googleusercontent.com/media/www.google.com/en//appsstatus/ir/l6t5py0jbfa7em8.pdf)

    Whoops, forgot to include this one last week.
1. [Postmortems and More With J. Paul Reed](https://www.pagerduty.com/blog/postmortems-more-2020/)

    My favorite part is the focus on blame awareness:Isabella Pontecorvo — PagerDuty
1. [Keeping Customers Streaming — The Centralized Site Reliability Practice at Netflix](https://netflixtechblog.com/keeping-customers-streaming-the-centralized-site-reliability-practice-at-netflix-205cc37aa9fb)

    Netflix has a team dedicated to the overall reliability of their service.Hank Jacobs– Netflix
1. [What is SRE?](https://www.firehydrant.io/blog/what-is-sre/)

    Another good reference if you’re looking to bootstrap SRE at your organization.Rich Burroughs — FireHydrant
1. [The Tail at Scale Approximation](https://billduncan.org/the-tail-at-scale-approximation/)

    Bill Duncan’s back with an easy and very close approximation for the “Tail at Scale” formula. The question it answers is: how many nines do you need on all of your backend microservices for X nines on the frontend?Bill Duncan
1. [The Essential List of Top SRE Resources](https://www.blameless.com/blog/top-sre-resources)

    Tons of great links in here with enticing descriptions to make you want to read them. Includes books, tools, hiring, certification, and general SRE goodness.Emily Arnot — Blameless
1. [Advocating for a Product Mindset within Platform Teams and How We Do it at HelloTech (Part 1)](https://engineering.hellofresh.com/advocating-for-a-product-mindset-within-platform-teams-and-how-we-do-it-at-hellotech-part-1-fc1fbf8ae015?source=rss----ac51da0a699a---4)

    SRE is all about keeping the user experience working, and working with product-focused folks can really help. For more on this, check out my former coworker Jen Wohlner’s awesome SRECon19 talk on SRE & product management.Samantha Coffman — HelloFresh
### Outages

1. [CloudflareDiscordPostmatesHosted GraphiteDownDetector](https://blog.cloudflare.com/cloudflare-outage-on-july-17-2020/)
    Cloudflare had a 50% drop in traffic served by their network subsequent to a BGP issue. Linked is their analysis including snippets of router configurations. Lots of services suffered contemporaneous outages possibly stemming from Cloudflare’s, including Discord, Postmates, Hosted Graphite, and DownDetector.John Graham-Cumming — Cloudflare
1. [Twitter](https://blog.twitter.com/en_us/topics/company/2020/an-update-on-our-security-incident.html)
    Twitter had a major security breach, and as part of their response, they temporarily cut off large parts of their service. Click for their post about what happened.
1. [GitHub](https://www.githubstatus.com/incidents/j597fw8kv04c)
1. [WhatsApp](https://www.sproutwired.com/whatsapp-experiences-worldwide-outage-with-1000s-impacted/)
1. [Hulu](https://babblesports.com/tech/hulu-down-audio-only-with-a-loading-bar/)
1. [Snapchat](https://www.express.co.uk/life-style/science-technology/1309028/Snapchat-down-black-screen-camera-not-working)
1. [Microsoft Outlook](https://www.computerworld.com/article/3567096/outlook-went-down-for-four-hours-wednesday-what-happened.html)
    Notably, the outage involved the Outlook application that people run on their computer, not the cloud version.
1. [Fastlycontrol plane incident](https://status.fastly.com/incidents/f7hg0w5c078c)
    Also a control plane incident later that day.Full disclosure: Fastly is my employer.

### [ << Prev ](sreweekly-227.md) ------------- [ Next >> ](sreweekly-229.md)