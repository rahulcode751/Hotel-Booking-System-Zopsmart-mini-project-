package datastore

import (
	"database/sql"
	"zopsmart-mini-project/model"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type booking struct{}

func New() *booking {
	return &booking{}
}

func (s *booking) GetByID(ctx *gofr.Context, id string) (*model.Booking, error) {
	var response model.Booking

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,hotelname,customername,date,price,customercontact,roomno FROM bookings where id=?", id).
		Scan(&response.ID, &response.HotelName, &response.CustomerName, &response.Date, &response.Price, &response.CustomerContact, &response.RoomNo)
	switch err {
	case sql.ErrNoRows:
		return &model.Booking{}, errors.EntityNotFound{Entity: "booking", ID: id}
	case nil:
		return &response, nil
	default:
		return &model.Booking{}, err
	}
}

func (s *booking) Create(ctx *gofr.Context, booking *model.Booking) (*model.Booking, error) {
	var response model.Booking

	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO bookings (hotelname,customername,date,price,customercontact,roomno) VALUES ( ?,?,?,?,?,?)", booking.HotelName, booking.CustomerName, booking.Date, booking.Price, booking.CustomerContact, booking.RoomNo)
	if err != nil {
		return &model.Booking{}, errors.DB{Err: err}
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return &model.Booking{}, errors.DB{Err: err}
	}
	err = ctx.DB().QueryRowContext(ctx, "SELECT id,hotelname,customername,date,price,customercontact,roomno FROM bookings WHERE id = ?", lastInsertID).Scan(&response.ID, &response.HotelName, &response.CustomerName, &response.Date, &response.Price, &response.CustomerContact, &response.RoomNo)
	if err != nil {
		return &model.Booking{}, errors.DB{Err: err}
	}
	return &response, nil
}

func (s *booking) Update(ctx *gofr.Context, booking *model.Booking) (*model.Booking, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE bookings SET hotelname=?,customername=?,date=?,price=?,customercontact=?,roomno=? WHERE id=?",
		booking.HotelName, booking.CustomerName, booking.Date, booking.Price, booking.CustomerContact, booking.RoomNo, booking.ID)
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
