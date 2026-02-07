package writers

import (
	"aeroline/src/domain/booking_domain"
	"context"

	"github.com/jackc/pgx/v5"
)

func saveTicket(ctx context.Context, tx pgx.Tx, ticket *booking_domain.Ticket, others ...*booking_domain.Ticket) error {
	tickets := append([]*booking_domain.Ticket{ticket}, others...)

	ids := make([]string, len(tickets))
	buyerIDs := make([]string, len(tickets))
	amounts := make([]int64, len(tickets)) // или decimal, зависит от домена
	currencies := make([]string, len(tickets))
	flightSeatIDs := make([]string, len(tickets))

	for i, t := range tickets {
		spt := t.Snapshot()
		ids[i] = spt.ID.String()
		buyerIDs[i] = spt.BuyerID.String()
		amounts[i] = spt.Price.Amount()
		currencies[i] = spt.Price.Currency().String()
		flightSeatIDs[i] = spt.FlightSeatID.String()
	}

	_, err := tx.Exec(ctx, `
		insert into tickets (id, buyer_id, price, flight_seat_id)
		select 
			u.id, u.bid, row(u.amt, u.cur)::price, u.fsid
		from unnest($1::uuid[], $2::uuid[], $3::int[], $4::text[], $5::uuid[]) 
			as u(id, bid, amt, cur, fsid)
		on conflict (id) do update set
			buyer_id = excluded.buyer_id,
			price = excluded.price,
			flight_seat_id = excluded.flight_seat_id;
	`, ids, buyerIDs, amounts, currencies, flightSeatIDs)

	return err
}
