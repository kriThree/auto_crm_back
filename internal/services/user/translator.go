package user_serivce

import (
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s UserService) fromStorageToDomain(user storage_models.User) models.User {

	return models.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
