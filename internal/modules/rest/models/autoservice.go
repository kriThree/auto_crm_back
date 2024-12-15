package rest_models

type Autoservice struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	OwnerId   int64  `json:"owner_id"`
	CreatedAt string `json:"created_at"`
}

type AddAutoserviceReq struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
type AddAutoserviceRes struct {
	Id int64 `json:"id"`
}
type GetAutoserviceRes struct {
	Autoservices []Autoservice `json:"autoservices"`
}

type UpdateAutoserviceReq struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type DeleteAutoserviceReq struct {
	Id int64 `json:"id"`
}
