## [#352 — March 5, 2021](https://golangweekly.com/issues/352)

1. [Creating a Self Contained Blog Server with Go 1.16embedding](https://golangweekly.com/link/104247/web)

     — Brian has been playing with the new embedding functionality in Go 1.16 and has come up with a practical way to create a Web server that embeds all the static assets needed to serve up a site in a single binary (though the actual content is pulled from elsewhere, but compiling everything in could be a fun experiment.) :-)
1. [An Introduction to Genericsgo2go playground](https://golangweekly.com/link/104250/web)

     — A nice tutorial that takes a step approach to teach generics starting from duplicating code for types, then to interface{}, and finally to generics. And you can run it all in the go2go playground.
1. [A Practical Understanding of Scheduler Semantics](https://golangweekly.com/link/104252/web)

     — Esteemed Go trainer Bill Kennedy gave a talk for the Golang NYC meetup recently covering what you need to understand to create effective multi-threaded concurrent code. An hour long session that focuses more on concepts under the hood than Go code itself.
1. [Lift-Off for Wasmer Go Embedding 1.0Wasmer](https://golangweekly.com/link/104253/web)

     — Good news though possibly hard to understand if you’re not into WebAssembly (the portable bytecode format). Basically, Wasmer is a runtime for WebAssembly, and this 1.0 news pushes forward Go as a first-class language for embedding Wasmer to run WebAssembly code within Go programs.
### 💻 Jobs

1. [Sr. Software Engineer at CrowdStrike (Remote)](https://golangweekly.com/link/104255/web)

     — CrowdStrike is the leader in cloud-delivered endpoint protection, which helps protect our customers from cybersecurity attacks.
1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/104256/web)

     — Join the most energizing community for developers and work on projects for Riot Games, FOX, Sony, Coinbase, and more.
1. [Find Your Next Job Through Hired](https://golangweekly.com/link/104257/web)

     — Create a profile on Hired to connect with hiring managers at growing startups and Fortune 500 companies. It's free for job-seekers.
### 📘 Tutorials and Stories

1. [Putting Text on an Image with Goa project](https://golangweekly.com/link/104258/web)

     — A developer new to Go decided to use Go to dynamically create images with quotes on them. This is both how he did it and also a project which explains how to use it on Google Cloud Functions, Heroku, or with Docker.
1. [Finding 'Evil' Go Packages](https://golangweekly.com/link/104260/web)

     — The specific search, in this case, is for packages ‘typosquatting’ popular packages. The search revealed two suspicious repos, one of which has seemingly been taken down by GitHub.
1. [Take Your First Steps with Go with Microsoft](https://golangweekly.com/link/104262/web)

     — It might not suit you but could be a handy link to pass on to anyone you know who would like to learn Go. This suite of tutorials from Microsoft goes from absolute basics and covers tooling as well as Go language basics.
1. [Using Delve to Debug Go Programs on Red Hat Enterprise Linux](https://golangweekly.com/link/104263/web)

     — From RHEL 8.2 onward the Delve debugger will come with the Go toolchain via the go-toolset package.
1. [A Few Go JSON Gotchas That Annoyed Me But I Have Learned to Deal With](https://golangweekly.com/link/104264/web)

     — In short, error checking struct tags and JSON input is on you.
### 🛠 Code & Tools

1. [generativeart: Go Code for Creating Generative ArtGenerative Art in Golang](https://golangweekly.com/link/104265/web)

     — Go full psychedelic with this library by creating swirls, skies, mazes, spirals, and more. In related news, Generative Art in Golang is an ebook we've seen recently on this very topic and we hope to review it soon.
1. [gosec 2.7: A Go Code Security Checker](https://golangweekly.com/link/104267/web)

     — Runs a series of rules over some Go code’s abstract syntax tree to spot potential issues.
1. [Inlets 3.0: Expose Your Local Endpoints to the Internetv3.0.0](https://golangweekly.com/link/104268/web)

     — A ‘cloud native tunnel written in Go’ that combines a reverse proxy and WebSocket tunnels to expose local endpoints to the public Internet via an exit node. v3.0.0 dropped recently.
1. [sqlc 1.7.0: Generate Type Safe Go From SQL1.7.0](https://golangweekly.com/link/104271/web)

     — sqlc supports a nice set of features, including transactions, array data types, and various migration tools. 1.7.0 adds initial support for Windows.
1. [sq: A General 'Swiss-Army' Tool for DataxsvjqGitHub repo](https://golangweekly.com/link/104273/web)

     — There are lot of multi-tool style data tools around (such as xsv for CSV or jq for JSON) – think of this one as being like jq but for all sorts of formats and systems, and it’s written in Go. GitHub repo.
1. [lungoDB: A MongoDB-Compatible Embeddable Database for Go](https://golangweekly.com/link/104277/web)

     — Imagine something like SQLite but MongoDB-flavored and aimed at Go developers. There are some use cases outlined in the README here.
1. [RoadRunner 2.0: A High Performance PHP Application ServerGitHub repo](https://golangweekly.com/link/104278/web)

     — PHP, you say? This is a Go-based system that keeps PHP under control with a load balancer and process manager built in. GitHub repo.
1. [Song2: Fast Gaussian Blur in Go](https://golangweekly.com/link/104280/web)

     — Woo-hoo. A linear time Gaussian blur implementation. Nice musical reference in the project name too, although I was more of an Oasis fan...

### [ << Prev ](golangweekly-351.md) ------------- [ Next >> ](golangweekly-353.md)