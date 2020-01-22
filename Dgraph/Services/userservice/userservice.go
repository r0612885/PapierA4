package userservice

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
	Needsservice bool   `json:"needsservice,omitempty"`
}

type User struct {
	Uid      string   `json:"uid,omitempty"`
	Name     string   `json:"name,omitempty"`
	Role     string   `json:"role,omitempty"`
	Vehicle  Vehicle  `json:"vehicle,omitempty"`
}

// CreateUser creates a user
func CreateUser() {

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

	u := User{
		Uid:  "_:user",
		Name: "Wesley Monten",
		Role: "Admin",
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
		user(func: uid($id)){
			name
			role
			vehicle {
				type
				latitude
				longitude
				needsservice
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

// UpdateUser updates a user
func UpdateUser(id string) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	pb, err := json.Marshal(User{Uid: id, Name: "Nick Wouters", Role: "Admin"})
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
		}
	}`

	resp, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil || resp == nil {
		log.Fatal(err)
	}
}

// DeleteUser deletes a user
func DeleteUser(id string) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
		}
	}`

	mu := &api.Mutation{}
	dgo.DeleteEdges(mu, id, "name", "role", "vehicle")
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

// DeleteConnectionBetweenUserAndVehicle deletes the connection between a user and a vehicle
func DeleteConnectionBetweenUserAndVehicle(id string) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
		}
	}`

	mu := &api.Mutation{}
	dgo.DeleteEdges(mu, id, "vehicle")
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

// GetUser gets a user
func GetUser(id string) {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		get(func: uid($id)){
			name
			role
			vehicle {
				type
				latitude
				longitude
				needsservice
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

// GetAllUsers gets all users
func GetAllUsers() {

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
		users(func: has(name)) {
		  name
		  role
		  vehicle {
			  type
			  latitude
			  longitude
			  needsservice
		  }
		}
	}`

	res, err := txn.Query(ctx, q)
	fmt.Printf("%s\n", res.Json)

	if err != nil {
		log.Fatal(err)
	}
}
