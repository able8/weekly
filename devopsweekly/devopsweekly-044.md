## [DevOps'ish 044](https://devopsish.com/044) - Sun Oct 8, 2017

I had the pleasure to travel to our South Florida office this week. I participated in five quarterly planning meetings and Bankrate’s 2nd Hackathon. The quarterly planning meetings were pretty cool. It was interesting to see what all our dev teams were going to be pushing forward with for the year. But, the highlight of the week was the Hackathon. There were some incredibly creative ideas built around the Bankrate platform that I’m sure a few will be brought into our products soon. I had the modest goal of standing up a Kubernetes cluster in AWS and deploying an example app to it. Luckily I had enough support from our team of Phippineers to make it happen. Phippineers you ask? Several members of the team have only heard me talk about Kubernetes. The first thing I asked the team to do was to watch <a href="https://kubernetes.io/blog/2016/06/illustrated-childrens-guide-to-kubernetes/">The Children’s Illustrated Guide to Kubernetes</a> to get an idea of the concepts at play. In the guide, a cute little giraffe named <a href="https://deis.com/phippy/">Phippy</a> represents a PHP app looking for a home. The team latched on to this and off we went using Terraform to roll out <a href="https://coreos.com/tectonic">CoreOS Tectonic</a> to AWS. There were some bumps along the way and at one point I was almost ready to throw in the towel. But, with some help from various <a href="https://twitter.com/ChrisShort/status/915643979470970880">CoreOS folks on Twitter</a> we got everything up and running. One of the devs on the team set out assembling a Laravel app to be deployed to the cluster. They marveled at the ability to roll out multiple versions of the app, scale up the whole cluster, and scale it back down by editing a variable here or there. It was an awesome experience! My team didn’t win any prizes or accolades but we definitely took a huge leap towards our future infrastructure.

<a href="https://docs.google.com/forms/d/e/1FAIpQLSdsxfQbVbuVVRizNaDmD1_6nyyG5WNn4pKtfHElzO9kblnz5Q/viewform"><strong>Join Our Research Group — GoCD</strong></a><br/>Take our short survey for the chance to join a great group of continuous delivery practitioners in our research group. You’ll be eligible to get your name on our contributors list, and win great schwag, and gift cards. <em>SPONSORED</em>

<strong>All Things Open</strong><br/>October 23 and 24, Raleigh, NC USA
2,500–3,000+ technologists will descend upon the City of Oaks to attend 200+ sessions from nearly as many speakers. Representative from nearly every major technology company in the U.S. will be in attendance as well.
To get 20% off enter code <strong><em>DevOpsish20</em></strong> when registering to attend.
Website: <a href="https://allthingsopen.org">https://allthingsopen.org</a>
To Register: <a href="https://allthingsopen.org/register-now">https://allthingsopen.org/register-now</a>

<strong>All Day DevOps, Live Online</strong><br/>October 24, 2017
When: October 24, 2017 (24 hours)
Where: From your desktop, laptop, or mobile device
Free Registration: All Day DevOps Registration (<a href="http://www.alldaydevops.com">http://www.alldaydevops.com</a>)

On October 24th, <a href="https://devopsish.com/">DevOps’ish</a> will be supporting the Live Online All Day DevOps Conference. This is a 24-hour event with 5 simultaneous tracks, delivering 96 sessions and 4 keynotes in 38 time zones. Session tracks include Automated Security, CI/CD, Modern Infrastructure, DevOps in Government, and the Tech Crawl, where companies will take you behind the scenes of their DevOps working environments.

### People

1. [Kubernetes Community Steering Committee Election Results](https://kubernetes.io/blog/2017/10/kubernetes-community-steering-committee-election-results/)

     Please congratulate Aaron Crickenberger, Derek Carr, Michelle Noorali, Phillip Wittrock, Quinton Hoole and Timothy St. Clair.
1. [Anxiety Update 1](https://becomingkris.com/anxiety-update-1/)

    by Kris Nova: “Most of these are NOT rational. I know that. Telling me to simply NOT worry about them will just hurt me.”
1. [Announcing more Open Source Peer Bonus winners](https://opensource.googleblog.com/2017/10/more-open-source-peer-bonus-winners.html)

     Google Open Source established this program six years ago to encourage Googlers to recognize and celebrate the external developers contributing to the open source ecosystem Google depends on.
1. [The Women Who Go Wildfire](https://code.likeagirl.io/the-women-who-go-widlfire-faf53d989d14)

    by Sarah Adams: Why has Women Who Go grown so fast? Because of the parasitic growth of mutual, undying love and respect.
### Process

1. [Are you stuck in the new DevOps matrix from hell?](http://sdtimes.com/stuck-new-devops-matrix-hell/)

    A DevOps assembly line platform can help avoid config sprawl and the DevOps matrix from hell by configuring the workflow above fairly easily, while also giving you repeatability and audit trails.
1. [Why everyone fails at monitoring; and what you can do about it](http://codearcana.com/posts/2017/10/05/why-everyone-fails-at-monitoring-and-what-you-can-do-about-it.html)

     People monitor their systems for two main reasons: to keep their system healthy and to understand its performance. Almost everyone does both wrong.
1. [PayPal API Design Guidelines](https://github.com/paypal/api-standards/blob/master/api-style-guide.md)

     PayPal is sharing these guidelines to help propagate good API design practices in general.
1. [3 billion Yahoo accounts affected by 2013 breach](https://nakedsecurity.sophos.com/2017/10/04/3-billion-yahoo-accounts-affected-by-2013-breach/)

    
1. [A Fast, Secure Migration to Google Cloud Platform using Cloudflare](https://blog.cloudflare.com/a-fast-secure-migration-to-google-cloud-platform-using-cloudflare/)

    
1. [VMware Escapology — How to Houdini the Hypervisor](https://www.zerodayinitiative.com/blog/2017/10/04/vmware-escapology-how-to-houdini-the-hypervisor)

    
1. [The Chaos Toolkit](http://chaostoolkit.org/)

     A free, open source project that enables you to create and apply Chaos Experiments to various types of infrastructure, platforms and applications.
1. [Vulnerability Note VU#973527](http://www.kb.cert.org/vuls/id/973527)

     Dnsmasq contains multiple vulnerabilities.
1. [Abandoning Iranian Nuclear Deal Could Lead to New Wave of Cyberattacks](http://foreignpolicy.com/2017/10/02/abandoning-iranian-nuclear-deal-could-lead-to-new-wave-of-cyberattacks/)

     If Trump walks away from the pact, Tehran could see “retribution against Western targets.”
### Tools

1. [Announcing Amazon EC2 per second billing](https://aws.amazon.com/about-aws/whats-new/2017/10/announcing-amazon-ec2-per-second-billing/)

     Amazon EC2 usage of Linux based instances that are launched in On-Demand, Reserved and Spot form will be billed on one second increments, with a minimum of 60 seconds.
1. [Keybase launches encrypted git](https://keybase.io/blog/encrypted-git-for-everyone)

     It is end-to-end encrypted. It’s hosted, like, say, GitHub, but only you (and teammates) can decrypt any of it. To Keybase, all is but a garbled mess. To you, it’s a regular checkout with no extra steps.
1. [Reasons Kubernetes is cool](https://jvns.ca/blog/2017/10/05/reasons-kubernetes-is-cool/)

     Kubernetes lets you do some amazing things (but isn’t easy)
1. [Start with Kubernetes in less than 5 minutes with Minikube](http://cloudmaniac.net/minikube-kubernetes-in-5-minutes/)

     Minikube creates a local, single-node Kubernetes cluster for development and testing. Setup is completely automated and doesn’t require a cloud provider account.
1. [The State of Go 1.9 — Francesc Campoy](https://www.youtube.com/watch?v=vFJkH4qDjJ0)

    
1. [Go 1.8.4 and Go 1.9.1 are released](https://groups.google.com/forum/#!msg/golang-nuts/sHfMg4gZNps/a-HDgDDDAAAJ)

     Two security-related issues were recently reported. To address this issue, we have just released Go 1.8.4 and Go 1.9.1.
1. [Compute Engine machine types with up to 96 vCPUs and 624GB of memory](https://cloudplatform.googleblog.com/2017/10/new-compute-engine-machine-types.html)

     If you need an insanely huge instance Google has you covered.
1. [PostgreSQL 10 Released](https://www.postgresql.org/about/news/1786/)

     The PostgreSQL 10 release includes significant enhancements to effectively implement the divide and conquer strategy, including native logical replication, declarative table partitioning, and improved query parallelism.
1. [OpenSSH 7.6 has just been released](http://www.openssh.com/txt/release-7.6)

     Deletes SSH protocol version 1 support, associated configuration options and documentation.
1. [twitchyliquid64/subnet](https://github.com/twitchyliquid64/subnet)

     Simple VPN server/client for the rest of us.
1. [PopSQL](https://popsql.io/)

     Modern, collaborative SQL editor for your team. Write queries, visualize data, and share your results.
1. [asciimoo/colly](https://github.com/asciimoo/colly)

     Lightning Fast and Elegant Scraping Framework for Gophers
1. [pcap2curl](https://isc.sans.edu/forums/diary/pcap2curl+Turning+a+pcap+file+into+a+set+of+cURL+commands+for+replay/22900/)

     Turning a pcap file into a set of cURL commands for “replay”
1. [cri-o v1.0.0-rc3](https://github.com/kubernetes-incubator/cri-o/releases/tag/v1.0.0-rc3)

     Close to releasing 1.0 as the number of bugs go down.
1. [Learn Go Variables — A Visual Guide](https://blog.learngoprogramming.com/learn-go-lang-variables-visual-tutorial-and-ebook-9a061d29babe)

     Easily understand Go variables with visual examples.
1. [Deploying Rancher 2.0 on Vagrant](http://www.joseluisgomez.com/automation/deploying-rancher-2-0-vagrant/)

     This post will drive you through the process to deploy without any effort a Rancher platform using Vagrant and VirtualBox.
1. [pearsontechnology/environment-operator](https://github.com/pearsontechnology/environment-operator)

     The purpose of Environment Operator is to provide a seamless application deployment capability for a given environment within Kubernetes. It can easily hook into existing CI/CD pipeline capabilities including our CI/CD pipeline as well as a typical Jenkins server through a Jenkins plugin.
1. [bitnami/sealed-secrets](https://github.com/bitnami/sealed-secrets)

     A Kubernetes controller and tool for one-way encrypted Secrets.
1. [clcollins/terraform-ansible-inventory](https://github.com/clcollins/terraform-ansible-inventory)

     An Ansible dynamic inventory script for generating an inventory from a Terraform terraform.tfstate file.

### [ << Prev ](devopsweekly-043.md) ------------- [ Next >> ](devopsweekly-045.md)