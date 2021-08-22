## [Kubelist Issue #129 for 2021-06-02](https://kubelist.com/issue/129)

#### New from the CNCF Sandbox 🏖

> It’s time again to do a roundup of new and exciting projects the CNCF TOC (Technical Oversight Committee) have recently accepted into the Sandbox. As always this isn’t an exhaustive list of what’s been added, so check out the <a href="https://www.cncf.io/sandbox-projects/">CNCF Site</a> for the latest list (and subscribe to the <a href="https://kubelist.com/podcast">Kubelist Podcast</a>).

1. [Antrea](https://blogs.vmware.com/opensource/2021/05/05/antrea-joins-cncf-sandbox/)

    Networking is hard, networking securely is very hard, networking securely with high performance is almost impossible. Thankfully we have Antrea implementing Container Network Interface (CNI) and Kubernetes NetworkPolicy at the pod level. Antrea is based on Open vSwitch project (OVS) which basically means it will run anywhere (Kernel >= 3.3+) and even works for Windows! The overlay is top notch, can’t say enough good things about this project! 🌐
1. [ChaosBlade](https://www.alibabacloud.com/blog/chaosblade---an-open-source-chaos-engineering-tool-by-alibaba_594850)

    Netllix introduced us to the concept of Chaos Engineering in 2011 and it has quickly become the de facto practice for many SRE teams. Tolerant failures is ultimately a big reason why you want to leverage tools such as kubernetes! ChaosBlade has some interesting ways to create specific scenarios and aims to make Chaos more accessible to the community as a whole. 🗡
1. [WasmEdge Runtime](https://wasmedge.org/)

    Security and portableness is what sticks out to us when looking at WebAssembly (WASM). WasmEdge is promising to take these concepts to the portable extremes, enabling you to run Wasm bytecode in a fully sandboxed and memory-safe environment that includes “(e.g., a SaaS, a car OS, an edge node, or even a blockchain node)”. If you are looking at k8s for more exotic edge-based locals this project seems like a great place to start. 🔒
1. [Vineyard](https://v6d.io/)

    Have you ever wanted to share data between two processes in memory but didn't know where to start? Vineyard is aiming to solve this problem with “in-memory immutable data manager that provides out-of-the-box high-level abstraction and zero-copy in-memory sharing for distributed data in big data tasks”. High level share your panda data frames without all the ETL! 🌱
1. [Fluid](https://github.com/fluid-cloudnative/fluid)

    A lot of companies are starting to move their DeepLearning and more specifically their datasets to k8s. Fluid aims to make the transition and manipulation of your data smoother by offering “Distributed Dataset Orchestrator and Accelerator for data-intensive applications''. Personally I am really excited to try out their caching and data-warming features and look forward to what this could turn into. 💧
1. [Submariner](https://submariner.io/getting-started/architecture/)

    As the complexity of your requirements grows (edge compute, fault isolation, region redundancies etc.) so will the number of disparate k8s clusters you manage. Creating cross-cluster connectivity is what Submariner does. It offers L3 connectivity, Service Discovery across clusters, and even its own CLI. While this project is in alpha, it seems well-baked and is one to keep an eye on! ⚓️
1. [Tweet of the week](https://twitter.com/SteveSmith_Tech/status/1399730488143659014)

    An interesting take on GitOps. Configuration management and infrastructure tools can often bring out some strong opinions and this is not different. There are some interesting comments in the thread, but is this the reason you haven’t adopted GitOps yet? Is this a misunderstanding of the value of GitOps or a good summary of the pattern?

### [ << Prev ](kubelist-128.md) ------------- [ Next >> ](kubelist-130.md)