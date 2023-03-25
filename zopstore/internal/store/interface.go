package store

import (
	"Day-19/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

// ProductStorer interface definition which is used in store layer
type ProductStorer interface {
	Get(ctx *gofr.Context, id int, brand string) (models.Product, error)
	Create(ctx *gofr.Context, prod *models.Product) (*models.Product, error)
	Update(ctx *gofr.Context, id int, prod *models.Product) (*models.Product, error)
	GetAll(ctx *gofr.Context, brand string) ([]models.Product, error)
	GetByName(ctx *gofr.Context, name string, brand string) ([]models.Product, error)
}

// BrandStorer interface definition which is used in store layer
type BrandStorer interface {
	Get(ctx *gofr.Context, id int) (models.Brand, error)
	Create(ctx *gofr.Context, brand models.Brand) (models.Brand, error)
	Update(ctx *gofr.Context, id int, brand models.Brand) (models.Brand, error)
}
