package work_service

import (
	"context"
	"log/slog"
	"server_crm/internal/services/models"
)

func (s WorkService) Get(ctx context.Context) ([]models.Work, error) {

	const op = "service.work.Get"

	log := s.l.With(
		slog.String("op", op),
	)

	log.Info("Start getting work")

	storageWorks, err := s.wkP.Get(ctx)
	if err != nil {
		log.Error("Get works error", slog.Any("error", err.Error()))
		return nil, err
	}

	works := make([]models.Work, 0, len(storageWorks))

	for _, work := range storageWorks {
		works = append(works, s.fromStorageToDomain(work))
	}

	log.Info("Get works success",
		slog.Int("count", len(works)),
	)

	return works, nil
}
func (s WorkService) GetById(ctx context.Context, id int64) (models.Work, error) {

	const op = "service.work.GetById"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start getting work by catalog id")

	storageWork, err := s.wkP.GetById(ctx, id)

	if err != nil {
		log.Error("Get works by catalog id error", slog.Int64("catalog_id", id), slog.Any("error", err.Error()))
		return models.Work{}, err
	}

	log.Info("Get works by catalog id success")

	return s.fromStorageToDomain(storageWork), nil
}
