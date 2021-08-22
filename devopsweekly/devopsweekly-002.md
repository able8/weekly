## [DevOps'ish 002](https://devopsish.com/002) - Sun Dec 18, 2016

It was the week before Christmas 🤶 🎅 and all through the <a href="https://devopsish.com/"><strong>DevOps</strong></a> world not a creature was stirring not even a mouse 🖱.

Hardly… You might be in a change freeze but the open source world is a glow with many gifts under the tree this week.

Linus’ gift to the DevOps world came in the form of <a href="http://www.mail-archive.com/linux-kernel@vger.kernel.org/msg1290645.html">Kernel 4.9</a>. This is a large release coming in at about <a href="http://www.phoronix.com/scan.php?page=news_item&amp;px=Linux-4.9-Git-Stats">22.3 million lines according to Michael Larabel</a> with, “<a href="http://www.mail-archive.com/linux-kernel@vger.kernel.org/msg1290645.html">a bit over two thirds drivers</a>”.

<a href="http://blog.kubernetes.io/2016/12/kubernetes-1.5-supporting-production-workloads.html">Kubernetes 1.5 is out</a>! PetSet is now called StatefulSet (beta), Windows server containers are available, and a slew of other features are now available.

<a href="https://blog.docker.com/2016/12/introducing-containerd/">Docker spun out containerd</a>, its core container runtime, into a standalone product, “and will be donating it to a neutral foundation early next year”. This is big news especially when you consider <a href="https://blog.docker.com/2016/12/containerd-core-runtime-component/">Alibaba, AWS, Google, IBM and Microsoft are all onboard the containerd train</a>.

<a href="https://coreos.com/blog/tectonic-self-driving.html">CoreOS Linux has rebranded itself to Container Linux</a>. The CoreOS team can do what they want as far as branding but I honestly think this is going to create some confusion in the marketplace. Think about explaining to an overzealous CTO who wants to use containers that Container Linux might not be the best option.

<a href="https://cloudplatform.googleblog.com/2016/12/Google-joins-the-Cloud-Foundry-Foundation.html">Google joined the Cloud Foundry Foundation</a>. Not really a surprise here as Google attempts to change the cloud playing field to a vendor agnostic landscape.

Yahoo announced <a href="https://yahoo.tumblr.com/post/154479236569/important-security-information-for-yahoo-users">ANOTHER hack affecting 1 billion users</a> (billion with a B). Interestingly the hashed passwords that were compromised used MD5 (which has been cryptographically compromised for years now). If you are in Ops, ask your Devs what they are doing to keep you out of the news.

This should not even have to be a thing but it is: <a href="http://neveragain.tech/">neveragain.tech</a>. It is a pledge from “engineers, designers, business executives, and others whose jobs include managing or processing data about people” who refuse to participate in overzealous government data collection programs. Where was this in 2001 or even after the Snowden reveal?

<a href="http://www.theregister.co.uk/2016/12/13/cisco_to_kill_its_intercloud_public_cloud_on_march_31st_2017/">Cisco announced it is going to kill its OpenStack based cloud platform</a>. This is more bad news for OpenStack. Rackspace helped invent OpenStack and invested a lot into it in error it would seem. <a href="http://www.computerworld.com/article/3146568/cloud-computing/and-there-she-goes-hpe-jettisons-both-openstack-and-cloud-foundry-initiatives.html">HP has bailed out of OpenStack</a>. <a href="https://twitter.com/KendrickColeman/status/809062687774806016">As Kendrick Coleman pointed out</a> to me on Twitter, <a href="http://www.theregister.co.uk/2016/12/02/dell_emc_kills_off_vxrack_neutrino/">Dell EMC ditched its OpenStack ambitions</a> as well recently.

<a href="https://www.bloomberg.com/news/articles/2016-12-15/github-is-building-a-coder-s-paradise-it-s-not-coming-cheap">Bloomberg did a great piece on GitHub</a> that showed the platform as an incredibly expensive venture. The piece also detailed what GitHub has done lately. But, I don’t think anyone realized how expensive it is to run GitHub. This is good exposure of a fundamental DevOps tool to the world outside of Dev and Ops for sure.

<a href="https://youtu.be/wewAC5X_CZ8">Jenn Schiffer’s talk at XOXO</a> is a little bit older but, it’s worth watching every week; timely and hilarious:

Should you ever need to figure out if an IP address is a TOR exit node or not just run this gem:

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Refreshment and Refurbishment

1. [Kernel 4.922.3 million lines according to Michael Larabela bit over two thirds drivers](http://www.mail-archive.com/linux-kernel@vger.kernel.org/msg1290645.html)

    Linus’ gift to the DevOps world came in the form of . This is a large release coming in at about  with, “”.
1. [Kubernetes 1.5 is out](http://blog.kubernetes.io/2016/12/kubernetes-1.5-supporting-production-workloads.html)

    ! PetSet is now called StatefulSet (beta), Windows server containers are available, and a slew of other features are now available.
1. [Docker spun out containerdAlibaba, AWS, Google, IBM and Microsoft are all onboard the containerd train](https://blog.docker.com/2016/12/introducing-containerd/)

    , its core container runtime, into a standalone product, “and will be donating it to a neutral foundation early next year”. This is big news especially when you consider .
1. [CoreOS Linux has rebranded itself to Container Linux](https://coreos.com/blog/tectonic-self-driving.html)

    . The CoreOS team can do what they want as far as branding but I honestly think this is going to create some confusion in the marketplace. Think about explaining to an overzealous CTO who wants to use containers that Container Linux might not be the best option.
### Department of Happy Little Clouds ☁️☁️☁️

1. [Google joined the Cloud Foundry Foundation](https://cloudplatform.googleblog.com/2016/12/Google-joins-the-Cloud-Foundry-Foundation.html)

    . Not really a surprise here as Google attempts to change the cloud playing field to a vendor agnostic landscape.
### Department of Dafuq

1. [ANOTHER hack affecting 1 billion users](https://yahoo.tumblr.com/post/154479236569/important-security-information-for-yahoo-users)

    Yahoo announced  (billion with a B). Interestingly the hashed passwords that were compromised used MD5 (which has been cryptographically compromised for years now). If you are in Ops, ask your Devs what they are doing to keep you out of the news.
1. [neveragain.tech](http://neveragain.tech/)

    This should not even have to be a thing but it is: . It is a pledge from “engineers, designers, business executives, and others whose jobs include managing or processing data about people” who refuse to participate in overzealous government data collection programs. Where was this in 2001 or even after the Snowden reveal?
### Department of Next Year’s Old Tech

1. [Cisco announced it is going to kill its OpenStack based cloud platformHP has bailed out of OpenStackAs Kendrick Coleman pointed outDell EMC ditched its OpenStack ambitions](http://www.theregister.co.uk/2016/12/13/cisco_to_kill_its_intercloud_public_cloud_on_march_31st_2017/)

    . This is more bad news for OpenStack. Rackspace helped invent OpenStack and invested a lot into it in error it would seem. .  to me on Twitter,  as well recently.
### Department of Sane Workplaces

1. [Bloomberg did a great piece on GitHub](https://www.bloomberg.com/news/articles/2016-12-15/github-is-building-a-coder-s-paradise-it-s-not-coming-cheap)

    that showed the platform as an incredibly expensive venture. The piece also detailed what GitHub has done lately. But, I don’t think anyone realized how expensive it is to run GitHub. This is good exposure of a fundamental DevOps tool to the world outside of Dev and Ops for sure.
1. [Jenn Schiffer’s talk at XOXO](https://youtu.be/wewAC5X_CZ8)

    is a little bit older but, it’s worth watching every week; timely and hilarious:
### DevOps’ish One-Liner of the Week

1. []()

    Should you ever need to figure out if an IP address is a TOR exit node or not just run this gem:
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-001.md) ------------- [ Next >> ](devopsweekly-003.md)