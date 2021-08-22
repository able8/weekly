## [Kubelist Issue #94 for 2020-08-05](https://kubelist.com/issue/94)

#### Serverless on Kubernetes ⚡

> Not everyone using Kubernetes writes PodSpecs and thinks about Ingress controllers, PersistentVolumeClaims and other Kubernetes fundamentals. While that type of Kubernetes-native workload is fine for some users, many use cases don’t need to know the details of how a pod is scheduled. Kubernetes is a great platform to run “serverless” workloads on. The idea of delivering a function or other interface to run on a cluster is appealing. This week, we look at a couple of popular options, and some tutorials to get started with each of them.

1. [Serverless Workflow Specification](https://serverlessworkflow.github.io/)

    The Serverless Workflow Specification is one of the newest CNCF Sandbox projects. It’s a little different from most other sandbox projects in that it’s a specification and packages for your application to use. They’ve already released early versions of Go and Java SDKs. If you are planning to run production-grade serverless projects on Kubernetes in the future, take some time to read this project spec as it will likely become the interoperability standard of the various platforms in the near future.
1. [Knative: Serving Your Serverless Services](https://www.openshift.com/blog/knative-serving-your-serverless-services)

    Knative has been around for a little while. It’s a Google sponsored platform to power serverless (and more). While serverless is just one part of Knative, this is an ambitious project that builds on top of Kubernetes, Istio, Fluentd and more. The tutorial linked here is a quick “101-style” walkthrough of building a Knative application and is a good place to get started. 📘
1. [Introduction to OpenFaaS: Serverless Made Accessible](https://www.objectif-libre.com/en/blog/2020/01/17/openfaas-made-accessible/)

    OpenFaaS is one of the original “serverless” platforms for Kubernetes, created by @alexellisuk as an open source project. This quick tutorial should be enough to show you what it’s like to run a FaaS provider on Kubernetes. If you have a local Kubernetes cluster available, you can deploy this Hello World python function to it, and quickly have a better understanding of what OpenFaaS can do. 🐳
1. [Host a CloudEvents compatible Event Gateway on Kubernetes](https://www.serverless.com/blog/host-your-own-cncf-cloudevents-compatible-event-gateway-kubernetes)

    There’s no question that functions as a service and other serverless platforms are an easy way to deploy REST APIs, but Event Gateway is an interesting approach to consuming additional event types. If you have Kubernetes, CloudEvents, and Event Gateway, you could build a quick IFTTT or Zapier for some internal events like glueing together systems. Think of the possibilities of building internal tools, reporting and business-team functionality based on events that occur in the database, queues or other external systems. This model lets non-core engineers extend the platform without modifying the core system. 🌤
1. [CloudEvents](https://cloudevents.io/)

    CloudEvents is not a serverless or functions-as-a-service platform. This is a specification and a component that helps make your serverless application more portable. Using CloudEvents will remove and avoid the vendor lock-in that happens with various cloud providers. ☁
1. [Nodeless Kubernetes](https://medium.com/elotl-blog/nodeless-kubernetes-on-eks-9fb14003b651)

    We’ve mentioned Nodeless back in Issue 80, but it’s worth bringing up again. Nodeless is decidedly not serverless. A nodeless cluster creates a new instance per pod, making nodeless really mean that there’s no dedicated Kubernetes nodes for co-tenancy of pods. But if you think about pairing a nodeless cluster with a FaaS provider, you have a very dynamic and inexpensive platform. Has anyone tried OpenFaaS on Nodeless yet? 💁‍
1. [Tweet of the week](https://twitter.com/bgrant0607/status/1290811274683588609)

    The extensibility of the Kubernetes API allows for a lot of use cases. 

### [ << Prev ](kubelist-93.md) ------------- [ Next >> ](kubelist-95.md)