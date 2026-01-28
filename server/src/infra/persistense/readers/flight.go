package readers

import (
	"aeroline/src/domain/flight"
	"aeroline/src/domain/shared"
	"aeroline/src/infra/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FlightReader struct {
	pool *pgxpool.Pool
}

func (ths FlightReader) GetFlightByID(ctx context.Context, id flight.FlightID) (*flight.Flight, error) {
	var row models.FlightRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from flights
		where id = $1
		limit 1;
	`, id.String()); err != nil {
		return nil, err
	}

	return flight.RestoreFlight(flight.FlightSnapshot{
		ID:            row.ID,
		Departure:     shared.GetAirportByCode(row.DepartureCode),
		DepartureTime: row.DepartureTime,
		Arrival:       shared.GetAirportByCode(row.ArrivalCode),
		ArrivalTime:   row.ArrivalTime,
		PlaneID:       row.PlaneID,
	}), nil
}

func (ths FlightReader) GetFlightSeatByID(ctx context.Context, id flight.FlightSeatID) (*flight.FlightSeat, error) {
	var row models.FlightSeatRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select *, (price).* from flight_seats
		where id = $1
		limit 1;
	`, id.String()); err != nil {
		return nil, err
	}

	return flight.RestoreFlightSeat(flight.FlightSeatSnapshot{
		ID:         row.ID,
		Price:      shared.RestorePrice(row.Amount, row.Currency),
		SeatID:     row.SeatID,
		FlightID:   row.FlightID,
		IsOccupied: row.IsOccupied,
	}), nil
}

func NewFlightReader(pool *pgxpool.Pool) FlightReader {
	return FlightReader{
		pool: pool,
	}
}
