package datastore

import (
	"zopsmart-mini-project/model"

	"gofr.dev/pkg/gofr"
)

type Booking interface {
	// Get booking detail by ID
	// METHOD GET
	// EndPoint = /hotelbooking/:id
	GetByID(ctx *gofr.Context, id string) (*model.Booking, error)
	// Get All Bookings
	// METHOD GET
	// EndPoint = /hotelbooking
	GetAllBookings(ctx *gofr.Context) ([]model.Booking, error)
	// Add Booking Details
	// METHOD POST
	// EndPoint = /hotelbooking
	Create(ctx *gofr.Context, model *model.Booking) (*model.Booking, error)
	// Update Booking Details by ID
	// METHOD PUT
	// EndPoint = /hotelbooking/:id
	Update(ctx *gofr.Context, model *model.Booking) (*model.Booking, error)
	// DELETE Booking Details by ID
	// METHOD DELETE
	// EndPoint = /hotelbooking/:id
	Delete(ctx *gofr.Context, id int) error
}
