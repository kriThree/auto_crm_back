package app_swaggerUpdater

import (
	"fmt"
	"log/slog"
	"os/exec"
	"time"
)

type App struct {
	l *slog.Logger
}

func New(l *slog.Logger) *App {
	return &App{l: l}
}

func (a *App) MustRun() {

	const op = "swaggerUpdaterApp.Run"

	log := a.l.With(slog.String("op", op))

	cmd := exec.Command("swag", "init", "-g", "cmd/crm/main.go", "--instanceName", "crm_swagger")
	log.Debug("cmd", slog.String("cmd", cmd.String()))
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Error("Ошибка:", slog.Any("error", err.Error()))
	}
	fmt.Println(string(output))
	time.Sleep(time.Second * 2)

}
