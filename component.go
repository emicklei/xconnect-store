package main

import dga "github.com/emicklei/dgraph-access"

func upsertComponent(d *dga.DGraphAccess, n *NetworkComponent) error {
	// if component does not exist then create it
	// if component does exist then update it (its scalar predicates)
	u := dga.UpsertNode{Node: n}
	u.InsertUnless("network_component_id", n.ID)
	_, err := d.Do(u)
	return err
}
