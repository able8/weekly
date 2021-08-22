## [DevOps'ish 026](https://devopsish.com/026) - Sun Jun 4, 2017

It was a very busy week for me personally and in the world of <a href="https://devopsish.com/">DevOps</a>. The <a href="../025/">news I hinted at last week</a> is my public announcement <a href="https://chrisshort.net/leaving-north-carolina/">I’m leaving SolarWinds MSP, moving to Detroit, and taking the role of Manager, Infrastructure &amp; Operations at Bankrate</a>. I am beyond excited about the position and the work I will be doing. If you want the job I’m vacating, <a href="http://solarwinds.jobs/durham-nc-nc/senior-aws-systems-engineer/C28A5E389CEE4983B580F6CD0B3799DB/job/">apply with SolarWinds</a>.

Not to downplay its significance, but, The Open Organization released its latest book, <a href="https://opensource.com/open-organization/resources/culture-change?sc_cid=7016000000127L3AAI"><em>The Open Organization Guide to IT Culture Change</em></a> this week. The book as a whole is <a href="https://opensource.com/open-organization/17/6/working-open-and-gsd">a feat of open source community</a> in and of itself. It is a quality piece of work based off what I read during the open editing process. I’m still trying to read through it between working, packing, staging, and sleeping. I was very fortunate to have been able to contribute to this book. It’s a free download and if you want a physical copy it will be available soon in print! Also, if you want to convert the ODT to an ePub <a href="https://github.com/open-organization-ambassadors/open-org-it-culture/issues/36">we are accepting pull requests</a> (my eyes will thank you).

<a href="https://www.cyberciti.biz/security/linux-security-alert-bug-in-sudos-get_process_ttyname-cve-2017-1000367/">There is/was a super gnarly bug in sudo allowing root access to any user with a shell account</a>. The <a href="http://www.openwall.com/lists/oss-security/2017/05/30/16">full details are insightful</a> but, patch, patch, patch!!!

British Airways had quite the debacle. Apparently, <a href="http://www.independent.co.uk/news/business/news/british-airways-system-outage-it-worker-power-supply-switch-off-accident-flights-delayed-cancelled-a7768581.html">a human error took down all BA flights</a> in and out of their primary hubs. Don’t worry, <a href="https://www.theguardian.com/business/2017/may/29/british-airways-ceo-will-not-resign-despite-catastrophic-it-failure">BA doesn’t think of themselves as a technology company so their CEO is staying put</a>. This is either blameless culture or incompetence. Where do we draw the line? Does blameless culture go all the way up to the CEO when we strand passengers in airports past their visas?

<a href="https://medium.com/@SergeyNuzhdin/going-open-source-in-monitoring-part-i-deploying-prometheus-and-grafana-to-kubernetes-e3d44460cec6">Going open-source in monitoring, part I: Deploying Prometheus and Grafana to Kubernetes</a>

<a href="https://github.com/qntm/base65536">base65536: Unicode’s answer to Base64</a>

<a href="http://blog.kubernetes.io/2017/05/managing-microservices-with-istio-service-mesh.html">Managing microservices with the Istio service mesh</a>

If you somehow think <a href="https://devopsish.com/">DevOps</a> is still a fad allow me to point out the <a href="https://news.ycombinator.com/item?id=14460777">June 2017 Hacker News “Who is hiring”</a> has 97 mentions of DevOps.

<a href="https://github.com/dzhou121/gonvim">gonvim</a> looks really cool.

DNS is everything. I love DNS and DNS hates me back and that’s totally okay. But to read <a href="https://githubengineering.com/dns-infrastructure-at-github/">how GitHub re-architected their DNS infrastructure</a> made me oddly happy. Total nerd out on this one.

<a href="https://jakearchibald.com/2017/h2-push-tougher-than-i-thought/">HTTP/2 push is tougher than I thought</a>: No shit?

<a href="https://www.infoq.com/news/2017/05/forsgren-devops-performance">DevOps and the Key to High Performance: Nicole Forsgren at the DevOps Enterprise Summit London</a>

I still want to be either Jess Frazelle or Kelsey Hightower when I grow up. Probably more Jess right now after seeing her <a href="https://docs.google.com/presentation/d/17Hml1iFqdXElxOcrh9caQSC5px5mDgaS015Vhaz42ZY/edit#slide=id.p">Container Linux on the Desktop</a> slides.

<a href="https://summitroute.com/blog/2017/05/30/free_tools_for_auditing_the_security_of_an_aws_account/">Free tools for auditing the security of an AWS account</a> is a must read.

Damn it! <a href="https://blogs.oracle.com/developers/kubernetes-community-engagement-time-to-roll">Who let Oracle discover Kubernetes?!?</a>

<a href="https://medium.com/@pritunl/pritunl-1-28-release-announcement-5e02c519d8ab">Pritunl 1.28 was released this week</a>. This fixed an odd issue I was having with licensing. Not sure if that was unique to this fix but it was nice to not have to open a support request. There’s other awesome stuff in this release too.

Hot on the heels of 3.6, <a href="https://alpinelinux.org/posts/Alpine-3.6.1-released.html">Alpine Linux 3.6.1 has been released</a> patching quite a few CVEs.

<a href="https://groups.google.com/forum/#!topic/ansible-announce/i3hsJKuVusQ">Ansible 2.3.1 and 2.1.6 FINAL have been released!</a>

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Data Defense

1. [There is/was a super gnarly bug in sudo allowing root access to any user with a shell accountfull details are insightful](https://www.cyberciti.biz/security/linux-security-alert-bug-in-sudos-get_process_ttyname-cve-2017-1000367/)

    . The  but, patch, patch, patch!!!
1. [a human error took down all BA flightsBA doesn’t think of themselves as a technology company so their CEO is staying put](http://www.independent.co.uk/news/business/news/british-airways-system-outage-it-worker-power-supply-switch-off-accident-flights-delayed-cancelled-a7768581.html)

    British Airways had quite the debacle. Apparently,  in and out of their primary hubs. Don’t worry, . This is either blameless culture or incompetence. Where do we draw the line? Does blameless culture go all the way up to the CEO when we strand passengers in airports past their visas?
### Department of Choice Concepts

1. [Going open-source in monitoring, part I: Deploying Prometheus and Grafana to Kubernetes](https://medium.com/@SergeyNuzhdin/going-open-source-in-monitoring-part-i-deploying-prometheus-and-grafana-to-kubernetes-e3d44460cec6)

    
1. [base65536: Unicode’s answer to Base64](https://github.com/qntm/base65536)

    
1. [Managing microservices with the Istio service mesh](http://blog.kubernetes.io/2017/05/managing-microservices-with-istio-service-mesh.html)

    
1. [DevOpsJune 2017 Hacker News “Who is hiring”](https://devopsish.com/)

    If you somehow think  is still a fad allow me to point out the  has 97 mentions of DevOps.
1. [gonvim](https://github.com/dzhou121/gonvim)

    looks really cool.
1. [how GitHub re-architected their DNS infrastructure](https://githubengineering.com/dns-infrastructure-at-github/)

    DNS is everything. I love DNS and DNS hates me back and that’s totally okay. But to read  made me oddly happy. Total nerd out on this one.
1. [HTTP/2 push is tougher than I thought](https://jakearchibald.com/2017/h2-push-tougher-than-i-thought/)

     No shit?
### Department of Discussion

1. [DevOps and the Key to High Performance: Nicole Forsgren at the DevOps Enterprise Summit London](https://www.infoq.com/news/2017/05/forsgren-devops-performance)

    
1. [Container Linux on the Desktop](https://docs.google.com/presentation/d/17Hml1iFqdXElxOcrh9caQSC5px5mDgaS015Vhaz42ZY/edit#slide=id.p)

    I still want to be either Jess Frazelle or Kelsey Hightower when I grow up. Probably more Jess right now after seeing her  slides.
### Department of Happy Clouds

1. [Free tools for auditing the security of an AWS account](https://summitroute.com/blog/2017/05/30/free_tools_for_auditing_the_security_of_an_aws_account/)

    is a must read.
1. [Who let Oracle discover Kubernetes?!?](https://blogs.oracle.com/developers/kubernetes-community-engagement-time-to-roll)

    Damn it!
### Department of Refreshment and Refurbishment

1. [Pritunl 1.28 was released this week](https://medium.com/@pritunl/pritunl-1-28-release-announcement-5e02c519d8ab)

    . This fixed an odd issue I was having with licensing. Not sure if that was unique to this fix but it was nice to not have to open a support request. There’s other awesome stuff in this release too.
1. [Alpine Linux 3.6.1 has been released](https://alpinelinux.org/posts/Alpine-3.6.1-released.html)

    Hot on the heels of 3.6,  patching quite a few CVEs.
1. [Ansible 2.3.1 and 2.1.6 FINAL have been released!](https://groups.google.com/forum/#!topic/ansible-announce/i3hsJKuVusQ)

    
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-025.md) ------------- [ Next >> ](devopsweekly-027.md)