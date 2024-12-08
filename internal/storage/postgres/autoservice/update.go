package storage_autoservice

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)
func (s AutoserviceStorage) Update(ctx context.Context, dto storage_models.UpdateAutoserviceDto) error {

	const op = "postgres.autoservice.Update"

	stmt, err := s.db.PrepareContext(ctx, `
		UPDATE autoservices
		SET name = $1, address = $2, phone = $3, email = $4
		WHERE id = $5;
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, dto.Name, dto.Address, dto.Phone, dto.Email, dto.Id)

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	

	return nil
}