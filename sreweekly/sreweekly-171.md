## [SRE Weekly Issue #171](https://sreweekly.com/sre-weekly-issue-171/) - May 5, 2019
### Articles

1. [Reliability That Works](https://www.linkedin.com/pulse/reliability-works-aaron-blohowiak/)

    Make failure a non-event, rather than trying to prevent it. You won’t succeed in fully preventing failures, and you’ll instead get out of practice of recovering.Aaron Blohowiak
1. [How I Almost Destroyed a £50 million War Plane and The Normalisation of Deviance.](https://fastjetperformance.com/podcasts/how-i-almost-destroyed-a-50-million-war-plane-when-display-flying-goes-wrong-and-the-normalisation-of-deviance/)

    They had me at “normalization of deviance”. I’ll read pretty much anything with that in the title. Tim Davies — Fast Jet Performance
1. [Monzo’s real-time incident response and reporting tool](https://github.com/monzo/response)

    Monzo’s system is directly integrated with Slack, helping you manage your incident and track what happens. Check out their video presentation for more details.Monzo
1. [Nolan Caudill on Twitter: “I think postmortem docs are an underused avenue for recruiting engineers.”](https://twitter.com/nolancaudill/status/1124532823900082177)

    Me too! Great thread.Nolan Caudill and others
1. [Incident Review: Caches are Good, Except When They Are Bad](https://www.honeycomb.io/blog/incident-review-caches-are-good-except-when-they-are-bad/)

    I love Honeycomb incident reviews, I really do.Douglas Soo
1. [Friday Deploy Freezes Are Exactly Like Murdering Puppies](https://charity.wtf/2019/05/01/friday-deploy-freezes-are-exactly-like-murdering-puppies/)

    Born from a Twitter argument thread, this article goes into depth about why Friday change freezes can cause much more trouble than good.Charity Majors
### Outages

1. [Amazon EC2](https://status.aws.amazon.com/)
    Network-related issues in Japan and Hong Kong (on separate days). It’s starting to become downright impossible to find historical incidents on their mile-long status page.
1. [Google Hangouts Meet](https://www.google.com/appsstatus#hl=en&v=issue&sid=44&iid=70c8bd72b9dfa3e16aadc1800207e1c0)
1. [Google Cloud Console](https://status.cloud.google.com/incident/developers-console/19003#19003003)
1. [Slack](https://status.slack.com/2019-05/6d7541f932dd628c)
1. [Azure, Microsoft 365, and Dynamics 365](https://azure.microsoft.com/en-us/status/history/)
    A DNS change went awry, resulting in one of their DNS zone’s four nameservers having an empty copy of the zone and serving NXDOMAIN. This is a really interesting incident report to read. Had the nameserver simply not had the zone at all, it would have returned a non-authoritative answer, and clients would have fallen back to one of the other three nameservers.
1. [Wells Fargo (bank)](https://www.newsweek.com/wells-fargo-app-down-not-working-website-mobile-deposit-1414148)
1. [Discord](https://discord.statuspage.io/incidents/kz4ky0cydym1)
1. [Google AdSense](https://news.alphastreet.com/google-adsense-suffers-massive-global-outage/)
1. [Facebook, Instagram, and WhatsApp](https://www.retailnews.asia/subscribers-to-facebook-instagram-and-whatsapp-faced-over-2-hours-downtime/)
1. [Coles (Supermarket chain)](https://www.dailymail.co.uk/news/article-6974701/Angry-customers-slam-Coles-supermarkets-online-shopping-website-goes-two-days.html)
1. [Hallifax and Lloyds (banks)](https://metro.co.uk/2019/04/29/natwest-lloyds-halifax-internet-banking-apps-9338739/)

### [ << Prev ](sreweekly-170.md) ------------- [ Next >> ](sreweekly-172.md)