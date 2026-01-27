package shared

type Price struct {
	amount   int64
	currency Currency
}

func NewPrice(amount int64, currency string) (Price, error) {

	currencyVO := Currency(currency)

	if !currencyVO.IsValid() {
		return Price{}, ErrUnknownCurrency
	}

	if amount < 0 {
		return Price{}, ErrNegativeAmount
	}

	return Price{
		amount:   amount,
		currency: currencyVO,
	}, nil
}

const (
	ErrUnknownCurrency = DomainError("Unknown currency")
	ErrNegativeAmount  = DomainError("Amount cannot be less than 0")
)
