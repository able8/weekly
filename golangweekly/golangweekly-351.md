## [#351 — February 26, 2021](https://golangweekly.com/issues/351)

1. ['Go is not an easy language.'some interesting discussion](https://golangweekly.com/link/103788/web)

     — The syntax is simple, the semantics are simple, but doing useful stuff is ‘not always easy in Go.’ The point made by the author here is rather simple IMHO, but it provoked some interesting discussion (423 comments!) on Hacker News. Maybe Go is like the guitar.. easy to learn but hard to master?
1. [A Proposal to Add Fuzzing to the Go Standard Librarydesign draft here](https://golangweekly.com/link/103790/web)

     — Fuzzing is a type of testing that mutates inputs to a program under test to catch issues that common testing or humans would miss. The proposal is targeting Go 1.17 as an experimental feature and there’s a design draft here.
1. [Contexts and Structs](https://golangweekly.com/link/103793/web)

     — A new post on the official Go blog digging into context.Context and why contexts should be explicitly passed around to functions that need it rather than tucked away inside a struct.
1. [A Rundown of What's New in Go 1.16release of Go 1.16](https://golangweekly.com/link/103794/web)

     — Last week we featured the release of Go 1.16 but this post takes a more accessible walk through a few of the updates.
### 📘 Tutorials and Stories

1. [The Life of an HTTP Request in a Go Server](https://golangweekly.com/link/103796/web)

     — Eli continues his streak of solid blog posts with a look at the route a typical HTTP request takes through a Go server.
1. [Injecting Values into Variables at Build Time-X option of Go’s linker](https://golangweekly.com/link/103797/web)

     — What if you want to set the value of certain variables without those values being present in source control? You could load them in at runtime, or… use the -X option of Go’s linker.
1. [Writing a Simple TCP Server Using Kqueue](https://golangweekly.com/link/103800/web)

     — Kqueue provides event notification functionality to FreeBSD-based operating systems and writing a basic echo server is a good way to learn how it works.
1. [Building Solid Go GraphQL Applications Quickly](https://golangweekly.com/link/103801/web)

     — Some thoughts from someone who’s been there. Testing is the key to confidence, Peter reckons.
1. [Context and Variables in the Hugo Static Site Generator](https://golangweekly.com/link/103802/web)

     — Hugo is a popular static site generator built in Go and this tutorial introduces some more advanced concepts around how Hugo handles context.
1. [Rapidly Developing Go Microservices on Kubernetes with TelepresenceTelepresence](https://golangweekly.com/link/103803/web)

     — Written by a developer advocate promoting Telepresence, a (commercial) tool this tutorial is oriented around, so keep that in mind.
1. [Why a Team is Switching From C# to Go for Backend Development](https://golangweekly.com/link/103805/web)

### 🛠 Code & Tools

1. [Dither: A Fast and 'Correct' Image Dithering Library](https://golangweekly.com/link/103806/web)

     — There’s something about a well dithered image that takes me back to the 1990s :-) This library supports quite a few approaches.
1. [GoMock 1.5.0: A Mocking Framework for Go](https://golangweekly.com/link/103807/web)

     — Has two modes of operation: source and reflect. Source mode generates mock interfaces from a source file. Reflect mode generates mock interfaces by building a program that uses reflection to understand interfaces.
1. [Redigo: A Go Client for Redis](https://golangweekly.com/link/103809/web)

     — Supports all the Redis commands through a Print-like API and also offers connection pooling, helper functions for handling replies, and pipelining support.
1. [go-ipfs 0.8.0: An IPFS Implementation for Gov0.8.0](https://golangweekly.com/link/103810/web)

     — IPFS (InterPlanetary File System) is a peer-to-peer protocol and network for organizing a distributed file system. v0.8.0 of this Go implementation has added quite a few new features.
1. [Sloc, Cloc and Code (scc) 3.0: A Fast Accurate 'Code Counter'v3.0](https://golangweekly.com/link/103812/web)

     — For multiple languages, counts lines of code and estimate code complexity. v3.0 moves things up to Go 1.16, adds M1 and ARM support, and more.
### 💻 Jobs

1. [Sr. Software Engineer at CrowdStrike (Remote)](https://golangweekly.com/link/103814/web)

     — CrowdStrike is the leader in cloud-delivered endpoint protection, which helps protect our customers from cybersecurity attacks.
1. [Senior Software Engineer (Go) - 100% Remote, UK/EU Only](https://golangweekly.com/link/103815/web)

     — Solving complex technical challenges using the latest, cloud-native technologies to help Banks and FinTechs move money faster.
1. [Find Your Next Job Through Hired](https://golangweekly.com/link/103816/web)

     — Create a profile on Hired to connect with hiring managers at growing startups and Fortune 500 companies. It's free for job-seekers.

### [ << Prev ](golangweekly-350.md) ------------- [ Next >> ](golangweekly-352.md)