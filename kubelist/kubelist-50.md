## [Kubelist Issue #50 for 2019-01-31](https://kubelist.com/issue/50)

#### DIY K8s vs Hosted, ML and Apache Spark, what else can we ask for?

> As our clusters grow in complexity and load, we start to question everything we thought we knew.  Don&#39;t worry, we&#39;ve got your back (and our own).  This week we bring you some  insights into scaling as you grow, a look inside of a truely large cluster, how to think about just hosted services, and finally, some great news from Apache Spark.

1. [Understanding Kubernetes Cluster Autoscaling ](https://medium.com/kubecost/understanding-kubernetes-cluster-autoscaling-675099a1db92)

    Since the beginning of cloud, the promise of elestic resources has helped us all sleep well at night knowing that if we suddenly went viral, our infra would just simply spin up more instances to take the load.  This promise, while being true, sometimes feels like you are adding entire death stars when all you need is a few more tie fighters.  Enter k8s, giving us the ability to scale up and down pods (aka tie fighters), and now thanks to K8s cluster scaling, nodes for when you actually need those death stars.In this article Ajay Tripathy does a great job of walking us through the power of node autoscaling, how combining pod and node scaling gives you that fine grained control you wish you always had.  
1. [How Periscope Uses Kubernetes to Power Data Science Services](https://thenewstack.io/how-periscope-uses-kubernetes-to-power-data-science-services/)

    Do you think of Airbnb as a data company? Tinder as a ranking problem?  Periscope does, as they are in the business of providing orgs like these the tools to mine and improve their products through data analysis tooling. Persicope sees 300,000 new lines of SQL in their platform daily, has 30TB of working memory across all its systems and manages over twenty-million queries a day.  They pull all of this off on top of K8s.In this article (with podcast!) they dive deep into all of the details.  It truely is fascinating to think how our own k8s clusters would manage similar load and keep on kicking.  The article describes it as a "fun challenge", that's one way to think of it...
1. [Think you can run Kubernetes better than a cloud provider? Think again](https://www.techrepublic.com/article/think-you-can-run-kubernetes-better-than-a-cloud-provider-think-again/)

    We have all fallen into this trap: "I'll do it myself because I can, and its <cheaper|more versatile|less locked in|more fun|better>". From building our own gaming platforms, to handcrafting our own clusters, this is a typical DIY'ers dilemma. TechRepublic and Matt Asay take a look at the pros and cons of building your own cluster vs just letting someone else do it.Speaking from our experience, it's not an easy choice, and sometimes you just can't understand something at depth if you hand it off to a service. On the other hand, it does make it awfully hard to move your actual business forward if you are constantly tweaking CRDs. Asay does a great job walking you through making the decision to build or to buy.
1. [Google announces Kubernetes Operator for Apache Spark](https://www.zdnet.com/article/google-announces-kubernetes-operator-for-apache-spark/)

    Speaking of CRDs and letting others invent and manage things for you, this week we found out Google is going to make it a lot easier for us to run our ML workloads in our clusters via their "Spark Operator" (aka Kubernetes Operator for Apache Spark).  This opens the world of kubectl to all of your favourite data scientists, something we all know they can't wait to use.In truely Google fashion, they have made the Spark Operator a one-click solution to those of us running on GCP. But if the last article left you thinking you should run your own cluster, it's also open-source here.  Should we hold our breath waiting to see if Amazon and Microsoft add it to their own once click offerings? We're not sure.
1. [Tweet of the Week](https://twitter.com/todaywasawesome/status/1088651763459203073)

    While this tweet inspired some of the stories in this week's issue, we're really including it because it reminds us of the Blue Man Group...

### [ << Prev ](kubelist-49.md) ------------- [ Next >> ](kubelist-51.md)