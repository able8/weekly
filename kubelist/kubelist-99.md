## [Kubelist Issue #99 for 2020-09-10](https://kubelist.com/issue/99)

#### All About Registries 📒

> If you are subscribing to Kubelist, you probably agree that containers have changed how we deploy software. A large enabler of containers was the creation of the Docker (now OCI) Registry. Container registries aren’t boring, and there are a lot of changes happening as the basic container registry grows into an OCI registry. This week, we take a close look at some interesting registry-related links.

1. [OCI Registry As Storage](https://github.com/deislabs/oras)

    The ORAS project, created by Deis/Microsoft, is one of the most unexpected but interesting uses of the OCI registry we’ve seen. This project lets you use an OCI compatible registry (many today are) for something that’s similar to object store. Is an OCI registry ready to replace S3 for some use cases? 🤔
1. [Cloud Native Artifact Registries evolve from Docker Container Registries](https://stevelasker.blog/2019/01/25/cloud-native-artifact-stores-evolve-from-container-registries/)

    This is a blog post that gives more background on the ORAS project. If you aren’t sure why you’d want to use an image registry for storage, give this a read. As the post says: “If containers are becoming the common unit of deployment for software, why not use a registry to store, secure and maintain these new artifacts?” 
1. [OCI Artifact Support In Amazon ECR](https://aws.amazon.com/blogs/containers/oci-artifact-support-in-amazon-ecr/)

    If you need validation that most Docker registries will evolve to support OCI artifacts (ORAS, Helm Charts and more), see this recent post from Amazon. The Amazon managed ECR (Elastic Container Registry) now supports OCI artifacts, enabling you to use these new tools in your EKS cluster (or any other Amazon-local workload) while relying on IAM Instance Roles for authentication. 
1. [Push and pull an OCI artifact using an Azure container registry](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-oci-artifacts)

    Azure has a walkthrough showing how to push and pull an OCI artifact to their Azure hosted registry (ACR). This uses the ORAS utility (see above) to push generic artifacts and store them in the registry. Also worth pointing out: ACR has supported OCI artifacts for a while. It’s getting harder to find registries that aren’t OCI compatible.
1. [Introducing GitHub Container Registry](https://github.blog/2020-09-01-introducing-github-container-registry/)

    GitHub hosting a container registry makes sense, right? Having this tightly coupled to the source code and specifically to the organization creates some interesting possibilities to ensure norms and common practices are maintained. It’s great to see GitHub embracing Docker registries, and we hope they introduce OCI Artifact support soon!
1. [DockerHub Update FAQ](https://www.docker.com/pricing/resource-consumption-updates)

    A lot has been written, shared and tweeted about the changes to DockerHub’s retention policies and pull limits. The changes create some uncertainty, but head over to this FAQ if you still have questions or are thinking about setting up a cron to pull your images on a regular basis.
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1303658173702852608?s=20)

    So much fresh content to start downloading and watching!

### [ << Prev ](kubelist-98.md) ------------- [ Next >> ](kubelist-100.md)