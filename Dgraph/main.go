package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r0612885/PapierA4/Dgraph/Services/userservice"
	"github.com/r0612885/PapierA4/Dgraph/Services/vehicleservice"
)

type User struct {
	Uid       string  `json:"uid,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname  string  `json:"lastname,omitempty"`
	Email     string  `json:"email,omitempty"`
	Password  string  `json:"password,omitempty"`
	Role      string  `json:"role,omitempty"`
	Vehicle   Vehicle `json:"vehicle,omitempty"`
}

type Vehicle struct {
	Uid          string  `json:"uid,omitempty"`
	Type         string  `json:"type,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Needsservice bool    `json:"needsservice"`
}

////////////////////////////////
/////////////user///////////////
////////////////////////////////

//get all the users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := userservice.GetUsers()

	w.Write(users)
}

//get all the active users
func getActiveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := userservice.GetActiveUsers()

	w.Write(users)
}

//get user -- pass userID
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	users := userservice.GetUser(params["id"])

	w.Write(users)
}

//create new user
func createUser(w http.ResponseWriter, r *http.Request) {

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg User
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")

	userservice.CreateUser(output)

	fmt.Print(output)
	w.Write(output)
}

//create a new connecten between vehicle and user
func createConnectionBetweenVehicleAndUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	users := userservice.CreateConnectionBetweenVehicleAndUser(params["vehicleID"], params["userID"])

	w.Write(users)
}

//update the user -- pass userID with the new data
func updateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg User
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")

	userservice.UpdateUser(params["id"], output)

	fmt.Print(output)
	w.Write(output)
}

//delete user -- pass userID
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	userservice.DeleteUser(params["id"])
}

//delete connection between user and vehicle -- pass the userID
func deleteConnectionBetweenUserAndVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	userservice.DeleteConnectionBetweenUserAndVehicle(params["id"])
}

//delete vehicle -- pass vehicleID
func authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	fmt.Print(params["email"])
	fmt.Print(params["password"])

	user := userservice.Authenticate(params["email"], params["password"])

	w.Write(user)
}

////////////////////////////////
/////////////vehicles///////////
////////////////////////////////

//get all the vehicles
func getVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vehicles := vehicleservice.GetVehicles()

	w.Write(vehicles)
}

// //get all the active vehicles
// func getActiveVehicles(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	vehicles := vehicleservice.GetActiveVehicles()

// 	fmt.Println("request get users")

// 	w.Write(vehicles)
// }

//get vehicle -- pass vehicleID
func getVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	vehicle := vehicleservice.GetVehicle(params["id"])

	w.Write(vehicle)
}

//create the vehicle -- pass the data
func createVehicle(w http.ResponseWriter, r *http.Request) {

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Vehicle
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")

	vehicleservice.CreateVehicle(output)

	w.Write(output)
}

//update the vehicle -- pass vehicleID with the new data
func updateVehicle(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Vehicle
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")

	vehicleservice.UpdateVehicle(params["id"], output)

	w.Write(output)
}

//delete vehicle -- pass vehicleID
func deleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	vehicleservice.DeleteVehicle(params["id"])
}

func main() {
	r := mux.NewRouter()

	////////////
	//user api//
	////////////
	r.HandleFunc("/user/all", getUsers).Methods("GET")
	r.HandleFunc("/user/active", getActiveUsers).Methods("GET")
	r.HandleFunc("/user/get/{id}", getUser).Methods("GET")
	r.HandleFunc("/user/create", createUser).Methods("POST")
	//pass vehicleID -- userID
	r.HandleFunc("/connection/create/{vehicleID}/{userID}", createConnectionBetweenVehicleAndUser).Methods("PUT")
	r.HandleFunc("/user/update/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/delete/{id}", deleteUser).Methods("DELETE")
	//pass userID
	r.HandleFunc("/connection/delete/{id}", deleteConnectionBetweenUserAndVehicle).Methods("DELETE")
	r.HandleFunc("/user/authenticate/{email}/{password}", authenticate).Methods("GET")

	///////////////
	//vehicle api//
	///////////////
	r.HandleFunc("/vehicle/all", getVehicles).Methods("GET")
	// r.HandleFunc("/vehicle/getallactive", getActiveVehicles).Methods("GET")
	r.HandleFunc("/vehicle/get/{id}", getVehicle).Methods("GET")
	r.HandleFunc("/vehicle/create", createVehicle).Methods("POST")
	r.HandleFunc("/vehicle/update/{id}", updateVehicle).Methods("PUT")
	r.HandleFunc("/vehicle/delete/{id}", deleteVehicle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
