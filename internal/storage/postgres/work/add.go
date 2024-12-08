package storage_work

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s workStorage) Add(ctx context.Context, dto storage_models.AddWorkDto) (id int64, err error) {

	const op = "postgres.work.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO works (cost, name, catalog_id)
		VALUES ($1, $2, $3)
		Returning id;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, dto.Cost, dto.Name, dto.CatalogId).Scan(&resId)

	if err != nil {
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
