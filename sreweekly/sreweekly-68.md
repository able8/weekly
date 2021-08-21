## [SRE Weekly Issue #68](https://sreweekly.com/sre-weekly-issue-68/) - April 17, 2017
### Articles

1. [Increment: On-Call](https://increment.com/on-call/)

    The big story this week is the release of the inaugural issue of Increment, a newsletter by Stripe, edited by Susan Fowler. They bill it as “A digital magazine about how teams build and operate software systems at scale” and the first issue, dedicated to on-call, certainly delivers. Below, I’ll include my short take on each article in the issue.
1. [What happens when the pager goes off?](https://increment.com/on-call/when-the-pager-goes-off/)

    Increment interviewed over thirty companies to build a picture of the common practices in incident response. I’m actually pretty surprised to hear that “it turns out that they all follow similar (if not completely identical) incident response processes”, but apparently the commonalities don’t stop at just process:Bonus content: Julia Evans shared her notes on Twitter.
1. [Who owns on-call?](https://increment.com/on-call/who-owns-on-call/)

    Next up, Increment addresses the dichotomy of ops teams versus developers on call for their code. It turns out that the latter practice is more prevalent than I’d realized.
1. [Crafting sustainable on-call rotations](https://increment.com/on-call/crafting-sustainable-on-call-rotations/)

    After laying a solid groundwork of suggestions for avoiding burn-out in on-call, this next Increment article raises a really important point: on-call affects people differently based on privilege. Example: single parents are going to have a much harder time of it.
1. [The benefits of transparency](https://increment.com/on-call/the-benefits-of-transparency/)

    Remember a couple of months back when GitLab live-streamed their incident response? Increment caught up with their CEO to give us this in-depth interview about their radical transparency.
1. [On-call at any size](https://increment.com/on-call/on-call-at-any-size/)

    Increment shares tips and key practices for setting up on-call, targeted to companies of size ranges varying from 0-10 employees all the way up to 10000+.
1. [How should startups approach on-call and incident response?](https://increment.com/on-call/ask-an-expert/)

    Increment rounds out their issue with advice in the form of quotes from six of the companies they interviewed.
1. [Introducing Honeycomb: A New Debugging Service to Address Flawed Systems](https://thenewstack.io/honeycomb-addresses-flawed-systems/)

    The other big news of the week is the official launch of Honeycomb.io. If you haven’t had a chance to check it out yet, here’s an introduction, and you can also sign up for a free one-month trial.
### Outages

1. [Melbourne IT](http://www.afr.com/technology/melbourne-it-suffer-denialofservice-attack-thousands-of-websites-inaccessible-20170413-gvkiok)
    A DDoS took out their DNS service, taking out customer domains and also sites they they host for customers. While this is a news article and not a formal post-analysis, it does include some pretty interesting technical detail from an interview with their CTO. I’m not sure that he did himself any favors by quoting the definition of their SLA:
“People look at 99.9 per cent and think that’s seconds of downtime, but you work it out and it’s 45 minutes.”
1. [Google Cloud HTTP(S) Load Balancer](http://status.cloud.google.com/incident/compute/17007#5659118702428160)
    Google Cloud LB threw 502s for 25% of requests in a 22-minute period. They released this post-analysis 7 days later, and I have to say, the root cause is pretty interesting – and sadly familiar.
A bug in the HTTP(S) Load Balancer configuration update process caused it to revert to a configuration that was substantially out of date.

### [ << Prev ](sreweekly-67.md) ------------- [ Next >> ](sreweekly-69.md)