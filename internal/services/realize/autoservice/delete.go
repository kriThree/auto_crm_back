package autoservice_service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s AutoserviceService) Delete(ctx context.Context, id int64) error {

	const op = "service.autoservice.Delete"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start to delete autoservice")
	
	err := s.auP.Delete(ctx, id)

	if err != nil {
		log.Error("Delete autoservice error", slog.Any("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Delete autoservice success")

	return nil
}