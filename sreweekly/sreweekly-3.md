## [SRE Weekly Issue #3](https://sreweekly.com/sre-weekly-issue-3/) - December 26, 2015
### Articles

1. [sysadvent: Day 22 – Simplicity in Complex Systems](http://feedproxy.google.com/~r/sysadvent/~3/bBemdgJyC9k/day-22-simplicity-in-complex-systems.html)

    I love this article! Simplicity is especially important to us in SRE, as we try to gain a holistic understanding of the service in order to understand its failure modes. Complexity is the enemy. Every bit of complexity can hide a design weakness that threatens to take down your service. I once heard that it takes Google SRE hires 6 months to come up to speed. I get that Google is huge, but is it possible to reduce this kind of spin-up time by aggressively simplifying?
1. [On Call for the Holidays](https://m.reddit.com/r/sysadmin/comments/2dj3ol/having_to_be_oncall_during_my_booked_off_holiday/)

    This Reddit thread highlights the importance of management’s duty to adequately staff on-call teams and make on-call not suck. My favorite quote:

1. [40 Percent of IT Pros Expect to Work on Christmas Eve and Christmas Day](http://mobile.esecurityplanet.com/network-security/40-percent-of-it-pros-expect-to-work-christmas-eve-and-christmas-day.html)

    Whether or not you celebrate, I hope this holiday is a quiet and incident-free one for all of you. Being on call during the holidays can be an especially quick path to burnout and pager fatigue. As an industry, it’s important that we come up with ways to keep our services operational during the holidays with minimal interruption to folks’ family/vacation time.
1. [Secret Code Found in Juniper’s Firewalls Shows Risk of Government Backdoors](http://www.wired.com/2015/12/juniper-networks-hidden-backdoors-show-the-risk-of-government-backdoors/)

    Even if you think you’ve designed your infrastructure to be bulletproof, there may weaknesses lurking.
1. [Oh, Molly!](http://www.redpill-linpro.com/sysadvent//2015/12/19/molly-guard.html)

    Molly-guard is a tool to help you avoid rebooting the wrong host. Back at ${JOB[0]}, I mistakenly rebooted the host running the first production trial of a deploy tool that I’d spent 6 months writing, when it was 95% complete. Oops.
1. [7 Steps to a Minimum Viable Runbook](https://victorops.com/blog/7-steps-minimum-viable-runbook/)

    The final installment in a series on writing runbooks. The biggest takeaway for me is the importance of including a runbook link in every automated alert. Especially useful for those 3am incidents.
1. [Every Minute Counts: Coordinating Heroku’s Incident Response](http://www.heavybit.com/library/video/2014-06-17-blake-gentry)

    In this talk from last year (video & transcript), Blake Gentry talks about how Heroku’s incident response had evolved. Full disclosure: I work for Heroku. We still largely do things the same way, although now there’s a entire team dedicated to only the IC role.
### Outages

1. [Minecraft’s Wii U release caused Nintendo’s eShop to experience an intermittent outage](http://venturebeat.com/2015/12/17/minecrafts-wii-u-release-caused-nintendos-eshop-to-experience-an-intermittent-outage/)
1. [Apple’s iCloud Hit by Global Outage](http://examinerunion.com/2015/12/apple-8217-s-icloud-hit-by-global-outage/)
1. [Moonfruit outage](https://support.moonfruit.com/hc/en-us/articles/207109805)
    Moonfruit has been the target of a series of nasty DDoS attacks and blackmail threats. They stood up to the attackers and were very open about the situation, and I really respect them for that. DDoS attacks can be incredibly exhausting, frustrating, and difficult to fight.
1. [‘Phantom Squad’ threatens Xbox Live and PSN outages over Christmas](http://www.neowin.net/news/phantom-squad-threatens-xbox-live-and-psn-outages-over-christmas)
    More blackmail. ‘Tis the season.
1. [SendGrid Outage](http://status.sendgrid.com/incidents/55dy0tbb8d62)
1. [Xbox Live intermittend outage](http://www.express.co.uk/entertainment/gaming/627755/Xbox-Live-down-Xbox-One-sign-in-Xbox-360-Microsoft-confirm-why-is-Xbox-Live-down)
1. [Microsoft’s December 3 Office 365 outage: What went wrong](http://www.zdnet.com/article/microsofts-december-3-office-365-outage-what-went-wrong/)
1. [WWE Network Crashes During TLC Kickoff](http://www.pwmania.com/wwe-network-crashes-during-tlc-kickoff-wwe-tlc-attendance-breaking-ground-preview-for-monday)

### [ << Prev ](sreweekly-2.md) ------------- [ Next >> ](sreweekly-4.md)