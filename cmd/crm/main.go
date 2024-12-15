package main

import (
	"log/slog"
	"os"
	"os/signal"
	"server_crm/internal/app"
	"server_crm/internal/config"
	"server_crm/internal/lib/log"
	"syscall"

)

// @title 				Swagger Example API
// @host localhost:8001
// @BasePath /api
// @Security http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	conf := config.MustLoad()
	log := log.LogInit()

	application := app.New(conf, log)

	application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application...", slog.String("signal", sign.String()))

	application.Stop()

	log.Info("aplication stopped")
}
