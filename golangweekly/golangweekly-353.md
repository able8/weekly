## [#353 — March 12, 2021](https://golangweekly.com/issues/353)

1. [The Go Developer Survey 2020 Results](https://golangweekly.com/link/104554/web)

     — Almost 10,000 gophers took the recently ran Go developer survey and the data has finally been crunched and various insights shared with us here. Some of the stand out points for me:

1. [netaddr.IP: A New IP Address Type for Go](https://golangweekly.com/link/104555/web)

     — This might sound boring at first glance but it’s both well written and it’s by Brad Fitzpatrick (formerly on the Go team at Google). Why did Tailscale, his new employer, need a new IP address type in Go? Find out why and how they approached the problem here.
1. [A Proposal for a New Security Policy for Go](https://golangweekly.com/link/104557/web)

     — You may have noticed that we’re mentioning a lot of Go updates based around fixing security issues lately (there’s even one just below this item). The Go core team are considering introducing a ‘severity scale’ for issues (so you can judge if a security release warrants a urgent upgrade, say) rather than them being treated as binary yes/no matters.
1. [Go 1.16.1/2 and 1.15.9/10 Releasedthen followed](https://golangweekly.com/link/104558/web)

     — Go 1.61.1 and 1.15.9 were two small releases to address two security issues: a infinite loop bug in encoding/xml and a potential panic bug in archive/zip. 1.16.2 and 1.15.10 then followed almost immediately as more typical bugfix releases.
1. [How I Build Web Frontends in Go](https://golangweekly.com/link/104559/web)

     — A developer shares his approach to using Go in situations where an API is oriented around a dynamically rendered Web-based frontend.
### 💻 Jobs

1. [Enjoy Building Scalable Infrastructure in Go? Stream Is HiringApply now](https://golangweekly.com/link/104560/web)

     — Like coding in Go? We do too. Stream is hiring in Amsterdam. Apply now.
1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/104561/web)

     — Join the most energizing community for developers and work on projects for Riot Games, FOX, Sony, Coinbase, and more.
1. [Find Golang Engineering Jobs with Hired](https://golangweekly.com/link/104562/web)

     — Take 5 minutes to build your free profile & start getting interviews for your next job. Companies on Hired are actively hiring right now.
### 📘 Tutorials and Stories

1. [Building a High-Scale Chat Server on Google Cloud Run](https://golangweekly.com/link/104563/web)

     — This isn’t necessarily the right tool for the job, but it’s interesting to think about what GCP’s serverless container compute service can handle.
1. [Finding a Problem in rtnetlink (or How to Nerdsnipe Me)rtnetlink](https://golangweekly.com/link/104564/web)

     — rtnetlink is a Go package that provides low level access to Linux’s rtnetlink API (used for working with the kernel’s routing tables).
1. [Go Maps vs Switches](https://golangweekly.com/link/104567/web)

     — Benchmarking switch statements vs looking values up in a map with spelunking into the code to figure out why one is faster than the other.
1. [gRPC Long-Lived Streaming](https://golangweekly.com/link/104568/web)

     — Omri creates an example using a service with multiple subscribers, demonstrating how to use gRPC for server and client, subscribe multiple clients, and handle errors for both.
1. [The Tradeoffs of Version Behavior in go.mod Module Files](https://golangweekly.com/link/104569/web)

     — The current version of Go will add a go 1.xx directive to the go.mod file, which comes with some tradeoffs that may surprise you.
1. [How I Organize Structs in Go Projects](https://golangweekly.com/link/104571/web)

     — Some conventions one gopher uses to organize structs in larger projects.
### 🛠 Code & Tools

1. [Pixelizer 2.0: An Efficient Image Pixelizer Written in Go](https://golangweekly.com/link/104572/web)

     — A fun little project that takes an image in and then pixelizes it.
1. [TinyGo 0.17.0: The Go Compiler for Small Places](https://golangweekly.com/link/104573/web)

     — Three months of work go into this ‘large release’ with LLVM 11 and Go 1.16 support amongst the headlines. If you’ve got microcontrollers or WebAssembly to work with, this is for you.
1. [Authelia: A Single Sign-On Multi-Factor Portal for Web Apps](https://golangweekly.com/link/104574/web)

     — An open-source (and, yes, written in Go) authentication and authorization server providing 2FA and SSO for your other apps.
1. [Marcel: An Email 'MIME Artist' Library](https://golangweekly.com/link/104575/web)

     — Apparently using Go and the mime/multipart library to send a properly encapsulated multipart email with text and HTML parts along with attachments is tricker than you’d think (I’ve not done it myself) and this will help orchestrate things.
1. [fsql: Search Through Your Filesystem with SQL-esque Queries](https://golangweekly.com/link/104576/web)

     — A tool written in Go to let you do things like SELECT name, size FROM . WHERE NOT name RLIKE *.go

### [ << Prev ](golangweekly-352.md) ------------- [ Next >> ](golangweekly-354.md)