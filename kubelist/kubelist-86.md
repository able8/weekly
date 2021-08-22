## [Kubelist Issue #86 for 2020-05-27](https://kubelist.com/issue/86)

#### Liveness and Readiness ⏰

> This week, we’ve been working to improve our application startup process so that it doesn’t receive traffic before it’s ready, and fails gracefully so that it can be restarted. These are all core features of a modern application, and Kubernetes supports this via Readiness Probes and Liveness Probes. But just because these are supported, doesn’t mean you necessarily should be using both of them. In this issue, we’ll cover the basics of the probes, when to use them, and more importantly, when not to.

1. [Liveness and Readiness Probes](https://www.openshift.com/blog/liveness-and-readiness-probes)

    Let’s start with an overview of what Liveness and Readiness probes are, and what they are designed to solve. The visuals of how various timing of failed probes can impact the service give some help when building your own probes and setting the intervals and timeouts. This post starts by explaining Liveness probes, but as you’ll see in the next couple of links, these might not be the first probes to add if you are just getting started.
1. [Kubernetes Liveness and Readiness Probes: How to Avoid Shooting Yourself in the Foot](https://blog.colinbreck.com/kubernetes-liveness-and-readiness-probes-how-to-avoid-shooting-yourself-in-the-foot/)

    In this post, Colin shows us how both Liveness and Readiness probes can cause harm and create outages. The recommendations at the bottom of each section are great and, if nothing else, listen to these because they aren’t just theoretical outage-causing patterns. Also, there’s advice to regularly restart and exercise the probes once you’ve set them to ensure that your probe timing is still compatible with your container startup times, network latency and other dependencies. 🦶
1. [Liveness Probes Are Dangerous](https://srcco.de/posts/kubernetes-liveness-probes-are-dangerous.html)

    This is an interesting post that shows the dangers of Liveness probes specifically. The recommendation is: “I recommend to avoid them unless you have a clear use case and understand the consequences.” The list of do’s and don’ts before implementing Liveness probes is a great place to start, if you decide you need Liveness probes. However, the overall message in this post is really to question if you actually need Liveness probes. 💣
1. [How should I answer a health check?](https://medium.com/polarsquad/how-should-i-answer-a-health-check-aa1fcf6e858e)

    This post is not written specifically about Kubernetes Readiness and Liveness probes, but the recommendations on how to design the health check endpoints that the probes will use is great. Kuebrnetes probes support various types of checks (http requests, exec a command, tcp socket). This post is a great place to help figure out the right health checks for your application pods, and then how you can map the right checks back to Readiness probes and maybe even Liveness probes. 🌡
1. [How to use readiness, liveness, and startup probes](https://www.innoq.com/en/blog/kubernetes-probes/)

    Another post about creating Liveness and Readiness probes, with some good info about exception handling. This is important to think about because things aren’t going as planned when a probe is failing, so hopefully you’ve designed for it. Liveness probes are designed to give Kubernetes the ability to terminate your pods, and if you haven’t thought about exception handling well enough, you might be creating bigger and harder to debug problems when implementing naive probe checks.
1. [Kubernetes 1.19 Release Delayed](https://github.com/kubernetes/sig-release/pull/1065)

    If you missed this one this past week, the Kubernetes 1.19 release date is slipping by a couple of weeks. It’s now scheduled to ship the week before Kubecon. This isn’t a trend or a new release cadence, but it’s good to see the flexibility and support from the #sig-release team when needed. Shipping Kubernetes every 3 months is hard work, so big thanks to everyone involved! 📆
1. [Tweet of the week](https://twitter.com/mitchellh/status/1265414263281029120)

    We need to do a deeper dive into this exciting new announcement from HashiCorp soon.

### [ << Prev ](kubelist-85.md) ------------- [ Next >> ](kubelist-87.md)