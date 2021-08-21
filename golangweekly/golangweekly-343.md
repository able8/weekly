## [👋 Welcome to the final Golang Weekly of 2020! This was meant to be a 2020 roundup issue but there have been a few bits of news to cover so we're leading with those before taking a look back.](https://golangweekly.com/issues/343)

1. [Go on ARM.. and Beyond!mentioneda beta available here](https://golangweekly.com/link/100580/web)

     — ARM-based systems have been growing in popularity over the past few years and Apple’s M1 chip has suddenly upped the importance of supporting the architecture in more ways (Docker mentioned a lack of stable-version Go support for the M1 as one of its engineering challenges in delivering Docker for M1-based Macs).While Go has supported both 32-bit and 64-bit ARM for years, M1 specific support is due in Go 1.16 early next year (there’s a beta available here if you want to try it out) and this post outlines Go's growing support for ARM and other architectures.
1. [Redirecting godoc.org Requests to pkg.go.devpkg.go.dev](https://golangweekly.com/link/100585/web)

     — The Go team are very keen for pkg.go.dev to become the place for finding Go packages and documentation and so godoc.org is being gradually phased out in favor of it.
1. [Ebiten in 2020](https://golangweekly.com/link/100587/web)

     — Ebiten is a popular open source game library for building 2D games in Go. I’m always happy to promote it as it’s both a neat bit of work and seems surrounded by such happiness and activity.
1. [An Interview with Go's Rob Pike](https://golangweekly.com/link/100588/web)

     — Go’s co-creator answered some big picture questions about Go’s status, history, and future.
1. [Go Standard Library Benchmarks: Intel vs Apple's M1](https://golangweekly.com/link/100589/web)

     — We've spoken about Go on M1 (above) but you were all very keen to see some initial benchmarks.. :-) As with all benchmarks, though, maintain a critical eye here, especially as this is against an older i5 CPU.
1. [Insights on Rust vs Go](https://golangweekly.com/link/100590/web)

     — John is a huge Go fan and has done a lot of work in the Go space, but he thinks both Rust and Go are awesome and takes a careful look at where Go and Rust each independently make the most sense. We wouldn't usually strongly feature a 'versus' type post, but several people from the Go and Rust spaces reviewed this piece for quality and fairness.
1. [Inlined Defers in Go](https://golangweekly.com/link/100592/web)

     — An optimization in 1.14 and later for simple use cases with defer that removes most of the performance hit for deferred functions, complete with the how and the why.
1. [Different Approaches to HTTP Routing in Go](https://golangweekly.com/link/100593/web)

     — “There are many ways to do HTTP path routing in Go – for better or worse.” And Ben took a nice tour through a variety of options here.
1. [I Want Off Mr. Golang's Wild Ride](https://golangweekly.com/link/100594/web)

     — Well, not everyone loves Go. This is a very detailed rant(?) that goes beyond generics and takes issue with the idea that Go is simple. Worth a read, as things like this can help us flesh out our own opinions and figure out just why we do love Go.
1. [Giu: A Cross Platform Rapid GUI Framework Based on Dear ImGuiDear ImGui](https://golangweekly.com/link/100595/web)

     — Another way to create simple GUI apps from Go. Dear ImGui is an interesting GUI library (for C++) aimed at creating quite idiosyncratic UIs aimed at power users rather than standard end-user UIs. I like the minimalist vibe.
1. [Errors: Go Errors, But With Network Portability](https://golangweekly.com/link/100597/web)

     — A drop-in replacement for the standard errors package that adds things like network portability (across versions!) and stripping personally identifiable information out of errors. Check out the feature table to see how it compares to other solutions.
1. [RobotGo: Native Cross-Platform GUI Automation](https://golangweekly.com/link/100598/web)

     — We’ve linked this a few times over the years but this year it had a huge update. Control the pointer, keyboard, read the screen, etc.. you could use this to automate many computer-based jobs with enough creativity 😁
1. [Fyne: Cross Platform GUI Toolkit and API](https://golangweekly.com/link/100600/web)

     — We first linked this when it was for desktop only but it supports Android and iOS now too. It claims to be inspired by Material Design but looks more generally Linux-y to me. A clean look either way though.
1. [Bubble Tea: A Powerful Elm-Inspired TUI FrameworkElm architectureBubbles](https://golangweekly.com/link/100601/web)

     — Based on the Elm architecture, this is aimed at interactive terminal applications. There’s also Bubbles which provides a few components (spinner, text input, paginator, viewport) to use in Bubble Tea apps.
### 💻 Jobs

1. [Join the AWS Lambda Team (Seattle, Washington)](https://golangweekly.com/link/100604/web)

     — We're looking for gophers who enjoy teamwork, global scale services, and building the next generation of AWS Lambda. #serverless
1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/100605/web)

     — Join the most energizing community for developers and work on projects for Riot Games, FOX, Sony, Coinbase, and more.
1. [Find a Job Through Vettery](https://golangweekly.com/link/100606/web)

     — Create a profile on Vettery to connect with hiring managers at startups and Fortune 500 companies. It's free for job-seekers.

### [ << Prev ](golangweekly-342.md) ------------- [ Next >> ](golangweekly-344.md)