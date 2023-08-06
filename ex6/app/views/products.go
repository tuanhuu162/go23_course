package views

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/tuanhuu162/go23_course/ex6/app/models"
	"github.com/tuanhuu162/go23_course/ex6/app/storage"
)

// ListProducts handles GET: http://localhost:8080/products.
func ListProducts(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status": http.StatusOK,
		"data":   storage.GetAllProducts(),
	})
}

// createProduct handles POST: http://localhost:8080/products.
func CreateProduct(ctx iris.Context) {
	var product models.Product
	err := ctx.ReadJSON(&product)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	create_product_err := storage.AddProduct(product)
	if create_product_err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": fmt.Sprintf("Error creating product: %s", create_product_err)})
		return
	}
	ctx.JSON(iris.Map{"message": "created"})
}

// deleteProduct handles DELETE: http://localhost:8080/products/1.
func DeleteProduct(ctx iris.Context) {
	product_id, err := ctx.Params().GetUint64("product_id")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	delete_product_error := storage.DeleteProduct(product_id)
	if delete_product_error != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": fmt.Sprintf("Error deleting product %d: %s", product_id, delete_product_error)})
		return
	}
	ctx.JSON(iris.Map{"message": fmt.Sprintf("Deleted product %d", product_id)})
}

// updateProduct handles PUT: http://localhost:8080/products/1.
func UpdateProduct(ctx iris.Context) {
	var product models.Product
	err := ctx.ReadJSON(&product)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	update_product_error := storage.UpdateProduct(product)
	if update_product_error != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": fmt.Sprintf("Error updating product %d: %s", product.ID, update_product_error)})
		return
	}
	ctx.JSON(iris.Map{"message": "updated"})

}
