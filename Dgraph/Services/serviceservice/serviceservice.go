package serviceservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

// type SubResult struct {
// 	DateCompleted string `json:"dateCompleted,omitempty"`
// }

// type Result struct {
// 	SubResult SubResult `json:"service,omitempty"`
// }

// type QueryResult struct {
// 	Result Result `json:"getVehicleLastService,omitempty"`
// }

type ID struct {
	Tijd string `json:""`
}

type Service struct {
	Uid           string `json:"uid,omitempty"`
	DateCompleted string `json:"datecompleted,omitempty"`
	Description   string `json:"description,omitempty"`
}

type Vehicle struct {
	Uid          string  `json:"uid,omitempty"`
	Type         string  `json:"type,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Needsservice bool    `json:"needsservice,omitempty"`
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

func CompleteService(s string, id string) (string, bool) {
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

	mu := &api.Mutation{SetJson: []byte(s)}
	req.Mutations = []*api.Mutation{mu}

	if _, err := dg.NewTxn().Do(ctx, req); err != nil {
		log.Fatal(err)
	}

	variables := map[string]string{"$id": id}
	q := `query getService($id: string){
		service(func: uid($id)) {
			dateCompleted
		}
	}`

	res, err := dg.NewReadOnlyTxn().QueryWithVars(ctx, q, variables)
	if err != nil || res == nil {
		error = true
		log.Fatal(err)
	}

	return string(res.Json), error
}

func GetTimeSinceLastService(id string) string {
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
		{vehicle(func: uid(` + id + `)){
			uid
			service(first: -1){
				datecompleted
				description
			  }
		  }}`

	res, err := txn.Query(ctx, q)

	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]interface{})
	json.Unmarshal(res.Json, &result)

	data := result["vehicle"].([]interface{})

	vehicle := data[0].(map[string]interface{})

	data2 := vehicle["service"].([]interface{})

	service := data2[0].(map[string]interface{})

	i := fmt.Sprint("%v", service["datecompleted"])

	x, _ := strconv.ParseInt(i[2:], 10, 64)

	t := time.Unix(x, 0)

	r := time.Since(t)

	fmt.Println(r)

	form, _ := time.ParseDuration("12h")
	rounded := r.Round(form)

	fmt.Println(rounded)

	days := rounded.Hours() / 24

	if days > 1.0 {
		return strconv.Itoa(int(days))
	} else {
		tmp, _ := time.ParseDuration("30m")
		tmp2 := r.Round(tmp)
		fmt.Println(tmp2)
		return strconv.Itoa(int(tmp2))
	}

}
