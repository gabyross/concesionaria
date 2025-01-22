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

func (r *VehicleMap) UpdateMaxSpeed(id int, newSpeed float64) (err error) {
	vehicle, exists := r.db[id]
	if !exists {
		return errors.New("Vehicle not found") // Si el vehículo no existe, devuelve un error
	}
	vehicle.MaxSpeed = newSpeed
	r.db[id] = vehicle
	return
}

func (r *VehicleMap) FindVehiclesByFuel(fuel string) (v map[int]models.Vehicle) {
	v = make(map[int]models.Vehicle)
	for key, value := range r.db {
		vehicle := r.db[key]
		if strings.EqualFold(vehicle.FuelType, fuel) {
			v[key] = value
		}
	}
	return v
}

func (r *VehicleMap) DeleteVehicle(id int) (err error) {
	_, exists := r.db[id]
	if !exists {
		return errors.New("Vehicle not found")
	}

	delete(r.db, id)
	return nil
}

func (r *VehicleMap) FindVehiclesByTransmission(transmisiion string) (v map[int]models.Vehicle) {
	v = make(map[int]models.Vehicle)
	for key, value := range r.db {
		vehicle := r.db[key]
		if strings.EqualFold(vehicle.Transmission, transmisiion) {
			v[key] = value
		}
	}
	return v
}

func (r *VehicleMap) UpdateFuel(id int, newFuel string) (err error) {
	vehicle, exists := r.db[id]
	if !exists {
		return errors.New("Vehicle not found")
	}

	vehicle.FuelType = newFuel
	r.db[id] = vehicle
	return nil
}

func (r *VehicleMap) GetVehiclesByBrand(brand string) (v map[int]models.Vehicle) {
	v = make(map[int]models.Vehicle)
	for key, value := range r.db {
		vehicle := r.db[key]
		if strings.EqualFold(vehicle.Brand, brand) {
			v[key] = value
		}
	}
	return v
}

func (r *VehicleMap) FindVehiclesByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) map[int]models.Vehicle {
	vehicles := make(map[int]models.Vehicle)
	for key, value := range r.db {
		vehicle := r.db[key]
		if vehicle.Length >= minLength && vehicle.Length <= maxLength && vehicle.Width >= minWidth && vehicle.Width <= maxWidth {
			vehicles[key] = value
		}
	}
	return vehicles
}

func (r *VehicleMap) FindVehiclesByWeigth(minWeigth float64, maxWeigth float64) map[int]models.Vehicle {
	vehicles := make(map[int]models.Vehicle)
	for key, value := range r.db {
		vehicle := r.db[key]
		if vehicle.Weight >= minWeigth && vehicle.Weight <= maxWeigth {
			vehicles[key] = value
		}
	}
	return vehicles
}
