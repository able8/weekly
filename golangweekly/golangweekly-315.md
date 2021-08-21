## [#315 — June 5, 2020](https://golangweekly.com/issues/315)

1. [ZZT in Go (Using a Pascal-to-Go Converter)ZZTPascal source code to Go](https://golangweekly.com/link/89670/web)

     — Back in the 90s I loved a text-based game called ZZT. It was built by Tim Sweeney who went on to create the Unreal Engine(!) and still has a bit of a cult following. So it’s fascinating to see this work done in translating the Pascal source code to Go and, of course, the similarities between Pascal and Go generally too.
1. [GoLand 2020.2 Early Access Program Is Open](https://golangweekly.com/link/89673/web)

     — GoLand is a Go IDE and its EAP build is free to use for 30 days. This EAP has some wicked new features, refactorings, and plugins.
1. ['The Go Compiler Needs to Be Smarter'?](https://golangweekly.com/link/89675/web)

     — A pretty opinionated piece from a reasonably academic perspective, but it raises some interesting points about the Go compiler’s tradeoffs.
1. [Writing Go CLIs With Just Enough Architecture](https://golangweekly.com/link/89676/web)

     — A developer inspired by some of the recent CLI app tutorials reflects on his own approach where he provides “just enough” architecture to build solid Go CLI apps.
1. [Bare Metal RISC-V Programming in Go](https://golangweekly.com/link/89677/web)

     — Did you know that Go 1.14 includes experiment support for 64 bit RISC-V? I didn’t, but the creator of Embedded Go was keen to try it out on a RISC-V microcontroller and here’s how it went.
### 💻 Jobs

1. [Want to Build a Platform Ecosystem in Go?Apply now](https://golangweekly.com/link/89678/web)

     — Skool is hiring its 2nd backend engineer in Los Angeles, CA. Go, PostgreSQL, Redis, Elasticsearch, Docker. Apply now.
1. [Enjoy Building Scalable Infrastructure in Go? Stream Is HiringApply now](https://golangweekly.com/link/89680/web)

     — Like coding in Go? We do too. Stream is hiring in Amsterdam. Apply now.
1. [Find A Job Through Vettery](https://golangweekly.com/link/89681/web)

     — Vettery specializes in tech roles and is completely free for job seekers. Create a profile to get started.
### 📚 Articles & Tutorials

1. [How to Write a Lexer in Go](https://golangweekly.com/link/89682/web)

     — A lexer (short for lexical analyzer) is the first phase in most modern compilers and is the process where textual code is turned from characters into a more useful construct of ‘tokens’.
1. [A Go RabbitMQ Beginners' Tutorial](https://golangweekly.com/link/89683/web)

     — A brief review of what RabbitMQ is followed by a coding session where you’ll create a queue, publish a message, and consume that message.
1. [Getting Hands-On with io_uring from Go](https://golangweekly.com/link/89684/web)

     — io_uring is a new way to perform high performance asynchronous I/O in Linux that removes some bottlenecks compared to existing methods. This is a long, technical, interesting post with Go examples towards the end.
1. [Unicode Support in Go Source Code](https://golangweekly.com/link/89686/web)

     — Unicode support is good enough to allow Ω to be a function name, but does anyone really do that kind of thing? Yes.. and this post looks at some other scary things you can do too 😆
1. [A Subtle Trap When Formatting time.Time Values](https://golangweekly.com/link/89687/web)

     — “Unless you’re sure you know what you’re doing and where your time.Time values came from, using a bare .Format() is probably a mistake.”
1. [Converting JSON to a Struct in Go](https://golangweekly.com/link/89688/web)

     — This covers a bit more than the basics, such as custom unmarshalling.
1. [The Trouble with Databaseswhat she wished more developers knew about databases](https://golangweekly.com/link/89689/web)

     — Google’s Jaana Dogan recently wrote a very popular article about what she wished more developers knew about databases and here she reflects on these topics and more in a solid hour with the Go Time podcast crew. Worthwhile weekend listening.
1. [Logging Without Losing Money or Context](https://golangweekly.com/link/89691/web)

     — Logging services are expensive, but logs are important. What’s a poor application developer to do?
1. [Waiting on Goroutines](https://golangweekly.com/link/89692/web)

     — Many people default to using WaitGroups for this, but there are multiple (and more common) ways to handle your waiting.
### 🛠 Code & Tools

1. [go-yaml: YAML Encoding, Decoding and Querying for Goexample](https://golangweekly.com/link/89693/web)

     — As of this week’s 1.7 release, it now supports YAMLPath, a way of extracting values or ASTs by a simple query language (example).
1. [Heimdall: An Enhanced HTTP Client for Go](https://golangweekly.com/link/89694/web)

     — Enhanced, how? A built in circuit breaker to control failing requests, multiple and custom retry strategies, a fluent API..
1. [Smocker: A Simple and Efficient HTTP Mock Server](https://golangweekly.com/link/89696/web)

     — Smocker uses YAML to define mocks and responses, but there is a handy user interface, as well. Use it in your dev and test environments to cut out external services.
1. [Oragono 2.1: A Modern IRC Server Written in Go](https://golangweekly.com/link/89697/web)

     — Supports UTF-8 and IRCv3 features. 2.1 adds native support for WebSockets.
1. [DBCore: Code Generation Powered by Your Database](https://golangweekly.com/link/89698/web)

     — If you’re thinking “I bet there is a YAML file involved.”, you’re right.

### [ << Prev ](golangweekly-314.md) ------------- [ Next >> ](golangweekly-316.md)