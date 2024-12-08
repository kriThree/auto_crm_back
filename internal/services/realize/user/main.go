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
	rsP RolesProvider
}

// Интрефейс по месту использования как пример, далее будут использоваться реализации с интерфейсами по месту объявления
type UserProvider interface {
	Add(ctx context.Context, dto storage_models.AddUserDto) (int64, error)
	GetOne(ctx context.Context, id int64) (storage_models.User, error)
	Get(ctx context.Context) ([]storage_models.User, error)
	FindByEmail(ctx context.Context, email string) (storage_models.User, error)
	Update(ctx context.Context, dto storage_models.UpdateUserDto) error
	Delete(ctx context.Context, id int64) error
}
type RoleProvider interface {
	Add(ctx context.Context, userId int64) (int64, error)
	GetByUserId(ctx context.Context, userId int64) (int64, error)
	Delete(ctx context.Context, userId int64) error
}
type RolesProvider struct {
	Admin  RoleProvider
	Owner  RoleProvider
	Client RoleProvider
}

type UserCrypter interface {
	Generate(ctx context.Context, userID int64,role string) (string, string, error)
	Validate(ctx context.Context, tokenStr string) (userId int64,role string, err error)
}

func New(l *slog.Logger, usP UserProvider, uc UserCrypter,rsP RolesProvider) *UserService {
	return &UserService{l: l, usP: usP, usC: uc, rsP: rsP}
}
