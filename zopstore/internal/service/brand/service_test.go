package brand

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

// TestGetBrand is a test function which uses mocks to test GetBrand function
func TestGetBrand(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockBrandStorer(ctrl)

	tests := []struct {
		desc   string
		input  int
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  6,
			output: models.Brand{ID: 6, Name: "Bru"},
			expErr: nil,
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 6).
					Return(models.Brand{ID: 6, Name: "Bru"}, nil),
			}},
		{desc: "Fail",
			input:  99,
			output: models.Brand{},
			expErr: errors.EntityNotFound{},
			calls: []*gomock.Call{
				storeMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), 99).
					Return(models.Brand{}, errors.EntityNotFound{}),
			}},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)
		out, err := s.GetBrand(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestCreateBrand is a test function which uses mocks to test CreateBrand function
func TestCreateBrand(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockBrandStorer(ctrl)

	tests := []struct {
		desc   string
		input  models.Brand
		output models.Brand
		expErr error
		Call   []*gomock.Call
	}{
		{desc: "Success",
			input:  models.Brand{ID: 3, Name: "Nike"},
			output: models.Brand{ID: 3, Name: "Nike"},
			expErr: nil,
			Call: []*gomock.Call{
				storeMock.EXPECT().Create(gomock.AssignableToTypeOf(&gofr.Context{}), models.Brand{ID: 3, Name: "Nike"}).
					Return(models.Brand{ID: 3, Name: "Nike"}, nil),
			}},
		{desc: "Fail",
			input:  models.Brand{},
			output: models.Brand{},
			expErr: errors.MissingParam{Param: []string{"body"}},
			Call:   nil,
		},
		{desc: "Fail",
			input:  models.Brand{ID: 3, Name: "Nike"},
			output: models.Brand{},
			expErr: errors.MissingParam{Param: []string{"body"}},
			Call: []*gomock.Call{
				storeMock.EXPECT().Create(gomock.AssignableToTypeOf(&gofr.Context{}), models.Brand{ID: 3, Name: "Nike"}).
					Return(models.Brand{}, errors.MissingParam{Param: []string{"body"}}),
			},
		},
		{desc: "Fail",
			input:  models.Brand{ID: 3, Name: ""},
			output: models.Brand{},
			expErr: errors.MissingParam{Param: []string{"body"}},
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
		out, err := s.CreateBrand(ctx, val.input)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestUpdateBrand is a test function which uses mocks to test UpdateBrand function
func TestUpdateBrand(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	storeMock := store.NewMockBrandStorer(ctrl)

	tests := []struct {
		desc   string
		input1 int
		input2 models.Brand
		output models.Brand
		expErr error
		Calls  []*gomock.Call
	}{
		{desc: "Success",
			input1: 6,
			input2: models.Brand{ID: 6, Name: "bru"},
			output: models.Brand{ID: 6, Name: "bru"},
			expErr: nil,
			Calls: []*gomock.Call{
				storeMock.EXPECT().Update(gomock.AssignableToTypeOf(&gofr.Context{}), 6, models.Brand{ID: 6, Name: "bru"}).
					Return(models.Brand{ID: 6, Name: "bru"}, nil),
			}},
		{desc: "Fail",
			input1: 11,
			input2: models.Brand{ID: 11, Name: "example"},
			output: models.Brand{},
			expErr: errors.EntityNotFound{},
			Calls: []*gomock.Call{
				storeMock.EXPECT().Update(gomock.AssignableToTypeOf(&gofr.Context{}), 11, models.Brand{ID: 11, Name: "example"}).
					Return(models.Brand{}, errors.EntityNotFound{}),
			}},
		{desc: "Fail",
			input1: 6,
			input2: models.Brand{},
			output: models.Brand{},
			expErr: errors.MissingParam{Param: []string{"body"}},
			Calls:  nil,
		},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		s := New(storeMock)

		out, err := s.UpdateBrand(ctx, val.input1, val.input2)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}
