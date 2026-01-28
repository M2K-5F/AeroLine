package readers

import (
	"aeroline/src/domain/user"
	"aeroline/src/infra/persistense/models"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserReader struct {
	pool *pgxpool.Pool
}

func (ths *UserReader) GetUserByID(ctx context.Context, id user.UserID) (*user.User, error) {
	var row models.UserRow
	if err := pgxscan.Get(ctx, ths.pool, &row, `
		select * from users 
		where id = $1 
		limit 1;
	`, id.String(),
	); err != nil {
		return nil, err
	}

	permissions := make([]user.Permission, len(row.Permissions))
	for i, permission := range row.Permissions {
		permissions[i] = user.Permission(permission)
	}

	return user.Restore(user.Snapshot{
		ID:           row.ID,
		Username:     user.Username(row.Username),
		PasswordHash: user.Password(row.PasswordHash),
		Permissions:  permissions,
	}), nil
}

func NewUserReader(pool *pgxpool.Pool) UserReader {
	return UserReader{
		pool: pool,
	}
}
