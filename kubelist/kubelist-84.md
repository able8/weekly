## [Kubelist Issue #84 for 2020-05-13](https://kubelist.com/issue/84)

#### Kubernetes Security

> This isn’t the first time we’ve written an entire issue about security in  Kubernetes. And it’s probably not the last time. The security landscape is constantly evolving and there’s a steady stream of best practices to implement and new attacks to protect against. Whether you are just starting to look at securing Kubernetes and the applications running on it, or want to catch up on some new info, the links here about applying security principles to Kubernetes should be interesting. Don’t stop at these 6 links though. Hopefully this inspires you to set a little time aside very soon to continue the research. 🛡️

1. [SELinux, Kubernetes & Udica](https://jaosorior.dev/2019/selinux-and-kubernetes/)

    Does your Kubernetes cluster have SELinux running in enforce mode on the nodes? Enforcing SELinux policies without interfering with application workloads is not necessarily easy to do, but it can provide substantial protection when executed well. In this post @jaosorior offers some tips and advice to extend your Kubernetes manifests to include SELinux policies. The selinux-operator looks super interesting.
1. [Threat Matrix for Kubernetes](https://www.microsoft.com/security/blog/2020/04/02/attack-matrix-kubernetes/)

    A group of security researchers at Microsoft, armed with operational expertise in Kubernetes. took some time to document techniques to exploit the tactics documented in the excellent MITRE ATT&CK framework. If you know Kubernetes and want to think like an attacker, or are a (white hat) attacker trying to apply your craft to a Kubernetes-world, this is an in-the-weeds place very practical guide to help you get started.
1. [How to Harden Your Kubernetes Cluster for Production](https://medium.com/better-programming/how-to-harden-your-kubernetes-cluster-for-production-7e47990efc2a)

    This guide includes 15 high-level suggestions that every cluster should implement before running production workloads. Some of these may be obvious but some might surprise you. Highlights include advice like “protect your etcd cluster like a treasure vault” (see the next link for what happens when you don’t 😬).  
1. [Using K3s for command and control on compromised Linux hosts](https://blog.christophetd.fr/using-k3s-for-command-and-control-on-compromised-linux-hosts/)

    We’ve shared the RSA talk that inspired this blog post back in Issue 79. If you think you follow all of the best practices to secure your applications and access to your cluster, this is still worth a quick read. If you haven’t considered how a “command and control node” could be used to steal all of your secrets, this post by @christophetd is a very practical example that does a great job explaining this attack. ⚔️
1. [Detect large-scale cryptocurrency mining attack against Kubernetes clusters](https://azure.microsoft.com/en-us/blog/detect-largescale-cryptocurrency-mining-attack-against-kubernetes-clusters/)

    If you need a reminder to build out your Kubernetes RBAC policies or to disable public internet access to the Kubernetes Dashboard, give this a read. There’s more here than just those tips, but this wasn’t that sophisticated of an attack. Kubernetes is hard folks, we have to try harder than this to make it secure.
1. [Five things CISOs can do to make containers secure and compliant](https://sysdig.com/blog/five-things-cisos-containers-secure-compliant/)

    These are pretty high level requirements to running any modern infrastructure securely from the Sysdig team. Kubernetes does provide a lot of the primitives so we can run applications securely (and of course, Sysdig ships some useful tools like Falco to help here). For the highest level post of the week, look here. Do you scan at build, deployment and past deployment time? Do you know how to check out the Kubernetes Audit Log to figure out how an object was created or what’s changed on it? 👩
1. [Tweet of the week](https://twitter.com/sometorin/status/1259955151735074820)

    Is there anything OPA can’t do?

### [ << Prev ](kubelist-83.md) ------------- [ Next >> ](kubelist-85.md)