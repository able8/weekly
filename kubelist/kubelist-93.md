## [Kubelist Issue #93 for 2020-07-29](https://kubelist.com/issue/93)

#### K8s 1.19 is pending, and some ecosystem news ⏳

> Kubernetes 1.19 is expected to be released in the next couple of weeks. Today, we start out by sharing some of the 1.19 features and improvements that are exciting to us. We’ve also got a few links regarding some timely CNCF news about KubeCon and certifications.

1. [Immutable Secrets and ConfigMaps Advances to Beta](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1412-immutable-secrets-and-configmaps/README.md#summary)

    Kubernetes 1.19 will see “Immutable Ephemeral Volumes” (aka Immutable Secrets and ConfigMaps) advance from alpha to beta. This means that most clusters will be able to access the feature without modifying flags in the api server (once upgraded to 1.19 of course). If you’re unfamiliar with the feature, the writeup discusses how immutability addresses a number of the challenges that arise when pods consume configuration as a volume. This feature represents one more step towards allowing application maintainers to manage configuration in a safe, predictable way. ⚙
1. [Improvements to kubectl debug](https://www.europeclouds.com/blog/debugging-with-ephemeral-containers-in-k8s)

    Ephemeral debug containers came to alpha in Kubernetes 1.18, and they are being extended in 1.19. Pull Request #92310 adds a node debug command that runs debug commands in the node’s PID, network and IPC namespaces; with access to the node’s filesystem. This should make it easier to use kubectl to debug node problems and run nodes without ssh access at all. The command is still alpha though, so turn on the feature gate to get access. 🔍
1. [A Bootiful Podcast: Kubernetes Release SIG Tim Pepper](https://spring.io/blog/2020/06/05/a-bootiful-podcast-kubernetes-release-sig-tim-pepper)

    As we are grateful for all of the work that the Kubernetes release team does to get us shiny new Kubernetes versions every 3 (or so) months, this podcast is a philosophical take on the process. It offers some insights about how decisions are made in the Kubernetes release cycle, and how the current state of the world may have affected it. Even if you aren’t a Spring developer, this is a good listen. 🔊
1. [New Sandbox Project: Serverless Workflow](https://serverlessworkflow.github.io/)

    Serverless Workflow is a new project added to the CNCF Sandbox. This is a specification for describing vendor-neutral markup for defining serverless workloads. This project is still very early, but a Go and Java SDK have been added to the org.
1. [Certified Kubernetes Security Specialist Announced](https://www.cncf.io/blog/2020/07/15/certified-kubernetes-security-specialist-cks-coming-in-november/)

    For anyone who has a current Certified Kubernetes Administrator (CKA) certification, a new program has been added if you want to demonstrate your knowledge in Kubernetes Security. The certification programs offered by the CNCF seem to be popular, and this is a great addition. Looking forward to reading the course material! 🔒
1. [KubeCon Europe 2020 (Virtual)](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/)

    The final schedule for the upcoming KubeCon EU has been published. The event will start at 1PM CET, in order to create more overlap with other time zones. Which sessions are you planning to attend? 
1. [Tweet of the week](https://twitter.com/castrojo/status/1288244264498081793)

    Some fun, completely made up KubeCon session submissions. 😂

### [ << Prev ](kubelist-92.md) ------------- [ Next >> ](kubelist-94.md)