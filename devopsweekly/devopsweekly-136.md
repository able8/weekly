## [DevOps'ish 136](https://devopsish.com/136) - Sun Jul 14, 2019

<a href="https://www.ibm.com/blogs/think/2019/07/ibm-red-hat/">Q&amp;A: IBM’s Landmark Acquisition of Red Hat</a>. It is done. Aside from extra meetings from too many Red Hat folks, I have seen zero change. I’d like to talk more about the whole acquisition process from my point of view since it was announced, at some point. Overall, from my perspective, the process was smooth. I commend Red Hat and IBM leadership for keeping the chaos to a minimum. Here’s to next week and hopefully fewer meetings!

<a href="https://www.oreilly.com/pub/cpc/238261"><strong>10 Kubernetes distributions leading the container revolution</strong></a><br/>Kubernetes and containers are changing how applications are built, deployed, and managed. Check out the distros leading the charge. <em>SPONSORED</em>

### DevOps’ish Last Week’s Top Five

### People

1. [DevOps for introverted people](https://opensource.com/article/19/7/devops-introverted-people)

     ”We asked the Opensource.com DevOps team to talk about their experience as DevOps introverts and to give DevOps extroverts some advice. Here are their answers.” Dan Barker, Catherine Louis, Abhishek Tamrakar, Elizabeth Joseph, and I were all quoted in this piece. It turned out really good.
1. [Here are the most popular programming languages used by the world’s largest unicorn startups](https://www.geekwire.com/2019/popular-programming-languages-used-worlds-largest-unicorn-startups/)

     If you’re looking to pick up a new language, maybe start with this list.
1. [The case for making the transition from sysadmin to DevOps engineer](https://opensource.com/article/19/7/devops-vs-sysadmin)

     ”There’s a learning curve, but there’s no time like the present to get started.”
1. [Five Things Developers Can Use to Go On-Call (And Not Hate Their Life)](https://channel9.msdn.com/Shows/5-Things/Five-Things-Developers-Can-Use-to-Go-On-Call-And-Not-Hate-Their-Life)

    
1. [Emotional exhaustion](https://en.wikipedia.org/wiki/Emotional_exhaustion)

     It’s a real thing.
1. [Opioid rules reassessed amid outcry from patients needing painkillers](https://www.usatoday.com/story/news/health/2019/07/12/opioid-rules-reassessed-amid-outcry-patients-needing-painkillers/1705026001/)

     Too often when a few people get together and think they’re doing the right thing, they’re harming another group of people. The changes to opioid laws of late have definitely increased the difficulty level on life for me and undoubtedly countless others. Make high impact decisions in public, please.
### Process

1. [No Slacking: Microsoft Teams Announces 13M DAUs, Mattermost Raises At Rapid Cliptweeted in relation to this story](https://news.crunchbase.com/news/no-slacking-microsoft-teams-announces-13m-daus-mattermost-raises-at-rapid-clip/)

     Microsoft Teams is so great that numerous teams at Microsoft flat out don’t use it. The edict from on high mentioned a few weeks ago to force folks out Microsoft products? Apparently, it had no teeth. Sadly there will be no getting away from Teams if you’re an Office 365 company it appears (hard to beat free). Mattermost CEO, Ian Tien, , “Messaging Collaboration is a $30B market—as big as email and web conferencing—with 3 key segments: First mover (Slack), incumbent (Microsoft Teams), and open source (Mattermost). Open source offers trust, flexibility and innovation like no other. Let’s see what happens next :)”
1. [Canonical Ltd source code repositories have been compromised](https://news.ycombinator.com/item?id=20369902)

     ”We can confirm that on 2019-07-06 there was a Canonical owned account on GitHub whose credentials were compromised and used to create repositories and issues among other activities. Canonical has removed the compromised account from the Canonical organisation in GitHub and is still investigating the extent of the breach, but there is no indication at this point that any source code or PII was affected.”
1. [Seriously, stop using RSA](https://blog.trailofbits.com/2019/07/08/fuck-rsa/)

     Originally titled, “Fuck RSA”, this article pulls no punches. “RSA is an intrinsically fragile cryptosystem containing countless foot-guns which the average software engineer cannot be expected to avoid.”
1. [Crypto-Mining Malware Outsmarting Image Scanners](https://blog.aquasec.com/crypto-mining-malware-container-security)

     It’s interesting to see everyone’s fears about sharing Vagrant images come to fruition with containers. You don’t need the whole OS to compromise a line of code (or infrastructure as code).
1. [Elastic CEO says Amazon Web Services isn’t slowing down its business](https://www.businessinsider.com/elastic-shay-banon-amazon-web-services-elasticsearch-2019-7)

     I’m a little worried about Elastic but, it doesn’t seem the folks there are. To be honest, I’m worried about any company that has to monkey around with open source licensing to protect their business model, Elastic or otherwise.
1. [Details of the Cloudflare outage on July 2, 2019](https://blog.cloudflare.com/details-of-the-cloudflare-outage-on-july-2-2019/)

     I’ve moved almost all of my web sites off Cloudflare’s CDN (hello new regions of the world!) but, their DNS lookup times are top-notch. This was a mess due to a configuration change. Guess what? Everyone tests in Prod.
1. [Blue Matador Closes $3.1 Million in Seed Funding to Transform Cloud Infrastructure Monitoring](https://finance.yahoo.com/news/blue-matador-closes-3-1-150000613.html)

     Congrats to my friends over at Blue Matador!
1. [A Very Cold Take on IBM, Red Hat and Their Hybrid Cloud Hyperbole](http://www.platformonomics.com/2019/07/a-very-cold-take-on-ibm-red-hat-and-their-hybrid-cloud-hyperbole/)

     Maybe a somewhat hyperbolic take on hybrid cloud hyperbole.
1. [DevOps’ish Deep Cuts 001: A bluer hue of red](https://devopsish.com/deep-cuts-001/)

    
### Tools

1. [Fire Up Your VMs with Weave Ignite](https://www.weave.works/blog/fire-up-your-vms-with-weave-ignite)

     ”Weave Ignite is an open source VM with a container UX and built-in GitOps management. [It] combines AWS Firecracker MicroVMs with Docker / OCI images to unify containers and VMs. [It] Works in a GitOps fashion and can manage VMs declaratively and automatically like Kubernetes and Terraform.”
1. [Developers don’t understand CORS](https://fosterelli.co/developers-dont-understand-cors)

     CORS is something I’m guaranteed to RTFM before even thinking about touching anything remotely close to cross origin related. The security implications are too high.
1. [Kubernetes productivity tips and tricks](https://www.padok.fr/en/blog/kubernetes-productivity-tips)

     ”Here’s are Kubernetes tips and tricks to code and deploy faster.”
1. [9 Steps to Awesome with Kubernetes](https://learning.oreilly.com/live-training/courses/9-steps-to-awesome-with-kubernetes/0636920283713/)

     Live Training from fellow Red Hatter, Burr Sutter. “For Developers, learn nine simple and practical steps that will take you from Kubernetes novice to cloud native application architect.”
1. [Dissecting the HAProxy Kubernetes Ingress Controller](https://www.haproxy.com/blog/dissecting-the-haproxy-kubernetes-ingress-controller/)

    
1. [Terraform: Up & Running, 2nd edition Early Release is now available!](https://blog.gruntwork.io/terraform-up-running-2nd-edition-early-release-is-now-available-b104fc29783f?gi=95020ed22d5)

     I know a lot of folks have been waiting very patiently for this.
1. [A simple example of git bisect command](https://www.pixelstech.net/article/1562942424-A-simple-example-of-git-bisect-command)

     I don’t think any solution to a problem as complex doing anything in git history should be called “simple” but, maybe that’s just me.
1. [There’s more than one way to write an IP address](https://ma.ttias.be/theres-more-than-one-way-to-write-an-ip-address/)

     All the fun ways to type an IP on the Linux CLI.
1. [The Illustrated TLS Connection: Every Byte Explained](https://tls.ulfheim.net/)

    
1. [Ramblings from Jessie: Linux Observability with BPF](https://blog.jessfraz.com/post/linux-observability-with-bpf/)

     Jessie Frazelle wrote the foreword to an upcoming book on BPF.
1. [Announcing AWS Toolkit for Visual Studio Code](https://aws.amazon.com/blogs/developer/announcing-aws-toolkit-for-visual-studio-code/)

     AWS Serverless Application Model (AWS SAM) focused extension for VSCode
1. [EKSphemeral](https://eksphemeral.info/)

     ”Meet EKSphemeral, the simple manager for ephemeral EKS clusters, allowing you to launch EKS clusters that auto-tear down after some time, and you can also prolong their lifetime if you want to continue to use them.”
1. [Announcing Docsy: A Website Theme for Technical Documentation](https://opensource.googleblog.com/2019/07/announcing-docsy-website-theme-for.html)

    
1. [crashloopbackoff.dev](https://crashloopbackoff.dev)

     How to describe CrashLoopBackOff in a way people understand (maybe ImagePullBackOff more accurately though).
1. [knadh/listmonk](https://github.com/knadh/listmonk)

     High performance, self-hosted newsletter and mailing list manager with a modern dashboard (as the maintainer of a newsletter, I’m VERY interested in this project)
1. [stakater/Reloader](https://github.com/stakater/Reloader)

     A Kubernetes controller to watch changes in ConfigMap and Secrets and then restart pods for Deployment, StatefulSet, DaemonSet and DeploymentConfig – [✩Star] if you’re using it!
1. [hjacobs/kube-ops-view](https://github.com/hjacobs/kube-ops-view)

     Kubernetes Operational View - read-only system dashboard for multiple K8s clusters
1. [dropbox/goebpf](https://github.com/dropbox/goebpf)

     Library to work with eBPF programs from Go
1. [edgelevel/lastpass-operator](https://github.com/edgelevel/lastpass-operator)

     A Kubernetes Operator to manage secrets stored in LastPass password manager
1. [paypal/hera](https://github.com/paypal/hera)

     Hera multiplexes connections for MySQL and Oracle databases. It supports sharding the databases for horizontal scaling.
1. [chris-short/DevOps-README.md](https://github.com/chris-short/DevOps-README.md)

     What to Read to Learn More About DevOps
1. [poseidon/typhoon](https://github.com/poseidon/typhoon)

     Minimal and free Kubernetes distribution
1. [here](./notes/)

    Notes from this week’s issue can be found .
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-135.md) ------------- [ Next >> ](devopsweekly-137.md)