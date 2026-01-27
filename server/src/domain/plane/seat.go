package plane

type Seat struct {
	id      SeatID
	planeID PlaneID
	tag     string
	serial  int
	class   Class
}

func (ths Seat) ID() SeatID {
	return ths.id
}

func (ths Seat) Tag() string {
	return ths.tag
}

func (ths Seat) Serial() int {
	return ths.serial
}

func (ths Seat) Class() Class {
	return ths.class
}
