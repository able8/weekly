## [SRE Weekly Issue #210](https://sreweekly.com/sre-weekly-issue-210/) - March 8, 2020
### Articles

1. [Introducing Dispatch](https://netflixtechblog.com/introducing-dispatch-da4b8a2a8072?source=rss----2615bd06b42e---4)

    Netflix open sourced their incident management system.Kevin Glisson, Marc Vilanova, Forest Monsen — Netflix
1. [Reading /proc/pid/cmdline can hang forever](https://rachelbythebay.com/w/2014/10/27/ps/)

    I wasn’t aware of this little pitfall of memory cgroups.rachelbythebay
1. [In space, no one can hear you kernel panic](https://increment.com/software-architecture/in-space-no-one-can-hear-you-kernel-panic/)

    Your failover DB instance is cute. Try 4x+ redundancy. That’s the kind of engineering required when designing systems to operate in space.Glenn Fleishman — Increment
1. [A single person on-call “rotation” is a critical vulnerability](https://www.firehydrant.io/blog/a-single-person-on-call-rotation-is-a-critical-vulnerability/)

    Daniel Condomitti — FireHydrant
1. [Experimental study on the effect of procedure under unexpected situations](https://resilienceroundup.com/issues/69/)

    This is a pretty nifty experiment showing the importance of letting folks use their judgement to handle unexpected situations rather than relying on adherence to procedures.Thai Wood — Resilience Roundup (summary)Makoto Takahashi, Daisuke Karikawa, Genta Sawasato and Yoshitaka Hoshii — Tohoku University (original paper)
1. [Coronavirus/COVID-19 and USENIX Conferences](https://www.usenix.org/conferences/coronavirus)

    FYI: SRECon Americas West has been rescheduled to June 2-4.
1. [Millions of tiny databases](https://blog.acolyer.org/2020/03/04/millions-of-tiny-databases/)

    This week, we have another summary of the Physalia paper. I especially like the bit about poison pills.Adrian Colyer — The Morning Paper (summary)Brooker et al. — NSDI’20 (original paper)
1. [How did software get so reliable without proof?](https://surfingcomplexity.blog/2020/03/02/how-did-software-get-so-reliable-without-proof/)

    In this case, “proof” means “formal proof”.Lorin Hochstein
### Outages

1. [Let’s Encrypt Statusinitial reportfull reportrevoke 3 million certificatesstatuspage.ioHeroku](https://status.io/pages/incident/55957a99e800baa4470002da/5e59d5d7a9ba9d04c65eb77f)
    Let’s Encrypt purposefully suspended certificate issuance to investigate a bug around validating CAA DNS records. See their initial report and subsequent full report for details.
Subsequently, they decided to revoke 3 million certificates with a pretty short warning. Both actions (the revocations and taking down issuance initially) were likely warranted and mandated under the compliance guidelines that CAs are subjected to.
I’ve found two third-party incidents so far that seem to stem from the revocations:
* statuspage.io
* Heroku
Got any more? Please do send them my way.
1. [Robinhood (stock trading platform)](https://status.robinhood.com/)
    Thanks to Daniel Lucas for this and a couple other recent ones. for this one.
1. [G Suite Status Dashboard](https://www.google.com/appsstatus#hl=en&v=issue&sid=1&iid=846ac4f49df126e45a93590496285ed6)
1. [PagerDuty](https://status.pagerduty.com/incidents/xhygfw6fczd2)
1. [Uber](https://www.theregister.co.uk/2020/03/05/uber_down/)
1. [Interactive Brokers (Stock Broker)](https://www.financemagnates.com/forex/brokers/interactive-brokers-experiences-technical-issues-website-crashes/)
1. [Binion’s and Four Queens (Las Vegas casinos)](https://www.casino.org/news/four-queens-binions-back-up-after-slots-outage-other-systems-affected/)
    Slot machines stopped working, and an eerie quiet descended.
1. [crates.io incident report for 2020-02-20](https://blog.rust-lang.org/inside-rust/2020/02/26/crates-io-incident-report.html)
    On 2020-02-20 at 21:28 UTC we received a report from a user of crates.io that their crate was not available on the index even after 10 minutes since the upload. This was a bug in the crates.io webapp exposed by a GitHub outage.
crates.io is the Rust language package registry.
Pietro Albini — crates.io
1. [Discord](https://discord.statuspage.io/incidents/8lqdtgs99vbs)

### [ << Prev ](sreweekly-209.md) ------------- [ Next >> ](sreweekly-211.md)