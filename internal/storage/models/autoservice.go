package storage_models

import "time"

type Autoservice struct {
	Id        int64
	Name      string
	Address   string
	Phone     string
	Email     string
	OwnerId   int64
	CreatedAt time.Time
}

type AddAutoserviceDto struct {
	Name      string
	Address   string
	Phone     string
	Email     string
	Owner_id  int64
	Created_at time.Time
}