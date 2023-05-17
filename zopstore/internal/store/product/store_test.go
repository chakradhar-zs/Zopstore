package product

import (
	"Day-19/internal/constants"
	"context"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"net/http"
	"net/http/httptest"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"Day-19/internal/models"
)

// TestGet is a test function which uses sql mocks to test Get function
func TestGet(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}
	defer db.Close()

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc    string
		input   int
		input2  string
		output  models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input:  3,
			input2: "true",
			output: models.Product{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available",
			},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Success",
			input:  5,
			input2: "false",
			output: models.Product{
				ID: 5, Name: "Bru", Description: "tasty", Price: 100, Quantity: 3, Category: "coffee",
				Brand: models.Brand{ID: 6}, Status: "Available",
			},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   333,
			input2:  "false",
			output:  models.Product{},
			mockErr: errors.EntityNotFound{Entity: "product"},
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
		{desc: "Fail",
			input:   22,
			input2:  "true",
			output:  models.Product{},
			mockErr: errors.EntityNotFound{Entity: "product"},
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
	}
	for _, val := range tests {
		ctx := context.WithValue(context.Background(), constants.CtxValue, "zs")
		r := httptest.NewRequest(http.MethodPost, "/", nil)
		r = r.WithContext(ctx)
		req := request.NewHTTPRequest(r)
		c := gofr.NewContext(nil, req, gofr.New())
		c.Context = r.Context()
		st := New()
		row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "bname", "status"}).
			AddRow(val.output.ID, val.output.Name, val.output.Description, val.output.Price,
				val.output.Quantity, val.output.Category, val.output.Brand.ID, val.output.Brand.Name, val.output.Status)
		mock.ExpectQuery("select").
			WithArgs(val.input).
			WillReturnRows(row).
			WillReturnError(val.mockErr)

		out, err := st.Get(c, val.input, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

// TestCreate is a test function which uses sql mocks to test Create function
func TestCreate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input   *models.Product
		output  *models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: "xyx"}, Status: "Available",
			},
			output: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: "xyx"}, Status: "Available",
			},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   &models.Product{},
			output:  &models.Product{},
			mockErr: errors.MissingParam{Param: []string{"body"}},
			expErr:  errors.MissingParam{Param: []string{"body"}},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		mock.ExpectExec("insert into").
			WithArgs(val.input.ID, val.input.Name, val.input.Description, val.input.Price, val.input.Quantity,
				val.input.Category, val.input.Brand.ID, val.input.Status).
			WillReturnResult(sqlmock.NewResult(6, 1)).
			WillReturnError(val.mockErr)
		mock.ExpectQuery("select").
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"bname"}).AddRow("xyx")).
			WillReturnError(nil)

		st := New()
		out, err := st.Create(ctx, val.input)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestCreateWithoutBrand(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	row := sqlmock.NewRows([]string{"bname"}).
		AddRow("maggi")
	row1 := sqlmock.NewRows([]string{"bname"}).
		AddRow(1)

	tests := []struct {
		desc   string
		input  *models.Product
		output *models.Product
		expErr error
		query1 *sqlmock.ExpectedExec
		query2 *sqlmock.ExpectedQuery
	}{
		{desc: "Success",
			input: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1}, Status: "Available",
			},
			output: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: "maggi"}, Status: "Available",
			},
			expErr: nil,
			query1: mock.ExpectExec("insert into").
				WithArgs(6, "maggi", "tasty", 50, 3, "noodles", 1, "Available").
				WillReturnResult(sqlmock.NewResult(6, 1)).
				WillReturnError(nil),
			query2: mock.ExpectQuery("select").
				WithArgs(1).
				WillReturnRows(row).
				WillReturnError(nil),
		},
		{desc: "Fail",
			input: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1}, Status: "Available",
			},
			output: &models.Product{},
			expErr: errors.EntityNotFound{Entity: "Brand Name"},
			query1: mock.ExpectExec("insert into").
				WithArgs(6, "maggi", "tasty", 50, 3, "noodles", 1, "Available").
				WillReturnResult(sqlmock.NewResult(6, 1)).
				WillReturnError(nil),
			query2: mock.ExpectQuery("select").
				WithArgs(1).
				WillReturnRows(row1).
				WillReturnError(errors.EntityNotFound{Entity: "Brand Name"}),
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		_ = val.query1
		_ = val.query2

		st := New()
		out, err := st.Create(ctx, val.input)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

// TestUpdate is a test function which uses sql mocks to test Update function
func TestUpdate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input1  int
		input2  *models.Product
		output  *models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input1: 6,
			input2: &models.Product{
				ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available",
			},
			output: &models.Product{
				ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available",
			},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input1:  2,
			input2:  &models.Product{},
			output:  &models.Product{},
			mockErr: errors.EntityNotFound{Entity: "product"},
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
		{desc: "Fail",
			input1: 333,
			input2: &models.Product{
				ID: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available",
			},
			output:  &models.Product{},
			mockErr: errors.EntityNotFound{Entity: "product"},
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		mock.ExpectExec("update ").
			WithArgs(val.input2.Name, val.input2.Description, val.input2.Price, val.input2.Quantity,
				val.input2.Category, val.input2.Brand.ID, val.input2.Status, val.input2.ID).
			WillReturnResult(sqlmock.NewResult(6, 1)).
			WillReturnError(val.mockErr)

		st := New()
		out, err := st.Update(ctx, val.input1, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

func TestUpdateInvalidId(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input1  int
		input2  *models.Product
		output  *models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Fail",
			input1: 1000,
			input2: &models.Product{
				ID: 1000, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available",
			},
			output:  &models.Product{},
			mockErr: nil,
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		mock.ExpectExec("update ").
			WithArgs(val.input2.Name, val.input2.Description, val.input2.Price, val.input2.Quantity,
				val.input2.Category, val.input2.Brand.ID, val.input2.Status, val.input2.ID).
			WillReturnResult(sqlmock.NewResult(1000, 0)).
			WillReturnError(nil)

		st := New()
		out, err := st.Update(ctx, val.input1, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

// TestGetAll is a test function which uses sql mocks to test GetAll function
func TestGetAll(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "bname", "status"}).
		AddRow(3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Nike", "Available")

	tests := []struct {
		desc    string
		input   string
		output  []models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input: "false",
			output: []models.Product{{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4}, Status: "Available",
			}},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:   "true",
			output:  nil,
			mockErr: errors.EntityNotFound{Entity: "product"},
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		if val.input == "false" {
			row = mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "bname", "status"}).
				AddRow(3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "", "Available")
		}

		mock.ExpectQuery("select").WillReturnRows(row).WillReturnError(val.mockErr)

		st := New()
		out, err := st.GetAll(ctx, val.input)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

// TestGetByName is a test function which uses sql mocks to test GetByName function
func TestGetByName(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	tests := []struct {
		desc    string
		input   string
		input2  string
		output  []models.Product
		mockErr error
		expErr  error
	}{
		{desc: "Success",
			input:  "zs_sneaker shoes",
			input2: "true",
			output: []models.Product{{
				ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available",
			}},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Success",
			input:  "zs_sneaker shoes",
			input2: "false",
			output: []models.Product{{
				ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4}, Status: "Available",
			}},
			mockErr: nil,
			expErr:  nil,
		},
		{desc: "Fail",
			input:  "bag",
			input2: "true",
			output: []models.Product{{
				ID: 0, Name: "", Description: "", Price: 0, Quantity: 0, Category: "",
				Brand: models.Brand{ID: 0, Name: ""}, Status: ""}},
			mockErr: errors.EntityNotFound{Entity: "product"},
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
		{desc: "Fail",
			input:  "chair",
			input2: "false",
			output: []models.Product{{
				ID: 0, Name: "", Description: "", Price: 0, Quantity: 0, Category: "",
				Brand: models.Brand{ID: 0, Name: ""}, Status: ""}},
			mockErr: errors.EntityNotFound{Entity: "product"},
			expErr:  errors.EntityNotFound{Entity: "product"},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()

		row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "bname", "status"}).
			AddRow(val.output[0].ID, val.output[0].Name, val.output[0].Description,
				val.output[0].Price, val.output[0].Quantity, val.output[0].Category,
				val.output[0].Brand.ID, val.output[0].Brand.Name, val.output[0].Status)

		if val.input2 == "false" {
			row = mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "bname", "status"}).
				AddRow(val.output[0].ID, val.output[0].Name, val.output[0].Description,
					val.output[0].Price, val.output[0].Quantity, val.output[0].Category, val.output[0].Brand.ID, "", val.output[0].Status)
		}

		mock.ExpectQuery("select").
			WithArgs(val.input).
			WillReturnRows(row).WillReturnError(val.mockErr)

		st := New()
		out, err := st.GetByName(ctx, val.input, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}

// TestByName is a test function which uses sql mocks to test GetByName function
func TestByName(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("Error while opening a mock db connection")
	}

	row := mock.NewRows([]string{"id", "name", "description", "price", "quantity", "category", "brand_id", "bname", "status"}).
		AddRow(0, "", "", 0, 0, "", 0, "", "")

	tests := []struct {
		desc   string
		input  string
		input2 string
		output []models.Product
		query  *sqlmock.ExpectedQuery
		expErr error
	}{
		{desc: "Fail",
			input:  "zs_nike",
			input2: "true",
			output: []models.Product{{
				ID: 0, Name: "", Description: "", Price: 0, Quantity: 0, Category: "",
				Brand: models.Brand{ID: 0, Name: ""}, Status: ""}},
			query: mock.ExpectQuery("select").WithArgs("zs_nike").
				WillReturnRows(row).WillReturnError(errors.EntityNotFound{Entity: "product"}),
			expErr: errors.EntityNotFound{Entity: "product"},
		},
	}

	for _, val := range tests {
		ctx.DataStore = datastore.DataStore{ORM: db}
		ctx.Context = context.Background()
		_ = val.query
		st := New()
		out, err := st.GetByName(ctx, val.input, val.input2)
		assert.Equal(t, val.output, out, "TEST failed.")
		assert.Equal(t, val.expErr, err, "TEST failed.")
	}
}
