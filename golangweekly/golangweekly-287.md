## [#287 — November 8, 2019](https://golangweekly.com/issues/287)

1. [Introducing Pkger: Static File Embedding in GoStatika 20 minute screencastGitHub repo](https://golangweekly.com/link/79660/web)

     — There are already a few ways to embed static files into Go binaries (such as Packr or Statik) but when the creator of Packr says he’s ‘rethought’ the problem and has come up with a new solution.. I listen! There’s a 20 minute screencast too, if you want a truly complete tour. GitHub repo.
1. [Go Modules: v2 and Beyond](https://golangweekly.com/link/79664/web)

     — The last in a series of posts on the official Go blog about using modules, as introduced in Go 1.11 and made a default feature in Go 1.13. On this outing, the focus is on versions and how to create and publish major new versions of your modules.
1. [Staticcheck in ActionStaticcheck](https://golangweekly.com/link/79666/web)

     — An introduction to the static analysis tool Staticcheck for analyzing and then improving Go code.
1. [How Many Go Developers Are There?](https://golangweekly.com/link/79668/web)

     — Russ has an update to some research he did a couple of years ago, complete with a new estimate: “As of November 2019, my best estimate is between 1.15 and 1.96 million.” To the 2.5% of you who subscribe to this newsletter, thank you! 😄
1. [Finding a Memory Leak in a Go App with cgo Bindings](https://golangweekly.com/link/79669/web)

     — Kir’s team had to find and fix a memory leak in a C extension that was being used by their Go app. Here’s how they did it.
1. [TiDB in the Browser: Running a Go-Based Database on WebAssembly](https://golangweekly.com/link/79670/web)

     — An interesting demo. The author wonders if a database like TiDB written in Go can run in a web browser, what about other complex Go apps? Go’s WebAssembly future is looking quite positive.
### 💻 Jobs

1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/79671/web)

     — Work with the world's leading brands, from anywhere. Travel the world while being part of the most energizing community of developers.
1. [Enjoy Building Scalable Infrastructure in Go? Stream Is Hiring](https://golangweekly.com/link/79672/web)

     — Like coding in Go? We do too. Stream is hiring in Amsterdam. Apply now.

1. [Find a Job Through Vettery](https://golangweekly.com/link/79673/web)

     — Vettery specializes in tech roles and is completely free for job seekers. Create a profile to get started.
### 📘 Articles & Tutorials

1. [Separate Your Go Tests with Build Tags](https://golangweekly.com/link/79674/web)

     — ‘Build tags’ are directives that can be placed at the top of Go source files and used as a way to logically separate tests.
1. [Building Sustainable Microservices, Our Opinions and Advice](https://golangweekly.com/link/79675/web)

     — A ‘semi opinionated guide’ that answers some questions that Echo is often asked internally by new developers who are new to the company, (micro)services, Go or distributed systems in general.
1. [Instrumenting Go Code En Masse via AST Manipulation](https://golangweekly.com/link/79676/web)

     — With a large number of API handlers to instrument, Mattermost took a novel approach of finding and editing them all.
1. [Unit Testing exec.Command](https://golangweekly.com/link/79678/web)

     — exec.Command executed a named program with arguments you supply, but how can you test functions that use it?
1. [Managing CPU Load in Golang](https://golangweekly.com/link/79679/web)

     — Patterns for handling requests, including worker pools, buffered channels, and offloading to third parties.
1. [Analyzing Go Executables with JEBJEB](https://golangweekly.com/link/79680/web)

     — JEB is a commercial reverse engineering tool and it can be used to analyse Go binaries in various ways, as seen here (where the focus is on analyzing malware written in Go).
1. [Realizing That Go Constants Are Always Materialized Into Values](https://golangweekly.com/link/79682/web)

1. [Instrumenting Your Go Distributed Application for Tracing with Jaeger](https://golangweekly.com/link/79683/web)

     — If you are building an application from microservices, tracing is paramount to production success.
### 🛠 Code & Tools

1. [goque: Persistent Stacks and Queues for Go Backed by LevelDB](https://golangweekly.com/link/79684/web)

     — Now also supports encoding/decoding objects as JSON which can make storing complex objects easier.
1. [lungoDB: A MongoDB-Compatible Embeddable Database and Toolkit](https://golangweekly.com/link/79685/web)

     — Imagine something like SQLite but MongoDB flavored. There are numerous possible use cases, well outlined in the README here.
1. [Go DNS: A Library to Work with DNS from Go](https://golangweekly.com/link/79686/web)

     — It’s technical and you need to be prepared to dig around a bit but this library supports an incredible amount of DNS related RFCs and features.
1. [Stripe CLI: A Command Line Development Environment for Stripe UsersGo powered!](https://golangweekly.com/link/79688/web)

     — Stripe has become somewhat ubiquitous in the payment processing space and their focus on developers is pretty neat, not least in this new (Go powered!) tool for building and testing integrations.
1. [sqlhooks: Attach Hooks to Any database/sql Driver](https://golangweekly.com/link/79690/web)

     — An unobtrusive way to instrument SQL queries. Has just had its first update in almost two years with go mod support added.
1. [Viper: A Complete Configuration Solution for Go Apps](https://golangweekly.com/link/79691/web)

     — A mature project used by projects like Hugo, Nanobox and Docker Notary for their configuration needs.

### [ << Prev ](golangweekly-286.md) ------------- [ Next >> ](golangweekly-288.md)