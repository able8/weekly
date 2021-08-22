## [Kubelist Issue #127 for 2021-05-12](https://kubelist.com/issue/127)

#### Always be thinking about security 🏰

> We started and wrapped up Kubecon EU last week, and what a conference it was. Between virtual CTF (Capture The Flag) and great keynotes, there was a lot to take in and digest! In this week’s newsletter, we’re sharing some of our favorite moments from Kubecon. Make sure to let us know on <a href="https://twitter.com/readkubelist">Twitter</a> if we missed anything!

1. [Cloud Native Security Day, Capture The Flag](https://www.youtube.com/watch?v=bFyYaECAPpo)

    Watch the pros try to hack around a Kubernetes cluster. If you’ve ever been curious how someone will try to infiltrate your cluster, watching this will be eye-opening. Folks were tasked with finding a hidden virtual flag in a cluster, and they had to search and hack around in order to find it. The conversation is informal and insightful. ⛳️
1. [Google’s kCTF](https://google.github.io/kctf/)

    If you want to start experimenting or play around with a CTF setup, but aren’t ready to join a public event yet, this is an interesting project from Google. You can just click a button and get a private cluster set up with a capture the flag competition. If you want to build a custom event, this infrastructure and setup are open source and you can build on it.
1. [Tutorial: Attacking and Defending Kubernetes](https://www.youtube.com/watch?v=UdMFTdeAL1s)

    Here’s a classic from KubeCon 2019 where you can watch a talented group of folks attacking and defending a cluster, live. If you have a Kubernetes cluster running on the internet, you should watch this. Someone either is or has attempted to break into the control plane, and while this is almost 2 years old, it’s still incredibly relevant! ⚔️
1. [sigstore demo with cosign](https://www.youtube.com/watch?v=gCi9_4NYyR0)

    If you haven’t heard, sigstore is a new project in the Linux Foundation to focus on supply chain. This is a video of one of the sigstore engineers sharing a live demo of using cosign and sigstore. This is technical, but it’s a deeply technical topic. Cosign doesn’t require any extra servers or storage, and looks to be a great way to verify that an image wasn’t tampered with after pushing, but before pulling. ⛓
1. [Verifying Container Image Signatures from an OCI Registry in Kubernetes](https://blog.sigstore.dev/verify-oci-container-image-signatures-in-kubernetes-33663a9ec7d8)

    Now that we’ve seen sigstore and cosign, you might be asking “how do I configure this to be part of my cluster automatically?” Glad you asked! This is a walk through of how to configure Kubernetes to verify image signatures automatically. 🖋
1. [Ensure Content Trust on Kubernetes using Notary and Open Policy Agent](https://siegert-maximilian.medium.com/ensure-content-trust-on-kubernetes-using-notary-and-open-policy-agent-485ab3a9423c)

    Notary is a mature CNCF project that is solving container image signing and part of the software supply chain around container images. This blog post shows how you can use Open Policy Agent to ensure that the images you are pulling aren’t tampered with.
1. [Tweet of the week](https://twitter.com/CloudNativeBot/status/1392117290284560388)

    Have you recovered from KubeCon EU? I hope so, because it’s time to get a proposal submitted for KubeCon US that’s happening in October!

### [ << Prev ](kubelist-126.md) ------------- [ Next >> ](kubelist-128.md)