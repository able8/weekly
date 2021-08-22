## [DevOps'ish 110](https://devopsish.com/110) - Sun Jan 13, 2019

This week I ran an unscientific Twitter poll asking, “Has DevOps left Ops behind?” It was a question that came up during a discussion with one of my co-workers.

There will be more discussion and thought around this for sure. But, it appears the software industry has gotten ahead of the skillsets of the majority of folks in it. Think about it though. We went from containers going mainstream via Docker in 2014 to Kubernetes going mainstream in 2017. Now it’s 2019 and we’re talking about Istio, serverless, and eBPF. If it seems like the pace of change is accelerating, it’s because it is.

In Thomas L. Friedman’s <a href="https://amzn.to/2SSD4mr"><em>Thank You For Being Late: An Optimist’s Guide to Thriving in the Age of Accelerations</em></a> a graph is drawn by Google X’s Eric “Astro” Teller like this one here.

<em>Note: DevOps’ish may earn compensation for sales from links on this post through <a href="../terms/">affiliate programs</a>.</em>

Friedman goes on to explain that the roles of education and government today should be to figure out how to lift the line of human adaptability to be able to keep pace with technology.

This sounds familiar… It’s no surprise that we in tech have been trying to do exactly this for years in tech with things like DevOps and SRE (and many things before that). I think the fact we’ve created “things” (DevOps, SRE, Agile, Lean, etc.) to help us adapt to change faster shows a clear sign the technology is outpacing human capacity to adapt to it. So what do we do about it? Looks like that’s the big question here.

<a href="http://www.oreilly.com/pub/cpc/187994"><strong>O’Reilly Software Architecture Conference NY 25% off</strong></a><br/>Save 25% on your Gold, Silver, or Bronze pass to the O’Reilly Software Architecture Conference in New York next month. Register using the code <strong>PCDEVOPSISH</strong>. Explore complex software architecture topics, learn best practices from industry experts, and exchange ideas and insights with your peers. <a href="http://www.oreilly.com/pub/cpc/187994">Register Now</a> <em>SPONSORED</em>

<a href="https://www.gocd.org/kubernetes"><strong>Continuous delivery on modern infrastructure - Run GoCD on Kubernetes</strong></a><br/>Model Docker-based build workflows more effectively with our GoCD Kubernetes integration. Run GoCD natively on Kubernetes, define your build workflow and let GoCD provision and scale build infrastructure on the fly. <em>SPONSORED</em>

I’m going to add this new top five from the previous newsletter. Hit reply and tell me what you think.

### Last Week’s Top Five

1. []()

    I’m going to add this new top five from the previous newsletter. Hit reply and tell me what you think.
### People

1. [Kolton Andrus, Gremlin Inc on On-Call Nightmares Podcast](https://www.podomatic.com/podcasts/oncallnightmares/episodes/2019-01-10T07_28_49-08_00)

     Kolton is the CEO of chaos engineering firm Gremlin, Inc.
1. [Jenn’s Guide to Thriving in Tech](https://dev.to/geekgalgroks/jenns-guide-to-thriving-in-tech-4k91)

     A well-rounded list of things to do to be successful in tech. One thing I need to work on is time boxing research.
1. [Mainframe brains-slurper sues IBM for ‘age discrim’, calls Ginny and biz ‘morally bankrupt’](https://www.theregister.co.uk/2019/01/07/ibm_age_discrimination_lawsuit/)

     This is disturbing to me because I think we all knew it was happening. But, no one could really do anything about it at the time.
1. [I’m a Red Hat Certified Specialist: Ansible Automation](https://chrisshort.net/im-a-red-hat-certified-specialist-ansible-automation/)

     I did a thing because I felt I had something to prove.
### Process

1. [oss-security - System Down: A systemd-journald exploit](https://www.openwall.com/lists/oss-security/2019/01/09/3)

     Patch all the 🍑
1. [Production Guideline](https://medium.com/@rakyll/production-guideline-9d5d10c8f1e)

     JBD gives us an absolutely fantastic list of things to go through before making it to production.
1. [IBM Began Buying Red Hat 20 Years Ago](https://www.linuxjournal.com/content/ibm-began-buying-red-hat-20-years-ago)

    
1. [Cloud Native Apps Need to be Managed in a Completely New Way](https://thenewstack.io/cloud-native-apps-need-to-be-managed-in-a-completely-new-way/)

     Always be shifting left.
1. [AWS gives open source the middle finger](https://techcrunch.com/2019/01/09/aws-gives-open-source-the-middle-finger/)

     Boy did AWS ever show all the license zealots out there how quickly their war was gonna last.
1. [China Is Shooting Itself in the Foot Over Huawei](https://foreignpolicy.com/2019/01/07/china-is-shooting-itself-in-the-foot-over-huawei/)

     I get it, Huawei is important to China. But, it’s getting to a point where folks are going to think twice regardless.
1. [What is DevSecOps?](https://opensource.com/article/19/1/what-devsecops)

     A nice, lightweight eBook on DevSecOps from opensource.com.
### Tools

1. [Weaveworks FlaggerFlagger Slide Deck](https://flagger.app/)

     “Flagger is a Kubernetes operator that automates the promotion of canary deployments using Istio routing for traffic shifting and Prometheus metrics for canary analysis.”
1. [Why Is Storage On Kubernetes So Hard?](https://softwareengineeringdaily.com/2019/01/11/why-is-storage-on-kubernetes-is-so-hard/)

     Storage is not trivial. Storage on Kubernetes is harder.
1. [How Tilt updates Kubernetes in Seconds, not Minutes](https://medium.com/windmill-engineering/how-tilt-updates-kubernetes-in-seconds-not-minutes-28ddffe2d79f)

     Tilt is kinda in the same vein as Skaffold. It’s a high-speed Kubernetes deployment/development tool. It introduces something they’ve dubbed a Synclet (a sidecar) that enables easy reuse of existing pods in Kubernetes as opposed to tearing the pod down, rebuilding it, and relaunching the containers in it. Pretty cool looking (I have not kicked the tires on it).
1. [5 open source tools to upgrade your next Kubernetes projectOpenFaaS](https://jaxenter.com/5-kubernetes-open-source-tools-154333.html)

     Tilt, Rook, Rancher, and Fission (I’d say )
1. [Kubernetes Authorization via Open Policy Agent](https://itnext.io/kubernetes-authorization-via-open-policy-agent-a9455d9d5ceb)

     RBAC isn’t the only auth tool in the Kubernetes world. “This blog post explains how to implement advanced authorization policies via Open Policy Agent (OPA) by leveraging the Webhook authorization module.”
1. [Adventures with a home Kubernetes cluster](https://blog.marshallbrekka.com/post/2019-01-05/home-kubernetes-cluster/)

     There’s more than one way to run Kubernetes, on ARM, in your house.
1. [Scaling Jupyter notebooks with Kubernetes and Tensorflow](https://learnk8s.io/blog/scaling-machine-learning-with-kubeflow-tensorflow)

    
1. [How to set up a Let’s Encrypt SSL Cert for Azure Web App with Linux in 5 steps](https://jessicadeen.com/how-to-manually-setup-a-lets-encrypt-ssl-cert-for-azure-web-app-with-linux/)

     “Have you ever wondered how you take a free Let’s Encrypt Certificate and use it with an Azure Web App (Linux)? This post is for you!”
1. [Google Adds DNS-over-TLS Support to Its Public DNS Service](https://www.bleepingcomputer.com/news/google/google-adds-dns-over-tls-support-to-its-public-dns-service/)

     This will come in handy.
1. [Picking communication tools for your community](https://funnelfiasco.com/blog/2019/01/07/picking-communication-tools-for-your-community/)

     I talked way too much about community tools this week. Sadly, no one tool is going to work for everyone, ever. But, there is a tool that meets the needs of most folks. This is why Jenkins is still a thing. ¯\_(ツ)_/¯
1. [New year, new GitHub: Announcing unlimited free private repos and unified Enterprise offering](https://blog.github.com/2019-01-07-new-year-new-github/)

     Free private repos with up to three contributors. That’s great but, GitHub Actions is what I’m really hoping will get better.
1. [GitHub now offers free private repos for up to three collaborators – here are our thoughts](https://about.gitlab.com/2019/01/07/github-offering-free-private-repos-for-up-to-three-collaborators/)

     Here’s what GitLab thinks of GitHub giving folks free private repos.
1. [8 Productivity Tips for GitHub](https://dev.to/_darrenburns/8-productivity-tips-for-github-44kn)

     I’m starting to get better about GitHub notifications, workflows, etc. I have a confession to make, I didn’t know what version control software was until 2011 and then it was svn. I was anti git until a few years ago. I know how to use git but, GitHub requires its own set of tools to manage.
1. [Git Explorer](https://gitexplorer.com/)

     Find the right commands you need without digging through the web.
1. [A YubiKey for iOS Will Soon Free Your iPhone From Passwords](https://www.wired.com/story/yubikey-lightning-ios-authentication-passwords/)

     iOS will need some work to make this happen. But, I literally can’t wait to be done with passwords.
1. [Update Elastic Stack with Ansible Playbooks](https://www.toptal.com/ansible/update-elastic-stack-ansible-playbooks)

    
1. [Bash-5.0 release available](http://lists.gnu.org/archive/html/bug-bash/2019-01/msg00063.html)

     Some new bits for bash.
1. [How to Monitor a Website’s Uptime for $0.75/month on AmazonPingdom is pulling the plug on their free monitoringUptimeRobot](https://danielmiessler.com/blog/how-to-monitor-a-website-for-0-75-month-on-amazon/)

     Since , here’s a reasonably priced AWS Route53 alternative. There’s also  which has a great free plan and reasonably priced entry tier.
1. [Introduction to eBPF in Red Hat Enterprise Linux 7](https://www.redhat.com/en/blog/introduction-ebpf-red-hat-enterprise-linux-7)

     eBPF, y’all.
1. [What’s New in Ansible Tower 3.4an interview with one of my all time fav PMs, Bill Nottingham](https://www.ansible.com/blog/whats-new-in-ansible-tower-3.4)

     “We’re most excited about workflows enhancements, job slices, and some other nifty features.” Here’s , about this release.
1. [Courier: Dropbox migration to gRPC](https://blogs.dropbox.com/tech/2019/01/courier-dropbox-migration-to-grpc/)

    
1. [mkcert: valid HTTPS certificates for localhost](https://blog.filippo.io/mkcert-valid-https-certificates-for-localhost/)

    
1. [Metasploit Framework 5.0 Released](https://blog.rapid7.com/2019/01/10/metasploit-framework-5-0-released/)

     I love me some Metasploit.
1. [1.9.0 — Homebrew](https://brew.sh/2019/01/09/homebrew-1.9.0/)

     “Homebrew 1.9.0 has beta support for Linux and Windows 10” YES!!! One step closer to having one Ansible Playbook for all my boxes!!!

### [ << Prev ](devopsweekly-109.md) ------------- [ Next >> ](devopsweekly-111.md)