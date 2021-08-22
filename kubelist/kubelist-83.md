## [Kubelist Issue #83 for 2020-05-06](https://kubelist.com/issue/83)

#### Operating State ✍️

> Kubernetes has long been pigeonholed as a platform for stateless applications, however a new generation of Kubernetes-native stateful applications have been emerging. This is largely due to the introduction of the operator pattern that allows for stateful applications to be managed through automated operations. Once installed, these operators extend Kubernetes to allow built-in support for new, custom resources.

1. [Confluent Operator](https://www.confluent.io/confluent-operator/)

    Chances are that you’ve been exposed to Kafka at some point in your career. While unmatched in runtime scalability and performance, Kafka is not exactly known to have an easy installation and management process. Enter Confluent and their Kubernetes Operator. If you want to run Kafka in your Kubernetes cluster, using the codified expertise built into this Operator would be the fast-track to a production grade Kafka deployment.
1. [MongoDB Operator](https://docs.mongodb.com/kubernetes-operator/master/tutorial/install-k8s-operator/)

    The MongoDB team has shipped an Operator to streamline provisioning a production-ready MongoDB Replica Set in Kubernetes. A powerful feature of this Operator is that it simplifies the multiple different topologies for deploying MongoDB in your cluster. Whether you want a cloud-connected instance that leverages CloudManager as a control plane for your on-premise data-plane or you need to deploy your own management plane alongside your database in an airgapped environment, this operator has you covered! 💪
1. [TiDB Operator](https://pingcap.com/docs/tidb-in-kubernetes/v1.0/tidb-operator-overview/)

    TiDB is a horizontally scalable, MySQL compliant, modern database engine. Running this at scale requires some pretty specific knowledge about how to configure everything. The TiDB Operator is a true K8s operator that provides built-in expertise to manage everything perfectly. This is another example of an Operator that handles installation well, but continues to provide ongoing expertise in daily automation for scaling, log rotation, configuration, failover and more.
1. [SchemaHero](https://schemahero.io)

    Once you have your database running in your cluster, you’ll need to create and manage your schema too. SchemaHero is a Kubernetes operator to declaratively create and manage database migrations. This is an example of an Operator that is more than install/upgrade, but instead attempts to automate a task that traditionally is tedious and error prone (creating ALTER TABLE statements and manually managing a database schema). 🦸
1. [Elasticsearch ECK Operator](https://www.elastic.co/blog/introducing-elastic-cloud-on-kubernetes-the-elasticsearch-operator-and-beyond)

    Elasticsearch isn’t a database, but it is often treated as a one. And in all reality, even if you aren’t using Elasticsearch as a database, you’ll probably bundle it into the database bucket when it comes to how to manage. The ECK operator is (of course) all open source, and is more than simply Elasticsearch. The operator contains CRDs to manage Elasticsearch, Kibana and Elastic’s APM service. ⚙️
1. [Celebrating Helm's CNCF Graduation](https://helm.sh/blog/celebrating-helms-cncf-graduation/)

    Congratulations to the entire Helm team (and everyone involved) for moving to a graduated CNCF project. This is not an easy feat considering there are only 10 projects that have achieved this status. 🎉
1. [Tweet of the week](https://twitter.com/CloudNativeFdn/status/1255518597448962053)

    It's official, KubeCon 2020 is a virtual event!

### [ << Prev ](kubelist-82.md) ------------- [ Next >> ](kubelist-84.md)