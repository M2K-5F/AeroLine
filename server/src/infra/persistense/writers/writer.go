package writers

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/booking_domain"
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/plane_domain"
	"aeroline/src/domain/user_domain"
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

func (ths TransactionWriter) SaveFlight(flight *flight_domain.Flight) error {
	return saveFlight(ths.ctx, ths.txn, flight)
}

func (ths TransactionWriter) SaveFlightSeat(flightSeat *flight_domain.FlightSeat) error {
	return saveFlightSeat(ths.ctx, ths.txn, flightSeat)
}

func (ths TransactionWriter) SavePlane(plane *plane_domain.Plane) error {
	return savePlane(ths.ctx, ths.txn, plane)
}

func (ths TransactionWriter) SaveSeat(seat *plane_domain.Seat) error {
	return saveSeat(ths.ctx, ths.txn, seat)
}

func (ths TransactionWriter) SaveTicket(ticket *booking_domain.Ticket) error {
	return saveTicket(ths.ctx, ths.txn, ticket)
}

func (ths TransactionWriter) SaveUser(user *user_domain.User) error {
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
