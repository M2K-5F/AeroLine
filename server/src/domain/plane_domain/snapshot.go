package plane_domain

type PlaneSnapshot struct {
	ID   PlaneID
	Name Name
}

func (ths Plane) Snapshot() PlaneSnapshot {
	return PlaneSnapshot{
		ID:   ths.id,
		Name: ths.name,
	}
}

func RestorePlane(spt PlaneSnapshot) *Plane {
	return &Plane{
		id:   spt.ID,
		name: spt.Name,
	}
}

type SeatSnapshot struct {
	Serial  int
	ID      SeatID
	PlaneID PlaneID
	Tag     string
	Class   Class
}

func (ths Seat) Snapshot() SeatSnapshot {
	return SeatSnapshot{
		Serial:  ths.serial,
		ID:      ths.id,
		PlaneID: ths.planeID,
		Tag:     ths.tag,
		Class:   ths.class,
	}
}

func RestoreSeat(spt SeatSnapshot) *Seat {
	return &Seat{
		id:      spt.ID,
		tag:     spt.Tag,
		planeID: spt.PlaneID,
		class:   spt.Class,
		serial:  spt.Serial,
	}
}
