package datastore

import (
	"zopsmart-mini-project/model"

	"gofr.dev/pkg/gofr"
)

type Booking interface {
	GetByID(ctx *gofr.Context, id string) (*model.Booking, error)
	Create(ctx *gofr.Context, model *model.Booking) (*model.Booking, error)
	Update(ctx *gofr.Context, model *model.Booking) (*model.Booking, error)
	Delete(ctx *gofr.Context, id int) error
}
