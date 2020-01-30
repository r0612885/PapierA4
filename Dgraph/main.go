package main

import (
	"fmt"

	// "fmt"

	"github.com/r0612885/PapierA4/Dgraph/Events/handlers"
	"github.com/r0612885/PapierA4/Dgraph/Events/mock"
	"github.com/r0612885/PapierA4/Dgraph/Events/models"
	// "github.com/r0612885/PapierA4/Dgraph/Services/authservice"
)

func main() {
	mockEvents := []models.Event{

		// // USER EVENTS

		// models.Event{
		// 	Topic:   "message",
		// 	Action:  "getUsers",
		// 	Payload: map[string]interface{}{},
		// 	Hash:    "GetUsers",
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
		// 		"content": `{ "user": [{"uid": "_:user","name": "Sponge", "password": "snicker","role": "Admin"}]}`,
		// 	},
		// 	Hash: "CreateUser",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "createConnectionUserAndVehicle",
		// 	Payload: map[string]interface{}{
		// 		"id":      "0x3",
		// 		"content": "0x68",
		// 	},
		// 	Hash: "CreateConnectionUserAndVehicle",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "updateUser",
		// 	Payload: map[string]interface{}{
		// 		"id":      "0x68",
		// 		"content": `{ "user": [{"uid": "0x68","role": "test"}]}`,
		// 	},
		// 	Hash: "UpdateUser",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "deleteUser",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x2d",
		// 	},
		// 	Hash: "DeleteUser",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "deleteConnectionUserAndVehicle",
		// 	Payload: map[string]interface{}{
		// 		"id": "0x5",
		// 	},
		// 	Hash: "DeleteConnectionUserAndVehicle",
		// },

		// // VEHICLE EVENTS

		// models.Event{
		// 	Topic:   "message",
		// 	Action:  "getVehicles",
		// 	Payload: map[string]interface{}{},
		// 	Hash:    "GetVehicles",
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
		// 		"id":      "0x65",
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

		// SERVICE EVENTS

		// models.Event{
		// 	Topic:  "message",
		// 	Action: "createService",
		// 	Payload: map[string]interface{}{
		// 		"id":      "0x17",
		// 		"content": `{"uid": "_:service","description": "Motorpech","datecompleted": "0"}`,
		// 	},
		// 	Hash: "CreateService",
		// },
		// models.Event{
		// 	Topic:  "message",
		// 	Action: "completeService",
		// 	Payload: map[string]interface{}{
		// 		"id":      "0x5f",
		// 		"content": `{ "service": [{"uid": "0x5f","datecompleted": "` + strconv.Itoa(int(time.Now().Unix())) + `"}]}`,
		// 	},
		// 	Hash: "CompleteService",
		// },
		models.Event{
			Topic:  "message",
			Action: "getLastService",
			Payload: map[string]interface{}{
				"id": "0x17",
			},
			Hash: "GetLastService",
		},
	}

	eventService := mock.EventService{
		MockedQueue:  mockEvents,
		EventChannel: make(chan models.Event),
	}

	// // USER SUBSCRIBES

	// eventService.Subscribe("message", "getUsers", handlers.GetUsersMessageHandler)
	// eventService.Subscribe("message", "getUser", handlers.GetUserMessageHandler)
	// eventService.Subscribe("message", "createUser", handlers.CreateUserMessageHandler)
	// eventService.Subscribe("message", "createConnectionUserAndVehicle", handlers.CreateConnectionBetweenUserAndVehicleMessageHandler)
	// eventService.Subscribe("message", "updateUser", handlers.UpdateUserMessageHandler)
	// eventService.Subscribe("message", "deleteUser", handlers.DeleteUserMessageHandler)
	// eventService.Subscribe("message", "deleteConnectionUserAndVehicle", handlers.DeleteConnectionBetweenUserAndVehicleMessageHandler)

	// // VEHICLE SUBSCRIBES

	// eventService.Subscribe("message", "getVehicles", handlers.GetVehiclesMessageHandler)
	// eventService.Subscribe("message", "getVehicle", handlers.GetVehicleMessageHandler)
	// eventService.Subscribe("message", "createVehicle", handlers.CreateVehicleMessageHandler)
	// eventService.Subscribe("message", "updateVehicle", handlers.UpdateVehicleMessageHandler)
	// eventService.Subscribe("message", "deleteVehicle", handlers.DeleteVehicleMessageHandler)

	// SERVICE SUBSCRIBES

	// eventService.Subscribe("message", "createService", handlers.CreateServiceMessageHandler)
	// eventService.Subscribe("message", "completeService", handlers.CompleteServiceMessageHandler)
	eventService.Subscribe("message", "getLastService", handlers.GetTimeSinceLastServiceMessageHandler)

	err := eventService.ListenForEvents()

	if err != nil {
		fmt.Println(err)
	}

	// content := `{"username":"test","password": "test123"}`

	// token := authservice.Login(content)
	// fmt.Println(token)
}
