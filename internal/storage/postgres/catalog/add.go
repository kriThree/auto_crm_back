package storage_catalog

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s catalogStorage) Add(ctx context.Context, dto storage_models.AddCatalogDto) (int64, error) {

	const op = "postgres.catalog.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO catalogs (name, creater_id)
		VALUES ($1, $2)
		Returning id;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, dto.Name, dto.AdminId).Scan(&resId)

	if err != nil {
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
