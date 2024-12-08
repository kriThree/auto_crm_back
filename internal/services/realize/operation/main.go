package operation_service

import (
	"log/slog"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

type OperationService struct {
	opP OperationProvider
	wkP WorkProvider
	crP CarProvider
	auP AutoserviceProvider
	l   *slog.Logger
}
type OperationProvider = storage_models.OperationDomain
type WorkProvider = storage_models.WorkDomain
type CarProvider = storage_models.CarDomain
type AutoserviceProvider = storage_models.AutoserviceDomain

func New(operationPrvoder OperationProvider, l *slog.Logger) models.OperationRepo {
	return OperationService{
		opP: operationPrvoder,
		l:   l,
	}
}
