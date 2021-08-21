## [#295 — January 17, 2020](https://golangweekly.com/issues/295)

1. [Porting the Go Pigo Face Detection Library to WebAssemblyPigo](https://golangweekly.com/link/82495/web)

     — Pigo is a neat face detection library we’ve linked to a few times now, but wouldn’t it be great to use directly in the browser too? Here’s what was involved.
1. [Plumbing At Scale: Event Sourcing and Stream Processing Pipelines at Grab](https://golangweekly.com/link/82497/web)

     — Grab is a Singapore-based food ordering app and this post digs into their  journey building and deploying an event sourcing platform with Go for servicing 300 billion events each week.
1. [How We Optimized Our DNS Server using Go Tools](https://golangweekly.com/link/82500/web)

     — It’s written in an idiosyncratic way but there’s a lot to enjoy here.
1. [Stream Is Deprecating Virtual Go (vg): It’s Time to Move to Go Modulesvg](https://golangweekly.com/link/82501/web)

     — vg was a project providing ‘workplace based’ development for Go to avoid problems with dependencies clashing between multiple projects.. but Go modules now solve that problem.
### 💻 Jobs

1. [Sr. Software Engineer at CrowdStrike (Remote)](https://golangweekly.com/link/82503/web)

     — CrowdStrike is the leader in cloud-delivered endpoint protection, which helps protect our customers from cybersecurity attacks.
1. [Enjoy Building Scalable Infrastructure in Go? Stream is HiringApply now](https://golangweekly.com/link/82504/web)

     — Like coding in Go? We do too. Stream is hiring in Amsterdam. Apply now.
1. [Find a Job Through Vettery](https://golangweekly.com/link/82505/web)

     — Vettery is completely free for job seekers. Make a profile, name your salary, and connect with hiring managers from top employers.
### 📘 Articles & Tutorials

1. [Exposing Interfaces](https://golangweekly.com/link/82506/web)

     — Writing good interfaces can be tricky, but there are some guidelines you can follow which this article outlines with examples.
1. [How GitLab Scaled Git Access with Go](https://golangweekly.com/link/82507/web)

     — A 25 minute talk on the story behind how GitLab switched from a Ruby on Rails app to using feature flags, protocol buffers, gRPC and Go.
1. [PubSub Using Channels in Go](https://golangweekly.com/link/82508/web)

     — As with most things, there are multiple possible approaches here. Eli covers several and chooses his favorite.
1. [Go Things I Love: Channels and Goroutines](https://golangweekly.com/link/82509/web)

     — A quick look at a few concurrency patterns.
1. [Go at Heroku](https://golangweekly.com/link/82510/web)

     — Take a peek into how Heroku uses Go for its services, and how they teach through style guides and pairing sessions.
1. [The 'await/async' Concurrency Pattern in Go](https://golangweekly.com/link/82511/web)

     — Another look at concurrency from the perspective of replicating the await/async type approach you might have used in something like Node.js.
1. [How Go's net.DialContext() Stops Things When The Context Is Cancelled](https://golangweekly.com/link/82512/web)

### 🛠 Code & Tools

1. [BLAKE3: A Pure Go Implementation of the BLAKE3 Cryptographic Hash Functionoriginally implemented in Rusthere are some ideas.](https://golangweekly.com/link/82513/web)

     — BLAKE3 is a Merkle-tree based cryptographic hash function (originally implemented in Rust) that’s fast, secure, and highly parallelizable. And in case you’re wondering why you’d want a faster hash, here are some ideas.
1. [sqlmw: Wrap a database/sql Driver with Middleware](https://golangweekly.com/link/82516/web)

     — Provides an abstraction similar to HTTP middleware or gRPC interceptors but for database/sql so you can do logging and tracing more easily.
1. [Avo: Better x86 Assembly Generation from Goadding two numbers example](https://golangweekly.com/link/82517/web)

     — Avo makes assembly easier to write by using Go control structures and virtual registers, with values handled for you. The basic adding two numbers example shows off how it works in a simple way.
1. [GoJSONQ: A Package to Query JSON Data](https://golangweekly.com/link/82519/web)

     — Provide a fluent chainable API to query JSON documents.
1. [asciigraph: Make Lightweight ASCII Line Graphs](https://golangweekly.com/link/82520/web)

     — We’ve linked this a few times over the years, but I still love the output from this package and it continues to get updates from time to time.
1. [Aini: Go Library for Parsing Ansible Inventory Files](https://golangweekly.com/link/82521/web)

1. [ok-zoomer: Turn Portraits Into Zooming Animated GIFs](https://golangweekly.com/link/82523/web)

     — Finds a face in an image using Pigo then zooms in and out of that face and turns the result into an animated GIF. Got it? 😂
1. [zero-width: Zero-Width Character Detection and Removalall sorts of headaches.](https://golangweekly.com/link/82524/web)

     — Things like zero width spaces in strings can cause all sorts of headaches.

### [ << Prev ](golangweekly-294.md) ------------- [ Next >> ](golangweekly-296.md)