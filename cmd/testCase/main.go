package main

import (
	"os"
	"os/signal"
	app_testcase "server_crm/internal/app/testCase"
	"server_crm/internal/config"
	"server_crm/internal/lib/log"
	"syscall"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title           Swagger Example API
// @host localhost:8001
// @BasePath /api
// @Security https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	conf := config.MustLoad()
	log := log.LogInit()

	application := app_testcase.New(conf, log)

	application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)


}
