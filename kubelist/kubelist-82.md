## [Kubelist Issue #82 for 2020-04-29](https://kubelist.com/issue/82)

#### Installing Kubernetes ⛵️

> As time goes on, we aren’t running out of new ways to bootstrap and install a Kubernetes cluster. We aren’t complaining though, the fact that there are so many options here is proof that this problem is not yet solved! 

1. [WKSctl - A New OSS Kubernetes Manager using GitOps](https://www.weave.works/blog/wksctl-a-new-oss-kubernetes-manager-using-gitops)

    GitOps everything, right? Why not GitOps the cluster too then. The team at Weave really has a good idea here. The first bullet point in this announcement is a desire to treat clusters like cattle. There’s a good section on “the problem with upstream Kubernetes” which is worth a read to explain why you can’t have your cake and eat it too when it comes to flexibility of installers. 🍰
1. [eksctl - The official CLI for Amazon EKS](https://eksctl.io/)

    When Amazon launched EKS (their managed Kubernetes offering), it wasn’t very easy to figure out how to get started. The team at Weave decided to build a CLI to simplify the onboarding here, and Amazon has since made this the “official” CLI for creating EKS clusters. It’s good — 1 command and a lot of CloudFormation happens, leaving you with a running cluster and not thinking about what just transpired to make all of it happen. 🚀
1. [kURL - A Custom Kubernetes Distro Creator](https://kurl.sh/)

    kURL is a distribution creator, for those of us who can’t find the perfect distribution. This is probably a growing trend — Forbes recently published about the challenges of vendor-supplied K8s distributions. kURL has a lot of add-ons available, and allows you to pick & choose what to include. After you design your installer, you can download and run it on an airgapped network. This is a pretty easy way to create your own custom Kubernetes installer, completely based on upstream Kubeadm and other popular CNCF tools. 🏗
1. [KinD (Kubernetes in Docker) (GitHub) Action](https://github.com/marketplace/actions/kind-kubernetes-in-docker-action)

    Not every cluster is designed to live a long life. This GitHub Action creates a KinD cluster right in your CI workflow, allowing you to run end to end tests directly on GitHub’s infrastructure. These small “clusters” are quick to create, can be added with just a few lines of YAML, and are a perfect, cheap and quick target to run tests against.
1. [HA Kubernetes Clusters on AWS with Cluster API v1alpha3](https://blog.scottlowe.org/2020/03/26/ha-kubernetes-clusters-on-aws-with-cluster-api-v1alpha3/)

    If you haven’t tried out the Cluster API way to deploy a new Kubernetes cluster yet, and you want to really see more than a toy-demo, this is an up-to-date and detailed walkthrough. It’s one thing to get a cluster set up locally to experiment with, but @scott_lowe shows how to use the latest release of Cluster API to provision a production-grade, highly available Kubernetes cluster on AWS. ✨
1. [Cluster API v1alpha3](https://kubernetes.io/blog/2020/04/21/cluster-api-v1alpha3-delivers-new-features-and-an-improved-user-experience/)

    Congrats to the sig-cluster-lifecycle team on the quick update to v1alpha3. There are some great new Cluster API additions — and functionality like MachineHealthCheck is starting to make it possible to manage a cluster completely, with less vendor-specific knowledge to learn and maintain.
1. [Tweet of the week](https://twitter.com/danielepolencic/status/1254330583380979712)

    When @danielepolencic finally built a tool to let his CEO control the cloud spend.

### [ << Prev ](kubelist-81.md) ------------- [ Next >> ](kubelist-83.md)