package models

type Cart struct {
	Items map[string]Item
	Total int64
}

type CartRemoveRequestPayload struct {
	ProductID string `json:"product_id"`
}

type Item struct {
	ProductID string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}
