## [DevOps'ish 134](https://devopsish.com/134) - Sun Jun 30, 2019

I spent a lot of time this week struggling with an odd problem. How do you make a multicloud platform without having to do a ton of work? I was trying to figure out why it’s so damn hard to get a multicloud platform going. People want this for various reasons but, struggle with it. Why? What was I trying to do? Host the DevOps’ish web site where it was less likely to be blocked based on a visitor’s geography.

The idea was if someone were in Hong Kong they’d hit a bucket in Alibaba Cloud. The US and Europe would likely draw from Google Cloud Storage. Everywhere else was going to fall to the algorithms. Turns out Route53 doesn’t let you route traffic based on geography to anything outside of AWS. Same for Google’s Cloud DNS. Digital Ocean spaces aren’t quite ready for prime time yet either.

Folks have to pick a cloud or build a platform across them. This is why Kubernetes is such a big deal. This is why Google’s dominance in the Kubernetes space matters so much. No one is going to be able to work a cloud providers primitives into the ANYCAST hybrid multicloud of their dreams. But if something is built on top of the primitives magic might be possible. For now, though, I’m using Google Cloud CDN. At least it’s not blocked in most places. Turns out there are some handy sites for testing a domain’s accessibility in other nations:

<a href="https://turbonomic.com/state-of-multicloud/?utm_campaign=7012o000001oRz6AAE"><strong>2019 State of Multicloud</strong></a><br/>A Report on the Underlying Dynamics Fueling Multicloud Strategies. <a href="https://turbonomic.com/state-of-multicloud/?utm_campaign=7012o000001oRz6AAE">Download Today!</a> <em>SPONSORED</em>

### DevOps’ish Last Week’s Top Five

### People

1. [Bill Gates on making “one of the greatest mistakes of all time”](https://techcrunch.com/2019/06/22/bill-gates-on-making-one-of-the-greatest-mistakes-of-all-time/)

     “You know, in the software world, in particular for platforms, these are winner-take-all markets. So, you know, the greatest mistake ever is the whatever mismanagement I engaged in that caused Microsoft not to be what Android is, [meaning] Android is the standard non-Apple phone form platform. That was a natural thing for Microsoft to win.”
1. [Without Any Warrants, CBP Searched My Devices at the Airport](https://theintercept.com/2019/06/22/cbp-border-searches-journalists/)

     Pretty disturbing country I’m living in these days.
1. [How I use Slack—alone—to get more done](https://char.gd/blog/2019/how-i-use-slack-alone-to-get-more-done)

     How Owen Williams gets less email with Slack. The Asana, Twitter, Google Calendar, and Google Drive Slack integrations are crucial to my work productivity. I am considering a similar setup to Owen because it would save me a ton of time.
1. [I spend too much time in Zoom…](https://medium.com/@joebeda/i-spend-too-much-time-in-zoom-5eedcea5cc90)

     Joe Beda takes nerd to the next level.
1. [I care about Ansible](https://jpmens.net/2019/06/21/i-care-about-ansible/)

     Debatable. In my opinion, if someone cares about Ansible they wouldn’t berate Ansible team members publicly. Nor would they heckle them at conferences, ever. That’s bare minimum civility, in my opinion. Any project matters WAY less than the people that keep it afloat. All the pull requests in the world don’t matter if no one is there to merge them.
1. [He Cyberstalked Teen Girls for Years—Then They Fought Back](https://www.wired.com/story/cyberstalked-teen-girls-for-years-fought-back/)

     This is some really vile stuff. Eight years in prison doesn’t seem like enough.
1. [Command Line Heroes: Season 3: Python’s Tale](https://www.redhat.com/en/command-line-heroes/season-3/pythons-tale)

     ”In this episode, Emily Morehouse makes the connection between Python’s technical extensibility and its inclusive community. Michael Kennedy explains how Python is both easy to learn and powerful enough to build YouTube and Instagram. And Diane Mueller highlights how the Python community took the lead on so many inclusive practices that are spreading in tech—including the rise of community-led decision-making.”
### Process

1. [New Kubernetes command-line flaw discovereddiscovered by Charles Holmes](https://techerati.com/news-hub/new-kubernetes-flaw-discovered-command-line/)

     ”Latest kubectl vulnerability linked to incomplete patch of previous flaw discovered in March.” Oops… This was  as part of a CNCF-sponsored security audit.
1. [Joint Chiefs CIO on JEDI: ‘Our warfighters need this capability now’acquisition process as convoluted as they do](https://www.fedscoop.com/joint-chiefs-cio-jedi-delay-letter/)

     Oracle is essentially tying the hands of US warfighters at this point. The DOD has done this to themselves though for having an . But, I question whether Oracle really has a leg to stand on here technologically.
1. [BGP super-blunder: How Verizon today sparked a ‘cascading catastrophic failure’ that knackered Cloudflare, Amazon, etc](https://www.theregister.co.uk/2019/06/24/verizon_bgp_misconfiguration_cloudflare/)

     ”Verizon sent a big chunk of the internet down a black hole… after it wrongly accepted a network misconfiguration from a small ISP in Pennsylvania, USA.” Verizon needs to be broken up. They’re so big they can’t manage their networks now. One of the worst outages I ever experienced was at Bankrate. Our site was down hard most of a long Independence Day weekend. Why? A faulty node someone in Verizon’s network that Fastly ended up identifying and routing around. I bought Fastly employees drinks whenever I saw them that year.
1. [How Verizon and a BGP Optimizer Knocked Large Parts of the Internet Offline Today](https://blog.cloudflare.com/how-verizon-and-a-bgp-optimizer-knocked-large-parts-of-the-internet-offline-today/)

     “This should never have happened because Verizon should never have forwarded those routes to the rest of the Internet. To understand why, read on.”
1. [What 36,000 OSS Projects and 12,000 Commercial Dev Teams Taught Us About Secure Coding Practices](https://blog.sonatype.com/2019ssc)

     Sonatype’s 2019 State of the Software Supply Chain report is now available.
1. [Kubernetes Security Myths Debunked](https://www.tremolosecurity.com/kubernetes-security-myths-debunked/)

     Some really good tips for making containers and Kubernetes security.
1. [How SRE teams are organized, and how to get started](https://cloud.google.com/blog/products/devops-sre/how-sre-teams-are-organized-and-how-to-get-started)

     It’s not a Deloitte or Accenture ratio of lines of code to the number of people. SRE is a process that when considered will help figure out what it’s going to take to get to the desired outcome. Now, with that being said, calculus is hard.
1. [CVE-less vulnerabilities](https://lwn.net/Articles/791855/)

     ”The sheer number of bugs being reported is overwhelming many (most?) free-software projects, which simply do not have enough eyeballs to fix, or even triage, many of the reports they receive. A discussion about that is currently playing out on the oss-security mailing list.”
1. [Pillaging The Jenkins Treasure Chest](https://dolosgroup.io/blog/2019/6/20/pillaging-the-jenkins-treasure-chest)

     “Jenkins tends to be a treasure trove of information in certain organizations, and it’s all too easy for a developer or operations team to leave something behind ‘just to get things done’.”
1. [What is DevOps](https://devopsish.com/what-is-devops/)

     “DevOps is the professional practice of frequent, continued, and iterative improvements through measurable changes, the goal of which is to become a high-velocity organization thus improving business outcomes.”
1. [European Commission - PRESS RELEASES - Press release - Mergers: Commission approves acquisition of Red Hat by IBM](http://europa.eu/rapid/press-release_IP-19-3433_en.htm)

     Only a handful of approvals left. I’ll still be a Red Hat employee when the deal closes so 😛😛😛.
### Tools

1. [Introducing Waxosuit](https://medium.com/@KevinHoffman/introducing-waxosuit-6ad754b48ed9)

     ”I didn’t want my development teams to have to re-write, copy/paste, or cargo cult stale implementations of NFRs for things like logging, tracing, contextual tracing (e.g. OpenTracing/Jaeger), Application Performance Monitoring (APM), health checking (e.g. live/readiness probes), message broker client wrappers, key-value store client wrappers, HTTP server endpoint wrappers, and so on.”
1. [A deep dive into Linux namespaces](http://ifeanyi.co/posts/linux-namespaces-part-1/)

     ”In this series of posts we will look closely at one of the main ingredients in a container - Namespaces”
1. [Raspberry Pi 4 on sale now from $35](https://www.raspberrypi.org/blog/raspberry-pi-4-on-sale-now-from-35/)

     A $55 (USD) version is available with 4 GB of RAM. This is a full computer now for a lot of folks. I can’t believe how far this platform has come.
1. [The Raspberry Pi 4 launch site runs on a Pi 4 cluster](https://arstechnica.com/information-technology/2019/06/the-raspberry-pi-4-launch-site-runs-on-a-pi-4-cluster/)

     ”The launch site for the Raspberry Pi 4 Model B is mostly running on a cluster of 18 of the little devices themselves. Fourteen handle PHP code execution, two serve static files, and two run memcached.” But, they’re using Cloudflare and an external database. I mean, I/O would be a nightmare.
1. [New property of light discovered](https://phys.org/news/2019-06-property.html)

     ”researchers have found that light can also be twisted, a property called angular momentum. Beams with highly structured angular momentum are said to have orbital angular momentum (OAM), and are called vortex beams.” This could potentially be an evolutionary step in communications.
1. [The in and outs of Microsoft’s new Windows Terminal](https://www.theregister.co.uk/2019/06/25/microsofts_new_terminal_put_through_paces/)

     Handy features in the Store applications, but the underlying infrastructure changes matter more
1. [After 4 years with nginx, we switched to Caddy - Here is why](https://engineering.hashnode.com/after-4-years-with-nginx-we-switched-to-caddy-here-is-why-cjxbv8eb2001ke8s1yl7ndroz)

    
1. [Security Crawl Maze: An Open Source Tool to Test Web Security Crawlers](https://opensource.googleblog.com/2019/06/security-crawl-maze-open-source-tool-to.html)

     ”Security Crawl Maze is a simple Python application built with the Flask framework that contains a wide variety of cases for ways in which a web based application can link other resources on the Web.”
1. [2001: Linux is cancer, says Microsoft. 2019: Hey friends, ah, can we join the official linux-distros mailing list, plz?](https://www.theregister.co.uk/2019/06/27/microsoft_linux_distro_list/)

     I know for a fact Microsoft was a little late to the party on a major Kubernetes vulnerability in the past year because they didn’t have solid representation on this list. This is a good thing for everyone in the end.
1. [Proposal: leave “if err != nil” alone? · Issue #32825 · golang/go](https://github.com/golang/go/issues/32825)

    
1. [chrisshort/homebrewLong story](https://quay.io/repository/chrisshort/homebrew)

     Yep. The Linux version of the venerable homebrew for Mac in a ubi8 container. Why? Well… . It’s not perfect but it’ll do.
1. [google/ko](https://github.com/google/ko)

     Build and deploy Go applications on Kubernetes
1. [alexrs/mobile-devops](https://github.com/alexrs/mobile-devops)

     A curated list of resources for mobile devops, including blog posts and talks about tooling and infrastructure for mobile apps

### [ << Prev ](sreweekly-133.md) ------------- [ Next >> ](sreweekly-135.md)