package repository

import "app/pkg/models"

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	AddVehicle(newVehicle models.Vehicle) (models.Vehicle, error)
	GetVehicleById(id int) (models.Vehicle, error)
	FindVehiclesByColorAndYear(color string, year int) (v map[int]models.Vehicle)
	FindVehiclesByBrandAndRangeYears(brand string, starYear int, endYear int) (v map[int]models.Vehicle, err error)
	FindVehiclesByBrand(brand string) (v map[int]models.Vehicle, err error)
}
