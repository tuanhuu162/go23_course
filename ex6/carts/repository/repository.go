package repository

import (
	"fmt"

	"github.com/tuanhuu162/go23_course/ex6/carts"
	"github.com/tuanhuu162/go23_course/ex6/models"
	productRepo "github.com/tuanhuu162/go23_course/ex6/products/repository"
)

type CartRepository struct {
	pr productRepo.ProductRepository
}

func NewCartRepository(pr productRepo.ProductRepository) carts.CartRepositoryInterface {
	return &CartRepository{
		pr: pr,
	}
}

func (c *CartRepository) AddProductToCart(payload models.CartRequestPayload, cart *models.Cart) error {
	product, err := c.pr.GetProduct(payload.ProductID)
	if err != nil {
		return err
	}
	if item, ok := (*cart).Items[payload.ProductID]; ok {
		item.Quantity += payload.Quantity
		(*cart).Items[payload.ProductID] = item
	} else {
		if (*cart).Items == nil {
			(*cart).Items = make(map[uint]models.Item)
		}
		(*cart).Items[payload.ProductID] = models.Item{
			ProductID:   payload.ProductID,
			ProductName: product.Name,
			Quantity:    payload.Quantity,
		}
	}
	cart.Total += product.Price * int64(payload.Quantity)
	return nil
}

func (c *CartRepository) DeleteProductFromCart(payload models.CartRequestPayload, cart *models.Cart) error {
	product, err := c.pr.GetProduct(payload.ProductID)
	if err != nil {
		return err
	}
	if item, ok := (*cart).Items[payload.ProductID]; ok {
		if (*cart).Items[payload.ProductID].Quantity > payload.Quantity {
			item.Quantity -= payload.Quantity
			(*cart).Items[payload.ProductID] = item
		} else if (*cart).Items[payload.ProductID].Quantity == payload.Quantity {
			delete((*cart).Items, payload.ProductID)
		} else {
			return fmt.Errorf("Quantity of product %d in cart is less than %d", payload.ProductID, payload.Quantity)
		}
		(*cart).Total -= product.Price * int64(payload.Quantity)
	} else {
		return fmt.Errorf("Product %d is not in cart", payload.ProductID)
	}
	return nil
}