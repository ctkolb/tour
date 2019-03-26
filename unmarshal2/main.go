package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type VehicleJson struct {
	Name string
	Description string
}

type Vehicle struct {
	Name string
	Description string
	CurrentSpeed int
}

func ConvertVehicleList(vehicleJsons []VehicleJson) []Vehicle {
	vehicles := []Vehicle{}

	for _, v := range vehicleJsons {
		vehicle := Vehicle{Name: v.Name, Description: v.Description}
		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

func main() {
	//read in our vehicles here
	vehiclesJsonFile, err := ioutil.ReadFile("vehicles.json") // just pass the file name
    if err != nil {
        panic(err)
	}
	var vehicleJsons []VehicleJson
	err = json.Unmarshal(vehiclesJsonFile, &vehicleJsons)
	if err != nil {
        panic(err)
	}
	fmt.Println("VehicleJsons loaded: ", len(vehicleJsons)) 

	vehicles := ConvertVehicleList(vehicleJsons)
	fmt.Println("Vehicles converted from json: ", len(vehicles))
}