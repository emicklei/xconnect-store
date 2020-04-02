package main

import dga "github.com/emicklei/dgraph-access"

func findOrCreateScope(d *dga.DGraphAccess, id string) (*NetworkScope, bool, error) {
	e := &NetworkScope{ID: id}
	c := new(dga.CreateNode)
	c.Node = e
	c.CreateUnless("network_scope_id", id)
	wasCreated, err := d.Do(c)
	return e, wasCreated, err
}
