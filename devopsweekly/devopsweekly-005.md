## [DevOps'ish 005](https://devopsish.com/005) - Sun Jan 8, 2017

Oh the weather outside is frightful,

But the <a href="https://devopsish.com/"><strong>DevOps</strong></a> is so delightful,

And since we’ve no place to go,

Let It Snow! Let It Snow! Let It Snow!

This song has been in my head all day as the US east coast has been coated in a blanket of frozen participation. We have not taken down our interior Christmas decorations yet either so that might be part of it too.

<a href="https://medium.com/learning-the-go-programming-language/why-i-wrote-a-book-on-go-programming-b67aa5d3067b#.utvbksmbn">Vladimir Vivien</a> wrote a book on Go. <a href="https://vladimirvivien.github.io/learning-go/"><em>Learning Go Programming</em></a> was written over the course of a year and a half and is an attempt to make it easier for newcomers. I bought it and intend to read it.

<a href="../004/">When I mentioned the leap second last week</a> I was hoping it would go off without a hitch. That was not the case when Cloudflare had a <a href="https://blog.cloudflare.com/how-and-why-the-leap-second-affected-cloudflare-dns/">globally impacting issue with its RRDNS software</a>. Developers have a hard job, I believe all of us in <a href="https://devopsish.com/">DevOps</a> recognize that. But, this is a good example of how <a href="https://chrisshort.net/take-ownership-plant-your-flag/">an assumption caused a big problem</a>.

Three hours after the New Year’s leap second, Ask Bjørn Hansen reported that <a href="https://community.ntppool.org/t/leap-second-2017-status/59">only 26% of servers</a> in the <a href="http://www.pool.ntp.org/en/">NTP pool</a> were not announcing the leap second. That is not that great but given that the bar to enter the pool is quite low it is not surprising.

You know it’s bad when The Register is talking about it. <a href="http://www.theregister.co.uk/2017/01/04/mongodb_installs_wiped_by_bitcoin_ransoming_script/">There are 25,000 publicly exposed MongoDB systems out on the Internet</a> and about 2,000 of them have been taken over by ransomware. I am shaking my head vigorously. Are these MongoDB admins the same people heralding “serverless” (did I go too far there)? If you happen to be a victim make sure you confirm <a href="https://www.databreaches.net/dont-pay-the-mongodb-ransom-until-you-check-to-see-if-its-a-scam/">whether or not the ransomware is a hoax</a>.

Ansible is launching a single day global conference tour called <a href="https://www.ansible.com/automates">Ansible Automates</a>. What is it? Ansible Automates is, “one-track conference that takes the best Ansible content and presents it to regional audiences around the world… we expect crowds of 200–400 people who are interested in learning more about IT automation solutions with Ansible…” Ansible Automates might be an answer to some feedback I provided after AnsibleFest Brooklyn. I did not really like the four track conference because I knew I was missing good content elsewhere.

<a href="https://blog.jessfraz.com/">Jess Frazelle</a> is <a href="https://twitter.com/jessfraz/status/815946957474721796">looking for underrepresented folks (minorities, women, veterans, etc.) who want help speaking at conferences this year</a>. If you want to get up in front of a group of people and share your experience but want some guidance let Jess know by e-mailing jess at linux dot com.

NTPsec is an effort to streamline and optimize the core NTP code into a sustainable effort for a critical core function of the Internet. The project is seriously <a href="https://blog.ntpsec.org/2017/01/03/getting-past-c.html">considering a move away from C</a> and the post about it is absolutely fascinating.

<a href="https://blog.openebs.io/openebs-the-containerized-storage-f76e394a9543#.tie13be63">OpenEBS is a block storage solution for container environments</a>. It features “Virtual Storage Machines, a concept similar to k8s PODs.” It seems pretty straight forward. But will it take off, get sucked up by another project, or be bought and added to the portfolio of another larger product line?

The highly publicized death of Flocker left some folks in a lurch. Thankfully the <a href="http://codedellemc.com/">{code} by Dell/EMC</a> folks have <a href="https://blog.codedellemc.com/2017/01/06/migrate_flocker_drivers_rexray/">a solution with their REX-Ray product</a>.

<a href="https://github.com/dxa4481/truffleHog">Truffle Hog</a> searches through git repositories for high entropy strings, digging deep into commit history and branches. This is effective at finding secrets accidentally committed that contain high entropy.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Choice Concepts

1. [Vladimir VivienLearning Go Programming](https://medium.com/learning-the-go-programming-language/why-i-wrote-a-book-on-go-programming-b67aa5d3067b#.utvbksmbn)

    wrote a book on Go.  was written over the course of a year and a half and is an attempt to make it easier for newcomers. I bought it and intend to read it.
### Department of Dafuq

1. [When I mentioned the leap second last weekglobally impacting issue with its RRDNS softwareDevOpsan assumption caused a big problem](../004/)

    I was hoping it would go off without a hitch. That was not the case when Cloudflare had a . Developers have a hard job, I believe all of us in  recognize that. But, this is a good example of how .
### Department of Data Defense

1. [only 26% of serversNTP pool](https://community.ntppool.org/t/leap-second-2017-status/59)

    Three hours after the New Year’s leap second, Ask Bjørn Hansen reported that  in the  were not announcing the leap second. That is not that great but given that the bar to enter the pool is quite low it is not surprising.
1. [There are 25,000 publicly exposed MongoDB systems out on the Internetwhether or not the ransomware is a hoax](http://www.theregister.co.uk/2017/01/04/mongodb_installs_wiped_by_bitcoin_ransoming_script/)

    You know it’s bad when The Register is talking about it.  and about 2,000 of them have been taken over by ransomware. I am shaking my head vigorously. Are these MongoDB admins the same people heralding “serverless” (did I go too far there)? If you happen to be a victim make sure you confirm .
### Department of Discussions

1. [Ansible Automates](https://www.ansible.com/automates)

    Ansible is launching a single day global conference tour called . What is it? Ansible Automates is, “one-track conference that takes the best Ansible content and presents it to regional audiences around the world… we expect crowds of 200–400 people who are interested in learning more about IT automation solutions with Ansible…” Ansible Automates might be an answer to some feedback I provided after AnsibleFest Brooklyn. I did not really like the four track conference because I knew I was missing good content elsewhere.
1. [Jess Frazellelooking for underrepresented folks (minorities, women, veterans, etc.) who want help speaking at conferences this year](https://blog.jessfraz.com/)

    is . If you want to get up in front of a group of people and share your experience but want some guidance let Jess know by e-mailing jess at linux dot com.
### Department of Refreshment and Refurbishment

1. [considering a move away from C](https://blog.ntpsec.org/2017/01/03/getting-past-c.html)

    NTPsec is an effort to streamline and optimize the core NTP code into a sustainable effort for a critical core function of the Internet. The project is seriously  and the post about it is absolutely fascinating.
1. [OpenEBS is a block storage solution for container environments](https://blog.openebs.io/openebs-the-containerized-storage-f76e394a9543#.tie13be63)

    . It features “Virtual Storage Machines, a concept similar to k8s PODs.” It seems pretty straight forward. But will it take off, get sucked up by another project, or be bought and added to the portfolio of another larger product line?
1. [{code} by Dell/EMCa solution with their REX-Ray product](http://codedellemc.com/)

    The highly publicized death of Flocker left some folks in a lurch. Thankfully the  folks have .
1. [Truffle Hog](https://github.com/dxa4481/truffleHog)

    searches through git repositories for high entropy strings, digging deep into commit history and branches. This is effective at finding secrets accidentally committed that contain high entropy.
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](sreweekly-4.md) ------------- [ Next >> ](sreweekly-6.md)