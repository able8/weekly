## [DevOps'ish 169](https://devopsish.com/169) - Sun Mar 1, 2020

The first CNCF Air Gapped Working Group meeting happened this week. I was fortunate enough to be able to attend. Air gapped or disconnected environments have a few variants. But, the air gaps I’m most familiar with are those between unclassified and classified US government systems (TEMPEST, EMSEC, COMSEC, etc.). Also, given the attendance, there is a lot of interest from folks across the public and private sectors. If you’re interested in getting involved, <a href="https://lists.cncf.io/g/cncf-sig-app-delivery/topic/airgap_working_group_meeting/71568152">we’re figuring out what the cadence of the meetings will be</a>. The current scope of our work is to provide a “cookbook” for maintaining an air gapped cloud native environment. Join the <a href="https://lists.cncf.io/g/cncf-sig-app-delivery/">CNCF SIG App Delivery mailing list</a> and keep an eye out for messages to see how to get involved.

The 2020 election for the <a href="https://wiki.opensource.org/bin/Main/OSI+Board+of+Directors/">Open Source Initiative Board of Directors</a> is rapidly approaching. I nominated myself for a seat on the board as an Individual Member. I spoke with a few friends in the months before submitting my nomination. It sounds like OSI could use my help <a href="http://lists.opensource.org/pipermail/license-discuss_lists.opensource.org/2020-February/021350.html">given the week they’ve had</a>. Open source software has done nothing but provide opportunity after opportunity in my life. It should be cared for and maintained so that everyone willing can benefit from it as much as I have (hopefully more). If you’re an OSI member, <a href="https://wiki.opensource.org/bin/Main/OSI+Board+of+Directors/Board+Member+Elections/2020+Individual+and+Affiliate+Elections/Short2020">please considering me for a seat on the board</a>. If you’re not an OSI member, you have until 1 March 2020 to join to vote. However, if you work with open source software, <a href="https://opensource.org/join">you should be an OSI member</a>.

### People

1. [Kubernetes Contributor Summit Amsterdam Schedule Announced](https://kubernetes.io/blog/2020/02/18/contributor-summit-amsterdam-schedule-announced/)

    “On top of the presentations, there will be a dedicated Docs Sprint as well as the New Contributor Workshop 101 and 201 Sessions. All told, we will have five separate rooms of content throughout the day on Monday.”
1. [Ansible Contributor Summit Europe](https://www.ansible.com/blog/announcement-ansible-contributor-summit-europe)

    “A European Contributor Summit will be held in Gothenburg, Sweden, ahead of the usual summit at AnsibleFest! On March 29 we’ll be welcoming new and old contributors alike. So if you already contribute to Ansible, or would like to, but don’t know how or where to start, this event is for you.”
1. [The Horrifically Dystopian World of Software Engineering Interviews](https://www.jarednelsen.dev/posts/The-horrifically-dystopian-world-of-software-engineering-interviews)

    “Why are we measuring people and not learning about people?”
1. [📕 The 25 most recommended programming books of all-time](https://www.daolf.com/posts/best-programming-books/)

    This is an excellent list of books. I’d like to see a similar method for DevOps books.
1. [DOJ Plans to Strike Against Encryption While the Techlash Iron Is Hot](https://cyberlaw.stanford.edu/blog/2020/02/doj-plans-strike-against-encryption-while-techlash-iron-hot)

    The US Department of Justice intends to strike while the iron is hot in the encryption war as public opinion of tech companies wanes. This is a pretty serious deal and the US government might get away with it given how dependent the US is upon mass surveillance.
### Process

1. [Container Security - Nobody Knows What It Means But It’s Provocative](https://capsule8.com/blog/container-security-nobody-knows-what-it-means-but-its-provocative/)

    “[T]his post helps delineate the different areas of container security and why they’re each important. As a community — not just security, but also embracing our friends in DevOps — we must work to secure the containers lest we become contained by our own failures.”
1. [Starting work on Operator Working Group](https://lists.cncf.io/g/cncf-sig-app-delivery/topic/71453445#151)

    CNCF is also starting up an Operator working group. I hope to be able to participate in this as it’s going to be a fun space for a while.
1. [How to solve the DevOps vs. ITSM culture clash](https://opensource.com/article/20/2/devops-vs-itsm)

    I know I faced this problem head on in a few organizations. “Declare a truce between ITSM and DevOps within your organization by learning these lessons.”
1. [Open Source-based Cloud Observabilitycloud observability platformfree trial today](https://logz.io/freeshirt/?utm_source=podcast&utm_medium=devopish&utm_campaign=freeshirt)

    If you are like many of your friends in DevOps, you might prefer using open source solutions. Logz.io’s  is designed to help deliver the open source experience you love, at the scale you need. Log management and Cloud SIEM based on ELK, and Infrastructure monitoring based on Grafana. All on one unified platform to help you correlate quickly between logs and metrics. Sign up for a  and receive a Logz.io t-shirt! SPONSORED
1. [Mirantis Acquires Kubernetes Assets from Kontena](https://containerjournal.com/topics/container-ecosystems/mirantis-acquires-kubernetes-assets-from-kontena/)

    “Kontena role-based access control (RBAC) and multi-cluster management tools will find their way into Docker Kubernetes Service (DKS) and Universal Control Plane (UCP).”
1. [Advanced Persistence Threats: The Future of Kubernetes Attacks](https://www.rsaconference.com/usa/agenda/advanced-persistence-threats-the-future-of-kubernetes-attacks-3)

    “Let’s explore the dark corners of clusters and shine a light on several new advanced attacks on Kubernetes, and learn how to detect and prevent them using practical, proven methods.”
1. [DevOps: 5 tips on how to manage stronger remote teams](https://enterprisersproject.com/article/2020/2/devops-how-manage-remote-teams-tips)

    “Are you leading or working in a DevOps team where some people work remote? Building a healthy culture presents some challenges - but experts say remote teams can become masters of collaboration and flexibility”
### Tools

1. [Your new web notifications experience is here](https://github.blog/2020-02-25-your-new-web-notifications-experience-is-here/)

    Pretty sure I already f’ed this up. I need to stop using GitHub on mobile. But, I’m hopeful I can re-manage my notifications as opposed to having to be so beholden to email.
1. [Free Software Foundation Aims To Launch Code Hosting / Collaboration Platform This Year](https://www.phoronix.com/scan.php?page=news_item&px=FSF-Forge-2020-Platform)

    “The Free Software Foundation’s lead contender right now is using Pagure as the basis for their platform, but have some reservations over it since JavaScript is required for a pleasant experience. Gitea and Sourcehut are other options they are also exploring but have pros/cons for all of them.”
1. [On-Demand Container Scanning API](https://jerrygamblin.com/2020/02/23/on-demand-container-scanning-api/)

    “scan.vulnerablecontainers.org is an open python API built using Trivy, Flask, Gunicorn, and Nginx that for now has two public endpoints (more endpoints and tools coming). From the start, I designed it to be easy to use in the browser or on the command line for integration with CI/CD.”
1. [Everyone might be a cluster-admin in your Kubernetes cluster](https://www.jeffgeerling.com/blog/2020/everyone-might-be-cluster-admin-your-kubernetes-cluster)

    “What the installation guide doesn’t tell you is that this manifest grants cluster-admin privileges to every single Pod in the default namespace!” This is going to become a big problem pretty quick. You need to make sure you’re applying RBAC and security policies to anything you pull down from the web.
1. [What happens behind the scenes of a rootless Podman container?](https://www.redhat.com/sysadmin/behind-scenes-podman)

    “Have you ever wondered what happens behind the scenes of a rootless Podman container? Let’s walk through an example.”
1. [Let’s Encrypt Has Issued a Billion Certificates](https://letsencrypt.org/2020/02/27/one-billion-certs.html)

    Huge thank you to the folks at Let’s Encrypt and congratulations on building one of the most fantastic infrastructure services on the internet.
1. [The Jellyfish-Inspired Database Under AWS Block Storage](https://www.nextplatform.com/2020/02/18/the-jellyfish-inspired-database-under-aws-block-storage/)

    “While it is highly unlikely that AWS will ever open source Physalia, it is possible that it will inspire others to re-create elements of its architecture that will be open source, much as has happened with many hyperscaler and cloud builder technologies in the past.”
1. [Monitor Your Docker Containers More Effectively](https://www.datadoghq.com/dg/monitor/docker-benefits-ts/?utm_source=Advertisement&utm_medium=Advertisement&utm_campaign=DevOpsish-Newsletter02&utm_content=Docker)

    See across all of your Docker containers, servers, apps, and services in one place with Datadog’s custom dashboards, distributed tracing, and 400+ vendor-backed integrations. Start your free trial today, create one dashboard, and receive a free Datadog t-shirt. SPONSORED
1. [Crisis? What crisis? Netflix opensources incident management framework](https://devclass.com/2020/02/25/crisis-what-crisis-netflix-opensources-incident-management-framework/)

    “Netflix has open sourced its crisis management framework, giving organisations a standardised way to manage incidents and get back to a state of chill.” I see what you did there. 👀
1. [Multicluster Kubernetes with Service Mirroring](https://linkerd.io/2020/02/25/multicluster-kubernetes-with-service-mirroring/)

    “With service mirroring in place, the full observability, security, and routing features of Linkerd apply uniformly to both in-cluster and cluster-calls, and the application does not need to distinguish between those situations.”
1. [Tig](https://jonas.github.io/tig/)

    Text-mode interface for Git
1. [Implementing A Reverse Proxy Server In Kubernetes Using The Sidecar](https://www.magalix.com/blog/implemeting-a-reverse-proxy-server-in-kubernetes-using-the-sidecar-pattern)

    DIY sidecar injection to help you understand the concept of having multiple containers in a single pod to form a fully functional web app.

### [ << Prev ](devopsweekly-168.md) ------------- [ Next >> ](devopsweekly-170.md)