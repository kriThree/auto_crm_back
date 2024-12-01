package user_controller

import (
	rest_models "server_crm/internal/modules/rest/models"
	"server_crm/internal/services/models"
)

func (uc UserController) fromDomainToRest(user models.User) rest_models.User {

	return rest_models.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

func (uc UserController) fromRestToDomain(user rest_models.User) models.User {

	return models.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}