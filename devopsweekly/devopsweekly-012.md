## [DevOps'ish 012](https://devopsish.com/012) - Sun Feb 26, 2017

This week has been highlighted by multiple security events

What an incredibly busy week news wise. No matter where you are in your <a href="https://devopsish.com/"><strong>DevOps</strong></a> journey it’s very likely one of the major events this week affected you. I am in the process of fighting off a sinus infection too so this week was a lot to handle.

Unless you live under a rock I am almost certain you have heard about the <a href="https://www.susanjfowler.com/blog/2017/2/19/reflecting-on-one-very-strange-year-at-uber">insanity Susan J. Fowler went through during her time at Uber</a>. This shit is #NotOkay, folks, period. If you proposition your coworkers, you are in the wrong. If you harass your coworkers, you are in the wrong. If you cover up sexual misconduct, workplace violence, or any other human resource issue in your organization, <strong>YOU ARE IN THE WRONG</strong>. <a href="https://chrisshort.net/what-the-military-taught-me-about-devops/">I have said this a lot in the past year</a>, I did not spend 11 years in the military so ignorant jerks can harass people. You can read the <a href="https://www.nytimes.com/2017/02/22/technology/uber-workplace-culture.html">incredibly well written piece by Mike Isaac</a> for an impartial point of view.

<a href="https://github.com/danapsimer">Dana P’Simer</a> released a pretty nifty tool to <a href="https://github.com/danapsimer/aws-api-to-lambda-shim">utilize Go with AWS Lambda</a>.

<a href="https://gitlab.com/gitlab-com/runbooks">GitLab released all of their runbooks for public consumption</a>. If you need to create documentation this is a good place to start.

<a href="https://bugs.chromium.org/p/project-zero/issues/detail?id=1139">Google’s Project Zero discovered</a> a <a href="https://blog.cloudflare.com/incident-report-on-memory-leak-caused-by-cloudflare-parser-bug/">memory leak in a CloudFlare parser</a>. This was a pretty gnarly leak. This summary from <a href="https://ma.ttias.be/cloudbleed-cloudflare-reverse-proxies-dumping-uninitialized-memory/">Mattias Geniar</a> is spot on, “any site behind Cloudflare <em>might</em> have leaked info from other sites hosted behind Cloudflare. Even if your HTML was perfectly balanced and strict, your site might’ve become the victim of another site with imbalanced HTML tags that leaked <em>your</em> data.”

<a href="https://security.googleblog.com/2017/02/announcing-first-sha1-collision.html">Google found a SHA1 collision</a> so the already limping hash function is now, for all intents and purposes, dead. Unless you are <a href="https://public-inbox.org/git/CA+55aFxJGDpJXqpcoPnwvzcn_fB-zaggj=w7P2At-TOt4buOqw@mail.gmail.com/">Linus Torvalds</a> who said, “I doubt the sky is falling for git as a source control management tool. Do we want to migrate to another hash? Yes. Is it ‘game over’ for SHA1 like people want to say? Probably not.”

<a href="http://techblog.netflix.com/2017/02/introducing-netflix-stethoscope.html">Netflix open sourced Stethoscope</a> this week. “Stethoscope is a web application that collects information for a given user’s devices and gives them clear and specific recommendations for securing their systems.”

<a href="https://techcrunch.com/2017/02/22/eff-half-the-web-is-now-encrypted/">EFF is reporting that over half of the web’s traffic is now encrypted</a>.

<a href="https://www.ansible.com/blog/happy-birthday-ansible">Ansible turned five years old this week!</a>

I end up sending a lot of work through AWS SQS. This little one liner is handy for blasting more than one message into SQS:

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Sane Workplaces

1. [insanity Susan J. Fowler went through during her time at UberI have said this a lot in the past yearincredibly well written piece by Mike Isaac](https://www.susanjfowler.com/blog/2017/2/19/reflecting-on-one-very-strange-year-at-uber)

    Unless you live under a rock I am almost certain you have heard about the . This shit is #NotOkay, folks, period. If you proposition your coworkers, you are in the wrong. If you harass your coworkers, you are in the wrong. If you cover up sexual misconduct, workplace violence, or any other human resource issue in your organization, YOU ARE IN THE WRONG. , I did not spend 11 years in the military so ignorant jerks can harass people. You can read the  for an impartial point of view.
### Department of Choice Concepts

1. [Dana P’Simerutilize Go with AWS Lambda](https://github.com/danapsimer)

    released a pretty nifty tool to .
### Department of Data Defense

1. [GitLab released all of their runbooks for public consumption](https://gitlab.com/gitlab-com/runbooks)

    . If you need to create documentation this is a good place to start.
1. [Google’s Project Zero discoveredmemory leak in a CloudFlare parserMattias Geniar](https://bugs.chromium.org/p/project-zero/issues/detail?id=1139)

    a . This was a pretty gnarly leak. This summary from  is spot on, “any site behind Cloudflare might have leaked info from other sites hosted behind Cloudflare. Even if your HTML was perfectly balanced and strict, your site might’ve become the victim of another site with imbalanced HTML tags that leaked your data.”
1. [Google found a SHA1 collisionLinus Torvalds](https://security.googleblog.com/2017/02/announcing-first-sha1-collision.html)

    so the already limping hash function is now, for all intents and purposes, dead. Unless you are  who said, “I doubt the sky is falling for git as a source control management tool. Do we want to migrate to another hash? Yes. Is it ‘game over’ for SHA1 like people want to say? Probably not.”
1. [Netflix open sourced Stethoscope](http://techblog.netflix.com/2017/02/introducing-netflix-stethoscope.html)

    this week. “Stethoscope is a web application that collects information for a given user’s devices and gives them clear and specific recommendations for securing their systems.”
1. [EFF is reporting that over half of the web’s traffic is now encrypted](https://techcrunch.com/2017/02/22/eff-half-the-web-is-now-encrypted/)

    .
### Department of Next Year’s Old Tech

1. [Ansible turned five years old this week!](https://www.ansible.com/blog/happy-birthday-ansible)

    
### DevOps’ish One-Liner of the Week

1. []()

    I end up sending a lot of work through AWS SQS. This little one liner is handy for blasting more than one message into SQS:
1. [Sponsor DevOps'ishDevOps'ish Sponsorship Prospectus](https://devopsish.com/sponsor/)

    and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the DevOps'ish Sponsorship Prospectus now!
1. [//devopsish](https://www.reddit.com/r/devopsish/)

    Join  for a stream of news and content throughout the week.

### [ << Prev ](devopsweekly-011.md) ------------- [ Next >> ](devopsweekly-013.md)