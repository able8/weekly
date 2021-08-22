## [DevOps'ish 173](https://devopsish.com/173) - Sun Mar 29, 2020

The impact of the coronavirus hit the tech events industry hard this week. In a shocking move on Tuesday, <a href="https://www.oreilly.com/conferences/from-laura-baldwin.html">O’Reilly announced it has immediately ceased all in-person events</a>. To add insult to injury, <a href="https://twitter.com/joshsimmons/status/1242596429924728832">O’Reilly laid off their entire events staff during a pandemic</a>. To make matters worse, under US law, they’ll get only a week of health insurance benefits, <a href="https://twitter.com/jhscott/status/1242604187294154752">meager support for COBRA (super expensive health insurance), and any severance is based solely on tenure</a>. DevOps’ish is an O’Reilly community partner. I heard from the O’Reilly partner team this week that they were full steam ahead. I was slated to speak at the <a href="https://conferences.oreilly.com/infrastructure-ops/io-ca/public/schedule/speaker/200908">O’Reilly Infrastructure &amp; Ops Conference in June</a>). I haven’t heard a peep out of O’Reilly about what happens now. <a href="https://www.oreilly.com/about/open-source-matters-more-than-ever.html">Tim O’Reilly mustered up this ‘lipstick on a pig’ post</a>, though, so we should all feel great about it, right? Guess what; I don’t.

### People

1. [The Hidden Vulnerabilities of Open Source Software](https://hbswk.hbs.edu/item/the-hidden-vulnerabilities-of-open-source-software)

    The increasing use of open source software in most commercial apps has revolutionized software development-but also created hidden vulnerabilities, say Frank Nagle and Jenny Hoffman.
1. [That Discomfort You’re Feeling Is Grief](https://hbr.org/2020/03/that-discomfort-youre-feeling-is-grief)

    I know an article is a big deal when my wife says I should read it. “If we can name it, perhaps we can manage it.” An HBR interview with David Kessler.
1. [14 Kubernetes interview questions: For hiring managers and job seekers](https://enterprisersproject.com/article/2020/3/14-kubernetes-interview-questions)

    Filling Kubernetes jobs can be tricky because the technology is relatively young. Experts share interview questions to help hiring managers sort candidates - and help job seekers prepare
1. [Former Linux Developer Hans Reiser To Remain Locked Up](https://www.phoronix.com/scan.php?page=news_item&px=Hans-Reiser-Locked-Up-No-Parole)

    I didn’t even realize this is what happened to ReiserFS. In other news, when you get 15 to life, don’t expect to be contributing to the kernel anytime soon. Also, never read the comments.
1. [Microsoft staff giggle beneath the weight of a 52,000-person Reply-All email stormWe even have a sticker](https://www.theregister.co.uk/2020/03/26/microsoft_reply_all_email_storm_52000/)

    We call these unsubsquirrels at Red Hat and they are generally hilarious. .
1. [Gresecurity maker finally coughs up $300k to foot open-source pioneer Bruce Perens’ legal bill in row over GPL](https://www.theregister.co.uk/2020/03/27/gresecurity_perens_gpl_settlement/)

    In other news, when you sue people because of their opinions, don’t get surprised when you’re forced to pay their legal fees.
1. [What happens when the maintainer of a JS library downloaded 26m times a week goes to prison for killing someone with a motorbike? Core-js just found out](https://www.theregister.co.uk/2020/03/26/corejs_maintainer_jailed_code_release/)

    What will be the fate of an open-source project relied upon by so many?
### Process

1. [Things no one tells you when you start a newsletter](https://chrisshort.net/things-no-one-tells-you-when-you-start-a-newsletter/)

    I took some of my notes and compared them with Gareth Rushgrove (Devops Weekly) and Mattias Geniar (cron.weekly) on what it’s like writing a tech-related newsletter. It’s very rewarding, but it is not trivial.
1. [Elizabeth Warren’s Campaign Is Making Its Software Open Source](https://www.wired.com/story/elizabeth-warren-campaign-open-source-tech/)

    One good thing from this US election cycle is that more people in politics will understand open source.
1. [Helping FOSS conferences in the face of a pandemic](https://lwn.net/SubscriberLink/815913/546f9bedce72c9ed/)

    “There are some resources listed on the web site, including organizations focused more widely on the problems that COVID-19 is bringing to our world (e.g. the World Health Organization response fund). There are also pointers to web pages covering virtual event organization and the mechanics of running an online event. The FOSS Responders web site makes it clear that many of the items posted there are simply meant as starting points; those with ideas are encouraged to get in touch with the project.”
1. [Definitive Guide to AWS EKS Security - Download eBook](https://devopsi.sh/guide-elastic-1c07c)

    When using Amazon’s Elastic Kubernetes Service (EKS), you must understand which pieces of the security management role fall on you. AWS takes responsibility for securing their infrastructure and addressing security issues in their software. The customer must ensure the security of their own applications while correctly using the controls offered to protect their data and workloads. In this 42-page eBook, StackRox covers general EKS cluster security, including the standard controls and best practices for minimizing the risk around cluster workloads, as well as specific requirements for securing an EKS cluster and its associated infrastructure. SPONSORED
1. [Zoom needs to clean up its privacy act](https://blogs.harvard.edu/doc/2020/03/27/zoom/)

    “A person whose personal data is being shed on Zoom doesn’t know that’s happening because Zoom doesn’t tell them. There’s no red light, like the one you see when a session is being recorded. If you were in a browser instead of an app, an extension such as Privacy Badger could tell you there are trackers sniffing your ass. And, if your browser is one that cares about privacy, such as Brave, Firefox or Safari, there’s a good chance it would be blocking trackers as well. But in the Zoom app, you can’t tell if or how your personal data is being harvested.”
1. [Pentesting a banking FTP service](https://blog.lexfo.fr/pentesting-pesit-ftp.html)

    If you’ve ever had the joy of handling large businesses’ financial records, you know the pleasure of banking FTP servers. How these things are STILL running is just beyond my comprehension.
### Tools

1. [Kubernetes 1.18: Fit & Finish](https://kubernetes.io/blog/2020/03/25/kubernetes-1-18-release-announcement/)

    Kubernetes 1.18 was released this week (after a slight delay). “Kubernetes 1.18 consists of 38 enhancements: 15 enhancements are moving to stable, 11 enhancements in beta, and 12 enhancements in alpha.” Thank you to all the contributors and especially the 1.18 release team!
1. [Extended and Improved WebAssemblyHub to Bring the Power of WebAssembly to Envoy and Istio](https://istio.io/blog/2020/wasmhub-istio/)

    “A place to build, publish, share, and deploy WebAssembly Envoy extensions”
1. [Serverless Service Mesh with Knative and Linkerd](https://linkerd.io/2020/03/23/serverless-service-mesh-with-knative-and-linkerd/)

    The Istio requirement has been dropped so you can use Linkerd with Knative now.
1. [Using UBI images to minimize container vulnerabilities](https://snyk.io/blog/ubi-to-minimize-container-vulnerabilities/)

    “From a technical perspective, they are nearly identical to Red Hat Enterprise Linux images, which means they have great security, performance, and life cycles. They are released under a different End User License Agreement - It’s possible to build a containerized application using UBI, push it to any registry server, easily share it with others - and because it’s freely redistributable - even deploy it on non-Red Hat platforms.”
1. [Knative Crowds out Other Serverless Software (and Other CNCF Survey Takeaways)](https://thenewstack.io/knative-crowds-out-other-serverless-software-packages-and-other-cncf-survey-takeaways/)

    A different and in some ways, better take on the CNCF survey data.
1. [Online Debug for AWS Lambda on your IDE!](https://devopsi.sh/aws-lambda-97890)

    Thundra allows you to natively debug your serverless applications on the cloud with their own permissions. Thundra’s online-debugger sets up a secure bridge between your AWS Lambda environment and your IDE. VSCode and IntelliJ IDEA are natively supported with plugins. For other IDEs, we provide a portable client to foster the integration with any IDEs. Start debugging Node.js, Python and Java functions for free today! SPONSORED
1. [spanner.fyi](https://spanner.fyi/)

    Bits of knowledge on Spanner from Jaana Dogan and the Cloud Spanner team at Google
1. [Amazon’s Arm-based Graviton2 Against AMD and Intel: Comparing Cloud Compute](https://www.anandtech.com/show/15578/cloud-clash-amazon-graviton2-arm-against-intel-and-amd)

    If you can cross-compile and run your code on ARM infrastructure, do it! You’ll save a ton of money.
1. [darkbitio/mkit](https://github.com/darkbitio/mkit)

    “MKIT is a ‘Managed Kubernetes Inspection Tool’ that validates several common security-related configuration settings of managed Kubernetes cluster objects and the workloads/resources running inside the cluster.”
1. [codingCoffee/fahclient](https://github.com/codingCoffee/fahclient)

    Dockerized Folding@home client with NVIDIA GPU support

### [ << Prev ](devopsweekly-172.md) ------------- [ Next >> ](devopsweekly-174.md)