package user_serivce

import (
	"context"
	"errors"
	"fmt"
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

	log.Info("Start logining user")

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

	role, roleId, err := s.identifyRole(ctx, user.Id)

	if err != nil {
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	accessToken, refreshToken, err := s.usC.Generate(ctx, roleId, role)

	if err != nil {
		log.Error("Generate token error", slog.Any("error", err.Error()))
		return "", "", models.User{}, err
	}
	log.Info("Login user success",
		slog.String("token", accessToken),
		slog.String("refreshToken", refreshToken),
		slog.Any("user", user),
	)

	return accessToken, refreshToken, s.fromStorageToDomain(
		user,
		s.writeUserRoles(roleId, role),
	), nil
}
