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
query NE {
  q(func: type(NetworkEndpoint)) {
   uid
   expand(_all_)
  }
}
-------------------------
query NC {
  q(func: type(NetworkComponent)) {
   uid
   expand(_all_)
  }
}
-------------------------
query X {
  q(func: type(NetworkEndpoint)) {
   uid
   expand(NetworkEndpoint)
   x {
     uid
     expand(_all_)
   }
  }
}
-------------------------
query X {
        q(func: type(NetworkScope)) {
    					  dgraph.type
    						network_scope_id
                uid
                components {
                      network_component_id
                      dgraph.type
                      uid
                    	connects {
                        network_endpoint_id
                        dgraph.type
                        uid
                        x {
                          dgraph.type
                          uid
                        }
                      }
                      listens {
                        network_endpoint_id
                        dgraph.type
                        uid
                        x {
                          dgraph.type
                          uid
                        }
                      }
                    }                	                    						
        }
}