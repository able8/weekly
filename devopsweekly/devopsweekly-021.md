## [DevOps'ish 021](https://devopsish.com/021) - Sun Apr 30, 2017

<a href="http://www.cnn.com/2017/04/25/us/north-carolina-flooding/">North Carolina, and particularly the Triangle area, has had an enormous amount of rain this week</a>. The tiny creek behind our house was a rapid early in the week. According to the USGS, <a href="https://waterdata.usgs.gov/nc/nwis/uv?cb_00065=on&amp;format=gif_stats&amp;site_no=02087182&amp;period=&amp;begin_date=2017-04-22&amp;end_date=2017-04-29">Falls Lake is almost ten feet above it’s average levels</a>. There was <a href="https://governor.nc.gov/news/governor-cooper-urges-caution-flooding-continues-threaten-eastern-nc-communities">one fatality as a result of someone driving around a barricade</a> onto a flooded roadway. Common sense is not common. Please stay safe out there, folks.

Meanwhile, back on the <a href="https://devopsish.com/"><strong>DevOps</strong></a> ranch, it was a VERY busy week. I’m on-call for two products this week, working on three projects, writing an article inspired by my VP, submitting CFPs, and trying to keep up with the news. With so much going on let’s get to the <a href="https://devopsish.com/">DevOps’ish</a>!

<a href="http://blog.alexellis.io/boot-linuxkit-in-10-mins/">Boot an OpenSSH server in 10 mins with LinuxKit</a>: Hands-on guide to use Docker’s LinuxKit to build, run and connect to a bootable Linux system image with OpenSSH.

<a href="https://jvns.ca/blog/2017/04/23/the-fish-shell-is-awesome/">Julia Evans wrote about her favorite shell this week</a>. I just setup <a href="https://fishshell.com/">fish shell</a> after re-reading Julia’s article. Install, add to /etc/shells, and you’re off and running. Out of the box it’s REALLY awesome.

<a href="https://blog.jessfraz.com/post/two-objects-not-namespaced-linux-kernel/">I had no idea the most important dimension of the human race, time, isn’t namespaced by the Linux kernel</a>.

<a href="https://cloudplatform.googleblog.com/2017/04/going-multi-cloud-with-Google-Cloud-Endpoints-and-AWS-Lambda.html">Going multi-cloud with Google Cloud Endpoints and AWS Lambda</a>: This is what cloud <em>should</em> be! You want to use them for some things? Fine. You want to use us for other things? Great! You want to use us both to make your product more awesome, resilient, and cost-effective? <strong>FANTASTIC</strong>!

<img src="https://d33wubrfki0l68.cloudfront.net/b93ceddf7c17955c0747f94feb2678b3c4b93e0b/b2d48/021/kubernetes-simplified-architecture.png" alt="Brian Grant&#39;s Kubernetes Simplified Architectural Summary"/><br/><em>Brian Grant’s Kubernetes Simplified Architectural Summary</em>

Brian Grant has laid out the <a href="https://groups.google.com/forum/#!topic/kubernetes-dev/yXMjoMAZRN4">Kubernetes Architecture in a way that is very consumable</a> yet sufficiently detailed.

Speaking of Kubernetes, here is a <a href="https://sematext.com/kubernetes/cheatsheet/">handy cheatsheet for Kubernetes</a> (available as PDF and Markdown).

<a href="https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back">Understand Go pointers in less than 800 words or your money back</a>

Go is a great programming language (especially for <a href="https://devopsish.com/">DevOps</a>). One feature I feel is missing is its ability to update itself or having some built-in process for doing so. <a href="https://groups.google.com/forum/#!topic/golang-devexp/-ejglEViXN0">I think we should be past the days of languages not at least having an integrated update process</a>. This is not to say that it should auto-update; <em>I am COMPLETELY against that</em>. But, the fact I have to use something outside of the golang environment to update go is a little awkward to me in this day and age. <a href="https://github.com/udhos/update-golang">update-golang</a> looks to simplify the process of updating Go.

<a href="http://www.sfchronicle.com/business/article/Suicide-of-an-Uber-engineer-widow-blames-job-11095807.php">An accomplished engineer made the horrible choice of taking a job at Uber and it likely cost him his life</a>. Can we all just uninstall Uber and delete our accounts now?

<a href="http://www.theverge.com/2017/4/26/15439622/fcc-net-neutrality-internet-freedom-isp-ajit-pai">The fight FOR <strong>net neutrality</strong> is on like Donkey Kong</a>. I don’t care which side of the debate you are on, <a href="https://www.fcc.gov/ecfs/search/filings?proceedings_name=17-108">**please share your thoughts with the FCC</a>**.

The <a href="http://www.kaporcenter.org/tech-leavers/"><em>Tech Leavers Study</em></a> is a first-of-its-kind national study examining why people voluntarily left their jobs in tech. The Kapor Center for Social Impact and Harris Poll surveyed a representative sample of more than 2,000 U.S. adults who have left a job in a technology-related industry or function within the last three years.

<a href="https://blog.hipchat.com/2017/04/24/hipchat-security-notice/">HipChat was breached last weekend</a>. According to Atlassian’s Chief Security Officer, Ganesh Krishnan, the incident involved a vulnerability in a third-party library. Patch your dependencies, folks!

Still trying to figure out CAA records? Check out <a href="https://isc.sans.edu/forums/diary/CAA+Records+and+Certificate+Issuance/22342/">J. Edward Durrett’s SANS ISC diary about CAA Records</a>.

<a href="http://www.circleid.com/posts/20170426_the_sysadmins_guide_to_securing_your_saas_apps/">The Sysadmin’s Guide to Securing Your SaaS Apps</a>

<a href="http://mailman.nginx.org/pipermail/nginx-announce/2017/000195.html">nginx 1.13.0 was released on Tuesday</a> and summarily broke my website due to an errant line break in a configuration I rarely touch. It includes TLSv1.3 support which makes me stupid happy.

<a href="https://twitter.com/b0rk/status/858077003244855298">Julia Evans made tcpdump zine</a> and it is absolutely fantastic. She is selling <a href="https://gumroad.com/l/LcKLx">early releases for $10</a>. Support her efforts to make us all smarter!

<a href="https://opensource.com/article/17/4/haters-gonna-hate">Jason van Gumster reminds us that haters gonna hate</a>. Jason also provides some methods of dealing with those that woke up and drank a tall glass of Haterade.

How do you look at this graphic from Docker and not think Docker is killing it’s branding for some reason? <a href="https://opensource.com/open-organization/17/4/how-branding-decisions-open">Docker’s recent rebranding efforts offer a lesson in public decision-making</a>.

<a href="https://www.debian.org/News/2017/20170425">Debian is shutting down its public FTP services</a>. That’s about a decade overdue. When I worked at <a href="http://www.mcclatchy.com/">McClatchy</a> we had a hell of a time maintaining FTP services. Photographers, journalists, and several other people depended upon FTP’s simplicity. But, its lack of security gave us all kinds of headaches. It is way past time for FTP to go the way of <a href="https://en.wikipedia.org/wiki/Gopher_(protocol)">Gopher</a>.

<a href="https://www.amazon.jobs/principles?&amp;tag=rnwap-20">Amazon has shared some Leadership Principles</a> that I agree with. I likely wouldn’t work at Amazon HQ but it’s encouraging to see this.

I got asked a very interesting question today by my VP, “<a href="https://twitter.com/ChrisShort/status/857356577417949184">How do you prep a good sysadmins to be a DevOps engineer in six months?</a>” E-mail or <a href="https://twitter.com/devopsish">tweet DevOps’ish</a> to give your thoughts.

<a href="http://www.businessinsider.com/wikipedia-cofounder-jimmy-wales-wikitribune-community-powered-news-2017-4">Wikipedia cofounder Jimmy Wales is launching a community-powered news site</a>

Sarah Novotny coins, “mean time to dopamine” while discussing <a href="https://opensource.com/article/17/4/podcast-kubernetes-sarah-novotny">how Kubernetes is making contributing easy</a> with Gordon Haff. I have to admit, the Kubernetes community is pretty great.

<a href="https://gist.github.com/bketelsen/3ebe4979e9ee563ed1a8c81435ddf767">Smoked Fried Chicken</a>, ‘nuff said.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Boot an OpenSSH server in 10 mins with LinuxKit](http://blog.alexellis.io/boot-linuxkit-in-10-mins/)

     Hands-on guide to use Docker’s LinuxKit to build, run and connect to a bootable Linux system image with OpenSSH.
1. [Julia Evans wrote about her favorite shell this weekfish shell](https://jvns.ca/blog/2017/04/23/the-fish-shell-is-awesome/)

    . I just setup  after re-reading Julia’s article. Install, add to /etc/shells, and you’re off and running. Out of the box it’s REALLY awesome.
1. [I had no idea the most important dimension of the human race, time, isn’t namespaced by the Linux kernel](https://blog.jessfraz.com/post/two-objects-not-namespaced-linux-kernel/)

    .
1. [Going multi-cloud with Google Cloud Endpoints and AWS Lambda](https://cloudplatform.googleblog.com/2017/04/going-multi-cloud-with-Google-Cloud-Endpoints-and-AWS-Lambda.html)

     This is what cloud should be! You want to use them for some things? Fine. You want to use us for other things? Great! You want to use us both to make your product more awesome, resilient, and cost-effective? FANTASTIC!
1. []()

    Brian Grant’s Kubernetes Simplified Architectural Summary
1. [Kubernetes Architecture in a way that is very consumable](https://groups.google.com/forum/#!topic/kubernetes-dev/yXMjoMAZRN4)

    Brian Grant has laid out the  yet sufficiently detailed.
1. [handy cheatsheet for Kubernetes](https://sematext.com/kubernetes/cheatsheet/)

    Speaking of Kubernetes, here is a  (available as PDF and Markdown).
1. [Understand Go pointers in less than 800 words or your money back](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back)

    
1. [DevOpsI think we should be past the days of languages not at least having an integrated update processupdate-golang](https://devopsish.com/)

    Go is a great programming language (especially for ). One feature I feel is missing is its ability to update itself or having some built-in process for doing so. . This is not to say that it should auto-update; I am COMPLETELY against that. But, the fact I have to use something outside of the golang environment to update go is a little awkward to me in this day and age.  looks to simplify the process of updating Go.
### Department of Woke Workplaces

1. [An accomplished engineer made the horrible choice of taking a job at Uber and it likely cost him his life](http://www.sfchronicle.com/business/article/Suicide-of-an-Uber-engineer-widow-blames-job-11095807.php)

    . Can we all just uninstall Uber and delete our accounts now?
1. [The fight FOR net neutrality is on like Donkey Kong**please share your thoughts with the FCC](http://www.theverge.com/2017/4/26/15439622/fcc-net-neutrality-internet-freedom-isp-ajit-pai)

    . I don’t care which side of the debate you are on, **.
1. [Tech Leavers Study](http://www.kaporcenter.org/tech-leavers/)

    The  is a first-of-its-kind national study examining why people voluntarily left their jobs in tech. The Kapor Center for Social Impact and Harris Poll surveyed a representative sample of more than 2,000 U.S. adults who have left a job in a technology-related industry or function within the last three years.
### Department of Data Defense

1. [HipChat was breached last weekend](https://blog.hipchat.com/2017/04/24/hipchat-security-notice/)

    . According to Atlassian’s Chief Security Officer, Ganesh Krishnan, the incident involved a vulnerability in a third-party library. Patch your dependencies, folks!
1. [J. Edward Durrett’s SANS ISC diary about CAA Records](https://isc.sans.edu/forums/diary/CAA+Records+and+Certificate+Issuance/22342/)

    Still trying to figure out CAA records? Check out .
1. [The Sysadmin’s Guide to Securing Your SaaS Apps](http://www.circleid.com/posts/20170426_the_sysadmins_guide_to_securing_your_saas_apps/)

    
### Department of Refreshment and Refurbishment

1. [nginx 1.13.0 was released on Tuesday](http://mailman.nginx.org/pipermail/nginx-announce/2017/000195.html)

    and summarily broke my website due to an errant line break in a configuration I rarely touch. It includes TLSv1.3 support which makes me stupid happy.
1. [Julia Evans made tcpdump zineearly releases for $10](https://twitter.com/b0rk/status/858077003244855298)

    and it is absolutely fantastic. She is selling . Support her efforts to make us all smarter!
### Department of Discussion

1. [Jason van Gumster reminds us that haters gonna hate](https://opensource.com/article/17/4/haters-gonna-hate)

    . Jason also provides some methods of dealing with those that woke up and drank a tall glass of Haterade.
### Department of Next Year’s Old Tech

1. [Docker’s recent rebranding efforts offer a lesson in public decision-making](https://opensource.com/open-organization/17/4/how-branding-decisions-open)

    How do you look at this graphic from Docker and not think Docker is killing it’s branding for some reason? .
1. [Debian is shutting down its public FTP servicesMcClatchyGopher](https://www.debian.org/News/2017/20170425)

    . That’s about a decade overdue. When I worked at  we had a hell of a time maintaining FTP services. Photographers, journalists, and several other people depended upon FTP’s simplicity. But, its lack of security gave us all kinds of headaches. It is way past time for FTP to go the way of .
### Department of Sane Workplaces

1. [Amazon has shared some Leadership Principles](https://www.amazon.jobs/principles?&tag=rnwap-20)

    that I agree with. I likely wouldn’t work at Amazon HQ but it’s encouraging to see this.
1. [How do you prep a good sysadmins to be a DevOps engineer in six months?tweet DevOps’ish](https://twitter.com/ChrisShort/status/857356577417949184)

    I got asked a very interesting question today by my VP, “” E-mail or  to give your thoughts.
### Not DevOps But Still Cool

1. [Wikipedia cofounder Jimmy Wales is launching a community-powered news site](http://www.businessinsider.com/wikipedia-cofounder-jimmy-wales-wikitribune-community-powered-news-2017-4)

    
1. [how Kubernetes is making contributing easy](https://opensource.com/article/17/4/podcast-kubernetes-sarah-novotny)

    Sarah Novotny coins, “mean time to dopamine” while discussing  with Gordon Haff. I have to admit, the Kubernetes community is pretty great.
1. [Smoked Fried Chicken](https://gist.github.com/bketelsen/3ebe4979e9ee563ed1a8c81435ddf767)

    , ‘nuff said.
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](sreweekly-20.md) ------------- [ Next >> ](sreweekly-22.md)