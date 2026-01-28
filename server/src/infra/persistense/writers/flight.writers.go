package writers

import (
	"aeroline/src/domain/flight"
	"context"

	"github.com/jackc/pgx/v5"
)

func saveFlight(ctx context.Context, tx pgx.Tx, flight *flight.Flight) error {
	spt := flight.Snapshot()

	_, err := tx.Exec(ctx, `
		insert into flights (
			id, 
			departure_aip_code, 
			arrival_aip_code, 
			plane_id, 
			arrival_time, 
			departure_time
		)
		values ($1, $2, $3, $4, $5, $6)
		on conflict (id) 
		do update set
			departure_aip_code = excluded.departure_aip_code,
			arrival_aip_code   = excluded.arrival_aip_code,
			plane_id           = excluded.plane_id,
			arrival_time       = excluded.arrival_time,
			departure_time     = excluded.departure_time;
		`,
		spt.ID.String(),
		spt.Departure.Code,
		spt.Arrival.Code,
		spt.PlaneID.String(),
		spt.ArrivalTime,
		spt.DepartureTime,
	)
	if err != nil {
		return err
	}

	return nil
}

func saveFlightSeat(ctx context.Context, tx pgx.Tx, flightSeat *flight.FlightSeat) error {
	spt := flightSeat.Snapshot()

	_, err := tx.Exec(ctx, `
		insert into flight_seats (
			id, 
			is_occupied, 
			price, 
			seat_id, 
			flight_id
		)
		values ($1, $2, ROW($3, $4)::price, $5, $6)
		on conflict (id) 
		do update set
			is_occupied = excluded.is_occupied,
			price       = excluded.price,
			seat_id     = excluded.seat_id,
			flight_id   = excluded.flight_id;
		`,
		spt.ID.String(),
		spt.IsOccupied,
		spt.Price.Amount(),
		spt.Price.Currency(),
		spt.SeatID.String(),
		spt.FlightID.String(),
	)
	if err != nil {
		return err
	}

	return nil
}
