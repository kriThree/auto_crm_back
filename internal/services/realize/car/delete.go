package car_service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s CarService) Delete(ctx context.Context, id int64) error {
	
	const op = "service.car.Delete"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start to delete car")
	
	err := s.crP.Delete(ctx, id)

	if err != nil {
		log.Error("Delete car error", slog.Any("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Delete car success")
	
	return nil
}
