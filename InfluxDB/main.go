package main

import (
	"fmt"

	"github.com/r0612885/PapierA4/InfluxDB/Services/locationservice"
)

func main() {
	client := locationservice.Init()

	location := locationservice.Location{Uid: "0xDA", Vid: "0xh1", Lat: "50", Lon: "60"}

	metric := locationservice.CreateMockMetric(location)

	locationservice.WriteRow(client, metric)

	lou := locationservice.ReadLastLocationOfUser(client, "0xAA")

	fmt.Printf("Last location of user %v:\n%v\n", "0xAA", lou)

	lov := locationservice.ReadLastLocationOfVehicle(client, "0xh1")

	fmt.Printf("Last location of vehicle %v:\n%v\n", "0xh1", lov)

	lovs := locationservice.ReadLastLocationOfVehicles(client)

	fmt.Printf("Last location of all vehicles:\n%v\n", lovs)

	locationservice.Exit(client)
}
