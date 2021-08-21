## [SRE Weekly Issue #56](https://sreweekly.com/sre-weekly-issue-56/) - January 22, 2017
### Articles

1. [On-Call Survey](https://www.surveymonkey.com/r/SVV6CQ7)

    If you have a minute (it’ll only take one!), would you please fill out this survey? Gabe Abinante (featured here previously) is gathering information about the on-call experience with an eye toward presenting it at Monitorama.
1. [Operations For Developers – you wrote it, we all run it!](http://some.ops4devs.info/)

    Wow, what a resource! As the URL says, this is “some ops for devs info”. Tons of links to useful background for developers that are starting to learn how to do operations. Thanks to the author for the link to SRE Weekly!
1. [AWS Lambda Performance and Cold Starts](https://dzone.com/articles/aws-lambda-performance-and-cold-starts)

    AWS Lambda response time can increase sharply if your function is accessed infrequently. I love the graphs in this post.
1. [Four load testing mistakes developers love to make](https://engineering.klarna.com/four-load-testing-mistakes-developers-love-to-make-68b443f7e8a2)

    A top-notch article on how to avoid common load-testing pitfalls. Great for SREs as well as developers!
1. [Instrumentation: Worst case performance matters](https://honeycomb.io/blog/2017/01/instrumentation-worst-case-performance-matters/)

    A description of an investigation into poor performance in a service with a 100% < 5ms SLA.
1. [Docker: InfraKit Under the Hood](https://blog.docker.com/2017/01/infrakit-hood-high-availability/)

    Docker posted this article on how they designed InfraKit for high availability.
1. [Should I block ICMP?](http://shouldiblockicmp.com/)

    A blanket block of ICMP on your network device breaks some important features like ping, traceroute, MTU discovery, and the like. MTU discovery (Fragmentation Required) is especially important, and ignoring it can cause connections to appear to time out for no obvious reason.
### Outages

1. [Gears of War 4](https://gamerant.com/gears-of-war-4-server-issues-microsoft/)
1. [Pokemon Go](http://heavy.com/games/2017/01/pokemon-go-servers-down-failed-to-log-in-retry-account-issue-offline-status/)
1. [Royal Canadian Mounted Police](http://www.ctvnews.ca/politics/computer-glitch-leaves-rcmp-without-crucial-info-1.3248341)
1. [Uber](http://www.product-reviews.net/2017/01/19/uber-app-users-ask-if-system-down-jan-19/)

### [ << Prev ](sreweekly-55.md) ------------- [ Next >> ](sreweekly-57.md)