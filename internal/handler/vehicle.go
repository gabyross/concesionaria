package handler

import (
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv service.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv service.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)
		for key, value := range v {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// create a vehicle and add it to the map
func (h *VehicleDefault) AddVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read the body
		body := r.Body
		vehicle := models.VehicleDoc{}
		err := json.NewDecoder(body).Decode(&vehicle)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err = h.sv.AddVehicle(vehicle)
		if err != nil {
			if err.Error() == "Identificador del vehículo ya existente" {
				response.JSON(w, http.StatusConflict, err.Error())
			} else if err.Error() == "Campos incompletos o mal formados" {
				response.JSON(w, http.StatusBadRequest, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusCreated, "Vehículo creado exitosamente")
	}
}

// FindVehiclesByColorAndYear is a method that returns vehicles filtered by color and year
func (h *VehicleDefault) FindVehiclesByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get parameters /vehicles/color/{color}/year/{year}
		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")

		// - get vehicles filtered by color and year
		v, err := h.sv.FindVehiclesByColorAndYear(color, year)
		if err != nil {
			// specify error
			if err.Error() == "No se encontraron vehículos con esos criterios" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)

		for key, value := range v {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) FindVhehiclesByBrandAndRangeYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		startYear, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Eror al convertir año de inicio")
		}

		endYear, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Eror al convertir año de finalización")
		}

		vehicles, err := h.sv.FindVehiclesByBrandAndRangeYears(brand, startYear, endYear)

		if err != nil {
			if err.Error() == "No se encontraron vehículos con esos criterios" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusOK, vehicles)
	}
}

func (h *VehicleDefault) FindAverageOfSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		average, err := h.sv.FindAverageOfSpeedByBrand(brand)

		if err != nil {
			if err.Error() == "No se encontraron vehículos de esa marca" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		response.JSON(w, http.StatusOK, average)
	}
}

func (h *VehicleDefault) AddMultipleVehicles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		var vehicles []models.VehicleDoc
		err := json.NewDecoder(body).Decode(&vehicles)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = h.sv.AddMultipleVehicles(vehicles)
		if err != nil {
			if err.Error() == "Algún vehículo tiene un identificador ya existente" {
				response.JSON(w, http.StatusConflict, err.Error())
			} else if err.Error() == "Datos de algún vehículo mal formados o incompletos" {
				response.JSON(w, http.StatusBadRequest, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusCreated, "Vehículos creados exitosamente")
	}
}

func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
		}

		body := r.Body

		var vehicleDoc models.VehicleDoc
		err = json.NewDecoder(body).Decode(&vehicleDoc)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
		}

		err = h.sv.UpdateMaxSpeed(id, vehicleDoc.MaxSpeed)
		if err != nil {
			if err.Error() == "Velocidad mal formada o fuera de rango." {
				response.JSON(w, http.StatusBadRequest, err.Error())
			} else if err.Error() == "Vehicle not found" {
				response.JSON(w, http.StatusNotFound, "No se encontró el vehículo")
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		response.JSON(w, http.StatusOK, "Velocidad del vehículo actualizada exitosamente")
	}
}

func (h *VehicleDefault) GetVehicleById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
		}

		vehicle, err := h.sv.GetVehicleById(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
		}

		response.JSON(w, http.StatusOK, vehicle)
	}
}

func (h *VehicleDefault) FindVehiclesByFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuel := chi.URLParam(r, "type")

		vehicles, err := h.sv.FindVehiclesByFuel(fuel)
		if err != nil {
			if err.Error() == "No se encontraron vehículos con ese tipo de combustible" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		response.JSON(w, http.StatusOK, vehicles)
	}
}

func (h *VehicleDefault) DeleteVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
		}

		err = h.sv.DeleteVehicle(id)
		if err != nil {
			if err.Error() == "Vehicle not found" {
				response.JSON(w, http.StatusNotFound, "No se encontró el vehículo")
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusNoContent, map[string]string{"message": "Vehículo eliminado exitosamente"})
	}
}

func (h *VehicleDefault) FindVehiclesBytransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmisiion := chi.URLParam(r, "type")

		vehicles, err := h.sv.FindVehiclesByTransmission(transmisiion)
		if err != nil {
			if err.Error() == "No se encontraron vehículos con ese tipo de transmisión" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusOK, vehicles)
	}
}

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
		}

		var vehicleDoc models.VehicleDoc
		body := r.Body
		err = json.NewDecoder(body).Decode(&vehicleDoc)

		err = h.sv.UpdateFuel(id, vehicleDoc)
		if err != nil {
			if err.Error() == "Vehicle not found" {
				response.JSON(w, http.StatusNotFound, "No se encontró el vehículo")
			} else if err.Error() == "Tipo de combustible mal formado o no admitido" {
				response.JSON(w, http.StatusBadRequest, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusOK, "Tipo de combustible del vehículo actualizado exitosamente")
	}

}

func (h *VehicleDefault) GetAveragePeopleCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		average, err := h.sv.GetAveragePeopleCapacityByBrand(brand)
		if err != nil {
			if err.Error() == "No se encontraron vehículos de esa marca" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusOK, average)
	}
}

func (h *VehicleDefault) FindVehiclesByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		length := r.URL.Query().Get("length")
		width := r.URL.Query().Get("width")

		lengthRangeArr := strings.Split(length, "-")
		widthRangeArr := strings.Split(width, "-")

		if len(lengthRangeArr) != 2 || len(widthRangeArr) != 2 {
			response.JSON(w, http.StatusBadRequest, "Rango de longitud o ancho mal formado")
			return
		}

		minLength, err := strconv.ParseFloat(lengthRangeArr[0], 64)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		maxLength, err := strconv.ParseFloat(lengthRangeArr[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Longitud máxima inválida")
			return
		}

		minWidth, err := strconv.ParseFloat(widthRangeArr[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Ancho mínimo inválido")
			return
		}

		maxWidth, err := strconv.ParseFloat(widthRangeArr[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Ancho máximo inválido")
			return
		}

		vehicles, err := h.sv.FindVehiclesByDimensions(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			if err.Error() == "No se encontraron vehículos con esas dimensiones" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusOK, vehicles)
	}
}

func (h *VehicleDefault) FindVehiclesByWeigth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		min := r.URL.Query().Get("min")
		max := r.URL.Query().Get("max")

		if min == "" || max == "" {
			response.JSON(w, http.StatusBadRequest, "Parámetros 'min' y 'max' son requeridos")
			return
		}

		minWeigth, err := strconv.ParseFloat(min, 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "peso minimo invalido")
		}

		maxWeigth, err := strconv.ParseFloat(max, 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "peso maximo invalido")
		}

		vehicles, err := h.sv.FindVehiclesByWeigth(minWeigth, maxWeigth)
		if err != nil {
			if err.Error() == "No se encontraron vehículos en ese rango de peso" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else {
				response.JSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		response.JSON(w, http.StatusOK, vehicles)
	}
}
