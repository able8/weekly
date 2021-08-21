## [#339 — November 20, 2020](https://golangweekly.com/issues/339)

1. [Go Standard Library Benchmarks: Intel i5 vs Apple's M1](https://golangweekly.com/link/98971/web)

     — If you’ve wondered how Apple’s newest Arm-based CPU fares with Go, these results are promising. As with all benchmarks, though, maintain a critical eye, especially as these are microbenchmarks against an older i5 CPU(!) but it's still worth reporting and there's more info below..
1. [The Pros of Conds](https://golangweekly.com/link/98972/web)

     — If you’ve never used sync.Cond (and it’s not that popular in Go code) Luke wants to sell you on why and where it’s useful. If you have events you need to sit around and wait for, this article is for you.
1. [Thread: "Let's See How The New M1 Chips Do On Go Benchmarks!"are mandatory on arm64](https://golangweekly.com/link/98974/web)

     — Filippo Valsorda of the Go team does eventually get to the benchmarks but the real takeaway here is his journey (signed binaries are mandatory on arm64?) to get there. If you’re thinking of buying a new Mac, read this.
1. [Go’s Recurring Security Problem](https://golangweekly.com/link/98976/web)

     — “go get does a lot under the hood, including invoking third-party tools like git and clang in ways that are heavily influenced by package configurations. Ensuring that these invocations are safe is an uphill battle that Go hasn’t quite won yet.”
### 💻 Jobs

1. [Mux Is Hiring Across the Board to Help Build the Stripe for Video](https://golangweekly.com/link/98982/web)

1. [Find a Job Through Vettery](https://golangweekly.com/link/98983/web)

     — Create a profile on Vettery to connect with hiring managers at startups and Fortune 500 companies. It's free for job-seekers.
### 📘 Tutorials and Stories

1. [Writing a Postgres Foreign Data Wrapper for Clickhouse in GoClickHouse](https://golangweekly.com/link/98984/web)

     — Foreign data wrappers (FDWs) provide a way for Postgres servers to interact with external services (including other database systems), such as the ClickHouse OLAP DBMS covered here.
1. [Using 'Dot Imports' in Views](https://golangweekly.com/link/98986/web)

     — Dot imports let you import a package into the local namespace (so no prefixing is needed!) – it’s not a good practice, but also not super commonly known, and Markus has speculated on an interesting use case for them at least.
1. [Good and Bad Practices - A Limiting Perspective?](https://golangweekly.com/link/98987/web)

     — Certain patterns or approaches can get painted as being ‘good’ or ‘bad’ practices (see ‘dot imports’ above!) but is it always so binary? No, says Preslav.
1. [An Introduction to AWS Lambda Functions in Go](https://golangweekly.com/link/98989/web)

     — This detailed look includes forays into logging, triggers (including API Gateway), using Context, and versioning the functions.
1. [Refactoring for Better Testability](https://golangweekly.com/link/98990/web)

     — The first in a series of posts on refactoring existing code for testability. This is good practice! ;-)
1. [How to Build Your Own Serverless Subscriber List with Go and AWSSimple Subscribe](https://golangweekly.com/link/98991/web)

     — How to build an email subscription sign up flow (complete with double opt in) using Lambda and DynamoDB. Can’t be bothered? Simple Subscribe is Victoria’s project with it all done for you ;-)
1. [Go Constants and JSON: To Iota and Back](https://golangweekly.com/link/98993/web)

     — Constants and iota are very useful…up until you combine them with JSON serialization. Thankfully, there’s a Marshal (and Unmarshal) in town to help.
1. [Checking If The Pi Is Done](https://golangweekly.com/link/98994/web)

     — Build a basic cluster monitoring system using a laptop, some Raspberry Pis, and Go’s RPC package.
1. [Building a WebRTC Video and Audio Broadcaster in Go using ION-SFU and Media Devices](https://golangweekly.com/link/98995/web)

### 🛠 Code & Tools

1. [go-getter: Download Things from Various Sources with a URL](https://golangweekly.com/link/98996/web)

     — Reading the title you might think.. “so it’s an HTTP client then?” Not exactly. A GitHub URL can become a git URL, local filenames work, S3, Mercurial and BitBucket URLs all work as intended too.
1. [Peer Calls: Group Peer-to-Peer Video Call Systemhere.](https://golangweekly.com/link/98997/web)

     — Want to get in on the action of rolling out a video chat service? The core of this one is written in Go. You can play with an example deployment here.
1. [eBPF: A Pure Go eBPF LibraryeBPF](https://golangweekly.com/link/98999/web)

     — Utilities for loading, compiling, and debugging eBPF programs.
1. [p5: An Initial Port of Processing to GoProcessing](https://golangweekly.com/link/99002/web)

     — It’s early days for this project, but this brings some initial concepts from Processing, a popular visual programming system, to Go.
1. [Fleets: Automatically Delete Tweets, Retweets, and LikesAnaconda](https://golangweekly.com/link/99004/web)

     — A new, craftily-named library for cleaning up your Twitter account that leans heavily on the Anaconda Go Twitter library.
1. [gdriver: Download Large Files From Google Drive](https://golangweekly.com/link/99006/web)

     — A command-line tool for downloading large files from Google Drive (using API v3).
1. [h2go: Apache H2 Go SQL Driver](https://golangweekly.com/link/99007/web)

     — Experimental, right now.

### [ << Prev ](golangweekly-338.md) ------------- [ Next >> ](golangweekly-340.md)