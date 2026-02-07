package shared

type Currency string

const (
	Rub Currency = "RUB"
	Usd Currency = "USD"
	Eur Currency = "EUR"
)

var validCurrencies = map[Currency]bool{
	Rub: true,
	Usd: true,
	Eur: true,
}

func (ths Currency) IsValid() bool {
	_, ok := validCurrencies[ths]
	return ok
}

func (ths Currency) String() string {
	return string(ths)
}
