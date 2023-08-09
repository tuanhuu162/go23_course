package main

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"

	"github.com/tuanhuu162/go23_course/ex6/app/database"
	"github.com/tuanhuu162/go23_course/ex6/app/models"
	"github.com/tuanhuu162/go23_course/ex6/app/views"
	// "github.com/iris-contrib/swagger/v12"
	// "github.com/iris-contrib/swagger/v12/swaggerFiles"
	// _ "github.com/tuanhuu162/go23_course/ex6/docs"
)

func LoggingMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func NewApp(database *database.Database) {
	app := iris.Default()

	// setup Database
	db, err := database.DB, database.Err
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

	// Some how swagger won't work, will investigate later
	// // setup swagger
	// config := swagger.Config{
	// 	// The url pointing to API definition.
	// 	URL:          "http://localhost:8080/swagger/doc.json",
	// 	DeepLinking:  true,
	// 	DocExpansion: "list",
	// 	DomID:        "#swagger-ui",
	// 	// The UI prefix URL (see route).
	// 	Prefix: "/swagger",
	// }
	// swaggerUI := swagger.Handler(swaggerFiles.Handler, config)

	// // Register on http://localhost:8080/swagger
	// app.Get("/swagger", swaggerUI)
	// // And the wildcard one for index.html, *.js, *.css and e.t.c.
	// app.Get("/swagger/{any:path}", swaggerUI)

	// setup middleware
	app.Use(LoggingMiddleware)
}

func main() {
	app := api.NewApp()
	app.Listen(":8080")
}
