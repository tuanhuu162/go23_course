package models

type Cart struct {
	Items map[uint]Item
	Total int64
}

type Item struct {
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    uint   `json:"quantity"`
}

type CartRequestPayload struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
