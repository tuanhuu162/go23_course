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
	process_err := processing_function(payload, &cart)
	if process_err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"message": fmt.Sprintf("Error process product %d to cart: %s", payload.ProductID, process_err)})
		return
	}
	cart_byte, encode_err := json.Marshal(cart)
	if encode_err != nil {
		ctx.StopWithError(http.StatusInternalServerError, encode_err)
	}
	encrypted_cart, encrypt_err := utils.Encrypt(cart_byte)
	if encrypt_err == nil {
		ctx.SetCookieKV("cart", encrypted_cart)
		ctx.JSON(iris.Map{
			"status": http.StatusOK,
			"data":   cart,
		})
	} else {
		ctx.StopWithError(http.StatusInternalServerError, encrypt_err)
		return
	}
}

// AddProductToCart godoc
//	@Summary	Add product to cart
//	@Description
//	@Tags		cart
//	@Accept		json
//	@Produce	json
//	@Param		product	body		models.CartRequestPayload	true	"CartRequestPayload"
//	@Success	200		{object}	models.Cart
//	@Failure	400		{object}	string
//	@Failure	404		{object}	string
//	@Failure	500		{object}	string
//	@Router		/cart/add [post]
func AddProductToCart(ctx iris.Context) {
	handleCartCookies(ctx, storage.AddProductToCart)
}

// DeleteProductFromCart godoc
//	@Summary	Delete product from cart
//	@Description
//	@Tags		cart
//	@Accept		json
//	@Produce	json
//	@Param		product	body		models.CartRequestPayload	true	"CartRequestPayload"
//	@Success	200		{object}	models.Cart
//	@Failure	400		{object}	string
//	@Failure	404		{object}	string
//	@Failure	500		{object}	string
//	@Router		/cart/remove [delete]
func DeleteProductFromCart(ctx iris.Context) {
	handleCartCookies(ctx, storage.DeleteProductFromCart)
}

// CheckoutCart godoc
//	@Summary	Checkout cart
//	@Description
//	@Tags		cart
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	models.Cart
//	@Failure	400	{object}	string
//	@Failure	404	{object}	string
//	@Failure	500	{object}	string
//	@Router		/cart/checkout [post]
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

	cart_byte, _ := json.Marshal(models.Cart{})
	encrypted_cart, _ := utils.Encrypt(cart_byte)
	ctx.SetCookieKV("cart", encrypted_cart)
	ctx.JSON(iris.Map{
		"status": http.StatusOK,
		"data":   cart,
	})

}
