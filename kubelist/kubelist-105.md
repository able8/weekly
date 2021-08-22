## [Kubelist Issue #105 for 2020-10-21](https://kubelist.com/issue/105)

#### SPIFFE and SPIRE: Don’t Trust The Network 🤞

> This week on the <a href="https://kubelist.com/podcast/">Kubelist Podcast</a>, we welcome Sunil James from Hewlett Packard Enterprise to chat about the SPIFFE and SPIRE projects. Sunil is one of the founders of Scytale, the company that brought SPIFFE into the CNCF, and has continued to work on it after being acquired by HPE. If you’re running microservices and are thinking about security and identity, check out the links in this issue to learn more about SPIFFE and SPIRE, and tune in to this week’s podcast episode!

1. [Attesting Istio workload identities with SPIFFE and SPIRE](https://developer.ibm.com/components/istio/articles/istio-identity-spiffe-spire/)

    Istio already uses Kubernetes RBAC for access control, so when we first saw this post from IBM, we wondered where it was going. They included a section titled “Why the current Istio mechanism is not enough” – which helped by illustrating a scenario where the SPIFFE model can work on top of Istio to solve a problem. For anyone with Istio, this is a recommended read.
1. [AWS IAM with SPIFFE & SPIRE](https://johnharris.io/2020/03/aws-iam-with-spiffe-spire/)

    An interesting approach to using SPIFFE with AWS IAM instance roles by John Harris of VMWare. This post shows how SPIFFE can deliver a secure, Kubernetes-native alternative to providing access to IAM roles for a Pod, instead of the cluster or the node. If you have questions about the capabilities or use of SPIFFE, this is an interesting way to use it, and shows the flexibility of the projects.
1. [Who's Calling? Production Identity in a Microservices World](http://slides.eightypercent.net/spiffe-intro/index.html#p1)

    Here’s a link to the Gluecon presentation where Joe Beda originally introduced SPIFFE. Keep in mind that this was presented before any SPIFFE code was written. While at Google, Joe identified this problem and wrote a proposal for the SPIFFE specification. This presentation is still one of the best walk-throughs to answer the question “why are strong microservice identities important?” (something that we have all asked ourselves on sleepless nights). ☎
1. [Using SPIRE to (Automatically) Deliver TLS Certificates to Envoy For Stronger Authentication](https://blog.envoyproxy.io/using-spire-to-automatically-deliver-tls-certificates-to-envoy-for-stronger-authentication-be5606ac9c75)

    If you’d had enough to read about the SPIFFE specification, you can dive into an implementation (SPIRE). This post on the Envoy blog is an informative guide on integrating SPIRE (so, SPIFFE) with Envoy to make your application more secure.
1. [Five Things You Didn’t Know You Could Do with SPIFFE and SPIRE](https://www.youtube.com/watch?v=5m6kjzdysBI)

    From KubeCon NA (San Diego 2019), this is a walk through of some use cases of SPIFFE and SPIRE. This presentation isn’t theoretical or experimental edge cases, but following some of the advice in this highly relevant talk will lead you to building a more secure production system.
1. [The SPIFFE Whitepaper](https://docs.google.com/document/d/1GjurNK2ROw4rXz-k-l68JtpGRkGj2fZcWqP6gksEriQ/edit#heading=h.pq1kki84bhak)

    The original whitepaper published to start the SPIFFE project. It’s an interesting read, even if somewhat outdated at this point because Scytale, HP and the entire community have continued to push and evolve this project. If you want to dig deeper into SPIFFE, this is the origin story. 📜
1. [Tweet of the week](https://twitter.com/justincormack/status/1318505386362753024)

    Justin does a good job explaining these recent CVEs and why we will continue to see issues like this unless we make some changes to how we write code.

### [ << Prev ](kubelist-104.md) ------------- [ Next >> ](kubelist-106.md)