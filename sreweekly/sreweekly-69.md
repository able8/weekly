## [SRE Weekly Issue #69](https://sreweekly.com/sre-weekly-issue-69/) - April 23, 2017
### Articles

1. [The tale of the flying gurney: MRI machines can turn objects into projectiles](https://www.bostonglobe.com/metro/2017/04/08/the-tale-flying-gurney/QwkE01J33BYYVF4vdMX4aP/story.html)

    In February of 2016, a metal hospital gurney was inadvertently wheeled* into an MRI room, resulting in a costly near-miss accident. Brigham and Women’s Hospital posted about the mishap on their Safety Matters blog and also released a Q&A with their chief quality officer about their dedication to an open and just culture.* Yes, I used the passive voice on purpose. See what I did there?
1. [Lies My Parents Told Me (About Logs)](https://honeycomb.io/blog/2017/04/lies-my-parents-told-me-about-logs/)

    Sometimes logs help us prevent outages or discover a cause. But raise your hand if you’ve seen logging cause an outage. Yeah, me too.
1. [Syscall Auditing at Scale](https://slack.engineering/syscall-auditing-at-scale-e6a3ca8ac1b8)

    Traditionally, auditd, together with Linux’s system call auditing support, has been used as part of security monitoring. Slack developed go-audit so that they could use system call auditing as a general monitoring tool. I can think of plenty of outages during which I’d have loved to be able to query system call patterns!
1. [Introducing Stormcrow](https://blogs.dropbox.com/tech/2017/03/introducing-stormcrow/)

    Dropbox has some pretty complex needs around feature gating. Apparently existing tools couldn’t satisfy their use case so they wrote and released a tool with sophisticated user segmentation support.
1. [Niffy: Perceptual Diffing to Catch Invisible Bugs](https://segment.com/blog/perceptual-diffing-with-niffy/)

    Why depend on fallible QA testing to spot regressions in a web UI? Computers are so much better at that kind of thing. Niffy spots the pixel changes between old and new code so you can see exactly what regressed before putting it in front of your users.
1. [Online migrations at scale](https://stripe.com/blog/online-migrations)

    In this beautifully-illustrated article, Stripe engineer Jacqueline Xu explains how Stripe safely rolled out a major database schema upgrade. Many code paths had to be refactored, and they took a methodical, incremental approach to avoid downtime. Thanks to this article, I now know about Scientist and can’t wait to use it.
1. [Scaling your API with rate limiters](https://stripe.com/blog/rate-limiters)

    Speaking of Stripe, they have another polished post on how and why to add load shedding to your API.
1. [GitHub – github/scientist: A Ruby library for carefully refactoring critical paths.](https://github.com/github/scientist)

    Scientist is such an awesome idea. The idea is to try out a new code path and see if it produces the same result as the old code path. It only returns the new code path, so you know you can safely prove to yourself whether the new code path is safe before exposing users to it.
1. [AWS status: The complete guide to monitoring status on the web’s largest cloud provider](http://blog.statuspage.io/aws-status-the-complete-guide-to-monitoring-status-on-the-web-s-largest-cloud-provider)

    I’m including this article at least in part due to its mention of the February S3 outage. AWS had difficulty reporting the outage on its status site because of a dependency on S3.
1. [CASE STUDY: Etsy, Sprouter and Conway’s Law](http://itrevolution.com/etsy-sprouter-and-conways-law/)

    Conway’s Law is extremely important to us as SREs. As we can see in the example of Sprouter, a poor organizational structure can produce unreliable software. My fellow SRE, Courtney Eckhardt, loves to say, “My job is applying Conway’s Law in reverse.”
### Outages

1. [AT&T VoIP](http://www.chron.com/business/bizfeed/article/AT-T-business-voice-customers-experience-outage-11046835.php)
    I received an anonymous anecdote from an SRE Weekly reader (thanks!) that this affected at least one hospital, with the result that critical phone communication was significantly hampered. What happened to the good old mostly-reliable traditional phone system? Irony: in the reader’s case, an announcement about the failure was sent out via email.
1. [Three](http://www.gizmodo.co.uk/2017/04/three-network-is-down-sms-texts-are-reportedly-going-to-the-wrong-people/)
    This is the second case this year of a telecom outage resulting in SMSes being delivered to the wrong people. Am I the only one that finds this an extremely surprising and concerning failure mode?
1. [eBay](http://www.express.co.uk/life-style/science-technology/795307/eBay-down-website-not-working-UK-buy-sell-log-in)
1. [Red Hat](https://linux.slashdot.org/story/17/04/21/1623233/red-hat-suffers-massive-data-center-network-outage)

### [ << Prev ](sreweekly-68.md) ------------- [ Next >> ](sreweekly-70.md)