package handlers

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/r0612885/PapierA4/Dgraph/Events/mock"
	"github.com/r0612885/PapierA4/Dgraph/Events/models"
	"github.com/r0612885/PapierA4/Dgraph/Services/vehicleservice"
)

// GetVehiclesMessageHandler is the event handler for topic "message" and action "getUser"
func GetVehiclesMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	vehicles := vehicleservice.GetVehicles()

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": vehicles,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// GetVehicleMessageHandler is the event handler for topic "message" and action "getUser"
func GetVehicleMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	vehicle := vehicleservice.GetVehicle(message.ID)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": vehicle,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// CreateVehicleMessageHandler is the event handler for topic "message" and action "createUser"
func CreateVehicleMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	vehicleID := vehicleservice.CreateVehicle(message.Content)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": vehicleID,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// UpdateVehicleMessageHandler is the event handler for topic "message" and action "createUser"
func UpdateVehicleMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	vehicle := vehicleservice.UpdateVehicle(message.ID, message.Content)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": vehicle,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// DeleteVehicleMessageHandler is the event handler for topic "message" and action "deleteUser"
func DeleteVehicleMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	vehicleservice.DeleteVehicle(message.ID)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": "Vehicle Deleted!",
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}