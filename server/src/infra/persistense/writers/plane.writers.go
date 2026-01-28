package writers

import (
	"aeroline/src/domain/plane"
	"context"

	"github.com/jackc/pgx/v5"
)

func savePlane(ctx context.Context, tx pgx.Tx, plane *plane.Plane) error {
	spt := plane.Snapshot()

	_, err := tx.Exec(ctx, `
		insert into 
		planes (id, name)
		values ($1, $2)
		on conflict (id) do
		update set 
			name = excluded.name;
		`,
		spt.ID.String(),
		spt.Name,
	)
	if err != nil {
		return err
	}

	return nil
}

func saveSeat(ctx context.Context, tx pgx.Tx, seat *plane.Seat) error {
	spt := seat.Snapshot()

	_, err := tx.Exec(ctx, `
		insert into 
		seats (id, plane_id, tag, serial, class)
		values ($1, $2, $3, $4, $5)
		on conflict (id) 
		do update set
			plane_id = excluded.plane_id,
			tag = excluded.tag,
			serial = excluded.serial,
			class = excluded.class;
		`,
		spt.ID.String(),
		spt.PlaneID.String(),
		spt.Tag,
		spt.Serial,
		spt.Class,
	)
	if err != nil {
		return err
	}

	return nil
}
