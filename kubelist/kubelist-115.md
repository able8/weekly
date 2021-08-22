## [Kubelist Issue #115 for 2021-02-03](https://kubelist.com/issue/115)

#### Better & More Secure Secrets 🔑

> Kubernetes has a built-in Kind called Secrets to store and pass sensitive configuration data to Pods. It’s popular – almost everyone running a cluster uses Secrets. Since it’s a Secret, it must be protected in a reasonable way by implementing best security practices, right? It can be, but it’s not always. It depends on the configuration of your cluster; and you shouldn’t assume that the defaults are the best settings from a security perspective.

1. [A Better Story For Kubernetes Secrets](https://www.youtube.com/watch?v=7jSfJombUeY)

    Here is a presentation by @sethvargo at StrangeLoop back in 2019 (at the last IRL StrangeLoop, before the pandemic). Anyone running a Kubernetes cluster not familiar with how etcd is storing your secrets should watch this before doing anything else. The TL;DR here is that secrets contain sensitive information, and are probably sitting unencrypted in etcd. This is bad, and especially bad if your etcd cluster is exposed to the internet. Good advice here – and a good primer on how to secure these secrets. 📖
1. [Securing Kubernetes secrets: How to efficiently secure access to etcd and protect your secrets](https://medium.com/opsguru/securing-kubernetes-secrets-how-to-efficiently-secure-access-to-etcd-and-protect-your-secrets-b147791da768)

    Another article demonstrating how to use EncryptionConfig to ensure that Secrets are encrypted at rest in etcd; but this article goes further to show some limitations and how to overcome them. Even if you read this and decide that EncryptionConfig is sufficient to secure your sensitive secrets in etcd, defense-in-depth is a best practice and there are some great tips about adding authentication to your etcd cluster here. 🔒
1. [Protecting Kubernetes Secrets: A Practical Guide](https://blog.aquasec.com/managing-kubernetes-secrets)

    Here’s a great guide to a high-level understanding of Kubernetes secrets by Rani Osnat at Aqua Security. A lot of the same information is here (etcd is not encrypted by default), but Rani starts to talk about the pros and cons of using third party delivery mechanisms and utilities to deliver sensitive information. Using a proven application like HashiCorp Vault can remove the security challenges of storing secrets in etcd, but take a look to understand the tradeoffs.
1. [Managing secrets deployment in Kubernetes using Sealed Secrets](https://aws.amazon.com/blogs/opensource/managing-secrets-deployment-in-kubernetes-using-sealed-secrets/)

    We are big fans of Bitnami’s SealedSecret controller to distribute secrets securely using GitOps workflows. But remember that you are writing encrypted data, and the controller is decrypting it to create a vanilla Secret at runtime. So SealedSecrets is a great way to distribute the secret values through the CI/CD pipeline, but you still need to secure etcd and the rest of the Kubernetes API server or else you will be leaking the unencrypted values in the end anyway. 🗄
1. [Using EKS encryption provider support for defense-in-depth](https://aws.amazon.com/blogs/containers/using-eks-encryption-provider-support-for-defense-in-depth/)

    A closer look at using the AWS key management EncryptionConfig provider, which will be a common, easy choice for EKS clusters. Enabling this takes a little bit of “day one” configuration and time, but key rotation and all ongoing work should be handled automatically, and without downtime of the cluster. If you are running an EKS cluster, start here to understand how to secure access to etcd and secrets. 🛡
1. [Using Azure Key Vault for Kubernetes Data Encryption](https://ritazh.com/using-azure-key-vault-for-kubernetes-data-encryption-d5eac8daee71)

    This is a little older, but still a great guide. You no longer need the --experimental flag passed to kubelet when starting to enable the EncryptionConfig; and it’s natively supported on Azure (AKS) clusters. There’s really no excuse to not have this data encrypted, but it’s important to enable it because etcd may not be encrypted by default. 🗝
1. [Tweet of the week](https://twitter.com/cloudnativefdn/status/1356656376395653123)

    Congrats to the new Technical Oversight Committee members!

### [ << Prev ](kubelist-114.md) ------------- [ Next >> ](kubelist-116.md)