package readers

import (
	"aeroline/src/domain/booking_domain"
	"aeroline/src/domain/shared"
	"aeroline/src/infra/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingReader struct {
	pool *pgxpool.Pool
}

func (ths BookingReader) GetTicketByID(ctx context.Context, id booking_domain.TicketID) (*booking_domain.Ticket, error) {
	var row models.TicketRow

	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select *, (price).* from tickets where id = $1;
	`, id.String(),
	); err != nil {
		return nil, err
	}

	return booking_domain.RestoreTicket(booking_domain.TicketSnapshot{
		ID:           row.ID,
		BuyerID:      row.BuyerID,
		Price:        shared.RestorePrice(row.Amount, row.Currency),
		FlightSeatID: row.FlightSeatID,
	}), nil
}

func NewBookingReader(pool *pgxpool.Pool) BookingReader {
	return BookingReader{
		pool: pool,
	}
}
