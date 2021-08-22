## [DevOps'ish 031](https://devopsish.com/031) - Sun Jul 9, 2017

This week was nuts. We dealt with a service degradation due to issues between our CDN and data center edge network. On a long, holiday weekend it’s the last thing you want to deal with. There were eighteen action items to come out of the post mortem (which was long enough to merit an executive summary). This was one of the weirder issues I have ever seen as it involved three different companies and a few bad hops in a network path. <a href="../029/">I quoted Baruch Sadogursky two weeks ago</a>, “Pain is instructional.” It definitely is! Meanwhile, my company, Bankrate, was <a href="http://www.thedrum.com/news/2017/07/04/red-ventures-digs-deep-bankrate-website-124bn-personal-finance-drive">getting acquired by Red Ventures</a> for $1.24 billion.

Hopefully, the week ahead will be calmer. Doubtful given I’ll be in Atlanta on Monday, Denver on Tuesday, back in North Carolina on Sunday, and then back to Michigan later the following. Next week’s edition of DevOps’ish Weekly will be written right after the conclusion of <a href="https://gophercon.com/">GopherCon 2017</a>! I’m attending GopherCon for the first time ever and am so thankful to have the opportunity to do so. If you’re there ask me for a DevOps’ish sticker.

<a href="https://2017.fullstackfest.com"><strong>Full Stack Fest 2017: Problems of today, wonders from the future.</strong></a><br/>Barcelona, 4–8 Sept. 2017<br/>Are you a curious mind? <a href="https://2017.fullstackfest.com">Full Stack Fest</a> is a week-long conference based in the amazing city of Barcelona that peeks into the web of tomorrow. Serverless, Blockchain, WebVR, Distributed Web, Progressive Web Apps… Come and see! Early bird tickets available with a 10% discount using the code DEVOPSISH. <em>SPONSORED</em>

<a href="https://letsencrypt.org//2017/07/06/wildcard-certificates-coming-jan-2018.html">Let’s Encrypt Wildcard Certificates Coming January 2018</a>: “Wildcard certificates will be offered free of charge via our upcoming ACME v2 API endpoint. We will initially only support base domain validation via DNS for wildcard certificates, but may explore additional validation options over time.”

<a href="https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go.html">HTTPS for free in Go, with little help of Let’s Encrypt</a>: Here’s how to start HTTPS web server that uses free SSL certificates from Let’s Encrypt.

<a href="https://github.com/PressLabs/zinc">zinc</a>: Route 53 zone manager. Policy Records on the Cheap.

<a href="https://github.com/mozilla/sops">sops</a>: sops is an editor of encrypted files that supports YAML, JSON and BINARY formats and encrypts with AWS KMS and PGP (by Mozilla).

<a href="https://github.com/goreleaser/goreleaser">goreleaser</a>: GoReleaser builds Go binaries for several platforms, creates a GitHub release and then pushes a Homebrew formula to a repository. All that wrapped in your favorite CI.

<a href="http://blog.ssdnodes.com/blog/ultimate-guide-self-hosted-alternatives-to-cloud-apps">Ultimate Guide: Self-Hosted Alternatives to Cloud Apps</a>

<a href="http://labs.play-with-k8s.com/">Kubernetes Playground</a>: This is so cool! Stand up a Kubernetes cluster with dashboard and everything in seconds to tinker with for up to four hours at a time.

<a href="http://ryanlue.com/posts/2017-06-29-gpg-for-ssh-auth">Using a GPG key for SSH Authentication</a>: PGP keys and SSH keys are different. It took me longer than I care to admit to understand this, given that the first StackExchange thread I found on the subject made it pretty clear that it’s not trivial to use one where you’re supposed to use the other.

<a href="https://medium.com/the-u-s-digital-service/how-a-20-year-old-kernel-feature-helped-usds-improve-vas-network-33109cbcb2e6">How a 20-year old kernel feature helped USDS improve VA’s network</a>: “this bug is fascinating because we started debugging Ruby on Rails from AWS and found the root cause in the NSOC Cisco Routers thanks to an ancient feature in the Linux kernel.”

<a href="https://github.com/cloudnativelabs/kube-shell">kube-shell</a>: Kubernetes shell - An integrated shell for working with the Kubernetes CLI

<a href="https://github.com/garethr/kubeval">kubeval</a>: Validate your Kubernetes configuration files, supports multiple Kubernetes versions

<a href="https://engineering.linkedin.com/blog/2017/06/open-sourcing-iris-and-oncall">Open Sourcing Iris and Oncall</a>: Iris, named after the Greek goddess of messages, is our open-sourced solution for incident escalation and reliable messaging, and has provided LinkedIn with fast, automated escalations for almost two years now.

<a href="https://www.bleepingcomputer.com/news/security/openbsd-will-get-unique-kernels-on-each-reboot-do-you-hear-that-linux-windows/">OpenBSD Will Get Unique Kernels on Each Reboot</a>: A new feature added in test snapshots for OpenBSD releases will create a unique kernel every time an OpenBSD user reboots or upgrades his computer.

<a href="https://medium.com/netflix-techblog/neflix-platform-engineering-were-just-getting-started-267f65c4d1a7">Netflix Platform Engineering — we’re just getting started</a>: As Netflix continues to evolve and grow, so do our engineering challenges. The nature of such challenges changes over time — from “greenfield” projects, to “scaling” activities, to “operationalizing” endeavors — all at great scale and break-neck speed.

<a href="https://www.informationweek.com/devops/rip-systems-administrator-welcome-devops/a/d-id/1329208">RIP, Systems Administrator, Welcome DevOps</a>: While the traditional sys admin role may fade away, those who are willing to change can take an exciting journey with cloud and DevOps.

<a href="https://redfin.engineering/two-commits-that-wrecked-the-user-experience-of-git-f0075b77eab1">Two Commits That Wrecked the User Experience of Git</a>: Git didn’t have to be so obnoxious, but these two commits set a terrible, unfixable precedent

<a href="http://squad-twelve.com/2017/07/03/7-specific-suggestions-to-sabotage-devops-simply/">7 Specific Suggestions to Sabotage DevOps Simply</a>: Have you ever felt like people were conspiring to make DevOps fail? They probably had good intentions but they just made DevOps more difficult than it should be. What if it was actually intentional? The resistance gets organized…

<a href="https://www.bloomberg.com/news/articles/2017-07-06/soundcloud-cuts-40-percent-of-staff-in-bid-to-remain-independent">SoundCloud Cuts 40% of Staff in Push for Profitability</a>: Music service will close offices in San Francisco, London

<a href="http://www.cnbc.com/2017/07/06/microsoft-will-layoff-thousands-of-employees.html">Microsoft plans thousands of job cuts in a sales staff overhaul to fuel cloud growth</a>: Microsoft’s layoffs will mostly affect sales, and thousands of jobs will be cut. Most of the changes will affect employees outside of the U.S.

<a href="https://medium.com/@ashleymcnamara/depression-is-not-a-sign-of-weakness-478d55ba66f9">Depression is not a sign of weakness</a>: “What you don’t see is how terribly afraid I am of my own code, what if it’s bad? What if they laugh at me? What if they see how much I actually don’t know? Will I ever know enough?”

<a href="https://medium.com/@OlarkLiveChat/its-2017-and-mental-health-is-still-an-issue-in-the-workplace-61efbef092f">It’s 2017 and Mental Health is still an issue in the workplace</a>: “We are in a knowledge economy. Our jobs require us to execute at peak mental performance. When an athlete is injured they sit on the bench and recover. Let’s get rid of the idea that somehow the brain is different.”

<a href="https://devops.com/finding-landing-right-job/">Finding, and Landing, the Right Job for You</a>: Finding the right job is really hard. Take a look at <a href="https://www.linkedin.com/in/christopherbshort/?lipi=urn%3Ali%3Apage%3Ad_flagship3_feed%3BxgmsXBp6RqOuzgj8Wke%2B7Q%3D%3D&amp;licu=urn%3Ali%3Acontrol%3Ad_flagship3_feed-identity_welcome_message">my LinkedIn profile</a> and you’ll see evidence of this fact. Finding the right balance of challenge and culture is very hard. These tips on navigating job hunting from <a href="https://twitter.com/ConorDevOps">Conor Delanbanque</a> are a good start.

<a href="http://blog.triplebyte.com/how-to-interview-engineers">How to Interview Engineers</a>: On the other side of job seeking is hiring. As a hiring manager, I am finding it very challenging to get people to open up during that initial phone call. I want details but not rambling, I want low ego but please tell me what you’re proud of. This how to is good and I especially think asking every candidate the same questions is very important.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Let’s Encrypt Wildcard Certificates Coming January 2018](https://letsencrypt.org//2017/07/06/wildcard-certificates-coming-jan-2018.html)

     “Wildcard certificates will be offered free of charge via our upcoming ACME v2 API endpoint. We will initially only support base domain validation via DNS for wildcard certificates, but may explore additional validation options over time.”
1. [HTTPS for free in Go, with little help of Let’s Encrypt](https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go.html)

     Here’s how to start HTTPS web server that uses free SSL certificates from Let’s Encrypt.
1. [zinc](https://github.com/PressLabs/zinc)

     Route 53 zone manager. Policy Records on the Cheap.
1. [sops](https://github.com/mozilla/sops)

     sops is an editor of encrypted files that supports YAML, JSON and BINARY formats and encrypts with AWS KMS and PGP (by Mozilla).
1. [goreleaser](https://github.com/goreleaser/goreleaser)

     GoReleaser builds Go binaries for several platforms, creates a GitHub release and then pushes a Homebrew formula to a repository. All that wrapped in your favorite CI.
1. [Ultimate Guide: Self-Hosted Alternatives to Cloud Apps](http://blog.ssdnodes.com/blog/ultimate-guide-self-hosted-alternatives-to-cloud-apps)

    
1. [Kubernetes Playground](http://labs.play-with-k8s.com/)

     This is so cool! Stand up a Kubernetes cluster with dashboard and everything in seconds to tinker with for up to four hours at a time.
1. [Using a GPG key for SSH Authentication](http://ryanlue.com/posts/2017-06-29-gpg-for-ssh-auth)

     PGP keys and SSH keys are different. It took me longer than I care to admit to understand this, given that the first StackExchange thread I found on the subject made it pretty clear that it’s not trivial to use one where you’re supposed to use the other.
1. [How a 20-year old kernel feature helped USDS improve VA’s network](https://medium.com/the-u-s-digital-service/how-a-20-year-old-kernel-feature-helped-usds-improve-vas-network-33109cbcb2e6)

     “this bug is fascinating because we started debugging Ruby on Rails from AWS and found the root cause in the NSOC Cisco Routers thanks to an ancient feature in the Linux kernel.”
1. [kube-shell](https://github.com/cloudnativelabs/kube-shell)

     Kubernetes shell - An integrated shell for working with the Kubernetes CLI
1. [kubeval](https://github.com/garethr/kubeval)

     Validate your Kubernetes configuration files, supports multiple Kubernetes versions
1. [Open Sourcing Iris and Oncall](https://engineering.linkedin.com/blog/2017/06/open-sourcing-iris-and-oncall)

     Iris, named after the Greek goddess of messages, is our open-sourced solution for incident escalation and reliable messaging, and has provided LinkedIn with fast, automated escalations for almost two years now.
1. [OpenBSD Will Get Unique Kernels on Each Reboot](https://www.bleepingcomputer.com/news/security/openbsd-will-get-unique-kernels-on-each-reboot-do-you-hear-that-linux-windows/)

     A new feature added in test snapshots for OpenBSD releases will create a unique kernel every time an OpenBSD user reboots or upgrades his computer.
1. [Netflix Platform Engineering — we’re just getting started](https://medium.com/netflix-techblog/neflix-platform-engineering-were-just-getting-started-267f65c4d1a7)

     As Netflix continues to evolve and grow, so do our engineering challenges. The nature of such challenges changes over time — from “greenfield” projects, to “scaling” activities, to “operationalizing” endeavors — all at great scale and break-neck speed.
### Department of Dafuq

1. [RIP, Systems Administrator, Welcome DevOps](https://www.informationweek.com/devops/rip-systems-administrator-welcome-devops/a/d-id/1329208)

     While the traditional sys admin role may fade away, those who are willing to change can take an exciting journey with cloud and DevOps.
1. [Two Commits That Wrecked the User Experience of Git](https://redfin.engineering/two-commits-that-wrecked-the-user-experience-of-git-f0075b77eab1)

     Git didn’t have to be so obnoxious, but these two commits set a terrible, unfixable precedent
1. [7 Specific Suggestions to Sabotage DevOps Simply](http://squad-twelve.com/2017/07/03/7-specific-suggestions-to-sabotage-devops-simply/)

     Have you ever felt like people were conspiring to make DevOps fail? They probably had good intentions but they just made DevOps more difficult than it should be. What if it was actually intentional? The resistance gets organized…
### Department of Assemblage Obtainment

1. [SoundCloud Cuts 40% of Staff in Push for Profitability](https://www.bloomberg.com/news/articles/2017-07-06/soundcloud-cuts-40-percent-of-staff-in-bid-to-remain-independent)

     Music service will close offices in San Francisco, London
1. [Microsoft plans thousands of job cuts in a sales staff overhaul to fuel cloud growth](http://www.cnbc.com/2017/07/06/microsoft-will-layoff-thousands-of-employees.html)

     Microsoft’s layoffs will mostly affect sales, and thousands of jobs will be cut. Most of the changes will affect employees outside of the U.S.
### Department of Sane Workplaces

1. [Depression is not a sign of weakness](https://medium.com/@ashleymcnamara/depression-is-not-a-sign-of-weakness-478d55ba66f9)

     “What you don’t see is how terribly afraid I am of my own code, what if it’s bad? What if they laugh at me? What if they see how much I actually don’t know? Will I ever know enough?”
1. [It’s 2017 and Mental Health is still an issue in the workplace](https://medium.com/@OlarkLiveChat/its-2017-and-mental-health-is-still-an-issue-in-the-workplace-61efbef092f)

     “We are in a knowledge economy. Our jobs require us to execute at peak mental performance. When an athlete is injured they sit on the bench and recover. Let’s get rid of the idea that somehow the brain is different.”
1. [Finding, and Landing, the Right Job for Youmy LinkedIn profileConor Delanbanque](https://devops.com/finding-landing-right-job/)

     Finding the right job is really hard. Take a look at  and you’ll see evidence of this fact. Finding the right balance of challenge and culture is very hard. These tips on navigating job hunting from  are a good start.
1. [How to Interview Engineers](http://blog.triplebyte.com/how-to-interview-engineers)

     On the other side of job seeking is hiring. As a hiring manager, I am finding it very challenging to get people to open up during that initial phone call. I want details but not rambling, I want low ego but please tell me what you’re proud of. This how to is good and I especially think asking every candidate the same questions is very important.

### [ << Prev ](sreweekly-30.md) ------------- [ Next >> ](sreweekly-32.md)