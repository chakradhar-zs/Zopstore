package main

import (
	"zopstore/internal/http/brand"
	"zopstore/internal/http/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	productstore "zopstore/internal/store/product"

	productservice "zopstore/internal/service/product"

	brandservice "zopstore/internal/service/brand"
	brandstore "zopstore/internal/store/brand"
)

func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false

	app.Server.UseMiddleware(product.Middle)

	productStore := productstore.New()
	productSvc := productservice.New(productStore)
	prodHTTP := product.New(productSvc)

	brandStore := brandstore.New()
	brandSvc := brandservice.New(brandStore)
	brandHTTP := brand.New(brandSvc)

	app.REST("product", prodHTTP)
	app.REST("brand", brandHTTP)
	app.Start()
}
