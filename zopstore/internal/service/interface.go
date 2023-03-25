package service

import (
	"Day-19/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

// Product interface definition which is used in service layer
type Product interface {
	GetProduct(ctx *gofr.Context, i int, brand string) (models.Product, error)
	CreateProduct(ctx *gofr.Context, p *models.Product) (*models.Product, error)
	UpdateProduct(ctx *gofr.Context, i int, p *models.Product) (*models.Product, error)
	GetAllProducts(ctx *gofr.Context, brand string) ([]models.Product, error)
	GetProductByNAme(ctx *gofr.Context, name string, brand string) ([]models.Product, error)
}

// Brand interface definition which is used in service layer
type Brand interface {
	GetBrand(ctx *gofr.Context, id int) (models.Brand, error)
	CreateBrand(ctx *gofr.Context, brand models.Brand) (models.Brand, error)
	UpdateBrand(ctx *gofr.Context, id int, brand models.Brand) (models.Brand, error)
}
