## [DevOps'ish 174](https://devopsish.com/174) - Sun Apr 5, 2020

I understand that there are some of you looking for something to learn while we’re in this odd time. I wrote this week’s newsletter with that in mind. <a href="https://youtu.be/K5qI9bTG9Ik">If you don’t know <code>git</code>, <strong>now</strong> is the time to learn</a> for sure. Please don’t hesitate to reach out if you have any feedback. Thank you for subscribing and reading.

### People

1. [We’re all in this Together: A Wellness Guide from the CNCF Well-Being Working Group](https://www.cncf.io/blog/2020/04/03/were-all-in-this-together-a-wellness-guide-from-the-cncf-well-being-working-group/)

    If you’re struggling, I’m right there with you. Don’t be afraid to reach out for help.
1. [IBM awards its second $50,000 Open Source Community Grant to internship and mentorship program Outreachy](https://devops.com/ibm-awards-its-second-50000-open-source-community-grant-to-internship-and-mentorship-program-outreachy/)

    “Our open source community nominated a number of nonprofits doing incredible work and, while voting was close with plenty of deserving organizations in the mix, we awarded Outreachy the most votes for their commitment to providing paid internships to underserved and underrepresented minorities.”
1. [Jeff Geerling’s DevOps books are free in April, thanks to Device42](https://www.jeffgeerling.com/blog/2020/my-devops-books-are-free-april-thanks-device42)

    Huge thanks for these books. Ansible for DevOps changed my career. Ansible for Kubernetes builds on that. Thank you!
1. [The best games to play while waiting out the coronavirus](https://www.washingtonpost.com/video-games/2020/04/03/best-video-games-modern-nintendo-playstation-xbox-pc/?utm_source=reddit.com)

    I will freely admit, I’m looking for distractions from the news and numbers.
1. [WeWork founder Adam Neumann loses out as SoftBank scraps share buyout](https://www.cnn.com/2020/04/02/tech/wework-softbank-bailout/index.html)

    points and laughs 🤣🤣🤣
### Process

1. [Migrating to Kubernetes](https://blog.sensu.io/migrating-to-kubernetes)

    “How to migrate depends, to an extent, on what you want to migrate from. The primary consideration is whether your existing infrastructure runs workloads in containers. If so, you’re already off to a quick start because you won’t have the containerization step to complete. Otherwise, you have a clear place to start.”
1. [A hacker has wiped, defaced more than 15,000 Elasticsearch servers](https://www.zdnet.com/article/a-hacker-has-wiped-defaced-more-than-15000-elasticsearch-servers/)

    “Hacker tries to pin the blame on Night Lion Security, a US cyber-security firm.”
1. [Ops by Pull Request: An Ansible GitOps Story](https://www.ansible.com/blog/ops-by-pull-request-an-ansible-gitops-story)

    “There are many tools you can use in your GitOps pipelines; however, Ansible provides some unique advantages that make it ideal for these workflows and extending their use beyond Kubernetes and cloud-native systems.”
1. [Who’s freezing hiring from coronavirus](https://candor.co/hiring-freezes/)

    A continuously updated AirTable of companies and whether or not they’re hiring.
1. [3 Tips to Avoid WFH Burnout](https://hbr.org/2020/04/3-tips-to-avoid-wfh-burnout)

    “Maintain physical and social boundaries. Maintain temporal boundaries as much as possible. Focus on your most important work.”
### Tools

1. [Kpt: Packaging up your Kubernetes configuration with git and YAML since 2014](https://opensource.googleblog.com/2020/03/kpt-packaging-up-your-kubernetes.html)

    “Today we’d like to introduce kpt, an OSS tool for Kubernetes packaging, which uses a standard format to bundle, publish, customize, update, and apply configuration manifests.”
1. [Forget AWS Lambda, Kubernetes AND Fargate - what we need is beyond all three](https://diginomica.com/forget-aws-lambda-kubernetes-fargate-need-beyond)

    The case for containers without thinking about capacity (which is exactly what we need right now).
1. [WireGuard 1.0.0 Christened As A Modern Secure VPN Alternative To OpenVPN/IPsec](https://www.phoronix.com/scan.php?page=news_item&px=WireGuard-1.0.0-Released)

    Welcome to the kernel, WireGuard.
1. [Online Debug for AWS Lambda on your IDE!](https://devopsi.sh/aws-lambda-97890)

    Thundra allows you to natively debug your serverless applications on the cloud with their own permissions. Thundra’s online-debugger sets up a secure bridge between your AWS Lambda environment and your IDE. VSCode and IntelliJ IDEA are natively supported with plugins. For other IDEs, we provide a portable client to foster the integration with any IDEs. Start debugging Node.js, Python and Java functions for free today! SPONSORED
1. [Deploying a Windows Server 2019 virtual machine using OpenShift Virtualization (YouTube)](https://youtu.be/Kx110kqoHo0)

    “OpenShift virtualization enables Red Hat OpenShift clusters deployed to bare metal physical servers to use the KVM hypervisor to deploy and manage virtual machines utilizing Kubernetes APIs and Operators. For applications which consist of a mix of containers and VMs, OpenShift virtualization provides a way to bring both components under the same management paradigm, allowing the application team to move at the speed of containers, but still retain the characteristics of a virtual machine when needed.”
1. [Improvements to the Ingress API in Kubernetes 1.18](https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/)

    “In Kubernetes 1.18, we’ve made 3 significant additions to this API”
1. [Introducing Kubie](https://blog.sbstp.ca/introducing-kubie/)

    “I kept being bothered by the fact that kubectx and kubens modify a global config file. That is, when you change your shell’s current context or namespace, it also affects all other shells…”
1. [Detailed Guide to Kubernetes Configurations - Download Today](https://devopsi.sh/guide-kubernetes-c346e)

    Accidental misconfiguration is one of the leading causes of security breaches. When it comes to Kubernetes, the configuration options are complicated, security features are not enabled by default, and most of the community is learning how to securely configure their infrastructure components. In this guide, StackRox takes a deep dive into key Kubernetes configurations and recommended best practices you should follow to secure your API server, kubelet, pod security policies, etc. - with code snippets to check the configs. SPONSORED
1. [Provisioning cloud resources (AWS, GCP, Azure) in Kubernetes](https://learnk8s.io/cloud-resources-kubernetes)

    “TL;DR: You can create and connect to managed cloud resources from Kubernetes using the Service Catalog, a tool such as Kubeform or cloud-specific operators such as Config Connector and AWS Operator Service.”
1. [Kubernetes Topology Manager Moves to Beta - Align Up!](https://kubernetes.io/blog/2020/04/01/kubernetes-1-18-feature-topoloy-manager-beta/)

    The TopologyManager feature enables NUMA alignment of CPUs and peripheral devices (such as SR-IOV VFs and GPUs), allowing your workload to run in an environment optimized for low-latency.
1. [kubevious](https://github.com/kubevious/kubevious)

    “Kubevious project, application centric Kubernetes UI”
1. [ankane/pgsync](https://github.com/ankane/pgsync)

    “Sync data from one Postgres database to another”

### [ << Prev ](devopsweekly-173.md) ------------- [ Next >> ](devopsweekly-175.md)