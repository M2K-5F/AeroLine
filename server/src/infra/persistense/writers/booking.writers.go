package writers

import (
	"aeroline/src/domain/booking"
	"context"

	"github.com/jackc/pgx/v5"
)

func saveTicket(ctx context.Context, tx pgx.Tx, ticket *booking.Ticket) error {
	spt := ticket.Snapshot()

	_, err := tx.Exec(ctx, `
		insert into tickets (
			id, 
			buyer_id, 
			price, 
			flight_seat_id
		)
		values ($1, $2, ROW($3, $4)::price, $5)
		on conflict (id) 
		do update set
			buyer_id       = excluded.buyer_id,
			price          = excluded.price,
			flight_seat_id = excluded.flight_seat_id;
		`,
		spt.ID.String(),
		spt.BuyerID.String(),
		spt.Price.Amount(),
		spt.Price.Currency(),
		spt.FlightSeatID.String(),
	)
	if err != nil {
		return err
	}

	return nil
}
