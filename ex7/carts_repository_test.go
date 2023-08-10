package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cartRepository "github.com/tuanhuu162/go23_course/ex6/carts/repository"
	"github.com/tuanhuu162/go23_course/ex6/models"
	"github.com/tuanhuu162/go23_course/ex6/products"
)

func TestAddProductToCart(t *testing.T) {
	mockProductRepository := new(products.MockProductRepositoryInterface)
	mockProductRepository.On("GetProduct", uint(1)).Return(models.Product{
		Name:  "abc",
		Price: 100,
	}, nil)

	cart := models.Cart{
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    1,
			},
		},
		Total: 100,
	}
	payload := models.CartRequestPayload{
		ProductID: 1,
		Quantity:  1,
	}
	CartRepository := cartRepository.NewCartRepository(mockProductRepository)
	err := CartRepository.AddProductToCart(payload, &cart)
	assert.NoError(t, err)
	assert.Equal(t, cart.Total, int64(200))
	assert.Equal(t, cart.Items[1].Quantity, uint(2))
	assert.Equal(t, cart.Items[1].ProductName, "abc")
}

func TestDeleteProductFromCart(t *testing.T) {
	mockProductRepository := new(products.MockProductRepositoryInterface)
	mockProductRepository.On("GetProduct", uint(1)).Return(models.Product{
		Name:  "abc",
		Price: 100,
	}, nil)

	cart := models.Cart{
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    1,
			},
		},
		Total: 100,
	}
	payload := models.CartRequestPayload{
		ProductID: 1,
		Quantity:  1,
	}
	CartRepository := cartRepository.NewCartRepository(mockProductRepository)
	err := CartRepository.DeleteProductFromCart(payload, &cart)
	assert.NoError(t, err)
	assert.Equal(t, cart.Total, int64(0))
	assert.Equal(t, len(cart.Items), 0)
}

func TestDeleteProductFromCartWithQuantityGreaterThanInCart(t *testing.T) {
	mockProductRepository := new(products.MockProductRepositoryInterface)
	mockProductRepository.On("GetProduct", uint(1)).Return(models.Product{
		Name:  "abc",
		Price: 100,
	}, nil)

	cart := models.Cart{
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    1,
			},
		},
		Total: 100,
	}
	payload := models.CartRequestPayload{
		ProductID: 1,
		Quantity:  3,
	}
	CartRepository := cartRepository.NewCartRepository(mockProductRepository)
	err := CartRepository.DeleteProductFromCart(payload, &cart)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Quantity of product 1 in cart is less than 3")
}

func TestDeleteProductFromCartWithProductNotInCart(t *testing.T) {
	mockProductRepository := new(products.MockProductRepositoryInterface)
	mockProductRepository.On("GetProduct", uint(1)).Return(models.Product{
		Name:  "abc",
		Price: 100,
	}, nil)
	mockProductRepository.On("GetProduct", uint(2)).Return(models.Product{
		Name:  "def",
		Price: 100,
	}, nil)

	cart := models.Cart{
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    1,
			},
		},
		Total: 100,
	}
	payload := models.CartRequestPayload{
		ProductID: 2,
		Quantity:  1,
	}
	CartRepository := cartRepository.NewCartRepository(mockProductRepository)
	err := CartRepository.DeleteProductFromCart(payload, &cart)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Product 2 is not in cart")
}
