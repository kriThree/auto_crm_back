package user_serivce

import (
	"context"
	"fmt"
	"log/slog"
	"server_crm/internal/auxiliary"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s UserService) Get(ctx context.Context) ([]models.User, error) {

	const op = "service.user.Get"

	log := s.l.With(
		slog.String("op", op),
	)

	log.Info("Get users")

	storageUsers, err := s.usP.Get(ctx)

	if err != nil {
		log.Error("Get users error", slog.Any("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	users, err := auxiliary.NewWorker(
		ctx,
		storageUsers,
		func(aCtx auxiliary.Context, storageUser storage_models.User) models.User {
			user, err := s.GetOne(context.Background(), storageUser.Id)
			if err != nil {
				aCtx.PushError(err)
				return models.User{}
			}
			return user
		})

	if err != nil {
		log.Error("Get users error", slog.Any("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return users, nil
}
