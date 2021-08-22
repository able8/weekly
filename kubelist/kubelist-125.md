## [Kubelist Issue #125 for 2021-04-28](https://kubelist.com/issue/125)

#### Kyverno

> <a href="https://twitter.com/jimbugwadia?lang=en">Jim Bugwadia</a> and <a href="https://twitter.com/riteshdp?lang=en">Ritesh Patel</a> from <a href="https://nirmata.com">Nirmata</a> joined me on episode 14 of <a href="https://kubelist.com/podcast">The Kubelist Podcast</a> to talk about <a href="https://kyverno.io">Kyverno</a>. This is another new CNCF Sandbox project, and Jim and Ritesh did a great job explaining the unique value proposition of Kyverno on the episode. If policy enforcement and simplifying some operations by using YAML is something you are looking at or working on, it might be good to give Kyverno a look.

1. [Kubernetes Policy Comparison: OPA/Gatekeeper vs Kyverno](https://neonmirrors.net/post/2021-02/kubernetes-policy-comparison-opa-gatekeeper-vs-kyverno/)

    Let’s dive right in with a comparison of Kyverno to another well-known policy management project: Open Policy Agent. While OPA can do more than K8s, this comparison is K8s-focused and can help explain some of the differences and strengths of each of the two projects. We can’t really start an entire issue on policy management without starting here! ⚖️
1. [Auto-labeling Kubernetes resources with Kyverno](https://www.cncf.io/blog/2020/12/30/auto-labeling-kubernetes-resources-with-kyverno/)

    Here’s a walkthrough of default functionality that Kyverno provides, complete with a simple policy. This guide helped me really understand the power of the declarative policy format and what’s possible. As a bonus, the guide isn’t a contrived scenario; it’s a useful policy to add so that if an unexpected namespace shows up in your cluster, you’ll know who created it! 🔖
1. [Easy as one-two-three policy management with Kyverno on Amazon EKS](https://aws.amazon.com/blogs/containers/easy-as-one-two-three-policy-management-with-kyverno-on-amazon-eks/)

    The title would lead you to believe that this is specific to EKS, but it’s not. This is another useful guide on how to get started with a couple of policies. The policies that are used in this post are quite useful: prevent privileged containers from running and create an allow-list of approved registries that images can be pulled from. Supply chain is critical, and this is a great policy to figure out how to add! 📋
1. [Exploring Kyverno](https://neonmirrors.net/post/2020-12/exploring-kyverno-part1/)

    A detailed, 3-part series that goes into quite a bit of depth on Kyverno. Part 1 is all about validation policies, or the kind that reject manifests that don’t meet a specific set of rules. Part 2 is about mutation, or altering resources at deploy-time. The classic example here is adding labels, but you can do a lot more. Part 3 is about generation, which is unique to Kyverno and definitely the real power of the platform. Instead of stopping at mutating a resource, Kyverno can trigger generation of entirely new resources. This is borderline-operator functionality, without writing code! 📘
1. [Kyverno: Kubernetes Native Policy Management Ritesh Patel & Jim Bugwadia (Nirmata) OpenShiftCommons](https://www.youtube.com/watch?v=27U0L-X3tN0)

    Ritesh and Jim gave a talk describing Kyverno on OpenShiftCommons last month. Starting around 10 minutes into this talk, Ritesh starts to explain the differences between Kyverno and Open Policy Agent. The two projects definitely have some overlap, but they approach the problem of policy enforcement differently, and they each offer some unique features. 
1. [The Way of the Future | Kubernetes Policy Management with Kyverno](https://www.youtube.com/watch?v=8fgrjBnxqi0)

    On this episode of The Way Of The Future, Abhay explains Kyverno, but then goes into a screenshare and demo. The demo and walkthrough start around 5 minutes in, and if you don’t have a cluster but want to actually see Kyverno in action, this is a good option. 🔮
1. [Tweet of the week](https://twitter.com/memenetes/status/1385337969062580231)

    Making complex things looking easy always require a little (or a lot) of planning!

### [ << Prev ](kubelist-124.md) ------------- [ Next >> ](kubelist-126.md)