## [#324 — August 7, 2020](https://golangweekly.com/issues/324)

1. [Boids in WebAssembly Using Go](https://golangweekly.com/link/93215/web)

     — A boid (bird-oid) is a “simulated bird-like object” and simulating a flock of them in Go and WebAseembly makes for an interesting demo. (The technique is of particular interest to me, rather than the output, though.)
1. [Go 1.14.7 and Go 1.13.15 Releasedthe issue in question](https://golangweekly.com/link/93216/web)

     — These releases address a recently reported security issue around ReadUvarint and ReadVarint (which is explained in more depth in the item below). Here’s the issue in question if you haven’t got a Google login.
1. [Comparing Go and Rust for Writing a CLI Tool](https://golangweekly.com/link/93218/web)

     — A developer, unfamiliar with both Go and Rust, decided to write an app in both and compared his experiences. It’s not deep but it’s quite balanced and he found good reasons to use both languages.
1. [Taking Advantage of a Bug in Go's Standard Library](https://golangweekly.com/link/93219/web)

     — Go 1.14.7 and 1.13.15 were just released (above) to fix this very bug which could be used as part of a DoS attack.
1. [Go 1.15 Release Candidate 2 Releaseddraft release notes.](https://golangweekly.com/link/93220/web)

     — Here are the draft release notes. Fingers crossed we get to announce the final release next week :-)
### 💻 Jobs

1. [Software Engineer - Want to Build a Platform Ecosystem in Go?Apply now.](https://golangweekly.com/link/93142/web)

     — Skool is hiring its 2nd backend engineer in Los Angeles, CA. Go, PostgreSQL, Redis, Elasticsearch, Docker. Apply now.
1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/93143/web)

     — Join the most energizing community for developers and work on projects for Riot Games, FOX, Sony, Coinbase, and more.
1. [One Application, Hundreds of Hiring Managers](https://golangweekly.com/link/93144/web)

     — Use Vettery to connect with hiring managers at startups and Fortune 500 companies. It's free for job-seekers.
### 📘 Tutorials

1. [An Introduction to Go's Escape Analysis](https://golangweekly.com/link/93222/web)

     — Escape analysis is the process of determining the scope of pointers, which in Go means whether variables should be allocated on the stack or the heap.
1. [Generators in Go](https://golangweekly.com/link/93223/web)

     — Generators here are for creating iterators, not generated code, that you can range over easily.
1. [Building Desktop Apps with Go: Webview vs. Lorca vs. Electron](https://golangweekly.com/link/93224/web)

     — A comparison of the three frameworks that (spoiler!) ends up with an odd recommended approach. Other options are available, however!
1. [Discussing Go with Google's Ian Lance Taylor](https://golangweekly.com/link/93225/web)

     — Yes, we’re linking to a C++ podcast! But they had Google’s Ian Lance Taylor (who works on GCC and Go’s frontend for it) on to talk about Go (and C++, a little).
1. [How to Notarize and Sign Go Binaries for MacOS with GitHub Actions](https://golangweekly.com/link/93226/web)

     — For jumping through all the hoops MacOS Catalina presents regarding trusted binaries.
1. [How Subqueries Were Implemented in a Go-Powered SQL EngineDoltgo-mysql-server](https://golangweekly.com/link/93227/web)

     — There’s a few moving parts here, but basically Dolt is like Git for data and the underlying SQL engine is go-mysql-server, a MySQL wire protocol compatible SQL server and engine.
1. [Tiny Go: Small Is Going Big](https://golangweekly.com/link/93230/web)

     — Ron Evans, the creator of TinyGo (a Go compiler for ‘small’ environments such as microcontrollers), gave a talk at QCon London 2020 (just before lockdown!). You can watch it here or read a transcript. Worth watching if you want to see how Go can be used in 'small' places :-)
### 🛠 Code & Tools

1. [sqlc v1.5.0 Released](https://golangweekly.com/link/93231/web)

     — A tool for creating fully type safe idiomatic Go code from SQL. Write SQL queries, run sqlc, then write application code that calls the methods sqlc generated.
1. [aws-lambda-go: Libraries, Samples, and Tools for Go and AWS Lambda](https://golangweekly.com/link/93232/web)

     — Libraries, samples and tools to help Go developers develop AWS Lambda functions.
1. [defaults: Initialize Structs with Default Values](https://golangweekly.com/link/93233/web)

     — Uses field tags to assign the default value. This seems very useful.
1. [html-to-markdown: Convert HTML to Markdown](https://golangweekly.com/link/93234/web)

     — Uses an HTML parser vs regex everywhere.
1. [Gearbox: A Web Framework with a Focus on High Performance](https://golangweekly.com/link/93235/web)

     — Written on top of fasthttp.
1. [go-mysql-server: An Extensible MySQL Server Implementation in Go](https://golangweekly.com/link/93229/web)

     — We linked to an article about this in the tutorials section (above) but basically this is an SQL engine and server that clones some of MySQL, so you get MySQL SQL syntax and wire protocol support but written in Go. It is not a complete database system in and of itself though (read the docs).

### [ << Prev ](golangweekly-323.md) ------------- [ Next >> ](golangweekly-325.md)