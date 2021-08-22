## [Kubelist Issue #1 for 2018-02-15](https://kubelist.com/issue/1)

#### Kubelist has launched!

> There’s a lot of news, tutorials, opinions and new projects around Kubernetes that comes out every day.

1. [The Twelve-Factors Kubernetes](https://blog.octo.com/en/the-twelve-factors-kubernetes/)

    Etienne Coutaud and Eric Favre provide a list of 12 rules for using Kubernetes, mostly related to application definition. With the 12 factors you need for your apps, that's either 24 or 144 factors total! Or maybe 8.9161004e+12?
1. [Level Triggering and Reconciliation in Kubernetes](https://hackernoon.com/level-triggering-and-reconciliation-in-kubernetes-1f17fe30333d)

    Our lead curator James Bowes uses concepts from electronics and systems programming to explain Edge and Level Triggering as it applies to the self-healing functionality of Kubernetes.
1. [CNCF publishes serverless landscape](https://www.cncf.io/blog/2018/02/14/cncf-takes-first-step-towards-serverless-computing/)

    The definition of Serverless has become diluted.  It's good to see the CNCF Working Group put out a whitepaper and landscape to help define and classify what Serverless is. The landscape includes a few Kubernetes native Serverless frameworks. Pick one, install it on your cluster, and use it to over-allocate that collection of chat bots and github webhooks your team has built up.
1. [Rainbow Deploys With Kubernetes](http://brandon.dimcheff.com/2018/02/rainbow-deploys-with-kubernetes/)

    Brandon Dimcheff describes how Olark uses multiple Deployment objects for a single application, to avoid disconnecting long-running clients. Gracefully handling long-running connections is often overlooked in orchestration and configuration management. Establishing a pattern for this in Kubernetes will be useful for many teams. Just don't forget to make your client tolerant of disconnections -- your deploys may be graceful, but the network isn't!
1. [Leader Election in Kubernetes Control Plane](https://blog.heptio.com/leader-election-in-kubernetes-control-plane-heptioprotip-1ed9fb0f3e6d)

    On the Heptio blog, Fabrizio Pandini gives a quick rundown of what leader election is, how its implemented in Kubernetes, and how to inspect the status of the election.I was hoping there'd be advice on how to prevent outside influence on your kube-scheduler leadership by observing and reacting to nascent patterns in social media. Maybe in a follow-up post?
1. [Tweet of the week](https://twitter.com/kelseyhightower/status/963413508300812295)

    Kelsey Hightower proclaims his opinion on running stateful workloads with Kubernetes.
Click through to read the whole thread.

### [ << Prev ](kubelist-0.md) ------------- [ Next >> ](kubelist-2.md)