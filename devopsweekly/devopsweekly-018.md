## [DevOps'ish 018](https://devopsish.com/018) - Sun Apr 9, 2017

I did so many different things this week! DevOps is really great and my role working between teams really brings me different challenges daily. Be it cultural or technical I am rarely doing the same thing hour-by-hour. Two things that did take a lot of my time this week were building new MySQL. database replicas (you know… state) and Mac server monitoring.

Surprisingly, I maintain a Mac server or two for package building for one of our products. There are literally no “out of the box” monitoring solutions for Mac OS X anymore. This is equally unsurprising and indicative of how a platform that is used by so many in DevOps is not used for production workloads. There is nothing wrong with this; it is just odd to me monitoring a Mac is such a cumbersome task (no it’s not just FreeBSD).

In the world of <a href="https://devopsish.com/"><strong>DevOps</strong></a> otherwise, it has been a pretty busy week. Let’s have a look!

<a href="https://medium.com/@snaveen">Naveen</a> shared a great list of <a href="https://medium.com/google-cloud/tools-that-make-my-life-easier-to-work-with-kubernetes-fce3801086c0">tools that make working with Kubernetes easier</a>. This is a kind of must have list. For whatever reason, I cannot get tab completion working with kubectl. ¯\_(ツ)_/¯

Speaking of <a href="https://kubernetes.io/">**Kubernetes</a>**, I have been tinkering with it more and more (who isn’t?). I really enjoy the tool and the concepts of k8s but, the documentation needs some serious help. I have volunteered to help the <a href="https://kubernetes.slack.com/archives/C1J0BPD2M/p1491607569009282">sig-docs</a> team to clean this up. If you have some cycles please consider helping out, the team is VERY receptive and accepting. For the record, <a href="https://twitter.com/jbeda/status/850377020538200064/photo/1">Joe Beda pushed me to do something</a>.

<a href="https://medium.com/devops-journeys/devops-has-gone-mainstream-is-that-a-good-thing-1d698baaa12b">DevOps has gone mainstream. Is that a good thing?</a>

Here is something really off the wall, <a href="http://code.kryo.se/iodine/">iodine: IPv4 over DNS tunnel</a>. This is something I would have used in various military networks of yesteryear to prove security is not as good as they thought it was. Enforce ports AND protocols, folks.

Jenkins announced their long overdue improvement to their UI, <a href="https://jenkins.io/blog/2017/04/05/welcome-to-blue-ocean/">Blue Ocean, went 1.0</a>. I am excited to get this out there and start tinkering with it. I feel like the Jenkins UI is daunting to new users.

A <a href="https://xenbits.xen.org/xsa/advisory-212.html">gnarly Xen vulnerability</a> was announced this week. Luckily this vulnerability only affects x86–64 guests. <a href="https://googleprojectzero.blogspot.com/2017/04/pandavirtualization-exploiting-xen.html">Google’s Project Zero has a detailed article detailing the vulnerability</a> if you are interested.

Security is an important aspect of every information worker’s daily routine. To ignore this is to end up a headline in a newspaper regarding your breach. I think it might be past time for the DevOps community to be concerned about <a href="https://blog.acolyer.org/2017/04/03/a-study-of-security-vulnerabilities-on-docker-hub/">container security</a>.

Obviously, VMware vCloud Air was a horrible idea. VMware Cloud on AWS put vCloud Air on the back burner in a big way. <a href="http://www.vmware.com/company/news/releases/vmw-newsfeed.OVH-Announces-Intent-to-Acquire-VMware-vCloud-Air-Business.2153983.html">The bastardized business unit is to be sold to OVH</a>.

Werner Vogels announced this week an <a href="http://www.allthingsdistributed.com/2017/04/aws-announces-eu-stockholm-region.html">AWS region is coming to Stockholm, Sweden</a>. “AWS EU (Stockholm) Region will have three Availability Zones and will be ready for customers to use in 2018.”

Don’t read “<a href="http://brson.github.io/2017/04/05/minimally-nice-maintainer">The Minimally-nice Open Source Software Maintainer</a>” as a Bible but it does merit a bookmark for when your open source works make you see red.

In a surprise to no one, <a href="http://www.cnbc.com/2017/04/06/san-francisco-cost-of-living-pricing-out-tech-companies-workers.html">San Francisco has gotten so expensive, some tech companies can’t convince employees to move there</a>. It would take an absurd salary to get me to move to SF at this point.

I wrote some Go and wrote about it! <a href="https://opensource.com/article/17/4/testing-certificate-chains-34-line-go-program">Golang to the rescue: Saving DevOps from TLS turmoil</a>.

Another tweet from Kelsey Hightower this week. I am particularly interested in seeing how this pans out:

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Naveentools that make working with Kubernetes easier](https://medium.com/@snaveen)

    shared a great list of . This is a kind of must have list. For whatever reason, I cannot get tab completion working with kubectl. ¯\_(ツ)_/¯
1. [**Kubernetessig-docsJoe Beda pushed me to do something](https://kubernetes.io/)

    Speaking of **, I have been tinkering with it more and more (who isn’t?). I really enjoy the tool and the concepts of k8s but, the documentation needs some serious help. I have volunteered to help the  team to clean this up. If you have some cycles please consider helping out, the team is VERY receptive and accepting. For the record, .
1. [DevOps has gone mainstream. Is that a good thing?](https://medium.com/devops-journeys/devops-has-gone-mainstream-is-that-a-good-thing-1d698baaa12b)

    
### Department of Refreshment and Refurbishment

1. [iodine: IPv4 over DNS tunnel](http://code.kryo.se/iodine/)

    Here is something really off the wall, . This is something I would have used in various military networks of yesteryear to prove security is not as good as they thought it was. Enforce ports AND protocols, folks.
1. [Blue Ocean, went 1.0](https://jenkins.io/blog/2017/04/05/welcome-to-blue-ocean/)

    Jenkins announced their long overdue improvement to their UI, . I am excited to get this out there and start tinkering with it. I feel like the Jenkins UI is daunting to new users.
### Department of Data Defense

1. [gnarly Xen vulnerabilityGoogle’s Project Zero has a detailed article detailing the vulnerability](https://xenbits.xen.org/xsa/advisory-212.html)

    A  was announced this week. Luckily this vulnerability only affects x86–64 guests.  if you are interested.
1. [container security](https://blog.acolyer.org/2017/04/03/a-study-of-security-vulnerabilities-on-docker-hub/)

    Security is an important aspect of every information worker’s daily routine. To ignore this is to end up a headline in a newspaper regarding your breach. I think it might be past time for the DevOps community to be concerned about .
### Department of Assemblage Obtainment

1. [The bastardized business unit is to be sold to OVH](http://www.vmware.com/company/news/releases/vmw-newsfeed.OVH-Announces-Intent-to-Acquire-VMware-vCloud-Air-Business.2153983.html)

    Obviously, VMware vCloud Air was a horrible idea. VMware Cloud on AWS put vCloud Air on the back burner in a big way. .
### Department of Happy Little Clouds

1. [AWS region is coming to Stockholm, Sweden](http://www.allthingsdistributed.com/2017/04/aws-announces-eu-stockholm-region.html)

    Werner Vogels announced this week an . “AWS EU (Stockholm) Region will have three Availability Zones and will be ready for customers to use in 2018.”
### Department of Sane Workplaces

1. [The Minimally-nice Open Source Software Maintainer](http://brson.github.io/2017/04/05/minimally-nice-maintainer)

    Don’t read “” as a Bible but it does merit a bookmark for when your open source works make you see red.
1. [San Francisco has gotten so expensive, some tech companies can’t convince employees to move there](http://www.cnbc.com/2017/04/06/san-francisco-cost-of-living-pricing-out-tech-companies-workers.html)

    In a surprise to no one, . It would take an absurd salary to get me to move to SF at this point.
### Department of Interior

1. [Golang to the rescue: Saving DevOps from TLS turmoil](https://opensource.com/article/17/4/testing-certificate-chains-34-line-go-program)

    I wrote some Go and wrote about it! .

### [ << Prev ](devopsweekly-017.md) ------------- [ Next >> ](devopsweekly-019.md)