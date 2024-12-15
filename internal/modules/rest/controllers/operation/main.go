package operation_controller

import (
	"net/http"
	"server_crm/internal/services/models"

	"github.com/gorilla/mux"
)

type OperationUsecase = models.OperationRepo

type OperationController struct {
	opUc OperationUsecase
}

func (h OperationController) Handle(router *mux.Router, authMiddleware mux.MiddlewareFunc, roleMiddleware mux.MiddlewareFunc) (*mux.Router, error) {

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
	operationUsecase OperationUsecase,
	authMiddleware mux.MiddlewareFunc,
	roleMiddleware mux.MiddlewareFunc,
) *mux.Router {

	operationController := &OperationController{
		opUc: operationUsecase,
	}

	operationController.Handle(router, authMiddleware, roleMiddleware)

	return router

}
