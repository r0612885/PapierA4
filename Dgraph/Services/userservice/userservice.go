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
	Uid          string  `json:"uid,omitempty"`
	Type         string  `json:"type,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Needsservice bool    `json:"needsservice"`
}

type User struct {
	Uid       string  `json:"uid,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname  string  `json:"lastname,omitempty"`
	Email     string  `json:"email,omitempty"`
	Password  string  `json:"password,omitempty"`
	Role      string  `json:"role,omitempty"`
	Vehicle   Vehicle `json:"vehicle,omitempty"`
}

// GetUsers gets all users
func GetUsers() []byte {

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
		users(func: has(firstname)) {
	      uid
		  firstname
		  lastname
		  email
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
		log.Fatal(err)
	}

	return res.Json
}

// GetActiveUsers gets all users who are currently operating a vehicle
func GetActiveUsers() []byte {

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
		activeUsers(func: has(vehicle)) {
			uid
			firstname
			lastname
			email
			role
		  		vehicle {
					uid
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
		log.Fatal(err)
	}

	return res.Json
}

// GetUser gets a user
func GetUser(id string) []byte {

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
		user(func: uid($id)){
			uid
			firstname
			lastname
			email
			role
		  		vehicle {
					uid
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
		log.Fatal(err)
	}

	return res.Json
}

// CreateUser creates a user
func CreateUser(u []byte) []byte {

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
		SetJson:   []byte(u),
	}

	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		fmt.Print(err)
	}

	variables := map[string]string{"$id": assigned.Uids["user"]}
	q := `query getUser($id: string){
		user(func: uid($id)){
			firstname
			lastname
			email
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
		fmt.Print(err)
	}

	return res.Json
}

// CreateConnectionBetweenVehicleAndUser connection between a user and a vehicle
func CreateConnectionBetweenVehicleAndUser(vehicleID string, userID string) []byte {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
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
			uid
			firstname
			lastname
			role
		  		vehicle {
					uid
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
		log.Fatal(err)
	}

	return res.Json
}

// UpdateUser updates a user
func UpdateUser(id string, u []byte) []byte {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
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
			firstname
			lastname
			role
		}
	}`

	res, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		log.Fatal(err)
	}

	return res.Json
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
		user(func: uid($id)){
			firstname
			lastname
			email
			role
		}
	}`

	mu := &api.Mutation{CommitNow: true}
	dgo.DeleteEdges(mu, id, "name", "role", "vehicle")

	res, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	res, err = dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
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
		user(func: uid($id)){
			firstname
			lastname
			email
			role
		}
	}`

	mu := &api.Mutation{CommitNow: true}
	dgo.DeleteEdges(mu, id, "vehicle")

	res, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	res, err = dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		log.Fatal(err)
	}
}

///////////////////////
////authentication/////
///////////////////////

func Authenticate(email string, password string) []byte {

	conn, err := grpc.Dial("192.168.99.100:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	variables := map[string]string{"$email": email, "$password": password}
	q := `query getUser($email : string, $password : string) {
		user(func: eq(email, $email)) @filter(eq(password, $password)){
			uid
			firstname
			lastname
			email
			role
		  		vehicle {
					uid
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
		log.Fatal(err)
	}

	return res.Json
}
