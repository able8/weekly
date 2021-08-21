## [#​359 — April 23, 2021](https://golangweekly.com/issues/359)

1. [Go 1.17 (..Hopefully) to Provide Better Performance via Register-Based Calling Convention](https://golangweekly.com/link/106913/web)

     — The details (which are quite technical) are all in this GitHub issue, but the takeaway is in the headline. More performance is good and a 6.5% bump is suggested here. This is on AMD64 for now, but ARM64 is the next architecture they’re planning to work on.
1. [Maybe Go Executable Files Aren't Full of 'Non-Useful Bits'Why are my Go executable files so large?become a retraction](https://golangweekly.com/link/106914/web)

     — Last week we featured an update to 2019’s Why are my Go executable files so large? where the author argued that modern Go binaries continue to be full of ‘non-useful bits.’ This article has now become a retraction as Russ Cox hit Hacker News to clear up what was going on, noting that the piece was ‘full of misinformation.’
1. [Excelize 2.4.0: A Library for Reading and Writing Excel Filesadds a lot](https://golangweekly.com/link/106918/web)

     — Read and write XLSX files, set and read cell values, add charts. This latest release adds a lot with support for 152 formula functions, a new API to get rich text from cells, and lots of bug fixes.
1. [Go 1.17 Will Allow Converting a Slice to an Array Pointer](https://golangweekly.com/link/106920/web)

     — Well, sometimes. 1.17 adds this ability without using reflect or unsafe, but there are some cautions to note.
### 📘 Tutorials and Stories

1. [Writing Good Unit Tests; Don't Mock Database Connections](https://golangweekly.com/link/106921/web)

     — While most of this post is 101 stuff, the take on mocking external dependencies might generate some opinions.
1. [Seven Years of Open-Source Database Development: Lessons Learnedrqlite](https://golangweekly.com/link/106922/web)

     — The developer of rqlite, a Go-powered, distributed SQLite-backed database system, has some high-level reflections on what it’s like to work on such a project.
1. [Go Generics Beyond the Playground](https://golangweekly.com/link/106925/web)

     — Sindre attempts to rewrite a test matcher library using generics, diving into how they solved the problem pre-generics and what is possible now.
1. [Talking TCP and UDP with Adam WoodbeckNetwork Programming with Go](https://golangweekly.com/link/106926/web)

     — Along with the author of Network Programming with Go, the Go Time podcast heads into the thorny weeds of networking with a look at two protocols that the entire Internet relies upon.
1. [Concurrent API Patterns in Go](https://golangweekly.com/link/106928/web)

     — Five simple rules (using three steps) to follow for building concurrent, leak-free API code.
1. [Senior Software Engineer at Even (Anywhere)](https://golangweekly.com/link/106929/web)

     — Help end the paycheck-to-paycheck cycle. Build w/ Go, React Native, GraphQL, Postgres, Bazel. Remote encouraged.
   

1. [Senior Software Engineer (Go) - 100% Remote, UK/EU Only](https://golangweekly.com/link/106930/web)

     — Solving complex technical challenges using the latest, cloud-native technologies to help Banks and FinTechs move money faster.
   

1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/106931/web)

     — Join the most energizing community for developers and work on long-term projects for Riot Games, FOX, Sony, Coinbase, and more.
   

### 🛠 Code & Tools

1. [Vugu: Vue-Like Frontends in Pure Go, HTML, and CSStwo years agohere on Hacker News](https://golangweekly.com/link/106932/web)

     — We called this ‘bleeding edge’ two years ago and it’s still a pretty novel idea. Vugu targets WebAssembly, so it will work in the latest major browsers but is also still a bit gooey in the middle. More discussion here on Hacker News.
1. [Tunny: A Goroutine Pool Library](https://golangweekly.com/link/106935/web)

     — A library for spawning and managing a pool of goroutines such as for when you need to limit the concurrent processing of jobs.
1. [xsek: An (Almost) Compliant XPath 1.0 Library](https://golangweekly.com/link/106937/web)

     — It’s really only missing one bit (but it adds others) and it’s extensible so you might be able to make “almost” into “entirely.”
1. [Fiber 2.8.0: An Express.js Inspired Web Framework](https://golangweekly.com/link/106938/web)

     — If you know Express (from the Node world) than Fiber will look very familiar.
1. [sx: Fast, Modern, Easy-to-Use Network Scanner](https://golangweekly.com/link/106939/web)

     — Handles ARP, TCP, UDP, ICMP scans and more.
1. [Raft 1.3.0: A Go Implementation of the Raft Consensus Protocol](https://golangweekly.com/link/106940/web)

1. [gronx: Fast, Dependency-Free Cron Expression Parser](https://golangweekly.com/link/106941/web)


### [ << Prev ](golangweekly-358.md) ------------- [ Next >> ](golangweekly-360.md)