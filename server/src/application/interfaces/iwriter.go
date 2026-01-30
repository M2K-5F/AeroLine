package interfaces

import (
	"aeroline/src/domain/booking_domain"
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/plane_domain"
	"aeroline/src/domain/user_domain"
	"context"
)

type IWriter interface {
	Execute(ctx context.Context, fn func(writer ITransactionWriter) error) error
}

type ITransactionWriter interface {
	user_domain.IWriter
	plane_domain.IWriter
	flight_domain.IWriter
	booking_domain.IWriter
}
