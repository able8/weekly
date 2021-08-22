## [Kubelist Issue #104 for 2020-10-14](https://kubelist.com/issue/104)

#### Observability with OpenTelemetry 🔭

> This week on the <a href="https://kubelist.com/podcast/">Kubelist Podcast</a>, we have a conversation with Ben Sigelman from LightStep about the OpenTelemetry project. This has us thinking about observability and collecting some good links we’ve been saving related to tracing, metrics, and logging using OpenTelemetry. On the podcast, Ben shares some thoughts on these “pillars of observability”. 

1. [OpenTelemetry Registry](https://opentelemetry.io/registry/)

    OpenTelemetry is just a specification, so the registry and implementations are a proxy to measure how quickly it’s being adopted. The registry is where you’ll find collectors and exporters to instrument your code and view the data using the platform of your choice. Remember that OpenTelemetry is an early project, but the number of integrations in this registry show the demand and rapid adoption of this platform. 📋
1. [Getting Started with OpenTelemetry for Java](https://thenewstack.io/getting-started-with-opentelemetry-for-java/)

    Like most OpenTelemetry tutorials, this starts out with the necessary architecture explanations to make sure you know how the system is going to work. If you’re a Java developer, this is a great tutorial on how to add the OpenTelemetry dependencies into your application and get started with the auto-instrumentation capabilities of the SDK. ☕
1. [From 0 to Insight with OpenTelemetry in Go](https://www.honeycomb.io/blog/from-0-to-insight-with-opentelemetry-in-go/)

    For Go developers, don’t expect auto-instrumentation. But Liz Fong-Jones from Honeycomb has shared a great tutorial on how to instrument your Go projects with the OpenTelemetry SDK. She shows using the exporter to Honeycomb, which is a good place to start, if you don’t already have an export target. 🍯
1. [Instrument your Python applications with Datadog and OpenTelemetry](https://www.datadoghq.com/blog/instrument-python-apps-with-datadog-and-opentelemetry/)

    DataDog supports and has contributed to the OpenTelemetry project. This post from their blog is a great walkthrough of what’s possible today with OpenTelemetry if you have Python and are a DataDog customer. 🐶
1. [zPages in OpenTelemetry](https://medium.com/opentelemetry/zpages-in-opentelemetry-2b080a81eb47)

    This is a welcome addition to a complex ecosystem. The idea of “zPages” (think healthz, metricsz, etc) is a lot simpler than setting up an external system like Zipkin or Jaeger. The internal diagnostics isn’t as powerful as dedicated systems, and isn’t where you want to end your integration efforts; however if you’re looking to get started with OpenTelemetry and add more observability to your application, this is a great way to deliver value in a few hours without needing a complex integration plan.
1. [Troubleshoot Your Applications Faster with OpenTelemetry and New Relic One](https://blog.newrelic.com/product-news/opentelemetry-user-experience/)

    The folks over at NewRelic have written a detailed walk through of using OpenTelemetry to send data into NewRelic. The NewRelic user interface is mature, and for teams that already have an investment here, connecting the standards-based OpenTelemetry SDK into NewRelic One might be a perfect fit.
1. [Tweet of the week](https://twitter.com/coreos/status/1314623658095075329)

    This marks the end of an era. Also, this is the best gif collection we’ve seen in a while.

### [ << Prev ](kubelist-103.md) ------------- [ Next >> ](kubelist-105.md)