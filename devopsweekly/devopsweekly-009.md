## [DevOps'ish 009](https://devopsish.com/009) - Sun Feb 5, 2017

What a week in <a href="https://devopsish.com/"><strong>DevOps</strong></a> ! There is so much news to share this week it took me an hour to collect it all! I also went to the <a href="http://opensource101.com/">Open Source 101</a> conference this week to pick up some knowledge on communities. It was a very good conference. If the future of open source software was represented there it will definitely not be a white male dominated field.

Like a lot of sysadmin and DevOps teams, I bet you have some random cron jobs that are critical to your daily/weekly/monthly workflows. The problem with these jobs is that they always seem to never have a great place to live. <a href="https://medium.com/@marccuva">Marc Cuva</a> and the Serverless movement <a href="https://blog.readme.io/writing-a-cron-job-microservice-with-serverless-and-aws-lambda/">might have given us a solution to consider</a>.

<a href="https://learntemail.sam.today/blog/stop-disabling-selinux:-a-real-world-guide/">Make SELinux Enforcing Again!</a> We all do it… To some extent. Let’s try to stop disabling SELinux in 2017.

18F at GSA ran a small experiment in September and it turned out that <a href="https://18f.gsa.gov/2017/01/19/the-dark-standup/">working more hours did not net more productivity</a>. Confining work schedules and after hours phone checking actually made a team more effective. The crazy thing is this analysis came from a US government agency. Kudos to <a href="https://medium.com/@USDigitalService">The U.S. Digital Service</a>!

<a href="https://medium.com/@kendrickcoleman">Kendrick Coleman</a> took us on a <a href="https://blog.codedellemc.com/2017/02/02/deep-dive-docker-1-13-storage-plugins-propagated-mounts/?cmp=soc-cor-glbl-us-sprinklr-TWITTER--codeDellEMC-796974590">Docker 1.13 journey detailing storage plugins and propagated mounts</a>. It might not be production ready but it is definitely interesting.

The biggest story of the week was the <a href="https://docs.google.com/document/d/1GCK53YDcBWQveod9kfzW-VCxIABGiryG7_z_6jHdVik/pub">unintentional meltdown at GitLab and the amazing response to it</a>. The world watched as a <a href="https://medium.com/@gitlab">GitLab</a> team member <code>rm -Rvf a PostgreSQL</code> directory. GitLab’s response was the most open disaster recovery in modern history. I was VERY impressed by the level of openness and honesty; this is how postmortems should work. February 1st was declared “<a href="http://checkyourbackups.work/">Check your backups Day!</a>“

<a href="https://medium.com/@etsy">Etsy</a> described how their Pattern product <a href="https://codeascraft.com/2017/01/31/how-etsy-manages-https-and-ssl-certificates-for-custom-domains-on-pattern/">manages HTTPS and stores SSL certificates</a>. This is a fascinating write up of what, to many, would be an insurmountable effort.

Facebook has released <a href="https://github.com/facebookincubator/DelegatedRecovery/">delegated-account-recovery</a>. It is, “protocol that allows an application to delegate the capability to recover an account (e.g. in the event of a credential loss or compromise) to an account controlled by the same user at a third party service provider.”

Ever needed a cross-platform password management tool for your team? <a href="https://www.justwatch.com/blog/post/announcing-gopass/">gopass</a> was announced this week and solves several password management problems. I am hoping to introduce to one of my teams in Scotland this week to replace KeePass.

<a href="http://kompose.io/">Kompose</a> was announced this week. Keep an eye on this one because it’s a very simple way to convert a Docker Compose file to run on Kubernetes or OpenShift.

<a href="https://medium.com/@VisualStudio">Visual Studio</a> (yes, Microsoft’s Visual Studio) released GVFS this week. What is GVFS? It stands for <a href="https://blogs.msdn.microsoft.com/visualstudioalm/2017/02/03/announcing-gvfs-git-virtual-file-system/">Git Virtual File System</a>. Its allows dev teams to check out massive repos without having to download them in their entirety. GVFS only pulls down the files you actually need when you’re working on them. Pretty smart for those working on big code bases.

I have not finished reading the Site Reliability Engineering book cover-to-cover yet. But, I have treated it as a sort of reference as of late. The good news is, now we all can call upon Google’s <a href="https://landing.google.com/sre/book.html">Site Reliability Engineering</a> book because it is available completely for free.

<a href="https://opensource.com/">Opensource.com</a> announced their <a href="https://opensource.com/article/17/2/community-awards-2017">2017 Community Award Winners</a> this week. Congrats to all those selected!

I had to blow away a temp directory that was consuming an enormous amount of inodes this week. rm -rf was running far too slow and I needed a faster solution:

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Marc Cuvamight have given us a solution to consider](https://medium.com/@marccuva)

    Like a lot of sysadmin and DevOps teams, I bet you have some random cron jobs that are critical to your daily/weekly/monthly workflows. The problem with these jobs is that they always seem to never have a great place to live.  and the Serverless movement .
1. [Make SELinux Enforcing Again!](https://learntemail.sam.today/blog/stop-disabling-selinux:-a-real-world-guide/)

    We all do it… To some extent. Let’s try to stop disabling SELinux in 2017.
1. [working more hours did not net more productivityThe U.S. Digital Service](https://18f.gsa.gov/2017/01/19/the-dark-standup/)

    18F at GSA ran a small experiment in September and it turned out that . Confining work schedules and after hours phone checking actually made a team more effective. The crazy thing is this analysis came from a US government agency. Kudos to !
1. [Kendrick ColemanDocker 1.13 journey detailing storage plugins and propagated mounts](https://medium.com/@kendrickcoleman)

    took us on a . It might not be production ready but it is definitely interesting.
### Department of Data Defense

1. [unintentional meltdown at GitLab and the amazing response to itGitLabCheck your backups Day!](https://docs.google.com/document/d/1GCK53YDcBWQveod9kfzW-VCxIABGiryG7_z_6jHdVik/pub)

    The biggest story of the week was the . The world watched as a  team member rm -Rvf a PostgreSQL directory. GitLab’s response was the most open disaster recovery in modern history. I was VERY impressed by the level of openness and honesty; this is how postmortems should work. February 1st was declared ““
### Department of Refreshment and Refurbishment

1. [Etsymanages HTTPS and stores SSL certificates](https://medium.com/@etsy)

    described how their Pattern product . This is a fascinating write up of what, to many, would be an insurmountable effort.
1. [delegated-account-recovery](https://github.com/facebookincubator/DelegatedRecovery/)

    Facebook has released . It is, “protocol that allows an application to delegate the capability to recover an account (e.g. in the event of a credential loss or compromise) to an account controlled by the same user at a third party service provider.”
1. [gopass](https://www.justwatch.com/blog/post/announcing-gopass/)

    Ever needed a cross-platform password management tool for your team?  was announced this week and solves several password management problems. I am hoping to introduce to one of my teams in Scotland this week to replace KeePass.
1. [Kompose](http://kompose.io/)

    was announced this week. Keep an eye on this one because it’s a very simple way to convert a Docker Compose file to run on Kubernetes or OpenShift.
1. [Visual StudioGit Virtual File System](https://medium.com/@VisualStudio)

    (yes, Microsoft’s Visual Studio) released GVFS this week. What is GVFS? It stands for . Its allows dev teams to check out massive repos without having to download them in their entirety. GVFS only pulls down the files you actually need when you’re working on them. Pretty smart for those working on big code bases.
### Department of Sane Workplaces

1. [Site Reliability Engineering](https://landing.google.com/sre/book.html)

    I have not finished reading the Site Reliability Engineering book cover-to-cover yet. But, I have treated it as a sort of reference as of late. The good news is, now we all can call upon Google’s  book because it is available completely for free.
### Not DevOps But Still Cool

1. [Opensource.com2017 Community Award Winners](https://opensource.com/)

    announced their  this week. Congrats to all those selected!
### DevOps’ish One-Liner of the Week

1. []()

    I had to blow away a temp directory that was consuming an enormous amount of inodes this week. rm -rf was running far too slow and I needed a faster solution:
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](sreweekly-8.md) ------------- [ Next >> ](sreweekly-10.md)