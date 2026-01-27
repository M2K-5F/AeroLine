package interfaces

import (
	"aeroline/src/domain/booking"
	"aeroline/src/domain/flight"
	"aeroline/src/domain/plane"
	"aeroline/src/domain/user"
	"context"
)

type IWriter interface {
	Execute(ctx context.Context, fn func(writer ITransactionWriter) error) error
}

type ITransactionWriter interface {
	user.IWriter
	plane.IWriter
	flight.IWriter
	booking.IWriter
}
