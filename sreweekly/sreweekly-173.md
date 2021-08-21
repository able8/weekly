## [SRE Weekly Issue #173](https://sreweekly.com/sre-weekly-issue-173/) - June 16, 2019
### Articles

1. [The Negotiability of “Severity” Levels](https://www.adaptivecapacitylabs.com/blog/2019/05/20/the-negotiability-of-severity-levels/)

    So many awesome concepts packed into this article. Here are just a couple:John Allspaw — Adaptive Capacity Labs
1. [How Webhook.site handles 100 mbit/s traffic on a single VPS](https://simonfredsted.com/1721)

    This shares more in common with the server behind sreweekly.com than I perhaps ought to admit to:Simon Fredsted
1. [PostgreSQL: pg_upgrade can result in early wraparound on databases with high
transaction load](https://www.postgresql.org/message-id/CALSof1GM6i21BLr8PsFiRYazakojvesdc+_MiR-L_V5NNkRuWg@mail.gmail.com)

    A Reddit engineer explains a hidden gotcha of pg_upgrade that caused an outage I reported here previously.Jason Harvey — Reddit
1. [Pilots at MIA’s Biggest Cargo Airline Warned Execs a Crash Was Coming. Then a Plane Went Down.](https://www.aviationpros.com/airlines/news/21084636/pilots-at-mias-biggest-cargo-airline-warned-execs-a-crash-was-coming-then-a-plane-went-down)

    This has “normalization of deviance” all over it.Taylor Dolven — The Miami Herald
1. [Boeing Built Deadly Assumptions Into 737 Max, Blind to a Late Design Change](https://www.nytimes.com/2019/06/01/business/boeing-737-max-crash.html)

    The deep details around MCAS are starting to come out. This article tells a tale that is all too familiar to me about organizational pressures and compartmentalization.Jack Nicas, David Gelles and James Glanz — New York Times
### Outages

1. [Google](https://cloud.google.com/blog/topics/inside-google-cloud/an-update-on-sundays-service-disruption)
    Click through for Google’s blog post about the outage that impacted Google Cloud Platform, YouTube, Gmail, Google Drive.A configuration change intended for a small number of servers was incorrectly applied more broadly, causing reduced network capacity. The similarity to the second Heroku outage below is striking.
1. [Heroku Incident #1776 Follow-up](https://status.heroku.com/incidents/1776)
    An expired SSL certificate caused control plane impact and some impact to running applications.
1. [Heroku Incident #1789 Follow-up](https://status.heroku.com/incidents/1789)
    A configuration change intended for a testing environment was mistakenly applied to production, resulting in 100% of requests in the EU failing.

### [ << Prev ](sreweekly-172.md) ------------- [ Next >> ](sreweekly-174.md)