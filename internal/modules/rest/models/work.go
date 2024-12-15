package rest_models

type Work struct {
	Id   int64  `json:"id"`
	Cost int    `json:"cost"`
	Name string `json:"name"`
}

type AddWorkReq struct {
	Name    string `json:"name"`
	Cost    int    `json:"cost"`
	CatalogId int64  `json:"catalog_id"`
}
type AddWorkRes struct {
	Id int64 `json:"id"`
}
type GetWorksRes struct {
	Works []Work `json:"works"`
}

type UpdateWorkReq struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type DeleteWorkReq struct {
	Id int64 `json:"id"`
}
