## [SRE Weekly Issue #239](https://sreweekly.com/sre-weekly-issue-239/) - October 11, 2020
### Articles

1. [ Respect your natural scaling limits](https://ayende.com/blog/189537-C/respect-your-natural-scaling-limits)

    Don’t scale up farther than you need to! If you won’t ever see more than 100 RPS, don’t architect for 100,000.Ayende Rahien
1. [The Many Shapes of Site Reliability Engineering](https://medium.com/slalom-build/the-many-shapes-of-site-reliability-engineering-468359866517)

    This one covers several common patterns of SRE practice and then offers insight on what to look for as you design your own SRE team.Rob Cummings — Slalom Build
1. [Abstractions and implicit preconditions](https://surfingcomplexity.blog/2020/08/09/abstractions-and-implicit-preconditions/)

    Lorin Hochstein
1. [Keeping CALM: When Distributed Consistency Is Easy](https://cacm.acm.org/magazines/2020/9/246941-keeping-calm/fulltext)

    Coordination between nodes in a distributed system can kill performance. What kinds of problems require coordination? The CALM theorem can tell us.Joseph M. Hellerstein and Peter Alvaro — Communications of the ACM
1. [The Ultimate, Free Incident Retrospective Template](https://www.blameless.com/blog/incident-retrospective-postmortem-template)

    Here’s another good post-incident analysis document template that you can use as inspiration for your own.Hannah Culver — Blameless
1. [4 Signs Software Reliability Should be Your Top Priority](https://www.blameless.com/blog/4-signs-software-reliability-has-become-a-top-priority-for-your-company)

    As your product ages, it transitions from “cool new thing” to “tool everyone uses and expects to Just Work”. Your reliability needs will change accordingly.Lyon Wong — Blameless
### Outages

1. [PagerDuty](https://status.pagerduty.com/incidents/1f2tfsyt09sw)
    95% of event submissions (your systems telling PagerDuty to trigger an alert) failed for about an hour. They posted some detail about what went wrong.
1. [Slack](https://status.slack.com//2020-10/e8c094cc99aabf64)
    Their latest update on this outage contains some detail about what went wrong.
1. [Telegram](https://mspoweruser.com/telegram-apologises-for-yesterdays-downtime-reveals-cause/)
1. [Microsoft Office 365](https://www.crn.com/news/cloud/-very-frustrating-microsoft-office-365-outage-hits-u-s-again)
1. [Coles Supermarkets](https://www.9news.com.au/national/coles-supermarkets-australia-reports-it-issues-causing-cash-registers-to-stop/942891ec-6d7b-4176-bd95-d5bdbb4b0049)
1. [Adobe Creative Cloud](https://bestgamingpro.com/its-not-just-you-adobe-creative-cloud-is-down-right-now/)
1. [GitHub](https://www.githubstatus.com/incidents/28js6274ybpb)

### [ << Prev ](sreweekly-238.md) ------------- [ Next >> ](sreweekly-240.md)