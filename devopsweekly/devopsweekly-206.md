## [DevOps'ish 206](https://devopsish.com/206) - Sun Feb 21, 2021

Sometimes you don’t know what the world needs until someone tells you. On Monday this week, a friend asked if I had any additional books to point them to for Kubernetes help. I have a mile-long list in my head. I said, yeah, let me punch that up for you real quick. But, instead of creating a locked down doc or dust bin email, I built a website. Behold, <a href="https://kubernetesreadme.com/"><strong>Kubernetes README</strong></a>. It’s nothing really fancy. A copy of an existing site, with a different name, and data to help folks make a selection that fits their needs.

But, this is the beauty of working in an open source environment. I didn’t even think of typing them an email. It was going to be a website the second the ask came in. There was no reason to go to the effort of creating a list of books that would live in a vacuum. Sharing knowledge helps us all. It’s how folks can figure out how to grow on their own. Thanks for the idea, Justin. I hope it helps.

Please note, the <a href="https://devopsish.com/solarwinds-supply-chain-compromise/">DevOps’ish Solarwinds supply chain compromise Index</a> has many updates this week. Including a White House briefing on the matter.

### People

1. [Why did I leave Google or, why did I stay so long?](https://paygo.media/p/25171)

    “So the question of why I stayed has many different aspects to it. When we were evaluating the offers to sell the company, we asked ourselves what would really change by being acquired? Due to a bunch of mistakes early on, we did not own substantial amounts of equity and had a pretty bad relationship with some of our board members. I remember the bottom line: “wouldn’t you rather work for Larry Page than our current board”? We were committed to our mission and saw this as a change in the cap table rather than a change of mission. This counted on the fact that Google had promised us autonomy to continue to act as Waze and we believed them.”
1. [Ephemeral Environments as a Service 🤯🤯🤯Get started now](https://releaseapp.io/?utm_source=devopsish&utm_medium=email&utm_content=get-started&utm_campaign=202102)

    Do you find that your engineers spend too much time creating and maintaining staging environments and yet, there never seems to be enough environments to go around? A shortage of environments is a top driver of low developer productivity and often impacts an engineering team’s ability to ship features on time. With Release, you can get a full instance of your app with all of its services with every pull request. Your developers will never have to fight over staging environments again. . SPONSORED
1. [“I will slaughter you”](https://daniel.haxx.se/blog/2021/02/19/i-will-slaughter-you/)

    People are bat shit crazy.
1. [Kenya Is Becoming a Global Hub of FinTech Innovation](https://hbr.org/2021/02/kenya-is-becoming-a-global-hub-of-fintech-innovation)

    “For more than half a century, the U.S. was the center of global innovation for financial technology, inventing credit cards, ATMs, and online banking. Now, however, it’s falling behind, as China has become a leader of mobile payments, and now African countries — namely Kenya — are making huge strides with familiar technologies such as mobile phones and SMS-style messaging, and rapidly expanding the circle of financial inclusion. Companies can learn three important lessons from this most recent wave of innovation: bundling services (think: banking and cellular as one offering) is essential to success; finance is about trust, but trusted companies can lend their credibility to newcomers with promising offerings; and technology that enables mass adoption or expansion is often old, not cutting edge.”
1. [This Cloud Computing Billing Expert Is Very Funny. Seriously.](https://www.nytimes.com/2021/02/17/technology/corey-quinn-amazon-aws.html)

    It’s weird and cool when people you know are in the NYT. It’s like, “What am I doing wrong?” mixed with “Really?!? Corey?!?” Regardless, Corey has put in the work and deserves as much praise as we can throw on him.
1. [Margaret Mitchell fired from Google](https://www.axios.com/google-fires-another-ai-ethics-leader-6ef7dcd5-4583-4396-b5b3-129547ff3091.html)

    Wait… I got it. Google should use its AI to fix its broken ethics around AI.“Okay Google, what are we doing wrong with our AI team?” DING “Stop ‘alienating Black women.’”
1. [How Oracle Sells Repression in China](https://theintercept.com/2021/02/18/oracle-china-police-surveillance/)

    “In its bid for TikTok, Oracle was supposed to prevent data from being passed to Chinese police. Instead, it’s been marketing its own software for their surveillance work.” Like, my hands are dirty and even I wouldn’t work at Oracle.
### Process

1. [Deep PostgreSQL Thoughts: Resistance to Containers is Futile](https://info.crunchydata.com/blog/deep-postgresql-thoughts-resistance-to-containers-is-futile)

    “The world of computing is inexorably moving toward automating everything and distributing all the bits in containers. Don’t fear it, embrace it.”
1. [“10 GitHub Security Best Practices from the founder of GitHub”](https://snyk.io/blog/ten-git-hub-security-best-practices/)

    Are you using GitHub securely? Learn tips on handling sensitive data and how to ensure your PRs don’t introduce new vulnerabilities. Read the cheat sheet. SPONSORED
1. [Engineering dependability and fault tolerance in a distributed system](https://www.ably.io/blog/engineering-dependability-and-fault-tolerance-in-a-distributed-system)

    “In this article, we discuss the concepts of dependability and fault tolerance in detail and explain how the Ably platform is designed with fault tolerant approaches to uphold its dependability guarantees.”
1. [Preparing to Issue 200 Million Certificates in 24 Hours](https://letsencrypt.org/2021/02/10/200m-certs-24hrs.html)

    “When we think about what essential infrastructure for the Internet needs to be prepared for though, we’re not thinking about normal days. We want to be prepared to respond as best we can to the most difficult situations that might arise. In some of the worst scenarios, we might want to re-issue all of our certificates in a 24 hour period in order to avoid widespread disruptions. That means being prepared to issue 200 million certificates in a day, something no publicly trusted CA has ever done.”
1. [Sonatype Spots 275+ Malicious npm Packages Copying Recent Software Supply Chain Attacks that Hit 35 Organizations](https://blog.sonatype.com/sonatype-spots-150-malicious-npm-packages-copying-recent-software-supply-chain-attacks)

    “Within 48 hours of the news reports, Sonatype’s automated malware detection systems, part of Nexus Intelligence, began flagging over 275 copycat npm packages published by different authors that imitate Alex Birsan’s proof-of-concept (PoC) research. We are actively seeing more of these packages coming in every few hours.” And people wonder why I’m not keen on Node.js
### Tools

1. [Introducing SSH zero trust, Identity aware TCP sockets](https://www.mysocket.io/post/introducing-ssh-zero-trust-identity-aware-tcp-sockets)

    Long time reader, Dustin Krysak pointed me to this. Seems VERY handy. I have not tried it yet.
1. [free copy](https://learn.launchdarkly.com/progressive-delivery/?utm_source=devopsish&utm_medium=news_pod&utm_campaign=21q1-newsletter)

    What if your team could safely ship software faster than your competitors? Learn how you can do just that with Progressive Delivery: The Next Iteration of Continuous Delivery. So you can reduce the risks of Continuous Delivery without sacrificing speed and deliver more value into your users’ hands. Download your  today! SPONSORED
1. [The world’s second-most popular desktop operating system isn’t macOS anymore](https://arstechnica.com/gadgets/2021/02/the-worlds-second-most-popular-desktop-operating-system-isnt-macos-anymore/)

    “Three quarters in a row of Chrome OS above macOS makes a trend. It’s safe to say that Chromebooks from PC makers like Acer, Asus, Dell, HP, and Lenovo now outsell Apple’s range of desktop Macs and laptop MacBooks. That said, 2020 was not a typical year.”
1. [Extending applications on Kubernetes with multi-container pods](https://learnk8s.io/sidecar-containers-patterns)

    “TL;DR: In this article you will learn how you can use the ambassador, adapter, sidecar and init containers to extend your apps in Kubernetes without changing their code.”
1. [Ansible Language - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=zbr.vscode-ansible)

    “Ansible extension aims to ease life of Ansible content creators by making easier to write Ansible playbooks, roles, collections, modules and plugins.”
1. [Introduction to GitOps on Kubernetes with Flux v2](https://blog.sldk.de/2021/02/introduction-to-gitops-on-kubernetes-with-flux-v2/)

    “Today we’re having a look at how to set up a GitOps pipeline for your Kubernetes cluster with Flux v2.”
1. [The New Space Age: IBM Develops A Unique, Custom Edge Computing Solution in Space](https://www.ibm.com/cloud/blog/ibm-develops-a-unique-custom-edge-computing-solution-in-space)

    “A ground-based solution on IBM Cloud will permit users to submit jobs, via a VPN connection, to the HPE ground systems that communicate securely with the HPE Spaceborne Computer-2 systems on the ISS, where Red Hat CodeReady Containers are installed.” Oh shit… Kubernetes is on the International Space Station!
1. [salesforce/cloudsplaining](https://github.com/salesforce/cloudsplaining)

    Cloudsplaining is an AWS IAM Security Assessment tool that identifies violations of least privilege and generates a risk-prioritized report.
1. [csantanapr/knative-kind](https://github.com/csantanapr/knative-kind)

    Knative on Kind (KonK)

### [ << Prev ](devopsweekly-205.md) ------------- [ Next >> ](devopsweekly-207.md)