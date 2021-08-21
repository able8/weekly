## [#328 — September 4, 2020](https://golangweekly.com/issues/328)

1. [GORM 2.0 Released: A 'Fantastic ORM' for GoGORMGitHub repo.](https://golangweekly.com/link/94636/web)

     — With version 2, GORM has been rewritten from scratch and boasts improved performance and modularity. It’s a huge upgrade (nested transactions, batch processing, dry runs, context support..) so updating existing GORM code may take a little work, but if you haven’t used it before, it’s well worth checking out. GitHub repo.
1. [An Update on the Proposal for Embedding Files in Go Binaries](https://golangweekly.com/link/94639/web)

     — While there are plenty of tools for embedding files in Go binaries already, the Go team has been discussing the potential of making it an official feature built into cmd/go itself. Here’s the latest discussion on the topic.
1. [Some Examples of Using Generics by the Go Team](https://golangweekly.com/link/94641/web)

     — go2go is a translation tool for experimenting with generics in Go and these examples follow the current generics design draft complete with square brackets.
1. [Even in Go, Concurrency Is Still Not Easy (With An Example)](https://golangweekly.com/link/94642/web)

     — Goroutines and channels are not all there is to concurrency and you can still make mistakes. Chris shows one such example to prove it’s still on the programmer to use the tools properly.
1. [Go 1.15.1 and 1.14.8 Releasedabout that issue heredownloads page.](https://golangweekly.com/link/94643/web)

     — Two releases that fix a recently reported security issue when you don’t explicitly set a Content-Type in an HTTP handler (when using net/http/cgi or net/http/fcgi). More about that issue here or hit the downloads page.
### 💻 Jobs

1. [Stream Provides APIs for Building Activity Feeds and Chatapply now](https://golangweekly.com/link/94646/web)

     — Stream is looking for a full time Backend Software Engineer to join our development team. If you are interested in becoming a part of what we do, apply now.
1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/94647/web)

     — Join the most energizing community for developers and work on projects for Riot Games, FOX, Sony, Coinbase, and more.
1. [Find a Job Through Vettery](https://golangweekly.com/link/94648/web)

     — Create a profile on Vettery to connect with hiring managers at startups and Fortune 500 companies. It's free for job-seekers.
### 📘 Tutorials

1. [Hugo Modules: Everything You Need to Know](https://golangweekly.com/link/94649/web)

     — Hugo, Go’s favorite static site generator, introduced a module system in 0.56.0. This allows you to mount any Hugo project and use some or all of its files.
1. [Introducing Clean Architecture by Refactoring a Go ProjectClean Architecture](https://golangweekly.com/link/94650/web)

     — Three Dots Labs has a series of posts building a real microservices project in Go. This post starts a refactoring toward Clean Architecture, an idea around defining clear boundaries within a system.
1. [An Introduction to the Go Client for Elasticsearchofficial Go client](https://golangweekly.com/link/94652/web)

     — If you’ve got a large amount of data you want to run full text search over, Elasticsearch has probably turned up on your radar. The official Go client now includes features such as retrying requests and discovering cluster nodes and this post digs into how it works.
1. [Discovering Alloc Size Classes in Go](https://golangweekly.com/link/94654/web)

     — Go 1.15 changed the way bytes are allocated for an object, rounding up to the nearest “size class.” Read on to find out what that means.
1. [Turn a Go REST API to GraphQL using Hasura Actions](https://golangweekly.com/link/94656/web)

     — How to convert an existing REST API to GraphQL by just defining types and configuring endpoints.
### 🛠 Code & Tools

1. [Echelon: Hierarchical Progress Bars in Terminal on Overdrive](https://golangweekly.com/link/94657/web)

     — A cross-platform library (yes, even on Windows) to organize progress bars in a hierarchical structure and that can be used from multiple goroutines.
1. [Apache Pulsar Go Client LibraryApache Pulsar](https://golangweekly.com/link/94658/web)

     — Apache Pulsar is a distributed messaging and streaming platform originally built at Yahoo, and it now has a pure-Go client library.
1. [gomplate: A Flexible CLI Tool for Template Rendering](https://golangweekly.com/link/94661/web)

     — A template renderer that supports a growing list of sources, such as: JSON, YAML, AWS EC2 metadata, BoltDB, Hashicorp Consul and Hashicorp Vault secrets. Worth looking at the examples here.
1. [gookit/color: Terminal Color Rendering Tool Library](https://golangweekly.com/link/94662/web)

     — No dependencies. Supports rich color output at 16-color, 256-color, or even true color (24-bit, RGB) levels.
1. [db: A Productive Data Access Layer for Different Data Sourcesa neat online tour](https://golangweekly.com/link/94663/web)

     — db is ORM-‘like’ and provides an agnostic way to work with numerous databases like Postgres, SQLite, MySQL, MongoDB, and SQL Server. There’s a neat online tour stepping through how it works too.
1. [pgzip 1.2.5: Parallel gzip Compression and Decompression](https://golangweekly.com/link/94665/web)

     — A drop in replacement for compress/gzip with added parallelism.
1. [Go Stripe: The Official Library for the Stripe APIv72](https://golangweekly.com/link/94666/web)

     — For when you’re using Stripe for your billing. v72 is out now – they certainly like their releases at Stripe :-)
1. [Zap 1.16: Fast, Structured, Leveled Logging in Go](https://golangweekly.com/link/94668/web)

1. [httpx 1.0: A Multi-Purpose HTTP Toolkit for Probing Serversretryablehttp library](https://golangweekly.com/link/94669/web)

     — A command line tool, written in Go, for running numerous HTTP ‘probers’ at the same time using the retryablehttp library under the hood.
### 🎲 Fun and Side Projects

1. [A Music Video Coded with GoHere’s the source code.](https://golangweekly.com/link/94671/web)

     — It’s all text-based but a neat experiment all the same. Here’s the source code.

### [ << Prev ](golangweekly-327.md) ------------- [ Next >> ](golangweekly-329.md)