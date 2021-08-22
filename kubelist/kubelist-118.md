## [Kubelist Issue #118 for 2021-03-10](https://kubelist.com/issue/118)

#### Kubernetes for local development 💻

> Sometimes you just want a local Kubernetes cluster on your laptop to use with development. Should you just enable the Kubernetes distribution that’s built into Docker? Or are there better options out there? In this week’s issue, we are sharing some links to posts and talks from people who have spent a lot of time solving this problem and are kind enough to share some advice. The tl;dr here is that there is no clear and obvious “best” cluster for local development, but if you are thinking about this problem, check out the links below for more data.

1. [Webinar: Navigating the Sea of Local Kubernetes Clusters](https://www.youtube.com/watch?v=q6kyHDleioA)

    Here is a presentation given by @arapulido, a Developer Advocate over @DatadogHQ, where she goes through most of the local cluster options and briefly touches on their architecture and design decisions. You can skip straight to her discussing and deploying MiniKube, kind, microK8s, k3s, k3d, or firekube. Or, if you’d prefer to read and not watch, you can check out the slides. 🧭⛵
1. [Kubernetes Setup Using Ansible and Vagrant](https://kubernetes.io/blog/2019/03/15/kubernetes-setup-using-ansible-and-vagrant/)

    If you are a developer that wants to automate your local laptop setup, Kubernetes installation doesn’t need to be a manual step performed at the end. There are plenty of reasons automation might be desirable; including reproducibility, setting up many laptops, or even the need/desire to reset your local environment frequently. From the official Kubernetes blog, there’s a guide on how to bootstrap your own local cluster with Ansible and Vagrant. BONUS: Here is a repo that has the playbooks already set up for you to get you started right away.🎛⚡
1. [Choosing a Local Dev Cluster](https://docs.tilt.dev/choosing_clusters.html)

    If you just read this post you are guaranteed to be making a relatively educated decision for the “choose your own” adventure that awaits you in getting kubernetes set up locally. The team over @tilt_dev gives you a great and insightful pros/cons list on many of the different options and even touches on remote dev clusters. 🎒
1. [How to Create a Kubernetes Cluster Locally](https://www.capitalone.com/tech/software-engineering/create-and-deploy-kubernetes-clusters/)

    There is a good chance you already have a one click solution installed for turning on kubernetes locally. The folks over at @capitalonetech offer a pretty straightforward tutorial on leveraging the kubernetes offering from docker for desktop. This is a good entry-level tutorial showing how to deploy a local app to the Kubernetes distribution that’s built into Docker Desktop. 🖥
1. [Minikube vs. kind vs. k3s - What should I use? ](https://brennerm.github.io/posts/minikube-vs-kind-vs-k3s.html)

    Here is another roundup that compares some of the top options for getting Kubernetes going locally. Max does an excellent job explaining the limitations and benefits of all three of these options. Each distribution covered has some unique strengths, and there’s no clear winner. This is a great article that provides data you can use as input to your decision about which distribution you want to run locally, based on your needs.
1. [Reddit Thread: K3s, minikube or microk8s?](https://www.reddit.com/r/kubernetes/comments/be0415/k3s_minikube_or_microk8s/)

    Here is a Reddit thread that has folks talking through the pros and cons of the various options specific to their workloads. The top answer (from u/petermbenjamin) is an incredibly detailed writeup of the differences between most of the local Kubernetes distribution options. We’ve shared some great blog posts this week, and this Reddit comment is one of our favorites because of how complete and thorough it is. 👾
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1367158436761395203)

    The KubeCon EU (Virtual) schedule is posted. There are some great talks on the schedule, and we are really excited about this event!

### [ << Prev ](kubelist-117.md) ------------- [ Next >> ](kubelist-119.md)