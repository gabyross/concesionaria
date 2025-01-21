package service

import "app/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	AddVehicle(vehicleDoc models.VehicleDoc) (models.Vehicle, error)
	FindVehiclesByColorAndYear(color string, year string) (vehicles map[int]models.Vehicle, err error)
	FindVehiclesByBrandAndRangeYears(brand string, starYear int, endYear int) (v map[int]models.Vehicle, err error)
	FindAverageOfSpeedByBrand(brand string) (average float64, err error)
	AddMultipleVehicles(v []models.VehicleDoc) (err error)
	UpdateMaxSpeed(id int, newSpeed float64) (err error)
	GetVehicleById(id int) (models.Vehicle, error)
	FindVehiclesByFuel(fuel string) (v map[int]models.Vehicle, err error)
	DeleteVehicle(id int) (err error)
	FindVehiclesByTransmission(transmisiion string) (v map[int]models.Vehicle, err error)
}


