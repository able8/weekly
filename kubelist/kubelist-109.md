## [Kubelist Issue #109 for 2020-11-18](https://kubelist.com/issue/109)

#### New Sandbox Projects 🏖

> Every couple of months, the CNCF TOC (Technical Oversight Committee) meets to review outstanding Sandbox applications. As a result of the meeting last week, 7 new projects were added to the Sandbox. In this issue, we are taking a quick look at each of the new additions, and we hope to feature them all on the Kubelist podcast soon!

1. [cert-manager](https://cert-manager.io/)

    Cert Manager is a widely used Kubernetes project to manage the lifecycle of Let’s Encrypt, Vault, Venafi, and other certs. This is a mature project that probably won’t last too long in the Sandbox before advancing to Incubation. But we are happy to see it donated, to show a more open governance and increased confidence in adopting cert-manager! 📜
1. [OpenKruise](https://openkruise.io/en-us/)

    OpenKruise is a set of CRDs that extend some of the built-in kinds in Kubernetes. For example, there’s the “Advanced StatefulSet” to introduce a new maxUnavailable attribute and adding a new in-place update strategy. It’s still to be seen if these extensions should be brought into Kubernetes and not simply extended via OpenKruise. But this is a great collection of extensions and you can install and use them now! 🛳
1. [Cloud Development Kit for Kubernetes](https://cdk8s.io/)

    CDK8s is a project from AWS Labs. This is a Kubernetes-native approach to building Kubernetes applications. (This makes sense, right?) With CDK8s, a developer can write an app in a familiar language (Typescript, Python, Java…) and get a deployable Kubernetes manifest. Sort of like a SDK for writing a K8s micro-service, without thinking about the YAML and manifests. 📦
1. [SchemaHero](https://schemahero.io)

    SchemaHero is a declarative approach to managing traditional database schemas. With this project, you can define the desired schema of the database, and deploy it. The SchemaHero operator then compares the desired schema to the actual running schema, generates all of the necessary CREATE and ALTER statements, and allows you to review and approve them. This is a great way to include traditional database schema management in a GitOps / Kubernetes workflow. 🗃
1. [Tinkerbell](https://tinkerbell.org)

    A project from Equinix Metal (they own Packet) to both provision and manage bare metal. This is a really exciting project because it allows layers below Kubernetes to be managed using the same concepts as we’ve started to use for Kubernetes. There’s a lot here, from workflows to metadata services. Tinkerbell is a project to keep an eye on if bare metal is something you are watching. ⚒
1. [Pravega](https://www.pravega.io/)

    Streams as storage. This is cool and should be watched. One use case to help explain the project is collecting logs and data from massively distributed systems, and then being able to query them in both realtime and historically. Traditionally, this would require writing a lot of code that’s different for the realtime vs historic use cases, and probably adding a new database into your stack. Pravega adds an API that you can use to query any of the data that you need, at any time. 🕊
1. [Kyverno](https://kyverno.io)

    Kyverno is a project that provides policy management in a Kubernetes native way. While you can do this with Open Policy Agent and Gatekeeper today, Kyverno uses YAML as a DSL instead of rego. It will be interesting to watch if this is powerful enough, but the sample policies provided in the repo are pretty impressive. Using YAML-based policies allows workflows that use Kustomize and other ecosystem tooling to modify the policies per environment. We love OPA, and Kyverno looks to be delivering a much more narrowly scoped solution to Kubernetes admission controller policies.
1. [Tweet of the week](https://twitter.com/datadoghq/status/1328702118019149825)

    Datadog is in an interesting position to see how a large number of end customers use Kubernetes and containers. We need to do a deep dive into this container report. Which is more interesting to you: Kubernetes 1.15 is the most popular version of K8s, or that 1 out of 3 AWS container environments run on Fargate?

### [ << Prev ](kubelist-108.md) ------------- [ Next >> ](kubelist-110.md)