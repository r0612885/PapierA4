package main

import(
	"github.com/r0612885/PapierA4/InfluxDB/Services/locationservice"
)

func main() {
	client := locationservice.Init()

	locations := []locationservice.Location{
		"0xa1","0xh1","50","60",
	}

	metrics := locationservice.CreateMockMetrics(locations)

	locationservice.WriteRow(client, metrics)

	locationservice.Exit(client)

}