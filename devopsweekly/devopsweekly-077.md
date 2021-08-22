## [DevOps'ish 077](https://devopsish.com/077) - Sun May 27, 2018

I have a lot to talk about this week. I’m putting sections in the introduction for ease of reading (feel free to skip parts you’re not mildly interested in). Also, thank you to this week’s sponsors; <a href="https://www.replicated.com/"><strong>Replicated</strong></a> and <a href="https://www.gocd.org/2018/05/08/continuous-delivery-microservices-test-strategy/"><strong>GoCD</strong></a>!

The highlight of my week was being able to meet 1,400 new friends at <a href="http://chefconf.chef.io/">ChefConf</a> 2018 in Chicago. If you are ever wondering how to run an amazing event, look no further than the good folks at Chef. From a speaker’s perspective, I have never felt more prepared for a talk. The event itself was meticulously planned and executed. I presented my talk, <em>DevOps is Not War</em> to a standing room only audience. Thanks to the fantastic guidance from the conference’s speaking coach, <a href="http://messageglue.com/">Anna Boynton</a>, I was able to take a forty minute talk and land it perfectly inside the twenty-minute time slot I had. The event morphed my bias of Chef from a niche tool to a formidable accompaniment of software to ease automation tasks. The messages from the keynotes were like those I’ve delivered to many audiences before, “You’re not an X company, you’re a technology company.” I highly recommend adding ChefConf 2019 to your schedules as soon as it is announced.

I won’t <del>bore</del> annoy you with more GDPR news. But, I do have one ask from readers: What do you want to know about the data gathered by this newsletter and accompanying web site? I think GDPR is important but, I have not changed much to implement it (and I’m not really sure if I need to either). With that being said, I present to you the <a href="https://gdprhallofshame.com/">👎 GDPR Hall of Shame</a>.

There were two announcements this week from the <a href="https://golang.org/">Go</a> community that merit attention: one is good, the other is disconcerting. First, Go has adopted a new <a href="https://golang.org/conduct">Code of Conduct</a> (based on Google’s new <a href="https://opensource.google.com/docs/releasing/template/CODE_OF_CONDUCT/">Code of Conduct template</a>) Go’s new Code of Conduct now applies to community members working in non-Go community spaces. This change is a good thing. You shouldn’t be able to prey on people in a community outside of that community’s space and expect to have your standing in the community protected. This allows for open secrets to be acted upon (and I feel transparency is good).

Sadly, this good news was offset by something that I’m deeply concerned about in the Go community. You may recall at GopherCon 2017, <a href="https://youtu.be/5LtMb090AZI">Sam Boyer gave a keynote</a> stating that <a href="https://github.com/golang/dep"><strong>dep</strong></a> was the “official experiment” for Go package management. According to <a href="https://codeengineered.com/blog/2018/golang-godep-to-vgo/">Matt Farina</a> at the very same GopherCon, “Russ Cox, the current lead of the Go team, came to the table. He made a comment that he could do better if he went off on his own and built something. That something was later announced as vgo.” This indicates that dep literally had no chance at being the official anything for Go. Russ Cox announced on 2018-04-25 that the <a href="https://github.com/golang/go/issues/24301#issuecomment-384349642">vgo proposal has been accepted</a>. Sam Boyer is a friend. I know he has spent countless hours working on dep as well as analyzing vgo. Sam has several articles in progress analyzing vgo; <a href="https://sdboyer.io/vgo/intro/">An Analysis of Vgo</a> and <a href="https://sdboyer.io/vgo/failure-modes/">Failure Modes</a> have already been published.

My concern as a somewhat outsider looking in, as well as others I’ve talked to in the Go community, is not in the solution itself but how the communication has been handled. Why announce dep as even a remote possibility as the one way forward at GopherCon 2017? Why was Russ the proposer, writer, and approver of something that clearly needed more debate? I understand Go is a Google project but, this seems rather dictatorial. It’s not a good look for Go’s core maintainers (to say the least). My worry is that this could hurt a young and growing community. I would hate to see that. I ❤️ Go!

It is Memorial Day weekend here in the US. One thing we veterans get mildly irritated about is when citizens don’t recognize the difference between Memorial Day and Veterans Day. Remember this: <strong>we remember the fallen on Memorial Day. We thank veterans on Veterans Day</strong>.

<a href="https://www.replicated.com/"><strong>Replicated: The modern way to ship and manage enterprise software</strong></a><br/>Replicated gives SaaS and software vendors a cloud-native platform for easily and securely deploying their applications inside customers’ data centers or VPC environments. Replicated provides tooling for automatic updates, license management, support, audit logs, LDAP integration and more. <a href="https://vendor.replicated.com/">Sign up for a free trial</a> and get started now. <em>SPONSORED</em>

<a href="https://www.gocd.org/2018/05/08/continuous-delivery-microservices-test-strategy/"><strong>Test Strategy for Microservices</strong></a><br/>Testing microservices systems is significantly more nuanced and complex than testing a traditional monolithic application. Is traditional testing pyramids still relevant? How to verify overall system behavior? Check out our new post: Test Strategy for Microservices. <em>SPONSORED</em>

<a href="https://www.devopsdays.org/events/2018-toronto/welcome/"><strong>DevOpsDays Toronto 2018</strong></a><br/>Dates: 2018-05-30 through 2018-05-31<br/>I’ll admit it, I’ve never been to Canada. But, I’m definitely going to <a href="https://www.devopsdays.org/events/2018-toronto/welcome/">DevOpsDays Toronto</a> this year to present <em>What the Military Taught Me about DevOps</em>.

<a href="https://devnationfederal.org/"><strong>DevNation Federal</strong></a><br/>Date: 2018-06-05<br/>Join us to learn about the revolutions happening in communities around containers, data, and application modernization. This is an opportunity for you to hear how visionary teams in the federal government are innovating with open source, and hear from leaders in the private sector doing the same.

<a href="https://chaosconf.splashthat.com/"><strong>Chaos Conf</strong></a><br/>Date: 2018-09-28<br/>Chaos Conf looks super awesome. Opening the event will be Adrian Cockcroft, VP AWS, who called 2018 “The year of #chaosengineering”. Closing out the night will be Jessie Frazelle, one of the top #containers experts on the planet currently at Microsoft.

### Events

1. [DevOpsDays Toronto 2018DevOpsDays Toronto](https://www.devopsdays.org/events/2018-toronto/welcome/)

    Dates: 2018-05-30 through 2018-05-31I’ll admit it, I’ve never been to Canada. But, I’m definitely going to  this year to present What the Military Taught Me about DevOps.
1. [DevNation Federal](https://devnationfederal.org/)

    Date: 2018-06-05Join us to learn about the revolutions happening in communities around containers, data, and application modernization. This is an opportunity for you to hear how visionary teams in the federal government are innovating with open source, and hear from leaders in the private sector doing the same.
1. [Chaos Conf](https://chaosconf.splashthat.com/)

    Date: 2018-09-28Chaos Conf looks super awesome. Opening the event will be Adrian Cockcroft, VP AWS, who called 2018 “The year of #chaosengineering”. Closing out the night will be Jessie Frazelle, one of the top #containers experts on the planet currently at Microsoft.
### People

1. [opensource.com is giving away a System76 laptop](https://opensource.com/article/18/5/system76-oryx-pro-laptop-giveaway)

     Enter for a chance to win a brand new laptop from System76. Not just any laptop either; it’s a first generation Oryx Pro!
1. [Would you wear this muzzle around the office for private calling?](https://www.cnet.com/news/would-you-wear-this-muzzle-around-the-office-for-private-calling/)

     You’ve got to be kidding me! When are we going to build walls again?
1. [IBM’s Watson Health wing left looking poorly after ‘massive’ layoffs](https://www.theregister.co.uk/2018/05/25/ibms_watson_layoffs/)

     Up to 70% of staff shown the door this week, insiders claim
1. [15 books for kids who (you want to) love Linux and open source](https://opensource.com/article/18/5/books-kids-linux-open-source)

     Plus, three books for babies.
1. [‘Elitist den of hate’: Silicon Valley pastor decries hypocrisy of area’s rich liberals](https://www.theguardian.com/technology/2018/may/22/silicon-valley-pastor-gregory-stevens-wealth-liberals)

     Gregory Stevens resigns after tweets about Palo Alto, slamming tech industry greed and empty social justice promises
### Process

1. [Observability+](https://medium.com/observability)

     A fantastic site about all things observability from the brilliant JBD.
1. [Why do Kubernetes clusters in AWS cost more than they should?](https://medium.com/@dyachuk/why-do-kubernetes-clusters-in-aws-cost-more-than-they-should-fa510c1964c6)

     Like all things AWS you have to actively manage your cloud spend. Dmytro Dyachuk shows us how.
1. [Writing Technical Articles](https://chrisshort.net/writing-technical-articles/)

     People struggle with writing. Here’s my attempt at getting you started. “Pick three points, add an introduction and conclusion and you have the beginnings of a great article.”
1. [Has Kubernetes Already Become Too Unnecessarily Complex for Enterprise IT?](https://thenewstack.io/has-kubernetes-already-become-too-unnecessarily-complex-for-enterprise-it/)

     Is Enterprise IT stupid? No. There’s the answer to that question.
1. [Making Kubernetes work for the average engineer—via PaaS](https://www.infoworld.com/article/3273104/containers/making-kubernetes-work-for-the-average-engineer-via-paas.html)

     Despite being the hottest thing since, well, Docker, Kubernetes remains a dark art for most mainstream enterprises
1. [Will Kubernetes Collapse Under the Weight of Its Complexity?](https://www.influxdata.com/blog/will-kubernetes-collapse-under-the-weight-of-its-complexity/)

     Is it? I’m pretty sure Kelsey Hightower said it was boring in Austin last year.
1. [Navigating the container security ecosystem](https://opensource.com/article/18/5/navigating-container-security-ecosystem)

     As container adoption increases, security automation will become a critical element of every organization’s workflow.
1. [Five Words that Get In The Way of DevSecOps](https://www.scmagazine.com/five-words-that-get-in-the-way-of-devsecops/article/761870/)

     “Vulnerability, risk, policy, compliance and governance are words that get lost in translation between development, security and operations and cause confusion.”
1. [Improve Security with Automated Image Scanning Through CI/CD](https://thenewstack.io/improve-security-with-automated-image-scanning-through-ci-cd/)

     “Using cloud-native security tools that hook right into Jenkins or your favorite CI/CD tool, enterprise security teams can set policies for developers who are building container images.”
1. [The DevOps Security Checklist](https://www.sqreen.io/checklists/devops-security-checklist.html)

     This security checklist aims to give DevOps professionals a list of DevOps security best practices they can follow to implement DevSecOps.
1. [CNCF To Host Telepresence in the Sandbox](https://www.cncf.io/blog/2018/05/22/telepresence-in-the-sandbox/)

     “Telepresence is an open source tool — licensed under the Apache 2.0 License — that lets developers run a single service locally, while connecting that service to a remote Kubernetes cluster.”
1. [Kubernetes best practices: terminating with grace](https://cloudplatform.googleblog.com/2018/05/Kubernetes-best-practices-terminating-with-grace.html)

     “When it comes to distributed systems, handling failure is key.”
1. [Your Success with Enterprise Kubernetes Isn’t About Kubernetes](https://medium.com/google-cloud/your-success-with-enterprise-kubernetes-isnt-about-kubernetes-1c2b18dedc9)

     Running like Google is more than Kubernetes products or a nice UI.
### Tools

1. [“I still don’t see the point of containers and Kubernetes when you can just have a good old VM. Why should I have multiple containers for each process I need when I can just run them all on the VM?”](https://www.linkedin.com/feed/update/urn:li:activity:6406136826708451328/?commentUrn=urn:li:comment:(activity:6406136826708451328,6406138601528193024))

    A healthy debate happening on my LinkedIn page (of all places) about why containers matter (and VMs still do too):
1. [Accessing Kubernetes Services Without Ingress, NodePort, or LoadBalancer](https://medium.com/@kyralak/accessing-kubernetes-services-without-ingress-nodeport-or-loadbalancer-de6061b42d72)

     ECMP can make Kubernetes services accessible without Ingress, NodePort, or LoadBalancer service types.
1. [Kubernetes won – so now what?](http://redmonk.com/jgovernor/2018/05/25/kubernetes-won-so-now-what/)

     Iterative improvements, that’s what!
1. [Open Sourcing Zuul 2](https://medium.com/netflix-techblog/open-sourcing-zuul-2-82ea476cb2b3)

     Netflix open sourced Zuul 2, their “cloud gateway”
1. [All aboard the gRPC train](http://lpan.io/migrating-to-grpc/)

     A Datadog intern migrated a legacy, homegrown RPC to gRPC.
1. [ianmiell/autotrace](https://github.com/ianmiell/autotrace)

     Runs a process, and gives you the output along with other telemetry on the process, all in one terminal window.
1. [Building a Microservices Application in Go Following the CQRS Pattern](https://outcrawl.com/go-microservices-cqrs-docker/)

     This article walks through the development of a simplistic social network application where anyone can post anonymous messages.
1. [Announcement: Scope communityScope](https://groups.google.com/forum/#!topic/scope-community/wLlOV_0Lt7Q)

     Weaveworks has decided to promote and foster  as a community project.
1. [Go2 status](https://github.com/golang/go/wiki/Go2)

     This page tracks the status of “Go 2”. Go 2 is in a very early, planning phase.
1. [atlassian/escalator](https://github.com/atlassian/escalator)

     Escalator is a batch or job optimized horizontal autoscaler for Kubernetes
1. [Playing battleships over BGP](https://blog.benjojo.co.uk/post/bgp-battleships)

     “Two communities were produced that would allow me run a game of battleships over BGP”
1. [Open Sourcing Coinbase’s Secure Deployment Pipeline](https://engineering.coinbase.com/open-sourcing-coinbases-secure-deployment-pipeline-ae6c78e25517)

     Odin takes a description of a project release and then safely and securely launches it into AWS using auto-scaling groups.
1. [facebookincubator/katran](https://github.com/facebookincubator/katran)

     A high performance layer 4 load balancer
1. [avantoss/vault-infra](https://github.com/avantoss/vault-infra)

     Terraform to create Vault infrastructure
1. [The Kata Containers project launches version 1.0 of its lightweight VMs for containers](https://techcrunch.com/2018/05/22/the-kata-containers-project-hits-1-0/)

    
1. [State of Cloud Native CI/CD Tools for Kubernetes](https://engineering.opsgenie.com/cloud-native-continuous-integration-and-delivery-tools-for-kubernetes-e6ea34d308c)

     A smorgasbord of cloud native tooling!
1. [A Docker Image in Less Than 1000 Bytes](https://zwischenzugs.com/2018/05/22/a-docker-image-in-less-than-1000-bytes/)

    
1. [kubernetes/kubernetes-template-project](https://github.com/kubernetes/kubernetes-template-project)

     A template for starting new projects on the github.com/kubernetes organization
1. [How OpenFaaS came to rescue us!](https://medium.com/iconscout/how-openfaas-came-to-rescue-us-ec129518cd46)

     An overview on our experience with serverless computing and speeding up Iconscout
1. [Automatically Generating InSpec Controls from Terraform](https://blog.chef.io/2018/05/23/automatically-generating-inspec-controls-from-terraform/)

    
1. [鯨魚水上飄的修煉 Day 4：Raspberry Pi上的Kubernetes](https://medium.com/mr-efacani-teatime/%E9%AF%A8%E9%AD%9A%E6%B0%B4%E4%B8%8A%E9%A3%84%E7%9A%84%E4%BF%AE%E7%85%89-day-4-raspberry-pi%E4%B8%8A%E7%9A%84kubernetes-b9f60cbc5517)

     兩個月學習Docker & Kubernetes 的心得分享
1. [CalVer](http://calver.org/)

     CalVer is a software versioning convention that is based on your project’s release calendar, instead of arbitrary numbers.

### [ << Prev ](devopsweekly-076.md) ------------- [ Next >> ](devopsweekly-078.md)