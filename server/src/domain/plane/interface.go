package plane

type IReader interface {
	GetPlaneByID(id PlaneID) (*Plane, error)
	GetSeatByID(id SeatID) (*Seat, error)
}

type IWriter interface {
	SavePlane(*Plane) error
	SaveSeat(*Seat) error
}
