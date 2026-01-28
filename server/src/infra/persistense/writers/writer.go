package writers

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/booking"
	"aeroline/src/domain/flight"
	"aeroline/src/domain/plane"
	"aeroline/src/domain/user"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Writer struct {
	pool *pgxpool.Pool
}

type TransactionWriter struct {
	txn pgx.Tx
	ctx context.Context
}

func (ths TransactionWriter) SaveFlight(flight *flight.Flight) error {
	return nil
}

func (ths TransactionWriter) SaveFlightSeat(flightSeat *flight.FlightSeat) error {
	return nil
}

func (ths TransactionWriter) SavePlane(plane *plane.Plane) error {
	return nil
}

func (ths TransactionWriter) SaveSeat(seat *plane.Seat) error {
	return nil
}

func (ths TransactionWriter) SaveTicket(ticket *booking.Ticket) error {
	return nil
}

func (ths TransactionWriter) SaveUser(user *user.User) error {
	return saveUser(ths.ctx, ths.txn, user)
}

func (txm *Writer) Execute(
	ctx context.Context,
	fn func(w interfaces.ITransactionWriter) error,
) (err error) {
	tx, err := txm.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if err := fn(&TransactionWriter{ctx: ctx, txn: tx}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func NewWriter(pool *pgxpool.Pool) *Writer {
	return &Writer{pool: pool}
}
