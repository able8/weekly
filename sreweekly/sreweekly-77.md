## [SRE Weekly Issue #77](https://sreweekly.com/sre-weekly-issue-77/) - June 18, 2017
### Articles

1. [Systemic brittleness, reactions to failure, and Conroy’s Law](https://www.snafucatchers.com/single-post/2017/06/04/BorkedTheDatabaseCase)

    Last week, I linked to a reddit story of an engineer that was unfairly fired for a mistake on their first day. Dr. Richard Cook picked this up and wrote up a great analysis of the underlying organizational issues.Thanks to John Allspaw for this one.
1. [Australian Tax Office’s post-incident report on the SAN outages](https://www.ato.gov.au/uploadedFiles/Content/CR/downloads/js39322_ATO-systems-report_w.pdf)

    This was released the week before last, but it took me awhile to digest it. The ATO did a very thorough post-analysis on their two outages and released this polished report. I like that they took full responsibility for the outage even though it was an issue with a fully-managed vendor SAN offering, and they clearly sought to learn as much as possible.
1. [Applications of (pin)trace data](http://t.dripemail2.com/c/eyJhY2NvdW50X2lkIjoiMjY1Njc0MyIsImRlbGl2ZXJ5X2lkIjoiOTM2MDMwNjYxIiwidXJsIjoiaHR0cHM6Ly9oYWNrZXJub29uLmNvbS9hcHBsaWNhdGlvbnMtb2YtcGluLXRyYWNlLWRhdGEtM2I5ZTZkYzI3NDRiP19fcz1id3lrd2sxa2NjZW9nc3pxOGFidCJ9)

    Pinterest tech lead Suman Karumuri explains how they use distributed tracing and the benefits it’s brought them.
1. [An Imaginary Apology Letter From Your Airline CEO](https://dzone.com/articles/an-imaginary-apology-letter-from-your-airline-ceo)

    Frustrated by British Airways’s Willie Walsh’s public statement regarding their major outage, TripWire founder Gene Kim took it upon himself to write an open letter of apology as if he were an airline CEO.  It’s pretty great.
1. [NGINX Plus High Availability on AWS](https://www.nginx.com/blog/nginx-plus-highly-available-aws-load-balancer/)

    This article explores several options for HA with Nginx: put an ELB in front of it, Route 53 with health checks, or an elastic IP switched either by keepalived or a Lambda function.
1. [On-Calliday: A Guide to Unsucking Your On-Call Experience](https://about.gitlab.com/2017/06/14/on-calliday-unsucking-your-on-call-experience/)

    I’ve been following GitLab’s blog since their engineer accidentally deleted their database earlier this year, and I’m glad I did. This article touches on all sorts of topics near to my heart: preventing burnout, examining incident response metrics, enforcing vacations, incident command, and having developers go on-call for what they wrote.
1. [The hidden cost of “Dark DR:” The economic argument for active/active operations](http://www.itproportal.com/features/the-hidden-cost-of-dark-dr-the-economic-argument-for-activeactive-operations/)

    
1. [A/B Testing and Beyond: Improving the Netflix Streaming Experience with Experimentation and Data Science](https://medium.com/netflix-techblog/a-b-testing-and-beyond-improving-the-netflix-streaming-experience-with-experimentation-and-data-5b0ae9295bdf?source=rss----2615bd06b42e---4)

    Netflix explains in depth the careful scientific experiments they perform in production in order to improve the QoE (quality of experience).
### Outages

1. [Google Cloud Services](http://status.cloud.google.com/incident/appengine/17006#5737979670691840)
    62-minute multiple-zone total internet outage in asia-northeast1. Postmortem linked, including a description of several contributing factors.
We apologize for the impact this issue had on our customers, and especially to those customers with deployments across multiple zones in the asia-northeast1 region. We recognize we failed to deliver the regional reliability that multiple zones are meant to achieve.
1. [Coinbase](http://www.pymnts.com/news/bitcoin-tracker/2017/high-bitcoin-trading-volume-causes-coinbase-outage/)
1. [YouTube](https://www.thesun.co.uk/tech/3816099/is-youtube-down-east-coast-of-america-hit-by-major-outage-as-video-site-goes-down/)

### [ << Prev ](sreweekly-76.md) ------------- [ Next >> ](sreweekly-78.md)