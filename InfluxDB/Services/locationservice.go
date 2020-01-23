package main

import(
	"fmt"
	"time"
	"log"
	"context"

	"github.com/influxdata/influxdb-client-go"
)

type Location struct {
	Uid				string		//User ID
	Vid				string		//Vehicle ID
	lat				string		//Latitude
	lon				string		//Longitude
}

func WriteRow(_influx influxdb.Client, _metrics []influxdb.Metric){
	n, err := _influx.Write(context.Background(), "tracking", "papierA4", _metrics...)
	if err != nil {
		log.Fatal(err) // as above use your own error handling here.
	}

	fmt.Println("Created rows: ", n)
}

func CreateMockMetrics(_locations []Location) ([]influxdb.Metric) {
	// we use client.NewRowMetric for the example because it's easy, but if you need extra performance
	// it is fine to manually build the []client.Metric{}.
	result := [len(_locations)]influxdb.Metric

	for index, _location := range _locations{
		result[index] := influxdb.NewRowMetric(
			map[string]interface{}{"lat": _location.lat, "lon": _location.lon},
			"location",
			map[string]string{"Uid": _location.Uid, "Vid": _location.Vid},
			time.Now())
	}

	return result
}

func Exit(_influx influxdb.Client){
	_influx.Close()
	fmt.Println("Program finished")
}

func Init() (_influx influxdb.Client){
	// You can generate a Token from the "Tokens Tab" in the UI
	influx, err := _influxdb.New("http://localhost:9999", "jyMT01jlRtd8El6HxAAELaTyNjr4yEZyLyjfWdOfOOmlj1j1NosiLEDt3READGawh8MPXg7u7ZGSYSRdrDbrvQ==", influxdb.WithUserAndPass("hark","welkpassword"))
	if err != nil {
	panic(err) // error handling here; normally we wouldn't use fmt but it works for the example
	}
}