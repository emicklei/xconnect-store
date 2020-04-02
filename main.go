package main

//go:generate dggen -s dgraph.schema > models1.go

import (
	"io/ioutil"
	"log"

	"github.com/emicklei/tre"
	"github.com/emicklei/xconnect"
	"gopkg.in/yaml.v2"

	"context"
	"flag"

	dgo "github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	dga "github.com/emicklei/dgraph-access"
	"google.golang.org/grpc"
)

var drop = flag.Bool("drop", false, "cleanup the database at startup")

func readXConnect(f string) xconnect.Config {
	content, _ := ioutil.ReadFile(f)
	var doc xconnect.Document
	if err := yaml.Unmarshal(content, &doc); err != nil {
		log.Fatalln(err)
	}
	return doc.Config
}

func main() {
	flag.Parse()
	ctx := context.Background()
	client := newClient()

	if *drop {
		log.Println("Cleaning up the database")
		if err := client.Alter(ctx, &api.Operation{DropAll: true}); err != nil {
			log.Fatal("drop all failed ", err)
		}
	}

	// create an accessor
	dac := dga.NewDGraphAccess(client)

	// for debugging only
	//dac = dac.WithTraceLogging()

	// set schema
	if err := dac.InTransactionDo(ctx, alterSchema); err != nil {
		log.Fatal(err)
	}

	app := readXConnect("app.yaml")
	db := readXConnect("db.yaml")

	// insert data
	if err := dac.InTransactionDo(ctx, func(d *dga.DGraphAccess) error {
		return upsertDeployment(d, "milano-prd/k8s-cluster", app)
	}); err != nil {
		log.Fatal(err)
	}
	if err := dac.InTransactionDo(ctx, func(d *dga.DGraphAccess) error {
		return upsertDeployment(d, "milano-prd/k8s-cluster", db)
	}); err != nil {
		log.Fatal(err)
	}
}

func alterSchema(da *dga.DGraphAccess) error {
	s, err := ioutil.ReadFile("dgraph.schema")
	if err != nil {
		return tre.New(err, "file", "dgraph.schema")
	}
	return da.Fluent().AlterSchema(string(s))
}

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
