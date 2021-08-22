## [Kubelist Issue #98 for 2020-09-02](https://kubelist.com/issue/98)

#### Virtual KubeCon EU 2020 & Kubernetes 1.19

> This week, we are sharing some new Kubernetes 1.19 features and a great technical post from Monzo about self-hosting Kubernetes. Let’s not forget that KubeCon EU 2020 wrapped up, so we made sure to include a couple of good recap links if you weren’t able to attend. We’ll share the videos when they are published! ✔

1. [Introducing Hierarchical Namespaces](https://kubernetes.io/blog/2020/08/14/introducing-hierarchical-namespaces/)

    The fine folks over at sig-multi-tenancy published this excellent deep dive into the new Hierarchical Namespaces concept in Kubernetes. It starts by acknowledging “no one tenancy model is likely to suit everyone”, and then dives into how inheritance and permission delegation might look in a cluster with subnamespaces. 🏊
1. [Customizing the Kubernetes Scheduler](https://kubernetes.io/docs/reference/scheduling/config/#profiles)

    If you ever felt limited by the Kubernetes built-in scheduler, 1.19 brings with it a number of interesting new features, including the long-awaited beta of KubeSchedulerConfiguration. This article covers the extension points and scheduling plugins that can be used to customize how Kubernetes decides which pods should run where.  🗓
1. [KubeCon Europe 2020 WrapUp](https://firehydrant.io/blog/kubecon-2020-europe-wrapup/)

    Now that KubeCon 2020 EU is behind us, this is a good recap of the first virtual KubeCon, with daily breakdowns. There were many good talks, and we can’t wait for all of them to be published!
1. [Summary Kubecon + CloudNativeCon EU 2020](https://elastisys.com/kubecon-cloudnativecon-eu-2020-summary/)

    If you are looking for another recap of the recent KubeCon, this is a well illustrated summary of CNCF and a video that covers the major take-aways from the event. 📺
1. [Envoy Looks to WebAssembly to Extend Microservices Monitoring](https://thenewstack.io/kubecon-eu-envoy-looks-to-webassembly-to-extend-microservices-monitoring/)

    Yaroslav Skopets from Tetrate gave a lightning talk on making envoy contributions feasible for everyone, in which WebAssembly is employed to lower entry cost for integration. ⚡
1. [Monzo: we learned a lot from self-hosting Kubernetes](https://www.computing.co.uk/news/4019233/monzo-learned-lot-self-hosting-kubernetes-wouldn%E2%80%99)

    The team at Monzo, presenters in this year’s KubeCon EU, have shared a bit of their history of working with Kubernetes since the 1.2 release, when managed cluster options were few and far between. For years Monzo’s team has been a source of fascinating post-mortems from in-place overlay network upgrades to etcd failures that took weeks to manifest. This post covers some of the lessons learned, some of the benefits of self-hosting, and why they wouldn’t do it again. 🙅‍
1. [Tweet of the Week](https://twitter.com/kubernetesio/status/1300767904250712065)

    In a move towards longer-term-support in Kubernetes, the support window for Kubernetes versions has increased from nine months to one year starting with 1.19. 🧰️

### [ << Prev ](kubelist-97.md) ------------- [ Next >> ](kubelist-99.md)