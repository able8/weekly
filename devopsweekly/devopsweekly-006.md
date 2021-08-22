## [DevOps'ish 006](https://devopsish.com/006) - Sun Jan 15, 2017

Welcome to this week’s edition of <a href="https://devopsish.com/"><strong>DevOps’ish</strong></a> where we cover Dev, Ops, and all the ish in between.

I hope you had a productive week and are looking forward to another productive week ahead!

So much <a href="https://devopsish.com/"><strong>DevOps</strong></a> , so little time.

<a href="http://www.forbes.com/sites/alexkonrad/2017/01/09/atlassian-acquires-popular-team-productivity-app-trello-for-425-million/">Trello was acquired by Atlassian for a whopping $425 million</a>. This is an acquisition that makes a lot of sense for Atlassian. Trello could be used as a sort of gateway drug into the much more complicated Atlassian software products.

<a href="https://9to5mac.com/2017/01/10/swift-chris-lattner-leaving-apple/">Chris Lattner, the creator of Swift, is leaving Apple and joining the Autopilot team at Tesla</a>. According to a <a href="http://www.businessinsider.com/how-apples-culture-of-secrecy-wears-down-its-top-developers-2017-1">Business Insider</a> source, Chris Lattner left because he felt, “He always felt constrained at Apple in terms of what he could discuss publicly.” But, <a href="https://twitter.com/clattner_llvm/status/819974025371787264">Chris vehemently denied this report</a> in a tweet saying, “Folk just want to make 🍎 look bad. 😠”.

<a href="https://twitter.com/ashleymcnamara">Ashley McNamara</a> maintains a “<a href="http://ashleymcnamara.github.io/learn_to_code/">Curated list of resources for budding developers</a>” that is a consolidated treasure trove of language resources.

<a href="https://www.blockloop.io/mastering-bash-and-terminal">blockloop has a good post on bash and the terminal</a>. It’s pretty basic stuff and a little macOS centric. But, it does discuss my two favorite CLI utilities in macOS: pbcopy and pbpaste.

I don’t know how I missed this last week but Flickr (yes, Flickr) has a great post about how they <a href="https://code.flickr.net/2017/01/05/a-year-without-a-byte/">tried to go through 2016 without buying new storage systems</a>.

<a href="https://theagileadmin.com/2017/01/10/but-how-can-it-do-devops/">LinkedIn has a vending machine for IT equipment</a> much like factories have vending of parts and PPE. That’s one way to automate IT functions.

<img src="https://d33wubrfki0l68.cloudfront.net/7d98d39e5728bcbf06616c8273dcb635ee603de5/40493/006/linkedin-vending-machine.jpg" alt="Photo Courtesy [the agile admin](https://theagileadmin.com/2017/01/10/but-how-can-it-do-devops/)"/><br/><em>Photo Courtesy <a href="https://theagileadmin.com/2017/01/10/but-how-can-it-do-devops/">the agile admin</a></em>

Tom Croucher of Uber brings us a “I’ve seen the light!” moment in <a href="https://www.oreilly.com/ideas/dont-gamble-when-it-comes-to-reliability"><em>Don’t gamble when it comes to reliability</em></a>

<a href="http://www.theregister.co.uk/2017/01/12/emc_layoffs/">Dell EMC is planning to layoff a big chunk of its workforce</a> post-merger. The cuts could affect as much as a quarter of Dell EMC employees (that is up to 35,000 people). Insiders say things have been a little awkward inside the company since the beginning of their fiscal Q4.

<a href="https://github.com/docker/docker/releases/tag/v1.12.6">Docker 1.12.6 was released</a> this week to address <a href="http://seclists.org/fulldisclosure/2017/Jan/21">CVE-2016–9962</a>. The <a href="https://bugzilla.suse.com/show_bug.cgi?id=1012568">vulnerability was pretty gnarly</a>: “runC passes a file descriptor from the host’s filesystem to the ‘runc init’ bootstrap process when joining a container. This allows a malicious process inside a container to gain access to the host
filesystem with its current privilege set.” <a href="http://rhelblog.redhat.com/2017/01/13/docker-0-day-stopped-cold-by-selinux/">Red Hat demonstrated that SELinux could have saved you</a> from this bug ever happening in your environment.

<a href="http://www.theregister.co.uk/2017/01/11/ansible_patches_own_the_farm_vulnerability/">Ansible had a pretty gnarly bug announced this week</a>. CVE-2016–9587 was discovered by <a href="https://www.computest.nl/advisories/CT-2017-0109_Ansible.txt">Computest</a>. According to James Cammarata, Ansible Lead/Sr. Principal Software Engineer, “<a href="https://groups.google.com/forum/#!topic/ansible-devel/SyrgcUySAIQ">a compromised remote system being managed via Ansible can lead to commands being run on the Ansible controller (as the user running the ansible or ansible-playbook command).</a>” Multiple release candidates came out during the week. <a href="https://groups.google.com/forum/?utm_medium=email&amp;utm_source=footer#!msg/ansible-project/ydfEh11hlXA/eaOPE3p_AQAJ">A stable 2.x release should be coming January 16, 2017</a>. If you are still running Ansible 1.x code, you should upgrade to the 2.1.4 or 2.2.1 release as soon as possible.

There was a great crowd at the <a href="https://www.meetup.com/Triangle-Kubernetes-Meetup/">Triangle Kubernetes and OpenShift Meetup</a> this past week. I got some helpful tips from Thomas Wiest about running Kubernetes in AWS. I have made <a href="https://www.evernote.com/l/AAV9Qz3Q7IhLCqRM05VZbis3oJm6duhVdi0">my notes from the Meetup available</a>.

<img src="https://d33wubrfki0l68.cloudfront.net/ff6fddd55983afd0782317e1bd9e7ccad1a4a730/d9485/006/triangle-openshift-meetup.jpg" alt="Photo courtesy of [Red Hat OpenShift](https://twitter.com/openshift/status/819690494393520131)"/><br/><em>Photo courtesy of <a href="https://twitter.com/openshift/status/819690494393520131">Red Hat OpenShift</a></em>

​<a href="http://www.zdnet.com/article/red-hat-enterprise-linux-6-9-beta-out-now/">Red Hat Enterprise Linux 6.9 beta has been released</a>. This release brings TLS 1.2 to the 6.x branch of RHEL which is a very welcomed (albeit late) improvement.

<a href="https://groups.google.com/forum/m/#!topic/golang-nuts/tr2ZKSQ42zE">Go 1.8 RC1 was released</a> this past week. The feature I’m most excited for is a default $GOPATH. I swear every time I install Go I never remember where I put my $GOPATH so the default will lower the bar to entry for newer folk (and help the senile).

The fantastic Hacker News alternative, <a href="https://lobste.rs/">**Lobsters</a>** is having a membership drive of sorts. If you would like an invite just let me know.

If you are a Linux DevOps person with AWS talents please <a href="http://solarwinds.jobs/durham-nc/senior-aws-systems-engineer/866A1F89D368408CB7F6F774409D961A/job/">consider joining my Global DevOps team</a>.

This week’s one-liner comes courtesy of <a href="https://twitter.com/jpetazzo/status/819288097376583680">Jérôme Petazzoni</a>:

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Assemblage Obtainment

1. [Trello was acquired by Atlassian for a whopping $425 million](http://www.forbes.com/sites/alexkonrad/2017/01/09/atlassian-acquires-popular-team-productivity-app-trello-for-425-million/)

    . This is an acquisition that makes a lot of sense for Atlassian. Trello could be used as a sort of gateway drug into the much more complicated Atlassian software products.
1. [Chris Lattner, the creator of Swift, is leaving Apple and joining the Autopilot team at TeslaBusiness InsiderChris vehemently denied this report](https://9to5mac.com/2017/01/10/swift-chris-lattner-leaving-apple/)

    . According to a  source, Chris Lattner left because he felt, “He always felt constrained at Apple in terms of what he could discuss publicly.” But,  in a tweet saying, “Folk just want to make 🍎 look bad. 😠”.
1. [Ashley McNamaraCurated list of resources for budding developers](https://twitter.com/ashleymcnamara)

    maintains a “” that is a consolidated treasure trove of language resources.
1. [blockloop has a good post on bash and the terminal](https://www.blockloop.io/mastering-bash-and-terminal)

    . It’s pretty basic stuff and a little macOS centric. But, it does discuss my two favorite CLI utilities in macOS: pbcopy and pbpaste.
### Department of Choice Concepts

1. [tried to go through 2016 without buying new storage systems](https://code.flickr.net/2017/01/05/a-year-without-a-byte/)

    I don’t know how I missed this last week but Flickr (yes, Flickr) has a great post about how they .
1. [LinkedIn has a vending machine for IT equipment](https://theagileadmin.com/2017/01/10/but-how-can-it-do-devops/)

    much like factories have vending of parts and PPE. That’s one way to automate IT functions.
1. [the agile admin](https://theagileadmin.com/2017/01/10/but-how-can-it-do-devops/)

    Photo Courtesy the agile admin
1. [Don’t gamble when it comes to reliability](https://www.oreilly.com/ideas/dont-gamble-when-it-comes-to-reliability)

    Tom Croucher of Uber brings us a “I’ve seen the light!” moment in
### Department of Dafuq

1. [Dell EMC is planning to layoff a big chunk of its workforce](http://www.theregister.co.uk/2017/01/12/emc_layoffs/)

    post-merger. The cuts could affect as much as a quarter of Dell EMC employees (that is up to 35,000 people). Insiders say things have been a little awkward inside the company since the beginning of their fiscal Q4.
### Department of Data Defense

1. [Docker 1.12.6 was releasedCVE-2016–9962vulnerability was pretty gnarlyRed Hat demonstrated that SELinux could have saved you](https://github.com/docker/docker/releases/tag/v1.12.6)

    this week to address . The : “runC passes a file descriptor from the host’s filesystem to the ‘runc init’ bootstrap process when joining a container. This allows a malicious process inside a container to gain access to the host
filesystem with its current privilege set.”  from this bug ever happening in your environment.
1. [Ansible had a pretty gnarly bug announced this weekComputesta compromised remote system being managed via Ansible can lead to commands being run on the Ansible controller (as the user running the ansible or ansible-playbook command).A stable 2.x release should be coming January 16, 2017](http://www.theregister.co.uk/2017/01/11/ansible_patches_own_the_farm_vulnerability/)

    . CVE-2016–9587 was discovered by . According to James Cammarata, Ansible Lead/Sr. Principal Software Engineer, “” Multiple release candidates came out during the week. . If you are still running Ansible 1.x code, you should upgrade to the 2.1.4 or 2.2.1 release as soon as possible.
### Department of Discussion

1. [Triangle Kubernetes and OpenShift Meetupmy notes from the Meetup available](https://www.meetup.com/Triangle-Kubernetes-Meetup/)

    There was a great crowd at the  this past week. I got some helpful tips from Thomas Wiest about running Kubernetes in AWS. I have made .
1. [Red Hat OpenShift](https://twitter.com/openshift/status/819690494393520131)

    Photo courtesy of Red Hat OpenShift
### Department of Refreshment and Refurbishment

1. [Red Hat Enterprise Linux 6.9 beta has been released](http://www.zdnet.com/article/red-hat-enterprise-linux-6-9-beta-out-now/)

    ​. This release brings TLS 1.2 to the 6.x branch of RHEL which is a very welcomed (albeit late) improvement.
1. [Go 1.8 RC1 was released](https://groups.google.com/forum/m/#!topic/golang-nuts/tr2ZKSQ42zE)

    this past week. The feature I’m most excited for is a default $GOPATH. I swear every time I install Go I never remember where I put my $GOPATH so the default will lower the bar to entry for newer folk (and help the senile).
### Not DevOps But Still Cool

1. [**Lobsters](https://lobste.rs/)

    The fantastic Hacker News alternative, ** is having a membership drive of sorts. If you would like an invite just let me know.
1. [consider joining my Global DevOps team](http://solarwinds.jobs/durham-nc/senior-aws-systems-engineer/866A1F89D368408CB7F6F774409D961A/job/)

    If you are a Linux DevOps person with AWS talents please .

### [ << Prev ](devopsweekly-005.md) ------------- [ Next >> ](devopsweekly-007.md)