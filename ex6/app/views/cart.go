package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"

	"github.com/tuanhuu162/go23_course/ex6/app/models"
	"github.com/tuanhuu162/go23_course/ex6/app/storage"
	"github.com/tuanhuu162/go23_course/ex6/app/utils"
)

func handleCartCookies(ctx iris.Context, processing_function func(models.CartRequestPayload, *models.Cart) error) {
	var payload models.CartRequestPayload
	var cart = models.Cart{}
	err := ctx.ReadJSON(&payload)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	cart_cookie := ctx.GetCookie("cart")
	if cart_cookie != "" {
		cart_json, err := utils.Decrypt(cart_cookie)
		if err != nil {
			ctx.StopWithError(http.StatusInternalServerError, err)
			return
		}
		json.Unmarshal([]byte(cart_json), &cart)
	}
	process_err := processing_function(payload, *cart)
	if process_err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": fmt.Sprintf("Error process product %d to cart: %s", payload.ProductID, process_err)})
		return
	}
	cart_byte, encode_err := json.Marshal(cart)
	if encode_err != nil {
		ctx.StopWithError(http.StatusInternalServerError, encode_err)
	}
	ctx.SetCookieKV("cart", utils.Encrypt(cart_byte))
	ctx.JSON(iris.Map{
		"status": http.StatusOK,
		"data":   cart,
	})
}

// cart/add handles POST: http://localhost:8080/cart/add.
func AddProductToCart(ctx iris.Context) {
	handleCartCookies(ctx, storage.AddProductToCart)
}

// cart/delete handles DELETE: http://localhost:8080/cart/remove/1.
func DeleteProductFromCart(ctx iris.Context) {
	handleCartCookies(ctx, storage.DeleteProductFromCart)
}

// cart/checkout handles POST: http://localhost:8080/cart/checkout.
func CheckoutCart(ctx iris.Context) {
	var cart = models.Cart{}
	cart_cookie := ctx.GetCookie("cart")
	if cart_cookie != "" {
		cart_json, err := utils.Decrypt(cart_cookie)
		if err != nil {
			ctx.StopWithError(http.StatusInternalServerError, err)
			return
		}
		json.Unmarshal([]byte(cart_json), &cart)
	}

	ctx.SetCookieKV("cart", utils.Encrypt(json.Marshal(models.Cart{})))
	ctx.JSON(iris.Map{
		"status": http.StatusOK,
		"data":   cart,
	})

}
