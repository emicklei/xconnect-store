Melisco (meta-listen-connect)


the connect problem ; finding the pink relations

from nemo-load-balancer to nemo-esp

# default value is "global" which indicates all used host names in this configuration are public resolveable.
# if not, provide a name that identifies a container in which hostnames could be resolved.

/v1/xconnect?networkscope=kramphub-valentio-production/valentio-cluster-v1/catlyn.catlyn

resource.labels.project_id="team-rebooted"
resource.labels.location="europe-west1-b"
resource.labels.cluster_name="dgraph-cluster-v1"
resource.labels.namespace_name="default"
resource.labels.pod_name="dgraph-ratel-59d8d7b848-vb7ch"


-------------------------
query X {
  q(func: type(NetworkEndpoint)) {
   uid
   dgraph.type
   expand(NetworkEndpoint)
   x {
     uid
   }
  }
}