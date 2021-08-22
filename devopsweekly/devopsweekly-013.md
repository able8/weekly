## [DevOps'ish 013](https://devopsish.com/013) - Sun Mar 5, 2017

I have been sick all week. I have pushed through it as much as possible but this was not a great week physically. However, there is plenty to discuss in the land of <a href="https://devopsish.com/"><strong>DevOps</strong></a>.

The Amazon S3 outage is the big story of the week. How many of you were impacted by it. I was to a minimal extent. Our products have been architected for failure pretty well. There was a little work to be done during the outage to work around one issue. Aside from that my time was spent monitoring things anticipating the catastrophe to expand (which it did to an extent). This did not cause any additional burden on my teams and I though. Why am I not mentioning <a href="https://aws.amazon.com/message/41926/">the reason behind the outage</a>? Because it should not have been as devastating as it was. Folks, the cloud is still something you have to manage. Remember <a href="https://landing.google.com/sre/book/chapters/introduction.html">error budgets from Google’s Site Reliability Engineering</a> book? You have to anticipate outages and down time for just about everything in this world. Your code should not depend on one cloud provider’s regional resource, ever. <strong>Diversify your cloud</strong>, folks!

Mattias Geniar announced that his <a href="https://dnsspy.io/?utm_source=devopsish.com">DNS Spy</a> tool has entered public beta. DNS Spy is a nice and handy tool to monitor and alert on changes to your DNS records. This is nice if you ever think you might be targeted for DNS hijacking (which you probably will be).

<a href="https://www.blackducksoftware.com/about/news-events/releases/black-duck-announces-2016-open-source-rookies-year">Black Duck Software announced its 2016 Open Source ‘Rookies of the Year’</a>. One notable mention on the list is <a href="https://www.ansible.com/ansible-container">Ansible Container</a>.

The US Department of Defense has spun up <a href="https://github.com/deptofdefense/code.mil#welcome-to-codemil---an-experiment-in-open-source-at-the-department-of-defense">Code.mil</a>. They are requesting feedback on their <a href="https://github.com/deptofdefense/code.mil/blob/master/Proposal/CONTRIBUTING.md">proposed open source strategy</a>. Please take a look at it and if you have expertise in this arena help them out.

Aaron Bieber shared a <a href="https://deftly.net/posts/2017-02-27-ssh-fp-verification-using-tor.html">unique method to use Tor to verify SSH fingerprints</a>. If you are worried about being subject to MitM attacks have a look.

Google stalled its TLS 1.3 rollout this week thanks to the <a href="https://bugs.chromium.org/p/chromium/issues/detail?id=694593">shoddy work of BlueCoat</a> (I am biased; I loathe BlueCoat’s products for legitimate reasons). When modern versions of Chrome connected through BlueCoat products to Google services the connection would just hang.

Since we are on the topic of AWS, <a href="https://medium.com/@cmcluck">Craig McLuckie</a> published Heptio’s <a href="https://blog.heptio.com/aws-quickstart-for-kubernetes-26ccaf7e1c8f#.io92vwkvr">AWS Quickstart for Kubernetes</a>. This is Heptio’s CloudFormation template and written guide to aid folks getting Kubernetes up and running quickly.

I keep coming back to this time and time again lately. The Cloud Native Computing Foundation maintains an incredibly handy <a href="https://github.com/cncf/landscape">Cloud Native Landscape</a>. If you are ever wondering what tools are available to handle a certain task look here first (and share this with your CIO/CTO).

<a href="https://news.netcraft.com/archives/2017/02/27/february-2017-web-server-survey.html">Apache Web Server is continuing to lose market share</a> at a pretty steady clip. Microsoft has the largest market share and nginx is continuing to gain on Apache and is within striking distance of overtaking the Apache behemoth. Interestingly, I am considering writing my own web server in Go with no configuration files; compile and launch a single binary.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Data Defense

1. [the reason behind the outageerror budgets from Google’s Site Reliability Engineering](https://aws.amazon.com/message/41926/)

    The Amazon S3 outage is the big story of the week. How many of you were impacted by it. I was to a minimal extent. Our products have been architected for failure pretty well. There was a little work to be done during the outage to work around one issue. Aside from that my time was spent monitoring things anticipating the catastrophe to expand (which it did to an extent). This did not cause any additional burden on my teams and I though. Why am I not mentioning ? Because it should not have been as devastating as it was. Folks, the cloud is still something you have to manage. Remember  book? You have to anticipate outages and down time for just about everything in this world. Your code should not depend on one cloud provider’s regional resource, ever. Diversify your cloud, folks!
### Department of Choice Concepts

1. [DNS Spy](https://dnsspy.io/?utm_source=devopsish.com)

    Mattias Geniar announced that his  tool has entered public beta. DNS Spy is a nice and handy tool to monitor and alert on changes to your DNS records. This is nice if you ever think you might be targeted for DNS hijacking (which you probably will be).
1. [Black Duck Software announced its 2016 Open Source ‘Rookies of the Year’Ansible Container](https://www.blackducksoftware.com/about/news-events/releases/black-duck-announces-2016-open-source-rookies-year)

    . One notable mention on the list is .
1. [Code.milproposed open source strategy](https://github.com/deptofdefense/code.mil#welcome-to-codemil---an-experiment-in-open-source-at-the-department-of-defense)

    The US Department of Defense has spun up . They are requesting feedback on their . Please take a look at it and if you have expertise in this arena help them out.
1. [unique method to use Tor to verify SSH fingerprints](https://deftly.net/posts/2017-02-27-ssh-fp-verification-using-tor.html)

    Aaron Bieber shared a . If you are worried about being subject to MitM attacks have a look.
### Department of Dafuq

1. [shoddy work of BlueCoat](https://bugs.chromium.org/p/chromium/issues/detail?id=694593)

    Google stalled its TLS 1.3 rollout this week thanks to the  (I am biased; I loathe BlueCoat’s products for legitimate reasons). When modern versions of Chrome connected through BlueCoat products to Google services the connection would just hang.
### Department of Happy Little Clouds

1. [Craig McLuckieAWS Quickstart for Kubernetes](https://medium.com/@cmcluck)

    Since we are on the topic of AWS,  published Heptio’s . This is Heptio’s CloudFormation template and written guide to aid folks getting Kubernetes up and running quickly.
1. [Cloud Native Landscape](https://github.com/cncf/landscape)

    I keep coming back to this time and time again lately. The Cloud Native Computing Foundation maintains an incredibly handy . If you are ever wondering what tools are available to handle a certain task look here first (and share this with your CIO/CTO).
### Department of Next Year’s Old Tech

1. [Apache Web Server is continuing to lose market share](https://news.netcraft.com/archives/2017/02/27/february-2017-web-server-survey.html)

    at a pretty steady clip. Microsoft has the largest market share and nginx is continuing to gain on Apache and is within striking distance of overtaking the Apache behemoth. Interestingly, I am considering writing my own web server in Go with no configuration files; compile and launch a single binary.

### [ << Prev ](sreweekly-12.md) ------------- [ Next >> ](sreweekly-14.md)