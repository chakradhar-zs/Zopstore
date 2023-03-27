package product

import (
	"fmt"
	"strconv"
	"strings"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/constants"
	"Day-19/internal/models"
	"Day-19/internal/service"
)

type Handler struct {
	svc service.Product
}

func New(s service.Product) *Handler {
	return &Handler{svc: s}
}

// Read handler takes gofr context and extracts id from url and calls GetProduct of service layer
func (h *Handler) Read(c *gofr.Context) (interface{}, error) {
	i := c.PathParam("id")
	brand := c.Param("brand")
	org := c.Param("organization")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.svc.GetProduct(c, id, brand)

	if err != nil {
		return nil, err
	}

	resp.Name = strings.TrimPrefix(resp.Name, org+"_")

	return resp, nil
}

// Create handler takes gofr context and extracts product details from request body
// Then adds organization as prefix to product name
// Then calls CreateProduct of service layer which returns product details and error if any
func (h *Handler) Create(c *gofr.Context) (interface{}, error) {
	var prod *models.Product

	org := c.Context.Value(constants.CtxValue)
	orgID := fmt.Sprint(org)
	err := c.Bind(&prod)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if orgID != "" {
		prod.Name = orgID + "_" + prod.Name
	}

	resp, err := h.svc.CreateProduct(c, prod)

	if err != nil {
		return nil, err
	}

	resp.Name = strings.TrimPrefix(resp.Name, orgID+"_")

	return resp, nil
}

// Update handler takes gofr context and extracts product details from request body
// Then adds organization as prefix to product name
// Then calls UpdateProduct of service layer which returns product details and error if any
func (h *Handler) Update(c *gofr.Context) (interface{}, error) {
	var prod models.Product

	org := c.Context.Value(constants.CtxValue)
	orgID := fmt.Sprint(org)

	i := c.PathParam("id")

	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	err = c.Bind(&prod)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if orgID != "" {
		prod.Name = orgID + "_" + prod.Name
	}

	resp, err := h.svc.UpdateProduct(c, id, &prod)

	if err != nil {
		return nil, err
	}

	resp.Name = strings.TrimPrefix(resp.Name, orgID+"_")

	return resp, nil
}

// Index handler takes gofr context calls GetProductByName and GetAllProducts of service layer based on url of request
func (h *Handler) Index(ctx *gofr.Context) (interface{}, error) {
	brand := ctx.Param("brand")
	org := ctx.Param("organization")
	name := ctx.Param("name")

	if org != "" && name != "" {
		name = org + "_" + name
		resp, _ := h.svc.GetProductByNAme(ctx, name, brand)

		for i := range resp {
			resp[i].Name = strings.TrimPrefix(resp[i].Name, org+"_")
		}

		return resp, nil
	}

	resp, err := h.svc.GetAllProducts(ctx, brand)

	for i := range resp {
		resp[i].Name = strings.TrimPrefix(resp[i].Name, org+"_")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
