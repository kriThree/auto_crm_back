package storage_work

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s workStorage) Update(ctx context.Context, dto storage_models.UpdateWorkDto) error {

	const op = "postgres.work.Update"

	stmt, err := s.db.PrepareContext(ctx, `
		UPDATE works
		SET cost = $1, name = $2, catalog_id=$3
		WHERE id = $4;
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, dto.Cost, dto.Name, dto.CatalogId, dto.Id)

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
