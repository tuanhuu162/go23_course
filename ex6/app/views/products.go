package views

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/tuanhuu162/go23_course/ex6/app/models"
	"github.com/tuanhuu162/go23_course/ex6/app/storage"
)

// ListProducts handles GET: http://localhost:8080/products.
func ListProducts(ctx iris.Context) {
	ctx.JSON(storage.GetAllProducts())
}

// createProduct handles POST: http://localhost:8080/products.
func createProduct(ctx iris.Context) {
	var product models.Product
	err := ctx.ReadJSON(&product)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	create_product_err := storage.AddProduct(product)
	if create_product_err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": "error creating product"})
		return
	}
	ctx.JSON(iris.Map{"message": "created"})
}

// deleteProduct handles DELETE: http://localhost:8080/products/1.
func deleteProduct(ctx iris.Context) {
	product_id, err := ctx.Params().GetUint64("product_id")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	delete_product_error := storage.DeleteProduct(product_id)
	if delete_product_error != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": fmt.Sprintf("error deleting product %s", product_id)})
		return
	}
	ctx.JSON(iris.Map{"message": fmt.Sprintf("Deleted product %s", product_id)})
}

// updateProduct handles PUT: http://localhost:8080/products/1.
func updateProduct(ctx iris.Context) {
	var product models.Product
	err := ctx.ReadJSON(&product)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	update_product_error := storage.UpdateProduct(product)
	if update_product_error != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": fmt.Sprintf("error updating product %s", &product.ID)})
		return
	}
	ctx.JSON(iris.Map{"message": "updated"})

}
