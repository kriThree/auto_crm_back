package storage_autoservice

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s AutoserviceStorage) Add(ctx context.Context, dto storage_models.AddAutoserviceDto, userId int64) (int64, error) {

	const op = "postgres.user.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO users (name, address, phone, owner_id, email)
		VALUES ($1, $2, $3, $4, $5)
		Returning id;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, dto.Name, dto.Address, dto.Phone, userId, dto.Email).Scan(&resId)

	if err != nil {
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
