Сборка сваггера:swag init -g cmd/crm/main.go  --instanceName crm_swagger
Goose 
GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=krithree password= dbname=crm host=localhost port=5432 sslmode=disable" goose down -dir=migrations/migrations