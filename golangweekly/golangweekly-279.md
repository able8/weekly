## [#279 — September 13, 2019](https://golangweekly.com/issues/279)

1. [The Epic Tale of Tracking Down a 'Memory Leak' in a Go Microservice](https://golangweekly.com/link/73791/web)

     — An interesting step-by-step tale of how Detectify’s backend team tracked down what seemed to be a memory leak.. but it was down to a change in how Go 1.12 works.
1. [Building a Better Go Linker](https://golangweekly.com/link/74338/web)

     — Go’s linker combines the outputs of the Go compiler into the eventual executables you run, but its (lack of) efficiency is causing problems as Go develops. Google’s Austin Clements explains the problem and looks at how to build a better linker for Go.
1. [Programming with errors](https://golangweekly.com/link/73794/web)

     — Some tips on using Go 1.13’s errors package (formerly xerrors).
1. [Don’t Force Allocations on the Callers of Your API](https://golangweekly.com/link/73792/web)

     — We all know about the ‘evils’ of premature optimization but when it comes to designing an API that you might not want to change later.. it’s worth thinking about the performance implications of the choices you make.
1. [Go is IEEE's 10th Top Programming Language of 2019](https://golangweekly.com/link/73795/web)

     — Python, Java, and C sit at the top of the table, but given Go’s relative youthfulness, seeing it at number ten is pretty neat.
1. [Discussing Generics with Ian Lance Taylor](https://golangweekly.com/link/73796/web)

     — What are generics and why are they useful? Why aren’t interfaces enough? How will the standard library change if generics are added to Go? These questions and more are discussed on the latest Go Time episode. (54 minutes.)
### 💻 Jobs

1. [Stream Is Hiring a Go Developer in Amsterdam](https://golangweekly.com/link/73797/web)

     — Like building scalable infrastructure in Go? We do too. Stream is hiring. Apply now.
1. [Sourcegraph Is Growing. Apply Now to Grow with Us](https://golangweekly.com/link/73798/web)

     — Build the best large scale developer tools. Product-market fit, clear goals, talented team, autonomy, open source, remote first.
1. [Find a Go job through Vettery](https://golangweekly.com/link/73799/web)

     — Make a free profile, name your salary, and connect with hiring managers from top employers.
### 📘 Articles & Tutorials

1. [How Mailgun Built a Lucene-Inspired Parser in Go](https://golangweekly.com/link/73800/web)

     — A team at Mailgun created a system to let other employees write custom queries that could match against real time events taking place within Mailgun’s systems.
1. [Go Modules and The Problem of Noticing Updates to Dependencies](https://golangweekly.com/link/73801/web)

     — Questions from a cautiously optimistic sysadmin about how Go modules will report changes (if at all) and what it does when a dependency disappears?
1. [Creating Custom Errors in Go](https://golangweekly.com/link/73802/web)

     — Sometimes errors.New and fmt.Errorf aren’t quite enough for communicating complicated error information to your users (or to your future self!) Here’s a look at what more you can do.
1. [Using SO_PEERCRED in Go](https://golangweekly.com/link/73803/web)

     — A practical look at a slightly obscure Unix domain socket option that, intriguingly, provides the server process with user, group, and process IDs of any connected client.
1. [Making Your Own Changes to Things That Use Go Modules](https://golangweekly.com/link/73804/web)

     — If you’re fiddling around with the dependencies of a program that uses Go modules, there’s some extra work to do to keep things working.
1. [The Basics of Implementing a Linked List in Go](https://golangweekly.com/link/73805/web)

1. [Using Go 1.13’s New ReportMetric API to Benchmark t-digest Implementations](https://golangweekly.com/link/73807/web)

1. [Bad Go: Not Sizing Slices](https://golangweekly.com/link/73808/web)

     — If you know how big a slice needs to be, set the size and enjoy a performance boost.
1. [if err != nil: Let's Talk About Errors in Gonow aborted](https://golangweekly.com/link/73809/web)

     — A podcast from earlier this year including folks like Dave Cheney and Peter Bourgon discussing error handling in Go and the now aborted try proposal.
### 🛠 Code & Tools

1. [Joe Bot: A General Purpose Chat Bot LibraryHubot](https://golangweekly.com/link/73811/web)

     — Inspired by GitHub’s Hubot, but written in Go, Joe Bot provides a framework and abstractions for building chat bots.
1. [GoVector: Vector Clock Logging Library](https://golangweekly.com/link/73813/web)

     — A vector clock algorithm is used to order events in a distributed system in lieu of a centralized clock. GoVector started as a teaching tool.
1. [Cadence: A Distributed Orchestration Engine Built in GoGo client](https://golangweekly.com/link/73815/web)

     — This system, built by Uber, simplifies the development of complex stateful distributed applications. This is the server but there are also Java and Go client libraries.
1. [quic-go: A QUIC Implementation in Pure Gosuch as in Caddy.](https://golangweekly.com/link/73817/web)

     — QUIC is a transport layer network protocol originally developed at Google that’s forming the basis for HTTP/3. This library may not be of direct interest to many but it can form the basis of HTTP/3 implementations in other projects, such as in Caddy.
1. [GoLand 2019.3 Opens Its Early Access Program](https://golangweekly.com/link/73819/web)

     — GoLand is a (commercial) IDE for Go developers from JetBrains but you can use/trial these ‘early access’ versions for 30 days.
1. [Using GitHub Actions as CI Effectively for Go](https://golangweekly.com/link/73820/web)

     — Docs and examples on how to do it.

### [ << Prev ](golangweekly-278.md) ------------- [ Next >> ](golangweekly-280.md)