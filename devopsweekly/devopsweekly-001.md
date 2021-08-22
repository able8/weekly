## [DevOps'ish 001](https://devopsish.com/001) - Sun Dec 11, 2016

Welcome to this week’s edition of DevOps’ish where we cover Dev, Ops, and all the ish in between.

I love the idea of <a href="https://blog.gladwell.me/html-microservices.html">dumping XML and JSON for HTML 5</a>. Despite having to block people trying to pick fights on Twitter, this idea has been pretty well received. I would love to see the frustrations of JSON cast aside for something less syntactically challenging and more human readable.

<a href="https://twitter.com/ryan_sb/status/806884324805115906">Ryan Scott Brown might have beat Amazon to the punch</a> but <a href="https://aws.amazon.com/blogs/aws/now-open-aws-canada-central-region/">AWS has launched a region in Canada</a>! 🇨🇦 🇨🇦 🇨🇦
The general consensus is if a company that is doing business in Canada they are likely going to be diving head first into this AWS region. <a href="http://www.servercloudcanada.com/2015/09/canadian-privacy-laws-canadian-cloud-primer-canadian-businesses/">Canadian privacy laws</a> make sure a lot of Canadian PII does not leave the county.

Speaking of Amazon and privacy… <a href="http://shivankaul.com/blog/2016/12/07/clean-your-desk-yet-another-amazon-interview-experience.html">A developer interviewing at Amazon had a rather unusual experience</a> taking a proctored test. He ended up terminating the interview early and then as the proctor was unable to clean things up on the dev’s machine the dev terminated the termination. “The normalization of privacy violation has never felt more real.”

There is another example where <a href="https://rajk.me/amazon-interview-experience/">the dev being interviewed rejected out right to a proctored test</a> over privacy concerns. Instead, the dev offered to do a face-to-face. Amazon has yet to respond.

LinkedIn found an issue with longer garbage collection times from their Java app running in cgroups. “<a href="https://engineering.linkedin.com/blog/2016/11/application-pauses-when-running-jvm-inside-linux-control-groups">We found that Java applications can have more and longer application pauses when using CFS (Completely Fair Scheduler) in conjunction with CFS Bandwidth Control quotas.</a>“

<a href="https://twitter.com/alicegoldfuss">Alice Goldfuss</a> sparked off a fire storm that she likely was not given enough credit for starting with the wonderfully written <a href="http://sysadvent.blogspot.com/">SysAdvent</a> piece titled, “<a href="http://sysadvent.blogspot.com/2016/12/day-6-no-more-on-call-martyrs.html">Day 6 — No More On-Call Martyrs</a>.” Reading <a href="https://lobste.rs/s/djmmbg/no_more_on_call_martyrs">some</a> <a href="https://twitter.com/alicegoldfuss/status/806328359768690688">of</a> <a href="https://twitter.com/alicegoldfuss/status/806329196683415561">the</a> <a href="http://naildrivin5.com/blog/2016/12/07/on-call.html">discussion</a> with <a href="https://www.youtube.com/watch?v=gFtb3EtjEic">Andy Williams’ “It’s The Most Wonderful Time Of The Year”</a> playing is kinda comical (in a sick sort of way). 🎄🎅🎄🎅

<a href="https://blog.travis-ci.com/2016-12-06-the-crons-are-here">Travis CI announced a new cron feature</a>! Test your stuff all day!

<a href="https://nodejs.org/en/blog/community/update-v8-5.4/">Node.js v7 has updated V8 to 5.4</a> and Node.js is trying to make itself <a href="https://nodejs.org/en/blog/weekly-updates/weekly-update.2016-12-02/">VM-neutral</a> (which is great news).

I wrote an article for DZone discussing <a href="https://dzone.com/articles/empathy-the-emerging-art-in-devops-1"><strong>empathy</strong> that is so desperately needed in DevOps</a>.

<a href="http://twitter.com/denoncourt">Don Denoncourt</a> over at <a href="http://corgibytes.com/">Corgibytes</a> has an interesting piece on what to do about a problem we all have… <a href="http://corgibytes.com/blog/2016/12/06/getting-old-er-in-tech/">Getting older</a>. Oh and Corgibytes has the best logo.

<a href="https://blog.docker.com/2016/12/docker-acquires-infinit/">Infinit was acquired by Docker</a> giving Docker the potential for distributed storage in their product offerings (which is long overdue, in my opinion). More on <a href="https://www.crunchbase.com/organization/infinit">Infinit from Crunchbase</a>.

I have had a few discussions this week about how configuration management platforms like Ansible, Puppet, Chef, etc. are going to start fading into the background as containers and Kubernetes gain popularity. What do you think?

Anil Dash and Fog Creek launched a new product called <a href="https://gomix.com/community/">Gomix</a>. Show it off to your relatives during the holidays and spark some intrigue in coding if you can.

Red Hat gave Duke University a $180,000 gift that will, “<a href="https://today.duke.edu/2016/11/red-hat-gift-spurs-open-source-ideas-duke">assist students who have project ideas that use open-source software to advance humanitarian issues</a>.” Two seniors are building a specialized prosthetic arm for a cake decorator using open source software.

I needed to figure out how big some Amazon S3 buckets were this week and found there was not an obvious way but at least <a href="http://serverfault.com/questions/84815/how-can-i-get-the-size-of-an-amazon-s3-bucket">serverfault</a> netted this beauty:

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [dumping XML and JSON for HTML 5](https://blog.gladwell.me/html-microservices.html)

    I love the idea of . Despite having to block people trying to pick fights on Twitter, this idea has been pretty well received. I would love to see the frustrations of JSON cast aside for something less syntactically challenging and more human readable.
### Department of Happy Little Clouds

1. [Ryan Scott Brown might have beat Amazon to the punchAWS has launched a region in CanadaCanadian privacy laws](https://twitter.com/ryan_sb/status/806884324805115906)

    but ! 🇨🇦 🇨🇦 🇨🇦
The general consensus is if a company that is doing business in Canada they are likely going to be diving head first into this AWS region.  make sure a lot of Canadian PII does not leave the county.
1. [A developer interviewing at Amazon had a rather unusual experience](http://shivankaul.com/blog/2016/12/07/clean-your-desk-yet-another-amazon-interview-experience.html)

    Speaking of Amazon and privacy…  taking a proctored test. He ended up terminating the interview early and then as the proctor was unable to clean things up on the dev’s machine the dev terminated the termination. “The normalization of privacy violation has never felt more real.”
1. [the dev being interviewed rejected out right to a proctored test](https://rajk.me/amazon-interview-experience/)

    There is another example where  over privacy concerns. Instead, the dev offered to do a face-to-face. Amazon has yet to respond.
### Department of Dafuq

1. [We found that Java applications can have more and longer application pauses when using CFS (Completely Fair Scheduler) in conjunction with CFS Bandwidth Control quotas.](https://engineering.linkedin.com/blog/2016/11/application-pauses-when-running-jvm-inside-linux-control-groups)

    LinkedIn found an issue with longer garbage collection times from their Java app running in cgroups. ““
1. [Alice GoldfussSysAdventDay 6 — No More On-Call MartyrssomeofthediscussionAndy Williams’ “It’s The Most Wonderful Time Of The Year”](https://twitter.com/alicegoldfuss)

    sparked off a fire storm that she likely was not given enough credit for starting with the wonderfully written  piece titled, “.” Reading     with  playing is kinda comical (in a sick sort of way). 🎄🎅🎄🎅
### Department of Continuous Continuum

1. [Travis CI announced a new cron feature](https://blog.travis-ci.com/2016-12-06-the-crons-are-here)

    ! Test your stuff all day!
1. [Node.js v7 has updated V8 to 5.4VM-neutral](https://nodejs.org/en/blog/community/update-v8-5.4/)

    and Node.js is trying to make itself  (which is great news).
### Department of Sane Workplaces

1. [empathy that is so desperately needed in DevOps](https://dzone.com/articles/empathy-the-emerging-art-in-devops-1)

    I wrote an article for DZone discussing .
1. [Don DenoncourtCorgibytesGetting older](http://twitter.com/denoncourt)

    over at  has an interesting piece on what to do about a problem we all have… . Oh and Corgibytes has the best logo.
### Department of Assemblage Obtainment

1. [Infinit was acquired by DockerInfinit from Crunchbase](https://blog.docker.com/2016/12/docker-acquires-infinit/)

    giving Docker the potential for distributed storage in their product offerings (which is long overdue, in my opinion). More on .
### Department of Next Year’s Old Tech

1. []()

    I have had a few discussions this week about how configuration management platforms like Ansible, Puppet, Chef, etc. are going to start fading into the background as containers and Kubernetes gain popularity. What do you think?
### Not DevOps But Still Cool

1. [Gomix](https://gomix.com/community/)

    Anil Dash and Fog Creek launched a new product called . Show it off to your relatives during the holidays and spark some intrigue in coding if you can.
1. [assist students who have project ideas that use open-source software to advance humanitarian issues](https://today.duke.edu/2016/11/red-hat-gift-spurs-open-source-ideas-duke)

    Red Hat gave Duke University a $180,000 gift that will, “.” Two seniors are building a specialized prosthetic arm for a cake decorator using open source software.
### DevOps’ish One-Liner of the Week

1. [serverfault](http://serverfault.com/questions/84815/how-can-i-get-the-size-of-an-amazon-s3-bucket)

    I needed to figure out how big some Amazon S3 buckets were this week and found there was not an obvious way but at least  netted this beauty:
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-000.md) ------------- [ Next >> ](devopsweekly-002.md)