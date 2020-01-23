package main

import (
	"fmt"

	"github.com/r0612885/PapierA4/Dgraph/Events/handlers"
	"github.com/r0612885/PapierA4/Dgraph/Events/mock"
	"github.com/r0612885/PapierA4/Dgraph/Events/models"
)

func main() {

	mockEvents := []models.Event{

		// USER EVENTS

		// models.Event{
		// 	Topic:  "message",
		// 	Action: "getUsers",
		// 	Payload: map[string]interface{}{
		// 	},
		// 	Hash: "GetUsers",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "getUser",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x5",
		// 	},
		// 	Hash: "GetUser",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "createUser",
		// 	Payload: map[string]interface{}{
		// 		"content": `{ "user": [{"uid": "_:user","name": "Sponge","role": "Admin"}]}`,
		// 	},
		// 	Hash: "CreateUser",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "updateUser",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x68",
		// 		"content": `{ "user": [{"uid": "0x68","name": "Bob Squarepants","role": "Admin"}]}`,
		// 	},
		// 	Hash: "UpdateUser",
		// },
		models.Event{
			Topic:  "message",
			Action: "deleteUser",
			Payload: map[string]interface{}{
				"id": "0x2d",
			},
			Hash: "DeleteUser",
		},
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "deleteConnectionUserAndVehicle",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x5",
		// 	},
		// 	Hash: "DeleteConnectionUserAndVehicle",
		// },

		// VEHICLE EVENTS

		// models.Event{
		// 	Topic:  "message",
		// 	Action: "getVehicles",
		// 	Payload: map[string]interface{}{
		// 	},
		// 	Hash: "GetVehicles",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "getVehicle",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x65",
		// 	},
		// 	Hash: "GetVehicle",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "createVehicle",
		// 	Payload: map[string]interface{}{
		// 		"content": `{ "vehicle": [{"uid": "_:vehicle","type": "Vrachtwagen ABC123","latitude": 415.2366,"longitude": 41.97791,"needsservice": false}]}`,
		// 	},
		// 	Hash: "CreateVehicle",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "updateVehicle",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x65",
		// 		"content": `{ "vehicle": [{"uid": "0x65","type": "Vrachtwagen DEF456","needsservice": true}]}`,
		// 	},
		// 	Hash: "UpdateVehicle",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "deleteVehicle",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x65",
		// 	},
		// 	Hash: "DeleteVehicle",
		// },
	}

	eventService := mock.EventService{
		MockedQueue:  mockEvents,
		EventChannel: make(chan models.Event),
	}

	// USER SUBSCRIBES

	// eventService.Subscribe("message", "getUsers", handlers.GetUsersMessageHandler)
	// eventService.Subscribe("message", "getUser", handlers.GetUserMessageHandler)
	// eventService.Subscribe("message", "createUser", handlers.CreateUserMessageHandler)
	// eventService.Subscribe("message", "updateUser", handlers.UpdateUserMessageHandler)
	// eventService.Subscribe("message", "deleteUser", handlers.DeleteUserMessageHandler)
	// eventService.Subscribe("message", "deleteConnectionUserAndVehicle", handlers.DeleteConnectionBetweenUserAndVehicleMessageHandler) NOG TESTEN!!!

	// VEHICLE SUBSCRIBES

	// eventService.Subscribe("message", "getVehicles", handlers.GetVehiclesMessageHandler)
	// eventService.Subscribe("message", "getVehicle", handlers.GetVehicleMessageHandler)
	// eventService.Subscribe("message", "createVehicle", handlers.CreateVehicleMessageHandler)
	// eventService.Subscribe("message", "updateVehicle", handlers.UpdateVehicleMessageHandler)
	// eventService.Subscribe("message", "deleteVehicle", handlers.DeleteVehicleMessageHandler)

	err := eventService.ListenForEvents()

	if err != nil {
		fmt.Println(err)
	}
}
