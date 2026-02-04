package readers

import (
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
	"aeroline/src/infra/persistense/models"
	"context"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserReader struct {
	pool *pgxpool.Pool
}

func (ths UserReader) GetUserByID(ctx context.Context, id user_domain.UserID) (*user_domain.User, error) {
	var row models.UserRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from users 
		where id = $1 
		limit 1;
	`, id.String(),
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, shared.ErrDataNotFound
		}
		return nil, err
	}

	return user_domain.Restore(user_domain.Snapshot{
		ID:           row.ID,
		Username:     user_domain.Username(row.Username),
		PasswordHash: user_domain.Password(row.PasswordHash),
		Permissions: shared.Map(row.Permissions,
			func(p string) user_domain.Permission {
				return user_domain.Permission(p)
			},
		),
	}), nil
}

func (ths UserReader) GetUserByUsername(ctx context.Context, username user_domain.Username) (*user_domain.User, error) {
	var row models.UserRow

	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from users where 
		username = $1
		limit 1;
		`,
		username.String(),
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, shared.ErrDataNotFound
		}
		return nil, err
	}

	return user_domain.Restore(user_domain.Snapshot{
		ID:           row.ID,
		Username:     user_domain.Username(row.Username),
		PasswordHash: user_domain.Password(row.PasswordHash),
		Permissions: shared.Map(row.Permissions,
			func(perm string) user_domain.Permission {
				return user_domain.Permission(perm)
			},
		),
	}), nil
}

func NewUserReader(pool *pgxpool.Pool) UserReader {
	return UserReader{
		pool: pool,
	}
}
