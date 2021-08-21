## [SRE Weekly Issue #63](https://sreweekly.com/sre-weekly-issue-63/) - March 12, 2017
### Articles

1. [ferd.ca -> Queues Don’t Fix Overload](http://ferd.ca/queues-don-t-fix-overload.html)

    I love the analogy: you can’t work around a slow drain with a bigger sink.
1. [Tenets of SRE](https://medium.com/@jerub/tenets-of-sre-8af6238ae8a8#.xay7qdifr)

    Stephen Thorne, a Google SRE, annotates the first chapter of the Google SRE book with his personal opinions and interpretations.
1. [Avoid An Oscar Epic Fail: 4 Practices For High Reliability](https://www.forbes.com/sites/davidmarquet/2017/03/10/avoid-an-oscar-epicfail-4-practices-for-high-reliability/)

    The author of this short article starts with the blooper during the Oscars and beautifully segues into a description of techniques organizations can use to halt the propagation of errors.
1. [Sandstorm or Significant: The evolving role of context in Incident Management](http://webinars.devops.com/incident-management-victorops1?portalId=1628905&hsFormKey=bf85b34b2275531f28092b7da6a63600&submissionGuid=375796e6-deee-473d-8bba-eaffb37f64ec#wizard_module_9818716207171313946259819182478846685)

    This webinar looks really interesting, and I’m going to try to see it. It’s about the importance of providing context to incident responders, how much to provide, and how to provide it.This article is published by my sponsor, VictorOps, but their sponsorship did not influence its inclusion in this issue.
### Outages

1. [AT&T 911 service](http://www.al.com/news/index.ssf/2017/03/att_911_outage_fcc_investigati.html)
    AT&T customers were unable to make emergency calls across the US. The Federal Communications Commission (FCC) is investigating.
1. [Bitly (link-shortening service)](http://www.zdnet.com/article/the-bitly-web-link-shortening-service-hiccuped/)
    The cause here is interesting: Comcast’s automated system decided Bitly was a phishing site.
1. [Post-mortem: Outages on 1/19/17 and 1/23/17 – Skyliner](https://blog.skyliner.io/post-mortem-outages-on-1-19-17-and-1-23-17-3f65cc6f693e?source=rss----fae7733bccbd---4)
    I really like their methodical hunt for the offending memory leak.
1. [HSBC (Bank)](http://www.silicon.co.uk/cloud/hsbc-internet-banking-outage-206292)
1. [Google accidentally resets OnHub and Google Wifi routers with server error](https://www.extremetech.com/electronics/245009-google-accidentally-resets-onhub-google-wifi-routers-server-error)
    The routers occasionally ping Google servers for authorization, and on February 23rd the server was sending back an error message. Through some esoteric fallback mechanism in the routers, this caused them to reset to factory settings. So, a problem on Google’s servers can reset your router. Oops.
1. [Incident 1059 | Heroku Status](https://status.heroku.com/incidents/1059)
    Heroku posted a followup regarding their outage on February 28th stemming from the Amazon S3 outage.
Full disclosure: Heroku is my employer and I was involved in writing this followup.

### [ << Prev ](sreweekly-62.md) ------------- [ Next >> ](sreweekly-64.md)