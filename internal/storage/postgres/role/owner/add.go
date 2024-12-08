package storage_owner

import (
	"context"
	"fmt"
)

func (s OwnerStorage) Add(ctx context.Context, userId int64) (id int64, err error) {

	const op = "postgres.roles.owner.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO owners (user_id)
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