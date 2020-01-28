package locationservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go"
)

type Location struct {
	Uid string //User ID
	Vid string //Vehicle ID
	Lat string //Latitude
	Lon string //Longitude
}

func WriteRow(_influx *influxdb.Client, _metric []influxdb.Metric) {
	n, err := _influx.Write(context.Background(), "tracking", "papierA4", _metric...)
	if err != nil {
		log.Fatal(err) // as above use your own error handling here.
	}

	fmt.Println("Created rows: ", n)
}

func CreateMockMetric(_location Location) []influxdb.Metric {
	// we use client.NewRowMetric for the example because it's easy, but if you need extra performance
	// it is fine to manually build the []client.Metric{}.

	metric := []influxdb.Metric{influxdb.NewRowMetric(
		map[string]interface{}{"lat": _location.Lat, "lon": _location.Lon},
		"location",
		map[string]string{"Uid": _location.Uid, "Vid": _location.Vid},
		time.Now()),
	}

	return metric
}

func Exit(_influx *influxdb.Client) {
	_influx.Close()
	fmt.Println("Program finished")
}

func Init() (_influx *influxdb.Client) {
	// You can generate a Token from the "Tokens Tab" in the UI
	_influx, err := influxdb.New("http://localhost:9999", "jyMT01jlRtd8El6HxAAELaTyNjr4yEZyLyjfWdOfOOmlj1j1NosiLEDt3READGawh8MPXg7u7ZGSYSRdrDbrvQ==", influxdb.WithUserAndPass("hark", "welkpassword"))
	if err != nil {
		panic(err) // error handling here; normally we wouldn't use fmt but it works for the example
	}

	return _influx
}

func ReadLastLocationOfUser(_influx *influxdb.Client, _uid string) string {
	stop := time.Now().Format(time.RFC3339)
	start := time.Now().AddDate(0, -1, 0).Format(time.RFC3339)

	var query string = `from(bucket: "tracking")
	|> range(start: ` + start + `, stop: ` + stop + `)
	|> filter(fn: (r) => r._measurement == "location")
	|> filter(fn: (r) => r._field == "lat" or r._field == "lon")
	|> filter(fn: (r) => r.Uid == "` + _uid + `")
	|> last()
	|> yield(name: "last")`

	reader, err := _influx.QueryCSV(
		context.Background(),
		query,
		"papierA4",
	)
	if err != nil {
		panic(err)
	}

	reader.Next()

	m := make(map[string]interface{})
	reader.Unmarshal(m)

	dataLon := fmt.Sprintf("%v", m["_value"])

	dataVid := fmt.Sprintf("%v", m["Vid"])

	reader.Next()

	n := make(map[string]interface{})
	reader.Unmarshal(n)

	dataLat := fmt.Sprintf("%v", n["_value"])

	result, err := json.Marshal(Location{Uid: _uid, Vid: dataVid, Lat: dataLat, Lon: dataLon})
	if err != nil {
		log.Fatal(err)
	}

	return string(result)
}

func ReadLastLocationOfVehicle(_influx *influxdb.Client, _vid string) string {
	stop := time.Now().Format(time.RFC3339)
	start := time.Now().AddDate(0, -1, 0).Format(time.RFC3339)

	var query string = `from(bucket: "tracking")
	|> range(start: ` + start + `, stop: ` + stop + `)
	|> filter(fn: (r) => r._measurement == "location")
	|> filter(fn: (r) => r._field == "lat" or r._field == "lon")
	|> filter(fn: (r) => r.Vid == "` + _vid + `")
	|> last()
	|> yield(name: "last")`

	reader, err := _influx.QueryCSV(
		context.Background(),
		query,
		"papierA4",
	)
	if err != nil {
		panic(err)
	}

	reader.Next()

	m := make(map[string]interface{})
	reader.Unmarshal(m)

	dataLon := fmt.Sprintf("%v", m["_value"])
	dataUid := fmt.Sprintf("%v", m["Uid"])

	reader.Next()

	n := make(map[string]interface{})
	reader.Unmarshal(n)

	dataLat := fmt.Sprintf("%v", n["_value"])

	result, err := json.Marshal(Location{Uid: dataUid, Vid: _vid, Lat: dataLat, Lon: dataLon})
	if err != nil {
		log.Fatal(err)
	}

	return string(result)
}

func ReadLastLocationOfVehicles(_influx *influxdb.Client) string {
	stop := time.Now().Format(time.RFC3339)
	start := time.Now().AddDate(0, -1, 0).Format(time.RFC3339)

	var query string = `from(bucket: "tracking")
	|> range(start: ` + start + `, stop: ` + stop + `)
	|> filter(fn: (r) => r._measurement == "location")
	|> filter(fn: (r) => r._field == "lat" or r._field == "lon")
	|> filter(fn: (r) => r.Vid != "")
	|> drop(columns: ["Uid"])
	|> unique(column: "Vid")
	|> last()`

	reader, err := _influx.QueryCSV(
		context.Background(),
		query,
		"papierA4",
	)
	if err != nil {
		panic(err)
	}

	count := 0
	result := `{ "locations": [`

	vid := [100]string{}
	lat := [100]string{}
	lon := [100]string{}

	for reader.Next() {
		if count%2 == 0 {
			m := make(map[string]interface{})
			reader.Unmarshal(m)
			vid[count] = fmt.Sprintf("%v", m["Vid"])
			lat[count] = fmt.Sprintf("%v", m["_value"])
		} else {
			n := make(map[string]interface{})
			reader.Unmarshal(n)
			lon[count-1] = fmt.Sprintf("%v", n["_value"])

			row, err := json.Marshal(Location{Uid: "", Vid: vid[count-1], Lat: lat[count-1], Lon: lon[count-1]})
			if err != nil {
				log.Panic(err)
			}
			result += string(row)
		}
		count++
	}

	return result + `]}`
}
