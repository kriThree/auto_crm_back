package storage_client

import (
	"context"
	"fmt"
)

func (s ClientStorage) Add(ctx context.Context, userId int64 ) (id int64, err error) {

	const op = "postgres.roles.client.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO clients (user_id)
		VALUES ($1)
		Returning id;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, userId).Scan(&resId)

	if err != nil {
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
