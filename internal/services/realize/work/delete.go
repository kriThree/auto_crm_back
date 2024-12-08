package work_service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s WorkService) Delete(ctx context.Context, id int64) error {

	const op = "service.work.Delete"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start to delete work")

	err := s.wkP.Delete(ctx, id)

	if err != nil {
		log.Error("Delete work error", slog.Any("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Delete work success")

	return nil
}
