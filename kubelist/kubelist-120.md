## [Kubelist Issue #120 for 2021-03-24](https://kubelist.com/issue/120)

#### Open Policy Agent Graduates 🎓

> It wasn’t that long ago (KubeCon 2018) that <a href="https://www.openpolicyagent.org/">Open Policy Agent (OPA)</a> was a Sandbox project, and now it&#39;s a <a href="https://www.cncf.io/announcements/2021/02/04/cloud-native-computing-foundation-announces-open-policy-agent-graduation/">CNCF graduate</a>! OPA is a declarative policy system that has its own DSL; so you can control what/who can access not only your cluster but all of your cloud native resources. In this week&#39;s issue we dive into why OPA exists, what problems OPA tackles, and of course some real world experiences and gotchas using the ecosystem that has formed around it.

1. [Better Kubernetes Security with Open Policy Agent (OPA)](https://www.stackrox.com/post/2020/04/enhancing-kubernetes-security-with-open-policy-agent-opa-part-1/)

    Fuzzykb provides us a great introduction to OPA. They cover a bit of everything that can get us off to a great start when diving in. On the theory side, it dives into the motivations behind the project; while on the pragmatic side it gives a tangible use case and examples of tooling in the ecosystem. 🛡
1. [Integrating Open Policy Agent (OPA) With Kubernetes](https://www.magalix.com/blog/integrating-open-policy-agent-opa-with-kubernetes-a-deep-dive-tutorial)

    This guide does an excellent job of laying out the different pieces you need when implementing OPA in your cluster. It touches on admission controllers, and gives examples of enforcing specific labels by namespace and limiting what registries your images can originate from. It also walks you through setting up your own constraints, and puts you in a position to start testing on your own clusters immediately. BONUS: How to use Kube-mgmt with OPA 🧩
1. [Five OPA and Styra Trends that Prove Kubernetes Adoption](https://blog.styra.com/blog/opa-styra-trends-prove-kubernetes-adoption)

    The people that brought us OPA in the first place give us an update on adoption within the k8s ecosystem. Downloads are officially a hockey stick, and large adopters (Intuit, Goldman Sachs, TripAdvisor) keep popping up daily. Other vendors are announcing integrations, and things seem to be moving rapidly towards adoption! 📈
1. [5 tips for using the Rego language for Open Policy Agent (OPA)](https://www.fugue.co/blog/5-tips-for-using-the-rego-language-for-open-policy-agent-opa)

    OPA actually has its own DSL (Rego) to enable you to define policy as code. Rego has a straightforward syntax and a small set of functions and operators, and is optimized for queries. The team over at Fugue has a great writeup of good practices for navigating Rego and brings you back to your CS days (recursion, tree navigation, references etc.) 📄
1. [OPA and Gatekeeper: OPA or Gatekeeper?](https://www.infracloud.io/blogs/opa-and-gatekeeper/)

    This article runs us through the gambit for OPA Gatekeeper. It briefly walks us through the design decisions for Gatekeeper and then discusses ConstraintTemplates and CI/CD integrations. Ggctwts does a good job of explaining the different use cases for OPA Gatekeeper and how to augment your existing workflows. 🚪
1. [Open Policy Agent 2020, year in review!](https://blog.openpolicyagent.org/open-policy-agent-2020-year-in-review-dc25b60308d7)

    2020 was a banner year for OPA and this article dives into the numbers and growth in the ecosystem. TL;DR: 29m downloads, 1500 gh stars, 1700 slack users, 750 commits, expansion of management services, a whole bunch of tooling, and much more! Add CNCF graduation to the list for 2021! 🎉
1. [Tweet of the week](https://twitter.com/estesp/status/1374073402454118422)

    Running containerd? Good news, nerdctl is now part of the containerd project!

### [ << Prev ](kubelist-119.md) ------------- [ Next >> ](kubelist-121.md)