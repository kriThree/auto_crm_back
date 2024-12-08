package autoservice_service

import (
	"log/slog"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)



// Ссылка на интерефейс в месте объявления, далее сервисы автосервиса будут ссылаться на
// autoserviceProvider позволяя без замены ссылки подменить реализации
type AutoserviceProvider = storage_models.AutoserviceDomain

type OwnerProvider = storage_models.RoleDomain

type AutoserviceService struct {
	auP AutoserviceProvider
	owP OwnerProvider
	l   *slog.Logger
}

func New(autoserviceProvider AutoserviceProvider, ownerProvider OwnerProvider, l *slog.Logger) models.AutoserviceRepo {
	return AutoserviceService{
		auP: autoserviceProvider,
		owP: ownerProvider,
		l:   l,
	}
}
