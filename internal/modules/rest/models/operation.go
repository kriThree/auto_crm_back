package rest_models

type Operation struct {
	Id          int64       `json:"id"`
	Desctiption string      `json:"description"`
	Car         Car         `json:"car"`
	Work        Work        `json:"work"`
	Autoservice Autoservice `json:"autoservice"`
}

type AddOperationReq struct {
	Description   string `json:"description"`
	CarID         int64  `json:"car_id"`
	WorkID        int64  `json:"work_id"`
	AutoserviceId int64  `json:"autoservice_id"`
}
type AddOperationRes struct {
	Id int64 `json:"id"`
}
type GetOperationsReq struct {
	CarId         int64 `json:"car_id"`
	WorkId        int64 `json:"work_id"`
	AutoserviceId int64 `json:"autoservice_id"`
}
type GetOperationsRes struct {
	Operations []Operation `json:"operations"`
}
type UpdateOperationReq struct {
	Id            int64  `json:"id"`
	Description   string `json:"description"`
	CarId         int64  `json:"car_id"`
	WorkId        int64  `json:"work_id"`
	AutoserviceId int64  `json:"autoservice_id"`
}

type DeleteOperationReq struct {
	Id int64 `json:"id"`
}
