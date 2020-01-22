package main

import (
	"github.com/r0612885/PapierA4/Dgraph/Services/userservice"
	"github.com/r0612885/PapierA4/Dgraph/Services/vehicleservice"
)

func main() {

	u := userservice.User{
		Uid:  "_:user",
		Name: "Din Vanwezemael",
		Role: "Admin",
	}

	v := vehicleservice.Vehicle{
		Uid:          "_:vehicle",
		Type:         "Vrachtwagen A978",
		Latitude:     41.1551,
		Longitude:    49.1255,
		Needsservice: false,
	}

	userservice.CreateUser(u)
	// userservice.UpdateUser("0x42", u)
	// userservice.DeleteVehicle("0x42")
	// userservice.GetUser("0x42")
	// userservice.GetAllService()
	// userservice.DeleteConnectionBetweenUserAndVehicle("0x42")

	vehicleservice.CreateVehicle(v)
	// vehicleservice.UpdateVehicle("0x3b", v)
	// vehicleservice.GetAllVehicles()
	// vehicleservice.GetVehicle("0x3b")
	// vehicleservice.DeleteVehicle("0x3b")

}
