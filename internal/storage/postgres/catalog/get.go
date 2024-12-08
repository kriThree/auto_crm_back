package storage_catalog

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s catalogStorage) GetByAdminId(ctx context.Context, adminId int64) ([]storage_models.Catalog, error) {

	const op = "postgres.catalog.Get"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, name, creater_id
		FROM catalogs
		WHERE creater_id = $1;
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx, adminId)

	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var catalogs []storage_models.Catalog
	for rows.Next() {
		var catalog storage_models.Catalog
		err = rows.Scan(&catalog.Id, &catalog.Name, &catalog.CreaterId)

		if err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}

		catalogs = append(catalogs, catalog)
	}

	return catalogs, nil
}

func (s catalogStorage) GetById(ctx context.Context, id int64) (catalog storage_models.Catalog, err error) {

	const op = "postgres.catalog.GetById"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, name, creater_id
		FROM catalogs
		WHERE id = $1;
	`)

	if err != nil {
		return storage_models.Catalog{}, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&catalog.Id, &catalog.Name, &catalog.CreaterId)

	if err != nil {
		return storage_models.Catalog{}, fmt.Errorf("%s : %w", op, err)
	}

	return catalog, nil
}
func (s catalogStorage) Get(ctx context.Context) ([]storage_models.Catalog, error) {

	const op = "postgres.catalog.Get"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, name, creater_id
		FROM catalogs;
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var catalogs []storage_models.Catalog
	for rows.Next() {
		var catalog storage_models.Catalog
		err = rows.Scan(&catalog.Id, &catalog.Name, &catalog.CreaterId)

		if err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}

		catalogs = append(catalogs, catalog)
	}

	return catalogs, nil
}
