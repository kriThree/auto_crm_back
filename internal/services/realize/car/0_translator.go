package car_service

import (
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s CarService) fromStorageToDomain(storage_car storage_models.Car) models.Car {
	return models.Car{
		Id:          storage_car.Id,
		Number:      storage_car.Number,
		Description: storage_car.Description,
		ClientId:    storage_car.ClientId,
	}
}