package handlers

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/r0612885/PapierA4/Dgraph/Events/mock"
	"github.com/r0612885/PapierA4/Dgraph/Events/models"
	"github.com/r0612885/PapierA4/Dgraph/Services/serviceservice"
)

// CreateServiceMessageHandler is the event handler for topic "message" and action "createService"
func CreateServiceMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	res := serviceservice.CreateService(message.Content, message.ID)

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

// CreateServiceMessageHandler is the event handler for topic "message" and action "createService"
func CompleteServiceMessageHandler(payload map[string]interface{}, hash string, metadata map[string]string) {
	var message models.Message

	mapstructure.Decode(payload, &message)

	var msg string

	res, err := serviceservice.CompleteService(message.ID)
	if err == true {
		msg = "completingservice failed"
	} else {
		msg = "Success"
	}

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
