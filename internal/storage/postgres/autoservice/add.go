package storage_autoservice

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s AutoserviceStorage) Add(ctx context.Context, dto storage_models.AddAutoserviceDto) (int64, error) {

	const op = "postgres.autoservice.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO autoservices (name, address, phone, email, owner_id)
		VALUES ($1, $2, $3, $4, $5)
		Returning id;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, dto.Name, dto.Address, dto.Phone, dto.Email, dto.Owner_id).Scan(&resId)

	if err != nil {
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
