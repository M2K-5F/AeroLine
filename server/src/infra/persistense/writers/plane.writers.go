package writers

import (
	"aeroline/src/domain/plane_domain"
	"context"

	"github.com/jackc/pgx/v5"
)

func savePlane(ctx context.Context, tx pgx.Tx, plane *plane_domain.Plane, others ...*plane_domain.Plane) error {
	planes := append([]*plane_domain.Plane{plane}, others...)

	ids := make([]string, len(planes))
	names := make([]string, len(planes))

	for i, p := range planes {
		spt := p.Snapshot()
		ids[i] = spt.ID.String()
		names[i] = spt.Name.String()
	}

	_, err := tx.Exec(ctx, `
		insert into planes (id, name)
		select * from unnest($1::uuid[], $2::text[])
		on conflict (id) do update set 
			name = excluded.name;
	`, ids, names)

	return err
}

func saveSeat(ctx context.Context, tx pgx.Tx, seat *plane_domain.Seat, others ...*plane_domain.Seat) error {
	seats := append([]*plane_domain.Seat{seat}, others...)

	ids := make([]string, len(seats))
	planeIDs := make([]string, len(seats))
	tags := make([]string, len(seats))
	serials := make([]int, len(seats))
	classes := make([]string, len(seats))

	for i, s := range seats {
		spt := s.Snapshot()
		ids[i] = spt.ID.String()
		planeIDs[i] = spt.PlaneID.String()
		tags[i] = spt.Tag
		serials[i] = spt.Serial
		classes[i] = spt.Class.String()
	}

	_, err := tx.Exec(ctx, `
		insert into seats (id, plane_id, tag, serial, class)
		select * from unnest($1::uuid[], $2::uuid[], $3::text[], $4::int[], $5::text[])
		on conflict (id) do update set
			plane_id = excluded.plane_id,
			tag = excluded.tag,
			serial = excluded.serial,
			class = excluded.class;
	`, ids, planeIDs, tags, serials, classes)

	return err
}
