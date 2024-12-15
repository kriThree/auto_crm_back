package work_controller

import (
	rest_models "server_crm/internal/modules/rest/models"
	"server_crm/internal/services/models"
)

func (h WorkController) fromDomainToRest(work models.Work) rest_models.Work {
	return rest_models.Work{
		Id:    work.Id,
		Name:  work.Name,
		Cost:  work.Cost,
	}
}
