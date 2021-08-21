## [SRE Weekly Issue #274](https://sreweekly.com/sre-weekly-issue-274/) - June 13, 2021
### Articles

1. [Chicken Soup for the SLO](https://devops.com/chicken-soup-for-the-slo/)

    The last section suggests selling SLOs to executives by likening them to OKRs or KPIs.Austin Parker — Devops.com
1. [How Lowe’s leverages Google SRE practices](https://cloud.google.com/blog/products/devops-sre/how-lowes-leverages-google-sre-practices/)

    Lowe’s is a home improvement retailer in North America. I often find it fascinating when I learn that a company that’s not seen as being in the tech-sector has a robust SRE practice.Vivek Balivada and Rahul Mohan Kola Kandy — Lowe’s
1. [Incident writeup as sociological storytelling](https://surfingcomplexity.blog/2021/06/11/incident-writeup-as-sociological-storytelling/)

    Lorin Hochstein
1. [DevOps & Autism Care](https://www.usenix.org/publications/loginonline/devops-autism-care)

    This is brilliant: they apply DevOps and SRE practices to the challenging work of raising two autistic children.Zac Nickens — USENIX ;login:
1. [Implementing ChatOps into our Incident Management Procedure](https://shopify.engineering/implementing-chatops-into-our-incident-management-procedure)

    I especially like how their bot automatically pages reinforcements after folks have been on an incident for long enough to become fatigued.Daniella Niyonkuru
1. [The MTTR that matters](https://firehydrant.io/blog/the-mttr-that-matters)

    Rather than measuring Mean Time To Recovery for incidents, let’s track our Mean Time To Retrospective.Robert Ross — FireHydrant
### Outages

1. [Fastlysummary](https://status.fastly.com/incidents/vpk0ssybt3bj)
    Fastly had a global outage of their CDN service, with many 5xx errors for around 40 minutes and diminished cache hit ratios following after. Many customers of Fastly experienced degradation, notably including Amazon, Reddit, and GitHub, among many others.
Fastly posted a summary shortly after the incident, describing a latent bug that was triggered by a customer’s (valid) configuration change.
Full disclosure: Fastly is my employer.
1. [Salesforce](https://status.salesforce.com/generalmessages/709)
1. [Facebook, Instagram, and WhatsApp](https://www.pakistanchristian.tv/facebook-crashes-again-and-with-it-instagram-and-whatsapp-its-not-fun-anymore/)

### [ << Prev ](sreweekly-273.md) ------------- [ Next >> ](sreweekly-275.md)