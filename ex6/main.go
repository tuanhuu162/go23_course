package main

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/tuanhuu162/go23_course/ex6/app/database"
	"github.com/tuanhuu162/go23_course/ex6/app/storage"
	"github.com/tuanhuu162/go23_course/ex6/app/views"
)

func main() {
	app := iris.Default()

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
		db.DropTableIfExists(&storage.Product{}) // drop table
	}
	db.AutoMigrate(&storage.Product{}) // create table: // AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data

	productsAPI := app.Party("/products")
	{
		productsAPI.Use(iris.Compression)

		// GET: http://localhost:8080/products
		productsAPI.Get("/", views.ListProducts)
		// POST: http://localhost:8080/products
		productsAPI.Post("/", views.createProduct)
		// DELETE: http://localhost:8080/products/1
		productsAPI.Delete("/{id:uint64}", views.deleteProduct)
		// PUT: http://localhost:8080/products/1
		productsAPI.Put("/{id:uint64}", views.updateProduct)
	}
	cartAPI := app.Party("/cart")
	{
		cartAPI.Use(iris.Compression)

		// POST: http://localhost:8080/cart/add
		cartAPI.Post("/add", views.AddProductToCart)
		// DELETE: http://localhost:8080/cart/remove/1
		cartAPI.Delete("/remove", views.DeleteProductFromCart)
		// POST: http://localhost:8080/cart/checkout
		cartAPI.Post("/checkout", views.CheckoutCart)
	}
	app.Listen(":8080")
}
