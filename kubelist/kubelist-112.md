## [Kubelist Issue #112 for 2021-01-06](https://kubelist.com/issue/112)

#### Validating For Better YAML 📄

> When running Kubernetes, the amount of YAML you’ll see on a daily basis is staggering. Joe Beda (one of the Kubernetes founders) <a href="https://twitter.com/jbeda/status/1306358044247298049">recently joked</a> about giving a talk titled “I’m sorry about the YAML”. While we can laugh about the amount of YAML, it is an error prone syntax, and it can be easy to make some mistakes. There are a few tools available to help validate and write your Kubernetes manifests. This week, we are looking at some of the methods of validating YAML before deploying, including <a href="https://www.openpolicyagent.org">Open Policy Agent</a>, and a couple of utilities from <a href="https://twitter.com/garethr">Gareth Rushgrove</a>. Hopefully the links below will help you see how to use these different tools to write production-grade Kubernetes manifests. 

1. [Integrating Open Policy Agent (OPA) With Kubernetes](https://www.magalix.com/blog/integrating-open-policy-agent-opa-with-kubernetes-a-deep-dive-tutorial)

    Let’s start with this detailed walkthrough of using the Open Policy Agent Gatekeeper project to enforce specific rules to be applied to all manifests deployed into a Kubernetes cluster. While this post starts out showing you some common scenarios including only allowed registries, the bottom shows how you can start to write your own rules with a little rego. Gatekeeper includes a rego-based Kubernetes Admission Controller, allowing you to enforce policies at deploy time. And, because it works as an admission controller, Gatekeeper policies will apply to in-cluster, Operator-created objects also. 🔱
1. [Kubeval](https://kubeval.instrumenta.dev)

    Did you know that Kubernetes publishes OpenAPI specs for all of the built-in kinds? Kubeval is a CLI that you can use to compare a given Kubernetes YAML to various versions of the Kubernetes OpenAPI specification, and look for errors and warnings. This is a good way to validate that your manifests are technically correct and will work. This is a handy CLI that will tell you when there’s an error in your YAML, such as using a number where a string is required or missing a required field. 🎓
1. [Conftest](https://www.conftest.dev)

    After using Kubeval to validate your manifests, how do you know that they are following best practices? Kubeval tells you if your manifests are correct, while Conftest can tell you if they are good. This project used Open Policy Agent, and has recently joined the OPA CNCF project. Because of the power of rego, Conftest is extremely flexible, and is a general-purpose tool that’s useful for more than just Kubernetes YAML.
1. [Validating Kubernetes YAML for best practice and policies](https://learnk8s.io/validating-kubernetes-yaml)

    A variety of tools that can be used to validate some YAML. Some of the ones mentioned here are covered in other links this week (Kubeval, Conftest). But there are some new utilities worth checking out also, and Amit describes a full set of utilities to validate Kubernetes manifests. If rego and YAML aren’t the languages you want to write rules in, check out Copper, where you can write policies in Javascript. 📐
1. [doc.crds.dev](https://doc.crds.dev)

    Kubernetes operators are popular, and most clusters have Custom Resource Definitions and custom types now. This site is a centralized location to download the OpenAPI docs for custom types. When using a utility like Kubeval, being able to provide the schemas for custom types will give you more complete coverage in your validation.
1. [Using Conftest and Kubeval With Helm](https://garethr.dev/2019/08/using-conftest-and-kubeval-with-helm/)

    Packaged as a Helm plugin, this is a way to use both Kubeval and Conftest to validate a Helm chart, ideally before deploying. You probably know that Helm charts are templated Kubernetes YAML, which makes them more difficult to lint and parse using normal YAML tools. But both Conftest and Kubeval are compatible with Helm charts, and this post shows how to use these utilities to test a chart.
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1343914259177222145)

    The CNCF Annual Report for 2020 was published. Through all of the challenges in this past year, cloud native adoption has continued to grow and at an accelerated pace!

### [ << Prev ](kubelist-111.md) ------------- [ Next >> ](kubelist-113.md)