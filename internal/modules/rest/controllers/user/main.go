package user_controller

import (
	"context"
	"net/http"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"

	"github.com/gorilla/mux"
)

type UserUsecase interface {
	Register(ctx context.Context, dto models.RegisterUserDto) (string, string, models.User, error)
	Update(ctx context.Context, userId int64, dto storage_models.UpdateUserDto) error
	Login(ctx context.Context, email string, password string) (string, string, models.User, error)
	GetOne(ctx context.Context, id int64) (models.User, error)
	Get(ctx context.Context) ([]storage_models.User, error)
	Authorize(ctx context.Context, accessToken string, refreshToken string) (string, string, models.User, error)
}
type UserController struct {
	uc UserUsecase
}

func (h UserController) Handle(router *mux.Router, authMiddleware mux.MiddlewareFunc) *mux.Router {

	router.HandleFunc("/register", http.HandlerFunc(h.Register)).Methods("POST")

	router.HandleFunc("/login", h.Login).Methods("POST")

	router.HandleFunc("/authorize", h.Authorize).Methods("GET")
	router.Use(authMiddleware)
	
	router.HandleFunc("/", h.Update).Methods("PUT")

	return router
}

func Register(uc UserUsecase, authMiddleware mux.MiddlewareFunc, router *mux.Router) *mux.Router {
	userController := UserController{
		uc: uc,
	}
	userController.Handle(router, authMiddleware)
	return router
}
