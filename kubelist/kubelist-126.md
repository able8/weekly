## [Kubelist Issue #126 for 2021-05-05](https://kubelist.com/issue/126)

#### Operator Framework ⚡️

> On this week’s <a href="https://kubelist.com/podcast">Kubelist Podcast</a> we sit down with Evan Cordell from RedHat and chat about the past, present and future of the incubating <a href="https://operatorframework.io/">Operator Framework project</a>. For the newsletter, we’re going to take a deep dive into operators and how they fit in the Kubernetes ecosystem; and look at the challenges and tools associated with them. Make sure to checkout <a href="https://operatorhub.io/">Operatorhub</a> for a list of what&#39;s already out there!

1. [Introducing Operators: Putting Operational Knowledge into Software](https://web.archive.org/web/20170224092119/https://coreos.com/blog/introducing-operators.html)

    We’re going to use the Wayback Machine to check out when Brandon Phillps and the CoreOS(RIP) team introduced us to operators! It’s still an awesome post that does a great job of outlining what operators can do, what they are for, and how they came to be. ✨
1. [TOC approves Operator Framework as Incubating Project](https://www.cncf.io/blog/2020/07/09/toc-approves-operator-framework-as-incubating-project/)

    Four years after the introduction of operators, RedHat donated the Operator Framework to the CNCF. It consists of two main components: Operator SDK and Operator Lifecycle Manager, ostensibly all the tooling you need to create and manage your own Operator.
1. [An in-depth walkthough of building and running a Go-based operator](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/)

    This tutorial guides you through setting up your own Operator and leveraging Operator SDK. Some GO knowledge is required, but it is quite educational and shows you fundamentally how to leverage operators for your day-1 and day-2 needs through the lens of Memcached. 🔍
1. [Operator Lifecycle Manager: Getting Started](https://olm.operatorframework.io/docs/getting-started/)

    Once you build your own Operator, you have to maintain and upgrade it–and that's where OLM comes in. We take you straight to the docs to get you going. By no means is OLM just about upgrading; it also helps you manage dependencies, discoverability, and stability for all of your operators. Bonus: Check out the community guiding this project! 🪄
1. [Build Your Kubernetes Operator with the Right Tool](https://hazelcast.com/blog/build-your-kubernetes-operator-with-the-right-tool/)

    There is more than one way to build an Operator, and here’s a nice write up that goes through a large chunk of your options. They dive into language choices and framework, along with handy code snippets and real-world implementations. 🛠
1. [Introducing Kubebuilder: an SDK for building Kubernetes APIs using CRDs](https://kubernetes.io/blog/2018/08/10/introducing-kubebuilder-an-sdk-for-building-kubernetes-apis-using-crds/)

    When one discusses operators, one must discuss CRDs! Evan brings up Kubebuilder, and it's a great project to take a look at to see if it’s helpful to you! You can jump straight to the getting started, or read up on the thinking behind the project!
1. [Tweet of the week](https://twitter.com/andrew_randall/status/1389872432148799488)

    Really cool. The sandbox is growing quickly and this was amazing to watch. We’ll share the recording link when it’s available, because the whirlwind tour of the sandbox was amazing.

### [ << Prev ](kubelist-125.md) ------------- [ Next >> ](kubelist-127.md)