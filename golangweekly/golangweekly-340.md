## [#340 — November 27, 2020](https://golangweekly.com/issues/340)

1. [The Latest Draft Design for Type Parameters (Enabling 'Generics'); Now Due in 2022 with Go 1.18mentioned](https://golangweekly.com/link/99343/web)

     — Russ Cox has mentioned that the Go compiler will require quite a lot of work to support type parameters so you’ll be waiting until Go 1.18 lands in about 14 months (which sounds closer than '2022' to me 😄)
1. [Anonymous Structs: What and Why?](https://golangweekly.com/link/99345/web)

     — An anonymous struct is a struct defined without a name meaning it can’t be referenced elsewhere. Pointless? No. There are some nice use cases like nested structs or marshalling JSON data, say, where Lane argues anonymous structs handily beat map[string]interface{}
1. [Why GitHub's CLI Team Switched From Ruby to Go](https://golangweekly.com/link/99347/web)

     — GitHub famously uses a lot of Ruby (and was founded by four Rubyists!) so it’s interesting to see where it doesn’t make sense. This 3 minute portion of a podcast with Mislav Marohnić of the GitHub CLI team breaks it down. Portability, basically.
1. [Blackrota: Looking Into a Heavily Obfuscated Backdoor Written in Go](https://golangweekly.com/link/99348/web)

     — Netlab is a Chinese security research company and have uncovered what they call “the most obfuscated Go-developed ELF-formatted malware we’ve found” although apparently such projects are not considered common. Lots of details on the obfuscation here though.
1. [An Online Go Scaling Conference Next Thursday (Dec 3)](https://golangweekly.com/link/99349/web)

     — Go Systems Conf SF is a free online Go event taking place next week focusing on building and scaling Go-powered systems. Nice speaker lineup from companies like Monzo, VMware, and DigitalOcean.
### 💻 Jobs

1. [Lead Backend Engineer (Go)](https://golangweekly.com/link/99350/web)

     — Our mission is to build collaborative software that’s Super-Fast and Nice-To-Use. We are seeking a talented Lead Backend Engineer (Golang) to help us achieve our mission.
1. [Golang Developer at X-Team (Remote)](https://golangweekly.com/link/99351/web)

     — Join the most energizing community for developers and work on projects for Riot Games, FOX, Sony, Coinbase, and more.
1. [Find Your Next Job Through Vettery](https://golangweekly.com/link/99352/web)

     — Create a profile on Vettery to connect with hiring managers at startups and Fortune 500 companies. It's free for job-seekers.
### 📘 Tutorials and Stories

1. [Python Extensions with Rust and Go](https://golangweekly.com/link/99353/web)

     — Two friends ponder if they should create Python extensions in Rust or Go. The answer is…wait for it…it depends!
1. [Connecting to RabbitMQ from Go](https://golangweekly.com/link/99354/web)

     — RabbitMQ is a popular message broker which has good Go support.
1. [Building a Blog with Micro in Go](https://golangweekly.com/link/99356/web)

     — A step-by-step tutorial creating a quick CRUD app for posts. Micro has a robust CLI utility that generates microservices for your cloud-native platform. That’s a lot of buzzwords, but Micro is worth a look.
1. [The Secret Life of Gophers](https://golangweekly.com/link/99357/web)

     — Mat Ryer, Kris Brandow, Angelica Hill, and Natalie Pistunovich talk about work/life balance, work environment must-haves, communication tips, developer tool recommendations, and more.
### 🛠 Code & Tools

1. [Clock: A Small Library for Mocking Time](https://golangweekly.com/link/99358/web)

     — Provides an interface around the time package so an app can use the realtime clock while tests use a mock clock.
1. [Box CLI Maker: Make Highly Customized Boxes for Your CLI](https://golangweekly.com/link/99359/web)

     — I think the author of this packages really likes boxes! Single lines, double lines, a mixture of the two.. rounded corners, only the corners, and a lot more options await you here.
1. [Machinery 2.0.0: An Async Job Queue Based on Distributed Message Passing](https://golangweekly.com/link/99360/web)

     — It can use Redis, Memcached, DynamoDB, RabbitMQ, or MongoDB as its backend store.
1. [Ethr: A Network Performance Measurement Tool for TCP, UDP and HTTP](https://golangweekly.com/link/99362/web)

     — Yep, this comes from Microsoft. It’s nice to see them writing Go, and it works across multiple platforms too.
1. [Smithy: A Tiny Git Forge Written in GoGitea](https://golangweekly.com/link/99363/web)

     — A simple frontend for your git repository if something like Gitea feels a bit too much?
1. [lambda-go-api-proxy: For Porting APIs on Go Frameworks (Such as 'Gin') to AWS Lambda + API Gateway](https://golangweekly.com/link/99365/web)

     — Also supports Negroni, GorillaMux, and plain HandlerFunc.
1. [CloudQuery: Transforms Your Cloud Infrastructure into SQL Queryable Tables](https://golangweekly.com/link/99366/web)

     — Open source, written in Go. select * from aws_ec2_instances seems nicer than a CLI. Azure and GCP support are still pending.

### [ << Prev ](golangweekly-339.md) ------------- [ Next >> ](golangweekly-341.md)