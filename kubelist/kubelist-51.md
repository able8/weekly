## [Kubelist Issue #51 for 2019-02-07](https://kubelist.com/issue/51)

#### K8s tools, secret management, maybe even write your own plugins?

> K8s is a great tool, but tools that make using K8s easier are even better. In this week&#39;s issue, we&#39;ll introduce a few tools that can improve your workflows for managing K8s resources and managing secrets. If you are feeling inspired by the end, the last article goes through how to write your own kubectl commands - so maybe you can write one too!

1. [If You K8s, Please Try K9s...](https://medium.com/@fernand.galiana/if-you-k8s-please-try-k9s-82ea30eb9aa3)

    Do you manage a ton of Kubernetes resources? Bonus question: do you like dogs? K9 developed by Fernand Galiana allows you to monitor Kubernetes resources in your clusters from single Terminal panel. You can set it to fetch the resources on a given interval. Your days of managing terminal windows just to see the state of your clusters is over.In this article, the author gives a great overview of the tool. For more information, visit https://k9ss.io/ and give it a try!
1. [Can Kubernetes Keep a Secret?](https://blog.solutotlv.com/can-kubernetes-keep-a-secret/)

    Kubernetes has a built-in secret management system, but there are flaws. It is difficult to audit changes, RBAC only allows users to either get or set secrets, secrets are encoded and not safe to be version controlled.The author in this article touches on these points and introduces a tool called Kamus. It allows you to encrypt your secret which can only be decrypted for consumption by the application, so that you can manage your secrets in version control.
1. [Bank Vaults](https://banzaicloud.com/bank-vaults/)

    Bank-vaults is another option for managing secrets within your Kubernetes clusters, using a well establish secrets platform: Hashicorp Vault. Bank-vaults allows you to run Vault inside a Kubernetes cluster, and comes with features like: automated token rotation, secret injection into pods, operator for interacting with Vault, and more. The article goes through the CLI tool, the Go library for interacting with Vault, using and deploying bank-vaults in great details.
1. [Write your own kubectl subcommands](https://ahmet.im/blog/kubectl-plugins/)

    Kubernetes is great, but tools that make using Kubernetes easier are even better.  Feature not known to many is Kubernetes' plugin ecosystem.  kubectl as of 1.12 allows you to invoke an executable file in your $PATH prefixed by kubectl-<subcommand> as kubectl <subcommand>.   In this article, the author explores how to write your own plugins, some of the challenges, and where the community is headed. If you're feeling inspired by some of the tools that make your life as a K8s developer easier,  this one is for you.
1. [Tweet of the Week](https://twitter.com/RealGophersShip/status/1092888278158708738)

    Hey! Speaking of useful K8s tools, Kustomize 2.0 is released!

### [ << Prev ](kubelist-50.md) ------------- [ Next >> ](kubelist-52.md)