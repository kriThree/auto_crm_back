package user_serivce

import (
	"context"
	"errors"
	"log/slog"
	services_errors "server_crm/internal/services/errors"
	"server_crm/internal/services/models"
	storage_errors "server_crm/internal/storage/errors"

	"golang.org/x/crypto/bcrypt"
)

func (s UserService) Login(ctx context.Context, email string, cryptedPassword string) (string, string, models.User, error) {

	const op = "service.user.Login"

	log := s.l.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	user, err := s.usP.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, storage_errors.ErrUserNotFound) {
			log.Error("User not found", slog.Any("error", err.Error()))
			return "", "", models.User{}, services_errors.ErrIncorrectPassword
		}
		log.Error("Find user error", slog.Any("error", err.Error()))
		return "", "", models.User{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cryptedPassword)); err != nil {
		log.Error("Password compare error", slog.Any("error", err.Error()))
		return "", "", models.User{}, services_errors.ErrIncorrectPassword
	}

	token, refreshToken, err := s.usC.Generate(ctx, user.Id)
	if err != nil {
		log.Error("Generate token error", slog.Any("error", err.Error()))
		return "", "", models.User{}, err
	}

	return token, refreshToken, s.fromStorageToDomain(user), nil
}
