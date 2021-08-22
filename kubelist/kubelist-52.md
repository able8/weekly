## [Kubelist Issue #52 for 2019-02-14](https://kubelist.com/issue/52)

#### To config or not to config

> That is not actually a question. You definitely have to config in Kubernetes. This week we have collected a bunch of recent articles on configuration management, security tooling, and more to help you think through your current kubernetes practice. 
A special treat this week is Lee Brigs. Why the fuck are we templating yaml? Which is well supplemented by background in The State of Kubernetes Configuration Management: An Unsolved Problem.

1. [Fully Automatic Certificate Bootstrapping for Kubelets](https://blog.kontena.io/fully-automatic-certificate-bootstrapping-for-kubelets/)

    Hi Pharos, thanks for making this. Secure Kubelets are important. Double that shoutout to the operator-sdk, it does make operator making easy, it does.
1. [Deploy your node.js app to Kubernetes with a single command. No config required.](https://github.com/kubesail/deploy-to-kube)

    Why we think this is cool: Not everyone should be forced to yaml in order kube
1. [Autocert](https://github.com/smallstep/certificates/blob/master/autocert/README.md)

    mmmmmmTLS, cert injection with the push of an… annotation. Great! With all this new and easy to use security tooling we think Kubernetes security is really upping its game.
1. [Kuberhealthy](https://github.com/Comcast/kuberhealthy)

    There is a lot that can go wrong with all the moving parts of Kubernetes. This handy tool builds up a set of tests around those more annoying things that go wrong. Plus it has a serviceMonitor+Grafana dashboard to make life even easier.
1. [Jib 1.0.0 is GA—building Java Docker images has never been easier](https://cloud.google.com/blog/products/application-development/jib-1-0-0-is-ga-building-java-docker-images-has-never-been-easier)

    Java containers build slow. Agreed. Java dockerfiles are messy. Agreed. Want pretty java dockerfiles with minimial effort? Agreed. See Jib.
1. [Why the fuck are we templating yaml?](https://leebriggs.co.uk/blog/2019/02/07/why-are-we-templating-yaml.html)

    Interesting thoughts about JSON, YAML, and templating with a cool interlude into the world of jsonnet. You may not agree (warning to Helm lovers) but you will definitely leave intrigued.
1. [The State of Kubernetes Configuration Management: An Unsolved Problem](https://blog.argoproj.io/the-state-of-kubernetes-configuration-management-d8b06c1205)

    Great overview of the different config mgmt tools available to us with solid pros and cons. We have been using Kustomize a lot recently and loving it, what about you?
1. [Tweet of the Week](https://twitter.com/dankohn1/status/1094937269687984130)

    Perfect time to become a Kubekid.

### [ << Prev ](kubelist-51.md) ------------- [ Next >> ](kubelist-53.md)