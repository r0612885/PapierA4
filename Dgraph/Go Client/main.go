package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Service struct {
	Uid          string `json:"uid,omitempty"`
	Needsservice bool   `json:"needsservice,omitempty"`
}

type Voertuig struct {
	Uid       string  `json:"uid,omitempty"`
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Service   Service `json:"service,omitempty"`
}

type User struct {
	Uid      string   `json:"uid,omitempty"`
	Name     string   `json:"name,omitempty"`
	Role     string   `json:"role,omitempty"`
	Voertuig Voertuig `json:"voertuig,omitempty"`
}

func main() {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	// CreateUserQuery(ctx, dg)
	// DeleteUserQuery(ctx, dg, "0x3a")
	// SearchAllUsersQuery(ctx, dg)
	// SearchUserQuery(ctx, dg, "0x37")
	DeleteUserQuery(ctx, dg, "0x32")
	// UpdateUserQuery(ctx, dg, "0x32")
}

// UpdateUserQuery updates a usernode
func UpdateUserQuery(ctx context.Context, dg *dgo.Dgraph, id string) {

	pb, err := json.Marshal(User{Uid: id, Name: "Blueface", Role: "VAPEGOD"})
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
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
			voertuig
		}
	}`

	resp, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil || resp == nil {
		log.Fatal(err)
	}
}

// DeleteConnectionBetweenUserAndVehicleQuery deletes the edge between a usernode and a vehiclenode
func DeleteConnectionBetweenUserAndVehicleQuery(ctx context.Context, dg *dgo.Dgraph, id string) {
	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
		}
	}`

	mu := &api.Mutation{}
	dgo.DeleteEdges(mu, id, "voertuig")
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

// DeleteUserQuery deletes a node
func DeleteUserQuery(ctx context.Context, dg *dgo.Dgraph, id string) {

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
		}
	}`

	mu := &api.Mutation{}
	dgo.DeleteEdges(mu, id, "name", "role", "voertuig")
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

// CreateUserQuery creates a node
func CreateUserQuery(ctx context.Context, dg *dgo.Dgraph) {

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	u := User{
		Uid:  "_:user",
		Name: "Wesley Monten",
		Role: "Admin",
		Voertuig: Voertuig{
			Type:      "Heftruck A978",
			Latitude:  41.1551,
			Longitude: 49.1255,
			Service: Service{
				Needsservice: true,
			},
		},
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	variables := map[string]string{"$id": assigned.Uids["user"]}
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
			voertuig {
				type
				latitude
				longitude
				service {
					needsservice
				}
			}
		}
	}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Json, &u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Json))

}

// SearchUserQuery searches a node
func SearchUserQuery(ctx context.Context, dg *dgo.Dgraph, id string) {
	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
			voertuig {
				type
				latitude
				longitude
				service {
					needsservice
				}
			}
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

// SearchQuery searches nodes
func SearchAllUsersQuery(ctx context.Context, dg *dgo.Dgraph) {

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := `
	{
		users(func: has(name)) {
		  name
		  role
		  voertuig {
			  type
			  latitude
			  longitude
			  service {
				  needsservice
			  }
		  }
		}
	}`

	res, err := txn.Query(ctx, q)
	fmt.Printf("%s\n", res.Json)

	if err != nil {
		log.Fatal(err)
	}
}
