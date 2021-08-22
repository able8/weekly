## [DevOps'ish 032](https://devopsish.com/032) - Mon Jul 17, 2017

I attended <a href="https://gophercon.com/">GopherCon</a> in Denver this week (and <a href="https://twitter.com/system76/status/885951487284621312">visited System 76’s Denver office</a>). It was an unbelievably positive experience. <a href="https://medium.com/@GolangBridge/welcome-to-our-gophercon-2017-diversity-scholarship-winners-afcced449ba2">I attended as a diversity scholarship recipient</a> (being a disabled veteran has *one *perk). I had friends there and met up with them the first day I was in town (I made a lot more friends as the week went on). I got to meet the wonderful Sarah Adams from <a href="http://www.womenwhogo.org/">Women Who Go</a> at a welcome breakfast at Modern Market the first day of the conference. Sarah is an incredible person and her personality put me at ease. Sarah showed me that I was among friends at GopherCon. The conference organizers are the most down to earth and nicest people you could ever meet. Dave Cheney and I talked over breakfast one morning about the premise of my Lightning Talk and how *crypto/tls *in the Go standard library, “just works.” I have never been hugged at a conference before but I got quite a few hugs at GopherCon. The Go community is a truly great thing to be a part of.

Russ Cox laid out the case for, “not abandoning Go 1.” Russ also turned on my thought bubbles around the idea of <a href="https://github.com/golang/go/wiki/ExperienceReports">Experience Reports</a>. Kris Nova gave a refreshing talk about *rm -rf’<em>ing</em> <em>a cumbersome piece of a project and starting over</em>.* I watched in amazement as Liz Rice built a strace tool in Go from scratch (and now I want to know everything about <a href="http://man7.org/linux/man-pages/man2/ptrace.2.html">ptrace</a>). But, the highlight of the conference talks came from Ashley McNamara. She gave a heartfelt keynote that brought members of the audience to tears (I got goosebumps). The Go community cheered her on as she began to choke up on stage discussing her journey to Go. This was a very moving keynote and it was Ashley’s first one (no pressure for the next one, Smash). It was a truly great time and I think I made some very good friends and met more people than I ever thought possible. From Estonia and Cameroon to Japan and Australia; the globe is covered in half a million Go developers and 1,500 were in Denver with me.

You might have noticed last week that I changed up the <a href="https://devopsish.com/">DevOps’ish</a> weekly headlines from an epoch to a description of what’s in the newsletter. This week, I am ditching the whimsical and creative section heads and consolidating into my <a href="https://speakerdeck.com/chrisshort/a-night-of-devops?slide=21">DevOps Three Things</a>. This will simplify the weekly creation process for the sake of speed.

<a href="https://2017.fullstackfest.com"><strong>Full Stack Fest 2017: Problems of today, wonders from the future.</strong></a><br/>Barcelona, 4–8 Sept. 2017<br/>Are you a curious mind? <a href="https://2017.fullstackfest.com">Full Stack Fest</a> is a week-long conference based in the amazing city of Barcelona that peeks into the web of tomorrow. Serverless, Blockchain, WebVR, Distributed Web, Progressive Web Apps… Come and see! Early bird tickets available with a 10% discount using the code DEVOPSISH. <em>SPONSORED</em>

### People

1. [I no longer feel like I can trust ‘init’ to do the sane thing](https://lkml.org/lkml/2017/7/6/577)

    Linus Torvalds said, “” when discussing rlimit defaults for setuid execs.
1. [Why Speak At A Conference?](https://itrevolution.com/why-speak-at-a-conference/)

    Because you have an interesting story to tell.
1. [URLs are UICool URIs don’t change](https://www.hanselman.com/blog/URLsAreUI.aspx)

     Tim Berners-Lee said “” in 1998.
### Process

1. [YES! YOUR SITE NEEDS HTTPS](https://doesmysiteneedhttps.com/)

    
1. [EFF Officially Appeals Tim Berners-Lee Decision On DRM In HTMLa notice of appeal](https://www.techdirt.com/articles/20170712/10262037770/eff-officially-appeals-tim-berners-lee-decision-drm-html.shtml)

     W3C member EFF has now filed  on the decision.
1. [Docker accidentally cuts off Ukraine](https://www.theregister.co.uk/2017/07/13/docker_accidentally_cuts_off_ukraine/)

     Did you know that US sanctions extend to the digital world too?
1. [Valuable Lessons in Over-Engineering the Core of Kubernetes kops, Kris Nova](https://about.sourcegraph.com/go/valuable-lessons-in-over-engineering-the-core-of-kubernetes-kops)

    at GopherCon
1. [Millions of Verizon customer records exposed in security lapse](http://www.zdnet.com/article/millions-verizon-customer-records-israeli-data/)

     Customer records for at least 14 million subscribers, including phone numbers and account PINs, were exposed.
1. [The .io Error](https://thehackerblog.com/the-io-error-taking-control-of-all-io-domains-with-a-targeted-registration/)

     Taking Control of All .io Domains With a Targeted Registration
1. [Design patterns for microservices](https://azure.microsoft.com/en-us/blog/design-patterns-for-microservices/)

     These nine patterns are particularly useful when designing and implementing microservices.
1. [How to Make a Strategic, Value-Driven Business Case for Your DevOps Initiative](https://www.contino.io/insights/how-to-make-a-strategic-value-driven-business-case-for-your-devops-initiative)

    
1. [gandi.net Detailed incident report](https://news.gandi.net/en/2017/07/detailed-incident-report/)

     On Friday July 7, an unauthorized connection to one of our technical partners resulted in the modification of the name servers [NS] of 751 domain names which then pointed traffic to the impacted domains to a malicious site.
### Tools

1. [Cluster Schedulers](https://medium.com/@cindysridharan/schedulers-kubernetes-and-nomad-b0f2e14a896)

     “I’ve invariably been skeptical about the necessity of schedulers, especially at smaller companies, albeit finding schedulers themselves fascinating to learn about out of curiosity.”
1. [Serverless Hosting Comparison](https://headmelted.com/serverless-showdown-4a771ca561d2)

     A head-to-head showdown of the providers fighting for your dollars.
1. [kubicorn](https://github.com/kris-nova/kubicorn)

     Create, manage, snapshot, and scale Kubernetes infrastructure in the public cloud. Kris Nova’s unofficial project that solves the Kubernetes infrastructure problem and gives users a rich golang library to work with infrastructure.
1. [The Benefits of Running Kubernetes on Google Container Engine](https://www.reactiveops.com/blog/benefits-of-running-kubernetes-on-google-container-engine/)

     “if you want to run Kubernetes on AWS, you have a lot of work ahead of you. If you’re looking for container-based workflows, GCP will make Kubernetes a lot easier to manage and keep up to date.”
1. [How to Build a Tiny Httpd ContainerSmithapache httpd](https://hackernoon.com/how-to-build-a-tiny-httpd-container-ae622c37db39)

     “ can shrink a 69MB  container down to under 2MB, resulting in space savings of 97%, while making it more secure and easier to operate.”
1. [githubdocs](https://www.npmjs.com/package/githubdocs)

     Easily build a searchable documentation Single Page Application (SPA) using markdown files in your Github Repo.
1. [Alibaba Cloud](https://www.alibabacloud.com/)

     Alibaba Cloud provides a comprehensive suite of global cloud computing services to power both our international customers’ online businesses and Alibaba Group’s own e-commerce ecosystem.
1. [It Turns Out, 2017 is the Year of Simply Secure PHP Cryptography](https://paragonie.com/blog/2017/07/it-turns-out-2017-is-year-simply-secure-php-cryptography)

     PHP is Getting ext/sodium in 7.2
1. [Upspin](https://upspin.io/)

    is an experimental project to build a framework for naming and sharing files and other data securely, uniformly, and globally: a global name system of sorts. It is not a file system, but a set of protocols and reference implementations that can be used to join things like file systems and other storage services to the name space.
1. [deadman-check](https://github.com/sepulworld/deadman-check)

     Monitoring companion for Nomad periodic jobs and Cron
1. [Cachet](https://cachethq.io/)

     The open source status page system
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-031.md) ------------- [ Next >> ](devopsweekly-033.md)