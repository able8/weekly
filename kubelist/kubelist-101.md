## [Kubelist Issue #101 for 2020-09-23](https://kubelist.com/issue/101)

#### What we’ve learned in [n] years

> How many years have you been running Kubernetes in production? Have you shared your lessons learned yet? This week, we gathered some stories about lessons learned after running Kubernetes in a year, in two years or in three years. There are similarities and common threads in these stories. Definitely a good read, and also some takeaways that might change how you run Kubernetes this year.

1. [What we learned after a year of GitLab.com on Kubernetes](https://about.gitlab.com/blog/2020/09/16/year-of-kubernetes/)

    The folks at GitLab shared a great writeup of their first year running GitLab.com on Kubernetes. There are a lot of good takeaways in this post. It’s also neat to see how they handled the migration slowly, without changing their current, monolithic codebase. The first lesson learned could be expensive for GitLab and anyone else with a lot of bandwidth costs. 💸
1. [3 Years of Kubernetes in Production–Here’s What We Learned](https://medium.com/better-programming/3-years-of-kubernetes-in-production-heres-what-we-learned-44e77e1749c8)

    Another “lessons learned” post. Here are a few takeaways: “kubernetes transformation is not cheap” and “Be prepared to redesign your entire build and deployment pipelines”. The word “transformation” is used extensively in this write up, and that’s the right way to think about adopting Kubernetes. 💫
1. [How we learned to improve Kubernetes CronJobs at Scale](https://eng.lyft.com/improving-kubernetes-cronjobs-at-scale-part-1-cf1479df98d4)

    Lyft has so many thoughts to share about running CronJobs in production that they’ve written a two part post about it. Start with part 1 to understand the problem, and then head to part 2 to learn how Lyft addressed these issues. We’ve seen other “lessons learned” talk about cost, new traffic patterns, and complexity. This one is all about scale. How CronJobs scale (not great out of the box), so what can be done (a lot). 📏
1. [Lessons Learned From Two Years Of Kubernetes](https://coderanger.net/lessons-learned/)

    Another great “lessons learned” post from @kantrn. GitOps is the way, but secret management is still hard. Thanos is not overkill, and everyone should write more operators. ⚡
1. [Kubernetes at Cruise: Two Years of Multitenancy](https://www.youtube.com/watch?v=m19D9vZ1QFQ)

    Cruise’s KubeCon talk is a little different from the other links today. This is about multitenancy in their cluster (technically clusters, but one of them serves most of the traffic). This is a really interesting and practical talk. Cruise is using really approachable and common tech, and learning how they handled multitenancy is helpful when thinking about scaling Kubernetes in your organization. 🏙
1. [One year using Kubernetes in production: Lessons learned](https://techbeacon.com/devops/one-year-using-kubernetes-production-lessons-learned)

    This covers a lot of topics, from canary deployments and logging to ingress and load balancing. This is sort of a catch-all, but a great read from the first year of running Kubernetes. This author chose to not run data (stateful) services in the cluster, however it’s very common and acceptable to run these in-cluster these days. Kubernetes has come a long way, and this article is very interesting.
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1308145727361576962?s=20)

    There are some great ideas to enhance CNCF projects this year! ✨

### [ << Prev ](kubelist-100.md) ------------- [ Next >> ](kubelist-102.md)