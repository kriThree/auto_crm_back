package car_controller

import (
	"net/http"
	"server_crm/internal/services/models"

	"github.com/gorilla/mux"
)

type CarUsecase = models.CarRepo

type CarController struct {
	crUc CarUsecase
}

func (h CarController) Handle(router *mux.Router, authMiddleware mux.MiddlewareFunc, roleMiddleware mux.MiddlewareFunc) (*mux.Router, error) {

	router.Use(authMiddleware)
	router.Use(roleMiddleware)

	router.HandleFunc("/", http.HandlerFunc(h.Add)).Methods("POST")
	router.HandleFunc("/", http.HandlerFunc(h.Get)).Methods("GET")
	router.HandleFunc("/", http.HandlerFunc(h.Update)).Methods("PATCH")
	router.HandleFunc("/", http.HandlerFunc(h.Delete)).Methods("DELETE")

	return router, nil
}

func Register(
	router *mux.Router,
	carUsecase CarUsecase,
	authMiddleware mux.MiddlewareFunc,
	roleMiddleware mux.MiddlewareFunc,
) *mux.Router {

	carController := &CarController{
		crUc: carUsecase,
	}

	carController.Handle(router, authMiddleware, roleMiddleware)

	return router

}
