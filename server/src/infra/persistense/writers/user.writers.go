package writers

import (
	"aeroline/src/domain/user"
	"context"

	"github.com/jackc/pgx/v5"
)

func saveUser(ctx context.Context, tx pgx.Tx, user *user.User) error {
	spt := user.Snapshot()

	_, err := tx.Exec(ctx, `
		insert into 
		users (id, username, permissions, password)
		values ($1, $2, $3, $4)
		on conflict (id) do
		update set 
			username = excluded.username,
			permissions = excluded.permissions,
			password = excluded.password;
		`,
		spt.ID.String(),
		spt.Username,
		spt.Permissions,
		spt.PasswordHash,
	)
	if err != nil {
		return err
	}

	return nil
}
