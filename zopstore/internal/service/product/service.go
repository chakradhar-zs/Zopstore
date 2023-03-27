package product

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
	"Day-19/internal/store"
)

type Service struct {
	store store.ProductStorer
}

func New(storer store.ProductStorer) *Service {
	return &Service{store: storer}
}

// GetProduct takes gofr context,id and brand as input and calls Get of store layer and returns product details and error
func (s *Service) GetProduct(ctx *gofr.Context, i int, brand string) (models.Product, error) {
	res, err := s.store.Get(ctx, i, brand)

	if err != nil {
		return models.Product{}, err
	}

	return res, nil
}

// GetProductByNAme takes gofr context , name and brand as input and calls GetByName of store layer and returns array of products and error
func (s *Service) GetProductByNAme(ctx *gofr.Context, name, brand string) ([]models.Product, error) {
	res, err := s.store.GetByName(ctx, name, brand)

	if err != nil {
		return []models.Product{}, err
	}

	return res, nil
}

// CreateProduct takes gofr context and product details as input
// Then checks for missing fields and calls Create of store layer
// Returns product details affected and error
func (s *Service) CreateProduct(ctx *gofr.Context, p *models.Product) (*models.Product, error) {
	err := isEmpty(p)
	if err != nil {
		return &models.Product{}, err
	}

	if p.ID == 0 {
		return &models.Product{}, errors.MissingParam{Param: []string{"id"}}
	}

	res, err := s.store.Create(ctx, p)

	if err != nil {
		return &models.Product{}, errors.MissingParam{Param: []string{"body"}}
	}

	return res, nil
}

// UpdateProduct takes gofr context, id and product structure as input
// Then checks for missing fields and calls Update of store layer
// Returns product details and error
func (s *Service) UpdateProduct(ctx *gofr.Context, id int, p *models.Product) (*models.Product, error) {
	err := isEmpty(p)
	if err != nil {
		return &models.Product{}, err
	}

	res, err := s.store.Update(ctx, id, p)

	if err != nil {
		return &models.Product{}, err
	}

	return res, nil
}

// GetAllProducts takes gofr context and brand as input
// Then calls GetAll of store layer and returns a list of all products and error
func (s *Service) GetAllProducts(ctx *gofr.Context, brand string) ([]models.Product, error) {
	res, err := s.store.GetAll(ctx, brand)

	if err != nil {
		return []models.Product{}, err
	}

	return res, nil
}

func isEmpty(b *models.Product) error {
	if b.Name == "" {
		return errors.MissingParam{Param: []string{"name"}}
	} else if b.Description == "" {
		return errors.MissingParam{Param: []string{"description"}}
	} else if b.Price == 0 {
		return errors.MissingParam{Param: []string{"price"}}
	} else if b.Quantity == 0 {
		return errors.MissingParam{Param: []string{"quantity"}}
	} else if b.Category == "" {
		return errors.MissingParam{Param: []string{"category"}}
	} else if b.Brand.ID == 0 {
		return errors.MissingParam{Param: []string{"Brand Id"}}
	} else if b.Status == "" {
		return errors.MissingParam{Param: []string{"Status"}}
	}

	return nil
}
