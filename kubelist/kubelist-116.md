## [Kubelist Issue #116 for 2021-02-24](https://kubelist.com/issue/116)

#### Crossplane 🍭

> Today’s episode of the <a href="https://kubelist.com/podcast">Kubelist Podcast</a> is with <a href="https://twitter.com/hasheddan">Daniel Mangum</a> from <a href="https://upbound.io">Upbound</a>, discussing the <a href="https://crossplane.io">Crossplane</a> project. Crossplane is a Sandbox project that allows users to define Compositions of infrastructure and then deploy the infrastructure right alongside an application. This is a really cool approach to managing infrastructure as data. In this week’s newsletter, we share some links for getting started and stories from users using Crossplane.

1. [Modernize with API driven Control Planes featuring Bassam Tabbara, Upbound](https://www.youtube.com/watch?v=-t0wCOH2X7I&feature=youtu.be)

    Let’s start with an excellent overview video from Bassam Tabbara, founder of Upbound. Here, Bassam explains the value of having the API-driven control plane that Crossplane provides. This is a great introduction and a higher-level discussion about the patterns that Crossplane provides, and the reasons for them. If you’re curious about why you should dig into Crossplane, this is a good place to start.
1. [The Power of Control Planes and the Kubernetes Resource Management Model](https://youtu.be/WGfYlssfIIk)

    This is a great conversation we had to share. Some of the founders and leaders from the Kubernetes ecosystem joined Bassam Tabbara on a panel discussion to talk about the concepts that Crossplane uses and implements. The conversation is between Kelsey Hightower from Google, Brian Grant from Google, Brendan Burns from Microsoft Azure, Joe Beda from VMware, and Bassam Tabbara from Upbound, and it took place on Crossplane Community Day at the end of 2020. 💬
1. [Crossplane vs Cloud Provider Infrastructure Addons](https://blog.crossplane.io/crossplane-vs-cloud-infrastructure-addons/)

    Here’s a blog post comparing design decisions of Crossplane to other projects in the Kubernetes ecosystem for managing infrastructure; including GCP Config Connector, AWS Controllers for Kubernetes, and more. If you are familiar with the native work that any of the cloud providers are doing in this space, this post will be a good read to understand the differences and value that Crossplane is providing and why Crossplane is worth a look. ☁
1. [TBS[25]: Composition 101](https://www.youtube.com/watch?v=f_2f1u-XOgA&feature=youtu.be)

    The Binding Status (hosted by our guest this week, Dan Mangum) had a recent episode featuring a nice overview of how Composition works in Crossplane. We think that Composition is one of the most important features of the project, and what makes Crossplane stand out when compared to the cloud provider’s native connectors for Kubernetes.
1. [TBS[26]: Provider Authentication](https://www.youtube.com/watch?v=tUOZ09xX7ng&feature=youtu.be)

    Also on The Binding Status, Dan shows how to use Crossplane providers with Vault to supply credentials to an application. This is a great ‘in the weeds’ example to help illustrate how Crossplane works at all levels of the application stack; from provisioning that Postgres database and a VPC (that we also talk about on the Podcast today), to making secrets easier to consume when a Pod starts. 🔐
1. [Integrating Crossplane with Cluster API](https://www.youtube.com/playlist?list=PL510POnNVaaazaBA7jNh1f6Qz1LxuB5Ku)

    A cool part of Crossplane that we didn’t talk too much about on the Podcast is that it can do more than provision databases and infrastructure. Crossplane can even provision a Kubernetes cluster. This link is a longer (4 part) series showing the work done to integrate Crossplane with Cluster API to make this happen. 🏗
1. [Tweet of the week](https://twitter.com/cloudnativefdn/status/1364258972606341125)

    CNCF End User Technology Radar is out sharing some information about how different end users are handling secrets in their Kubernetes clusters.

### [ << Prev ](kubelist-115.md) ------------- [ Next >> ](kubelist-117.md)