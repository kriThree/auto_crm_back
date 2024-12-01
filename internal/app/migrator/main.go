package app_migrator

import (
	"fmt"
	"log/slog"
)

type Migrator interface {
	Migrate() error
}

type App struct {
	migrator Migrator
	log      *slog.Logger
}

func New(log *slog.Logger, migrator Migrator) *App {
	return &App{
		log:      log,
		migrator: migrator,
	}
}

func (a App) Run() error {
	const op = "migratorApp.Run"

	if err := a.migrator.Migrate(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (a App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
