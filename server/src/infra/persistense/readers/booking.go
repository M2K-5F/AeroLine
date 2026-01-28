package readers

import (
	"aeroline/src/domain/booking"
	"aeroline/src/domain/shared"
	"aeroline/src/infra/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingReader struct {
	pool *pgxpool.Pool
}

func (ths BookingReader) GetTicketByID(ctx context.Context, id booking.TicketID) (*booking.Ticket, error) {
	var row models.TicketRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select *, (price).* from tickets where id = $1;
	`, id.String(),
	); err != nil {
		return nil, err
	}

	return booking.RestoreTicket(booking.TicketSnapshot{
		ID:           row.ID,
		BuyerID:      row.BuyerID,
		Price:        shared.RestorePrice(row.Amount, row.Currency),
		FlightSeatID: row.FlightSeatID,
	}), nil
}
