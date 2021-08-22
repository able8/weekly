## [DevOps'ish 102](https://devopsish.com/102) - Sun Nov 18, 2018

It’s weird how things decay online. A 404 here, a failed git submodule update there can lead to weird things breaking. I noticed that happening with <a href="https://chrisshort.net/">chrisshort.net</a>. The site had gotten slower too despite it being a static site served through Cloudflare. I spent an evening looking for new <a href="http://gohugo.io">Hugo</a> themes. Testing their performance, figuring out what design would work best for me, and looked pleasant.

Last night, I ended up picking <a href="https://themes.gohugo.io/theme/beautifulhugo/">the updated version of the theme I was already using</a>. Today, I tuned it for a balance of features and performance. Then I went about tweaking the features to my liking. After cleaning up some 404s, I promoted the branch to master, and the CI was off and running.

At what point was I doing web development? I don’t know. But, that’s how I got my second big job in tech. Internet Explorer (back when it was on a coveted CD) came with <a href="https://en.wikipedia.org/wiki/Microsoft_FrontPage">FrontPage Express</a>. Say what you want about FrontPage. But, I knew how to use Microsoft Word and FrontPage Express gave me a WYSIWYG that showed me how to write HTML at the flip of a switch. Then I started writing web sites to W3C specs, discovered CSS, and ended up writing more web sites in Notepad then later <a href="https://www.jedsoft.org/jed/">jed</a> and <a href="https://www.vim.org/">vim</a>.

I was a senior in high school and needed a job after work. It was a fortuitous event that I needed internet access and they needed me. I took what I learned in my first tech job as an MIS technician and what I learned about Web 1.0 on my own and turned it into one of the best jobs I ever had. I learned about enterprise IT, Windows, and Linux. I read RFCs, ran internet services and loved my massive desk with two customized computers. One computer ran Windows 98 (seldom used) and my daily driver ran <a href="https://en.wikipedia.org/wiki/Red_Hat_Linux">Red Hat Linux</a>. The best part was the full T-1 speed late at night.

Then the dot-com bubble burst and off to the Air Force I went. Where I did do some upgrading of Air Combat Command’s websites. At no point though did I consider myself a web developer but, maybe that should change. I’ve done it longer than I’ve done DevOps. Ha! Who am I kidding? This is hard stuff and I’d rather leave it up to the experts. <strong>Web development is hard</strong>. Fight me.

<a href="https://www.gocd.org/test-drive-gocd/"><strong>Spin up a Continuous Delivery server in less than 5 minutes</strong></a><br/>Trying out new continuous delivery tools can take some time, so the team at GoCD created a quick start option for new users to spin up a GoCD server in under 5 minutes. <a href="https://www.gocd.org/test-drive-gocd/">Try now</a>. <em>SPONSORED</em>

<a href="https://www.influxdata.com/blog/writing-logs-directly-to-influxdb/?utm_campaign=devopsish&amp;utm_medium=partner&amp;utm_source=email&amp;utm_content=&amp;utm_term="><strong>Writing Logs Directly to InfluxDB</strong></a><br/>We consider logs to be just another form of time series data, and we want to give users better tools for ingesting and viewing that data. Learn how to use open source tools such as Telegraf and InfluxDB to take a metrics-first approach to log analysis. <em>SPONSORED</em>

<a href="https://info.signalsciences.com/devsecops-finding-the-adversaries-in-our-midst-webinar?utm_medium=newsletter&amp;utm_source=devopsish"><strong>Webinar with Shannon Lietz: DevSecOps: Finding the Adversaries in Our Midst</strong></a><br/>Discover how you can get greater visibility into how and where you’re being attacked, and use that visibility to build, create, and deploy sites that protect what your business cares about. <em>SPONSORED</em>

### People

1. [How old were you when you first started using Linux?](https://opensource.com/article/18/11/linux-age-poll)

     Old or young: New Linux users come in all ages.
1. [The great “DevOps engineer” title debate](https://enterprisersproject.com/article/2018/11/what-s-job-title-which-we-call-devops-engineer)

     “DevOps engineer” is currently the most recruited engineering job. But what does this job title mean exactly – and do we need it at all?
1. [We all love open source, but hiring based on contributions could be harmful](https://about.gitlab.com/2018/11/16/hiring-based-on-open-source-contributions-could-be-harmful/)

     An industry expert from Indeed says it’s a bad idea to make hiring decisions based on GitHub activity.
1. [Google Cloud CEO Diane Greene is out, to be replaced by former Oracle exec Thomas Kurian](https://www.cnbc.com/2018/11/16/google-cloud-ceo-greene-being-replaced-by-former-oracle-exec-kurian.html)

     Diane Greene became head of Google’s cloud three years ago after her start-up was acquired. She’s being replaced by Thomas Kurian, who was previously president of Oracle.
1. [What Diane Greene’s Departure Means for Google Cloud](https://www.wired.com/story/what-diane-greenes-departure-means-google-cloud/)

     I’m sure this change will bring with it a shift in how Google Cloud attacks enterprise deals.
1. [The true cost of a new employee: compensation calculator for startups](https://serverless.com/blog/true-cost-employee-calculator/)

    
1. [VA Officials Reluctant to Say When G.I. Benefit Delays Will Be Fixed](https://www.nextgov.com/it-modernization/2018/11/va-officials-reluctant-say-when-gi-benefit-delays-will-be-fixed/152895/)

     When government IT fails, the people trying to get services they’ve paid for and are legally obligated to receive them are the ones that get hurt.
1. [Veterans Find New Roles in Enterprise Cybersecurity](https://www.darkreading.com/risk/veterans-find-new-roles-in-enterprise-cybersecurity/d/d-id/1333250)

     Facebook and Synack create programs to educate vets and grow employment opportunities while shrinking the cybersecurity talent gap.
1. [A Roundup of Venture Capital Firms Focused On Investing In Veterans](https://news.crunchbase.com/news/veteran-focused-investment-firms/)

    
1. [How to stand out, and get hired, at Grace Hopper Celebration](https://www.redhat.com/en/blog/how-stand-out-and-get-hired-grace-hopper-celebration)

     “Here’s how to stand out when you’re trying to leave GHC with opportunities to chart your own path in technology.”
1. [This Chemical Is So Hot It Destroys Nerve Endings—in a Good Way](https://www.wired.com/story/resiniferatoxin/)

     Wired stories like this give me a lot of hope. If you know anyone working with RTX, have them give me a call. I’ve got some periphery nerve damage they can test.
### Process

1. [Google Kubernetes Engine’s third consecutive day of service disruption](https://news.ycombinator.com/item?id=18428497)

     Running complex systems at enormous scale is hard work. If something goes wrong it has high chance of escalating into a very big problem very quickly. Thankfully, I was not impacted by this in a significant way because I do not put all my eggs in any one provider’s basket.
1. [Google goes down after major BGP mishap routes traffic through China](https://arstechnica.com/information-technology/2018/11/major-bgp-mishap-takes-down-google-as-traffic-improperly-travels-to-china/)

     Google says it doesn’t believe leak was malicious despite suspicious appearances.
1. [How a Nigerian ISP Accidentally Knocked Google Offline](https://blog.cloudflare.com/how-a-nigerian-isp-knocked-google-offline/)

     “Last Monday evening — 12 November 2018 — Google and a number of other services experienced a 74 minute outage.”
1. [Kubernetes Office Hours for 21 November](https://discuss.kubernetes.io/t/office-hours-for-21-november/3528)

     Join some of the great folks on the livestream asking and answering questions.
1. [KubeCon + CloudNativeCon Barcelona 2019 Call for Proposals (CFP) Is Open](https://www.cncf.io/blog/2018/11/16/kubecon-barcelona-2019-call-for-proposals-cfp-is-open/)

     Sharpen those pencils and submit to KubeCon Barcelona 2019!
1. [Cloud Native Computing Foundation Opens Inaugural KubeCon + CloudNativeCon China and Welcomes 53 New Members](https://www.cncf.io/announcement/2018/11/13/cncf-welcomes-53-new-members/)

     New members, including R&D China Information Technology, Shanghai Qiniu Information Technologies, Beijing Yan Rong Technology and Yuan Ding Technology, to accelerate cloud native adoption across global markets
1. [Kubernetes 202 — Making It Fully Operational](https://medium.com/uptime-99/kubernetes-202-making-it-fully-operational-7416e4bb15ab)

     This article should serve as a guide to taking a bare Kubernetes cluster and making it “fully operational.”
1. [3 new DevOps lessons learned](https://enterprisersproject.com/article/2018/11/3-new-devops-lessons-learned)

     Lawyers (yes, lawyers) can help you kill legacy systems - and more wisdom from the recent DevOps Enterprise Summit
1. [How to Get Buy-In for Addressing Technical Debt](https://blog.gitprime.com/buy-in-addressing-technical-debt/)

     “If anyone understands the immediate impacts of technical debt, it’s development teams. Engineers and engineering leaders experience firsthand the slowdowns, the death-by-a-thousand-paper-cuts, even in drastic cases the soul-sucking repercussions of unmanaged technical debt.”
1. [What is machine learning? We drew you another flowchart](https://www.technologyreview.com/s/612437/what-is-machine-learning-we-drew-you-another-flowchart/)

     It pretty much runs the world.
1. [Real World DevOps by Jeff Geerling](https://www.jeffgeerling.com/blog/2018/real-world-devops)

    
1. [IBM, ServiceNow team up to expedite customers’ multicloud migrations](https://siliconangle.com/2018/11/16/ibm-servicenow-team-expedite-customers-multicloud-migrations/)

     IBM Corp. is stepping up its pitch to be the preferred multicloud manager for enterprise businesses.
1. [FCC paves the way for improved GPS accuracy](https://www.theverge.com/2018/11/15/18097154/fcc-ajit-pai-gps-location-accuracy-galileo)

     “The commission approved an order that will allow consumer devices to tap into a European satellite system.” I wonder what kind of improvements we’ll see to timing sources as a result of this.
1. [Uber Joins the Linux Foundation as a Gold Member](https://eng.uber.com/linux-uber-gold-membership/)

     The Linux Foundation becomes the next stop on the Uber apology tour. It’ll work and it’ll better the company as a result. But, this should never have had to have happened.
1. [Spectre, Meltdown researchers unveil 7 more speculative execution attacks](https://arstechnica.com/gadgets/2018/11/spectre-meltdown-researchers-unveil-7-more-speculative-execution-attacks/)

     Systematic analysis reveals a range of new issues and a need for new mitigations. This. Will. Never. End.
1. [What you need to know about the GPL Cooperation Commitment](https://opensource.com/article/18/11/gpl-cooperation-commitment)

     The GPL Cooperation Commitment fosters innovation by freeing developers from fear of license termination.
1. [Tracking Down Exposed Kubernetes Instances in the Wild](https://www.twistlock.com/labs-blog/tracking-exposed-kubernetes-instances-wild/)

     “After gathering close to 80,000 servers and analyzing them, I found that about 1000 servers are affected by this problem.”
1. [Corporate America’s blockchain and bitcoin fever is over](https://www.axios.com/corporate-america-blockchain-bitcoin-fervor-over-fb13bc5c-81fd-4c12-8a7b-07ad107817ca.html)

     S&P 500 executives are dropping blockchain buzzwords less on earnings calls and during presentations to analysts and investors. Analysts are also asking about it less.
1. [Salesforce makes undisclosed “strategic investment” in Docker, companies will cross-sell MuleSoft and Docker Enterprise](https://www.geekwire.com/2018/salesforce-makes-undisclosed-strategic-investment-docker-companies-will-cross-sell-mulesoft-docker-enterprise/)

    👀
### Tools

1. [sr.ht, the hacker’s forge, now open for public alpha](https://drewdevault.com/2018/11/15/sr.ht-general-availability.html?utm_source=tldrnewsletter)

     “The 500 foot view is that sr.ht is a 100% free and open source software forge, with a hosted version of the services running at sr.ht for your convenience. Unlike GitHub, which is almost entirely closed source, and Gitlab, which is mostly open source but with a proprietary premium offering, all of sr.ht is completely open source, with a copyleft license.”
1. [The State of the Octoverse: top programming languages of 2018](https://blog.github.com/2018-11-15-state-of-the-octoverse-top-programming-languages/)

     TypeScript arrives with a bang. Ruby’s popularity is in a long decline still.
1. [Amazon Web Services in Plain English](https://www.expeditedssl.com/aws-in-plain-english)

     What all this AWS nonsense should really be called.
1. [Amazon S3 Block Public Access – Another Layer of Protection for Your Accounts and Buckets](https://aws.amazon.com/blogs/aws/amazon-s3-block-public-access-another-layer-of-protection-for-your-accounts-and-buckets/)

     See that bucket over there with the blinking red light? Might want to check that.
1. [Amazon announces Corretto, a free, production-ready distribution of OpenJDK with LTS](https://jaxenter.com/amazon-corretto-free-openjdk-lts-151872.html)

     Some people like their old versions of Java to be supported.
1. [30 Second Kubernetes Concepts Cheat Sheet](https://medium.com/hashmapinc/30-second-kubernetes-concepts-cheat-sheet-98ba813194cb)

     A quick-reference guide to the basic concepts and resources available in Kubernetes.
1. [Scaling CoreDNS in Kubernetes Clusters](https://coredns.io/2018/11/15/scaling-coredns-in-kubernetes-clusters/)

     A guide for tuning CoreDNS resources/requirements in Kubernetes clusters
1. [Announcing Knative v0.2 Release](https://medium.com/knative/https-medium-com-knative-v0-2-963f276af58e)

     Improved pluggability, autoscaling, stability, and performance
1. [chris-short.ansible_role_adguard_homeAdGuard Home](https://galaxy.ansible.com/chris-short/ansible_role_adguard_home)

     This Ansible Role will install  on target system(s). This shows off some of the new features we’ve been working on in Ansible Galaxy too.
1. [Integrating Ansible and Red Hat Enterprise Linux 8 Beta](https://www.ansible.com/blog/integrating-ansible-and-red-hat-enterprise-linux-8-beta)

     RHEL 8 represents another significant change in RHEL. Python isn’t generally available by default (it’s still there at /usr/libexec/platform-python).
1. [Monitoring With Prometheus Using Ansible](https://itnext.io/monitoring-with-prometheus-using-ansible-812bf710ef43)

     Get going with Prometheus using Ansible
1. [7 open source platforms to get started with serverless computing](https://opensource.com/article/18/11/open-source-serverless-platforms)

     Serverless computing is transforming traditional software development. These open source platforms will help you get started.
1. [4 tips for learning Golang](https://opensource.com/article/18/11/learning-golang)

     Arriving in Golang land: A senior developer’s journey.
1. [GoAWK, an AWK interpreter written in Go](https://benhoyt.com/writings/goawk/)

     “Summary: After reading The AWK Programming Language I was inspired to write an interpreter for AWK in Go. This article gives an overview of AWK, describes how GoAWK works, how I approached testing, and how I measured and improved its performance.”
1. [Azure/golua](https://github.com/Azure/golua)

     A Lua 5.3 engine implemented in Go
1. [thecasualcoder/tztail](https://github.com/thecasualcoder/tztail)

     tztail (TimeZoneTAIL) allows you to view logs in the timezone you want
1. [Why and How to Use Git LFS](https://dzone.com/articles/git-lfs-why-and-how-to-use)

     Learn how Git LFS, an open-source Git extension, will help you handle large repositories.
1. [HomelabOS](https://gitlab.com/NickBusey/HomelabOS)

     Keep being gross tech companies. We could totally decentralize the web again with an Ansible Playbook.
1. [Distribution Release: Raspbian 2018-11-13](https://distrowatch.com/index-mobile.php?newsid=10376)

     Nothing earth shattering from the server side of things (or so it seems).
1. [Running ASP.NET Core on minikube](https://itnext.io/running-asp-net-core-on-minikube-ad69472c4c95)

    
1. [SSH Agents In Depth](https://dev.to/samuyi/ssh-agents-in-depth-4116)

     A deep dive into how SSH agents work including some edge cases.
1. [Find “Toxicity/Profanity” of a comment using Google Perspective](https://dev.to/sandeshsuvarna/find-toxicityprofanity-of-a-comment-using-google-perspective-3m22)

     I need this built into my brain ASAP!

### [ << Prev ](sreweekly-101.md) ------------- [ Next >> ](sreweekly-103.md)