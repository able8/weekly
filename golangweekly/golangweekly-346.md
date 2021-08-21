## [#346 — January 22, 2021](https://golangweekly.com/issues/346)

1. [Go 1.15.7 and 1.14.14 Releasedan errorhere](https://golangweekly.com/link/101752/web)

     — (Requires Google login.) Two new releases to primarily fix two reported security issues around arbitrary code execution (covered in more depth in the item below) and an error in crypto/elliptic. Downloads here.
1. [Command PATH Security in Go](https://golangweekly.com/link/101753/web)

     — Go 1.15.7 and Go 1.14.14 address a arbitrary code execution bug that can occur around go get and other commands that build code. This bug is significant enough to get a full write up with where it came from and ways to improve the situation.
1. [AWS SDK for Go Version 2 Now Generally Available](https://golangweekly.com/link/101755/web)

     — There’s a lot to like in version 2, including modularized clients per service, better config and pagination, and performance improvements. Requires Go 1.15 or higher.
1. [Go's ioutil Package to Be Deprecated in 1.16by Russ Cox](https://golangweekly.com/link/101756/web)

     — Most of the useful functions in this package (called a “poorly defined and hard to understand collection of things” by Russ Cox) have been moved to io and os, so this is really just clean up.
### 📘 Tutorials and Stories

1. [Packages as Layers, Not Groups](https://golangweekly.com/link/101757/web)

     — Looking at your packages as layers is a step toward a hexagonal architecture and can clean up your code structure significantly.
1. [How to Debug Go-Powered AWS Lambda Functions with GebugGebug](https://golangweekly.com/link/101758/web)

     — Go is a fantastic language for serverless use cases but local debugging can be a pain point. Here are some pointers that involve Gebug, a tool for debugging Dockerized Go apps.
1. [Non-Blocking Parallelism for Services in Go](https://golangweekly.com/link/101760/web)

     — Dubbed “The Tickler Pattern”, it’s a way to handle a queue using channels and semaphores that isn’t too complicated and won’t block the client adding work to the queue.
1. [Go Pointers Explained, Once and for All](https://golangweekly.com/link/101762/web)

     — A simple, practical introduction to pointers in a 14-minute screencast covering their use in implementing stacks and heaps.
1. [Errors vs. Exceptions in Go and C++ in 2020](https://golangweekly.com/link/101763/web)

     — Subtitled “Why and how exceptions are still better for performance, even in Go”, this is a return to a study done in 2018 and the conclusions are consistent.
1. [But How, Exactly, Do Databases Use mmap?BoltDB](https://golangweekly.com/link/101797/web)

     — A look through the BoltDB source code to understand how memory mapping can work in a database system. BoltDB, by the way, is a Go-powered key/value story inspired by LMDB.
### 🛠 Code & Tools

1. [go-enry: A Faster File Programming Language DetectorLinguist](https://golangweekly.com/link/101764/web)

     — When they say ‘faster’, they’re comparing to the Ruby-powered Linguist, as used by GitHub for its repo language detection.
1. [Goridge: A High-Performance PHP-to-Go IPC Bridge](https://golangweekly.com/link/101766/web)

     — Using native sockets in PHP and net/rpc in Go, this lets you call Go methods from PHP, if you so require.
1. [Dragonboat 3.3: A High Performance Multi-Group Raft Librarywhat’s new](https://golangweekly.com/link/101767/web)

     — Consensus algorithms (such as Raft) provide fault-tolerance by allowing a system to continue to operate as long as a majority of member servers are available. Here’s what’s new in 3.3.
1. [BFE: A Modern Layer 7 Load Balancer From Baidu](https://golangweekly.com/link/101770/web)

     — Baidu operates a popular Chinese search engine so they’re used to working at scale.
1. [wuzz: Interactive CLI Tool for HTTP Inspection](https://golangweekly.com/link/101771/web)

     — Think curl but with a text UI to show off request and response headers, manage params, etc.
### 💻 Jobs

1. [Senior Software Engineer (Go) - 100% Remote, UK/EU Only](https://golangweekly.com/link/101772/web)

     — Solving complex technical challenges using the latest, cloud-native technologies to help Banks and FinTechs move money faster.
1. [Find a Job Through Hired](https://golangweekly.com/link/101773/web)

     — Create a profile on Hired to connect with hiring managers at growing startups and Fortune 500 companies. It's free for job-seekers.
### 🥇 A Golden Oldie Revisited

1. [termenv: Advanced ANSI Style and Color Support for Terminal Apps](https://golangweekly.com/link/101776/web)

     — Offers convenient methods to colorize and style output, without needing to deal with weird ANSI escape sequences and color conversions.

### [ << Prev ](golangweekly-345.md) ------------- [ Next >> ](golangweekly-347.md)