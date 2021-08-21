## [#312 — May 15, 2020](https://golangweekly.com/issues/312)

1. [What's Coming in Go 1.15Go 1.15 slidedeck](https://golangweekly.com/link/88449/web)

     — We linked to a Go 1.15 slidedeck a couple of weeks ago, but I know most of you prefer articles, so here you go! You know about smaller binaries, but did you know there’s a new linker?
1. ['Ensmallening' Go Binaries by Prohibiting Comparisons](https://golangweekly.com/link/88451/web)

     — Most people don’t need to squeeze Go binary sizes by this mild extent, but it’s an interesting spelunk nonetheless into how Go’s compilation process feeds a lot of sugar to your app, causing it to be a bit fatter and how you can put it on a (very modest) diet.
1. [go-elasticsearch: The Official Go Client for ElasticsearchElasticsearchv7.7.0example here.](https://golangweekly.com/link/88453/web)

     — Elasticsearch is a verrry popular search engine/document database used to implement search features in apps, and with v7.7.0 the Go client has added support for efficient, parallel ‘bulk indexing’ with very high ingest rates (example here.)
1. [immudb: A Lightweight, High-Speed Immutable Database](https://golangweekly.com/link/88457/web)

     — Data can’t be changed in this database so it’s well suited for storing every update to other databases for auditing purposes, perhaps, or maybe log streams or public certificates. Written in Go.
1. [Go 1.14.3 and Go 1.13.11 Releaseda handful](https://golangweekly.com/link/88458/web)

     — Minor point releases fixing a handful of minor bugs.
### 💻 Jobs

1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/88460/web)

     — Join X-Team and work on projects for companies like Riot Games, FOX, Coinbase, and more. Work from anywhere.
1. [Find a Job Through Vettery](https://golangweekly.com/link/88461/web)

     — Vettery specializes in tech roles and is completely free for job seekers. Create a profile to get started.
### 📚 Articles & Tutorials

1. [Faking stdin and stdout in Go](https://golangweekly.com/link/88462/web)

     — Redirecting standard input and output is common in testing situations but has other use cases as well.
1. [A Beginner Friendly Introduction to PrometheusPrometheus](https://golangweekly.com/link/88463/web)

     — Prometheus is a popular system monitoring and alerting system that’s built in Go and which you can therefore extend and enhance with Go too.
1. [How I Write My Unit Tests in Go Quickly](https://golangweekly.com/link/88465/web)

     — This covers everything from designing for easy testing to mocks to helpful libraries.
1. [Building Uber’s Go Monorepo with Bazel](https://golangweekly.com/link/88467/web)

     — Uber’s monorepo is the largest that’s using Bazel, leading to many challenges leading to many opportunities for Uber to work with Bazel and improve the platform.
1. [Getting My Head Around What Things Aren't Comparable in Go](https://golangweekly.com/link/88468/web)

     — Builds upon Dave Cheney’s featured article above.
### 🛠 Code & Tools

1. [Termdash: A Go Terminal Based Dashboard System](https://golangweekly.com/link/88469/web)

     — A terminal-based dashboard before that supports window resizing, customizable layouts, dynamic layout changes at runtime, and more.
1. [goaction: Write GitHub Actions in Go](https://golangweekly.com/link/88470/web)

     — “The idea is: write a standard Go script, one that works with go run, and use it as a GitHub action.”
1. [Shotizam: Analyze The Size of Go Binaries](https://golangweekly.com/link/88472/web)

     — Shotizam gives you a SQLite prompt with the results so you can use SQL to query them.
1. [Excelize 2.2: A Library for Reading and Writing Excel Files2.2](https://golangweekly.com/link/88473/web)

     — Read and write XLSX files, set and read cell values, add charts. 2.2 is a key step forward with easier access to numerous features.
1. [jsonparser 1.0: An Alternative JSON Parser for Go](https://golangweekly.com/link/88476/web)

     — Schema-less so it’ll parse whatever you throw at it. Claims to be 10 times faster than encoding/json.
1. [Squirrel 1.4.0: Fluent SQL Generation for Go](https://golangweekly.com/link/88477/web)

     — Think about chainable methods to build a SQL query. For example: sq.Select("*").From("users").Join("emails USING (email_id)")
1. [sqlc 1.3.0: Generate Type Safe Go From SQL1.3.0 just dropped.](https://golangweekly.com/link/88478/web)

     — sqlc supports a nice set of features, including transactions, array data types, and various migration tools. 1.3.0 just dropped.
1. [go-toml: Go Library for the TOML FormatTOMLwhat it looks like.](https://golangweekly.com/link/88480/web)

     — TOML is a configuration file format invented by one of GitHub’s founders. Here’s what it looks like.
### 🎲 Fun and Side Projects

1. [Profefe: Continuously Collect Profiling Data for Long-Term Postmortem Analysis](https://golangweekly.com/link/88483/web)

     — Vladimir writes in: 

1. [Orchestra: A New Library to Manage Long Running Go Processes](https://golangweekly.com/link/88475/web)


### [ << Prev ](golangweekly-311.md) ------------- [ Next >> ](golangweekly-313.md)