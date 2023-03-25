package product

import (
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

const val = "true"

// Get takes gofr context, id and brand as input
// Then query in the database and returns product details with brand name based on brand value and error if any
func (s *Store) Get(ctx *gofr.Context, id int, brand string) (models.Product, error) {
	var p models.Product

	resp := ctx.DB().QueryRowContext(ctx, "select id,name,description,price,quantity,category,brand_id,status from products where id=?", id)
	fmt.Println(resp)
	err := resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Status)

	if err != nil {
		return models.Product{}, errors.EntityNotFound{Entity: "product"}
	}

	if brand == val {
		res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.ID)
		_ = res.Scan(&p.Brand.Name)
	}

	return p, nil
}

// GetByName takes gofr context, name and brand as input
// Then query the database and returns list of products with brand name based on brand value and error if any
func (s *Store) GetByName(ctx *gofr.Context, name, brand string) ([]models.Product, error) {
	res := []models.Product{}

	resp, _ := ctx.DB().QueryContext(ctx,
		"select id,name,description,price,quantity,category,brand_id,status from products where name=?",
		name)
	if resp == nil {
		return []models.Product{{}}, errors.EntityNotFound{Entity: "product"}
	}

	for resp.Next() {
		var p models.Product
		_ = resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Status)

		if brand == val {
			res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.ID)
			_ = res.Scan(&p.Brand.Name)
		}

		res = append(res, p)
	}

	return res, nil
}

// Create takes gofr context and product details as input
// Then insert into database and returns product details and error if any
func (s *Store) Create(ctx *gofr.Context, prod *models.Product) (*models.Product, error) {
	_, err := ctx.DB().ExecContext(ctx, "insert into products values(?,?,?,?,?,?,?,?)",
		prod.ID, prod.Name, prod.Description, prod.Price, prod.Quantity, prod.Category, prod.Brand.ID, prod.Status)

	if err != nil {
		return &models.Product{}, errors.MissingParam{Param: []string{"body"}}
	}

	return prod, nil
}

// Update takes gofr context,id and product details as input
// Then updates the product in database and returns product details and error if any
func (s *Store) Update(ctx *gofr.Context, id int, prod *models.Product) (*models.Product, error) {
	_, err := ctx.DB().ExecContext(ctx,
		"update products set name=?,description=?,price=?,quantity=?,category=?,brand_id=?,status=? where id =?",
		prod.Name, prod.Description, prod.Price, prod.Quantity, prod.Category, prod.Brand.ID, prod.Status, id)

	if err != nil {
		return &models.Product{}, errors.EntityNotFound{Entity: "product"}
	}

	prod.ID = id

	return prod, nil
}

// GetAll takes gofr context and brand as input
// Then returns list of all products with brand name based on brand and error if any
func (s *Store) GetAll(ctx *gofr.Context, brand string) ([]models.Product, error) {
	res := []models.Product{}

	resp, err := ctx.DB().QueryContext(ctx, "select * from products")

	if err != nil {
		return nil, errors.EntityNotFound{Entity: "product"}
	}

	for resp.Next() {
		var p models.Product
		_ = resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Status)

		if brand == val {
			res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.ID)
			_ = res.Scan(&p.Brand.Name)
		}

		res = append(res, p)
	}

	return res, nil
}
