package main

import dga "github.com/emicklei/dgraph-access"

func findOrCreateEndpoint(d *dga.DGraphAccess, id string) (*NetworkEndpoint, bool, error) {
	e := &NetworkEndpoint{ID: id}
	c := new(dga.CreateNode)
	c.Node = e
	c.CreateUnless("network_endpoint_id", id)
	wasCreated, err := d.Do(c)
	return e, wasCreated, err
}

func upsertEndpoint(d *dga.DGraphAccess, e *NetworkEndpoint) error {
	// if component does not exist then create it
	// if component does exist then update it (its scalar predicates)
	u := dga.UpsertNode{Node: e}
	u.InsertUnless("network_endpoint_id", e.ID)
	_, err := d.Do(u)
	return err
}
