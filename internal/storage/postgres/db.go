package storage_postgres

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
	storage_autoservice "server_crm/internal/storage/postgres/autoservice"
	storage_car "server_crm/internal/storage/postgres/car"
	storage_catalog "server_crm/internal/storage/postgres/catalog"
	storage_operation "server_crm/internal/storage/postgres/operation"
	storage_role "server_crm/internal/storage/postgres/role"
	storage_user "server_crm/internal/storage/postgres/user"
	storage_work "server_crm/internal/storage/postgres/work"
	"server_crm/migrations"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB   *sql.DB
	log  *slog.Logger
	User storage_user.UserStorage
	Role storage_role.RolesStorage
	Autoservice storage_models.AutoserviceDomain
	Car storage_models.CarDomain
	Catalog storage_models.CatalogDomain
	Work storage_models.WorkDomain
	Operation storage_models.OperationDomain
}
type DBSettings struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func New(cfg DBSettings, l *slog.Logger) (*Storage, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.DbName,
		cfg.Host,
		cfg.Port,
	))
	l.Info(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.DbName,
		cfg.Host,
		cfg.Port,
	))
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %s", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %s", err)
	}

	storage := &Storage{
		log:  l,
		DB:   db,
		User: storage_user.New(db),
		Role : storage_role.New(db),
		Autoservice: storage_autoservice.New(db),
		Car: storage_car.New(db),
		Catalog: storage_catalog.New(db),
		Work: storage_work.New(db),
		Operation: storage_operation.New(db),

	}
	return storage, nil
}

func (db *Storage) Stop() error {
	if db.DB != nil {
		err := db.DB.Close()
		{
			if err != nil {
				log.Fatalf("Failed to closed database: %s", err)
				return err
			}
		}
	}
	return nil
}

func (db Storage) Migrate() error {

	const op = "db.Migrate"

	log := db.log.With(slog.String("op", op))

	log.Info("Statrting migrate")

	err := migrations.Migrations(db.DB)

	if err != nil {
		return fmt.Errorf("%s %e", op, err)
	}
	log.Info("Migrate success")

	return nil

}
