package storage_catalog

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s catalogStorage) Update(ctx context.Context, dto storage_models.UpdateCatalogDto) error {

	const op = "postgres.catalog.Update"

	stmt, err := s.db.PrepareContext(ctx, `
		UPDATE catalogs
		SET name = $1
		WHERE id = $2;
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, dto.Name, dto.Id)

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
