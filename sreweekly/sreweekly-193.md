## [SRE Weekly Issue #193](https://sreweekly.com/sre-weekly-issue-193/) - November 10, 2019
### Articles

1. [The Consul outage that never happened](https://about.gitlab.com/blog/2019/11/08/the-consul-outage-that-never-happened/)

    Ever had a Sev 1 non-impacting incident? This team’s Consul cluster was balanced on a razor’s edge: one false move and quorum would be lost. Read about their incident response and learn how they avoided customer impact.Devin Sylva — GitLab
1. [Hot SRE trends in 2019 (brought to you from SREcon EMEA)](https://making.pusher.com/hot-sre-trends-in-2019/)

    This SRECon EMEA highlight reel is giving me serious FOMO.Will Sewell — Pusher
1. [Resilience Roundup – Handoff strategies in settings with high consequences for failure: lessons for health care operations](https://resilienceroundup.us7.list-manage.com/track/click?e=c7bb46d5fc&id=75ae885324&u=3195f1d3ece4512b9491eb783)

    Emily Patterson, Emilie Roth, David Woods, and Renee Chow (original paper)Thai Wood (summary)
1. [Scaling in the presence of errors—don’t ignore them](https://programmingisterrible.com/post/188942142748)

    This is an interesting essay on handling errors in complex systems.tef
1. [Fast dimensional analysis for root cause analysis at scale](https://engineering.fb.com/developer-tools/fast-dimensional-analysis/)

    To be clear: this is about assisting incident responders in gaining an understanding of an incident in the moment, not about finding a “root cause” to present in an after-action report.I’m not going to pretend to understand the math, but the concept is intriguing.Nikolay Pavlovich Laptev, Fred Lin, Keyur Muzumdar, Mihai-Valentin Curelea, Seunghak Lee, and Sriram Sankar — Facebook
1. [The inflection point hypothesis: a principled approach to finding the root cause of a failure](https://blog.acolyer.org/2019/11/08/the-inflection-point-hypothesis/)

    This one’s about assisting humans in debugging, when they have a reproduction case for a bug but can’t see what’s actually going wrong.That’s two different uses of “root cause” this week, and neither one is the troublesome variety that John Allspaw has debunked repeatedly.Zhang et al. (original paper)Adrian Colyer (summary)
### Outages

1. [HoneycombHere](http://status.honeycomb.io/incidents/l0m559jd5xkv)
    Here‘s an unroll of an interesting Twitter thread by Honeycomb’s Liz Fong-Jones during and after the incident.
1. [GitHub](https://www.githubstatus.com/incidents/42hkbtl63nmn)
1. [Amazon Prime Video](https://www.trustedreviews.com/news/amazon-prime-video-streaming-service-wont-let-stream-anything-3952780)
1. [Google Compute Engine](http://status.cloud.google.com/incident/compute/19008#19008020)
    Network administration functions were impacted. Click for their post-incident analysis.
1. [Squarespace](https://status.squarespace.com/incidents/19154mxssb79)
    On Wednesday November 6th, many Squarespace websites were unavailable for 102 minutes between 14:13 and 15:55 ET.
Click through for their post-incident analysis.

### [ << Prev ](sreweekly-192.md) ------------- [ Next >> ](sreweekly-194.md)