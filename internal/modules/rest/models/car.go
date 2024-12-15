package rest_models

type Car struct {
	Id          int64  `json:"id"`
	Number      string `json:"number"`
	Description string `json:"description"`
}

type AddCarReq struct {
	Number      string `json:"number"`
	Description string `json:"description"`
	ClientId    int64  `json:"client_id"`
}
type AddCarRes struct {
	Id int64 `json:"id"`
}
type GetCarsRes struct {
	Cars []Car `json:"cars"`
}

type UpdateCarReq struct {
	Id          int64  `json:"id"`
	Number      string `json:"number"`
	Description string `json:"description"`
}

type DeleteCarReq struct {
	Id int64 `json:"id"`
}
