package user_serivce

import (
	"context"
	"errors"
	services_errors "server_crm/internal/services/errors"
	"server_crm/internal/services/models"
	storage_errors "server_crm/internal/storage/errors"
	storage_models "server_crm/internal/storage/models"
)

func (s UserService) addRoleForUser(ctx context.Context, userId int64, role string) (int64, error) {

	roleId := int64(0)
	err := error(nil)

	if role == storage_models.ROLE_ADMIN {
		roleId, err = s.rsP.Admin.Add(ctx, userId)

	} else if role == storage_models.ROLE_OWNER {
		roleId, err = s.rsP.Owner.Add(ctx, userId)

	} else if role == storage_models.ROLE_CLIENT {
		roleId, err = s.rsP.Client.Add(ctx, userId)

	} else {
		return 0, services_errors.ErrRoleNotFound
	}

	return roleId, err
}
func (s UserService) identifyRole(ctx context.Context, userId int64) (role string, roleId int64, err error) {

	roleId, err = s.rsP.Admin.GetByUserId(ctx, userId)

	if err != nil && !errors.Is(err, storage_errors.ErrAdminRoleNotFound) {
		return "", 0, err
	} else if err == nil {
		return storage_models.ROLE_ADMIN, roleId, nil
	}

	roleId, err = s.rsP.Owner.GetByUserId(ctx, userId)

	if err != nil && !errors.Is(err, storage_errors.ErrOwnerRoleNotFound) {
		return "", 0, err
	} else if err == nil {
		return storage_models.ROLE_OWNER, roleId, nil
	}

	roleId, err = s.rsP.Client.GetByUserId(ctx, userId)

	if err != nil && !errors.Is(err, storage_errors.ErrClientRoletNotFound) {
		return "", 0, err
	} else if err == nil {
		return storage_models.ROLE_CLIENT, roleId, nil
	}

	return "", 0, services_errors.ErrRoleNotFound
}
func (s UserService) writeUserRoles(roleId int64, role string) models.UserRoles {

	if role == storage_models.ROLE_ADMIN {
		return models.UserRoles{Admin: roleId}
	} else if role == storage_models.ROLE_OWNER {
		return models.UserRoles{Owner: roleId}
	} else if role == storage_models.ROLE_CLIENT {
		return models.UserRoles{Client: roleId}
	}

	return models.UserRoles{}
}
