## [Kubelist Issue #12 for 2018-05-03](https://kubelist.com/issue/12)

#### kubelist issue 12 - The Stateful Edition

> Hello! This is the Stateful edition of Kubelist. So called because the resources below can help you get started on your stateful journey or help you expand your mind on how you might think about incorporating state into your Kubernetes app (AWS service broker! gVisor! Digital Ocean primitives!) .

1. [Kubernetes: The State of Stateful Apps](https://www.cockroachlabs.com/blog/kubernetes-state-of-stateful-apps/)

    A good basic overview of how and why you might consider DaemonSets and StatefulSets for managing stateful services in Kubernetes. The post doesn’t mention how Stateful services in Kubernetes deal with taints, tolerations, or antiaffinity on nodes.
1. [DigitalOcean Introduces Kubernetes Product for Simple, Scalable Container Deployment and Orchestration](https://www.digitalocean.com/press/releases/digitalocean-introduces-kubernetes-product/)

    If you are a digital ocean fan (we all have our preferred cloud providers), now you can use their managed Kubernetes that takes advantage of the Digital Ocean cloud primitives that you know and love. Check out early the access signup here.
1. [Open-sourcing gVisor, a sandboxed container runtime](https://cloudplatform.googleblog.com/2018/05/Open-sourcing-gVisor-a-sandboxed-container-runtime.html)

    Google is continuing its container security efforts. This one is really cool. A container runtime that runs containers unprivileged and provides a more secure and trusted layer of access to your systems kernel. Now that I can get into.
1. [Introducing the Operator Framework: Building Apps on Kubernetes](https://coreos.com/blog/introducing-operator-framework)

    An operator is a service you build for managing your kubernetes application. This is particular useful for stateful applications because, for instance, a postgres operator can do backups based on all kinds of conditions that represent how your operators would have done backups The operator sdk is now open source! For more on operators check this page.
1. [Developing on Kubernetes](https://kubernetes.io/blog/2018/05/01/developing-on-kubernetes/)

    Did you know you can do hot reloading with Kubernetes? Ksync looks particularly interesting! Take a look at this sweet post on boosting your Kubernetes development workflow.P.s. they included walkthroughs on some of the tools they discuss!
1. [Provision AWS Services Through Kubernetes Using the AWS Service Broker](https://aws.amazon.com/blogs/opensource/provision-aws-services-kubernetes-aws-service-broker/)

    You can provision AWS services (well, some) through Kubernetes? Sounds cool! It seems to be quite a bit of overhead to manage (a la IAM policies and set up with Kops). Wondering if AWS Kubernetes Engine will have a managed service catalog and service instance plugin all set up already so you can just write your YAML and be done with it?
1. [Tweet of the week](https://twitter.com/monadic/status/991593075184107520)

    Notice how it says 3,300 Users...

### [ << Prev ](kubelist-11.md) ------------- [ Next >> ](kubelist-13.md)