package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	autoservice_controller "server_crm/internal/modules/rest/controllers/autoservice"
	car_controller "server_crm/internal/modules/rest/controllers/car"
	catalog_controller "server_crm/internal/modules/rest/controllers/catalog"
	operation_controller "server_crm/internal/modules/rest/controllers/operation"
	user_controller "server_crm/internal/modules/rest/controllers/user"
	work_controller "server_crm/internal/modules/rest/controllers/work"
	"server_crm/internal/modules/rest/middlewares"
	storage_models "server_crm/internal/storage/models"

	_ "server_crm/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Controller struct {
	Router *mux.Router
}

func New(
	userUsecase user_controller.UserUsecase,
	autoserviceUsecase autoservice_controller.AutoserviceUsecase,
	carUsecase car_controller.CarUsecase,
	catalogUsecase catalog_controller.CatalogUsecase,
	operationUsecase operation_controller.OperationUsecase,
	workUsecase work_controller.WorkUsecase,
	usC middlewares.UserDecrypter,
) *Controller {

	router := mux.NewRouter()

	// router.Use(middleware.Logger())
	// router.Use(middleware.Recover())
	router.Use(middlewares.CorsMiddleware())

	router.Methods(http.MethodOptions).HandlerFunc(OptionsOK)

	apiRouter := router.PathPrefix("/api").Subrouter()

	authMiddleware := middlewares.AuthMiddleware(usC)

	ownerMiddleware, _ := middlewares.CheckRole([]string{storage_models.ROLE_OWNER})

	clientMiddleware, _ := middlewares.CheckRole([]string{storage_models.ROLE_CLIENT})

	adminMiddleware, _ := middlewares.CheckRole([]string{storage_models.ROLE_ADMIN})
	
	allMiddleware, _ := middlewares.CheckRole([]string{storage_models.ROLE_ADMIN, storage_models.ROLE_OWNER, storage_models.ROLE_CLIENT})

	autoservice_controller.Register(apiRouter.PathPrefix("/autoservices").Subrouter(), autoserviceUsecase, authMiddleware, ownerMiddleware)
	
	user_controller.Register(userUsecase, authMiddleware, apiRouter.PathPrefix("/user").Subrouter())

	car_controller.Register(apiRouter.PathPrefix("/cars").Subrouter(), carUsecase, authMiddleware, clientMiddleware)
	catalog_controller.Register(apiRouter.PathPrefix("/catalogs").Subrouter(), catalogUsecase, authMiddleware, adminMiddleware)	
	operation_controller.Register(apiRouter.PathPrefix("/operations").Subrouter(), operationUsecase, authMiddleware, allMiddleware)
	work_controller.Register(apiRouter.PathPrefix("/works").Subrouter(), workUsecase, authMiddleware, adminMiddleware)


	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

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
