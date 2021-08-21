## [#278 — September 5, 2019](https://golangweekly.com/issues/278)

1. [Go 1.13 Releasedthe release notes.this slidedeck](https://golangweekly.com/link/69492/web)

     — Hurrah, a new version of Go! The big news for this release is the go commands will use Google’s module mirror and checksum database to grab modules (more on this below). TLS 1.3 support is also now enabled by default, along with numerous other things covered in the release notes. Alternatively, this slidedeck is also another neat way to tour the new features.
1. [Go's Official Module Mirror and Checksum Database Launched](https://golangweekly.com/link/69495/web)

     — The Go team have been working on a module mirror, index, and checksum database in order to improve performance, resiliency (as otherwise if dependencies disappear at their original sources, you’ll have trouble), and security when using third party modules. It’s used by default by Go 1.13 (above).
1. [Context Package Semantics In Goof Contexts](https://golangweekly.com/link/69497/web)

     — “Understanding the semantics of Contexts is critical if your goal is to write reliable software with integrity.” This is a great article with many practical examples.
1. [Two Go Programs, Three Different Profiling Techniques](https://golangweekly.com/link/69498/web)

     — A GopherCon 2019 ‘must watch’ if you’re interested in digging into the performance of your programs.
### 💻 Jobs

1. [Are You a Go Dev? Want to Join an Awesome Startup in Amsterdam?](https://golangweekly.com/link/69499/web)

     — Stream is hiring a talented Go developer to join our team in Amsterdam. Apply now.

1. [Find a Golang job through Vettery](https://golangweekly.com/link/69500/web)

     — Make a free profile, name your salary, and connect with hiring managers from top employers.
### 📘 Articles & Tutorials

1. [Instrumenting Go Applications](https://golangweekly.com/link/69501/web)

     — Covers all the practices that make up good observability: logging, tracing, error handling, and metrics.
1. [Discussing Go's Role in a Serverless World](https://golangweekly.com/link/69502/web)

     — The latest episode of the popular Go Time podcast covers serverless approaches to development, going from the basics to how Go is specifically well suited to the task.
1. [Breadth-First Search using the Standard Library](https://golangweekly.com/link/69505/web)

     — BFS is one of the two ways to search a queue and it is easiest enough to do in Go, provided you know about certain packages…
1. [Logging HTTP Requests in Go](https://golangweekly.com/link/69507/web)

     — A one-pager that covers HTTP logging: what to log, where to get it, how to format it, and when to toss it.
1. [RSA - Theory and Implementation](https://golangweekly.com/link/69509/web)

     — Everything you want to know about RSA encryption and decryption, punctuated with an example in Go.
1. [Using Go Modules at Scale](https://golangweekly.com/link/69511/web)

### 🛠 Code & Tools

1. [SCS: HTTP Session Management for Go](https://golangweekly.com/link/69513/web)

     — Loads and saves session data via middleware, works with various stores (Postgres, Redis, etc.) and easy to extend and customize.
1. [Dgraph 1.1: A Fast, Distributed Graph Database Built in Gov1.1.0](https://golangweekly.com/link/69515/web)

     — Aims to provide “Google production level scale and throughput” at low latency and supports a GraphQL-like query syntax. v1.1.0 is just out.
1. [Litter: A Pretty Printer Library for Go Data Structures to Aid in Debugging](https://golangweekly.com/link/69519/web)

1. [🇯🇵 Kagome: A Japanese Morphological Analyzer](https://golangweekly.com/link/69520/web)

     — I imagine the audience for this is very niche, but we like to feature stuff like that sometimes! So if you have some Japanese words you need to break apart..
1. [Curlie: A CLI Tool for HTTP Requests](https://golangweekly.com/link/69521/web)

     — If you like the interface of HTTPie but miss the features of curl, curlie is what you are searching for.
1. [xxhash: A Go Implementation of the 64-bit XxHash Algorithm, XXH64](https://golangweekly.com/link/69522/web)

     — “a high-quality hashing algorithm that is much faster than anything in the Go standard library.”
1. [bed: A Binary/Hex Editor Written in Go](https://golangweekly.com/link/69523/web)


### [ << Prev ](golangweekly-277.md) ------------- [ Next >> ](golangweekly-279.md)