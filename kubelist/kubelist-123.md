## [Kubelist Issue #123 for 2021-04-14](https://kubelist.com/issue/123)

#### Kubernetes 1.21 Drops

> This week, we’re going to dive into some new <a href="https://kubernetes.io/blog/2021/04/08/kubernetes-1-21-release-announcement/">Kubernetes 1.21</a> features, and see what’s the latest and greatest in the ecosystem. Bonus: Docker 1.20 is supported by kubeadm now! Let’s get going!

1. [Kubernetes Pod Security Policy Deprecation: All You Need to Know](https://blog.aquasec.com/kubernetess-policy)

    | Deprecated | Amit over at Aqua Security does a great job of explaining what Pod Security Policy is and why it’s being deprecated. He dives deep into Pod Security Standards (PSS) and some of the things to keep an eye out for when migrating over. 👀
1. [ Kubernetes 1.21 released!](https://jasworks.org/kubernetes-1-21-released/)

    We now turn to Michael Lam with his writeup on upgrading 1.20 -> 1.21. He dives into the process he went through upgrading the various nodes and tooling needed to upgrade. Spoiler: it all went pretty smoothly, but make sure to check out the notes as the article can save you some headaches and gotchas! 🪄
1. [Graceful node shutdown #2000](https://github.com/kubernetes/enhancements/issues/2000)

    | Graduating to Beta | In a little change of pace, we wanted to highlight PR #2000 on K8s. The actual thread starts in Sep 2020, and if we follow it along, we can see how to contribute directly to the codebase and the governance associated with it. It’s also pretty awesome that Kubelets can now gracefully terminate pods during a node shutdown, and will prove to be quite valuable for those using Spot Instances while running their k8s! 🛑
1. [IPv4/IPv6 dual-stack](https://kubernetes.io/docs/concepts/services-networking/dual-stack/)

    | Graduating to Beta | Now enabled by default, dual IPv4/IPv6 network stacks are fully supported and could end up being quite impactful. First-class native support for IPv6 routing to pods and services will help alleviate some of the limitations of scaling multi-service workloads inside K8s, while still allowing clusters to leverage IPv4 where needed. Be sure to check your filtering settings once you upgrade, now that all services / pods have a IPv6! 🥞
1. [Deep dive into Kubernetes CronJob](https://michael.bouvy.net/post/deep-dive-kubernetes-cronjob)

    | Graduating to Stable | Making sure a computer does something on a schedule is DevOps/SRE 101. Traditional CronJobs along with NTP have been the de facto convention for running periodic tasks for decades. With the release of 1.21 CronJobs has graduated to stable. Check out this article to see how and why they work, and even set one up yourself! 📆
1. [Sysdig Wrap Up](https://sysdig.com/blog/kubernetes-1-21-whats-new/)

    Another Kubernetes version shipped; so we get another great deep-dive from the folks over at Sysdig. There’s plenty we haven’t covered in Kubelist this week (because we can only make the newsletter so long) but head over to Sysdig for the most comprehensive writeup, and get your full k8s fix! 🍫
1. [Tweet of the week](https://twitter.com/stephenaugustus/status/1381695422113144838)

    Big Congrats to Stephen. Looking forward to what #OpenCisco is gonna be working on...

### [ << Prev ](kubelist-122.md) ------------- [ Next >> ](kubelist-124.md)