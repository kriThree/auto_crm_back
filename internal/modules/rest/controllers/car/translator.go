package car_controller

import (
	rest_models "server_crm/internal/modules/rest/models"
	"server_crm/internal/services/models"
)

func (h CarController) fromDomainToRest(car models.Car) rest_models.Car {
	return rest_models.Car{
		Id:    car.Id,
		Number: car.Number,
		Description: car.Description,
	}
}
