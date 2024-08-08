package main

type VehicleType string
type FuelType string

const (
	Gasoline FuelType = "Gasoline"
	Diesel   FuelType = "Diesel"
)

const (
	Motorcycle VehicleType = "Motorcycle"
	Car        VehicleType = "Car"
	Truck      VehicleType = "Truck"
	Van        VehicleType = "Van"
)

type Maker struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	ID      string `json:"id"`
	MakerID string `json:"makerId"`
	Name    string `json:"name"`
	Year    int    `json:"year"`
}

type Vehicle struct {
	ID                 string      `json:"id"`
	ModelID            string      `json:"modelId"`
	VehicleType        VehicleType `json:"vehicleType"`
	PlateNo            string      `json:"plateNo"`
	FuelType           FuelType    `json:"fuelType"`
	Capacity           int         `json:"capacity"`
	DepartmentAssigned string      `json:"deptAssigned"`
}
