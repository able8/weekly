## [#​361 — May 7, 2021](https://golangweekly.com/issues/361)

1. [The Art of Solving Problems with Monte Carlo SimulationsMonte Carlo simulations](https://golangweekly.com/link/107684/web)

     — This is pretty math-heavy, but it does show several examples of using Monte Carlo simulations to converge on a value that could prove useful in this data-driven age.
1. [Proposal: A New Package to Provide Generic Slice Functions](https://golangweekly.com/link/107685/web)

     — Off the back of the eventual arrival of generics comes a proposal for a package containing a lot of the functions you’d expect around comparison and enumeration.
1. ['How Litestream Eliminated My Database Server for $0.03/month'Litestream](https://golangweekly.com/link/107687/web)

     — Litestream replicates data from a local (to your app) SQLite database to an S3-compatible set of SQLite snapshots. In effect, this removes the need for backups and data migrations. Oh, and it runs in its own process.
### 📘 Tutorials and Stories

1. [Branchless Coding in Go](https://golangweekly.com/link/107692/web)

     — Matt had some bit-packing to do and wondered if he could pull it off without any conditional branching. He gave it a go and reports on the pros and cons of the approach.
1. [Between Go and Elixir](https://golangweekly.com/link/107693/web)

     — “Reason wanted me to make a choice, and I am so glad I didn’t. Because the more I kept delving into both Elixir and Go, the more I found out how complementary the two can be to one another.”
1. [Building a Bingo Game Backend with EncoreEncore](https://golangweekly.com/link/107695/web)

     — Encore does, as promised, provide a lot of ‘magic’ when it comes to getting a service developed and deployed quickly.
1. [Building Tamper-Proof Systems with ImmuDB and GoImmuDB](https://golangweekly.com/link/107697/web)

     — ImmuDB is an Go-powered database that boasts no data mutation APIs at all and a ‘tamper-evident’ history system. (16 minutes.)
1. [A Note on Worker Pools in Go](https://golangweekly.com/link/107699/web)

     — If you’re using goroutines to handle requests then you’ll likely hit memory limits at scale, so worker pools to the rescue. However, there are sharks in those waters, too.
1. [How to Model JSON Data in Go with CockroachDB](https://golangweekly.com/link/107700/web)

1. [Building a Simple Terminal Emulator in 100 Lines of Go](https://golangweekly.com/link/107701/web)

1. [Platform Senior Software Engineer [Remote]](https://golangweekly.com/link/107702/web)

     — Doximity is on a mission to transform the US healthcare system. Join us and build the platform our teams use to ship software.
   

1. [Senior Software Engineer (Go) - 100% Remote, UK/EU Only](https://golangweekly.com/link/107703/web)

     — Solving complex technical challenges using the latest, cloud-native technologies to help Banks and FinTechs move money faster.
   

### 🛠 Code & Tools

1. [Gomponents: View Components in Pure Go](https://golangweekly.com/link/107704/web)

     — Write view components in the Go you know and love and they get rendered to HTML. I’ve never been a fan of this approach myself but it’s a neat project all the same.
1. [goyek: Create Build Pipelines in Go](https://golangweekly.com/link/107705/web)

     — This package was heavily influenced by the testing package which is apparent when you run a pipeline.
1. [Twirp 8.0: Simple RPC Framework with Protobuf Service Definitions8.0.0](https://golangweekly.com/link/107707/web)

     — Define your service in a Protobuf file and this autogenerates Go code with a server interface and fully functional clients. Think gRPC, but without the custom HTTP server and transport implementations - it uses net/http instead. 8.0.0 brings Twirp up to Protobuf APIV2 standards.
1. [sqlc 1.8: Generate Type Safe Go(lang) from SQL1.8.0](https://golangweekly.com/link/107709/web)

     — You write SQL queries, run sqlc to generate code and interfaces for those queries, then write Go code that calls the aforementioned code. 1.8.0 adds support for Postgres 12 and 13 features by upgrading an underlying dependency.

### [ << Prev ](golangweekly-360.md) ------------- [ Next >> ](golangweekly-362.md)