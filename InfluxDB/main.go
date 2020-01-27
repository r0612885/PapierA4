package main

import (
	"fmt"

	influxdb "github.com/influxdata/influxdb-client-go"
	"github.com/r0612885/PapierA4/InfluxDB/Services/locationservice"
)

func main() {
	client := locationservice.Init()

	location := locationservice.Location{Uid: "0xDA", Vid: "0xh1", Lat: "50", Lon: "60"}

	metric := locationservice.CreateMockMetric(location)

	locationservice.WriteRow(client, metric)

	res := locationservice.ReadLocationOfUser(client, "0xAA")

	fmt.Println(res)

	var csvRes influxdb.QueryCSVResult

	_interface := map[string]interface{}{
		"row": csvRes.Row,
		// "colnames": colnames,
	}

	err := res.Unmarshal(_interface)
	if err != nil {
		panic(err)
	}

	fmt.Println(csvRes.Row)
	// fmt.Println(colnames)

	locationservice.Exit(client)

}
