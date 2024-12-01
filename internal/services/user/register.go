package user_serivce

import (
	"context"
	"fmt"
	"log/slog"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s UserService) Register(ctx context.Context, dto models.RegisterUserDto) (string, string, models.User, error) {

	const op = "service.user.Register"

	log := s.l.With(
		slog.String("op", op),
		slog.String("email", dto.Email),
		slog.String("password", dto.Password),
		slog.String("name", dto.Name),
	)

	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error("Hashing password error", slog.Any("error", err.Error()))
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	time := time.Now()
	id, err := s.usP.Add(ctx, storage_models.AddUserDto{
		Name:       dto.Name,
		Email:      dto.Email,
		Password:   string(password),
		Created_at: time,
	})

	if err != nil {
		log.Error("Register user error", slog.Any("error", err.Error()))
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	accessToken, refreshToken, err := s.usC.Generate(ctx, id)

	if err != nil {
		log.Error("Generate token error", slog.Any("error", err.Error()))
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return accessToken, refreshToken, models.User{
		Id:        id,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  string(password),
		CreatedAt: time,
	}, nil

}
