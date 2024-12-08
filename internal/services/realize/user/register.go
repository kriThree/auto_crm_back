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
		slog.String("name", dto.Name),
	)

	log.Info("Start registering user")

	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error("Hashing password error", slog.Any("error", err.Error()))
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	time := time.Now()

	userId, err := s.usP.Add(ctx, storage_models.AddUserDto{
		Name:       dto.Name,
		Email:      dto.Email,
		Password:   string(password),
		Created_at: time,
	})

	if err != nil {
		log.Error("Register user error", slog.Any("error", err.Error()))
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	roleId, err := s.addRoleForUser(ctx, userId, dto.Role)

	if err != nil {
		log.Error("Get role error", slog.Any("error", err.Error()))
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	accessToken, refreshToken, err := s.usC.Generate(ctx, roleId, dto.Role)

	if err != nil {
		log.Error("Generate token error", slog.Any("error", err.Error()))
		return "", "", models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Register user success",
		slog.String("token", accessToken),
		slog.String("refreshToken", refreshToken),
	)

	return accessToken, refreshToken, models.User{
		Id:        userId,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  string(password),
		CreatedAt: time,
		Roles:     s.writeUserRoles(roleId, dto.Role),
	}, nil

}
