package handler

import (
	"strconv"
	"zopsmart-mini-project/datastore"
	"zopsmart-mini-project/model"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store datastore.Booking
}

func New(s datastore.Booking) handler {
	return handler{store: s}
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	// if ID is not present in path variables
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := validateID(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.store.GetByID(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "booking",
			ID:     id,
		}
	}

	return resp, nil
}
func (h handler) GetAllBookings(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.store.GetAllBookings(ctx)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "bookings",
		}
	}
	return resp, nil
}
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var booking model.Booking

	if err := ctx.Bind(&booking); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, &booking)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var booking model.Booking
	if err = ctx.Bind(&booking); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	booking.ID = id

	resp, err := h.store.Update(ctx, &booking)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.store.Delete(ctx, id); err != nil {
		return nil, err
	}

	return "Deleted successfully", nil
}

// Validating ID
func validateID(id string) (int, error) {
	// converting string id to integer
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return res, err
}
