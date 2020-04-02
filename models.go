// Generated by github.com/emicklei/dgraph-access/cmd/dggen. DO NOT EDIT
package main
import (
	dga "github.com/emicklei/dgraph-access"		
)

type NetworkScope struct {
	dga.Node `json:",inline"` 
        ID string `json:"network_scope_id,omitempty"` 
}

type NetworkEndpoint struct {
	dga.Node `json:",inline"` 
        ID string `json:"network_endpoint_id,omitempty"` 
        URL string `json:"url,omitempty"` 
        Host string `json:"host,omitempty"` 
        Protocol string `json:"protocol,omitempty"` 
        TLS bool `json:"tls"` 
        Port int64 `json:"port"` 
}

type NetworkComponent struct {
	dga.Node `json:",inline"` 
        ID string `json:"network_component_id,omitempty"` 
        Product string `json:"product,omitempty"` 
}

