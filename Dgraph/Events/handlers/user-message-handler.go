package handlers

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/r0612885/PapierA4/Dgraph/Events/mock"
	"github.com/r0612885/PapierA4/Dgraph/Events/models"
	"github.com/r0612885/PapierA4/Dgraph/Services/userservice"
)

// GetUsersMessageHandler is the event handler for topic "message" and action "getUser"
func GetUsersMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	users := userservice.GetUsers()

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": users,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// GetUserMessageHandler is the event handler for topic "message" and action "getUser"
func GetUserMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	user := userservice.GetUser(message.ID)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": user,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// CreateUserMessageHandler is the event handler for topic "message" and action "createUser"
func CreateUserMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	userID := userservice.CreateUser(message.Content)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": userID,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// UpdateUserMessageHandler is the event handler for topic "message" and action "updateUser"
func UpdateUserMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	user := userservice.UpdateUser(message.ID, message.Content)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": user,
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// DeleteUserMessageHandler is the event handler for topic "message" and action "deleteUser"
func DeleteUserMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	userservice.DeleteUser(message.ID)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": "User Deleted!",
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}

// DeleteConnectionBetweenUserAndVehicleMessageHandler is the event handler for topic "message" and action "deleteConnectionUserAndVehicle"
func DeleteConnectionBetweenUserAndVehicleMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {

	var message models.Message

	mapstructure.Decode(payload, &message)

	userservice.DeleteConnectionBetweenUserAndVehicle(message.ID)

	echoNotification := &models.Notification{
		Type:   "global",
		Target: "any",
		Payload: map[string]interface{}{
			"message": "Connection Deleted!",
		},
		Hash: fmt.Sprintf("n:%s", hash),
	}

	var ns mock.NotificationService

	ns.Publish(echoNotification)
}