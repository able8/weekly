## [DevOps'ish 027](https://devopsish.com/027) - Sun Jun 11, 2017

Quite the week in the greater world of technology. <a href="https://developer.apple.com/wwdc/">Apple had an event</a> and <a href="http://www.businessinsider.com/apple-planet-of-the-apps-ad-developer-rarely-saw-his-kids-2017-6">folks had a meltdown about their new show</a>. GitHub had a rough week, Uber sunk to new lows, and Intel has fired a massive warning shot at the future of computing. A lot of <a href="https://devopsish.com/">DevOps</a> happened too. Let’s break it all down.

<a href="https://2017.fullstackfest.com"><strong>Full Stack Fest 2017: Problems of today, wonders from the future.</strong></a><br/>Barcelona, 4–8 Sept. 2017<br/>Are you a curious mind? <a href="https://2017.fullstackfest.com">Full Stack Fest</a> is a week-long conference based in the amazing city of Barcelona that peeks into the web of tomorrow. Serverless, Blockchain, WebVR, Distributed Web, Progressive Web Apps… Come and see! <em>SPONSORED</em>

Intel took us down memory lane this week with a <a href="https://newsroom.intel.com/editorials/x86-approaching-40-still-going-strong/">blog post discussing how X86 is almost forty years old</a>. Remember the 8086? MMX? SSE 1 through 4? VMX? Good times! Oh and remember that Intel has 1,600 patents and they are not afraid to sue the bejesus out of you.

Intel is firing a shot across the bow for anyone trying to do X86 emulation without ponying up licensing fees to Intel. “…we fully expect other companies to continue to respect Intel’s intellectual property rights.” <a href="https://www.axios.com/intel-steer-clear-of-our-patents-2437931228.html">Axios points out</a>, “Microsoft and Qualcomm have announced plans for a version of Windows 10 on Qualcomm’s Snapdragon 835 that uses emulation to run older applications designed for x86-based Windows machines.”

Emulation (software mimicking hardware) is important for the advancement of microprocessors. Emulation can be used for older workloads instead having to buy expensive, overpowered hardware to maintain something that could run on a much cheaper component. Emulation comes in so that you don’t have to completely rewrite your legacy stack because you can’t get new hardware for it to run on. This will be a story that won’t end anytime soon.

<a href="https://itrevolution.com/imaginary-apology-letter-airline-ceo/">Gene Kim weighed in on last week’s British Airways debacle</a>. Gene wrote a fictitious letter that is what <a href="https://en.wikipedia.org/wiki/Willie_Walsh_(businessman)">Willie Walsh</a> should have written. “The accident last week was not due to a power failure, or an IT failure — this was a business failure. After all, we were unable to perform some of our most critical business operations for nearly three days.” Bravo! 👏👏👏

Also from Gene Kim land, <a href="https://puppet.com/resources/whitepaper/state-of-devops-report">2017 State of DevOps Report</a> is out. There are a lot of takeaways from this report but one of the key findings stands out to me: Dimensions of transformational leadership. High performing teams have leadership that shares five characteristics; vision, intellectual stimulation, recognition, inspirational communication, and “supportive leadership”.

Red Hat is offering free <a href="https://www.edx.org/course/fundamentals-red-hat-enterprise-linux-red-hat-rh066x#!">Fundamentals of Red Hat Enterprise Linux</a> training through edX. This is going to be helpful when people ask me how to get started with Linux.

Brendan Gregg has updated his awesome <a href="http://www.brendangregg.com/linuxperf.html">Linux Performance Tools</a> diagrams. This is something I print out and have very handy at all times.

Julia Evans shares, “<a href="https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/">A few things I’ve learned about Kubernetes</a>” complete with awesome sketches.

<a href="https://arstechnica.com/security/2017/06/russian-hackers-turn-to-britney-spears-for-help-concealing-espionage-malware/">Why set up a complex process to communicate with command and control systems for your malware when you could just use Britney Spears’ Instagram comments?</a> This is a fairly brilliant tactic; good luck taking down Britney Spears social media profiles and risk pissing tens of millions of fans.

As previously mentioned, GitHub has had a hard time of late. <a href="https://status.github.com/messages/2017-06-10">GitHub experienced two “major service outage” events this week</a>. I feel for the team there; it’s not easy running the services that distribute a ton of the world’s code.

<a href="http://www.theregister.co.uk/2017/06/06/windows_defender_competition_complaint/">Kaspersky is suing Microsoft claiming that Microsoft’s heavy marketing of Windows Defender violates antitrust law</a>. Kaspersky called Windows Defender, “inferior” in its claim. Damn!

<a href="https://www.reddit.com/r/cscareerquestions/comments/6ez8ag/accidentally_destroyed_production_database_on/">This reddit thread exploded this week</a>. A new, junior member of a development team was given shoddy documentation with production connection string, credentials, and table names to setup their dev environment. No one told them that they shouldn’t actually use that information and they ended up blowing away the production database. It’s a great read (yes, even the comments, to an extent).

Uber sunk to new lows this week. <a href="http://www.businessinsider.com/uber-fired-more-than-20-employees-as-part-of-its-sex-harassment-probe-2017-6">Uber fired more than 20 employees after receiving 215 claims in probe of sex harassment and other incidents</a>. <a href="https://www.recode.net/2017/6/7/15754316/uber-executive-india-assault-rape-medical-records">A top Uber executive, who obtained the medical records of a customer who was a rape victim, has been fired</a>. <a href="https://www.recode.net/2017/6/8/15765514/2013-miami-letter-uber-ceo-kalanick-employees-sex-rules-company-celebration">Uber CEO Travis Kalanick advised employees on sex rules for a company celebration in 2013 ‘Miami letter’</a>. <a href="https://www.cnbc.com/amp/2017/06/07/arianna-huffington-says-uber-ceo-travis-kalanick-has-started-meditating.html">Kalanick has started meditating in an office lactation room, says Arianna Huffington</a>. If you still are working at Uber, I’d like to know why? Seriously.

<a href="https://devopsish.com/sponsor/" title="Sponsor DevOps&#39;ish"><strong>Sponsor DevOps&#39;ish</strong></a> and put your brand in front of thousands of highly skilled operators, maintainers, developers, and leaders from Amazon, Apple, Google, IBM, Intel, Microsoft, Red Hat, many of the Fortune 100, and beyond. Download the <strong><a href="https://devopsi.sh/prospectus">DevOps&#39;ish Sponsorship Prospectus</a></strong> now!

Join <a href="https://www.reddit.com/r/devopsish/">/<span class="fa fa-reddit-alien fa-sm" aria-hidden="true"></span>/devopsish</a> for a stream of news and content throughout the week.

### Department of Patent Pedantry

1. [blog post discussing how X86 is almost forty years old](https://newsroom.intel.com/editorials/x86-approaching-40-still-going-strong/)

    Intel took us down memory lane this week with a . Remember the 8086? MMX? SSE 1 through 4? VMX? Good times! Oh and remember that Intel has 1,600 patents and they are not afraid to sue the bejesus out of you.
1. [Axios points out](https://www.axios.com/intel-steer-clear-of-our-patents-2437931228.html)

    Intel is firing a shot across the bow for anyone trying to do X86 emulation without ponying up licensing fees to Intel. “…we fully expect other companies to continue to respect Intel’s intellectual property rights.” , “Microsoft and Qualcomm have announced plans for a version of Windows 10 on Qualcomm’s Snapdragon 835 that uses emulation to run older applications designed for x86-based Windows machines.”
1. []()

    Emulation (software mimicking hardware) is important for the advancement of microprocessors. Emulation can be used for older workloads instead having to buy expensive, overpowered hardware to maintain something that could run on a much cheaper component. Emulation comes in so that you don’t have to completely rewrite your legacy stack because you can’t get new hardware for it to run on. This will be a story that won’t end anytime soon.
### Department of Choice Concepts

1. [Gene Kim weighed in on last week’s British Airways debacleWillie Walsh](https://itrevolution.com/imaginary-apology-letter-airline-ceo/)

    . Gene wrote a fictitious letter that is what  should have written. “The accident last week was not due to a power failure, or an IT failure — this was a business failure. After all, we were unable to perform some of our most critical business operations for nearly three days.” Bravo! 👏👏👏
1. [2017 State of DevOps Report](https://puppet.com/resources/whitepaper/state-of-devops-report)

    Also from Gene Kim land,  is out. There are a lot of takeaways from this report but one of the key findings stands out to me: Dimensions of transformational leadership. High performing teams have leadership that shares five characteristics; vision, intellectual stimulation, recognition, inspirational communication, and “supportive leadership”.
1. [Fundamentals of Red Hat Enterprise Linux](https://www.edx.org/course/fundamentals-red-hat-enterprise-linux-red-hat-rh066x#!)

    Red Hat is offering free  training through edX. This is going to be helpful when people ask me how to get started with Linux.
1. [Linux Performance Tools](http://www.brendangregg.com/linuxperf.html)

    Brendan Gregg has updated his awesome  diagrams. This is something I print out and have very handy at all times.
1. [A few things I’ve learned about Kubernetes](https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/)

    Julia Evans shares, “” complete with awesome sketches.
### Department of Data Defense

1. [Why set up a complex process to communicate with command and control systems for your malware when you could just use Britney Spears’ Instagram comments?](https://arstechnica.com/security/2017/06/russian-hackers-turn-to-britney-spears-for-help-concealing-espionage-malware/)

    This is a fairly brilliant tactic; good luck taking down Britney Spears social media profiles and risk pissing tens of millions of fans.
1. [GitHub experienced two “major service outage” events this week](https://status.github.com/messages/2017-06-10)

    As previously mentioned, GitHub has had a hard time of late. . I feel for the team there; it’s not easy running the services that distribute a ton of the world’s code.
1. [Kaspersky is suing Microsoft claiming that Microsoft’s heavy marketing of Windows Defender violates antitrust law](http://www.theregister.co.uk/2017/06/06/windows_defender_competition_complaint/)

    . Kaspersky called Windows Defender, “inferior” in its claim. Damn!
### Department of Sane Workplaces

1. [This reddit thread exploded this week](https://www.reddit.com/r/cscareerquestions/comments/6ez8ag/accidentally_destroyed_production_database_on/)

    . A new, junior member of a development team was given shoddy documentation with production connection string, credentials, and table names to setup their dev environment. No one told them that they shouldn’t actually use that information and they ended up blowing away the production database. It’s a great read (yes, even the comments, to an extent).
1. [Uber fired more than 20 employees after receiving 215 claims in probe of sex harassment and other incidentsA top Uber executive, who obtained the medical records of a customer who was a rape victim, has been firedUber CEO Travis Kalanick advised employees on sex rules for a company celebration in 2013 ‘Miami letter’Kalanick has started meditating in an office lactation room, says Arianna Huffington](http://www.businessinsider.com/uber-fired-more-than-20-employees-as-part-of-its-sex-harassment-probe-2017-6)

    Uber sunk to new lows this week. . . . . If you still are working at Uber, I’d like to know why? Seriously.

### [ << Prev ](sreweekly-26.md) ------------- [ Next >> ](sreweekly-28.md)