## [#325 — August 14, 2020](https://golangweekly.com/issues/325)

1. [Go 1.15 Releasedthe release notesran into issues with TLS connectionsnow fixed](https://golangweekly.com/link/93510/web)

     — Hit up the release notes to go through everything that’s new, but the highlights include substantial improvements to the linker, better allocation in some cases, and various core library improvements. This is a step forward for Go, however, rather than a feature-packed leap. Do note, though, that some people  ran into issues with TLS connections involving certificates without SANs (now fixed in this case but may cause issues elsewhere).
1. [Proposal: Register-Based Go Calling ConventionABI](https://golangweekly.com/link/93512/web)

     — This proposal hopes to increase performance (well, throughput) by changing the Go ABI’s stack-based calling convention to a register-based version. This won’t affect most Go developers at a day to day level, but feel free to read on for the gritty details.
1. [Why Go Modules Are Faster Than GOPATH](https://golangweekly.com/link/93515/web)

     — The short answer is modules usually end up having to download much less code (in some cases, zero code) for a given set of dependencies. Hurray for progress!
1. [AWS Lambda Now Supports Go on Amazon Linux 2](https://golangweekly.com/link/93516/web)

     — AWS Lambda, Amazon’s serverless platform, has supported Go directly for a while now, but support has now been extended to the newer Amazon Linux 2 runtime.
### 💻 Jobs

1. [Senior Software Engineer at Even (Anywhere)](https://golangweekly.com/link/93517/web)

     — Help people break the paycheck-to-paycheck cycle. Build w/ Go, React Native, GraphQL, Postgres, Bazel. Remote encouraged.
1. [Software Engineer - Want to Build a Platform Ecosystem in Go?Apply now.](https://golangweekly.com/link/93518/web)

     — Skool is hiring its 2nd backend engineer in Los Angeles, CA. Go, PostgreSQL, Redis, Elasticsearch, Docker. Apply now.
1. [Find a Job Through Vettery](https://golangweekly.com/link/93519/web)

     — Use Vettery to connect with hiring managers at startups and Fortune 500 companies. It's free for job-seekers.
### 📘 Tutorials

1. [How Go 1.15 Improved Converting Small Integer Values to Interfaces](https://golangweekly.com/link/93520/web)

     — This gets in the weeds a bit, but the short answer is a static array of integers that can be pointed to instead of a heap allocation.
1. [Create a Serverless Go App for AWS Lambda](https://golangweekly.com/link/93521/web)

     — A basic, entry level tutorial to using Go in a serverless context.
1. [Deploying Serverless Go Functions with Netlify](https://golangweekly.com/link/93522/web)

     — It’s a bit convoluted, but the end process of having Go functions in your JAMstack might just be worth it.
1. [Lessons from Go: Keep Calm and Use the Byte Array](https://golangweekly.com/link/93525/web)

     — Byte arrays are ubiquitious in Go, despite being hidden being abstractions in other languages. And this is a good thing.
1. [Let's Build a Concurrent Download ManagerGitHub repo.](https://golangweekly.com/link/93542/web)

     — A practical 20 minute screencast. GitHub repo.
1. [What Is the Extension Interface Pattern in Go?](https://golangweekly.com/link/93524/web)

### 🛠 Code & Tools

1. [Multiple Progress Bar: Progress Bars for Go CLI Appsv5.3.0](https://golangweekly.com/link/93526/web)

     — Straightforward textual progress bars with support for rendering multiple bars at the same time (such as for tasks running concurrently). v5.3.0 just came out.
1. [immudb: A Lightweight, High-Speed Immutable DatabaseGitHub repo.](https://golangweekly.com/link/93528/web)

     — An open source data that boasts no data mutation APIs at all and a ‘tamper-evident’ history system. It’s built in Go. GitHub repo.
1. [Excelize 2.3: A Library for Reading and Writing Excel Filesv2.3.0](https://golangweekly.com/link/93531/web)

     — Read and write XLSX files, set and read cell values, add charts. v2.3.0 lets you set cell values concurrently, and provides some new APIs for improved access to formatting and columns.
1. [Dynamo: An Expressive DynamoDB Libraryimproves support for null and empty values.](https://golangweekly.com/link/93533/web)

     — This week’s new release improves support for null and empty values.
1. [Oscar: Next Generation Building Tool for Nothingobligatory xkcd reference](https://golangweekly.com/link/93535/web)

     — Yes, this is a joke/lighthearted project! You’d use this to make it look like you’re being productive when you’re not (obligatory xkcd reference).
1. [Got: Go Package and CLI Tool for Downloading Large Files Over HTTP Faster](https://golangweekly.com/link/93537/web)

     — The technique used is to concurrently download numerous ranges of the desired file which can be quicker in many cases (but may also be frowned upon if it causes extra load).
1. [httpmock: Easy Mocking of HTTP Responses From External Resources](https://golangweekly.com/link/93538/web)

1. [Sarama: A Go Client Library for Apache Kafka 0.8+](https://golangweekly.com/link/93539/web)


### [ << Prev ](golangweekly-324.md) ------------- [ Next >> ](golangweekly-326.md)