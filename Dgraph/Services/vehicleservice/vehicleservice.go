package vehicleservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)


type Vehicle struct {
	Uid       string  `json:"uid,omitempty"`
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Needsservice bool   `json:"needsservice"`	
}

// CreateVehicle creates a vehicle
func CreateVehicle(v Vehicle) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
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
	}
	pb, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	variables := map[string]string{"$id": assigned.Uids["vehicle"]}
	q := `query getVehicle($id: string){
		vehicle(func: uid($id)){
				type
				latitude
				longitude
				needsservice
			}
	}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Json, &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Json))

}


// UpdateVehicle updates a vehicle
func UpdateVehicle(id string, v Vehicle) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	pb, err := json.Marshal(Vehicle{Uid: id, Type: v.Type, Latitude: v.Latitude, Longitude: v.Longitude, Needsservice: v.Needsservice})
	if err != nil {
		log.Fatal(err)
	}

	req := &api.Request{CommitNow: true}

	mu := &api.Mutation{SetJson: pb}
	mu.SetJson = pb
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
		}
	}`

	resp, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil || resp == nil {
		log.Fatal(err)
	}
}

// DeleteVehicle deletes a vehicle
func DeleteVehicle(id string) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
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
		}
	}`

	mu := &api.Mutation{}
	dgo.DeleteEdges(mu, id, "type", "latitude", "longitude", "needsservice")
	mu.CommitNow = true

	resp, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || resp == nil {
		log.Fatal(err)
	}
}

// GetVehicle gets a vehicle
func GetVehicle(id string) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
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
		}
	}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Json, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Json))
}

// GetAllVehicles gets all vehicles
func GetAllVehicles() {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
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
		  type
		  latitude
		  longitude
		  needsservice
		}
	}`

	res, err := txn.Query(ctx, q)
	fmt.Printf("%s\n", res.Json)

	if err != nil {
		log.Fatal(err)
	}
}