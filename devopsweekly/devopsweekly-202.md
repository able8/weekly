## [DevOps'ish 202](https://devopsish.com/202) - Sun Jan 24, 2021

Unpopular opinion alert (and <a href="https://devopsish.com/terms/#disclaimer">Disclaimer</a>)…

Call me old fashioned, but I thought two of the top tenets of open source were candor and goodwill. I thought it was good practice to contribute to a project before baking it into a product. This was often the case for open source friendly vendors. But, it feels like AWS came along and never got that memo. I feel like AWS has done a lot more taking and productizing (aka making AWS a trillion-dollar, with a T, business) than contributing back to open source.

They keep shooting themselves in the foot as they take more and more projects into their bevy of services. A couple of weeks ago, I saw a tweet of a dev finding their code in the codebase of an AWS project without any credit being given, among other license violations. Someone at AWS said they’d look into it. But, consumption without credit incidents keeps happening (this wasn’t the first such incident I’d observed). There’s a culture problem, it seems. Then AWS hires a journalist to cover its open source work. I feel like that doesn’t help its case at all either. It acknowledges awareness of a problem. Pay for play is a negative thing in the radio business. It’s duplicitous at best in the tech industry.

I get it. AWS and others think <a href="https://aws.amazon.com/blogs/opensource/stepping-up-for-a-truly-open-source-elasticsearch/">forking Elastic’s code</a> was inevitable and acceptable. Maybe? On a different timeline, with candor and goodwill sprinkled on it, this is never an issue because AWS would be submitting heaping helpings of PRs that Elastic would gladly accept. It gives them the same capabilities in their product that a hyperscale cloud provider has. AWS wouldn’t care because they rake in billions upon billions in profits every year. Sure they are technically assisting a competing product but, how many people using a non-AWS Elastic are also using higher-level AWS services? I imagine that’s a tiny overlap in a Venn diagram. AWS and Elastic cater to <em>slightly</em> different personas. AWS was the most prominent cloud six years ago and had enough market share and growth to stay that way. Instead of helping an open source project, they consumed it to make money. This gave birth to the <a href="https://opensource.org/node/1099">bullshit SSPL license</a>, and the rest is history. Major takeaway: open source isn’t a business model.

No one is right here, and customers pretty much get hosed on the software and services’ price anyway. <a href="https://www.elastic.co/blog/licensing-change">Elastic makes its versions</a> of ElasticSearch, LogStash, and Kibana. AWS makes a fork of the previously Apache-licensed code. Who wins here? Elastic and AWS. Developers don’t win because they potentially have to make up for deficiencies in one version or the other. I feel like this fork and SSPL drama could very well have been preventable. But, what do I know? I’ve just been around open source for most of my life. It all feels gross. Why? Because there’s more than enough for everyone to get a piece of the pie.

<strong>Note</strong>: I’m looking for an intern this summer to help with <a href="https://OpenShift.tv?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=202">OpenShift.tv</a> (live streaming). If you know anyone that may be interested, please ask them to apply. If they have questions, feel free to send them my way (<a href="https://twitter.com/ChrisShort">Twitter DMs</a>, <a href="https://t.me/ChrisShort">Telegram</a>). Please <a href="https://us-redhat.icims.com/jobs/83032/openshift.tv-associate-producer-internship/job?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=202">apply ASAP</a> as I’m already reviewing resumes this weekend.

### People

1. [5 Reasons Why Disability Issues Should Be A Higher Priority, Even Now](https://www.forbes.com/sites/andrewpulrang/2021/01/11/5-reasons-why-disability-issues-should-be-a-higher-priority-even-now/?sh=4051337342f8)

    Yes, please.
1. [Gigantic Asshole Ajit Pai Is Officially Gone. Good Riddance (Time of Your Life)](https://www.vice.com/en/article/bvxpja/gigantic-asshole-ajit-pai-is-officially-gone-good-riddance-time-of-your-life)

    The headline says it all.
1. [Home alarm tech backdoored security cameras to spy on customers having sex](https://arstechnica.com/information-technology/2021/01/home-alarm-tech-backdoored-security-cameras-to-spy-on-customers-having-sex/)

    Holy smokes… This is an embarrassing incident from a company that usually maintains a low key and rock-solid business.
1. [Why Tech Is Still Toxic for Women (and What to Do About it)](https://thenewstack.io/why-tech-is-still-toxic-for-women-and-what-to-do-about-it/)

    It’s a start.
1. [Why I don’t use Stack Overflow‍](https://dev.to/facundocorradini/why-i-don-t-use-stack-overflow-1f0l)

    There’s a lot wrong with Stack Overflow. A LOT.
1. [Scoop: Google is investigating the actions of another top AI ethicist](https://www.axios.com/scoop-google-is-investigating-the-actions-of-another-top-ai-ethicist-50030739-ea3d-4ea2-b452-c228b4fc9773.html)

    AGAIN?!? Sundar… Come on.
1. [Stop Making Excuses for Toxic Bosses](https://hbr.org/2021/01/stop-making-excuses-for-toxic-bosses)

    “If you’ve ever worked for a toxic boss, you know how damaging it can be. So should you forgive a manager who tries to make amends for their bad behavior? A new study shows that most abusive bosses care more about their social image than actually changing how they act.”
1. [Update on an employee matter](https://github.blog/2021-01-17-update-on-an-employee-matter/)

    HR has been shaken, not stirred at GitHub.
### Process

1. [FAA Files Reveal a Surprising Threat to Airline Safety: the U.S. Military’s GPS Tests](https://spectrum.ieee.org/aerospace/aviation/faa-files-reveal-a-surprising-threat-to-airline-safety-the-us-militarys-gps-tests)

    This is a horrific accident waiting to happen.
1. [Download today: Kubernetes security ebook - tips, tricks, best practicesDownload this ebook](https://security.stackrox.com/kubernetes-security-ebook-tips-tricks-best-practices.html?Source=DevOpsIsh&LSource=DevOpsIsh)

    The rapid adoption of Kubernetes to manage containerized workloads is driving great efficiencies in application development, deployment, and scalability. However, when security becomes an afterthought, you risk diminishing the greatest gain of containerization - agility.  to learn how to (1) build secure images and prevent untrusted/vulnerable code, (2) configure RBAC, network policies, and runtime privileges, (3) detect unauthorized runtime activity, and (4) secure your Kubernetes infrastructure components such as the API server. SPONSORED
1. [In the Digital Economy, Your Software Is Your Competitive Advantage](https://hbr.org/2021/01/in-the-digital-economy-your-software-is-your-competitive-advantage)

    If Harvard Business Review is trying to tell you this, maybe your execs will finally listen to you.
1. [Drupal’s journey from dorm-room project to billion-dollar exit](https://techcrunch.com/2021/01/22/drupals-journey-from-dorm-room-project-to-billion-dollar-exit/)

    “Twenty years ago Drupal and Acquia founder Dries Buytaert was a college student at the University of Antwerp. He wanted to put his burgeoning programming skills to work by building a communications tool for his dorm. That simple idea evolved over time into the open-source Drupal web content management system, and eventually a commercial company called Acquia built on top of it.”
1. [Monitoring as Code: What It Is and Why You Need It](https://thenewstack.io/monitoring-as-code-what-it-is-and-why-you-need-it/)

    “Integrate Observability into the CI/CD Pipeline with Monitoring as Code”
1. [Remediation and Hardening Strategies for Microsoft 365 to Defend Against UNC2452](https://www.fireeye.com/blog/threat-research/2021/01/remediation-and-hardening-strategies-for-microsoft-365-to-defend-against-unc2452.html)

    I got out of the Windows sysadmin business because of stuff like this.
### Tools

1. [Why I migrated from Traefik to Caddy](https://p1ngouin.com/posts/why-i-migrated-from-traefik-to-caddy)

    The process and research behind ones choice to use Caddy. I like this article because it walks through the warts and implementation.
1. [Red Hat Announces No-Cost RHEL For Small Production Environments](https://www.phoronix.com/scan.php?page=news_item&px=Red-Hat-RHEL-No-Cost-16-Systems)

    “The 16 system limit handling is still done through the registration-based developer program and does include deployments to major cloud providers. These developer program changes begin on 1 February.”
1. [PostgreSQL on ARM-based AWS EC2 Instances: Is It Any Good?](https://www.percona.com/blog/2021/01/22/postgresql-on-arm-based-aws-ec2-instances-is-it-any-good/)

    “We decided to take an independent look at the price/performance of the new instances from the standpoint of running PostgreSQL.”
1. [HTML and CSS still isn’t about painting with code](https://christianheilmann.com/2021/01/21/teaching-html-and-css/)

    This article is so good. Imitating art is not programming.
1. [Painless services: implementing serverless with rootless Podman and systemd](https://www.redhat.com/en/blog/painless-services-implementing-serverless-rootless-podman-and-systemd)

    Title pretty much explains it all. But, it’s a guide to using systemd and Podman to make serverless services. Kinda nifty if you don’t want to muck with a framework and utilize closer to core abstractions.
1. [Kubernetes Begins Year With A Bang — And You Can Expect More](https://news.crunchbase.com/news/kubernetes-begins-year-with-a-bang-and-you-can-expect-more/)

    ”‘About 18 to 24 months ago, you really saw Kubernetes take over,’ said Ubaid Dhiyan, a director at investment bank Union Square Advisors. ‘Now companies are building on top of Kubernetes.’”
1. [Why I like Rust - and why it’s annoying](https://jrhawley.ca/2021/01/22/why-i-like-rust)

    “They don’t approach the situation with the perspective of ‘what do we stand to gain from, and what are the costs that come with, switching to Rust?’.”
1. [Dnsmasq, used in only a million or more internet-facing devices globally, patches not-so-secret seven spoofing, hijacking flaws](https://www.theregister.com/2021/01/20/dns_cache_poisoning/)

    Patch your everything.
1. [Introduction to eBPF](https://oswalt.dev/2021/01/introduction-to-ebpf/)

    “[U]nlike a lot of hyped-up technology buzzwords, this one’s momentum doesn’t seem to be unwarranted, or even ahead of the curve.”
1. [fireeye/Mandiant-Azure-AD-Investigator](https://github.com/fireeye/Mandiant-Azure-AD-Investigator)

    
1. [rqlite/rqlite](https://github.com/rqlite/rqlite)

    The lightweight, distributed relational database built on SQLite.
1. [treenotation/dumbdown](https://github.com/treenotation/dumbdown)

    The dumb alternative to markdown. The keyword for title is title.

### [ << Prev ](devopsweekly-201.md) ------------- [ Next >> ](devopsweekly-203.md)