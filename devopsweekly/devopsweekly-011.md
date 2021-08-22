## [DevOps'ish 011](https://devopsish.com/011) - Sun Feb 19, 2017

Alexander Graham Bell Birthplace, Edinburgh, Scotland, United Kingdom

My apologies for the delay in this week’s DevOps’ish newsletter. When I opened the ole MacBook last night to hammer it out it dawned on me how exhausted I was. Instead of this newsletter I spent time with family and went to bed early. Sorry, not sorry. It was an eventful week in the world of <a href="https://devopsish.com/"><strong>DevOps</strong></a> though.

<a href="https://cloud.google.com/spanner/">Google Spanner</a> was introduced as, “The first horizontally scalable, globally consistent, relational database service.” After some initial reading on <a href="https://www.wired.com/2017/02/spanner-google-database-harnessed-time-now-open-everyone/">Wired about Spanner</a> my interest has piqued. There is a white paper in regards to <a href="https://cloud.google.com/spanner/docs/whitepapers/SpannerAndCap.pdf">CAP Theorem and Spanner</a> that is like reading a sci-fi novel in which laws of physics are bent to humanity’s will. But the true gem in Spanner is <a href="http://static.googleusercontent.com/media/research.google.com/en/us/archive/spanner-osdi2012.pdf">Google’s handling of time</a> (dubbed TrueTime) to create consistency.

Software licensing is becoming an incredibly important factor in DevOps. You might have cobbled together some tools and some code to solve a problem but can your employer legally use it? <a href="https://lwn.net/SubscriberLink/714524/0c495ef932cd49a8/">Tom Callaway is the person behind making Fedora legally legit</a>. The story about his work in the Fedora teams is a must read.

<a href="https://www.openssl.org/news/vulnerabilities.html#2017-3733">OpenSSL 1.1.0 received an update this week to mitigate CVE-2017–3733</a>. Not a show stopper for most people. Has the Heartbleed stress finally gone its course surrounding OpenSSL updates?

<a href="https://medium.com/making-instapaper/instapaper-outage-cause-recovery-3c32a7e9cc5f#.tz9u7z33m">Instapaper published their recent outage postmortem this week</a>. It reminded me of an ever present problem in DevOps; <strong>who owns state</strong>? As much as we want to abstract state further away from development, we have to maintain state in a very hands on way or we’re doomed. Instapaper offloaded their state to Amazon RDS. An unknown and undocumented (by Amazon) 2TB limit in ext3 doomed the Instapaper RDS instance.

I am a sucker for a good cheatsheet or command reference. <a href="http://linoxide.com/linux-how-to/docker-commands-cheat-sheet/">Bobbin Zachariah created a good Docker cheatsheet</a> (albeit aesthetically awful) that is worth keeping around.

<a href="https://www.theregister.co.uk/2017/02/15/think_different_shut_up_and_work_harder_says_linus_torvalds/">Linus Torvalds, creator of the Linux kernel, made some headlines this week</a> at Open Source Leadership Summit. But, the real story is not Linus’ gruff attitude it’s the fact that he acknowledged that process (not code) is the real problem in large projects.

<a href="https://medium.com/@anne_e_currie/kubernetes-the-destroyer-of-worlds-4615dec3027b#.jv62ksuu8">Anne Currie reminds us that strictly Operations jobs are an outmoded profession</a>. Automation will eventually eliminate tier 2 type positions and those folks will need to either become DevOps thinking people or change professions entirely. Technology is evolving around us, what are you doing to keep up?

<a href="https://krebsonsecurity.com/2017/02/men-who-sent-swat-team-heroin-to-my-home-sentenced/">Brian Krebs had some news on folks that have messed with him personally.</a> Justice is being served and some of the attackers are apologetic to Krebs.

This week’s one-liner is pretty simple but one that I actually used this week. We needed to run <a href="https://www.percona.com/software/database-tools/percona-toolkit">pt-query-digest</a> against a week’s worth of MySQL slow query logs. The problem was the logs were rotated daily and I needed a single file for input.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Google SpannerWired about SpannerCAP Theorem and SpannerGoogle’s handling of time](https://cloud.google.com/spanner/)

    was introduced as, “The first horizontally scalable, globally consistent, relational database service.” After some initial reading on  my interest has piqued. There is a white paper in regards to  that is like reading a sci-fi novel in which laws of physics are bent to humanity’s will. But the true gem in Spanner is  (dubbed TrueTime) to create consistency.
1. [Tom Callaway is the person behind making Fedora legally legit](https://lwn.net/SubscriberLink/714524/0c495ef932cd49a8/)

    Software licensing is becoming an incredibly important factor in DevOps. You might have cobbled together some tools and some code to solve a problem but can your employer legally use it? . The story about his work in the Fedora teams is a must read.
### Department of Data Defense

1. [OpenSSL 1.1.0 received an update this week to mitigate CVE-2017–3733](https://www.openssl.org/news/vulnerabilities.html#2017-3733)

    . Not a show stopper for most people. Has the Heartbleed stress finally gone its course surrounding OpenSSL updates?
1. [Instapaper published their recent outage postmortem this week](https://medium.com/making-instapaper/instapaper-outage-cause-recovery-3c32a7e9cc5f#.tz9u7z33m)

    . It reminded me of an ever present problem in DevOps; who owns state? As much as we want to abstract state further away from development, we have to maintain state in a very hands on way or we’re doomed. Instapaper offloaded their state to Amazon RDS. An unknown and undocumented (by Amazon) 2TB limit in ext3 doomed the Instapaper RDS instance.
1. [Bobbin Zachariah created a good Docker cheatsheet](http://linoxide.com/linux-how-to/docker-commands-cheat-sheet/)

    I am a sucker for a good cheatsheet or command reference.  (albeit aesthetically awful) that is worth keeping around.
### Department of Discussion

1. [Linus Torvalds, creator of the Linux kernel, made some headlines this week](https://www.theregister.co.uk/2017/02/15/think_different_shut_up_and_work_harder_says_linus_torvalds/)

    at Open Source Leadership Summit. But, the real story is not Linus’ gruff attitude it’s the fact that he acknowledged that process (not code) is the real problem in large projects.
### Department of Next Year’s Old Tech

1. [Anne Currie reminds us that strictly Operations jobs are an outmoded profession](https://medium.com/@anne_e_currie/kubernetes-the-destroyer-of-worlds-4615dec3027b#.jv62ksuu8)

    . Automation will eventually eliminate tier 2 type positions and those folks will need to either become DevOps thinking people or change professions entirely. Technology is evolving around us, what are you doing to keep up?
### Not DevOps But Still Cool

1. [Brian Krebs had some news on folks that have messed with him personally.](https://krebsonsecurity.com/2017/02/men-who-sent-swat-team-heroin-to-my-home-sentenced/)

    Justice is being served and some of the attackers are apologetic to Krebs.
### DevOps’ish One-Liner of the Week

1. [pt-query-digest](https://www.percona.com/software/database-tools/percona-toolkit)

    This week’s one-liner is pretty simple but one that I actually used this week. We needed to run  against a week’s worth of MySQL slow query logs. The problem was the logs were rotated daily and I needed a single file for input.
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](sreweekly-10.md) ------------- [ Next >> ](sreweekly-12.md)