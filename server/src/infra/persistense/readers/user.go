package readers

import (
	"aeroline/src/domain/user_domain"
	"aeroline/src/infra/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserReader struct {
	pool *pgxpool.Pool
}

func (ths *UserReader) GetUserByID(ctx context.Context, id user_domain.UserID) (*user_domain.User, error) {
	var row models.UserRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from users 
		where id = $1 
		limit 1;
	`, id.String(),
	); err != nil {
		return nil, err
	}

	permissions := make([]user_domain.Permission, len(row.Permissions))
	for i, permission := range row.Permissions {
		permissions[i] = user_domain.Permission(permission)
	}

	return user_domain.Restore(user_domain.Snapshot{
		ID:           row.ID,
		Username:     user_domain.Username(row.Username),
		PasswordHash: user_domain.Password(row.PasswordHash),
		Permissions:  permissions,
	}), nil
}

func NewUserReader(pool *pgxpool.Pool) UserReader {
	return UserReader{
		pool: pool,
	}
}
