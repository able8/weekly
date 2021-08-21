## [SRE Weekly Issue #197](https://sreweekly.com/sre-weekly-issue-197/) - December 8, 2019
### Articles

1. [Why build an online community around “learning from incidents”?](https://www.learningfromincidents.io/blog/learning-from-incidents-in-software)

    Here’s an intro to the Learning From Incidents community. I can’t wait to see what these folks write. They’re coming out of the gate fast, with a post every day for the first week.Nora Jones
1. [OOPS! Learning from the incident you didn’t have](https://www.learningfromincidents.io/blog/oops-learning-from-the-incident-you-didnt-have)

    I love the move toward using the term “operational surprise” rather than “incident”.Lorin Hochstein
1. [The Space Review: All in the family](https://www.thespacereview.com/article/3842/1)

    Fascinating detail about the space shuttle Columbia’s accident, and the confusing jargon at NASA that may have contributed.Dwayne A. Day — The Space Review
1. [The Art of SLOs](https://landing.google.com/sre/resources/practicesandprocesses/art-of-slos/)

    Google released free material (slides, handbooks, worksheets) to help you run a workshop on effective SLOs.
1. [Eliminating toil with fully automated load testing](https://engineering.linkedin.com/blog/2019/eliminating-toil-with-fully-automated-load-testing)

    Lots of really interesting detail about how LinkedIn routes traffic to datacenters and what happens when a datacenter goes down.Nishant Singh — LinkedIn
1. [Patience in Implementing Effective Incident Reviews](http://willgallego.com/2019/12/07/patience-in-implementing-effective-incident-reviews/)

    Our field is learning a ton, and it can be tempting to short-circuit that learning.  It takes time to really grok and integrate what we’re learning.Will Gallego
1. [Shrinking the time to mitigate production incidents](https://cloud.google.com/blog/products/management-tools/shrinking-the-time-to-mitigate-production-incidents/)

    I like the distinction between “unmanaged” and “untrained” incident response.Author: Jesus Climent — Google
1. [Journey into Observability: Reading material](https://mads-hartmann.com/sre/2019/08/04/journey-into-observability-reading-material.html)

    This chronicle of learning about observability makes for an excellent reading list to those just diving in.Mads Hartmann
### Outages

1. [GitLab — Analysis of November 28th outage](https://gitlab.com/gitlab-com/gl-infra/production/issues/1421)
    A change to roll out ip-tables to other non gitlab.com hosts was inadvertently applied to the database hosts. That change to host firewalling caused all web and api hosts to lose connectivity to the database. The change has been rolled back and we are now restarting host processes.
1. [Disney+](https://popculture.com/streaming/2019/12/04/disney-plus-down-streamer-experiences-technical-problems/)
1. [Dexcom Diabetes Alerts](https://fortune.com/2019/12/02/dexcom-outage-blackout-diabetes-patients-blood-sugar-monitor/)
    Blood sugar monitors failed to send alerts for days. Parents use these monitors for monitoring their diabetic children’s blood sugar levels.
1. [AOL Mail](https://babblesports.com/tech/aol-mail-down-due-to-outage-user-getting-502-bad/)
1. [DRS black-out during Abu Dhabi GP](https://www.racefans.net/2019/12/01/server-crash-caused-drs-black-out-during-abu-dhabi-gp/)
    A Service failure prevented drivers from being allowed to use the Drag Reduction System (DRS).
1. [Discordanother incident](https://discord.statuspage.io/incidents/vs0q8cvycg46)
    Related to the Google Compute Engine incident below.
They also had another incident today.
1. [Heroku incident #1930 followup](https://status.heroku.com/incidents/1930)
1. [Heroku](https://status.heroku.com/incidents/1933)
1. [Google Compute Engine](https://status.cloud.google.com/incident/compute/19012)
    High latency in IO operations to SSD-based persistent disks.

### [ << Prev ](sreweekly-196.md) ------------- [ Next >> ](sreweekly-198.md)