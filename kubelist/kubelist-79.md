## [Kubelist Issue #79 for 2020-04-08](https://kubelist.com/issue/79)

#### Kubernetes 1.18 And More

> A lot has happened in the world since our last issue, including some Kubernetes developments worth talking about. With everyone quarantined right now, it felt like a good time to send some recommended reading (and a video) your way.

1. [Kubernetes 1.18.0 Is Here](https://sysdig.com/blog/whats-new-kubernetes-1-18/)

    Last week, Kubernetes 1.18.0 shipped, and as usual, the folks at Sysdig delivered a complete writeup explaining what’s new. I’m pretty excited about debug containers, but you can get your complete 1.18 fix here.
1. [kpt: Google’s new OSS packaging format](https://opensource.googleblog.com/2020/03/kpt-packaging-up-your-kubernetes.html)

    Google recently introduced kpt, a GitOps-friendly way to package up (and share) a collection of Kubernetes manifests. It’s still early, but add this to the shelf on your workbench that has Helm, Kustomize, CNAB and other packaging tools. It’s great to see more work being done here, and we are looking forward to a good deep dive into kpt soon.
1. [HashiCorp Joins the CNCF](https://www.hashicorp.com/blog/hashicorp-joins-the-cncf/)

    It’s pretty rare to run a modern workload without using some of the HashiCorp stack. And now it’s great that HashiCorp has made the commitment to join the CNCF. It will be interesting to see what changes this brings to the ecosystem.
1. [The Future of Kubernetes Attacks](https://www.youtube.com/watch?v=CH7S5rE3j8w)

    Ian Coldwater and Brad Geesaman gave an eye-opening talk at RSA talking about some of the future of attacks as we all adopt Kubernetes and run it in prod. Start right around the 10 minute mark here to see some great examples of how to get secrets and more from a cluster. Pay attention to the takeaways at the end of the presentation.
1. [A Model For Multicluster Workloads](https://timewitch.net/post/2020-03-31-multicluster-workloads/)

    A bit theoretical, but managing applications and workloads across disparate clusters is becoming a real-world problem that folks are starting to solve in interesting ways.
1. [What Happens When, Kubernetes Edition](https://github.com/jamiehannaford/what-happens-when-k8s/blob/master/README.md)

    This is a fun and highly informative explanation of what happens when you run a typical kubectl command. Kubernetes does a lot of work to execute (or try to execute) your command. This document does a great job of explaining the internal workings and helps provide a deeper understanding of the “magic”.
1. [Tweet of the week](https://twitter.com/pczarkowski/status/1244995897261469697)

    As opinions and tools abound for how to package an application in Kubernetes, this tweet hits a little close to home for some of us.

### [ << Prev ](kubelist-78.md) ------------- [ Next >> ](kubelist-80.md)