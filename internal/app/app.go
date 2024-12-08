package app

import (
	"log/slog"
	app_crm "server_crm/internal/app/crm"
	app_migrator "server_crm/internal/app/migrator"
	app_swaggerUpdater "server_crm/internal/app/swaggerUpdater"
	"server_crm/internal/config"
	user_crypt "server_crm/internal/lib/crypt"
	user_serivce "server_crm/internal/services/realize/user"
	storage_postgres "server_crm/internal/storage/postgres"
)

type App struct {
	l      *slog.Logger
	crmApp *app_crm.App
}

func New(conf *config.Config, l *slog.Logger) *App {

	storage, err := storage_postgres.New(storage_postgres.DBSettings{
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		DbName:   conf.DB.DbName,
	}, l)

	if err != nil {
		panic(err)
	}

	userCrypter := user_crypt.NewJWTManager(conf.Jwt.Secret, conf.Jwt.Duration)
	
	userService := user_serivce.New(l, storage.User, userCrypter, user_serivce.RolesProvider{
		Admin:  storage.Role.Admin,
		Owner:  storage.Role.Owner,
		Client: storage.Role.Client,
	})

	app_migrator.New(l, storage).MustRun()
	app_swaggerUpdater.New(l).MustRun()

	return &App{
		l:      l,
		crmApp: app_crm.New(conf.Rest.Port, l, userService, userCrypter),
	}
}

func (a App) MustRun() {

	a.crmApp.MustRun()

}

func (a App) Stop() {

}
