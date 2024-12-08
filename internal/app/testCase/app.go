package app_testcase

import (
	"context"
	"log/slog"
	app_migrator "server_crm/internal/app/migrator"
	"server_crm/internal/config"
	user_crypt "server_crm/internal/lib/crypt"
	"server_crm/internal/services/models"
	autoservice_service "server_crm/internal/services/realize/autoservice"
	car_service "server_crm/internal/services/realize/car"
	catalog_service "server_crm/internal/services/realize/catalog"
	operation_service "server_crm/internal/services/realize/operation"
	user_serivce "server_crm/internal/services/realize/user"
	work_service "server_crm/internal/services/realize/work"
	storage_models "server_crm/internal/storage/models"
	storage_postgres "server_crm/internal/storage/postgres"
	"time"
)

type App struct {
	l        *slog.Logger
	services Services
}
type Services struct {
	userService        user_serivce.UserService
	autoserviceService models.AutoserviceRepo
	carService         models.CarRepo
	catalogService     models.CatalogRepo
	workService        models.WorkRepo
	operatinoService   models.OperationRepo
}

func New(conf *config.Config, l *slog.Logger) *App {

	storage, err := storage_postgres.New(storage_postgres.DBSettings{
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		DbName:   conf.DB.DbName,
	}, l)

	userCrypter := user_crypt.NewJWTManager(conf.Jwt.Secret, conf.Jwt.Duration)

	userService := user_serivce.New(l, storage.User, userCrypter, user_serivce.RolesProvider{
		Admin:  storage.Role.Admin,
		Owner:  storage.Role.Owner,
		Client: storage.Role.Client,
	})

	autoserviceService := autoservice_service.New(storage.Autoservice, storage.Role.Owner, l)
	carService := car_service.New(storage.Car, storage.Role.Client, l)
	catalogService := catalog_service.New(storage.Catalog, storage.Role.Admin, storage.Work, l)
	workService := work_service.New(storage.Work, l)
	operationService := operation_service.New(storage.Operation, l)

	if err != nil {
		panic(err)
	}

	app_migrator.New(l, storage).MustRun()

	return &App{
		l: l,
		services: Services{
			userService:        *userService,
			autoserviceService: autoserviceService,
			carService:         carService,
			catalogService:     catalogService,
			workService:        workService,
			operatinoService:   operationService,
		},
	}
}

func (a App) MustRun() {

	log := a.l.With(slog.String("op", "tests"))

	log.Info("Starting tests")

	_, _, owner, err := a.services.userService.Register(context.Background(), models.RegisterUserDto{
		Name:     "Owner",
		Email:    time.Now().String() + "owner@ya.ru",
		Password: "test",
		Role:     storage_models.ROLE_OWNER,
	})
	log.Info("Register owner success",
		slog.Int64("id", owner.Roles.Owner))
	if err != nil {
		log.Error("Register owner error", slog.Any("error", err.Error()))
		return
	}
	_, _, client, err := a.services.userService.Register(context.Background(), models.RegisterUserDto{
		Name:     "Client",
		Email:    time.Now().String() + "client@ya.ru",
		Password: "test",
		Role:     storage_models.ROLE_CLIENT,
	})
	log.Info("Register client success",
		slog.Int64("id", client.Roles.Client))
	if err != nil {
		log.Error("Register owner error", slog.Any("error", err.Error()))
		return
	}
	_, _, admin, err := a.services.userService.Register(context.Background(), models.RegisterUserDto{
		Name:     "Admin",
		Email:    time.Now().String() + "admin@ya.ru",
		Password: "test",
		Role:     storage_models.ROLE_ADMIN,
	})
	log.Info("Register admin success",
		slog.Int64("id", admin.Roles.Admin))
	if err != nil {
		log.Error("Register owner error", slog.Any("error", err.Error()))
		return
	}

	autoserviceId, err := a.services.autoserviceService.Add(context.Background(), storage_models.AddAutoserviceDto{
		Name:     "Autoservice with Pertovich",
		Address:  "Test street",
		Phone:    "122131912",
		Email:    "test@ya.ru",
		Owner_id: owner.Roles.Owner,
	})

	if err != nil {
		log.Error("Add autoservice error", slog.Any("error", err.Error()))
		return
	}

	carId, err := a.services.carService.Add(context.Background(), storage_models.AddCarDto{
		Number:      "A462KL21",
		Description: time.Now().String() + "Blue BMW",
		ClientId:    client.Roles.Client,
	})

	if err != nil {
		log.Error("Add car error", slog.Any("error", err.Error()))
		return
	}
	catalogId, err := a.services.catalogService.Add(context.Background(), storage_models.AddCatalogDto{
		Name:    "Test catalog",
		AdminId: admin.Roles.Admin,
	})

	if err != nil {
		log.Error("Add catalog error", slog.Any("error", err.Error()))
		return
	}

	workId, err := a.services.workService.Add(context.Background(), storage_models.AddWorkDto{
		Name:      "Test work",
		Cost:      100,
		CatalogId: catalogId,
	})

	if err != nil {
		log.Error("Add work error", slog.Any("error", err.Error()))
		return
	}

	operationId, err := a.services.operatinoService.Create(context.Background(), storage_models.AddOperationDto{
		Description:   "Test operation",
		CarId:         carId,
		WorkId:        workId,
		AutoserviceId: autoserviceId,
	})

	if err != nil {
		log.Error("Add operation error", slog.Any("error", err.Error()))
		return
	}

	log.Info("",
		slog.Int64("owner_id", owner.Id),
		slog.Int64("client_id", client.Id),
		slog.Int64("admin_id", admin.Id),
		slog.Int64("autoservice_id", autoserviceId),
		slog.Int64("car_id", carId),
		slog.Int64("catalog_id", catalogId),
		slog.Int64("work_id", workId),
		slog.Int64("operation_id", operationId),
	)

	autoservice, err := a.services.autoserviceService.Get(context.Background(), owner.Id)

	if err != nil {
		log.Error("Get autoservice error", slog.Any("error", err.Error()))
		return
	}

	log.Info("",
		slog.Any("autoservice", autoservice),
	)

	err = a.services.autoserviceService.Update(context.Background(), storage_models.UpdateAutoserviceDto{
		Id:      autoserviceId,
		Name:    "test2",
		Address: "test2",
		Phone:   "test2",
		Email:   "test2",
	})

	if err != nil {
		log.Error("Update autoservice error", slog.Any("error", err.Error()))
		return
	}

	err = a.services.operatinoService.Delete(context.Background(), operationId)

	if err != nil {
		log.Error("Delete operation error", slog.Any("error", err.Error()))
		return
	}

	err = a.services.autoserviceService.Delete(context.Background(), autoserviceId)

	if err != nil {
		log.Error("Delete autoservice error", slog.Any("error", err.Error()))
		return
	}

	err = a.services.workService.Delete(context.Background(), workId)

	if err != nil {
		log.Error("Delete work error", slog.Any("error", err.Error()))
		return
	}

	err = a.services.carService.Delete(context.Background(), carId)

	if err != nil {
		log.Error("Delete car error", slog.Any("error", err.Error()))
		return
	}

	err = a.services.catalogService.Delete(context.Background(), catalogId)

	if err != nil {
		log.Error("Delete catalog error", slog.Any("error", err.Error()))
		return
	}

	err = a.services.userService.Delete(context.Background(), owner.Id)

	if err != nil {
		log.Error("Delete user error", slog.Any("error", err.Error()))
		return
	}
}
