## [Kubelist Issue #113 for 2021-01-20](https://kubelist.com/issue/113)

#### Backing up Kubernetes Clusters 💾

> Now that it’s become more common to run stateful workloads on Kubernetes, backup (and restore) starts to become more important. Remember that the scheduler can stop, move, restart and control the lifecycle of a component that has critical and irreplaceable data. To deal with this, it’s often recommended to choose a cloud-native database or stateful component that is built to operate in a Kubernetes environment. But even so, having a good and reliable backup stored outside of the cluster (or even the region) is at least a good idea, and possibly even a requirement to maintain SOC-2 compliance.

1. [Kubernetes Backup Tools: Comparing Cohesity, Kasten, OpenEBS, Portworx, Rancher Longhorn, and Velero](https://portworx.com/kubernetes-backup-tools/)

    If you don’t have previous (non-K8s) experience designing a backup solution, this guide from Ryan Wallner at Portworx will help you understand the three types of backups that can be implemented. Then it offers a quick explanation of six popular backup tools, and an explanation of which types of backup and restore they each offer. There are some newer solutions not covered on this writeup, but it’s an incredibly useful guide to read before designing your own backup or disaster recovery solution. 💽
1. [Kubernetes Backup vs Disaster Recovery](https://portworx.com/kubernetes-backup-vs-kubernetes-dr/)

    Another post from Ryan Wallner at Portworx; this one digs in deeper to explain the differences between backup/restore and disaster recovery. While these two terms might sound like the synonyms, they often aren’t. Many times you’ll want different solutions for different backup/restore vs disaster recovery. And they can be implemented very differently. Whether you need to restore to an existing cluster or plan for a broader outage, this is a good guide to help think about how to approach each. 🌪
1. [Backup and Restore of Kubernetes Stateful Application Data with CSI Volume Snapshots](https://itnext.io/backup-and-restore-of-kubernetes-stateful-application-data-with-csi-volume-snapshots-14ce9e6f3778)

    Ok, so this one is a little technical. If you missed it, a feature called Volume Snapshots was moved to GA in December’s release of Kubernetes 1.20. This post is one of the first we’ve seen showcasing this new feature. Here, Zhimin Wen walks through using this built-in feature on a cluster with Rook/Ceph handling persistence storage. It’s a good example to showcase how the core feature works. Then Zhimin shows how to use Velero to drive the CSI Volume Snapshot functionality in the cluster, which is a little more approachable. 📸
1. [How to use CSI Volume Snapshotting with Velero](https://velero.io/blog/csi-integration/)

    So now that there’s a native CSI Volume Snapshot feature in Kubernetes, should you remove Velero? As we just saw in the previous link, Velero plays very nicely with the now-GA CSI Volume Snapshot functionality in Kubernetes 1.20. Note that this article was written for K8s 1.17, when the feature was beta. This quick and short guide from the Velero team at VMWare shows how to enable the CSI Snapshot feature in Velero and use it with your normal workflows to back up using Velero.
1. [Kubernetes Backup: Improve already awesome Velero with OpenEBS](https://blog.mayadata.io/openebs/suggesting-ways-to-improve-already-awesome-velero)

    From Harshita Sharma at Mayadata (creators of OpenEBS and Litmus Chaos), this is a detailed walkthrough of using Velero, specifically when using OpenEBS as a PVC provisioner. The end of this post shows off a really cool feature of OpenEBS if you take this approach – easy migration between on-prem and IaaS or different cloud providers. Note that this guide was written before the CSI Volume Snapshot feature, and instead relies on Restic with Velero.
1. [Tutorial: Kubernetes-Native Backup and Recovery With Stash](https://appfleet.com/blog/kubernetes-native-backup-and-recovery-with-stash/)

    Last but not least, Stash is a tool that we haven’t mentioned yet. This is an open source backup and restore solution that we think is worth sharing. The link above is an intro-level, learn Stash guide. What’s interesting is the walkthrough shows how to back up and restore a Postgresql database; something that more mature backup utilities don’t always handle very well because of locked files. 🗃
1. [Tweet of the week](https://twitter.com/stefanprodan/status/1351451466075152385)

    If you haven’t heard, Pod Security Policies are going away in 1.25 (you still have plenty of time to handle this). But start reading up on options if you depend on PSPs today so that you aren’t caught off guard.

### [ << Prev ](kubelist-112.md) ------------- [ Next >> ](kubelist-114.md)