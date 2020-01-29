package main

import (
	"log"

	"github.com/r0612885/PapierA4/InfluxDB/Events/handlers"
	"github.com/r0612885/PapierA4/InfluxDB/Events/mock"
	"github.com/r0612885/PapierA4/InfluxDB/Events/models"
)

func main() {
	mockEvents := []models.Event{

		models.Event{
			Topic:  "message",
			Action: "writeRow",
			Payload: map[string]interface{}{
				"content": `{"Uid": "0xFF","Vid": "1x23","Lat": "51.6594", "Lon": "45.2397"}`,
			},
			Hash: "WriteRow",
		},
		models.Event{
			Topic:  "message",
			Action: "getLastUserLocation",
			Payload: map[string]interface{}{
				"id": "0xAA",
			},
			Hash: "GetLastUserLocation",
		},
		models.Event{
			Topic:  "message",
			Action: "getLastVehicleLocation",
			Payload: map[string]interface{}{
				"id": "0x02",
			},
			Hash: "GetLastVehicleLocation",
		},
		models.Event{
			Topic:   "message",
			Action:  "getAllVehiclesLastLocation",
			Payload: map[string]interface{}{},
			Hash:    "GetAllVehiclesLastLocation",
		},
	}

	eventService := mock.EventService{
		MockedQueue:  mockEvents,
		EventChannel: make(chan models.Event),
	}

	eventService.Subscribe("message", "writeRow", handlers.WriteRowMessageHandler)
	eventService.Subscribe("message", "getLastUserLocation", handlers.GetLastLocationOfUserMessageHandler)
	eventService.Subscribe("message", "getLastVehicleLocation", handlers.GetLastLocationOfVehicleMessageHandler)
	eventService.Subscribe("message", "getAllVehiclesLastLocation", handlers.GetAllVehiclesLastLocationMessageHandler)

	err := eventService.ListenForEvents()
	if err != nil {
		log.Fatal(err)
	}
}
