package app_crm

import (
	"fmt"
	"log/slog"
	"server_crm/internal/modules/rest"
	user_controller "server_crm/internal/modules/rest/controllers/user"
	"server_crm/internal/modules/rest/middlewares"
	"server_crm/internal/services/models"
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
	autoserviceService models.AutoserviceRepo,
	carService models.CarRepo,
	catalogService models.CatalogRepo,
	operationService models.OperationRepo,
	workService models.WorkRepo,
	userCrypter middlewares.UserDecrypter,
) *App {

	restController := rest.New(userService,
		autoserviceService,
		carService,
		catalogService,
		operationService,
		workService,
		userCrypter,
	)

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
