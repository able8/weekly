## [Kubelist Issue #131 for 2021-06-16](https://kubelist.com/issue/131)

#### Running Databases in Kubernetes 👟

> For years the rule of thumb in Kubernetes land was “Don’t run your own DB!”  Things have begun to change and this week we are going to look at a few of the options out there. Using a managed database service has always served us well but at Kubelist we always want to stay on top of the latest recommendations, and we think it’s totally acceptable to run databases (and other stateful services) in Kubernetes these days! 🏃🏽‍♀️

1. [To run or not to run a database on Kubernetes: What to consider](https://cloud.google.com/blog/products/databases/to-run-or-not-to-run-a-database-on-kubernetes-what-to-consider)

    The folks over at Google Cloud give us a nice writeup of the pros and cons of managed versus self hosted databases. This blog post provides us with tools to help decide what types of databases are good options for self hosting, and they even provide us with a basic decision tree if we should take the plunge.
1. [Using GitOps to Self-Manage Postgres in Kubernetes](https://blog.crunchydata.com/blog/gitops-postgres-kubernetes)

    Dive right into setting up your own highly available Postgresql cluster and get your hands dirty with some Statefulsets. In this tutorial the Crunchydata folks give us the raw K8s manifests we need to get up and running and evaluating. 🤿
1. [Production Checklist for Redis on Kubernetes](https://medium.com/swlh/production-checklist-for-redis-on-kubernetes-60173d5a5325)

    Lest us not forget Redis is one of the most valuable tools in most application developers toolboxes, and it too can be set up in our clusters. This post walks us through a whole list of different configurations for Redis and gives us a few protips and Helm charts to start with. 🧰️
1. [Running CockroachDB on Kubernetes](https://www.cockroachlabs.com/blog/running-cockroachdb-on-kubernetes/)

    CockroachDB is as OSS as it gets and designed from the ground up to live in the cloud. This blog post from Cockroach Labs does a thorough job of walking us through setting up CockroachDB on Kubernetes and even sprinkles in a bit of chaos to help test your install. If I were starting from scratch, there is a very good chance CockroachDB would be my first choice. And, of course, we recommend the SchemaHero operator to manage CockroachDB schemas!
1. [Modernizing Oracle operations with Kubernetes and El Carro ](https://opensource.googleblog.com/2021/05/modernizing-oracle-operations-with-kubernetes-el-carro.html)

    If we are being serious about the transition to Kubernetes we need to understand our options for ALL databases, not just OSS… El Carro adopts the Operator pattern and gives us an early path to setting up Oracle in a cloud native environment. This project is early but quite interesting if you need (or want) to run Oracle in Kubernetes. ⛅️
1. [Introducing Tobs: Deploy a full observability suite for Kubernetes in two minutes ](https://blog.timescale.com/blog/introducing-tobs-deploy-a-full-observability-suite-for-kubernetes-in-two-minutes/)

    It’s likely that one of the most important databases you are going to run in your cluster is the one that monitors it and all your services. The folks over at TimeScale introduce us to Tobs, an easy way to install the tools to collect, store, and visualize metrics. ⏱
1. [Tweet of the Week](https://twitter.com/jbeda/status/1401580679415009285)

    Happy 7 years old, @kubernetesio! This is a bit of a special birthday as 7 is a special number for k8s. Huge thank you to everyone that joined in the journey and made it happen!

### [ << Prev ](kubelist-130.md) ------------- [ Next >> ](kubelist-132.md)