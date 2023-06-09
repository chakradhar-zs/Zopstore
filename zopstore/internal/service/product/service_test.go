package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"Day-19/internal/models"
	"Day-19/internal/store"
)

// TestGetProduct is a test function which uses mocks to test GetProduct
func TestGetProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  int
		input2 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{
			desc:   "Success",
			input:  3,
			input2: "true",
			output: models.Product{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 3, "true").
					Return(models.Product{
						ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}, nil),
			}},
		{desc: "Fail",
			input:  333,
			input2: "true",
			output: models.Product{},
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 333, "true").
					Return(models.Product{}, errors.EntityNotFound{}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetProduct(ctx, val.input, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestGet is a test function which uses mocks to test GetProduct
func TestGet(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  int
		input2 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  4,
			input2: "false",
			output: models.Product{
				ID: 4, Name: "maggi", Description: "yum", Price: 100, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 5}, Status: "Available"},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 4, "false").
					Return(models.Product{
						ID: 4, Name: "maggi", Description: "yum", Price: 100, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 5}, Status: "Available"}, nil),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetProduct(ctx, val.input, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestCreateProduct is a test function which uses mocks to test CreateProduct
func TestCreateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  *models.Product
		output *models.Product
		expErr error
		Call   []*gomock.Call
	}{
		{desc: "Success",
			input: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			expErr: nil,
			Call: []*gomock.Call{
				storeMock.EXPECT().
					Create(gomock.AssignableToTypeOf(&gofr.Context{}), &models.Product{
						ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(&models.Product{
						ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}, nil),
			}},
		{desc: "Fail",
			input: &models.Product{
				ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"body"}},
			Call: []*gomock.Call{
				storeMock.EXPECT().
					Create(gomock.AssignableToTypeOf(&gofr.Context{}), &models.Product{
						ID: 6, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(&models.Product{}, errors.MissingParam{Param: []string{"body"}}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.CreateProduct(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestCreateWithInvalidBody is a test function which uses mocks to test CreateProduct with invalid body
func TestCreateWithInvalidBody(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input  *models.Product
		output *models.Product
		expErr error
		Call   []*gomock.Call
	}{
		{desc: "Fail",
			input: &models.Product{
				ID: 1, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 0, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"Brand Id"}},
			Call:   nil,
		},
		{desc: "Fail",
			input: &models.Product{
				ID: 1, Name: "maggi", Description: "tasty", Price: 0, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"price"}},
			Call:   nil,
		},
		{desc: "Fail",
			input: &models.Product{
				ID: 1, Name: "maggi", Description: "tasty", Price: 50, Quantity: 3, Category: "",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"category"}},
			Call:   nil,
		},
		{desc: "Fail",
			input: &models.Product{
				ID: 1, Name: "", Description: "tasty", Price: 50, Quantity: 3, Category: "",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"name"}},
			Call:   nil,
		},
		{desc: "Fail",
			input: &models.Product{
				ID: 1, Name: "maggi", Description: "", Price: 50, Quantity: 3, Category: "",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"description"}},
			Call:   nil,
		},
		{desc: "Fail",
			input: &models.Product{
				ID: 1, Name: "maggi", Description: "tasty", Price: 50, Quantity: 0, Category: "",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"quantity"}},
			Call:   nil,
		},
		{desc: "Fail",
			input: &models.Product{
				ID: 0, Name: "maggi", Description: "tasty", Price: 50, Quantity: 1, Category: "xyx",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"id"}},
			Call:   nil,
		},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.CreateProduct(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestUpdateProduct is a test function which uses mocks to test UpdateProduct
func TestUpdateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		input1 int
		input2 *models.Product
		output *models.Product
		expErr error
		Calls  []*gomock.Call
	}{
		{desc: "Success",
			input1: 6,
			input2: &models.Product{
				ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{
				ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			expErr: nil,
			Calls: []*gomock.Call{
				storeMock.EXPECT().
					Update(gomock.AssignableToTypeOf(&gofr.Context{}), 6, &models.Product{
						ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(&models.Product{
						ID: 6, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}, nil),
			}},
		{desc: "Fail",
			input1: 6,
			input2: &models.Product{
				ID: 6, Name: "maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: ""},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"Status"}},
			Calls:  nil,
		},
		{desc: "Fail",
			input1: 2,
			input2: &models.Product{},
			output: &models.Product{},
			expErr: errors.MissingParam{Param: []string{"name"}},
			Calls:  nil,
		},
		{desc: "Fail",
			input1: 333,
			input2: &models.Product{
				ID: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
				Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"},
			output: &models.Product{},
			expErr: errors.EntityNotFound{Entity: "id"},
			Calls: []*gomock.Call{
				storeMock.EXPECT().
					Update(gomock.AssignableToTypeOf(&gofr.Context{}), 333, &models.Product{
						ID: 333, Name: "Maggi", Description: "yummy", Price: 50, Quantity: 3, Category: "noodles",
						Brand: models.Brand{ID: 1, Name: ""}, Status: "Available"}).
					Return(&models.Product{}, errors.EntityNotFound{Entity: "id"}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.UpdateProduct(ctx, val.input1, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestGetAllProducts is a test function which uses mocks to test GetAllProducts
func TestGetAllProducts(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		path   string
		input  string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			path:  "/product/?brand=true",
			input: "true",
			output: []models.Product{{
				ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().GetAll(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return([]models.Product{{
						ID: 3, Name: "sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}}, nil),
			}},
		{desc: "Fail",
			path:   "/product/?brand=true",
			input:  "true",
			output: []models.Product{},
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().GetAll(gomock.AssignableToTypeOf(&gofr.Context{}), "true").
					Return([]models.Product{}, errors.EntityNotFound{}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, val.path, nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetAllProducts(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestGetProductByName is a test function which uses mocks to test GetProductByName
func TestGetProductByName(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockProductStorer(ctrl)
	tests := []struct {
		desc   string
		path   string
		input1 string
		input2 string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			path:   "/product/?brand=true",
			input1: "zs_sneaker shoes",
			input2: "true",
			output: []models.Product{{
				ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
				Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().GetByName(gomock.AssignableToTypeOf(&gofr.Context{}), "zs_sneaker shoes", "true").
					Return([]models.Product{{
						ID: 3, Name: "zs_sneaker shoes", Description: "stylish", Price: 1000, Quantity: 3, Category: "shoes",
						Brand: models.Brand{ID: 4, Name: "Nike"}, Status: "Available"}}, nil),
			}},
		{desc: "Fail",
			path:   "/product/?brand=true",
			input1: "",
			input2: "true",
			output: []models.Product{},
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().GetByName(gomock.AssignableToTypeOf(&gofr.Context{}), "", "true").
					Return([]models.Product{}, errors.EntityNotFound{}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, val.path, nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetProductByNAme(ctx, val.input1, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}
