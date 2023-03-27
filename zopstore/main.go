package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	productstore "Day-19/internal/store/product"

	productservice "Day-19/internal/service/product"

	brandservice "Day-19/internal/service/brand"
	brandstore "Day-19/internal/store/brand"

	"Day-19/internal/http/brand"
	"Day-19/internal/http/product"
	"Day-19/middleware"
)

// Dependency injection for product and brand
func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false

	app.Server.UseMiddleware(middleware.Middle, middleware.MiddleOrg)

	productStore := productstore.New()
	productSvc := productservice.New(productStore)
	prodHTTP := product.New(productSvc)

	brandStore := brandstore.New()
	brandSvc := brandservice.New(brandStore)
	brandHTTP := brand.New(brandSvc)

	app.REST("products", prodHTTP)
	app.REST("brands", brandHTTP)
	app.Start()
}
