package views

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/tuanhuu162/go23_course/ex6/app/models"
	"github.com/tuanhuu162/go23_course/ex6/app/storage"
)

// ListProduct godoc
//	@Summary	Show all Product we have
//	@Description
//	@Tags		products
//	@Produce	json
//	@Success	200	{object}	[]models.ProductSerializer
//	@Failure	400	{object}	string
//	@Failure	404	{object}	string
//	@Failure	500	{object}	string
//	@Router		/products/ [get]
func ListProducts(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status": http.StatusOK,
		"data":   storage.GetAllProducts(),
	})
}

// CreateProduct godoc
//	@Summary	Create product with name and price
//	@Description
//	@Tags		products
//	@Accept		json
//	@Produce	json
//	@Param		product	body		models.Product	true	"Product"
//	@Success	200		{object}	string
//	@Failure	400		{object}	string
//	@Failure	404		{object}	string
//	@Failure	500		{object}	string
//	@Router		/products/ [post]
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

// DeleteProduct godoc
//	@Summary	Delete product with id
//	@Description
//	@Tags		products
//	@Produce	json
//	@Param		product_id	path		int	true	"Product ID"
//	@Success	200			{object}	string
//	@Failure	400			{object}	string
//	@Failure	404			{object}	string
//	@Failure	500			{object}	string
//	@Router		/products/{product_id} [delete]
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

// UpdateProduct godoc
//	@Summary	Update product with id
//	@Description
//	@Tags		products
//	@Accept		json
//	@Produce	json
//	@Param		product	body		models.Product	true	"Product"
//	@Success	200		{object}	string
//	@Failure	400		{object}	string
//	@Failure	404		{object}	string
//	@Failure	500		{object}	string
//	@Router		/products/ [put]
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
