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
	// check mandatory fields
	ok := areMandatoryFieldsOK(newVehicle)
	if !ok {
		return models.Vehicle{}, errors.New("Campos incompletos o mal formados")
	}

	// check if the vehicle (id) already exists
	_, err := r.GetVehicleById(newVehicle.Id)

	// if vehicle does not exists in the db
	if err != nil {
		// add new vehicle to db and return it
		r.db[newVehicle.Id] = newVehicle
		return newVehicle, nil
	}

	// in case the vehicle already exists, return an error
	return models.Vehicle{}, errors.New("Identificador del vehículo ya existente")
}

func (r *VehicleMap) GetVehicleById(id int) (models.Vehicle, error) {
	// try to get the vehicle form the db by its key
	vehicle, ok := r.db[id]

	// check if the vehicle was found
	if !ok {
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

func areMandatoryFieldsOK(vehicle models.Vehicle) bool {
	if vehicle.Id == 0 ||
		vehicle.Brand == "" ||
		vehicle.Model == "" ||
		vehicle.Registration == "" ||
		vehicle.Color == "" ||
		vehicle.FabricationYear == 0 ||
		vehicle.Capacity == 0 ||
		vehicle.MaxSpeed == 0 ||
		vehicle.FuelType == "" ||
		vehicle.Transmission == "" ||
		vehicle.Weight == 0 ||
		vehicle.Height == 0 ||
		vehicle.Length == 0 ||
		vehicle.Width == 0 {
		return false
	}
	return true
}
