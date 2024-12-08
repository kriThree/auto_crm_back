package operation_service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s OperationService) Delete(ctx context.Context, id int64) error {

	const op = "service.operation.Delete"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start to delete operation")

	err := s.opP.Delete(ctx, id)

	if err != nil {
		log.Error("Delete operation error", slog.Any("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Delete operation success")

	return nil
}
