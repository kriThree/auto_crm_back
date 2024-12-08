package car_service

import (
	"log/slog"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

type CarService struct {
	crP CarProvider
	clP storage_models.RoleDomain
	l   *slog.Logger
}
type ClientProvider = storage_models.RoleDomain
type CarProvider = storage_models.CarDomain

func New(carPrvoder CarProvider, clientProvider ClientProvider, l *slog.Logger) models.CarRepo {
	return CarService{
		crP: carPrvoder,
		l:   l,
		clP: clientProvider,
	}
}
