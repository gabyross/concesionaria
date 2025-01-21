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
	UpdateMaxSpeed(id int, newSpeed float64) (err error)
	FindVehiclesByFuel(fuel string) (v map[int]models.Vehicle)
	DeleteVehicle(id int) (err error)
	FindVehiclesByTransmission(transmisiion string) (v map[int]models.Vehicle)
	UpdateFuel(id int, newFuel string) (err error)
	GetVehiclesByBrand(brand string) (v map[int]models.Vehicle)
}
