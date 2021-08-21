## [SRE Weekly Issue #168](https://sreweekly.com/sre-weekly-issue-168/) - April 14, 2019
### Articles

1. [How to Get Into SRE](https://blog.alicegoldfuss.com/how-to-get-into-sre/)

    This one’s great for folks that are new to SRE, and it’s also an enlightening read for seasoned SREs. What caught me most was the Definition section, on what it means to be an SRE.Alice Goldfuss
1. [Chaos Engineering Traps](https://medium.com/@njones_18523/chaos-engineering-traps-e3486c526059)

    In this articlization of a conference talk, the author lays out 8 common pitfalls in chaos engineering, with detailed example stories related to them. It goes much deeper than mere chaos engineering into the theory of how to operate complex systems.Nora Jones
1. [Ghosts in the machines](https://www.oreilly.com/ideas/ghosts-in-the-machines)

    Automation can have unintended effects — and can tend to not have the effect we hope it will.Thanks to Greg Burek for this one.Courtney Nash
1. [ What SREs can learn from Aviation industry? · ](https://www.anshulpatel.in/post/sre_aviation_process/)

    Anshul Patel
1. [Notes on running production code](https://kgrz.io/notes-on-running-production-code)

    Lessons learned by a software engineer on supporting their code in production.Kashyap Kondamudi
1. [The CASE Method: Better Monitoring For Humans](http://onemogin.com/monitoring/case-method-better-monitoring-for-humans.html)

    CASE stands for Context-heavy, Actionable, Symptom-based, and Evaluated. That last one’s really key. The author proposes setting an expiration time for your alerts after which time you should evaluate them to make sure that they still make sense.Cory Watson
### Outages

1. [Heroku: (EU) routing issues for ssl:endpoint applications](https://status.heroku.com/incidents/1762)
    Heroku posted this followup for an outage on April 2.
1. [The Travis CI Blog: Incident review for slow booting Linux builds outage](http://blog.travis-ci.com/2019-04-11-incident-review-slow-booting-Linux-builds-outage)
    The outage happened March 27-28.
1. [Azure VMs — North Central US](https://azure.microsoft.com/en-us/status/history/)
    Since deep-linking to Azure incident summaries doesn’t work and this one is especially interesting, I’ll quote it here:
Azure Storage team made a configuration change on 9 April 2019 at 21:30 UTC to our back-end infrastructure in North Central US to improve performance and latency consistency for Azure Disks running inside Azure Virtual Machines. This change was designed to be transparent to customers. It was enabled following our normal deployment process, first to our test environment, and lower impact scale units before being rolled out to the North Central US region. However, this region hit bugs which impacted customer VM availability. Due to a bug, VM hosts were able to establish session with the storage scale unit but hit issues when trying to receive/send data from/to storage scale unit. This situation was designed to be handled with fallback to our existing data path, but an additional bug led to failure in the fallback path and resulted in in VM reboots.
1. [Facebook, Instagram, and WhatsApp](https://www.theverge.com/2019/4/14/18310069/facebook-instagram-whatsapp-down-outage-issues)

### [ << Prev ](sreweekly-167.md) ------------- [ Next >> ](sreweekly-169.md)