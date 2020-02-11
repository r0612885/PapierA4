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

	var msg string

	// if err == true {
	// 	msg = "getting users failed"
	// } else{
	// 	msg = "Success"
	// }

	fmt.Println(msg)

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

	user, err := userservice.GetUser(message.ID)

	var msg string

	if err == true {
		msg = "getting users failed"
	} else{
		msg = "Success"
	}

	fmt.Println(msg)

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

	userID, err := userservice.CreateUser(message.Content)

	var msg string

	if err == true {
		msg = "getting users failed"
	} else{
		msg = "Success"
	}

	fmt.Println(msg)

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

	user, err := userservice.UpdateUser(message.ID, message.Content)

	var msg string

	if err == true {
		msg = "getting users failed"
	} else{
		msg = "Success"
	}

	fmt.Println(msg)

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

	err := userservice.DeleteUser(message.ID)

	var msg string

	if err == true {
		msg = "getting users failed"
	} else{
		msg = "Successfull delete of user"
	}

	fmt.Println(msg)

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

	err := userservice.DeleteConnectionBetweenUserAndVehicle(message.ID)

	var msg string

	if err == true {
		msg = "getting users failed"
	} else{
		msg = "Successfull delete between user and vehicle"
	}

	fmt.Println(msg)

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

// CreateConnectionBetweenUserAndVehicleMessageHandler is the event handler for topic "message" and action "createConnectionUserAndVehicle"
func CreateConnectionBetweenUserAndVehicleMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {

	var message models.Message

	mapstructure.Decode(payload, &message)

	user, err := userservice.CreateConnectionBetweenVehicleAndUser(message.ID, message.Content)

	var msg string

	if err == true {
		msg = "getting users failed"
	} else{
		msg = "Success"
	}

	fmt.Println(msg)

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