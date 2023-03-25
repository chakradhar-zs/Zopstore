package brand

import (
	"bytes"
	"encoding/json"
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
	"Day-19/internal/service"
)

// TestRead is a test function which uses mocks to test Read handler
func TestRead(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockBrand(ctrl)

	tests := []struct {
		desc   string
		input  string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  "6",
			output: models.Brand{ID: 6, Name: "Bru"},
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().GetBrand(gomock.AssignableToTypeOf(&gofr.Context{}), 6).Return(models.Brand{ID: 6, Name: "Bru"}, nil),
			}},
		{desc: "Fail",
			input:  "99",
			output: models.Brand{},
			expErr: errors.EntityNotFound{Entity: "brand"},
			calls: []*gomock.Call{
				serviceMock.EXPECT().GetBrand(gomock.AssignableToTypeOf(&gofr.Context{}), 99).Return(models.Brand{}, errors.EntityNotFound{}),
			}},
		{desc: "Fail",
			input:  "",
			output: models.Brand{},
			expErr: errors.MissingParam{Param: []string{"id"}},
			calls:  nil,
		},
		{desc: "Fail",
			input:  "abc",
			output: models.Brand{},
			expErr: errors.InvalidParam{Param: []string{"id"}},
			calls:  nil,
		},
	}

	for i, val := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": val.input,
		})

		h := New(serviceMock)
		out, err := h.Read(ctx)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestCreate is a test function which uses mocks to test Create handler
func TestCreate(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockBrand(ctrl)

	tests := []struct {
		desc   string
		input  interface{}
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input:  models.Brand{ID: 3, Name: "Nike"},
			output: models.Brand{ID: 3, Name: "Nike"},
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().CreateBrand(gomock.AssignableToTypeOf(&gofr.Context{}), models.Brand{ID: 3, Name: "Nike"}).
					Return(models.Brand{ID: 3, Name: "Nike"}, nil),
			}},
		{desc: "Fail",
			input:  "nike",
			output: nil,
			expErr: errors.InvalidParam{Param: []string{"body"}},
			calls:  nil,
		},
		{desc: "Fail",
			input:  models.Brand{},
			output: models.Brand{},
			expErr: errors.MissingParam{Param: []string{"body"}},
			calls: []*gomock.Call{
				serviceMock.EXPECT().CreateBrand(gomock.AssignableToTypeOf(&gofr.Context{}), models.Brand{}).
					Return(models.Brand{}, errors.MissingParam{Param: []string{"body"}}),
			}},
	}
	for i, val := range tests {
		body, _ := json.Marshal(val.input)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(body))

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)

		h := New(serviceMock)
		out, err := h.Create(ctx)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}

// TestUpdate is a test function which uses mocks to test Update handler
func TestUpdate(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	serviceMock := service.NewMockBrand(ctrl)
	tests := []struct {
		desc   string
		input1 string
		input2 interface{}
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{
		{desc: "Success",
			input1: "6",
			input2: models.Brand{ID: 6, Name: "bru"},
			output: models.Brand{ID: 6, Name: "bru"},
			expErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().UpdateBrand(gomock.AssignableToTypeOf(&gofr.Context{}), 6, models.Brand{ID: 6, Name: "bru"}).
					Return(models.Brand{ID: 6, Name: "bru"}, nil),
			}},
		{desc: "Fail",
			input1: "11",
			input2: models.Brand{},
			output: models.Brand{},
			expErr: errors.EntityNotFound{Entity: "brand"},
			calls: []*gomock.Call{
				serviceMock.EXPECT().UpdateBrand(gomock.AssignableToTypeOf(&gofr.Context{}), 11, models.Brand{}).
					Return(models.Brand{}, errors.EntityNotFound{Entity: "brand"}),
			}},
		{desc: "Fail",
			input1: "abc",
			input2: models.Brand{},
			output: models.Brand{},
			expErr: errors.InvalidParam{Param: []string{"id"}},
			calls:  nil,
		},
		{desc: "Fail",
			input1: "1",
			input2: "nike",
			output: nil,
			expErr: errors.InvalidParam{Param: []string{"body"}},
			calls:  nil,
		},
		{desc: "Fail",
			input1: "",
			input2: models.Brand{},
			output: models.Brand{},
			expErr: errors.MissingParam{Param: []string{"id"}},
			calls:  nil,
		},
	}

	for i, val := range tests {
		body, _ := json.Marshal(val.input2)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(body))

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": val.input1,
		})

		h := New(serviceMock)
		out, err := h.Update(ctx)
		assert.Equalf(t, val.output, out, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}
