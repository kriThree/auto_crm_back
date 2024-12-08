package autoservice_service

import (
	"context"
	"log/slog"
	"server_crm/internal/services/models"
)

func (s AutoserviceService) Get(ctx context.Context, userId int64) ([]models.Autoservice, error) {

	const op = "service.autoservice.Get"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("user_id", userId),
	)

	log.Info("Start getting autoservices")

	storageAutoservices, err := s.auP.GetByOwnerId(ctx, userId)
	if err != nil {
		return nil, err
	}

	autoservices := make([]models.Autoservice, 0, len(storageAutoservices))

	for _, autoservice := range storageAutoservices {
		autoservices = append(autoservices, s.fromStorageToDomain(autoservice))
	}

	log.Info("Get autoservices success",
		slog.Int("count", len(autoservices)),
	)

	return autoservices, nil
}
func (s AutoserviceService) GetById(ctx context.Context, id int64) (models.Autoservice, error) {
	return models.Autoservice{}, nil
}
