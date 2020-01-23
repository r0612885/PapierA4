package main

import (
	"fmt"

	"github.com/r0612885/PapierA4/InfluxDB/Services/locationservice"
)

func main() {
	client := locationservice.Init()

	location := locationservice.Location{Uid: "0xa1", Vid: "0xh1", Lat: "50", Lon: "60"}

	metric := locationservice.CreateMockMetric(location)

	locationservice.WriteRow(client, metric)

	res := locationservice.ReadLocationOfUser(client, "0xa1")

	fmt.Println(res)

	locationservice.Exit(client)

}
