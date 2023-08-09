package products

import (
	"github.com/tuanhuu162/go23_course/ex6/models"
)

//go:generate mockery --name ProductRepositoryInterfaceMock --output ~/work/go23_course/ex6/products/mock --outpkg mock
type ProductRepositoryInterface interface {
	GetProduct(product_id uint) (models.Product, error)
	GetAllProducts() []models.Product
	AddProduct(product models.Product) error
	DeleteProduct(product_id uint64) error
	UpdateProduct(product models.Product) error
}
