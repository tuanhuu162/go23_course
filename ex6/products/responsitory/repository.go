package responsitory

import (
	"github.com/jinzhu/gorm"
	"github.com/tuanhuu162/go23_course/ex6/models"
	"github.com/tuanhuu162/go23_course/ex6/products"
)

type ProductRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) products.ProductRepositoryInterface {
	return &ProductRepository{
		Conn: conn,
	}
}

func (ps *ProductRepository) GetProduct(product_id uint) (models.Product, error) {
	result := models.Product{}
	err := ps.Conn.First(&result, product_id).Error
	if err != nil {
		return models.Product{}, err
	} else {
		return result, nil
	}
}

func (ps *ProductRepository) GetAllProducts() []models.Product {
	result := []models.Product{}
	ps.Conn.Find(&result)
	return result
}

func (ps *ProductRepository) AddProduct(product models.Product) error {
	result := ps.Conn.Create(&product)
	return result.Error
}

func (ps *ProductRepository) DeleteProduct(product_id uint64) error {
	err := ps.Conn.First(&models.Product{}, product_id).Error
	if err != nil {
		return err
	} else {
		result := ps.Conn.Delete(&models.Product{}, product_id)
		return result.Error
	}
}

func (ps *ProductRepository) UpdateProduct(product models.Product) error {
	err := ps.Conn.First(&models.Product{}, product.ID).Error
	if err != nil {
		return err
	} else {
		result := ps.Conn.Save(&product)
		return result.Error
	}
}
