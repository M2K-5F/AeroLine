package plane

import "context"

type IReader interface {
	GetPlaneByID(ctx context.Context, id PlaneID) (*Plane, error)
	GetSeatByID(ctx context.Context, id SeatID) (*Seat, error)
}

type IWriter interface {
	SavePlane(*Plane) error
	SaveSeat(*Seat) error
}
