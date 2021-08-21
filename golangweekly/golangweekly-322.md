## [#322 — July 24, 2020](https://golangweekly.com/issues/322)

1. [A Draft Design for Go Command Support for Embedded Static Assetsa proposalPackrrecorded this video demo](https://golangweekly.com/link/92418/web)

     — Last December we linked to a proposal for bringing the embedding of static assets/files in Go binaries to the main Go toolset (versus using existing tools like Packr) – Brad Fitzpatrick is now back with a draft design for us to chew over (which includes links to 15 alternative implementations, if you’re curious!). Russ Cox has also recorded this video demo of how the draft design works in practice.
1. [A Chat with Brian Kernighan of Unix, C, AWK, and Go FameBrian Kernighan58:45](https://golangweekly.com/link/92422/web)

     — Lex Fridman, a well known AI researcher, had an extensive chat with Brian Kernighan, the co-author of The Go Programming Language (book). Brian touches on lots of neat stuff that you’ll enjoy, but the Go specific section is between 58:45 and 1:01:55.
1. [A Draft Design for New File System Interfaces for Govideo presentationAfero](https://golangweekly.com/link/92426/web)

     — You don’t get just one new draft design this week, you get two. This one similarly has a video presentation and introduces io/fs, a new package that defines an interface for read-only file trees. If this idea is ringing a bell, Afero, which we linked last week, covers similar ground.
1. [The go2go Playground Now Supports New Style Genericsthe next step for generics](https://golangweekly.com/link/92429/web)

     — Want to play with the next step for generics without lots of messing around? Enter this special build of the online Go Playground.
1. [A Draft Design for First Class Fuzzing](https://golangweekly.com/link/92431/web)

     — What’s better than two design drafts? Three. Fuzzing is a type of testing where inputs are manipulated continually in a search for issues. It is automated and mutatable and an excellent way to harden code.
### 💻 Jobs

1. [Sr. Software Engineer at CrowdStrike (Remote)](https://golangweekly.com/link/92432/web)

     — CrowdStrike is the leader in cloud-delivered endpoint protection, which helps protect our customers from cybersecurity attacks.
1. [Software Engineer - Want to Build a Platform Ecosystem in Go?Apply now](https://golangweekly.com/link/92433/web)

     — Skool is hiring its 2nd backend engineer in Los Angeles, CA. Go, PostgreSQL, Redis, Elasticsearch, Docker. Apply now.
1. [One Application, Hundreds of Hiring Managers](https://golangweekly.com/link/92435/web)

     — Use Vettery to connect with hiring managers at startups and Fortune 500 companies. It's free for job-seekers.
### 🛠 Code & Tools

1. [GoLand 2020.2 Reaches Beta](https://golangweekly.com/link/92436/web)

     — The popular (though commercial) IDE takes another step forward, though note this beta is the last release you can try out without a subscription. Improved Go module support is included here, along with initial, experimental support for generics.
1. [todocheck: A Static Code Analyzer for Annotated TODO Comments](https://golangweekly.com/link/92438/web)

     — I don’t know about you, but I’m a fan of dumping ‘TODO’ comments in my code but this tool will actually let you do something about them (in several languages, including Go).
1. [Gabs: Parse, Create and Edit 'Unknown' or Dynamic JSON in Go](https://golangweekly.com/link/92439/web)

     — A wrapper for navigating hierarchies of map[string]interface{} objects provided by encoding/json.
1. [Stats: A Comprehensive Go Statistics Library Package](https://golangweekly.com/link/92440/web)

     — No dependencies, well tested, work with concepts like averages, sums, percentiles, standard deviation, etc.
1. [Goldmark 1.2.0: A Markdown Parser Written in GoCommonMark](https://golangweekly.com/link/92441/web)

     — Boasts that it’s easy to extend and CommonMark compliant (we’ll debate the role of standards in Markdown another day 😏)
1. [Evergreen: A Distributed Continuous Integration System from MongoDB](https://golangweekly.com/link/92444/web)

     — Built by MongoDB to test MongoDB, but it’s written in Go and dynamically allocates hosts to run tasks in parallel across many machines.
1. [fluff: Fast Web Fuzzer Written in Go](https://golangweekly.com/link/92445/web)

     — This is for testing systems/apps of your own, of course, but will let you quickly fuzz POST data, URL parameters, paths, and more.

### [ << Prev ](golangweekly-321.md) ------------- [ Next >> ](golangweekly-323.md)