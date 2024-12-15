package autoservice_controller

import (
	"net/http"
	"server_crm/internal/services/models"
	"github.com/gorilla/mux"
)

type AutoserviceUsecase = models.AutoserviceRepo

type AutoserviceController struct {
	autUc AutoserviceUsecase
}

func (h AutoserviceController) Handle(router *mux.Router, authMiddleware mux.MiddlewareFunc, roleMiddleware mux.MiddlewareFunc) (*mux.Router, error) {

	router.Use(authMiddleware)
	router.Use(roleMiddleware)

	router.HandleFunc("/", http.HandlerFunc(h.Add)).Methods("POST")
	router.HandleFunc("/", http.HandlerFunc(h.Get)).Methods("GET")
	router.HandleFunc("/", http.HandlerFunc(h.Update)).Methods("PATCH")
	router.HandleFunc("/", http.HandlerFunc(h.Delete)).Methods("DELETE")

	return router, nil
}

func Register(router *mux.Router, autUc AutoserviceUsecase, authMiddleware mux.MiddlewareFunc, roleMiddleware mux.MiddlewareFunc) *mux.Router {

	autoserviceController := &AutoserviceController{
		autUc: autUc,

	}

	autoserviceController.Handle(router, authMiddleware, roleMiddleware)

	return router

}
