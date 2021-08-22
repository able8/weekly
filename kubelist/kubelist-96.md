## [Kubelist Issue #96 for 2020-08-19](https://kubelist.com/issue/96)

#### ZeroTrust 🔒

> This week we’ll be focusing on the Zero Trust security model and how it applies to Kubernetes. This issue is split between blogs that explain what zero trust is, and hands-on practical tutorials and examples. Have your kubectl ready to go!

1. [Zero Trust for Beginners](https://www.stackrox.com/wiki/zero-trust-networks-in-kubernetes-cloud-native-applications/)

    If you’ve never heard of the zero trust model, or have heard of it but don’t know what fits under this umbrella, this blog post will help you ease into the subject. ☂ Here you will also find links to additional materials and practical tutorials that are suitable for someone who is getting started. 
1. [Adopt a Zero Trust Network Model for Security](https://docs.projectcalico.org/security/adopt-zero-trust)

    Learn how to model a zero trust network model to secure your Kubernetes workloads using NetworkPolicy and Istio. This is a great primer into the mindset of Zero Trust as well as its definition as applied to Kubernetes.
1. [Keynote: Progress Toward Zero Trust Kubernetes Networks](https://www.youtube.com/watch?v=Agxt9Vg-YP4&t=12s)

    Spike Curtis, a leading contributor to Istio, discusses how organizations are adapting Zero Trust and how the Kubernetes security posture can be increased with Service Mesh and other tools.
1. [EKS: Best Practices for Network Security](https://aws.github.io/aws-eks-best-practices/network/)

    This article from the EKS best practices guide goes into more detail about securing a network that is more specific to EKS. Focusing on network policies, it brings together security groups, encryption, network traffic logging, and more using tools that are native to or recommended for AWS. ☁
1. [Securing a GKE Cluster with Istio](https://cloud.google.com/blog/products/networking/the-service-mesh-era-securing-your-environment-with-istio)

    This blog post from Google reviews all aspects of securing a cluster with Istio, with links to practical examples, like this tutorial. ⛵️
1. [Network Policy Visualizer](https://orca.tufin.io/netpol/)

    Nobody likes to read YAML. If you have a long and complex list of NetworkPolicy definitions, this tool from Tufin can help visualize it in a way that may look familiar to those who have spent time configuring AWS security groups. 👓
1. [Tweet of the week](https://twitter.com/alexellisuk/status/1295783066489835523)

    Due to the recent changes in DockerHub rate limiting, Alex Ellis has created an Operator for propagating your imagePullSecret to all of your namespaces. ⚡

### [ << Prev ](kubelist-95.md) ------------- [ Next >> ](kubelist-97.md)