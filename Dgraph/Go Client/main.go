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

	CreateQuery(ctx, dg)
	SearchQuery(ctx, dg)
	// DeleteQuery(ctx, dg, "0x21")
	SearchQuery(ctx, dg)
}

// DeleteQuery deletes a node
func DeleteQuery(ctx context.Context, dg *dgo.Dgraph, id string) {

	// variables := map[string]string{"$id": id}

	variables := map[string]string{"$id": id}
	q := `query createUser($id: string){
		create(func: uid($id)){
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

	mu := &api.Mutation{}
	dgo.DeleteEdges(mu, id, "voertuig")
	dgo.DeleteEdges(mu, id, "name", "role")
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

// CreateQuery creates a node
func CreateQuery(ctx context.Context, dg *dgo.Dgraph) {

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
	q := `query createUser($id: string){
		create(func: uid($id)){
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

// SearchQuery searches nodes
func SearchQuery(ctx context.Context, dg *dgo.Dgraph) {

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := `
	{
		users(func: has(name)) {
		  name
		  role
		  voertuig
		}
	}`

	res, err := txn.Query(ctx, q)
	fmt.Printf("%s\n", res.Json)

	if err != nil {
		log.Fatal(err)
	}
}

// Schema voor dgraph:

// name: string @index(exact) .
// role: string .
// needsservice: bool .
// type: string .
// latitude: float .
// longitude: float .

// type Service {
// 	needsservice
// }

// type Voertuig {
// 	type
// 	latitude
// 	longitude
// }

// type User {
// 	name
// 	role
// }
