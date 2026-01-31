package shared

type Price struct {
	amount   int64
	currency Currency
}

func (ths Price) Amount() int64 {
	return ths.amount
}

func (ths Price) Currency() Currency {
	return ths.currency
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

func RestorePrice(amount int64, currency string) Price {
	return Price{
		amount:   amount,
		currency: Currency(currency),
	}
}

var (
	ErrUnknownCurrency = &AppError{
		Type: TypeValidation,
		Msg:  "Unknown currency",
	}
	ErrNegativeAmount = &AppError{
		Type: TypeBusinessLogic,
		Msg:  "Amount cannot be less than 0",
	}
)
