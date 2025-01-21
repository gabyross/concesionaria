package repository

import (
	"app/pkg/models"
	"errors"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]models.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]models.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]models.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) AddVehicle(newVehicle models.Vehicle) (models.Vehicle, error) {
	r.db[newVehicle.Id] = newVehicle
	return newVehicle, nil
}

func (r *VehicleMap) GetVehicleById(id int) (models.Vehicle, error) {
	// try to get the vehicle form the db by its key
	vehicle, exists := r.db[id]

	// check if the vehicle was found
	if !exists {
		return models.Vehicle{}, errors.New("Vehicle not found")
	}

	// return vehicle without an error
	return vehicle, nil
}

// FindVehiclesByColorAndYear implements VehicleRepository.
func (r *VehicleMap) FindVehiclesByColorAndYear(color string, year int) (v map[int]models.Vehicle) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		vehicle := r.db[key]

		// if match, copy vehicle
		if strings.ToLower(vehicle.Color) == strings.ToLower(color) && vehicle.FabricationYear == year {
			v[key] = value
		}
	}
	return v
}

func (r *VehicleMap) FindVehiclesByBrandAndRangeYears(brand string, starYear int, endYear int) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)
	for key, value := range r.db {
		vehicle := r.db[key]
		if strings.EqualFold(vehicle.Brand, brand) && (vehicle.FabricationYear >= starYear && vehicle.FabricationYear <= endYear) {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindVehiclesByBrand(brand string) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)
	for key, value := range r.db {
		vehicle := r.db[key]
		if strings.EqualFold(vehicle.Brand, brand) {
			v[key] = value
		}
	}
	return
}
