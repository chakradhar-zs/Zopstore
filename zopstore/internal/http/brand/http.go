package brand

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
	"Day-19/internal/service"
)

type Handler struct {
	svc service.Brand
}

func New(s service.Brand) *Handler {
	return &Handler{svc: s}
}

// Read handler takes gofr context and extracts id from url and calls GetBrand of service layer
func (h *Handler) Read(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.svc.GetBrand(ctx, id)

	if err != nil {
		return nil, errors.EntityNotFound{Entity: "brand"}
	}

	return resp, nil
}

// Create handler takes gofr context and extract  brand details from request body
// Then calls CreateBrand of service layer which returns brand details and error if any
func (h *Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var b models.Brand

	err := ctx.Bind(&b)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.svc.CreateBrand(ctx, b)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Update handler takes gofr context and extract brand details from request body
// Then calls UpdateBrand of service layer which returns brand details and error if any
func (h *Handler) Update(ctx *gofr.Context) (interface{}, error) {
	var b models.Brand

	i := ctx.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err = ctx.Bind(&b)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.svc.UpdateBrand(ctx, id, b)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
