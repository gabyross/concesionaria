package service

import "app/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	AddVehicle(vehicleDoc models.VehicleDoc) (models.Vehicle, error)
	FindVehiclesByColorAndYear(color string, year string) (vehicles map[int]models.Vehicle, err error)
}
