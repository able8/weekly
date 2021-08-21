## [#350 — February 19, 2021](https://golangweekly.com/issues/350)

1. [Go 1.16 Releasedrelease notes](https://golangweekly.com/link/103407/web)

     — A significant release for our favorite language! The release notes cover the details but faster and more efficient builds, the use of Go modules by default, macOS Arm64 support, and native file embedding support in executables are the order of the day.
1. [A Go Security Cheatsheet: 8 Security Best Practices](https://golangweekly.com/link/103409/web)

     — Covers things like the use of modules, reflection, containers, and crypto packages.
1. [Block Profiling in Go](https://golangweekly.com/link/103411/web)

     — Block profiling tracks when goroutines are put into a waiting state, so think select or chan receive but not time.Sleep or garbage collection.
1. [New Module Changes in Go 1.16](https://golangweekly.com/link/103437/web)

     — A key change in Go 1.16 is that module-aware mode is enabled by default even if no go.mod is present. This is in preparation for the old GOPATH mode being dropped completely in Go 1.17. There are several other changes to take into account, however.
### 📘 Tutorials and Stories

1. [Generic Functions on Slices with Go Type Parametersissue 349](https://golangweekly.com/link/103415/web)

     — Cue up the articles on how generics make things better/easier in Go! Here, Eli revisits an older post (which we linked in issue 349) to do just that.
1. [Compiling a Gameboy Advance ROM with TinyGocan target](https://golangweekly.com/link/103417/web)

     — Did you know that TinyGo can target the Gameboy Advance?
1. [A (Maybe) Unexpected Usage of the -count Flag in Go Tests](https://golangweekly.com/link/103419/web)

     — go test -count 3 will run your tests three times, but what about -count 1? Once. But it also disables test caching.
1. [Preventing SQL Injections in Go (and Other Vulnerabilities)](https://golangweekly.com/link/103421/web)

     — Handle your errors, kids, and use open-source tools to scan your code and projects for vulnerabilities.
1. [How Buffer Pool Works: An Implementation in Go](https://golangweekly.com/link/103422/web)

     — Exploring how buffer pool management works in databases by building one.
1. ['My Go Mistakes'](https://golangweekly.com/link/103423/web)

     — We learn more from failure than success, right? So why not share your learnings with other devs? Be like Henrique.
1. [Go 1.16's embed Package Explained in 5 Minutes](https://golangweekly.com/link/103412/web)

     — If you like to learn things from videos then this is the embed content you’ve been waiting for.
1. [Leveraging the Go Type System](https://golangweekly.com/link/103424/web)

     — A specific example around getting the efficiency of numeric constants with the readability of strings.
### 🛠 Code & Tools

1. [EGo: A Way to Build 'Confidential' Go Apps for Intel SGX EnclavesSGX](https://golangweekly.com/link/103425/web)

     — SGX is a set of security oriented instructions on Intel CPUs to manage private regions of memory called ‘enclaves.’ EGo wants to bring all of this to Go developers. Plus it has a cute logo :-)
1. [Logrus 1.8: A Feature Rich Structured Logger](https://golangweekly.com/link/103427/web)

     — Supports JSON formatting, hooks (for sending certain log entries to external services, say), plus it’s API compatible with the standard library’s logger too, so try dropping it in.
1. [pgCenter: Command Line Tool for Observing and Troubleshooting Postgres](https://golangweekly.com/link/103429/web)

     — Imagine something like top but for Postgres. You can keep your eye on vital statistics and running queries in real time. Linux only but written in Go.
1. [Dapr 1.0: A Portable, Event-Driven, Runtime for Distributed Appsbuilt in Gooriginal announcement blog post from 2019.](https://golangweekly.com/link/103430/web)

     — Dapr, an open source project out of Microsoft, provides best practices and building blocks for microservices and is language-agnostic (though it’s built in Go itself). It also provides bindings for common infrastructure items, such as pub/sub messaging. There’s more detail in the original announcement blog post from 2019.
1. [textnote: Simple Tool for Daily Note-Taking from the Command Line](https://golangweekly.com/link/103433/web)

     — Supports monthly archiving of notes, as well.
### 💻 Jobs

1. [Sr Back-End Software Engineer at Mode (Remote)apply here.](https://golangweekly.com/link/103434/web)

     — Mode is a powerful analytics platform, turning data into actionable insights in one, central location. To learn more, apply here.
1. [Sr. Software Engineer at CrowdStrike (Remote)](https://golangweekly.com/link/103473/web)

     — CrowdStrike is the leader in cloud-delivered endpoint protection, which helps protect our customers from cybersecurity attacks.
1. [Find Your Next Job Through Hired](https://golangweekly.com/link/103436/web)

     — Create a profile on Hired to connect with hiring managers at growing startups and Fortune 500 companies. It's free for job-seekers.

### [ << Prev ](golangweekly-349.md) ------------- [ Next >> ](golangweekly-351.md)