package autoservice_controller

import (
	rest_models "server_crm/internal/modules/rest/models"
	"server_crm/internal/services/models"
)

func (h AutoserviceController) fromDomainToRest(a models.Autoservice) rest_models.Autoservice {
	return rest_models.Autoservice{
		Id:        a.Id,
		Name:      a.Name,
		Address:   a.Address,
		Phone:     a.Phone,
		Email:     a.Email,
		OwnerId:   a.OwnerId,
		CreatedAt: a.CreatedAt.UTC().String(),
	}
}