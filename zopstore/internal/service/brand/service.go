package brand

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
	"Day-19/internal/store"
)

type Service struct {
	store store.BrandStorer
}

func New(storer store.BrandStorer) *Service {
	return &Service{store: storer}
}

// GetBrand function takes id and gofr context as input and calls Get of store layer
func (svc *Service) GetBrand(ctx *gofr.Context, id int) (models.Brand, error) {
	res, err := svc.store.Get(ctx, id)
	if err != nil {
		return models.Brand{}, err
	}

	return res, nil
}

// CreateBrand takes gofr context and brand structure as input
// Then checks for missing fields and calls Create of store layer which returns brand details and error if any
func (svc *Service) CreateBrand(ctx *gofr.Context, brand models.Brand) (models.Brand, error) {
	err := isEmpty(brand)
	if err != nil {
		return models.Brand{}, err
	}

	if brand.ID == 0 {
		return models.Brand{}, errors.MissingParam{Param: []string{"id"}}
	}

	res, err := svc.store.Create(ctx, brand)
	if err != nil {
		return models.Brand{}, err
	}

	return res, nil
}

// UpdateBrand takes gofr context ,id and brand structure as input
// Then checks for missing fields and calls Update of store layer which returns brand details and error if any
func (svc *Service) UpdateBrand(ctx *gofr.Context, id int, brand models.Brand) (models.Brand, error) {
	err := isEmpty(brand)
	if err != nil {
		return models.Brand{}, err
	}

	res, err := svc.store.Update(ctx, id, brand)

	if err != nil {
		return models.Brand{}, err
	}

	return res, nil
}

func isEmpty(b models.Brand) error {
	if b.Name == "" {
		return errors.MissingParam{Param: []string{"name"}}
	}

	return nil
}
