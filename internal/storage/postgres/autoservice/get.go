package storage_autoservice

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s AutoserviceStorage) GetByOwnerId(ctx context.Context, userId int64) ([]storage_models.Autoservice, error) {

	const op = "postgres.autoservice.Get"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, name, address, phone, owner_id, email
		FROM autoservices
		WHERE owner_id = $1;
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx, userId)

	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var autoservices []storage_models.Autoservice
	for rows.Next() {
		var autoservice storage_models.Autoservice
		err = rows.Scan(&autoservice.Id, &autoservice.Name, &autoservice.Address, &autoservice.Phone, &autoservice.OwnerId, &autoservice.Email)

		if err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}

		autoservices = append(autoservices, autoservice)
	}

	return autoservices, nil
}

func (s AutoserviceStorage) GetById(ctx context.Context, id int64) (autoservice storage_models.Autoservice, err error) {

	const op = "postgres.autoservice.GetById"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, name, address, phone, owner_id, email
		FROM autoservices
		WHERE id = $1;
	`)

	if err != nil {
		return storage_models.Autoservice{}, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&autoservice.Id, &autoservice.Name, &autoservice.Address, &autoservice.Phone, &autoservice.OwnerId, &autoservice.Email)

	if err != nil {
		return storage_models.Autoservice{}, fmt.Errorf("%s : %w", op, err)
	}

	return autoservice, nil
}
