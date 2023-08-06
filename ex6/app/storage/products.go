package storage

import (
	"github.com/tuanhuu162/go23_course/ex6/app/database"
	"github.com/tuanhuu162/go23_course/ex6/app/models"
)

func GetProduct(product_id string) (models.Product, error) {
	result := models.Product{}
	err := database.DB.First(&result, product_id).Error
	if err != nil {
		return models.Product{}, err
	} else {
		return result, nil
	}
}

func GetAllProducts() []models.Product {
	result := []models.Product{}
	database.DB.Find(&result)
	return result
}

func AddProduct(product models.Product) error {
	result := database.DB.Create(&product)
	return result.Error
}

func DeleteProduct(product_id uint64) error {
	result := database.DB.Delete(&models.Product{}, product_id)
	return result.Error
}

func UpdateProduct(product models.Product) error {
	result := database.DB.Save(&product)
	return result.Error
}
