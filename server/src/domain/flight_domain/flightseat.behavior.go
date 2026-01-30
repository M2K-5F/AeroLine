package flight_domain

func (ths *FlightSeat) OccupieSeat() error {
	if ths.isOccupied {
		return ErrAlreadyOccupied
	}

	ths.isOccupied = true
	return nil
}
