package user_serivce

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

type UserService struct {
	l   *slog.Logger
	usP UserProvider
	usC UserCrypter
}
type UserProvider interface {
	Add(ctx context.Context, dto storage_models.AddUserDto) (int64, error)
	ReadOne(ctx context.Context, id int64) (storage_models.User, error)
	Read(ctx context.Context) ([]storage_models.User, error)
	FindByEmail(ctx context.Context, email string) (storage_models.User, error)
	Update(ctx context.Context, dto storage_models.UpdateUserDto) error
}
type UserCrypter interface {
	Generate(ctx context.Context, userID int64) (string, string, error)
	Validate(ctx context.Context, tokenStr string) (userId int64, err error)
}

func New(l *slog.Logger, usP UserProvider, uc UserCrypter) *UserService {
	return &UserService{l: l, usP: usP, usC: uc}
}
