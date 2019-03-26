package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type VehicleConstant struct {
	Name string
	Description string
}

type Vehicle struct {
	Properties VehicleConstant
	CurrentSpeed int
	CurrentDirection string
}

func ConvertVehicleList(vehicleJsons []VehicleConstant) []Vehicle {
	vehicles := []Vehicle{}

	for _, v := range vehicleJsons {
		vehicle := Vehicle{Properties: v, CurrentSpeed: 0, CurrentDirection: "N"}
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
	var vehicleConstants []VehicleConstant
	err = json.Unmarshal(vehiclesJsonFile, &vehicleConstants)
	if err != nil {
        panic(err)
	}
	fmt.Println("VehicleJsons loaded: ", len(vehicleConstants)) 

	//convert from the constants provided in the json over to our struct that has the dynamic fields
	vehicles := ConvertVehicleList(vehicleConstants)
	fmt.Println("Vehicles converted from json: ", len(vehicles))

	//list out the details of the vehicles
	for i, v := range vehicles {
		fmt.Println("Array ID: ", i, "\nVehicle Name: ", v.Properties.Name, "\nVehicle Description: ", v.Properties.Description)
		fmt.Println("Vehicle Speed: ", v.CurrentSpeed, "\nVehicle Direction: ", v.CurrentDirection)
	}
}