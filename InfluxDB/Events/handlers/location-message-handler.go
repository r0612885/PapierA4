package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/r0612885/PapierA4/InfluxDB/Events/mock"
	"github.com/r0612885/PapierA4/InfluxDB/Events/models"
	"github.com/r0612885/PapierA4/InfluxDB/Services/locationservice"
)

func WriteRowMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	client := locationservice.Init()

	var m locationservice.Location
	json.Unmarshal([]byte(message.Content), &m)

	fmt.Println(m)

	metric := locationservice.CreateMetric(m)

	locationservice.WriteRow(client, metric)

	locationservice.Exit(client)

	var msg string

	fmt.Println(msg)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": "Location Saved!!",
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

func GetLastLocationOfUserMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	client := locationservice.Init()

	id := message.ID

	res := locationservice.GetLastLocationOfUser(client, id)

	locationservice.Exit(client)

	var msg string

	fmt.Println(msg)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": res,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

func GetLastLocationOfVehicleMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	client := locationservice.Init()

	id := message.ID

	res := locationservice.GetLastLocationOfVehicle(client, id)

	locationservice.Exit(client)

	var msg string

	fmt.Println(msg)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": res,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

func GetAllVehiclesLastLocationMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	client := locationservice.Init()

	res := locationservice.GetLastLocationOfVehicles(client)

	locationservice.Exit(client)

	var msg string

	fmt.Println(msg)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": res,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}
