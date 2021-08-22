## [Kubelist Issue #128 for 2021-05-26](https://kubelist.com/issue/128)

#### OCI 📦

> Tune in to episode 16 of the Kubelist Podcast with Josh Dolitsky. Josh is a maintainer of the Open Container Initiative (OCI) project. There’s a surprising amount of cool work happening in this project that’s taking Docker registries and adding a pile of great new use cases on top. 

1. [Open Container Initiative](https://opencontainers.org/)

    Let’s get started this week with a link directly to the OCI website where you can learn more about the project. The Open Container Initiative project is a Linux Foundation project that has three distinct tasks: distribution, runtime and image specs. Each of these have their own release process and timelines, but they obviously work closely together. Oh, and the OCI specs aren’t static, there’s some really awesome work happening here.
1. [OCI distribution-spec v 1.0.0 release notice](https://opencontainers.org/release-notices/v1-0-0-distribution-spec/)

    Huge release from earlier this month. We all use the OCI distribution spec daily, and it’s now at a 1.0. On the podcast Josh also talks about the effort that it took to get here. Getting a widely adopted, community supported spec to a v1.0 is no small undertaking, so congrats to everyone involved! 🎉
1. [Oras](https://github.com/deislabs/oras)

    Oras is a Go-based project that is a CLI enabling you to store anything in an OCI compliant registry. With this CLI, you can “oras push” anything (a PDF, a source file, a spreadsheet, anything!) to your favorite OCI-compliant registry, and treat it effectively the same as you’d treat a Docker image. 📄
1. [OCI Images for More Than Containers](https://www.youtube.com/watch?v=ExyWAhS2zBA)

    A great talk from a recent Crossplane Community Day event showing how you can use an OCI registry for more than just docker containers. This talk has a detailed walkthrough of the different specs included in the Open Container Initiative project, and is a wealth of information. Also worth mentioning from this talk is the awesome go-containerregistry project, if you have a go project that needs to interact with a container registry.
1. [Bundle.bar](https://bundle.bar)

    A project from Josh, this is a way to use the OCI registry. Bundle.bar is a hosted service for storing small things in a registry. Small, like just a few MB. If you have a service that needs to store small objects and are considering S3, bundle.bar looks like a cool alternative that opens the door to some integrations that aren’t possible on S3. 💾
1. [Distributing WebAssembly modules using OCI registries](https://radu-matei.com/blog/wasm-to-oci/)

    Another example of non-standard use of an OCI registry. There are a ton of use cases when we think about using registries for data that isn’t a container image, but here’s a great walkthrough of using a registry to store (and distribute) WebAssembly modules. 🗞
1. [Tweet of the week](https://twitter.com/rawkode/status/1396883036508237833)

    Klustered is amazing and it’s back! Super excited to see the next episodes.

### [ << Prev ](kubelist-127.md) ------------- [ Next >> ](kubelist-129.md)