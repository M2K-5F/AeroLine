package plane

type PlaneSnapshot struct {
	ID   PlaneID
	Name string
}

func (ths Plane) Snapshot() PlaneSnapshot {
	return PlaneSnapshot{
		ID:   ths.id,
		Name: ths.name.String(),
	}
}

func RestorePlane(spt PlaneSnapshot) *Plane {
	return &Plane{
		id:   spt.ID,
		name: Name(spt.Name),
	}
}

type SeatSnapshot struct {
	Serial  int
	ID      SeatID
	PlaneID PlaneID
	Tag     string
	Class   string
}

func (ths Seat) Snapshot() SeatSnapshot {
	return SeatSnapshot{
		Serial:  ths.serial,
		ID:      ths.id,
		PlaneID: ths.planeID,
		Tag:     ths.tag,
		Class:   ths.class.String(),
	}
}

func RestoreSeat(spt SeatSnapshot) *Seat {
	return &Seat{
		id:      spt.ID,
		tag:     spt.Tag,
		planeID: spt.PlaneID,
		class:   Class(spt.Class),
		serial:  spt.Serial,
	}
}
