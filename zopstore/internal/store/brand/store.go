package brand

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

// Get takes gofr context, id as input and query through database and returns brand details and error if any
func (s *Store) Get(ctx *gofr.Context, id int) (models.Brand, error) {
	var b models.Brand

	resp := ctx.DB().QueryRowContext(ctx, "select id,name from brands where id=?", id)

	err := resp.Scan(&b.ID, &b.Name)

	if err != nil {
		return models.Brand{}, errors.EntityNotFound{Entity: "brand", ID: "id"}
	}

	return b, nil
}

// Create takes gofr context, brand details as input and insert in database and returns brand details and error if any
func (s *Store) Create(ctx *gofr.Context, brand models.Brand) (models.Brand, error) {
	_, err := ctx.DB().ExecContext(ctx, "insert into brands values(?,?)", brand.ID, brand.Name)

	if err != nil {
		return models.Brand{}, errors.MissingParam{Param: []string{"body"}}
	}

	return brand, nil
}

// Update takes gofr context, id and brand detailsas input and update brand details database and returns brand details and error if any
func (s *Store) Update(ctx *gofr.Context, id int, brand models.Brand) (models.Brand, error) {
	_, err := ctx.DB().ExecContext(ctx, "update brands set name=? where id=?", brand.Name, id)

	if err != nil {
		return models.Brand{}, errors.EntityNotFound{Entity: "brand", ID: "id"}
	}

	brand.ID = id

	return brand, nil
}
