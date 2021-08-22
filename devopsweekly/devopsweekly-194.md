## [DevOps'ish 194](https://devopsish.com/194) - Sun Nov 29, 2020

There is usually a lot of hype surrounding Apple announcements. The recent report of Apple starting to build Macs with their own ARM-based silicon is no exception. But, there’s some meat to this hype; let me explain.

It wasn’t long ago that my iPhone 8 Plus with its <a href="https://www.macrumors.com/2017/09/13/a11-bionic-chip-geekbench-scores/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">A11 Bionic chip</a> could leave my MacBook Air I was using as a daily driver in the dust. The ARM-based phone you’re carrying around (Apple or otherwise) <a href="https://theconversation.com/would-your-mobile-phone-be-powerful-enough-to-get-you-to-the-moon-115933?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">probably has more computing power than the entire Apollo space program</a>.

In 2017, I was linked up with Edward Vielmetti of <a href="https://www.worksonarm.com/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">Works on ARM</a> fame. I was already tinkering and building Kubernetes clusters with Raspberry Pis. Ed validated my thinking and while it’s going to be quite some time before we’re all running ARM chips, that day is coming. I saw ARM as a fine alternative thanks to modern languages multiarch compilers (may The Maker bless Go). But there was always a cross-compilation tax for users. But, it was a one-time task if you were smart about things.

To further prove ARM’s market penetration, <a href="https://devopsish.com/042/">Docker Hub supported “multi-platform” images</a> in 2017. <a href="https://blog.cloudflare.com/arm-takes-wing/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">Cloudflare was also benchmarking the ARM chips</a> they had their hands on in 2017 against their Intel fleet. Red Hat announced an <a href="https://www.redhat.com/en/blog/red-hat-introduces-arm-server-support-red-hat-enterprise-linux?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">ARM-based distro of RHEL back in 2017</a> (Note: Red Hat is my employer). In June 2020, an <a href="https://www.zdnet.com/article/arm-and-linux-take-supercomputer-top500-crown/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">ARM-powered supercomupter became the fastest supercomputer in the world</a>.

Yesterday, I was reading the news like I normally do (in my favorite newsreader, <a href="https://www.inoreader.com/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">Inoreader</a>). I kept seeing <a href="https://steipete.com/posts/apple-silicon-m1-a-developer-perspective/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">reports about the new M1 ARM Macs</a> making their users happy. I came across one article showing a very bright future for <a href="https://www.tomsguide.com/news/macbook-pro-16-inch-m1x-chip-just-leaked-and-its-game-over-for-intel?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">the next generation of ARM-based Macs (the M1X chip)</a>. It’s supposedly a 12 core monster that will arrive in the 16” MacBook Pro next year. All this Apple talk is the exact opposite of what I usually discuss. But, it shows the progression of ARM over a breathtaking pace. Which led me to tweet a brief synopsis of the recent history of ARM…

Yes, the servers came before the supercomputers, obviously, but I was in the middle of like three things at once. Needless to say, ARM is here. We’re talking about it because Apple has the design and fabrication capacity to make it happen. But, there will be many opportunities for the ole Intel and AMD allies to make the switch to offering yet another processor architecture. Dell, HP, Asus, Acer, etc., have to show they aren’t lagging behind somehow. Make no mistake, Apple has the luxury of comparing itself to itself for now and potentially for all time. But, I can guarantee Microsoft does not want to lose the ARM market completely. It’s why they have an <a href="https://docs.microsoft.com/en-us/windows/arm/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194">ARM version of Windows</a>. Hold on tight, folks. With more ways to crack an Intel chip open seemingly springing to life every month (math done by a person who is bad at it), change is afoot (or an ARM if you’re into dad jokes).

<strong>DevOps’ish is brought to you by</strong> <a href="https://www.accurics.com/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=194"><strong>Accurics</strong></a>

### People

1. []()

    No shortage of KubeCon Virtual 2020 recaps this week, let me know if I missed some:
1. [newsletter sponsor](https://devopsish.com/sponsor/)

    Want to see your ad in DevOps’ish? Check out the  page for all the details. SPONSORED
1. [‘Tokenized’: Inside Black Workers’ Struggles at Coinbase](https://www.nytimes.com/2020/11/27/technology/coinbase-cryptocurrency-black-employees.html)

    “[A]ccording to 23 current and former Coinbase employees, five of whom spoke on the record, as well as internal documents and recordings of conversations, the start-up has long struggled with its management of Black employees.” Some truly repugnant behavior from Coinbase.
1. [Tony Hsieh, founder of Zappos, online shoe retailer, dies at 46](https://www.washingtonpost.com/local/obituaries/tony-hsieh-dead/2020/11/28/9669a22a-3197-11eb-bae0-50bb17126614_story.html)

    “[Tony] Hsieh was reported injured in a house fire on Nov. 18 while visiting family members in New London, Conn. New London Fire Chief Thomas Curcio told the Day newspaper in Connecticut that firefighters rescued a victim trapped in a waterfront house. The victim, not specifically identified as Mr. Hsieh, was flown by helicopter to the burn unit of a hospital in Bridgeport, Conn.”
1. [Linus Torvalds would like to use an M1 Mac for Linux, but…](https://www.zdnet.com/article/linus-torvalds-would-like-to-use-an-m1-mac-for-linux-but/)

    Speaking on the new ARM-based Macs, Linus Torvalds said, “The main problem with the M1 for me is the GPU and other devices around it, because that’s likely what would hold me off using it because it wouldn’t have any Linux support unless Apple opens up.” He’s 100% right. I remember the days when only certain GPUs would work out of the box with Linux distros. ATI (now AMD) was friendly at the time. Apple needs to open up some code so Linux is a possibility. This also helps with those pesky anti-trust subpoenas potentially heading Apple’s way.
1. [Common Performance Review Biases: How to Spot and Counter Them](https://blog.pragmaticengineer.com/performance-review-biases/)

    “This article summarizes the 8 common biases I’ve observed in performance reviews and how you - as someone receiving this review - can counter them.”
1. [Building a healthy relationship between security and sysadmins](https://www.redhat.com/sysadmin/security-and-sysadmins)

    “Learn how to bridge the gap between operations/development and security.” It’s one of the most important relationships you can maintain in your organization.
### Process

1. [Download today: container security ebook - going beyond image scanningDownload this eBook](https://security.stackrox.com/container-security-going-beyond-image-scanning.html?Source=DevOpsish&LSource=DevOpsish)

    Vulnerability scanning is often the first step to securing containers and Kubernetes, but it’s insufficient by itself. , from StackRox, to learn about the security checks, controls, and best practices DevOps should consider to secure container images and CI build pipeline, Kubernetes deployments, running workloads, and the underlying Kubernetes infrastructure. SPONSORED
1. [Summary of the Amazon Kinesis Event in the Northern Virginia (US-EAST-1) Region](https://aws.amazon.com/message/11201/)

    When a mistake in your capacity increase process brings down half the internet… A great postmortem as usual from AWS. Now, it’d be nice if half the internet didn’t solely depend on us-east-1.
1. [Stripe Said to Be in Talks for Funding at $70 Billion Value](https://www.bloomberg.com/news/articles/2020-11-24/payments-startup-stripe-is-said-in-talks-to-raise-new-funding)

    “Valuation could double – or more – valuation in April round”
1. [Salesforce is in talks to buy Slack, deal could be announced next week](https://www.cnbc.com/2020/11/25/slack-shares-jump-following-report-of-possible-salesforce-acquisition.html)

    “Salesforce is in talks to acquire Slack, and a deal could be announced as soon as next week, according to a person familiar with the matter.” I’m not sure if this is a good or bad thing. Salesforce could royally screw this up when Slack’s popularity is waning. Now is as good a time as ever to sell and Salesforce is going to pay the premium Slack likely wants.
1. [Addressing The Developer Experience Gap. Newcode or Gluecode](https://redmonk.com/jgovernor/2020/11/26/addressing-the-developer-experience-gap-newcode-or-gluecode/)

    “So in conclusion, Newcode or Gluecode, there is a ton of interesting stuff going on to address The Developer Gap. Addressing this gap will define success in the industry going forward. Companies that makes their developers most successful are going to win in their markets.”
1. [An online place of recorddon’t use Mediumchrisshort.net](https://rubenerd.com/an-online-place-of-record/)

    Medium isn’t the place where people write anymore () which is good. But, what’s replacing it is another tech startup. Take control of your online brand.  is sixteen years old. The site or its contents have come up in every interview since 2017. It’s wild but, shows the power of owning your content.
1. [Someone attacked our company](https://usefathom.com/blog/ddos-attack)

    Someone is indeed spamming the ever loving shit out of Fathom Analytics.
### Tools

1. [OpenTelemetry Starter Kit at the Honeycomb booth](https://info.honeycomb.io/honeycomb-and-kubecon)

    🎺 Looking to up your hands-on skills at KubeCon? We are premiering our . Instrument a sample app or your own production data. There are prize-drawings for completed steps, but the real prize is launching with OTel and getting valuable insights right away.
1. [Honeycomb.io](https://ui.honeycomb.io/signup/?&utm_source=devopsish&utm_medium=newsletter&utm_campaign=ad&utm_content=product-signup)

    OpenTelemetry + observability = path to production excellence. Get the observability you deserve from . SPONSORED
1. [Understanding Kubernetes Operators](https://caylent.com/understanding-kubernetes-operators)

    “Kubernetes operator is one of the tools designed to push automation past its limits. You can do so much more without having to rely on manual inputs every time.”
1. [Cloud Native Computing Foundation Announces etcd Graduation](https://www.cncf.io/announcements/2020/11/24/cloud-native-computing-foundation-announces-etcd-graduation/)

    When the call for votes went out, I thought, “How hasn’t this graduated yet? Whoops.”
1. [Primer: How XDP and eBPF Speed Network Traffic via the Linux Kernel](https://thenewstack.io/primer-how-xdp-and-ebpf-speed-network-traffic-via-the-linux-kernel/)

    “XDP isn’t a language. Instead, XDP uses the eBPF programming language to gain access to the lower-level kernel hook. That hook is then implemented by the network device driver within the ingress traffic processing function, before a socket buffer can be allocated for the incoming packet.”
1. [The Pen Testing Tools We’re Thankful for in 2020](https://labs.bishopfox.com/industry-blog/pen-testing-tools-were-thankful-for-in-2020)

    A very nice list of security testing tooling.
1. [WireGuard For Windows Updated With Improved Installer, ARM/ARM64 Support](https://www.phoronix.com/scan.php?page=news_item&px=WireGuard-For-Windows-0.3)

    “WireGuard For Windows 0.3 is out as a big update bringing documentation improvements, ARM/ARM64 support so Windows users on the Raspberry Pi or relevant Microsoft Surface devices can enjoy this functionality, installer infrastructure improvements, improved configuration handling, support for split DNS tunneling, support for concurrently handling multiple tunnels, and performance/stability improvements.”
1. [Hackensack Board Member Who Opposed LGBTQ Curriculum Resigns After Embarrassing Zoom IncidentZoomsterbating](https://dailyvoice.com/new-jersey/hackensack/news/hackensack-board-member-who-opposed-lgbtq-curriculum-resigns-after-embarrassing-zoom-incident/798530/)

    Not quite , but close enough. “A short time later, Vice President Scott James-Vickery reportedly told [Frances] Cogelja: ‘You need to go. We’re here trying to get work done while you’re sitting on the toilet.’”
1. [awslabs/karpenter](https://github.com/awslabs/karpenter)

    Kubernetes Autoscaling: built for scalability and extensibility from the ground up.
1. [inovex/illuminatio](https://github.com/inovex/illuminatio)

    The kubernetes network policy validator.

### [ << Prev ](devopsweekly-193.md) ------------- [ Next >> ](devopsweekly-195.md)