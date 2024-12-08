package autoservice_service

import (
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s AutoserviceService) fromStorageToDomain(user storage_models.Autoservice) models.Autoservice {

	return models.Autoservice{
		Id:        user.Id,
		Name:      user.Name,
		Address:   user.Address,
		Phone:     user.Phone,
		Email:     user.Email,
		OwnerId:   user.OwnerId,
		CreatedAt: user.CreatedAt,
	}
}
