package storage_role

import (
	"database/sql"
	storage_models "server_crm/internal/storage/models"
	storage_admin "server_crm/internal/storage/postgres/role/admin"
	storage_client "server_crm/internal/storage/postgres/role/client"
	storage_owner "server_crm/internal/storage/postgres/role/owner"
)

type RolesStorage struct {
	Admin  storage_models.RoleDomain
	Client storage_models.RoleDomain
	Owner  storage_models.RoleDomain
}

func New(db *sql.DB) RolesStorage {
	return RolesStorage{
		Admin:  storage_admin.New(db),
		Client: storage_client.New(db),
		Owner:  storage_owner.New(db),
	}
}
