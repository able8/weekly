## [Kubelist Issue #97 for 2020-08-26](https://kubelist.com/issue/97)

#### Developing in a Kubernetes environment 🖼

> As we build more Kubernetes-native applications and run more and more on Kubernetes, developers need to be able to have an environment that closely resembles production. This means Kubernetes for developers and Kubernetes-friendly development environments. This week, we are sharing a few of our favorite projects that help developers build for Kubernetes environments.

1. [Blimp](https://kelda.io/blog/blimp-introduction/)

    Blimp is a drop-in replacement for docker-compose that runs your development environment on Kubernetes in the cloud. It's fully backward compatible with Docker Compose, the most popular local development tool today. Take any existing docker-compose.yml file, run blimp up, and get a development experience with infinite resources at your disposal. ☁
1. [Garden.io](https://garden.io/)

    Garden is a cloud native testing platform that lets developers run fast integration tests and spin up production-like preview environments for Kubernetes, on demand. Garden’s  mission is to ensure that developers get fast test feedback when they need it most, and that DevOps teams no longer have to spend their time provisioning and managing ad hoc environments. 🌷
1. [Tilt](https://tilt.dev/)

    The team at Tilt identified that developing on Kubernetes is far from perfect, and set off to build some tools to help. Tilt excels at quickly updating microservices into a cluster, and its slick UI can help your team correct failing services without having to “play 20 questions with kubectl”. Open your editor, write some code, and your Starlark-powered Tiltfile will shorten the iteration loop for your team. 🧰
1. [Skaffold](https://skaffold.dev/)

    Skaffold is a CLI to set up a local dev environment that watches your local filesystem and redeploys to the cluster when you save. It’s a little more bare-bones than Tilt, in that you need to figure out how to build optimized Dockerfiles if you want a quick deployment cycle. 
1. [Telepresence](https://www.telepresence.io/)

    Telepresence is a little different from the other tools in this space. With telepresence, you don’t need Kubernetes to build a Kubernetes application. With some clever networking and port forwarding, you can “check out” a service from a cluster and have it routed to your local dev environment, even if you aren’t running your local environment in Kubernetes.
1. [Loft.sh](https://loft.sh/)

    Different from the dev tool, loft.sh is the updated version of devspace.cloud, a managed “namespace-as-a-service” Kubernetes offering to give easy dev environments to all developers. Once you pick one of the tools above, consider loft.sh to run it. 🏠
1. [Tweet of the week](https://twitter.com/kelseyhightower/status/1298343780307755008)

    Kelsey Hightower points to the importance of the Kubernetes Resource Model - a consistent and interoperable pattern for interacting with Kubernetes objects.

### [ << Prev ](kubelist-96.md) ------------- [ Next >> ](kubelist-98.md)