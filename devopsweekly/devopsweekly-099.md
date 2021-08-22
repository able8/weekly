## [DevOps'ish 099](https://devopsish.com/099) - Sun Oct 28, 2018

This was an incredibly busy week in tech news and events. First, <a href="https://allthingsopen.org/">AllThingsOpen</a> was this week. For me, this included the annual <a href="https://opensource.com/">opensource.com</a> Moderators meeting in Red Hat Tower. The conference was amazing but, I feel like it might have outgrown the Raleigh Convention Center. 4,100 people registered to attend AllThingsOpen and I’m pretty sure half of them tried to get to two talks I wasn’t able to attend due to crowds. I got to spend and even monopolize time with some truly amazing folks. If you were at AllThingsOpen and we weren’t able to connect my deepest apologies.

GitHub had two significant outages this week that are discussed later. It’s amazing to me how technology professionals treat these complex services they use every day as if they’re air, water, and light. If you think running your code is hard, imagine how much harder it is to run the thing your code relies on. GitHub’s week definitely ended on a high note though. The “<a href="https://blog.github.com/2018-10-26-github-and-microsoft/">Microsoft acquisition of GitHub is complete. 🎉</a>“

Google admitted to <a href="https://www.nytimes.com/2018/10/25/technology/google-sexual-harassment-andy-rubin.html">giving golden handshakes to folks on their way out the door while keeping silent about misconduct claims</a>. This is a horrendous display of stereotypical tech culture from a company that once marketed itself as the exact opposite of. As I asked on Twitter:

More details and reaction regarding Google’s gross behavior follows (they’re also doing some shady things with hiring too). I’m considering abandoning as many Google products as I can get away with at this point (and this is after pretty much moving all my cloud resources over to Google Cloud).

SQLite is making light of Code of Conducts <a href="https://www.theregister.co.uk/2018/10/22/sqlite_code_of_conduct/">by using its own Code of Conduct to mock them</a>. Then after a ton of blowback from the community, that original Code of Conduct was renamed to <a href="https://sqlite.org/codeofethics.html">Code of Ethics of the Project Founder</a> and SQLite has adopted a <a href="https://sqlite.org/codeofconduct.html">new Code of Conduct</a>, the Mozilla Community Participation Guidelines.

This is on top of beefing up the DevOps’ish <a href="../terms/">Terms of Service</a>, implementing a <a href="../privacy/">Privacy Policy</a> at the request of a possible reader, tweaking the production process for the newsletter to keep it from getting too long, <a href="https://www.ansible.com/blog/author/chris-short">working a full-time job</a>, and <a href="http://bit.ly/2qd1jPJ">being a dad</a>. Continue reading for more of this week’s DevOps, cloud native, and open source news.

<a href="https://info.thoughtworks.com/Actionable_CD_Metrics.html"><strong>Webinar: Actionable Continuous Delivery Metrics</strong></a><br/>Want to deliver faster? Join our free webinar: Actionable Continuous Delivery Metrics. Gain insights into software delivery pipeline and learn to use metrics to improve your path to production. <em>SPONSORED</em>

<a href="https://raygun.com/"><strong>Finally, it’s time for actionable APM!</strong></a><br/>Raygun delivers what New Relic and other APM tools can’t: an intuitive and easy to use interface with unparalleled detail into app transactions. <a href="https://raygun.com/">Now available for .NET.</a> <em>SPONSORED</em>

<a href="https://info.signalsciences.com/crash-course-kubernetes-security?utm_medium=newsletter&amp;utm_source=devopsish"><strong>Webinar: Kubernetes and Security Crash Course</strong></a><br/>Kubernetes is so powerful, it sometimes seems like magic. However, if not setup correctly, it can lead to serious security issues. In this webinar, learn how to avoid these issues and unlock the full benefits of Kubernetes. <em>SPONSORED</em>

### People

1. [How Google Protected Andy Rubin, the ‘Father of Android’](https://www.nytimes.com/2018/10/25/technology/google-sexual-harassment-andy-rubin.html)

     The internet giant paid Mr. Rubin $90 million and praised him, while keeping silent about a misconduct claim.
1. [An Insider’s View of How Alphabet Is Dealing With Harassment](https://www.bloomberg.com/news/videos/2018-10-25/an-insider-s-view-of-how-alphabet-is-dealing-with-harassment-video)

     Liz Fong Jones, an engineer at Alphabet Inc., discusses how the company has been handling sexual harassment claims. She speaks with Bloomberg’s Emily Chang on “Bloomberg Technology.” (Source: Bloomberg)
1. [Google CEO admits company had a sexual harassment problem — says it has fired 48 employees for sexual misconduct](https://www.cnbc.com/2018/10/25/google-ceo-memo-says-48-fired-for-sexual-misconduct.html)

     Google CEO Sundar Pichai sent an email to employees in response to a report in The New York Times about sexual misconduct at the company.
1. [Linus Torvalds is back in charge of Linux](https://www.zdnet.com/article/linus-torvalds-is-back-in-charge-of-linux/)

     After a few weeks off to reconsider his role in the Linux community, Linus Torvalds is back in the saddle.
1. [Silicon Valley’s dirty secret: Using a shadow workforce of contract employees to drive profits](https://www.cnbc.com/2018/10/22/silicon-valley-using-contract-employees-to-drive-profits.html)

     “This year at Google, contract workers outnumbered direct employees for the first time in the company’s 20-year history.”
1. [Why the NSA Called Me After Midnight and Requested My Source Code](https://medium.com/datadriveninvestor/why-the-nsa-called-me-after-midnight-and-requested-my-source-code-f7076c59ab3d)

     The story behind my top secret coffee cup
1. [The case for all-remote companies](https://about.gitlab.com/2018/10/18/the-case-for-all-remote-companies/)

     Remote teams offer flexibility, reduce company costs, and increase productivity.
1. [DevOps Salaries Hit New High](https://www.technative.io/devops-salaries-hit-new-high/)

     “DevOps is growing at a remarkable pace” Ya don’t say?
1. [Open Source Veterans Join source{d} as advisorssource{d}](https://medium.com/sourcedtech/open-source-veterans-join-source-d-as-advisors-2b48717b1748)

     Wow! What a great group of folks to have on the side of
1. [What is a site reliability engineer and why you should consider this career path](https://opensource.com/article/18/10/what-site-reliability-engineer)

     If you want a challenging, in-demand role that goes beyond DevOps, consider becoming an SRE.
### Process

1. [SQLite Code of Conduct Sparks Outrage in Dev Community](https://dzone.com/articles/sqlite-code-of-conduct-sparks-outrage-in-the-dev-c)

     Members of the dev community are shocked by the latest SQLite Code of Conduct found on their website.
1. [Serverless DevOps](https://www.serverlessops.io/download-the-serverless-devops-ebook)

     What do we do when the server goes away?
1. [Chaos Monkey Guide for Engineers](https://www.gremlin.com/chaos-monkey/)

     “The Chaos Monkey Guide for Engineers is a full how-to of Chaos Monkey, including what it is, its origin story, its pros and cons, its relation to the broader topic of Chaos Engineering, and much more.”
1. [Making the GPL more scary](https://lwn.net/SubscriberLink/768670/21b0a9ecf1337105/)

     Strange times in the world of software licenses.
1. [Fed up with cloud giants ripping off its database, MongoDB forks new ‘open-source license’](https://www.theregister.co.uk/2018/10/16/mongodb_licensning_change/)

     Paperwork demands code from internet goliaths
1. [Unexpected MySQL database meltdown fingered in GitHub’s 24-hour website wobble](https://www.theregister.co.uk/2018/10/23/github_outage_ends/)

     GitHub had a ROUGH beginning of the week.
1. [Ep. #6, Customer Reliability Engineering with Google’s Liz Fong-Jones](https://www.heavybit.com/library/podcasts/o11ycast/ep-6-customer-reliability-engineering-with-googles-liz-fong-jones/)

     In episode 6 of O11ycast, Charity and Rachel talk to Google developer advocate Liz Fong-Jones about ways to build systems to be more transparent and explainable.
1. [Documentation as an Open Source Practice](https://blog.digitalocean.com/documentation-as-an-open-source-practice/)

     Develop Documentation Like Code
1. [GitHub is FedRAMP Authorized](https://blog.github.com/2018-10-24-github-is-fedramp-authorized/)

     That didn’t take long.
1. [Amazon’s move off Oracle caused Prime Day outage in one of its biggest warehouses, internal report says](https://www.cnbc.com/2018/10/23/amazon-move-off-oracle-caused-prime-day-outage-in-warehouse.html)

     Amazon’s move off Oracle’s database software was the main reason for an outage in one of its biggest warehouses on Prime Day, according to internal documents obtained by CNBC.
1. [Why IT Ticketing Systems Don’t Work with Microservices](https://blog.getambassador.io/why-it-ticketing-systems-dont-work-with-microservices-18e2be509bf6)

    
### Tools

1. [How do you keep up with Kubernetes?](https://dev.to/petermbenjamin/how-do-you-keep-up-with-kubernetes-2209)

     A good starting point on how to manage your Kubernetes knowledge.
1. [Kubernetes: Getting Started](https://cshort.co/start-k8s)

     This Kubernetes Getting Started guide I wrote almost a year ago is still getting a share every once in a while.
1. [Mirantis Takes Kubernetes to the Telco Edge](https://www.lightreading.com/the-edge/mirantis-takes-kubernetes-to-the-telco-edge/d/d-id/747077)

    
1. [Why Is Securing Kubernetes so Difficult?](https://blog.giantswarm.io/why-is-securing-kubernetes-so-difficult/)

     The first of a series of articles Giant Swarm is rolling out on Kubernetes security
1. [Atlassian’s overhaul of Jira is complete, with a more user-friendly interface for the cloud-native era](https://www.geekwire.com/2018/atlassians-overhaul-jira-complete-user-friendly-interface-cloud-native-era/)

     Anyone using it yet?
1. [gRPC-Web is going GA](https://www.cncf.io/blog/2018/10/24/grpc-web-is-going-ga/)

    
1. [Build a Single-Page App with Go and Vue](https://developer.okta.com/blog/2018/10/23/build-a-single-page-app-with-go-and-vue)

     Whoa… This doesn’t seem as hard as I thought it might be.
1. [Jerry Hargrove - Cloud Diagrams & Notes](https://www.awsgeek.com/)

    
1. [The Challenges Behind Rolling Out Security Updates To Your Docker Images](https://eng.lyft.com/the-challenges-behind-rolling-out-security-updates-to-your-docker-images-86106de47ece)

    
1. [HashiCorp Product Announcements at HashiConf 2018](https://www.hashicorp.com/blog/hashicorp-product-announcements-at-hashiconf-2018)

    
1. [Linux Kernel /etc/sysctl.conf Security Hardening](https://www.cyberciti.biz/faq/linux-kernel-etcsysctl-conf-security-hardening/)

    
1. [namreg/godown](https://github.com/namreg/godown)

     Distributed, fault-tolerant key-value storage written in go.
1. [k88hudson/git-flight-rules](https://github.com/k88hudson/git-flight-rules)

     Flight rules for git
1. [IBM/mac-ibm-enrollment-app](https://github.com/IBM/mac-ibm-enrollment-app)

     The Mac@IBM enrollment app makes setting up macOS with Jamf Pro more intuitive for users and easier for IT.

### [ << Prev ](sreweekly-98.md) ------------- [ Next >> ](sreweekly-100.md)