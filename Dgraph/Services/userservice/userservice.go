package userservice

import (
	"context"
	"encoding/json"

	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Vehicle struct {
	Uid          string  `json:"uid,omitempty"`
	Type         string  `json:"type,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Needsservice bool    `json:"needsservice"`
}

type User struct {
	Uid     string  `json:"uid,omitempty"`
	Name    string  `json:"name,omitempty"`
	Role    string  `json:"role,omitempty"`
	Vehicle Vehicle `json:"vehicle,omitempty"`
}

// GetUsers gets all users
func GetUsers() (string, bool) {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
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
	      uid
		  name
		  role
		  vehicle {
			  type
			  latitude
			  longitude
			  needsservice
			  	service {
					dateCompleted
					description  
				}
		  	}
		}
	}`

	res, err := txn.Query(ctx, q)

	if err != nil {
		error = true
		log.Fatal(err)
	}

	return string(res.Json), error
}

// GetActiveUsers gets all users who are currently operating a vehicle
func GetActiveUsers() (string, bool) {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
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
		query(func: has(vehicle)) {
			name
			password
			role
		  		vehicle {
					type
					needsservice
					latitude
					longitude
						service {
							dateCompleted
							description  
					}
				}
			}
	  }`

	res, err := txn.Query(ctx, q)

	if err != nil {
		error = true
		log.Fatal(err)
	}

	return string(res.Json), error
}

// GetUser gets a user
func GetUser(id string) (string, bool) {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		user(func: uid($id)){
			name
			password
			role
		  		vehicle {
					type
					needsservice
					latitude
					longitude
						service {
							dateCompleted
							description  
						}
				  }
			}
	}`

	res, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		error = true
		log.Fatal(err)
	}

	return string(res.Json), error
}

// CreateUser creates a user
func CreateUser(u string) (string, bool) {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
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
		SetJson:   []byte(u),
	}

	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	variables := map[string]string{"$id": assigned.Uids["user"]}
	q := `query getUser($id: string){
		user(func: uid($id)){
			name
			password
			role
		  		vehicle {
					type
					needsservice
					latitude
					longitude
						service {
							dateCompleted
							description  
					}
				}
		}
	}`

	res, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		error = true
		log.Fatal(err)
	}

	return string(res.Json), error
}

// CreateConnectionBetweenVehicleAndUser connection between a user and a vehicle
func CreateConnectionBetweenVehicleAndUser(vehicleID string, userID string) (string, bool) {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	pb, err := json.Marshal(User{Uid: userID, Vehicle: Vehicle{Uid: vehicleID}})
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

	variables := map[string]string{"$id": userID}

	q := `query getUser($id: string){
		user(func: uid($id)){
			name
			password
			role
		  		vehicle {
					type
					needsservice
					latitude
					longitude
						service {
							dateCompleted
							description  
					}
			}
		}
	}`

	res, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		error = true
		log.Fatal(err)
	}

	return string(res.Json), error
}

// UpdateUser updates a user
func UpdateUser(id string, u string) (string, bool) {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	if err != nil {
		log.Fatal(err)
	}

	req := &api.Request{CommitNow: true}

	mu := &api.Mutation{SetJson: []byte(u)}
	req.Mutations = []*api.Mutation{mu}

	if _, err := dg.NewTxn().Do(ctx, req); err != nil {
		log.Fatal(err)
	}

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		user(func: uid($id)){
			name
			password
			role
		}
	}`

	res, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		error = true
		log.Fatal(err)
	}

	return string(res.Json), error
}

// DeleteUser deletes a user
func DeleteUser(id string) bool {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		user(func: uid($id)){
			name
			password
			role
		}
	}`

	mu := &api.Mutation{CommitNow: true}
	dgo.DeleteEdges(mu, id, "name", "role", "vehicle")

	res, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		error = true
		log.Fatal(err)
	}

	res, err = dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		error = true
		log.Fatal(err)
	}

	return error
}

// DeleteConnectionBetweenUserAndVehicle deletes the connection between a user and a vehicle
func DeleteConnectionBetweenUserAndVehicle(id string) bool {

	error := false

	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$id": id}
	q := `query getUser($id: string){
		user(func: uid($id)){
			name
			password
			role
		}
	}`

	mu := &api.Mutation{CommitNow: true}
	dgo.DeleteEdges(mu, id, "vehicle")

	res, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		error = true
		log.Fatal(err)
	}

	res, err = dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		error = true
		log.Fatal(err)
	}

	return error
}
