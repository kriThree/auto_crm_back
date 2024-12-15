package car_service

import (
	"context"
	"log/slog"
	"server_crm/internal/services/models"
)

func (s CarService) Get(ctx context.Context, clientId int64) ([]models.Car, error) {

	const op = "service.car.Get"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("client_id", clientId),
	)

	log.Info("Start getting car")

	storageCars, err := s.crP.GetByClientId(ctx, clientId)
	if err != nil {
		log.Error("Get cars error", slog.Any("error", err.Error()))
		return nil, err
	}

	cars := make([]models.Car, 0, len(storageCars))

	for _, car := range storageCars {
		cars = append(cars, s.fromStorageToDomain(car))
	}

	log.Info("Get cars success",
		slog.Int("count", len(cars)),
	)

	return cars, nil
}


func (s CarService) GetById(ctx context.Context, id int64) (models.Car, error) {

	const op = "service.car.GetById"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start getting car by id")

	storageCar, err := s.crP.GetById(ctx, id)
	if err != nil {
		return models.Car{}, err
	}

	log.Info("Get car by id success",
		slog.Int64("id", id),
	)

	return s.fromStorageToDomain(storageCar), nil
}