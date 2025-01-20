package service

import (
	"app/internal/repository"
	"app/pkg/models"
	"errors"

	"strconv"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp repository.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp repository.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) AddVehicle(vehicleDoc models.VehicleDoc) (models.Vehicle, error) {
	// convert vehicleDoc to vehicle
	newVehicle := mapDocToVehicle(vehicleDoc)

	vehicle, err := s.rp.AddVehicle(newVehicle)

	// set error based on resulted string
	if err != nil {
		if err.Error() == "Campos incompletos o mal formados" {
			return models.Vehicle{}, err
		}
		if err.Error() == "Identificador del vehículo ya existente" {
			return models.Vehicle{}, err
		}

		return models.Vehicle{}, err
	}
	return vehicle, nil
}

// FindVehiclesByColorAndYear implements VehicleService.
func (s *VehicleDefault) FindVehiclesByColorAndYear(color string, year string) (vehicles map[int]models.Vehicle, err error) {
	yearParsed, err := strconv.Atoi(year)
	if err != nil {
		return make(map[int]models.Vehicle), err
	}

	vehicles = s.rp.FindVehiclesByColorAndYear(color, yearParsed)
	if len(vehicles) == 0 {
		return make(map[int]models.Vehicle), errors.New("No se encontraron vehículos con esos criterios")
	}
	return vehicles, nil
}

func (s *VehicleDefault) FindVehiclesByBrandAndRangeYears(brand string, starYear int, endYear int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindVehiclesByBrandAndRangeYears(brand, starYear, endYear)

	if err != nil {
		return make(map[int]models.Vehicle), err
	}

	if len(v) == 0 {
		return v, errors.New("No se encontraron vehículos con esos criterios")
	}

	return v, nil
}

func (s *VehicleDefault) FindAverageOfSpeedByBrand(brand string) (average float64, err error) {
	vehicles, err := s.rp.FindVehiclesByBrand(brand)
	if err != nil {
		return 0, err
	}

	for key := range vehicles {
		average += vehicles[key].MaxSpeed
	}

	if average != 0 {
		return average / float64(len(vehicles)), nil
	}

	return 0, errors.New("No se encontraron vehículos de esa marca")
}

func mapDocToVehicle(doc models.VehicleDoc) models.Vehicle {
	vehicle := models.Vehicle{
		Id: doc.ID,
		VehicleAttributes: models.VehicleAttributes{
			Brand:           doc.Brand,
			Model:           doc.Model,
			Registration:    doc.Registration,
			Color:           doc.Color,
			FabricationYear: doc.FabricationYear,
			Capacity:        doc.Capacity,
			MaxSpeed:        doc.MaxSpeed,
			FuelType:        doc.FuelType,
			Transmission:    doc.Transmission,
			Weight:          doc.Weight,
			Dimensions: models.Dimensions{
				Height: doc.Height,
				Length: doc.Length,
				Width:  doc.Width,
			},
		},
	}
	return vehicle
}
