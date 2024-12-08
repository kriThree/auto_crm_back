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
