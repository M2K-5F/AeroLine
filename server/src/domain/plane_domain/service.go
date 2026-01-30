package plane_domain

import (
	"aeroline/src/domain/shared"
	"fmt"
)

func NewPlane(name string) (*Plane, []*Seat, error) {
	plane, err := newPlane(name)
	if err != nil {
		return nil, nil, err
	}

	seats := make([]*Seat, 52)
	for i := range len(seats) {
		switch {
		case i < 12:
			seat, err := newBusinessSeat(plane.ID(), i+1)
			if err != nil {
				return nil, nil, err
			}
			seats[i] = seat

		case i >= 12:
			seat, err := newEconomySeat(plane.ID(), i+1)
			if err != nil {
				return nil, nil, err
			}
			seats[i] = seat
		}
	}

	return plane, seats, nil
}

func newPlane(name string) (*Plane, error) {
	nameVO, err := NewName(name)
	if err != nil {
		return nil, err
	}
	return &Plane{id: PlaneID{ID: shared.NewID()}, name: nameVO}, nil
}

func newBusinessSeat(planeID PlaneID, serial int) (*Seat, error) {
	if serial < 1 {
		return nil, ErrUnknownSeatClass
	}

	return &Seat{
		id:      SeatID{ID: shared.NewID()},
		serial:  serial,
		planeID: planeID,
		tag:     fmt.Sprint("A", serial),
		class:   BusinessSeat,
	}, nil
}

func newEconomySeat(planeID PlaneID, serial int) (*Seat, error) {
	if serial < 1 {
		return nil, ErrUnknownSeatClass
	}

	return &Seat{
		id:      SeatID{ID: shared.NewID()},
		serial:  serial,
		planeID: planeID,
		tag:     fmt.Sprint("B", serial),
		class:   EconomySeat,
	}, nil
}
