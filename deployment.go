package main

import (
	"log"

	dga "github.com/emicklei/dgraph-access"
	"github.com/emicklei/xconnect"
)

func upsertDeployment(d *dga.DGraphAccess, networkPath string, cfg xconnect.Config) error {
	scope, created, err := findOrCreateScope(d, networkPath)
	if err != nil {
		return err
	}
	if created {
		log.Println("NetworkScope ", scope.Node.GetUID(), ":", scope.ID)
	}
	n := new(NetworkComponent)
	n.ID = networkPath + "#" + cfg.Meta.Name
	if err := upsertComponent(d, n); err != nil {
		return err
	}
	log.Println("NetworkComponent ", n.Node.GetUID(), ":", n.ID)

	if err := d.Fluent().CreateEdge(scope, "component", n); err != nil {
		return err
	}

	for key, entry := range cfg.Listen {
		e := new(NetworkEndpoint)
		e.ID = n.ID + "#" + key
		e.Host = entry.Host
		if err := upsertEndpoint(d, e); err != nil {
			return err
		}
		log.Println("NetworkEndpoint ", e.Node.GetUID(), ":", e.ID)
		if err := d.Fluent().CreateEdge(n, "listen", e); err != nil {
			return err
		}
	}

	for key, entry := range cfg.Connect {
		e := new(NetworkEndpoint)
		e.ID = n.ID + "#" + key
		e.Host = entry.Host
		if err := upsertEndpoint(d, e); err != nil {
			return err
		}
		log.Println("NetworkEndpoint ", e.Node.GetUID(), ":", e.ID)
		if err := d.Fluent().CreateEdge(n, "connect", e); err != nil {
			return err
		}
	}

	return nil
}
