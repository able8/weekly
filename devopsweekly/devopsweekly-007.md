## [DevOps'ish 007](https://devopsish.com/007) - Sun Jan 22, 2017

What a busy week in the land of <a href="https://devopsish.com/">DevOps</a>! It seems like everyone has shaken off the holidays and snow days. So without further adieu let’s get to it!

Breanne Boland has a great piece on <a href="http://breanneboland.com/blog/2017/01/11/11-lessons-first-year-software-engineering/">lessons learned during her first year as a software engineer</a>. “Timing is everything” as always but her point on curiosity is spot on.

The Cloudcast had a fantastic episode titled, “<a href="http://www.thecloudcast.net/2017/01/the-cloudcast-286-balancing-monolithic.html">Balancing Monolithic Apps and Microservices</a>”. I had a lightbulb click on for one of the products I support while listening to this episode. Everyone has legacy infrastructure. This episode discusses a strategy to deal with the older parts of your applications.

Grant Shipley shared <a href="https://blog.openshift.com/openshift-developers-set-full-cluster-30-minutes/">how to get a full blown OpenShift cluster up and running in thirty minutes</a>. OpenShift is looking more and more appealing.

Jessie Frazelle hacked the bejesus out of CoreOS’ Container Linux. <a href="https://blog.jessfraz.com/post/ultimate-linux-on-the-desktop/">She put graphics drivers in CoreOS, y’all</a>… Graphics drivers! Jess is a glutton for punishment and I admire the hell out of her for it.

Eric Raymond is making the next version of NTP, aptly named NTPsec. He recently made <a href="https://blog.ntpsec.org/2017/01/18/rust-vs-go.html">a comparison of Rust vs. Go for the NTPsec project</a> that is a good assessment of both languages.

Remember <a href="../006/">last week</a>, when I linked to a piece from Red Hat about <a href="http://rhelblog.redhat.com/2017/01/13/selinux-mitigates-container-vulnerability/">SELinux stopping a Docker bug dead in its tracks</a>? <a href="http://thenewstack.io/docker-calls-out-red-hat/">Docker is not too happy about that</a> for some reason.

<a href="https://psiloveyou.xyz/help-me-im-on-tinder-i-don-t-want-to-see-your-code-711de4986ab8#.plk8m2y9b">Bruh… Please</a>. No one on Tinder wants to look at your code.

<a href="https://cloud.google.com/security/security-design/">Google shared their infrastructure security design</a> with the world. They go through a lot of effort to secure the data that makes them dominate. Google’s security measures are deep, “We also design custom chips, including a hardware security chip that is currently being deployed on both servers and peripherals. These chips allow us to securely identify and authenticate legitimate Google devices at the hardware level.”

<a href="http://www.theregister.co.uk/2017/01/19/cumulus_networks_writes_its_name_on_a_white_box/">Cumulus Networks is releasing its own hardware</a>. Cumulus is an open networking operating system and their teams make heavy use of Ansible. It is great to see them branding their own hardware.

yossorion shared their thoughts on <a href="https://gist.github.com/yossorion/4965df74fd6da6cdc280ec57e83a202d">what they wish they knew before they joined a unicorn</a>. Thinking about taking equity as compensation? I would recommend reading this.

Jérôme Petazzoni penned a fantastic piece on sexism in our industry. <a href="http://jpetazzo.github.io/2017/01/15/yes-all-men/">This shit is not okay</a> and our industry HAS to fix this right the fuck now (RTFN).

The one true way to grep IP addresses:

<code>egrep -o &#39;[0–9]{1,3}\.[0–9]{1,3}\.[0–9]{1,3}\.[0–9]{1,3}&#39;</code>

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [lessons learned during her first year as a software engineer](http://breanneboland.com/blog/2017/01/11/11-lessons-first-year-software-engineering/)

    Breanne Boland has a great piece on . “Timing is everything” as always but her point on curiosity is spot on.
1. [Balancing Monolithic Apps and Microservices](http://www.thecloudcast.net/2017/01/the-cloudcast-286-balancing-monolithic.html)

    The Cloudcast had a fantastic episode titled, “”. I had a lightbulb click on for one of the products I support while listening to this episode. Everyone has legacy infrastructure. This episode discusses a strategy to deal with the older parts of your applications.
1. [how to get a full blown OpenShift cluster up and running in thirty minutes](https://blog.openshift.com/openshift-developers-set-full-cluster-30-minutes/)

    Grant Shipley shared . OpenShift is looking more and more appealing.
1. [She put graphics drivers in CoreOS, y’all](https://blog.jessfraz.com/post/ultimate-linux-on-the-desktop/)

    Jessie Frazelle hacked the bejesus out of CoreOS’ Container Linux. … Graphics drivers! Jess is a glutton for punishment and I admire the hell out of her for it.
1. [a comparison of Rust vs. Go for the NTPsec project](https://blog.ntpsec.org/2017/01/18/rust-vs-go.html)

    Eric Raymond is making the next version of NTP, aptly named NTPsec. He recently made  that is a good assessment of both languages.
### Department of Dafuq

1. [last weekSELinux stopping a Docker bug dead in its tracksDocker is not too happy about that](../006/)

    Remember , when I linked to a piece from Red Hat about ?  for some reason.
1. [Bruh… Please](https://psiloveyou.xyz/help-me-im-on-tinder-i-don-t-want-to-see-your-code-711de4986ab8#.plk8m2y9b)

    . No one on Tinder wants to look at your code.
### Department of Data Defense

1. [Google shared their infrastructure security design](https://cloud.google.com/security/security-design/)

    with the world. They go through a lot of effort to secure the data that makes them dominate. Google’s security measures are deep, “We also design custom chips, including a hardware security chip that is currently being deployed on both servers and peripherals. These chips allow us to securely identify and authenticate legitimate Google devices at the hardware level.”
### Department of Refreshment and Refurbishment

1. [Cumulus Networks is releasing its own hardware](http://www.theregister.co.uk/2017/01/19/cumulus_networks_writes_its_name_on_a_white_box/)

    . Cumulus is an open networking operating system and their teams make heavy use of Ansible. It is great to see them branding their own hardware.
### Department of Sane Workplaces

1. [what they wish they knew before they joined a unicorn](https://gist.github.com/yossorion/4965df74fd6da6cdc280ec57e83a202d)

    yossorion shared their thoughts on . Thinking about taking equity as compensation? I would recommend reading this.
1. [This shit is not okay](http://jpetazzo.github.io/2017/01/15/yes-all-men/)

    Jérôme Petazzoni penned a fantastic piece on sexism in our industry.  and our industry HAS to fix this right the fuck now (RTFN).
### DevOps’ish One-Liner of the Week

1. []()

    The one true way to grep IP addresses:
1. []()

    egrep -o '[0–9]{1,3}\.[0–9]{1,3}\.[0–9]{1,3}\.[0–9]{1,3}'
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](sreweekly-6.md) ------------- [ Next >> ](sreweekly-8.md)