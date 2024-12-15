package rest_models

type Catalog struct {
	Id        int64           `json:"id"`
	Name      string          `json:"name"`
	Works     []WorkInCatalog `json:"works"`
}
type WorkInCatalog struct {
	Id   int64  `json:"id"`
	Cost int    `json:"cost"`
	Name string `json:"name"`
}

type AddCatalogReq struct {
	Name string `json:"name"`
}
type AddCatalogRes struct {
	Id int64 `json:"id"`
}
type GetCatalogRes struct {
	Catalogs []Catalog `json:"catalogs"`
}

type UpdateCatalogReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type DeleteCatalogReq struct {
	Id int64 `json:"id"`
}
