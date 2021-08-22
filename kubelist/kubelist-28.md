## [Kubelist Issue #28 for 2018-08-23](https://kubelist.com/issue/28)

#### the almost entirely new releases issue

> Welcome to a very special almost entirely new release announcements issue of kubelist! This week, we&#39;ve got a handful of updates and new projects. For those of you that don&#39;t enjoy NEW PRODUCT BULLETINS we&#39;ve also included a guide for installing kubernetes on bare metal and CentOS, and a way to use Helm <em>without actually using Helm</em>. Enjoy!

1. [Install Kubernetes on bare-metal CentOS7](https://itnext.io/install-kubernetes-on-bare-metal-centos7-fba40e9bb3de)

    The kubelist editors would be hard-pressed to tell you what distribution of Linux is running in their cluster, let alone what version it is. The cloud handles all that for us, right? If getting your hands dirty from the silicon up sounds good, give this guide a read.
1. [Deploy only what you trust: introducing Binary Authorization for Google Kubernetes Engine](https://cloud.google.com/blog/products/identity-security/deploy-only-what-you-trust-introducing-binary-authorization-for-google-kubernetes-engine)

    On the Google Cloud blog, Jianing Sandra Guo introduces Binary Authorization for GKE. It's always great to see new features on top of Kubernetes that use Open Source, and Binary Authorization builds on Grafeas. Grafeas looks powerful. Beyond simply signing your container image, it provides signed attestation that some action has been done with the container image. With this, you can ensure that no container image is deployed to production unless it's provably run through every phase of your deployment pipeline. 
1. [Introducing Heighliner](https://blog.manifold.co/introducing-heighliner-1fb233c577ad)

    Disclaimer: the kubelist editors are employees of Manifold.As with last week's post on our credentials integration (⚠️yes, we're shilling again), we wanted to share another tool we've built and use at Manifold. Heighliner is a collection of Custom Resource Definitions and Controllers that reduce the YAML you need to write, codify some best practices for container security, and create preview deployments for pull request reviews. We use Heighliner for most of our deployments.Heighliner is still early, but we do hope you give it a look, and can find some value in it 🙂.
1. [Cilium 1.2: DNS Security Policies, EKS Support, ClusterMesh, kube-router integration, ...](https://cilium.io/blog/2018/08/21/cilium-12/)

    Jessie Frazelle's favorite software defined networking stack released version 1.2 this week. There's a load of new features in this release; most of them we don't understand! DNS based security policies sound pretty good though.
1. [Make a Kubernetes Operator in 15 minutes with Helm](https://blog.openshift.com/make-a-kubernetes-operator-in-15-minutes-with-helm/)

    On the OpenShift blog, Rob Szumski introduces the Helm Operator kit, part of the larger Operator framework. The Helm Operator kit seems like it could provide a reasonable incremental migration away from the Helm CLI and Tiller, should you wish to do so. Nobody will judge you if you do stop using the Helm CLI; it's your cluster!
1. [Tweet of the Week](https://twitter.com/michellenoorali/status/1032257089618542593)

    For as much as we focus on the tech (are you tired of Custom Resource Definitions yet?), Kubernetes' true strength lies in its welcoming and inclusive community. We hope it only gets better.

### [ << Prev ](kubelist-27.md) ------------- [ Next >> ](kubelist-29.md)