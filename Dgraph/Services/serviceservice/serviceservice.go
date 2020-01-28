package serviceservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"time"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Service struct {
	Uid           string    `json:"uid,omitempty"`
	DateCompleted time.Time `json:"datecompleted,omitempty"`
	Description   string    `json:"description,omitempty"`
}

type Vehicle struct {
	Uid          string  `json:"uid,omitempty"`
	Type         string  `json:"type,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Needsservice bool    `json:"needsservice"`
	Service      Service `json:"service,omitempty"`
}

// CreateService creates a service
func CreateService(s string, vehicleID string) string {

	var service Service
	json.Unmarshal([]byte(s), &service)

	fmt.Println(service)

	if service.Description == "" {
		panic("paniiieeeekk")
	}

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
		SetJson:   []byte(s),
	}

	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	variables := map[string]string{"$id": assigned.Uids["service"]}
	q := `query getService($id: string){
					service(func: uid($id)) {
						dateCompleted
						description
				}

	}`

	res, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		log.Fatal(err)
	}

	pb, err := json.Marshal(Vehicle{Uid: vehicleID, Service: Service{Uid: assigned.Uids["service"]}})
	if err != nil {
		log.Fatal(err)
	}

	req := &api.Request{CommitNow: true}

	mu = &api.Mutation{SetJson: pb}
	mu.SetJson = pb
	req.Mutations = []*api.Mutation{mu}

	res, err = dg.NewTxn().Do(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	return string(res.Json)
}
