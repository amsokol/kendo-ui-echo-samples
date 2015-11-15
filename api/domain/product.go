package domain

type Product struct {
	Id   int    `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type ProductStore interface {
	GetAll() ([]Product, error)
}
