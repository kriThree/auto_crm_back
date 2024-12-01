package app_crm

import (
	"fmt"
	"log/slog"
	"server_crm/internal/modules/rest"
	user_controller "server_crm/internal/modules/rest/controllers/user"
	"server_crm/internal/modules/rest/middlewares"
)

type App struct {
	port           int
	log            *slog.Logger
	restController *rest.Controller
}

func New(
	port int,
	log *slog.Logger,
	userService user_controller.UserUsecase,
	userCrypter middlewares.UserDecrypter,
) *App {

	restController := rest.New(userService, userCrypter)

	return &App{
		port:           port,
		log:            log,
		restController: restController,
	}
}

func (a App) Run() error {

	const op = "crmApp.Run"

	if err := a.restController.Run(a.log, a.port); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}

}
