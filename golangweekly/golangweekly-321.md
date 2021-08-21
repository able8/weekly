## [#321 — July 17, 2020](https://golangweekly.com/issues/321)

1. [Generics and Parentheseshere’s a screenshot of the post](https://golangweekly.com/link/92076/web)

     — Using parentheses () to declare or instantiate types raised numerous concerns but the core team is now experimenting with square brackets [] instead – can you find any issues with this? For example, what might have previously been func f((T(int)) could instead be func f(T[int]) which, frankly, reads a lot better to us.(Note: Annoyingly, a Google login is required to read Groups posts. If you don’t want to do this, here’s a screenshot of the post you can read instead with the text in the description.)
1. [Making Pong with Go and WebAssembly](https://golangweekly.com/link/92078/web)

     — Another fantastic example of using Go to build something that works in the browser by way of WebAssembly.
1. [Go 1.14.6 and Go 1.13.14 ReleasedMore info here.](https://golangweekly.com/link/92079/web)

     — 1.14.5 and 1.13.13 also came out a few days ago which covered two security fixes: data race in net/http and a X.509 verification issue on Windows. These two extra releases are more focused on minor fixes. More info here.
### 💻 Jobs

1. [Sr. Software Engineer at CrowdStrike (Remote)](https://golangweekly.com/link/91966/web)

     — CrowdStrike is the leader in cloud-delivered endpoint protection, which helps protect our customers from cybersecurity attacks.
1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/91967/web)

     — Join the most energizing community for developers and work on projects for Riot Games, FOX, Sony, Coinbase, and more.
1. [Find a Job Through Vettery](https://golangweekly.com/link/91968/web)

     — Use Vettery to connect with growing tech teams at startups and Fortune 500 companies.
### 📚 Articles & Tutorials

1. [How Are Deadlocks Triggered?](https://golangweekly.com/link/92085/web)

     — Go has a deadlock detector, but how does it work? And what are its limitations?
1. [Unit Test Outbound HTTP Requests](https://golangweekly.com/link/92086/web)

     — Go makes it very easy to mock services, but Erik goes one step further and records external service responses so you can run your tests dependency-free.
1. ["Interface Smuggling", a Go Design Pattern for Expanding APIs](https://golangweekly.com/link/92087/web)

     — In short, check to see if a type supports a broader API and use it.
1. [PopCount on ARM64 in Go Assembler](https://golangweekly.com/link/92088/web)

     — A useful CPU extension, a really interesting architecture, and Go’s offbeat assembler
1. [Under The Hood of context](https://golangweekly.com/link/92089/web)

     — A guided tour through context’s source with the intention being to reduce the author’s worries about when and why to use it.
1. [Use Cases for Channels](https://golangweekly.com/link/92090/web)

     — This isn’t new but you might find it useful nonetheless and it’s delightfully code and example heavy.
1. [Your First Week with Go](https://golangweekly.com/link/92091/web)

     — Two developers (Jacquie Grindrod and DaShaun Carter) join Jon Calhoun to talk about their first week with Go, what worked for them and what didn’t.
### 🛠 Code & Tools

1. [Afero: A Filesystem Abstraction System for Go](https://golangweekly.com/link/92092/web)

     — A single consistent API for accessing a variety of filesystems. Also lets you create mock and testing filesystems that don’t rely on disk at all.
1. [errcheck: Checking That You Checked for Errors 🐞](https://golangweekly.com/link/92093/web)

     — Checking for errors is a fundamental part of the Go experience and this tool will help you check that you checked!
1. [go-ordered-map: An Implementation of Ordered Maps](https://golangweekly.com/link/92094/web)

     — Think of a map that remembers the order key-value pairs were added, so you can iterate over map.Newest().
1. [gRPCurl: Like curl, But for gRPC](https://golangweekly.com/link/92095/web)

     — A command line tool for interacting with gRPC servers.
1. [oapi-codegen: Generate Go Client and Server Boilerplate From OpenAPI 3 Specifications](https://golangweekly.com/link/92096/web)

1. [gorush: A Push Notification Server Written in Go](https://golangweekly.com/link/92097/web)

     — Supports APNS (Apple Push Notification Service) and Firebase. Now supports concurrent pushes for iOS.
1. [Testify: Assertions and Mocks That Play Nicely with testing](https://golangweekly.com/link/92098/web)

1. [go-tagexpr: An Interesting Way to Add Tag Expressions to Structs](https://golangweekly.com/link/92099/web)

     — The main use case so far is to define validations.

### [ << Prev ](golangweekly-320.md) ------------- [ Next >> ](golangweekly-322.md)