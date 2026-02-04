package readers

import (
	"aeroline/src/domain/plane_domain"
	"aeroline/src/domain/shared"
	"aeroline/src/infra/persistense/models"
	"context"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PlaneReader struct {
	pool *pgxpool.Pool
}

func (ths PlaneReader) GetPlaneByID(ctx context.Context, id plane_domain.PlaneID) (*plane_domain.Plane, error) {
	var row models.PlaneRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
			select * from planes 
			where id = $1
			limit 1;
		`, id.String(),
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, shared.ErrDataNotFound
		}
		return nil, err
	}

	return plane_domain.RestorePlane(plane_domain.PlaneSnapshot{
		ID:   row.ID,
		Name: plane_domain.Name(row.Name),
	}), nil
}

func (ths PlaneReader) GetSeatByID(ctx context.Context, id plane_domain.SeatID) (*plane_domain.Seat, error) {
	var row models.SeatRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from seats
		where id = $1
		limit 1;
	`, id.String()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, shared.ErrDataNotFound
		}
		return nil, err
	}

	return plane_domain.RestoreSeat(plane_domain.SeatSnapshot{
		ID:      row.ID,
		Serial:  row.Serial,
		PlaneID: row.PlaneID,
		Tag:     row.Tag,
		Class:   plane_domain.Class(row.Class),
	}), nil
}

func NewPlaneReader(pool *pgxpool.Pool) PlaneReader {
	return PlaneReader{
		pool: pool,
	}
}
