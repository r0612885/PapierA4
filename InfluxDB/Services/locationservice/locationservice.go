package locationservice

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

func ReadLocationOfUser(_influx *influxdb.Client, _uid string) *influxdb.QueryCSVResult {
	var timeStop = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	var tmp = time.Now().AddDate(0, -1, 0)
	var timeStart = strconv.FormatInt(tmp.UTC().UnixNano(), 10)

	var query string = `from(bucket: "tracking")  |> range(start: ` + timeStart + `, stop: ` + timeStop + `)  |> filter(fn: (r) => r._measurement == "location")  |> filter(fn: (r) => r.Uid == "` + _uid + `")  |> last()`

	fmt.Println(timeStart)
	fmt.Println(timeStop)

	result, err := _influx.QueryCSV(
		context.Background(),
		query,
		"papierA4",
	)

	fmt.Println(result)

	if err != nil {
		panic(err)
	}
	return result
}
