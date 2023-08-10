package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"

	"github.com/tuanhuu162/go23_course/ex6/models"
	"github.com/tuanhuu162/go23_course/ex6/products"
	prodcutRepo "github.com/tuanhuu162/go23_course/ex6/products/repository"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	repo products.ProductRepositoryInterface
}

func (su *ProductRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to open gorm v2 db, got error: %v", err)
	}
	db.AutoMigrate(&models.Product{})
	su.repo = prodcutRepo.NewProductRepository(db)
}

func (su *ProductRepositoryTestSuite) TestProductRepository_AddProduct() {
	add_err := su.repo.AddProduct(models.Product{
		Name:  "test",
		Price: 1,
	})
	defer su.repo.DeleteProduct(1)
	assert.Nil(su.T(), add_err)
	product := su.repo.GetAllProducts()
	assert.Equal(su.T(), product[0].Name, "test")
	assert.Equal(su.T(), product[0].Price, int64(1))

}

func (su *ProductRepositoryTestSuite) TestProductRepository_GetProduct() {
	su.repo.AddProduct(models.Product{
		Name:  "test",
		Price: 1,
	})
	defer su.repo.DeleteProduct(1)
	product, err := su.repo.GetProduct(1)
	assert.Nil(su.T(), err)
	assert.Equal(su.T(), product.ID, uint(1))
	assert.Equal(su.T(), product.Name, "test")
}

func (su *ProductRepositoryTestSuite) TestProductRepository_GetAllProducts() {
	su.repo.AddProduct(models.Product{
		Name:  "test",
		Price: 1,
	})
	defer su.repo.DeleteProduct(1)
	products := su.repo.GetAllProducts()
	assert.Equal(su.T(), products[0].ID, uint(1))
	assert.Equal(su.T(), products[0].Name, "test")
	assert.Equal(su.T(), len(products), 1)
}

func (su *ProductRepositoryTestSuite) TestProductRepository_DeleteProduct() {
	su.repo.AddProduct(models.Product{
		Name:  "test",
		Price: 1,
	})
	err := su.repo.DeleteProduct(1)
	assert.Nil(su.T(), err)
	products := su.repo.GetAllProducts()
	assert.Equal(su.T(), len(products), 0)
}

func (su *ProductRepositoryTestSuite) TestProductRepository_UpdateProduct() {
	su.repo.AddProduct(models.Product{
		Name:  "test",
		Price: 1,
	})
	defer su.repo.DeleteProduct(1)
	input_product, _ := su.repo.GetProduct(1)
	input_product.Name = "test2"
	input_product.Price = 2
	err := su.repo.UpdateProduct(input_product)
	assert.Nil(su.T(), err)
	product, err := su.repo.GetProduct(1)
	assert.Equal(su.T(), product.ID, uint(1))
	assert.Equal(su.T(), product.Name, "test2")
	assert.Equal(su.T(), product.Price, int64(2))

}

func TestRunAll(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}
