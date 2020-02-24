package vehicleservice

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Vehicle struct {
	Uid          string `json:"uid,omitempty"`
	Type         string `json:"type,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
	Needsservice string `json:"needsservice,omitempty"`
}

// GetVehicles gets all vehicles
func GetVehicles() []byte {

	conn, err := grpc.Dial("10.244.1.9:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := `
	{
		vehicles(func: has(type)) {
		  uid
		  type
		  latitude
		  longitude
		  needsservice
      			service{
					dateCompleted
					description  
			}
		}
	}`

	res, err := txn.Query(ctx, q)

	if err != nil {
		log.Fatal(err)
	}

	return res.Json
}

// GetVehicle gets a vehicle
func GetVehicle(id string) []byte {

	conn, err := grpc.Dial("10.244.1.9:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getVehicle($id: string){
		vehicle(func: uid($id)){
			uid
			type
			latitude
			longitude
			needsservice
				service {
				dateCompleted
				description  
			}
		}
	}`

	res, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	return res.Json
}

// CreateVehicle creates a vehicle
func CreateVehicle(v []byte) []byte {

	conn, err := grpc.Dial("10.244.1.9:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   []byte(v),
	}

	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		fmt.Print(err)
	}

	variables := map[string]string{"$id": assigned.Uids["vehicle"]}
	q := `query getVehicle($id: string){
		vehicle(func: uid($id)){
				type
				latitude
				longitude
				needsservice
					service {
						dateCompleted
						description  
				}
			}
	}`

	res, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		fmt.Print(err)
	}

	return res.Json
}

// UpdateVehicle updates a vehicle
func UpdateVehicle(id string, v []byte) []byte {

	conn, err := grpc.Dial("10.244.1.9:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	req := &api.Request{CommitNow: true}

	mu := &api.Mutation{SetJson: []byte(v)}
	req.Mutations = []*api.Mutation{mu}

	if _, err := dg.NewTxn().Do(ctx, req); err != nil {
		log.Fatal(err)
	}

	variables := map[string]string{"$id": id}
	q := `query getVehicle($id: string){
		vehicle(func: uid($id)){
			type
			latitude
			longitude
			needsservice
				service {
					dateCompleted
					description  
			}
		}
	}`

	res, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		log.Fatal(err)
	}

	return res.Json
}

// DeleteVehicle deletes a vehicle
func DeleteVehicle(id string) {

	conn, err := grpc.Dial("10.244.1.9:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getVehicle($id: string){
		vehicle(func: uid($id)){
			type
			latitude
			longitude
			needsservice
				service {
					dateCompleted
					description  
			}
		}
	}`

	mu := &api.Mutation{CommitNow: true}
	dgo.DeleteEdges(mu, id, "type", "latitude", "longitude", "needsservice")

	resp, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || resp == nil {
		log.Fatal(err)
	}
}
