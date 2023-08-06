package views

// import (
// 	"github.com/kataras/iris/v12"
// 	"github.com/kataras/iris/v12/sessions"

// 	"github.com/tuanhuu162/go23_course/ex6/app/models"
// )

// var (
// 	cookieNameForSessionID = "mycookiesessionnameid"
// 	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
// )

// // cart/add handles POST: http://localhost:8080/cart/add.
// func AddProductToCart(ctx iris.Context) {
// 	var item models.Item
// 	err := ctx.ReadJSON(&item)
// 	if err != nil {
// 		ctx.StopWithError(iris.StatusBadRequest, err)
// 		return
// 	}
// 	if ProductData, ok := ctx.Values().Get("ProductData").(ProductData); ok {
// 		if _, err := ProductData.GetProduct(item.ProductID); err != nil {
// 			ctx.StopWithError(iris.StatusBadRequest, err)
// 			return
// 		}
// 		ProductData.AddItemToCart(item.ProductID, item.Quantity)
// 		ctx.JSON(ProductData.GetCart())
// 	}
// }

// // cart/delete handles DELETE: http://localhost:8080/cart/remove/1.
// func DeleteProductFromCart(ctx iris.Context) {
// 	var payload models.CartRemoveRequestPayload
// 	err := ctx.ReadJSON(&payload)
// 	if err != nil {
// 		ctx.StopWithError(iris.StatusBadRequest, err)
// 		return
// 	}
// 	if ProductData, ok := ctx.Values().Get("ProductData").(ProductData); ok {
// 		if _, err := ProductData.GetProduct(payload.product_id); err != nil {
// 			ctx.StopWithError(iris.StatusBadRequest, err)
// 			return
// 		}
// 		ProductData.DeleteItemFromCart(payload.product_id)
// 		ctx.JSON(ProductData.GetCart())
// 	}
// }

// // cart/checkout handles POST: http://localhost:8080/cart/checkout.
// func CheckoutCart(ctx iris.Context) {

// }
