package storage_postgres

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	storage_user "server_crm/internal/storage/postgres/user"
	"server_crm/migrations"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB   *sql.DB
	log  *slog.Logger
	User storage_user.UserStorage
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
