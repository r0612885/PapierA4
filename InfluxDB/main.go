package main

import (
	"github.com/r0612885/PapierA4/InfluxDB/Services/locationservice"
)

func main() {
	client := locationservice.Init()

	location := locationservice.Location{Uid: "0xa1", Vid: "0xh1", Lat: "50", Lon: "60"}

	metrics := locationservice.CreateMockMetric(location)

	locationservice.WriteRow(client, metrics)

	locationservice.Exit(client)

}
