package work_controller

import (
	"net/http"
	"server_crm/internal/services/models"

	"github.com/gorilla/mux"
)

type WorkUsecase = models.WorkRepo

type WorkController struct {
	wkUc WorkUsecase
}

func (h WorkController) Handle(router *mux.Router, authMiddleware mux.MiddlewareFunc, roleMiddleware mux.MiddlewareFunc) (*mux.Router, error) {

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
	workUsecase WorkUsecase,
	authMiddleware mux.MiddlewareFunc,
	roleMiddleware mux.MiddlewareFunc,
) *mux.Router {

	workController := &WorkController{
		wkUc: workUsecase,
	}

	workController.Handle(router, authMiddleware, roleMiddleware)

	return router

}
