package writers

import (
	"aeroline/src/domain/flight_domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func saveFlight(ctx context.Context, tx pgx.Tx, flight *flight_domain.Flight, others ...*flight_domain.Flight) error {
	flights := append([]*flight_domain.Flight{flight}, others...)

	ids := make([]string, len(flights))
	depCodes := make([]string, len(flights))
	arrCodes := make([]string, len(flights))
	planeIDs := make([]string, len(flights))
	arrTimes := make([]time.Time, len(flights))
	depTimes := make([]time.Time, len(flights))

	for i, f := range flights {
		spt := f.Snapshot()
		ids[i] = spt.ID.String()
		depCodes[i] = spt.Departure.Code
		arrCodes[i] = spt.Arrival.Code
		planeIDs[i] = spt.PlaneID.String()
		arrTimes[i] = spt.ArrivalTime
		depTimes[i] = spt.DepartureTime
	}

	_, err := tx.Exec(ctx, `
		insert into flights (id, departure_aip_code, arrival_aip_code, plane_id, arrival_time, departure_time)
		select * from unnest($1::uuid[], $2::text[], $3::text[], $4::uuid[], $5::timestamp[], $6::timestamp[])
		on conflict (id) do update set
			departure_aip_code = excluded.departure_aip_code,
			arrival_aip_code   = excluded.arrival_aip_code,
			plane_id           = excluded.plane_id,
			arrival_time       = excluded.arrival_time,
			departure_time     = excluded.departure_time;
	`, ids, depCodes, arrCodes, planeIDs, arrTimes, depTimes)

	return err
}

func saveFlightSeat(ctx context.Context, tx pgx.Tx, flightSeat *flight_domain.FlightSeat, others ...*flight_domain.FlightSeat) error {
	fSeats := append([]*flight_domain.FlightSeat{flightSeat}, others...)

	ids := make([]string, len(fSeats))
	occupied := make([]bool, len(fSeats))
	amounts := make([]int64, len(fSeats))
	currencies := make([]string, len(fSeats))
	seatIDs := make([]string, len(fSeats))
	flightIDs := make([]string, len(fSeats))

	for i, fs := range fSeats {
		spt := fs.Snapshot()
		ids[i] = spt.ID.String()
		occupied[i] = spt.IsOccupied
		amounts[i] = spt.Price.Amount()
		currencies[i] = spt.Price.Currency().String()
		seatIDs[i] = spt.SeatID.String()
		flightIDs[i] = spt.FlightID.String()
	}

	_, err := tx.Exec(ctx, `
		insert into flight_seats (id, is_occupied, price, seat_id, flight_id)
		select 
			u.id, u.occ, row(u.amt, u.cur)::price, u.sid, u.fid
		from unnest($1::uuid[], $2::bool[], $3::numeric[], $4::text[], $5::uuid[], $6::uuid[]) 
			as u(id, occ, amt, cur, sid, fid)
		on conflict (id) do update set
			is_occupied = excluded.is_occupied,
			price       = excluded.price,
			seat_id     = excluded.seat_id,
			flight_id   = excluded.flight_id;
	`, ids, occupied, amounts, currencies, seatIDs, flightIDs)

	return err
}
