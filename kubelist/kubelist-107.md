## [Kubelist Issue #107 for 2020-11-04](https://kubelist.com/issue/107)

#### Kubectl Plugins 🔌

> Most of us have kubectl installed on our laptops. Building on this, kubectl plugins are a simple and convenient way to package additional CLIs that depend on Kubernetes clusters. And with Krew (see today’s first link), we no longer need to package separate MacOS, Windows, and Linux binaries into homebrew, apt, rpm, snap, and other package managers. If you’ve shipped a CLI before, you know how difficult it is to keep it updated on everyone’s computer. If you haven’t, trust us, it’s difficult. Read on for some new links on creating and packaging a CLI and then delivering it to your Kubernetes users.

1. [Krew](https://krew.sigs.k8s.io/)

    Krew is “the plugin manager for kubectl”. While it’s possible to distribute a kubectl plugin without Krew, this project has made it a lot easier to both publish and consume these plugins. End users need to install krew once, and then they can easily install or update any of the plugins in the krew index. It’s just like homebrew, but for kubectl plugins. 🐙
1. [Krew Plugin Stats](https://datastudio.google.com/u/0/reporting/f74370a0-adcf-4cec-b7bd-a58c638948f5/page/Ufl7)

    Because krew is a centralized index of plugins, we can get this list of download count by plugin. If you are curious about the health of the kubectl plugin ecosystem, or debating if you want to package a utility up as a kubectl plugin or a different format, check out this dashboard. 📈
1. [Krew Release Bot](https://github.com/rajatjindal/krew-release-bot)

    Automation is better than manual process, right? This is a bot that will automatically update your plugin in the Krew index when you create a new release in GitHub. This bot is highly recommended to reduce error-prone copying of SHAs into manifests and manually creating pull requests into the index. 🤖
1. [Krew Plugin Template](https://github.com/replicatedhq/krew-plugin-template)

    For GitHub users, this is a template repo to help you get started with a new plugin. This template will create a new project with Goreleaser, krew manifests and everything needed so you can just start coding your plugin without having to set up the scaffolding and wire up the release process.
1. [Writing your first kubectl plugin with Go](https://bmuschko.com/blog/writing-your-first-kubectl-plugin/)

    If you’d rather learn by doing it “the hard way” once, this is a wonderful tutorial on getting started building a kubectl plugin. Keep in mind that there is a lot that the plugin template repo and other projects do to handle cloud-provider specific authentication options, and you’ll have to implement these yourself when taking this approach. 📝
1. [More plugins](https://github.com/topics/kubectl-plugins)

    The Krew package manager is a great way to discover plugins. But not every plugin is in there, and it’s up to the owner of the plugin to decide to include it. This link is a list of open source repos on GitHub that contain kubectl plugins. There is overlap with what’s in Krew, but there’s also a lot of good standalone plugins that have not been added to the package manager yet.	
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1323599287612284928)

    The CNCF shared a statement about the passing of Dan Kohn, executive director who helped build and shape the organization into what it is today.

### [ << Prev ](kubelist-106.md) ------------- [ Next >> ](kubelist-108.md)