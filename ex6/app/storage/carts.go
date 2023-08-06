package storage

import (
	"github.com/tuanhuu162/go23_course/ex6/app/models"
)

func AddProductToCart(payload models.CartRequestPayload, cart *models.Cart) error {
	product, err := GetProduct(payload.ProductID)
	if err != nil {
		return err
	}
	if item, ok := (*cart).Items[payload.ProductID]; ok {
		item.Quantity += payload.Quantity
		(*cart).Items[payload.ProductID] = item
	} else {
		(*cart).Items[payload.ProductID] = models.Item{
			ProductID:   payload.ProductID,
			ProductName: product.Name,
			Quantity:    payload.Quantity,
		}
	}
	return nil
}

func DeleteProductFromCart(payload models.CartRequestPayload, cart *models.Cart) error {
	if _, err := GetProduct(payload.ProductID); err != nil {
		return err
	}
	if item, ok := (*cart).Items[payload.ProductID]; ok {
		if (*cart).Items[payload.ProductID].Quantity < payload.Quantity {
			item.Quantity -= payload.Quantity
			(*cart).Items[payload.ProductID] = item
		} else {
			delete((*cart).Items, payload.ProductID)
		}
	}
	return nil
}
