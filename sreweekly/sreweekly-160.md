## [SRE Weekly Issue #160](https://sreweekly.com/sre-weekly-issue-160/) - February 17, 2019
### Articles

1. [Operable Software](https://ferd.ca/operable-software.html)

    This is a long one, but trust me, it’s worth the read. My favorite part is where the author gets into mental models, hearkening back to the Stella Report.Fred Hebert
1. [Multi-CDN support in Mux for improved performance and reliability](https://mux.com/blog/multi-cdn-support-in-mux-video-for-improved-performance-and-reliability/)

    A multi-CDN approach can be tricky to pull off, but as these folks explain, it can be critical for reliability and performance.Scott Kidder — mUXFull disclosure: Fastly, my employer, is mentioned.
1. [Towards an understanding of technical debt](https://kellanem.com/notes/towards-an-understanding-of-technical-debt)

    This article explains five different phenomena that people mean when they say “technical debt”, and advocates understanding the full context rather than just assuming the folks that came before were fools./thanks Greg BurekKellan Elliott-McCrea
1. [How We Prepared New York Times Engineering for the Midterm Elections](https://open.nytimes.com/how-we-prepared-new-york-times-engineering-for-the-midterm-elections-2a615fe4196e)

    Kriton Dolias and Vinessa Wan — The New York Times
1. [@mipsytipsy on Twitter: what to alert on](https://twitter.com/mipsytipsy/status/1096278212257054720)

    There’s a real gem in here, definitely worth a read.Charity Majors (and Liz Fong-Jones in reply)
1. [Notes from On-call Adjacency – Honeycomb](https://www.honeycomb.io/blog/notes-from-on-call-adjacency/)

    Rachel Perkins — Honeycomb
1. [ How we used delayed replication for disaster recovery with PostgreSQL](https://about.gitlab.com/2019/02/13/delayed-replication-for-disaster-recovery-with-postgresql/)

    Andreas Brandl — GitLab
### Outages

1. [Azure Kubernetes Service (US East)](https://azure.microsoft.com/en-us/status/history/)
    There’s a pretty interesting incident description in their history page.
1. [VFEmail](https://twitter.com/VFEmail/status/1095038701665746945)
    Via Twitter:
At this time, the attacker has formatted all the disks on every server. Every VM is lost. Every file server is lost, every backup server is lost. NL was 100% hosted with a vastly smaller dataset. NL backups by the provideer were intact, and service should be up there.
My sympathies, folks.
1. [Slack](https://status.slack.com/2019-02/f0db7f15a8ffd08e)
    Emails into slack were failing due to an expired TLS certificate.
1. [Squarespace](https://status.squarespace.com/incidents/pbn8bwh6w8wk)
    Linked is their followup post explaining more about the incident.
1. [JPMorgan Chase](https://piunikaweb.com/2019/02/16/chase-bank-app-website-down-servers-not-working-online-and-mobile-banking-suffers/)
1. [Instagram](https://www.mirror.co.uk/tech/instagram-down-valentines-day-outage-14000720)
1. [Strava and Garmin Connect](https://piunikaweb.com/2019/02/17/strava-servers-down-and-not-working-users-reports-service-issues/)
1. [Microsoft Windows Update](https://support.microsoft.com/en-au/help/4464619/windows-10-update-history)
1. [Snapchat](https://www.express.co.uk/life-style/science-technology/1086314/Snapchat-down-server-status-latest-not-working)
1. [Sydney, AU Train Network](https://www.watoday.com.au/national/nsw/sydney-trains-digital-timetable-screens-hit-by-it-disruption-20190211-p50x45.html)
1. [Lloyds Bank](https://www.theweek.co.uk/banking/99518/lloyds-online-banking-down-bank-apologises-following-major-outage)

### [ << Prev ](sreweekly-159.md) ------------- [ Next >> ](sreweekly-161.md)