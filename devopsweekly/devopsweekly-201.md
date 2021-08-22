## [DevOps'ish 201](https://devopsish.com/201) - Sun Jan 17, 2021

I lost a co-worker from the Ansible team this week. I’ve been struggling to get past the insanity of people younger than dying. 2021 is off to a real shit start. But, I think the biggest tech story of the week comes from Elastic. Keep reading for the details on Elastic’s idiocy. Here’s your weekly reminder that open source isn’t a business model, though.

But, there’s been a moment of justice for those here in Michigan who were impacted by the <a href="https://en.wikipedia.org/wiki/Flint_water_crisis?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=201">Flint Water Crisis</a>. <a href="https://arstechnica.com/tech-policy/2021/01/ex-michigan-governor-indicted-for-willful-neglect-in-flint-water-crisis/?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=201">Ex-Michigan governor indicted for ‘willful neglect’ in Flint water crisis</a>. Here’s the <a href="https://www.michigan.gov/documents/ag/GJ_Indictment_-_Snyder_R_712955_7.pdf?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=201">grand jury indictment</a> of Former Michigan Governor, Rick Snyder. The Judicial System better not mess this up.

I also learned about the term <a href="https://en.wikipedia.org/wiki/Sealioning?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=201">Sealioning</a> this week. I’ve seen it done before but did not know it had a definition that I could shut folks down with. Nice!

<strong>Note</strong>: I’m looking for an intern this summer to help with <a href="https://OpenShift.tv?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=201">OpenShift.tv</a> (live streaming). If you know anyone that may be interested, please ask them to apply. If they have questions, feel free to send them my way (<a href="https://twitter.com/ChrisShort">Twitter DMs</a>, <a href="https://t.me/ChrisShort">Telegram</a>). Please <a href="https://us-redhat.icims.com/jobs/83032/openshift.tv-associate-producer-internship/job?utm_source=newsletter&amp;utm_medium=devopsish&amp;utm_campaign=201">apply ASAP</a> as I’m already reviewing resumes this weekend.

### People

1. [GitHub still won’t explain if it fired someone for saying ‘Nazi,’ and employees are pissed](https://www.theverge.com/2021/1/15/22232766/github-employees-protest-jewish-employee-firing-warn-nazi?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    “Now in protest, they’re using Slack to call Nazis what they are” I’m totally dumbfounded by this. A Jewish employee tells his Jewish friends and coworkers to be safe during the insurrection of 2021 because nazi flags are clearly on display then gets fired. “Now, GitHub workers are saying ‘Nazi’ repeatedly in Slack, in regards to the US Capitol rioters, to protest what is being perceived as unfair treatment.” Fix this GitHub, I know GitHub employees are reading this (Pro tip: you don’t even have to have a position in the matter, just highlight this and forward it up).
1. [I’m an Impostor - Incarceration and Living a Lie](https://theworst.dev/im-an-impostor/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    “Every day, I walk around telling little lies so I can project this false image of myself. I would like to tell you this is a story about how I feel like I don’t know enough, and then I realized that people don’t know shit either, but this isn’t that story.” This is an amazing story of hitting rock bottom and bouncing back. Now, Kurt Kemple is trying to help inmates while bringing awareness to us in tech about those in incarceration. Mad props to you, Kurt.
1. [Confronting our own racism as white engineering managers](https://leaddev.com/diversity-inclusion/confronting-our-own-racism-white-engineering-managers?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    “But there is very little formal management advice for White people about recognizing and understanding our racism in managing and leading Black and other employees of color.” This article makes the argument that if you’re a white manager you have not only the organization giving you power but, the echoes and often drum beat of White Supremacy as well. I would highly encourage sharing this with your managers. Tell them Chris Short told you to.
1. [Millions Flock to Telegram and Signal as Fears Grow Over Big TechRedditSignal went down on Friday](https://www.nytimes.com/2021/01/13/technology/telegram-signal-apps-big-tech.html?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    People dropping WhatsApp like a bad habit. Join DevOps’ish on . We’re looking into alternative platforms as well in case these become unstable.  under the deluge of new users.
1. [Intel lured new CEO Pat Gelsinger with a package valued at $116 millionIntel tag here on DevOps’ish](https://www.oregonlive.com/silicon-forest/2021/01/intel-lured-new-ceo-pat-gelsinger-with-a-package-valued-at-116-million.html?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    For a cool $116 million you too can have a smooth-talking, boisterous, and hopefully transformative CEO. I’m considering this Intel’s hail mary against ARM and RISC-V after MANY bad years of not so great outcomes. Check the  if you need a reminder.
1. [Dropbox to cut 11% of its global workforce](https://www.cnbc.com/2021/01/13/dropbox-to-cut-11percent-of-its-global-workforce.html?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    This sucks.
### Process

1. [Secondary stock sale to Sequoia values Zapier at more than $4 billion](https://siliconflorist.com/2021/01/15/secondary-stock-sale-to-sequoia-values-zapier-at-more-than-4-billion/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    I like Zapier and use it for a lot of things that power this newsletter. While an interesting way to invest in a company, it’s not uncommon, and Sequoia knows what it’s doing. I see this as a very good sign for Zapier’s future.
1. [Download today: Kubernetes security ebook - tips, tricks, best practicesDownload this ebook](https://security.stackrox.com/kubernetes-security-ebook-tips-tricks-best-practices.html?Source=DevOpsIsh&LSource=DevOpsIsh)

    The rapid adoption of Kubernetes to manage containerized workloads is driving great efficiencies in application development, deployment, and scalability. However, when security becomes an afterthought, you risk diminishing the greatest gain of containerization - agility.  to learn how to (1) build secure images and prevent untrusted/vulnerable code, (2) configure RBAC, network policies, and runtime privileges, (3) detect unauthorized runtime activity, and (4) secure your Kubernetes infrastructure components such as the API server. SPONSORED
1. [[SECURITY] CVE-2021-24122 Apache Tomcat Information Disclosure](http://mail-archives.apache.org/mod_mbox/www-announce/202101.mbox/%3Cf3765f21-969d-7f21-e34a-efc106175373%40apache.org%3E)

    “When serving resources from a network location using the NTFS file system it was possible to bypass security constraints and/or view the source code for JSPs in some configurations. The root cause was the unexpected behaviour of the JRE API File.getCanonicalPath() which in turn was caused by the inconsistent behaviour of the Windows API (FindFirstFileW) in some circumstances.” I know for a fact there are quite a few orgs using NTFS as mount points for Tomcat to consume.
1. [A security researcher commandeered a country’s expired top-level domain to save it from hackers](https://techcrunch.com/2021/01/15/congo-comandeered/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Good triumphs over evil yet again.
1. [CISA tells agencies to consider ad blockers to fend off ‘malvertising’EFFPrivacy Badger](https://www.cyberscoop.com/ad-blockers-security-nsa-dhs-wyden/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Here’s a suggestion: Ask your company to donate to  and start using  everywhere.
### Tools

1. [Elasticsearch and Kibana are now business riskschanges in licenses for Elasticsearch and KibanaSSPLOSIELKEFK](https://anonymoushash.vmbrasseur.com/2021/01/14/elasticsearch-and-kibana-are-now-business-risks?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    My good friend, VM (Vicky) Brasseur, points out the  are, “As licenses go, it’s pretty problematic from a business perspective.” Also, don’t put your faith in an FAQ when there’s a legally binding document that concerns every intellectual property lawyer, “when you agree to a license you are agreeing to the text of that license document and not to a FAQ. If the text of that license document is ambiguous, then so are your rights and responsibilities under that license.”  is not endorsed by the , period. It’s not an open source license. It’s not permissive. This means that if you stick with the open versions of these projects, you’re not getting security updates. This is the business problem. Adopting the new SSPL’d projects will require orgs to have to release the entire stack under SSPL, which is kinda bullshit. Read Vicky’s post for all the semantics and possibilities that are no longer possible with this change. Needless to say,  and  stacks are potential liabilities now. Thanks for nothing, Elastic.
1. []()

    We need your voice!
1. [Honeycomb.io](https://www.honeycomb.io/?&utm_source=devopsish&utm_medium=newsletter&utm_campaign=ad&utm_content=honeycomb-homepage-devopish)

    In partnership with the team at ClearPath Strategies,  is collecting insights for changes in software development and operation practices across our industry. How do you see the world and what your team is doing?
1. [Take the survey](https://clearpathstrategies.sjc1.qualtrics.com/jfe/form/SV_cMAECZ6jv5wmjrL?&utm_source=devopsish&utm_medium=newsletter&utm_campaign=ad&utm_keyword=&utm_content=software-production-excellence-survey-clearpath-devopsish&utm_adgroup=)

    for a chance to win $500 from Apple, HelloFresh, or Fender. SPONSORED
1. [Sysdig 2021 container security and usage report: Shifting left is not enough](https://sysdig.com/blog/sysdig-2021-container-security-usage-report/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Now I need to update my Security at Cloud Native speed talk. I haven’t had time to read it yet but, this report gave me a ton of ideas and insight last year. I hope for the same this year.
1. [The Final Report on the Slack Outage](https://devopsish.com/pdf/Slack-Incident-Jan-04-2021-RCA-Final.pdf?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    This one jumped out at me, “We have increased the open filehandle limit on our provisioning service workers.” That’s a tough pill to swallow. I wouldn’t imagine this NOT causing a problem of some sort frequently. Mainly capacity issues too.
1. [Apple removes feature that allowed its apps to bypass macOS firewalls and VPNs](https://www.zdnet.com/article/apple-removes-feature-that-allowed-its-apps-to-bypass-macos-firewalls-and-vpns/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    “The ContentFilterExclusionList has been removed in macOS 11.2 beta 2.”
1. [Criminals are Bypassing MFA to Access Organisation’s Cloud Services“pass-the-cookie” attackAttackers Exploit Poor Cyber Hygiene to Compromise Cloud Security Environments](https://www.tripwire.com/state-of-security/featured/hackers-bypassing-mfa-to-access-organisations-cloud-services/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    “In the case it cited, CISA said it believed the malicious hackers may have used a  to waltz around MFA.” The full CISA report is here:
1. [Analyze Kubernetes files for errors with KubeLinter](https://opensource.com/article/21/1/kubelinter?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    “Find and fix errors in your Helm charts and Kubernetes configuration files with KubeLinter.”
1. [Tutorial: Encrypting Kubernetes Secrets with Sealed Secrets](https://www.arthurkoziel.com/encrypting-k8s-secrets-with-sealed-secrets/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Sealed Secrets is a solution to store encrypted Kubernetes secrets in version control. This also applies to the GitOps world as well. Rotate these frequently.
1. [A “no math” (but seven-part) guide to modern quantum mechanics](https://arstechnica.com/science/2021/01/the-curious-observers-guide-to-quantum-mechanics/?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    “Welcome to ‘The curious observer’s guide to quantum mechanics’–featuring particle/wave duality.”
1. [Issues · kubernetes/kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/labels/help%20wanted?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    kube-state-metrics has some help wanted issues that could be a great way for readers to get involved with this critical Kubernetes component. Plus, you’ll get to collaborate with some really awesome people.
1. [davidhampgonsalves/life-dashboard](https://github.com/davidhampgonsalves/life-dashboard?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Heads up Display for every day life
1. [lorin/awesome-limits](https://github.com/lorin/awesome-limits?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Examples of OS / system limits (cough Slack cough)
1. [akinomyoga/ble.sh](https://github.com/akinomyoga/ble.sh?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Bash Line Editor – a full-featured line editor written in pure Bash! Syntax highlighting, auto suggestions, vim modes, etc. are available in Bash interactive sessions!
1. [sandstorm/sku](https://github.com/sandstorm/sku?utm_source=newsletter&utm_medium=devopsish&utm_campaign=201)

    Sandstorm Kubernetes Client - Convenience tools to interact with Kubernetes

### [ << Prev ](sreweekly-200.md) ------------- [ Next >> ](sreweekly-202.md)