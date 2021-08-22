## [Kubelist Issue #134 for 2021-07-14](https://kubelist.com/issue/134)

#### K3S, Keith Basil, And Removing The Cognitive Load Of Kubernetes 🧠

> On this week’s episode of the <a href="https://kubelist.com/podcast">Kubelist podcast</a>, Keith Basil from SUSE (via Rancher) joined to talk about K3s, simplifying software, and much more. We talked about other projects (Longhorn, RKE) but really kept coming back to the impressive K3s project. Personal highlight was running K3s on Raspberry PIs in SPACE in a Satellite to perform image recognition against a baby Yoda doll... what else is there to say?

1. [K3S](https://k3s.io)

    Now in the sandbox, k3s is already seeing great adoption and is considered a mature project. A simple, conformant, production-grade Kubernetes distribution that runs in a single Go binary.
1. [RKE2](https://docs.rke2.io)

    K3s is a popular and a great simple distribution, but don’t forget about RKE2. This is a secure, compliant, conformant distribution of Kubernetes that’s designed for really sensitive environments. Don’t discount the effort in this project–the team at Rancher has had to use a different Go compiler in order to get FIPS compliant. 🏛
1. [Longhorn](https://longhorn.io)

    Discussed in the podcast this week, Longhorn is another CNCF Sandbox project from SUSE, designed to simplify distributed block storage in Kubernetes. This is an exceptionally hard problem that does get in the way of Kubernetes adoption. Super cool project. 🐮
1. [Kine](https://github.com/k3s-io/kine)

    Kine is not etcd (get it). This is one of the big blocks that makes K3s possible. Kine is a shim layer that lets your application write to etcd, but kine will store the data in a different backend database engine such as Postgresql or MySQL. 🛢
1. [Rancher Desktop](https://rancherdesktop.io)

    Basil talks about Rancher Desktop and the work happening here to make it super easy for developers to install and operate a local Kubernetes cluster for their dev environment. 
1. [Introduction to k3d: Run K3s in Docker](https://www.suse.com/c/introduction-k3d-run-k3s-docker-src/)

    A quick intro to a related (sub?) project of K3s. K3d is a project that runs K3s in a Docker container. This post walks you through a quick start to get a Kubernetes cluster running in K3d.  We covered this a good amount back in issue 118.
1. [Tweet of the Week](https://twitter.com/puerco/status/1414784210800422929)

    This is huge. A Software Bill Of Materials (SBOM) will start to be included with future Kubernetes releases!

### [ << Prev ](kubelist-133.md) ------------- [ Next >> ](kubelist-135.md)