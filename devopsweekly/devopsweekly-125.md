## [DevOps'ish 125](https://devopsish.com/125) - Sun Apr 28, 2019

What a series of unfortunate events for <a href="http://petty.company"><strong>Docker</strong></a> in 2019. In what appeared to be a massive talent flush due to what looks like a potential earnings miss, the Great Docker Culling of 2019 happened. Docker appears to have laid off the vast majority of its well-known talent. <a href="https://twitter.com/aluzzard">Andrea Luzzardi</a>, <a href="https://www.linkedin.com/in/samalba/">Sam Alba</a>, and <a href="https://twitter.com/garethr">Gareth Rushgrove</a> are among a slew of recent <a href="https://devopsish.com/117/">Docker layoffs</a> discussed in this newsletter earlier this year. According to one source teams were, “killed,” and Docker, “missed their number, and by a lot.”

Fast forward to Friday night on the US east coast (like we weren’t going to notice?!?). Many people (myself included) <a href="../pdf/docker-hub-breach-email.pdf">received an e-mail from Docker</a> about a <strong>Docker Hub breach</strong> impacting at least 190,000 accounts. According to the e-mail, “Data includes usernames and hashed passwords… as well as Github and Bitbucket tokens for Docker autobuilds.” Audit any Docker Hub tokens right now. Docker also, “revoked GitHub tokens and access keys. This means your autobuilds will fail.” Nothing like a page on a weekend because Docker broke your builds. Check your <a href="https://cloud.docker.com/settings">Docker Hub Linked Accounts</a> and re-link them. You’ll then likely have to do a weird do-si-do in the Build config of one of your image pages to get everything working as is.

This Docker Hub breach is a <strong>significant breach</strong>. If any of the tokens of any of the common base images had been compromised — packages like Alpine, busybox, Node.JS, or any of the major databases — these could have easily permeated into the wild with little or no knowledge. From my point of view, the only way to be sure you’re not affected is to somehow verify with the image provider that their account has been cleaned up and redeploy all your containers. Cleaned up as <a href="https://success.docker.com/article/docker-hub-user-notification">directed by Docker</a> (note the URL, “success”). Why? Because it’s likely some <a href="https://chrisshort.net/upstream-vs-downstream/">upstream</a> used Docker Hub even if you didn’t. In other words, “Nuke the entire site from orbit. It’s the only way to be sure.” Yes, it’s that bad until it’s confirmed otherwise.

<em>What a Docker freaking mess we’re in</em>. At the risk of being extra petty, I can’t help but mention I started using <a href="https://quay.io/">Quay</a> when I joined Red Hat and I’m pretty happy with it. It’s a container registry, not an Alexander Wang piece. Quay is not perfect but, I’m not expecting a whole lot here. It looks like I’ll be moving more images off Docker Hub in the future.

<a href="https://www.indeedprime.com/devopsish/?sid=us_other-EmailSponsor_JS_ACQ&amp;kw=Devopsish_Email1"><strong>Hit send on your last tech job application</strong></a><br/>One application, many tech opportunities. Indeed Prime makes job hunting quick and easy. Save time and let us do the heavy lifting for you by matching you with top tech companies. <a href="https://www.indeedprime.com/devopsish/?sid=us_other-EmailSponsor_JS_ACQ&amp;kw=Devopsish_Email1">Join for free today</a>! <em>SPONSORED</em>

<a href="https://logdna.com/sign-up/?utm_medium=Syndication&amp;utm_campaign=DevOpsish&amp;utm_source=DevOpsish"><strong>Log Management Modernized</strong></a><br/>With LogDNA’s fast, multi-cloud logging platform, DevOps and Engineering teams can easily and quickly aggregate all system and application logs into one efficient platform.<br/>Whether on-premise, in the cloud, or a hybrid solution, we have you covered. Don’t take our word for it. Try it yourself.

<a href="https://logdna.com/sign-up/?utm_medium=Syndication&amp;utm_campaign=DevOpsish&amp;utm_source=DevOpsish">Get started logging in a few minutes with a free trial</a>. <em>SPONSORED</em>

### DevOps’ish Top Five from Last Week

### People

1. [Google Walkout Organizers Say They’re Facing Retaliationdispute intensified#NotOkGoogle](https://www.wired.com/story/google-walkout-organizers-say-theyre-facing-retaliation/)

     ”Two employee activists at Google say they have been retaliated against for helping to organize a walkout among thousands of Google workers in November.” I’m pretty sure this is illegal, no? The  at a town hall on Friday.
1. [Hire People or Optimize Processes: A cost-benefit analysis for engineering leaders](https://codeclimate.com/blog/scale-engineering-calculator/)

     Help to figure out if you should hire more people or optimize your widget making.
1. [Kubernetes jobs hunt: How to land that role](https://enterprisersproject.com/article/2019/4/kubernetes-jobs-hunt-how-land-role)

     Trying to get a job working with Kubernetes? Consider these five tips.
1. [The Difference Between Goals, Strategies, Metrics, OKRs, KPIs, and KRIs](https://danielmiessler.com/blog/the-difference-between-goals-strategies-metrics-okrs-kpis-and-kris/)

     The differences and similarities between the most common types of business measurement systems
1. [DNS over HTTPS is coming whether ISPs and governments like it or not](https://nakedsecurity.sophos.com/2019/04/24/dns-over-https-is-coming-whether-isps-and-governments-like-it-or-not/)

     Encrypting your DNS queries in the payload of an HTTPS packet means that countries and companies can’t as easily hijack DNS to control internet access or to monitor employee activity. Conversely, blocking known bad DNS entries (sinkhole) and using DNS query logs to hunt for indicators of compromise are common security measures. This becomes a MUCH harder problem.
1. [The Science Behind DevOps with Dr. Nicole Forsgren — Real World DevOps](https://share.transistor.fm/s/433ecde0)

    
### Process

1. [Accelerate: State of DevOps 2019 Surveythe 2018 reportAccelerate: The Science of Lean Software and DevOps: Building and Scaling High Performing Technology Organizations](https://google.qualtrics.com/jfe/form/SV_0v2VZMeA2Eha365?sp=5)

     Nicole Forsgren, PhD is conducting the State of DevOps 2019 Survey. Your input is incredibly important. On several occasions, I have referenced  since its release for real-world work that impacts real numbers. Nicole’s group also wrote, , which I cannot recommend enough either.
1. [affiliate programs](../terms/)

    Note: DevOps’ish may earn compensation for sales from links on this post through affiliate programs.
1. [How You Can Help Localize Kubernetes Docs](https://kubernetes.io/blog/2019/04/26/how-you-can-help-localize-kubernetes-docs/)

     If you’re interested in helping an existing Kubernetes localization please help us out!
1. [7 advantages of open source for agile teams](https://enterprisersproject.com/article/2019/4/7-advantages-open-source-agile)

     “Adopting DevOps and agile practices is how companies are managing the velocity of change. I imagine it’s possible to do agile or DevOps in a closed source world. The question then becomes, why put that hurdle there?”
1. [Open Source Software: The Complete Wired Guide](https://www.wired.com/story/wired-guide-open-source-software/)

     Some knowledge about open source software that’s easily consumable by the unknowing.
1. [Forge Your Future with Open Source: Build Your Skills. Build Your Network. Build the Future of Technology. by VM (Vicky) Brasseur](https://pragprog.com/book/vbopens/forge-your-future-with-open-source)

     And if you want to get going in Open Source, I highly recommend Vicky’s book.
1. [Slack files to go public, reports $138.9M in losses on revenue of $400.6M](https://techcrunch.com/2019/04/26/slack-files-to-go-public/)

     And the tech IPOs just keep on coming.
1. [Some internet outages predicted for the coming month as ‘768k Day’ approachesBGP and DNS are the bubble gum and duct tape of the internet](https://www.zdnet.com/article/some-internet-outages-predicted-for-the-coming-month-as-768k-day-approaches/)

     Some routers will crash soon due to ternary content-addressable memory needs of over 768K. .
1. [Apple spends more than $30 million a month on Amazon Web ServicesApple Slashed Amazon Cloud Spending 50 Percent in Bid for Self-Sufficiency](https://www.cnbc.com/2019/04/22/apple-spends-more-than-30-million-on-amazon-web-services-a-month.html)

     People freaked out when they heard this number. Apple uses a lot of different cloud providers for various reasons (I used to work for one). But, I think the bigger story is,
1. [A Roadmap to Convergence – OpenTracingmerge the OpenTracing and OpenCensus projects](https://medium.com/opentracing/a-roadmap-to-convergence-b074e5815289)

     “We are creating a new, unified set of libraries and specifications for observability telemetry. It will , and provide a supported migration path.”
1. [The DevOps Institute Has Been BrandjackedQuickStart](https://devops.com/the-devops-institute-has-been-brandjacked/)

    This is some pretty gross behavior by  but, trademark and copyright law exist for a reason. I certainly hope QuickStart Learning Inc. finds out about the law very soon.
1. [Accenture sued over website redesign so bad it Hertz: Car hire biz demands $32m+ for ‘defective’ cyber-revampThis Twitter thread](https://www.theregister.co.uk/2019/04/23/hertz_accenture_lawsuit/)

     This is so bad y’all.  is also gold.
1. [DevOps’ish with Chris Short – Newsletterers – The Tim Show – S02E01](https://devopsi.sh/timshow)

     Ever wonder how the DevOps’ish newsletter started, how I build it, or what it’s like to write a newsletter? Check out fellow Red Hatter Tim Hildred’s podcast on that very topic!
### Tools

1. [Running Drupal in Kubernetes with Docker in production](https://www.jeffgeerling.com/blog/2019/running-drupal-kubernetes-docker-production)

     ”You really have to be on your game in the world of containerized-Drupal-in-production!”
1. [Hardware Accelerated SSL/TLS Termination in Ingress Controllers using Kubernetes Device Plugins and RuntimeClass](https://kubernetes.io/blog/2019/04/24/hardware-accelerated-ssl-tls-termination-in-ingress-controllers-using-kubernetes-device-plugins-and-runtimeclass/)

    
1. [Packets-per-second limits in EC2](https://stressgrid.com/blog/pps_limits_in_ec2/)

     ”[W]e determined that each EC2 instance type has a packet-per-second budget. Surprisingly, this budget goes toward the total of incoming and outgoing packets. Even more surprisingly, the same budget gets split between multiple network interfaces, with some additional performance penalty.”
1. [How to run systemd in a container](https://developers.redhat.com/blog/2019/04/24/how-to-run-systemd-in-a-container/)

     Trying to run systemd in a container is hard af. I’ve had to do it to test deployments before and I have to say, Dan’s method is probably better than any I’ve ever seen. I want to change the Podman name so bad though, Man.
1. [Kubernetes Tutorial - Step by Step Guide to Basic Kubernetes Concepts](https://auth0.com/blog/kubernetes-tutorial-step-by-step-introduction-to-basic-concepts/)

     This is a nice introduction to k8s from Auth0. Good for them for getting a little out of their comfort zone to help lower the barrier to entry a little.
1. [Leveraging Jenkins on Kubernetes](https://caylent.com/leveraging-jenkins-on-kubernetes/)

    
1. [Kubernetes Network Policy APIs](https://octetz.com/posts/k8s-network-policy-apis)

     “This post explores multiple ways network policy can be expressed in Kubernetes.” In other words, way more than I want to know about Kubernetes networking. But, good to have as a reference.
1. [Improving the security of Kubernetes clusters using Istio](https://blog.giantswarm.io/Improving-security-with-Istio/)

     Istio does SO MUCH.
1. [Istio the Easy Way](https://medium.com/solo-io/istio-the-easy-way-de66e6eba4a1)

     Oh good there’s an easy way! Thank you, Christian Posta and solo.io for this.
1. [Python Project Tooling explained](https://write.as/chobeat/python-project-tooling-explained)

     Instant bookmark. Please share with others that are new or even a little old to Python.
1. [All That You Need to Know About Microsoft’s New Programming Language: Bosque](https://dev.to/0xrumple/all-what-you-need-to-know-about-microsoft-s-new-programming-language-bosque-38c0)

     ”The Bosque programming language is a Microsoft Research project that is investigating language designs for writing code that is simple, obvious, and easy to reason about for both humans and machines.”
1. [Why Go? – Key advantages you may have overlookedGophergopher](https://yourbasic.org/golang/advantages-over-java-python/)

     I hear good things about these  people. Not those  people.
1. [Open Sourcing Jingo, a Faster JSON Encoder for Go](https://bet365techblog.com/open-sourcing-jingo-a-faster-json-encoder-for-go)

     This package provides the ability to encode golang structs to a buffer as JSON very quickly.
1. [Building platform stacks with in-house scripts vs. Kubernetes Operators](https://medium.com/@cloudark/building-platform-stacks-with-in-house-scripts-vs-kubernetes-operators-95541e555025)

     STOP HAND ROLLING SCRIPTS! YOUR ARTISANAL SCRIPTS DON’T BELONG! GET OFF MY LAWN!
1. [ricardbejarano/haproxy](https://github.com/ricardbejarano/haproxy)

     🏎 Built-from-source container image of the HAProxy load balancer
1. [cdr/sshcode](https://github.com/cdr/sshcode)

     Run VS Code on any server over SSH
1. [bxcodec/gotcha](https://github.com/bxcodec/gotcha)

     gotcha: inmemory-cache in Go (Golang) with customizable algorithm
1. [mhausenblas/kboom](https://github.com/mhausenblas/kboom)

     The Kubernetes scale & soak load tester (do you want to shred some clusters with me?!?)
1. [GoogleCloudPlatform/berglas](https://github.com/GoogleCloudPlatform/berglas)

     A tool for managing secrets on Google Cloud (I hope Google does End of Life this before the newsletter goes live)
1. [operator-framework/community-operatorsOKD](https://github.com/operator-framework/community-operators)

     The canonical source for Kubernetes Operators that appear on OperatorHub.io, OpenShift Container Platform, and

### [ << Prev ](sreweekly-124.md) ------------- [ Next >> ](sreweekly-126.md)