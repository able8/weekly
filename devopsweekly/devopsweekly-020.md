## [DevOps'ish 020](https://devopsish.com/020) - Sun Apr 23, 2017

Not quite sure what Docker is doing? Few people are. Docker is still Docker except when it’s <a href="https://mobyproject.org/">Moby</a>. Moby is open source and Docker isn’t (kind of). According to the <a href="https://blog.docker.com/2017/04/introducing-the-moby-project/">official Moby Project announcement</a>, Moby Project is <strong>NEW</strong>! But, Moby Project is actually the new <a href="https://chrisshort.net/upstream-vs-downstream/">upstream</a> for the Docker project. Interestingly enough, the Moby Project exists to create “An open framework to assemble specialized container systems without reinventing the wheel.” (Emphasis added.) The best analysis of the whole S.N.A.F.U. <a href="https://www.theregister.co.uk/2017/04/21/docker_renames_open_source_code_moby/">fittingly comes from The Register</a>.

I used <a href="https://en.wikipedia.org/wiki/Comic_Sans">Comic Sans</a> in the above image because this is a PR fustercluck. The worst part of the announcement was that it was timed REALLY poorly (or not at all). I learned of it when I click a GitHub issue for Docker and it took me to <a href="https://github.com/moby/moby/issues/4717">github.com/moby/moby/issues/4717</a>. I legitimately thought something was horribly wrong with GitHub. Nope! All the branding shenanigans appeared to happen all at once during DockerCon.

Containers are hard to explain to people because <a href="https://blog.jessfraz.com/post/containers-zones-jails-vms/">containers aren’t really a thing</a>. Linux containers are a combination of pre-existing Linux kernel components. The Founder of Docker, Solomon Hykes, <a href="https://twitter.com/solomonstre/status/855918630915133440">tweeted out two sketches to explain the Moby Project</a>. The fact that this was necessary shows how awkward this is to explain. Now, the container waters are even muddier than they already were by this rough re-branding effort.

<a href="https://blog.bluematador.com/posts/mini-guide-google-golang-why-its-perfect-for-devops/">Mini-Guide to Google’s Golang and Why It’s Perfect for DevOps</a>

<a href="https://ma.ttias.be/mit-no-longer-owns-18-0-0-08/">MIT no longer owns 18.0.0.0/8</a>: MIT is <a href="https://gist.github.com/simonster/e22e50cd52b7dffcf5a4db2b8ea4cce0">selling off it’s IPv4 block to fund its transition to IPv6</a>.

<a href="https://yalantis.com/blog/whats-wrong-telegram-open-api/">What’s Wrong With Open Source Telegram?</a> Apparently a lot!

<a href="https://blog.docker.com/2017/04/introducing-linuxkit-container-os-toolkit/">Announcing LinuxKit: A Toolkit for Building Secure, Lean AND Portable Linux Subsystems</a> (<a href="https://github.com/linuxkit/linuxkit">GitHub</a>)

<a href="https://github.com/digitalrebar/provision">Digitial Rebar Provision</a> looks to be an awesome replacement to the venerable Cobbler.

Here is a handy reference for enterprise Linux distros’ life cycles: <a href="https://linuxlifecycle.com/?utm_source=cronweekly.com">linuxlifecycle.com</a>

<a href="https://linux.slashdot.org/story/17/04/21/1623233/red-hat-suffers-massive-data-center-network-outage">Red Hat Suffers Massive Data Center Network Outage</a>

Friend of DevOps’ish <a href="https://ma.ttias.be/dns-spy-launched/">Mattias Geniar has officially launched DNS Spy</a>! In case you don’t recall, <a href="http://readwrite.com/2013/08/27/new-york-times-twitter-huffpo-hacked-syrian-electronic-army/">back in 2013, several news and media websites were hacked by exploiting their domain name registrars</a>. <a href="https://dnsspy.io/"><strong>DNS Spy</strong></a> will alert you almost immediately to any changes to your DNS records. This could indicate something is afoot or, at the very least, misconfigured.

<a href="https://lwn.net/SubscriberLink/720215/867f710e3d145b55/">LWN.net has a fantastic analysis of a talk at KubeCon 2017 that covers security in Kubernetes</a>.

According to Akamai, <a href="http://www.darkreading.com/attacks-breaches/new-breed-of-ddos-attack-on-the-rise/d/d-id/1328620">Connectionless LDAP is being exploited in a new breed of DDoS attack</a>.

<a href="https://containerhardening.org/">Linux Container Hardening</a>: “This project will focus on hardening of Linux containers. It will help contribute patches to the <a href="https://kernsec.org/wiki/index.php/Kernel_Self_Protection_Project">Kernel Self Protection Project</a> that evolve the primitives in the Linux kernel used by containers (namespaces, cgroups, etc) to be more secure.”

<a href="https://aws.amazon.com/about-aws/whats-new/2017/04/aws-lambda-supports-python-3-6/">AWS Lambda Supports Python 3.6</a>

<a href="https://aws.amazon.com/codestar/">AWS CodeStar</a> is a <a href="https://aws.amazon.com/blogs/aws/new-aws-codestar/">cloud service designed to make it easier to develop, build, and deploy applications on AWS by simplifying the setup of your entire development project</a>.

<img src="https://d33wubrfki0l68.cloudfront.net/6678bdeea7ddc4f9cd13166a86f27fbaed842b0e/db30e/020/run-k8s-tshirt.jpg" alt="Get this awesome shirt at [runk8s.io](https://runk8s.io/)"/><br/><em>Get this awesome shirt at <a href="https://runk8s.io/">runk8s.io</a></em>

<a href="https://insights.ubuntu.com/2017/04/05/ubuntu-on-aws-gets-serious-performance-boost-with-aws-tuned-kernel/">Ubuntu on AWS gets serious performance boost with AWS-tuned kernel</a>

<a href="https://www.washingtonpost.com/news/business/wp/2017/04/18/the-newest-silicon-valley-perk-paid-time-off-to-protest-trump/?utm_term=.3d2e13c39de5">The newest Silicon Valley perk? Paid time off to protest Trump.</a>

A great episode of a great podcast with friend, Ashley McNamara, talking about her new job, being a woman and mom in the world today, and how to become a coder. <a href="https://overcast.fm/+GnKbwnxKI">Go Time 43: Getting Better, Mentoring, and Drawing Gophers with Ashley McNamara</a>

This <a href="https://blogs.msdn.microsoft.com/oldnewthing/20170418-00/?p=95985">Microsoft Flight Simulator support case</a> reminds me of a call I answered when I worked at a dial-up ISP. The customer called in late at night to figure out why he couldn’t get online. It was late but, I was in the office because of a bad thunderstorm potentially affecting phone lines, modems, etc. I went through several troubleshooting steps only to find out that there was nothing being displayed on his monitor. When I asked if his monitor was turned on he informed me the power was out due to the storm. I was dumbfounded that this seemingly intelligent human did not realize his computer wasn’t powered by the phone line. He asked, “If I can call you, why can’t I get online?” I can’t…

Putting the period on the end of the Docker/Moby change:

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Mini-Guide to Google’s Golang and Why It’s Perfect for DevOps](https://blog.bluematador.com/posts/mini-guide-google-golang-why-its-perfect-for-devops/)

    
1. [MIT no longer owns 18.0.0.0/8selling off it’s IPv4 block to fund its transition to IPv6](https://ma.ttias.be/mit-no-longer-owns-18-0-0-08/)

     MIT is .
1. [What’s Wrong With Open Source Telegram?](https://yalantis.com/blog/whats-wrong-telegram-open-api/)

    Apparently a lot!
### Department of Refreshment and Refurbishment

1. [Announcing LinuxKit: A Toolkit for Building Secure, Lean AND Portable Linux SubsystemsGitHub](https://blog.docker.com/2017/04/introducing-linuxkit-container-os-toolkit/)

    ()
1. [Digitial Rebar Provision](https://github.com/digitalrebar/provision)

    looks to be an awesome replacement to the venerable Cobbler.
1. [linuxlifecycle.com](https://linuxlifecycle.com/?utm_source=cronweekly.com)

    Here is a handy reference for enterprise Linux distros’ life cycles:
### Department of Data Defense

1. [Red Hat Suffers Massive Data Center Network Outage](https://linux.slashdot.org/story/17/04/21/1623233/red-hat-suffers-massive-data-center-network-outage)

    
1. [Mattias Geniar has officially launched DNS Spyback in 2013, several news and media websites were hacked by exploiting their domain name registrarsDNS Spy](https://ma.ttias.be/dns-spy-launched/)

    Friend of DevOps’ish ! In case you don’t recall, .  will alert you almost immediately to any changes to your DNS records. This could indicate something is afoot or, at the very least, misconfigured.
1. [LWN.net has a fantastic analysis of a talk at KubeCon 2017 that covers security in Kubernetes](https://lwn.net/SubscriberLink/720215/867f710e3d145b55/)

    .
1. [Connectionless LDAP is being exploited in a new breed of DDoS attack](http://www.darkreading.com/attacks-breaches/new-breed-of-ddos-attack-on-the-rise/d/d-id/1328620)

    According to Akamai, .
1. [Linux Container HardeningKernel Self Protection Project](https://containerhardening.org/)

     “This project will focus on hardening of Linux containers. It will help contribute patches to the  that evolve the primitives in the Linux kernel used by containers (namespaces, cgroups, etc) to be more secure.”
### Department of Happy Little Clouds

1. [AWS Lambda Supports Python 3.6](https://aws.amazon.com/about-aws/whats-new/2017/04/aws-lambda-supports-python-3-6/)

    
1. [AWS CodeStarcloud service designed to make it easier to develop, build, and deploy applications on AWS by simplifying the setup of your entire development project](https://aws.amazon.com/codestar/)

    is a .
1. [runk8s.io](https://runk8s.io/)

    Get this awesome shirt at runk8s.io
1. [Ubuntu on AWS gets serious performance boost with AWS-tuned kernel](https://insights.ubuntu.com/2017/04/05/ubuntu-on-aws-gets-serious-performance-boost-with-aws-tuned-kernel/)

    
### Department of Woke Workplaces

1. [The newest Silicon Valley perk? Paid time off to protest Trump.](https://www.washingtonpost.com/news/business/wp/2017/04/18/the-newest-silicon-valley-perk-paid-time-off-to-protest-trump/?utm_term=.3d2e13c39de5)

    
1. [Go Time 43: Getting Better, Mentoring, and Drawing Gophers with Ashley McNamara](https://overcast.fm/+GnKbwnxKI)

    A great episode of a great podcast with friend, Ashley McNamara, talking about her new job, being a woman and mom in the world today, and how to become a coder.
### Department of Story Time

1. [Microsoft Flight Simulator support case](https://blogs.msdn.microsoft.com/oldnewthing/20170418-00/?p=95985)

    This  reminds me of a call I answered when I worked at a dial-up ISP. The customer called in late at night to figure out why he couldn’t get online. It was late but, I was in the office because of a bad thunderstorm potentially affecting phone lines, modems, etc. I went through several troubleshooting steps only to find out that there was nothing being displayed on his monitor. When I asked if his monitor was turned on he informed me the power was out due to the storm. I was dumbfounded that this seemingly intelligent human did not realize his computer wasn’t powered by the phone line. He asked, “If I can call you, why can’t I get online?” I can’t…

### [ << Prev ](sreweekly-19.md) ------------- [ Next >> ](sreweekly-21.md)