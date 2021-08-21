## [#311 — May 8, 2020](https://golangweekly.com/issues/311)

1. [Caddy 2: The Go-Powered Web Server with Automatic TLSa new architecturethis Hacker News thread](https://golangweekly.com/link/87948/web)

     — After over a year of redesign, Caddy 2 has a new architecture to v1. If you want a new HTTPS server that ‘just works’, Caddy is well worth a look IMO. Its lead creator, Matt Holt, answered lots of questions on this Hacker News thread about the release.
1. [Rek: An Easy HTTP Client for GoRequests](https://golangweekly.com/link/87951/web)

     — The inspiration here is from Python’s very well known and highly esteemed Requests library.. so the Pythonistas among you might like this!
1. [Life Without Line Numbers](https://golangweekly.com/link/87954/web)

     — There’s a lot of buzz around reducing the size of Go binaries (1.15 does so by ~6%) and here’s another tactic: reduce the precision of the position information. The gain is 2-6%, depending on how far you take it.
1. [Discussing Black Hat GoBlack Hat Go](https://golangweekly.com/link/88062/web)

     — “Are you excited to learn about hacking and that?” Got an hour? Roberto Clapis, a security engineer at Google, and Tom Steele, a co-author of Black Hat Go, join the Go Time team to discuss security, penetration testing, and more.
### 💻 Jobs

1. [Enjoy Building Scalable Infrastructure in Go? Stream Is HiringApply now](https://golangweekly.com/link/87955/web)

     — Like coding in Go? We do too. Stream is hiring in Amsterdam. Apply now.
1. [Find a Job Through Vettery](https://golangweekly.com/link/87956/web)

     — Vettery specializes in tech roles and is completely free for job seekers. Create a profile to get started.
### 📚 Articles & Tutorials

1. [Mid-Stack Inlining in Go](https://golangweekly.com/link/87957/web)

     — Inlining a function can lead to serious performance gains, so why not do it for everything? Well, there are always trade-offs.
1. [Asynchronous Preemption in Go 1.14](https://golangweekly.com/link/87958/web)

     — How the new preemption implementation works, including the use of a lesser-known signal (SIGURG).
1. [Accelerating Aggregate MD5 Hashing Up to 800% with AVX512md5-simd](https://golangweekly.com/link/87963/web)

     — The culmination of this work is md5-simd, a Go library that performs such rapid MD5 hashing (when running concurrently). The use cases here are quite restricted but you may appreciate seeing how such things are implemented for any high end SIMD wrangling you need to do one day.
1. [A Beginner's Guide to gRPC in Gowritten version of the tutorial](https://golangweekly.com/link/87959/web)

     — There’s a written version of the tutorial if you dislike videos.
1. [Four Steps to Daemonize Your Go Programs](https://golangweekly.com/link/87961/web)

     — Daemons are programs that run as non-interactive background processes (e.g. background job processors, Web servers, database systems).
1. [Go as a Scripting Language?](https://golangweekly.com/link/87965/web)

     — There’s plenty of folks that use Go as a scripting language, but there are challenges around REPLs and shebang support. Some of these challenges are being addressed today.
### 🛠 Code & Tools

1. [UUID 3.3: A Pure Go Implementation of UUIDsRFC-4122](https://golangweekly.com/link/87966/web)

     — A pure Go implementation of Universally Unique Identifiers (UUID) as defined in RFC-4122 covering versions 1 through 5.
1. [Reed-Solomon: A Reed-Solomon Erasure Coding LibraryReed Solomon erasure coding](https://golangweekly.com/link/87968/web)

     — A Go port of a Java library built by Backblaze that does Reed Solomon erasure coding (a way to send or store data in a larger form that’s resilient to data loss). Boasts operation of over 1GB/sec per core.
1. [ko 0.5: Build and Deploy Go Apps on Kubernetes](https://golangweekly.com/link/87970/web)

     — ko’s objective is to “to make containers invisible infrastructure.” It’s been rapidly maturing in the past few months too.
1. [Tengo 2.2: A Fast Embeddable Script Language for Go](https://golangweekly.com/link/87971/web)

     — Quite a mature project now and worth a look if you need to add some dynamic scripting to your code.
1. [UniPDF 3.7: A Library for Creating and Processing PDF Files](https://golangweekly.com/link/87974/web)

     — Pure Go, which is neat, but note it’s dual licensed: AGPL for open source, commercial for closed source projects.
1. [Mockery: A Mock Code Generator for Go Interfaces](https://golangweekly.com/link/87976/web)

1. [Dynamo: An Expressive DynamoDB Library](https://golangweekly.com/link/87977/web)

### 🎲 Two Fun Side Projects

1. [gasm: An Experimental WASM Virtual Machine for Gophers](https://golangweekly.com/link/87973/web)

     — “I did this just for fun and for learning WASM specification.” Nonetheless, it works with basic examples.
1. [thdwb: A Homebrew Web Browser and Rendering Engine](https://golangweekly.com/link/87972/web)

     — Another experimental, fun learning project. You won’t be using it for your day to day browsing any time soon but projects like this keep the imagination fueled up.

### [ << Prev ](golangweekly-310.md) ------------- [ Next >> ](golangweekly-312.md)