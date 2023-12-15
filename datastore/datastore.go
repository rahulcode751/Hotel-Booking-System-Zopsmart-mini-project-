package datastore

import (
	"database/sql"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"zopsmart-mini-project/model"
)

type booking struct{}

func New() *booking {
	return &booking{}
}

func (s *booking) GetByID(ctx *gofr.Context, id string) (*model.Booking, error) {
	var resp model.Booking

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,hotelname,customername,date,price,customercontact FROM bookings where id=?", id).
		Scan(&resp.ID, &resp.HotelName, &resp.CustomerName, &resp.Date, &resp.Price, &resp.CustomerContact)
	switch err {
	case sql.ErrNoRows:
		return &model.Booking{}, errors.EntityNotFound{Entity: "booking", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Booking{}, err
	}
}

func (s *booking) Create(ctx *gofr.Context, booking *model.Booking) (*model.Booking, error) {
	// var resp model.Booking

	var resp model.Booking

	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO bookings (hotelname,customername,date,price,customercontact) VALUES ( ?,?,?,?,?)", booking.HotelName, booking.CustomerName, booking.Date, booking.Price, booking.CustomerContact)
	if err != nil {
		return &model.Booking{}, errors.DB{Err: err}
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return &model.Booking{}, errors.DB{Err: err}
	}
	// ctx.Logger.Log(lastInsertID)
	err = ctx.DB().QueryRowContext(ctx, "SELECT id,hotelname,customername,date,price,customercontact FROM bookings WHERE id = ?", lastInsertID).Scan(&resp.ID, &resp.HotelName, &resp.CustomerName, &resp.Date, &resp.Price, &resp.CustomerContact)
	if err != nil {
		return &model.Booking{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *booking) Update(ctx *gofr.Context, booking *model.Booking) (*model.Booking, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE bookings SET hotelname=?,customername=?,date=?,price=?,customercontact=? WHERE id=?",
		booking.HotelName, booking.CustomerName, booking.Date, booking.Price, booking.CustomerContact, booking.ID)
	if err != nil {
		return &model.Booking{}, errors.DB{Err: err}
	}

	return booking, nil
}

func (s *booking) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM bookings where id=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
