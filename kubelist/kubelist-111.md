## [Kubelist Issue #111 for 2020-12-16](https://kubelist.com/issue/111)

#### Telepresence 🕶

> <a href="https://twitter.com/rdli">Richard Li</a> from <a href="https://www.getambassador.io">Ambassador Labs</a> was on the <a href="https://kubelist.com/podcast">Kubelist Podcast</a> this week to discuss <a href="https://telepresence.io">Telepresence</a>, one of the first CNCF Sandbox projects. Telepresence makes sense for large teams to manage the cost of running many dev clusters better; but in this episode Richard explains how the project makes sense for small teams as well. The conversation about Telepresence 2.0 (rewritten in Go) is also great. Hopefully we get to see that version soon!

1. [Telepresence](https://www.telepresence.io)

    The Telepresence project home. Start here to understand the project (and, of course, listen to this week’s podcast). Telepresence is a clever way to run part of a development environment locally, while everything else is in a remote cluster. This is useful when the entire environment is simply too large or it’s just not practical to run locally. The collaboration aspects of a shared dev environment are probably too often overlooked, and worth checking out.
1. [Fast development workflow with Docker and Kubernetes](https://www.telepresence.io/tutorials/docker)

    Sort of the “quick start” for Telepresence. If you are curious to try it out, this is a good straightforward place to start. This guide will walk you through setting up your local environment and installing the Telepresence bits. Hopefully this will help the project be a little less abstract and more real, if you are trying to see how Telepresence works. 🏁
1. [Debug a Kubernetes service locally](https://www.telepresence.io/tutorials/kubernetes)

    Here’s where the real power of Telepresence starts to show for folks building on Kubernetes. If you have Docker for Mac or another local K8s cluster because you want to have everything local, give this guide a look. The Telepresence team shows that you don’t lose the power of the local development environment, but can run it in the cloud if you have a decent enough Internet connection. 👾
1. [Use Your Favorite Developer Tools in Kubernetes With Telepresence](https://youtu.be/JyZfWMmeVlU)

    Here’s a talk from KubeCon San Diego last year where Abhay Saxena (from the Telepresence team) talked about how you don’t need to adopt a new IDE or new workflow in order to move your development environment to the cloud. This is the video that really helped me understand how the project worked.
1. [Seamless Development Environments on Kubernetes using Telepresence](https://youtu.be/8Dl8U-AbJN0)

    From an earlier KubeCon, Ara Pulido from Bitnami talked about Telepresence. This is a fun, live demo where Ara starts using Telepresence to configure her development environment, but starts to dig into the cluster and show what’s actually happening. In this demo, she uses kubectl to examine the cluster while using Telepresence, and this is a great way to see what’s happening “under the hood”.
1. [Eliminating Local Resource Constraints for Building Cloud Native Applications](https://www.getambassador.io/resources/eliminate-local-resource-constraints/)

    This is the story, right? Most of us end up running some version of K8s locally because that’s just how we’ve known to test code. We build and run it all locally, and with Kubernetes, we pulled all of that down too. Then, as this post calls out, our local K8s installation has made the fan run continuously on our laptop, and we question our decisions. So we tune in to the next Mac event hoping that more RAM will come. But we get an M1 with 16 GB max, so that’s not likely to help our local K8s installation in the short term. What else can you do? Move the dev environment to the cloud, and don’t even run Docker and K8s locally? Yes. ☁
1. [Tweet of the week](https://twitter.com/crossplane_io/status/1338906842093654017)

    Yesterday, the team at Crossplane hosted a Crossplane Community Day. Check out this tweet for an idea of what took place.

### [ << Prev ](kubelist-110.md) ------------- [ Next >> ](kubelist-112.md)