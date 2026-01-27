package plane

type Class string

const (
	EconomySeat  Class = "ECONOMY"
	BusinessSeat Class = "BUSINESS"
)

func (ths Class) String() string {
	return string(ths)
}

func (ths Class) IsValid() bool {
	switch ths {
	case EconomySeat, BusinessSeat:
		return true
	default:
		return false
	}
}
