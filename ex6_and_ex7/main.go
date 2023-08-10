package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	cartRepository "github.com/tuanhuu162/go23_course/ex6_and_ex7/carts/repository"
	cartTransporter "github.com/tuanhuu162/go23_course/ex6_and_ex7/carts/transporter"
	"github.com/tuanhuu162/go23_course/ex6_and_ex7/models"
	productRepository "github.com/tuanhuu162/go23_course/ex6_and_ex7/products/repository"
	productTransporter "github.com/tuanhuu162/go23_course/ex6_and_ex7/products/transporter"
	// "github.com/iris-contrib/swagger/v12"
	// "github.com/iris-contrib/swagger/v12/swaggerFiles"
	// _ "github.com/tuanhuu162/go23_course/ex6_and_ex7/docs"
)

func LoggingMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func main() {
	app := iris.Default()

	// setup Database
	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})

	if err != nil {
		app.Logger().Fatalf("connect to sqlite3 failed")
		return
	}

	db_connection, dberr := db.DB()

	if dberr != nil {
		app.Logger().Fatalf("get db failed")
		return
	}

	iris.RegisterOnInterrupt(func() {
		defer db_connection.Close()
	})

	db.AutoMigrate(&models.Product{}) // create table: // AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data

	// setup authentication
	auth := basicauth.Default(map[string]string{
		"admin": "admin",
	})

	// setup handlers
	productAPI := app.Party("/products", auth)

	productRepository := productRepository.NewProductRepository(db)
	productTransporter.NewProductTransport(productAPI, productRepository)

	cartAPI := app.Party("/cart")
	cartRepository := cartRepository.NewCartRepository(productRepository)
	cartTransporter.NewCartTransport(cartAPI, cartRepository)

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
	app.Listen(":8080")
}
