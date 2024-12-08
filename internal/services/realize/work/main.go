package work_service

import (
	"log/slog"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

type WorkService struct {
	wkP WorkProvider
	l   *slog.Logger
}
type WorkProvider = storage_models.WorkDomain

func New(workPrvoder WorkProvider, l *slog.Logger) models.WorkRepo {
	return WorkService{
		wkP: workPrvoder,
		l:   l,
	}
}
