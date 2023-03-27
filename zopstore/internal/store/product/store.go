package product

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

const val = "true"

// Get takes gofr context, id and brand as input
// Then query in the database and returns product details with brand name based on brand value and error if any
func (s *Store) Get(ctx *gofr.Context, id int, brand string) (models.Product, error) {
	var p models.Product

	resp := ctx.DB().QueryRowContext(ctx, "select products.id,products.name,products.description,products.price,"+
		"products.quantity,products.category,brands.id,brands.name,products.status "+
		"from products join brands on products.brand_id=brands.id where products.id=?", id)
	err := resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Brand.Name, &p.Status)

	if err != nil {
		return models.Product{}, errors.EntityNotFound{Entity: "product"}
	}

	if brand != val {
		p.Brand.Name = ""
	}

	return p, nil
}

// GetByName takes gofr context, name and brand as input
// Then query the database and returns list of products with brand name based on brand value and error if any
func (s *Store) GetByName(ctx *gofr.Context, name, brand string) ([]models.Product, error) {
	res := []models.Product{}

	resp, _ := ctx.DB().QueryContext(ctx,
		"select products.id,products.name,products.description,products.price,"+
			"products.quantity,products.category,brands.id,brands.name,products.status "+
			"from products join brands on products.brand_id=brands.id where products.name=?",
		name)
	if resp == nil {
		return []models.Product{{}}, errors.EntityNotFound{Entity: "product"}
	}

	for resp.Next() {
		var p models.Product
		_ = resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Brand.Name, &p.Status)

		if brand != val {
			p.Brand.Name = ""
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

	if prod.Brand.Name == "" {
		row := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", prod.Brand.ID)

		err := row.Scan(&prod.Brand.Name)
		if err != nil {
			return &models.Product{}, errors.EntityNotFound{Entity: "Brand Name"}
		}
	}

	return prod, nil
}

// Update takes gofr context,id and product details as input
// Then updates the product in database and returns product details and error if any
func (s *Store) Update(ctx *gofr.Context, id int, prod *models.Product) (*models.Product, error) {
	resp, err := ctx.DB().ExecContext(ctx,
		"update products set name=?,description=?,price=?,quantity=?,category=?,brand_id=?,status=? where id =?",
		prod.Name, prod.Description, prod.Price, prod.Quantity, prod.Category, prod.Brand.ID, prod.Status, id)

	if err != nil {
		return &models.Product{}, errors.EntityNotFound{Entity: "product"}
	}

	row, _ := resp.RowsAffected()

	if row == 0 {
		return &models.Product{}, errors.EntityNotFound{Entity: "product"}
	}

	prod.ID = id

	return prod, nil
}

// GetAll takes gofr context and brand as input
// Then returns list of all products with brand name based on brand and error if any
func (s *Store) GetAll(ctx *gofr.Context, brand string) ([]models.Product, error) {
	res := []models.Product{}

	resp, err := ctx.DB().QueryContext(ctx, "select products.id,products.name,products.description,products.price,"+
		"products.quantity,products.category,brands.id,brands.name,products.status "+
		"from products join brands on products.brand_id=brands.id")

	if err != nil {
		return nil, errors.EntityNotFound{Entity: "product"}
	}

	for resp.Next() {
		var prod models.Product
		_ = resp.Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.Quantity,
			&prod.Category, &prod.Brand.ID, &prod.Brand.Name, &prod.Status)

		if brand != val {
			prod.Brand.Name = ""
		}

		res = append(res, prod)
	}

	return res, nil
}
