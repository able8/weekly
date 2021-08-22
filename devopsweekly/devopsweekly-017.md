## [DevOps'ish 017](https://devopsish.com/017) - Sun Apr 2, 2017

I started a new on-call rotation this week and it has really sucked. The team whose rotation I was added on to is mature in a few areas but monitoring and alerting seems to be lacking a little. They acknowledge that and my coming on to the rotation is partially to identify areas of improvement. The rotation I was on a couple weeks ago for a different team resulted in no pages all week and my bringing the team donuts.

This week’s rotation has resulted in a page every night (and two shorter maintenance events). The team that got donuts is very weak in some areas that this week’s team is not. But the donut team focused on monitoring and alerting as well as not waking people up and it shows.

Part of my job is to bring consistency to these teams that are separated 3,676 miles apart. Sharing expertise, processes, and procedures is part of the gig. The deficiencies in tooling have been identified. The need for personnel has been fulfilled. Now it’s time to bring on the processes and results.

The good news is I organized and lead my first <a href="https://tridevops.com">Triangle DevOps</a> meetup on Wednesday night. It was a blast! <a href="https://twitter.com/Josh_Atwell">Josh Atwell</a> did a superb job speaking (the GIF is glorious), <a href="http://www.selectgroup.com/">The Select Group</a> fed us well, and <a href="http://transloc.com/">Transloc</a> was a fantastic host. Thanks to everyone for a successful first rodeo!

I think it’s save to say Jessie Frazzle broke the Internet with her exquisite comparison of <a href="https://blog.jessfraz.com/post/containers-zones-jails-vms/">containers vs. Zones vs. Jails vs. VMs</a>. Also, dudes need to stop telling Jess how containers work.

Julia Evans put together a wonderful article on <a href="http://jvns.ca/blog/2017/03/26/bash-quirks/">Bash scripting quirks and safety tips</a>.

freeCodeCamp has put together a list of the <a href="https://medium.freecodecamp.com/the-10-github-repos-people-mention-the-most-in-freecodecamps-main-chat-room-189750600fa4">10 most mentioned GitHub repos for new developers</a>. My personal favorite is the <a href="https://github.com/vhf/free-programming-books">free-programming-books</a> repo.

Get Into DevOps put together a very useful piece on <a href="https://getintodevops.com/blog/keeping-the-whale-happy-how-to-clean-up-after-docker">tidying up after the rather messy Docker</a>.

<a href="https://thenewstack.io/alpine-linux-heart-docker/">Meet Alpine Linux, Docker’s Distribution of Choice for Containers</a>

<a href="https://daniel.haxx.se/blog/2017/03/27/curl-is-c/">Curl is written in C and here is why</a>.

Upheaval in the Linux Kernel world! <a href="https://lkml.org/lkml/2017/3/31/641">A cadre of kernel developers are going to Make Linux Great Again</a>. 😉😉😉

Congress made VPNs all the rage this week when they <a href="http://www.dslreports.com/shownews/The-GOP-Just-Killed-Consumer-Broadband-Privacy-Protections-139244">repealed broadband privacy protections</a>. The amount of FUD out there regarding VPNs is really high. When my wife is reading about <a href="https://www.buzzfeed.com/nicolenguyen/how-to-keep-your-browsing-history-actually-private">VPNs on Buzzfeed</a> we have reached a story’s critical mass. Remember, a VPN will only change your jurisdiction, not what you are actually surfing. At some point in your packets’ journey, your traffic is still trackable and traceable. You can reduce the risk but you cannot completely eliminate it. I recommend using an ad blocker like <a href="https://www.ghostery.com/">Ghostery</a>, use a VPN you are actually invested in (read the terms and pay for it), and generating a ton of bogus traffic might be a good idea too. <a href="https://cs.nyu.edu/trackmenot/">NYU has launched TrackMeNot to generate noise and obfuscation</a>.

<a href="http://www.pcworld.com/article/3186748/security/millions-of-websites-affected-by-unpatched-flaw-in-microsoft-iis-6-web-server.html">IIS 6 has a gnarly zero day that Microsoft has stated it will not patch</a>.

<a href="https://access.redhat.com/support/policy/updates/errata">EL5 (RHEL and CentOS) is officially end of life</a>. You should have purged it from your infrastructure already.

<a href="http://events.linuxfoundation.org/events/cloudnativecon-and-kubecon-europe">KubeCon Europe</a> wrapped up this past week. There were quite a few nuggets of great information from the event. <a href="https://www.thenewstack.io/tag/KubeCon-Europe-2017">The New Stack had fantastic coverage of KubeCon Europe</a>. Part of this coverage included a great piece on <a href="https://thenewstack.io/kubernetes-federation-post-configuration-management-universe/">Kelsey Hightower’s talk on Kubernetes Federation</a>.

<a href="https://blog.codedellemc.com/2017/03/29/cloud-native-computing-foundation-announces-dell-technologies-platinum-member/">Dell EMC has joined the Cloud Native Computing Foundation</a>.

<a href="https://blogs.msdn.microsoft.com/bharry/2017/03/31/shutting-down-codeplex/">Microsoft is shutting down CodePlex and recommending people use GitHub</a>.

<a href="http://larahogan.me/blog/first-one-on-one-questions/">Lara Hogan has put together a list of 1-on-1 questions</a> that I feel would have helped me stay at some company’s longer if my manager had known the answers to these questions.

<a href="https://venturebeat.com/2017/03/31/cloudera-files-to-raise-200-million-in-ipo/">Cloudera has filed for an IPO on the NYSE</a> (CLDR). They are trying to raise $200 million.

<a href="https://opensource.googleblog.com/2017/03/a-new-home-for-google-open-source.html?m=1">Google announced</a> their new home for open source, <a href="https://opensource.google.com/">Google Open Source</a>. Google has opened a treasure trove of documentation to help companies advance their open source initiatives. <a href="https://changelog.com/podcast/245">The Changelog episode with Will Norris</a> announcing Google Open Source is a great listen.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [containers vs. Zones vs. Jails vs. VMs](https://blog.jessfraz.com/post/containers-zones-jails-vms/)

    I think it’s save to say Jessie Frazzle broke the Internet with her exquisite comparison of . Also, dudes need to stop telling Jess how containers work.
1. [Bash scripting quirks and safety tips](http://jvns.ca/blog/2017/03/26/bash-quirks/)

    Julia Evans put together a wonderful article on .
1. [10 most mentioned GitHub repos for new developersfree-programming-books](https://medium.freecodecamp.com/the-10-github-repos-people-mention-the-most-in-freecodecamps-main-chat-room-189750600fa4)

    freeCodeCamp has put together a list of the . My personal favorite is the  repo.
1. [tidying up after the rather messy Docker](https://getintodevops.com/blog/keeping-the-whale-happy-how-to-clean-up-after-docker)

    Get Into DevOps put together a very useful piece on .
1. [Meet Alpine Linux, Docker’s Distribution of Choice for Containers](https://thenewstack.io/alpine-linux-heart-docker/)

    
1. [Curl is written in C and here is why](https://daniel.haxx.se/blog/2017/03/27/curl-is-c/)

    .
1. [A cadre of kernel developers are going to Make Linux Great Again](https://lkml.org/lkml/2017/3/31/641)

    Upheaval in the Linux Kernel world! . 😉😉😉
### Department of Data Defense

1. [repealed broadband privacy protectionsVPNs on BuzzfeedGhosteryNYU has launched TrackMeNot to generate noise and obfuscation](http://www.dslreports.com/shownews/The-GOP-Just-Killed-Consumer-Broadband-Privacy-Protections-139244)

    Congress made VPNs all the rage this week when they . The amount of FUD out there regarding VPNs is really high. When my wife is reading about  we have reached a story’s critical mass. Remember, a VPN will only change your jurisdiction, not what you are actually surfing. At some point in your packets’ journey, your traffic is still trackable and traceable. You can reduce the risk but you cannot completely eliminate it. I recommend using an ad blocker like , use a VPN you are actually invested in (read the terms and pay for it), and generating a ton of bogus traffic might be a good idea too. .
1. [IIS 6 has a gnarly zero day that Microsoft has stated it will not patch](http://www.pcworld.com/article/3186748/security/millions-of-websites-affected-by-unpatched-flaw-in-microsoft-iis-6-web-server.html)

    .
1. [EL5 (RHEL and CentOS) is officially end of life](https://access.redhat.com/support/policy/updates/errata)

    . You should have purged it from your infrastructure already.
### Department of Discussion

1. [KubeCon EuropeThe New Stack had fantastic coverage of KubeCon EuropeKelsey Hightower’s talk on Kubernetes Federation](http://events.linuxfoundation.org/events/cloudnativecon-and-kubecon-europe)

    wrapped up this past week. There were quite a few nuggets of great information from the event. . Part of this coverage included a great piece on .
### Department of Happy Little Clouds

1. [Dell EMC has joined the Cloud Native Computing Foundation](https://blog.codedellemc.com/2017/03/29/cloud-native-computing-foundation-announces-dell-technologies-platinum-member/)

    .
1. [Microsoft is shutting down CodePlex and recommending people use GitHub](https://blogs.msdn.microsoft.com/bharry/2017/03/31/shutting-down-codeplex/)

    .
### Department of Sane Workplaces

1. [Lara Hogan has put together a list of 1-on-1 questions](http://larahogan.me/blog/first-one-on-one-questions/)

    that I feel would have helped me stay at some company’s longer if my manager had known the answers to these questions.
### Department of Assemblage Obtainment

1. [Cloudera has filed for an IPO on the NYSE](https://venturebeat.com/2017/03/31/cloudera-files-to-raise-200-million-in-ipo/)

    (CLDR). They are trying to raise $200 million.
### Not DevOps But Still Cool

1. [Google announcedGoogle Open SourceThe Changelog episode with Will Norris](https://opensource.googleblog.com/2017/03/a-new-home-for-google-open-source.html?m=1)

    their new home for open source, . Google has opened a treasure trove of documentation to help companies advance their open source initiatives.  announcing Google Open Source is a great listen.

### [ << Prev ](sreweekly-16.md) ------------- [ Next >> ](sreweekly-18.md)