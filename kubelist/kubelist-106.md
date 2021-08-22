## [Kubelist Issue #106 for 2020-10-28](https://kubelist.com/issue/106)

#### Flux for GitOps ⚙

> Flux is the original GitOps operator for Kubernetes. Back in episode 103, we focused on the ArgoCD project as an implementation of GitOps in Kubernetes. This week, we are diving into Flux, and featuring a new episode of the <a href="https://kubelist.com/podcast">Kubelist Podcast</a> with Michael Bridgen, one of the creators of the project. This issue of the Kubelist newsletter contains links to Flux tutorials, guides, and a few pages that explain the Flux v1 to Flux v2 transition.

1. [GitOps is about to get better with Flux v2](https://www.weave.works/blog/gitops-with-flux-v2)

    Here’s a recent post on the Weave blog that discusses the changes between Flux v1 and Flux v2. The Kubelist Podcast this week is all about Flux, but if you’re hungry for information about the changes coming in Flux v2, give this a read. It’s a short, concise explanation of what to expect when you upgrade.
1. [Flux v2 Roadmap](https://toolkit.fluxcd.io/roadmap/)

    We decided to include the roadmap from the GitOps Toolkit since it paints a very easy-to-grok picture of the state of the Flux v2 project. If it’s incomplete in the GitOps Toolkit project, it’s certainly incomplete in Flux v2. At the time of this writing, the HelmOperator is 100% complete, but there’s a lot left to do on the image update side. If you are a Go developer and are waiting for Flux v2, maybe pick up an issue? 🗺
1. [GitOps Toolkit](https://toolkit.fluxcd.io/)

    The GitOps Toolkit is the core of Flux v2. This all used to be built into Flux v1 as a monolith, but is now separated into a composable and reusable library. If you are looking at adding anything related to GitOps functionality into an application you are developing, make sure to explore these docs before starting with a blank editor. 🧰
1. [New Kubernetes GitOps Toolkit — Flux CD v2](https://medium.com/@berndonline/new-kubernetes-gitops-toolkit-flux-cd-v2-229644233ea9)

    This post shows some of the new flexibility and power of Flux v2 when compared to v1. In this post, Bernd Malmqvist shares an example installation of Flux v2 where he’s pinned a semver range for a production cluster, ensuring that he can deploy newer versions to a pre-production cluster, without affecting the version on production. This opens the door to many interesting and powerful workflows. 📌
1. [How we do GitOps @ Mettle](https://itnext.io/how-we-do-gitops-mettle-4cc771a6c029)

    This post from the team at Mettle shows how they are using Flux v1 to install Kubernetes manifests, Helm charts, and Kustomize patches all at the same time. Infrastructure is never as neat as we’d like, and the challenges of implementing and deploying Flux come down to the complexities and variations in what you are looking to deploy.
1. [How to Continuously Deliver Kubernetes Applications With Flux CD](https://medium.com/better-programming/how-to-continuously-deliver-kubernetes-applications-with-flux-cd-502e4fb8ccfe)

    We like this post as a good intro to Flux. It covers the architecture of Flux v1, and proceeds to show how flux actually works, not just an example of a GitOps workflow. If you are comparing Flux vs other GitOps tools, or just want a better understanding of how Flux is working internally, this is a good place to start. 🔄
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1320790972574420993)

    Reminder to get signed up for KubeCon this week. It’s coming up quicker than you realize! 📢

### [ << Prev ](kubelist-105.md) ------------- [ Next >> ](kubelist-107.md)