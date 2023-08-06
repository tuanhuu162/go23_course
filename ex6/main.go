package main

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"

	"github.com/tuanhuu162/go23_course/ex6/app/database"
	"github.com/tuanhuu162/go23_course/ex6/app/models"
	"github.com/tuanhuu162/go23_course/ex6/app/views"
)

func LoggingMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func main() {
	app := iris.Default()

	// setup Database
	db, err := database.New()
	db.LogMode(true) // show SQL logger

	if err != nil {
		app.Logger().Fatalf("connect to sqlite3 failed")
		return
	}

	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})

	if os.Getenv("ENV") != "" {
		db.DropTableIfExists(&models.Product{}) // drop table
	}
	db.AutoMigrate(&models.Product{}) // create table: // AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data

	// setup authentication
	auth := basicauth.Default(map[string]string{
		"admin": "admin",
	})

	// setup views
	productsAPI := app.Party("/products", auth)
	{
		productsAPI.Use(iris.Compression)

		// GET: http://localhost:8080/products
		productsAPI.Get("/", views.ListProducts)
		// POST: http://localhost:8080/products
		productsAPI.Post("/", views.CreateProduct)
		// DELETE: http://localhost:8080/products/1
		productsAPI.Delete("/{product_id:uint64}", views.DeleteProduct)
		// PUT: http://localhost:8080/products/1
		productsAPI.Put("/", views.UpdateProduct)
	}
	// cartAPI := app.Party("/cart")
	// {
	// 	cartAPI.Use(iris.Compression)

	// 	// POST: http://localhost:8080/cart/add
	// 	cartAPI.Post("/add", views.AddProductToCart)
	// 	// DELETE: http://localhost:8080/cart/remove/1
	// 	cartAPI.Delete("/remove", views.DeleteProductFromCart)
	// 	// POST: http://localhost:8080/cart/checkout
	// 	cartAPI.Post("/checkout", views.CheckoutCart)
	// }

	// setup middleware
	app.Use(LoggingMiddleware)

	app.Listen(":8080")
}
