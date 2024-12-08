package storage_work

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s workStorage) Get(ctx context.Context) ([]storage_models.Work, error) {

	const op = "postgres.work.Get"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, cost,name,catalog_id
		FROM works;
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	var works []storage_models.Work
	for rows.Next() {
		var work storage_models.Work
		err = rows.Scan(&work.Id, &work.Cost, &work.Name, &work.CatalogId)

		if err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}

		works = append(works, work)
	}

	return works, nil
}

func (s workStorage) GetById(ctx context.Context, id int64) (work storage_models.Work, err error) {

	const op = "postgres.work.GetById"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, cost, name, catalog_id
		FROM works
		WHERE id = $1;
	`)
	if err != nil {
		return storage_models.Work{}, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&work.Id, &work.Cost, &work.Name, &work.CatalogId)

	if err != nil {
		return storage_models.Work{}, fmt.Errorf("%s : %w", op, err)
	}

	return work, nil
}
func (s workStorage) GetByCatalogId(ctx context.Context, catalogId int64) (work []storage_models.Work, err error) {

	const op = "postgres.work.GetByCatalogId"

	works, err := s.getByQuery(ctx, op, "catalog_id = $1", catalogId)

	if err != nil {
		return []storage_models.Work{}, fmt.Errorf("%s: %w", op, err)
	}


	return works, nil
}
func (s workStorage) getByQuery(ctx context.Context, op string, query string, args ...any) ([]storage_models.Work, error) {

	stmt, err := s.db.PrepareContext(ctx, fmt.Sprintf(
		`SELECT id, cost, name, catalog_id
		FROM operations
		WHERE %s `, query))

	if err != nil {
		return []storage_models.Work{}, err
	}
	rows, err := stmt.QueryContext(ctx, args)
	if err != nil {
		return []storage_models.Work{}, err
	}

	var works []storage_models.Work
	for rows.Next() {
		var work storage_models.Work
		err = rows.Scan(&work.Id, &work.Cost, &work.Name, &work.CatalogId)

		if err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}

		works = append(works, work)
	}
	return works, nil
}
