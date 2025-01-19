package service

import (
	"app/internal/repository"
	"app/pkg/models"
	"fmt"
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
			return models.Vehicle{}, fmt.Errorf("400 Bad Request: %s", err.Error())
		}
		if err.Error() == "Identificador del vehículo ya existente" {
			return models.Vehicle{}, fmt.Errorf("409 Conflict: %s", err.Error())
		}

		return models.Vehicle{}, err
	}
	return vehicle, nil
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
