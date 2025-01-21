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

	// check mandatory fields
	fieldsAreOk := areMandatoryFieldsOK(newVehicle)
	if !fieldsAreOk {
		return models.Vehicle{}, errors.New("Campos incompletos o mal formados")
	}

	// check if the vehicle (id) already exists
	_, err := s.rp.GetVehicleById(newVehicle.Id)

	// if vehicle does not exists in the db
	if err != nil {
		// add new vehicle to db and return it
		_, err = s.rp.AddVehicle(newVehicle)
		if err != nil {
			return models.Vehicle{}, err
		}
		return newVehicle, nil

	} else {
		return models.Vehicle{}, errors.New("Identificador del vehículo ya existente")
	}
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

func (s *VehicleDefault) AddMultipleVehicles(v []models.VehicleDoc) (err error) {
	for _, vehicle := range v {
		newVehicle := mapDocToVehicle(vehicle)
		// check mandatory fields
		fieldsAreOk := areMandatoryFieldsOK(newVehicle)
		if !fieldsAreOk {
			return errors.New("Datos de algún vehículo mal formados o incompletos")
		}

		// check if the vehicle (id) already exists
		_, err := s.rp.GetVehicleById(newVehicle.Id)

		// if vehicle exists in the db
		if err == nil {
			return errors.New("Algún vehículo tiene un identificador ya existente")
		} else {
			// add new vehicle to db and return it
			_, err = s.rp.AddVehicle(newVehicle)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *VehicleDefault) UpdateMaxSpeed(id int, newSpeed float64) (err error) {
	if newSpeed <= 0 {
		return errors.New("Velocidad mal formada o fuera de rango.")
	}

	_, err = s.rp.GetVehicleById(id)
	if err != nil {
		return err
	}

	err = s.rp.UpdateMaxSpeed(id, newSpeed)
	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleDefault) GetVehicleById(id int) (models.Vehicle, error) {
	vehicle, err := s.rp.GetVehicleById(id)
	if err != nil {
		return models.Vehicle{}, err
	}
	return vehicle, nil
}

func (s *VehicleDefault) FindVehiclesByFuel(fuel string) (map[int]models.Vehicle, error) {
	v := s.rp.FindVehiclesByFuel(fuel)
	if len(v) == 0 {
		return make(map[int]models.Vehicle), errors.New("No se encontraron vehículos con ese tipo de combustible")
	}

	return v, nil
}

func (s *VehicleDefault) DeleteVehicle(id int) (err error) {
	_, err = s.rp.GetVehicleById(id)
	if err != nil {
		return err
	}

	err = s.rp.DeleteVehicle(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *VehicleDefault) FindVehiclesByTransmission(transmisiion string) (v map[int]models.Vehicle, err error) {
	v = s.rp.FindVehiclesByTransmission(transmisiion)
	if len(v) == 0 {
		return v, errors.New("No se encontraron vehículos con ese tipo de transmisión")
	}

	return v, nil
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
