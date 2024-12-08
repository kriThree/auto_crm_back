package test_functional_user

import (
	"context"
	"log/slog"
	"server_crm/internal/services/models"
	user_serivce "server_crm/internal/services/realize/user"
	storage_models "server_crm/internal/storage/models"
	test_functional "server_crm/internal/tests/functional"
	"time"
)

type UserTest struct {
	userService user_serivce.UserService
	log         *slog.Logger
	test_functional.TestEntity
}
func New(userService user_serivce.UserService, log *slog.Logger) UserTest {

	return UserTest{
		userService: userService,
		log:         log,
	}
}
func (a UserTest) Create() models.User {

	email := time.Now().String() + "test@ya.ru"
	password := "test"

	_, _, userRegister, err := a.userService.Register(context.Background(), models.RegisterUserDto{
		Name:     "test",
		Email:    email,
		Password: password,
		Role:     storage_models.ROLE_OWNER,
	})

	if err != nil {
		a.log.Error("Register user error", slog.Any("error", err.Error()))
		return models.User{}
	}

	_, _, userLogin, err := a.userService.Login(context.Background(), userRegister.Email, "test")

	if err != nil {
		a.log.Error("Login user error", slog.Any("error", err.Error()))
		return models.User{}
	}

	if userRegister.Id != userLogin.Id {
		a.log.Error("User id not equal", slog.Int64("register_id", userRegister.Id), slog.Int64("login_id", userLogin.Id))
		return models.User{}
	}
	
	userGet,err := a.userService.GetOne(context.Background(), userRegister.Id)

	if err != nil {
		a.log.Error("Get user error", slog.Any("error", err.Error()))
		return models.User{}
	}

	if userRegister.Id != userGet.Id {
		a.log.Error("User id not equal", slog.Int64("register_id", userRegister.Id), slog.Int64("get_id", userGet.Id))
		return models.User{}
	}

	err = a.userService.Update(context.Background(), userRegister.Id, storage_models.UpdateUserDto{
		Name:     "test",
		Email:    email,
		Password: password,
	})

	if err != nil {
		a.log.Error("Update user error", slog.Any("error", err.Error()))
		return models.User{}
	}


	return userGet

}
