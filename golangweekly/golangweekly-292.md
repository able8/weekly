## [#292 — Decemeber 13, 2019](https://golangweekly.com/issues/292)

1. [Introducing sqlc: Compile SQL Queries to Type-Safe Go](https://golangweekly.com/link/81287/web)

     — sqlc generates readable, idiomatic Go from SQL so you can interact with your database without field tags or sacrificing type safety.
1. [Using Makefile(s) with Go](https://golangweekly.com/link/81288/web)

     — We’ve seen a few articles around using make, but this one is especially detailed with an eye on being productive in a real scenario.
1. [Versioning with Branching Strategy](https://golangweekly.com/link/81290/web)

     — The Go team recently recommended copying different versions of your library into different directories. Here is an alternate (and, seemingly, better) approach based on Git tags and branches.
1. [Dynamically Scoped Variables in Go?](https://golangweekly.com/link/81291/web)

     — “This is a bad idea, no argument there. This is not a pattern you should ever use in production code. But, this isn’t production code, it’s a test, and perhaps there are different rules that apply to test code.” An interesting experiment, even if it doesn’t smell quite right..
1. [Bill Kennedy's Go Reading List](https://golangweekly.com/link/81292/web)

     — Whenever esteemed Go trainer Bill Kennedy comes across a significant interview, blog, historical piece, book, or similar thing about Go, he puts it in this list. If you’re looking for some Go-related Xmas reading, this could be ideal.
### 💻 Jobs

1. [Are You a Go Developer? Stream Has a Job for YouApply now](https://golangweekly.com/link/81293/web)

     — Stream is looking for a talented Go developer with a passion for building scalable infrastructure. Apply now.
1. [Sr. Software Engineer at CrowdStrike (Remote)](https://golangweekly.com/link/81294/web)

     — CrowdStrike is the leader in cloud-delivered endpoint protection, which helps protect our customers from cybersecurity attacks.
1. [Find a Job Through Vettery](https://golangweekly.com/link/81295/web)

     — Make a profile, name your salary, and connect with hiring managers from top employers. Vettery is completely free for job seekers.
### 📘 Articles & Tutorials

1. [Dependency Injection with WireWire](https://golangweekly.com/link/81296/web)

     — Wire is a compile-time DI solution built by Google themselves that uses code generation and reflection to produce injection code.
1. [The Go Runtime Scheduler's Clever Way of Dealing with System Calls](https://golangweekly.com/link/81299/web)

     — When scheduling goroutines onto processors, Go has an optimistic and an pessimistic approch. In other words, is your processor half full or half empty?
1. [Processing Parquet Files in Go](https://golangweekly.com/link/81300/web)

     — Parquet is a binary format that stores data in a columnar fashion and is commonly associated with Hadoop.
1. [Go Advanced Concurrency Patterns: Channels](https://golangweekly.com/link/81301/web)

     — To prove the power of using select with channels, Rob rewrites the sync package only using those constructs. Part 3 in a series.
1. [Discussing Concurrency, Parallelism, and Async Design](https://golangweekly.com/link/81302/web)

     — If you’ve ever felt intimidated by Go features like goroutines and channels or just concurrency in general, this will be 50 minutes well spent.
1. [Parsing CSV Files with Go](https://golangweekly.com/link/81303/web)

1. [Going Serverless with OpenFaaS and Go - Building Optimized Templates](https://golangweekly.com/link/81304/web)

1. [Parsing 18 Billion Lines JSON with Go](https://golangweekly.com/link/81305/web)

     — While this isn’t all about Go, the problem is interesting and using Go at that scale is impressive.
### 🛠 Code & Tools

1. [Slog: Minimal Structured Logging Library](https://golangweekly.com/link/81306/web)

     — Features first class context.Context and testing.TB support, both nice human readable text and machine readable JSON output, a minimal APIU surface, and more.
1. [go-locale: Cross Platform Locale Detection](https://golangweekly.com/link/81367/web)

     — Checks a variety of things from environment variables on Linux to using the Win32 API and more.
1. [golicense: Analyze OSS Dependencies and Licenses from Go Binaries](https://golangweekly.com/link/81307/web)

     — If you need to check licensing compliance, this tool could be invaluable.
1. [MongoDB Go Driver 1.2: The Official MongoDB Driver](https://golangweekly.com/link/81312/web)

     — The big feature here is support for MongoDB’s new client-side field level encryption (as supported in MongoDB 4.2+)
1. [ghw: A Go Hardware Discovery/Inspection Library](https://golangweekly.com/link/81310/web)

     — Find out things about the memory, CPU, storage, network support, and similar things about the host computer. Aimed at Linux, with partial support for macOS.
1. [go-circuitbreaker: A Context Aware Circuit Breaker Library](https://golangweekly.com/link/81311/web)

1. [DockerSlim: A Tool to Slim Down Docker Containers](https://golangweekly.com/link/81309/web)

     — A tool that aims to ‘throw away what you don’t need’ both reducing the attack surface of your containers and their size too.

### [ << Prev ](golangweekly-291.md) ------------- [ Next >> ](golangweekly-293.md)