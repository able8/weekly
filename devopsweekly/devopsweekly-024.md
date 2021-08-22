## [DevOps'ish 024](https://devopsish.com/024) - Sun May 21, 2017

Ever have one of those weeks where you feel like more of a consultant than a full-time employee? That was my week. I do not mind it at all. But, sometimes it’s nice to not be bombarded with questions all week long. Also, I hate building VPNs.

In other news, I received a request to make the newsletter more responsive and readable. If you are reading this in a mail client then you should be benefitting from the hour or so I put into the new template. Reply to let me know what you think.

<a href="https://bsima.me/clog/server-tools.html">Ben Sima shared the first three things he installs on any new server.</a> I’m surprised he doesn’t have this in Ansible for easy deployment.

<a href="https://techbeacon.com/being-agile-working-smart-are-not-same-thing">Being agile and working smart are not the same thing</a> (TechBeacon)

Matt Micene reminds us all that DevOps isn’t some apparition that appeared out of nowhere. <a href="https://opensource.com/open-organization/17/5/what-is-the-point-of-DevOps">DevOps is the course correction to a culture that is a byproduct of Taylorism and Sloanianism</a>.

<a href="https://dev.to/mikesimons/mikes-monster-list-of-docker-tips">Mike Simmons shares his monster list of Docker tips</a>. There are definitely some container nuggets in here.

<a href="https://medium.com/@anthonypjshaw/ansible-v-s-salt-saltstack-v-s-stackstorm-3d8f57149368">Anthony Shaw weighs the strengths and weaknesses of Ansible vs. Salt vs. StackStorm</a>. Ansible fairs very well showing that it is standing the test of time.

<a href="http://heap.engineering/basic-performance-analysis-saved-us-millions/">Heap shared a solid article on Postgres performance tuning</a> (despite the click-bait title).

<a href="http://blog.kubernetes.io/2017/05/kubernetes-monitoring-guide.html">Kubernetes put out a piece on monitoring</a> and how k8s shifts the mindset required to effectively monitor your infrastructure.

<a href="https://cloud.google.com/spanner/">Cloud Spanner</a>, the world’s first horizontally-scalable and strongly-consistent relational database service, is now generally available for your mission-critical OLTP applications. This is great news for folks needing five-nines of availability, strong consistency, and low latency.

<a href="https://www.hashicorp.com/blog/vagrant-1-9-5/">Vagrant 1.9.5 is out</a>! This release features Docker Compose support among other things.

Ever want a completely open source software forge? <a href="http://allura.apache.org/">Apache Allura</a> offers a git repo, wiki, and ticketing system with 2FA.

<a href="https://ma.ttias.be/centos-7-4-ship-tls-1-2-alpn/">HTTP/2 coming to EL7 (finally)</a>: OpenSSL has been rebased in RHEL 7 to OpenSSL 1.0.2k. This will allow for Application-Layer Protocol Negotiation (ALPN) which is necessary for HTTP/2.

<a href="https://stackoverflow.blog/2017/05/16/stack-overflow-official-app-launches-ios-android/">Stack Overflow has a new iOS and Android app</a>. I’m not sure how useful this will be to me; Google has always been my conduit to Stack Overflow. But, if you slap it on your phone and use it more power to you.

<a href="https://open.dgraph.io/post/badger/">Introducing Badger: A fast key-value store written natively in Go</a> (Dgraph)

<a href="https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG.md/#v164">Kubernetes 1.6.4 is out</a> as the k8s squad pushes towards 1.7.

<a href="http://www.csoonline.com/article/3196974/data-protection/unmanaged-orphaned-ssh-keys-remain-a-serious-enterprise-risks.html">Unmanaged SSH keys are truly awful</a>. How are you going to handle employees coming and going (sometimes abruptly) without central management of SSH credentials? It’s surprisingly a bigger problem than you’d think.

What happens where you get hacked? <a href="https://panic.com/blog/stolen-source-code/">A developer at Panic shared their experience when, through a random series of events, got pwned</a>. It is an insightful postmortem of personal accountability.

<a href="https://blogs.microsoft.com/on-the-issues/2017/05/14/need-urgent-collective-action-keep-people-safe-online-lessons-last-weeks-cyberattack/#sm.0000llg81y90hdbs115975pt98jep">The need for urgent collective action to keep people safe online: Lessons from last week’s cyberattack</a> (Microsoft)

My employer, <a href="http://www.solarwinds.com/">SolarWinds</a>, has acquired <a href="https://scoutapp.com/">Scout</a>. Scout was in the server monitoring and APM space. <a href="http://www.solarwinds.com/company/press-releases/solarwinds-acquires-scouts-saas-based-server-monitoring-technology-and-launches-it-as-solarwinds-pingdom-server-monitor">This acquisition</a> dovetails nicely into both the traditional SolarWinds monitoring space and its cloud business.

<a href="https://twitter.com/kelseyhightower/status/864857551942307840">Kelsey Hightower shared how he goes about learning new things</a>. Reading RFCs is truly insightful.

<a href="https://begriffs.com/posts/2017-05-17-linux-workstation-guide.html#core">Joe Nelson broke down what the quintessential Linux workstation would look like today</a>. Asus is still the manufacturer of choice for the motherboard.

The University of Washington has a new class, <a href="http://callingbullshit.org/syllabus.html">Calling Bullshit in the Age of Big Data</a>. If I were anywhere near Seattle I would be in this class.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Ben Sima shared the first three things he installs on any new server.](https://bsima.me/clog/server-tools.html)

    I’m surprised he doesn’t have this in Ansible for easy deployment.
1. [Being agile and working smart are not the same thing](https://techbeacon.com/being-agile-working-smart-are-not-same-thing)

    (TechBeacon)
1. [DevOps is the course correction to a culture that is a byproduct of Taylorism and Sloanianism](https://opensource.com/open-organization/17/5/what-is-the-point-of-DevOps)

    Matt Micene reminds us all that DevOps isn’t some apparition that appeared out of nowhere. .
1. [Mike Simmons shares his monster list of Docker tips](https://dev.to/mikesimons/mikes-monster-list-of-docker-tips)

    . There are definitely some container nuggets in here.
1. [Anthony Shaw weighs the strengths and weaknesses of Ansible vs. Salt vs. StackStorm](https://medium.com/@anthonypjshaw/ansible-v-s-salt-saltstack-v-s-stackstorm-3d8f57149368)

    . Ansible fairs very well showing that it is standing the test of time.
1. [Heap shared a solid article on Postgres performance tuning](http://heap.engineering/basic-performance-analysis-saved-us-millions/)

    (despite the click-bait title).
1. [Kubernetes put out a piece on monitoring](http://blog.kubernetes.io/2017/05/kubernetes-monitoring-guide.html)

    and how k8s shifts the mindset required to effectively monitor your infrastructure.
### Department of Happy Clouds

1. [Cloud Spanner](https://cloud.google.com/spanner/)

    , the world’s first horizontally-scalable and strongly-consistent relational database service, is now generally available for your mission-critical OLTP applications. This is great news for folks needing five-nines of availability, strong consistency, and low latency.
### Department of Refreshment and Refurbishment

1. [Vagrant 1.9.5 is out](https://www.hashicorp.com/blog/vagrant-1-9-5/)

    ! This release features Docker Compose support among other things.
1. [Apache Allura](http://allura.apache.org/)

    Ever want a completely open source software forge?  offers a git repo, wiki, and ticketing system with 2FA.
1. [HTTP/2 coming to EL7 (finally)](https://ma.ttias.be/centos-7-4-ship-tls-1-2-alpn/)

     OpenSSL has been rebased in RHEL 7 to OpenSSL 1.0.2k. This will allow for Application-Layer Protocol Negotiation (ALPN) which is necessary for HTTP/2.
1. [Stack Overflow has a new iOS and Android app](https://stackoverflow.blog/2017/05/16/stack-overflow-official-app-launches-ios-android/)

    . I’m not sure how useful this will be to me; Google has always been my conduit to Stack Overflow. But, if you slap it on your phone and use it more power to you.
1. [Introducing Badger: A fast key-value store written natively in Go](https://open.dgraph.io/post/badger/)

    (Dgraph)
1. [Kubernetes 1.6.4 is out](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG.md/#v164)

    as the k8s squad pushes towards 1.7.
### Department of Data Defense

1. [Unmanaged SSH keys are truly awful](http://www.csoonline.com/article/3196974/data-protection/unmanaged-orphaned-ssh-keys-remain-a-serious-enterprise-risks.html)

    . How are you going to handle employees coming and going (sometimes abruptly) without central management of SSH credentials? It’s surprisingly a bigger problem than you’d think.
1. [A developer at Panic shared their experience when, through a random series of events, got pwned](https://panic.com/blog/stolen-source-code/)

    What happens where you get hacked? . It is an insightful postmortem of personal accountability.
1. [The need for urgent collective action to keep people safe online: Lessons from last week’s cyberattack](https://blogs.microsoft.com/on-the-issues/2017/05/14/need-urgent-collective-action-keep-people-safe-online-lessons-last-weeks-cyberattack/#sm.0000llg81y90hdbs115975pt98jep)

    (Microsoft)
### Department of Assemblage Obtainment

1. [SolarWindsScoutThis acquisition](http://www.solarwinds.com/)

    My employer, , has acquired . Scout was in the server monitoring and APM space.  dovetails nicely into both the traditional SolarWinds monitoring space and its cloud business.
### Not DevOps But Still Cool

1. [Kelsey Hightower shared how he goes about learning new things](https://twitter.com/kelseyhightower/status/864857551942307840)

    . Reading RFCs is truly insightful.
1. [Joe Nelson broke down what the quintessential Linux workstation would look like today](https://begriffs.com/posts/2017-05-17-linux-workstation-guide.html#core)

    . Asus is still the manufacturer of choice for the motherboard.
1. [Calling Bullshit in the Age of Big Data](http://callingbullshit.org/syllabus.html)

    The University of Washington has a new class, . If I were anywhere near Seattle I would be in this class.
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-023.md) ------------- [ Next >> ](devopsweekly-025.md)