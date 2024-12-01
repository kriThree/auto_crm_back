package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	user_controller "server_crm/internal/modules/rest/controllers/user"
	"server_crm/internal/modules/rest/middlewares"

	_ "server_crm/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Controller struct {
	Router *mux.Router
}

func New(
	uc user_controller.UserUsecase,
	usC middlewares.UserDecrypter,
) *Controller {

	router := mux.NewRouter()

	// router.Use(middleware.Logger())
	// router.Use(middleware.Recover())
	router.Use(middlewares.CorsMiddleware())
	
	router.Methods(http.MethodOptions).HandlerFunc(OptionsOK)

	apiRouter := router.PathPrefix("").Subrouter()

	authMiddleware := middlewares.AuthMiddleware(usC)

	user_controller.Register(uc, authMiddleware, apiRouter.PathPrefix("/user").Subrouter())

	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	return &Controller{
		Router: router,
	}
}

func (c Controller) Run(l *slog.Logger, port int) error {

	const op = "rest.Run"

	log := l.With(slog.String("op", op))

	log.Info("Starting server", slog.Int("port", port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), c.Router)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("REST server is running",
		slog.String("port", fmt.Sprintf("%d", port)),
	)

	return nil
}
