## [#300 — February 21, 2020](https://golangweekly.com/issues/300)

1. [pkg.go.dev To Be Open Sourcedmore on where pkg.go.dev is headed](https://golangweekly.com/link/84272/web)

     — The “central source of information about Go packages and modules” is set to have its code released in a few months/weeks, and there is some rejoicing. Here’s more on where pkg.go.dev is headed.
1. [Advanced Go Concurrency](https://golangweekly.com/link/84323/web)

     — Goes beyond the usual candidates (goroutines, using context and sync) and digs into x/sync and things like singleflight and errgroup. If you’ve not dug far into concurrency in Go, you should pick up a thing or two here.
1. [Go Tops List of In-Demand Software Skills](https://golangweekly.com/link/84277/web)

     — The IEEE have picked up on a report from Hired (the recruiting marketplace) which places Go at the top of the pile when it comes to ‘most in-demand coding languages’ worldwide, though Scala(!) just pips it in the Bay Area.
1. [ESR's Notes on the Go Translation of ReposurgeonGo could replace C](https://golangweekly.com/link/84274/web)

     — A few years ago, Eric S Raymond thought Go could replace C in most situations, but now he’s done a rather extensive write up of his experiences in porting a complex program from Python to Go. The result is faster and easier to read, but he ran into a variety of problems.
1. [What's Involved in Building a 'Roguelike' in GoNetHack](https://golangweekly.com/link/84279/web)

     — A “roguelike” is a type of fantastical game, a la NetHack or Diablo. Hunter wrote this enjoyable conversational piece on game design in Go just prior to being eaten by a dragon.
### 💻 Jobs

1. [Software Engineer / Founding Teammate (Remote OK)](https://golangweekly.com/link/84281/web)

     — A well funded, fast-moving startup that is changing the way large teams collaborate on documents. Think: GitHub for documents.
1. [Find a Dev Job Through Vettery](https://golangweekly.com/link/84282/web)

     — Vettery is completely free for job seekers. Make a profile, name your salary, and connect with hiring managers from top employers.
### 📘 Articles & Tutorials

1. [Moving Towards Domain Driven Design in Go](https://golangweekly.com/link/84278/web)

     — An abstract (meaning, not runnable) tutorial where the DDD code evolves from the tightly coupled version and why evolving to the domain approach is recommended.
1. [What's New In Go 1.14: Test Cleanup](https://golangweekly.com/link/84283/web)

     — 1.14 adds a Cleanup() method to the testing objects that is run at the end of each test. This is good for cleaning up a database, for example.
1. [Rendering a Bloom Effect in Go](https://golangweekly.com/link/84284/web)

     — Basically makes things look like they’re glowing.
1. [Modeling MongoDB Documents with Native Go Data Structures](https://golangweekly.com/link/84285/web)

     — How to use the MongoDB Go driver to add BSON annotations to your native Go data structures, providing a way to seamlessly interact with data in the app as well as the database.
1. [Running GoLand on a Raspberry Pi 4](https://golangweekly.com/link/84287/web)

     — Someone at JetBrains wanted to see if they could get GoLang (a popular commercial Go IDE) running on a Raspberry Pi 4.. turns out, you can, if you’re okay with a pretty lengthy startup time.
1. [Discussing Interfaces in Go](https://golangweekly.com/link/84288/web)

     — The latest episode of the Go Time podcast digs into interfaces at length, what they are, how they’re being used, plus tips for how you can use them too.
### 🛠 Code & Tools

1. [Spinner: 70+ Configurable Terminal Spinner/Progress Indicators](https://golangweekly.com/link/84289/web)

     — If you want to add a little pizazz to your CLI app when it’s processing or loading something, you can’t ask for more variety than this. Check out all of the GIF demos!
1. [Operator: Go Operators as Functions](https://golangweekly.com/link/84324/web)

     — Implements logical, arithmetic, bitwise and comparison operators as functions (like Python’s operator module). Includes unary, binary, and nary functions with overflow checked variants.
1. [Hugo 0.65.0: The Go-Based Static Site Builder](https://golangweekly.com/link/84290/web)

     — If you’re a fan of Hugo for building sites, this seems to be a significant release if you want more flexibility in how pages are managed.
1. [Nodebook: A Minimalist Multi Language REPL with a Web-Based UI](https://golangweekly.com/link/84291/web)

     — I think this started life as a Node specific REPL but it now supports Go, C, C++, Elixir, and more. It’s principally written in Go, though.
1. [go-memdb: An In-Memory Database Built on Immutable Radix Trees](https://golangweekly.com/link/84293/web)

1. [goxygen: Generate a Go and MongoDB-Backed React Project in Seconds](https://golangweekly.com/link/84294/web)

     — An opinionated full stack app generator that builds a skeleton backend that uses Go and MongoDB with React on the front end.
1. [Kivik: A Common Interface to CouchDB or CouchDB-like Databases for Go](https://golangweekly.com/link/84295/web)

     — A mature library that has just been updated to use modules.
1. [go-blurhash: A 'BlurHash' Implementation in Pure GoBlurHash](https://golangweekly.com/link/84296/web)

     — A BlurHash is an ultra compact representation of a placeholder for an image that, as a blur, vaguely represents the original image visually.
1. [REST: An HTTP Client Specifically for RESTful APIs](https://golangweekly.com/link/84298/web)

     — SendGrid uses this in their own client library.
1. [Twirp: A Simple RPC Framework with Protobuf Service Definitions](https://golangweekly.com/link/84299/web)

     — Think gRPC but without the HTTP server and transport stuff. This runs direct on net/http and can even run over HTTP 1.1.
1. [SeqKit 0.12: A Go Toolkit for FASTA/Q File Manipulation](https://golangweekly.com/link/84300/web)

     — Perhaps a handful of you will find this useful but it’s neat to see Go being used for working with biochemistry. FASTA/FASTQ are formats for storing “nucleotide and protein sequences”..

### [ << Prev ](golangweekly-299.md) ------------- [ Next >> ](golangweekly-301.md)