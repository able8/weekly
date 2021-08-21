## [SRE Weekly Issue #6](https://sreweekly.com/sre-weekly-issue-6/) - January 17, 2016
### Articles

1. [Designed to Fail – Brave New Geek](http://bravenewgeek.com/designed-to-fail/)

    A discussion of failing fast, degrading gracefully, and applying back-pressure to avoid cascading failure in a service-oriented architecture.

1. [Kernel Patching 101: How to Make Repairs Without System Downtime/](https://appdevelopermagazine.com/3486/2016/1/5/Kernel-Patching-101:-How-to-Make-Repairs-Without-System-Downtime/)

    A SUSE developer introduces kGraft, SUSE’s system for live kernel patching.  Anyone who survived the AWS reboot-a-thon is probably a big fan of live kernel patching solutions.
1. [Not Everything Critical is Urgent. Learn the Difference.](https://www.pagerduty.com/blog/critical-versus-urgent-alert-fatigue/)

    One thing that is critical is avoiding burnout in on-call.  This article is a description of the “urgency” feature in Pagerduty, but they make a generally applicable point: don’t wake someone for something just because it’s critical; only wake them if it needs immediate action.
1. [Fallacies of Distributed Computing Explained](http://www.rgoarchitects.com/Files/fallacies.pdf)

    This is a review/update of the 1994 article.  The fallacies still hold true, and anyone designing a large-scale service should heed them.  The fallacies:
  As I get into SRE Weekly, I repeatedly run across articles that I probably should have read long since in my career.  Hopefully they’re new to some of you, too.
1. [Delivering safer cars faster through automation and continuous delivery](http://cdn.electric-cloud.com/wp-content/uploads/Automotive-World-delivering-safer-cars-faster-automation-continuous-delivery.pdf)

    Every position I’ve held has involved supporting reliability in a 24/7 service, but let’s be realistic: it’s unlikely someone would have died as a result of an outage.  In cars, reliability takes a whole new meaning.  I first got interested in MISRA and the other standards surrounding the code running in cars when I read some technical write-ups of the investigation surrounding the “unintended acceleration” incidents a few years back.  This article discusses how devops practices are being applied in the development of vehicle code.
1. [Security experts confirm Ukraine power grid blackout a ‘coordinated intentional attack’](http://www.v3.co.uk/v3-uk/news/2440469/ukraine-investigating-suspected-russian-cyber-attack-on-power-grid)

    Evidence has come out that the recent major power outage in Ukraine was a network-based attack (I can’t make myself say “cyber-” anything).
1. [PS4 porn viewers rocket during PSN outage](http://www.psu.com/news/29103/PS4-porn-viewers-rocket-during-PSN-outage-)

    I should have seen this coming.
1. [Verizon grounds JetBlue — how could that happen? Another plan B gone bad](https://bobsullivan.net/cybercrime/verizon-grounds-jetblue-how-could-that-happen-another-plan-b-gone-bad/)

    One blogger’s take on the JetBlue outage.

1. [SRECon16 Call for Participation](https://www.usenix.org/conference/srecon16/call-for-participation)

    The SRECon call for participation is now open!
1. [LostPass](https://www.seancassidy.me/lostpass.html)

    Sean Cassidy has discovered an easy and indistinguishable phishing method for LastPass in Chrome, with a slightly less simple and effective method for Firefox.  This one’s important for availability because many organizations rely heavily on LastPass.  Compromising the right Employee’s vault could spell big trouble and possibly downtime.
### Outages

1. [GTA Online](http://www.gamespot.com/articles/gta-online-ps4-servers-suffer-outage/1100-6433668/)
1. [EE (phone network)](http://leadercall.com/2016/01/severe-outage-with-ee-phone-network-leaves-people-unable-to/)
1. [Amplitude](http://status.amplitude.com/incidents/njfbgbbpcdjg)
    A truly heinous multi-day outage for Amplitude.  The root cause appears to be inadvertent deletion of data in DynamoDB.  Thanks to the folks at Amplitude for the extremely detailed status and analysis.  Get some sleep, folks.
1. [PlayStation Network](http://www.express.co.uk/entertainment/gaming/633755/PlayStation-Network-down-PSN-status-PS4-Sony)
1. [Xbox Live Down](http://www.ibtimes.com/xbox-live-down-online-multiplayer-party-chat-among-affected-services-2264074)
1. [JetBlue](http://www.bostonherald.com/news/local_coverage/2016/01/jetblue_says_flight_booking_system_back_up_after_power_outage)
    This one was all over the news.  JetBlue points the finger at a Verizon datacenter outage.
1. [TalkTalk](http://www.gizmodo.co.uk/2016/01/major-talktalk-outage-leaves-customers-without-phone-or-internet-access/)
1. [Yahoo Mail](http://www.ibtimes.co.uk/yahoo-mail-down-email-service-suffered-worldwide-outage-did-anyone-really-care-1538173)

### [ << Prev ](sreweekly-5.md) ------------- [ Next >> ](sreweekly-7.md)