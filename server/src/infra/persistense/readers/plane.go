package readers

import (
	"aeroline/src/domain/plane"
	"aeroline/src/infra/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PlaneReader struct {
	pool *pgxpool.Pool
}

func (ths PlaneReader) GetPlaneByID(ctx context.Context, id plane.PlaneID) (*plane.Plane, error) {
	var row models.PlaneRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from planes 
		where id = $1
		limit 1;
	`, id.String()); err != nil {
		return nil, err
	}

	return plane.RestorePlane(plane.PlaneSnapshot{
		ID:   row.ID,
		Name: plane.Name(row.Name),
	}), nil
}

func (ths PlaneReader) GetSeatByID(ctx context.Context, id plane.SeatID) (*plane.Seat, error) {
	var row models.SeatRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from seats
		where id = $1
		limit 1;
	`, id.String()); err != nil {
		return nil, err
	}

	return plane.RestoreSeat(plane.SeatSnapshot{
		ID:      row.ID,
		Serial:  row.Serial,
		PlaneID: row.PlaneID,
		Tag:     row.Tag,
		Class:   plane.Class(row.Class),
	}), nil
}

func NewPlaneReader(pool *pgxpool.Pool) PlaneReader {
	return PlaneReader{
		pool: pool,
	}
}
