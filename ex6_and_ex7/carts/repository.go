package carts

import (
	"github.com/tuanhuu162/go23_course/ex6_and_ex7/models"
)

//go:generate mockery --name CartRepositoryInterfaceMock --output ~/work/go23_course/ex6_and_ex7/carts/mock --outpkg mock
type CartRepositoryInterface interface {
	AddProductToCart(payload models.CartRequestPayload, cart *models.Cart) error
	DeleteProductFromCart(payload models.CartRequestPayload, cart *models.Cart) error
}
