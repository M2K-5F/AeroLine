package fligth

func (ths *FlightSeat) OccupieSeat() error {
	if ths.isOccupied {
		return ErrAlreadyOccupied
	}

	ths.isOccupied = true
	return nil
}
