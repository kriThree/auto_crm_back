package operation_controller

import (
	rest_models "server_crm/internal/modules/rest/models"
	"server_crm/internal/services/models"
)

func (h OperationController) fromDomainToRest(
	operation models.Operation,
) rest_models.Operation {
	return rest_models.Operation{
		Id:          operation.Id,
		Desctiption: operation.Description,
		Car: rest_models.Car{
			Id:          operation.Car.Id,
			Number:      operation.Car.Number,
			Description: operation.Car.Description,
		},
		Work: rest_models.Work{
			Id:        operation.Work.Id,
			Name:      operation.Work.Name,
			Cost:      operation.Work.Cost,
		},
		Autoservice: rest_models.Autoservice{
			Id:        operation.Autoservice.Id,
			Name:      operation.Autoservice.Name,
			Address:   operation.Autoservice.Address,
			Phone:     operation.Autoservice.Phone,
			Email:     operation.Autoservice.Email,
			OwnerId:   operation.Autoservice.OwnerId,
			CreatedAt: operation.Autoservice.CreatedAt.UTC().String(),
		},
	}
}
