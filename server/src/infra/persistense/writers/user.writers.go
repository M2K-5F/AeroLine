package writers

import (
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
	"context"

	"github.com/jackc/pgx/v5"
)

func saveUser(ctx context.Context, tx pgx.Tx, user *user_domain.User, others ...*user_domain.User) error {
	users := append([]*user_domain.User{user}, others...)

	ids := make([]string, len(users))
	usernames := make([]string, len(users))
	permissions := make([][]string, len(users))
	passwords := make([]string, len(users))

	for i, u := range users {
		spt := u.Snapshot()
		ids[i] = spt.ID.String()
		usernames[i] = spt.Username.String()
		permissions[i] = shared.Map(spt.Permissions, func(p user_domain.Permission) string { return p.String() })
		passwords[i] = string(spt.PasswordHash)
	}

	_, err := tx.Exec(ctx, `
		insert into users (id, username, permissions, password)
		select * from unnest($1::uuid[], $2::text[], $3::text[][], $4::text[])
		on conflict (id) do update set
			username = excluded.username,
			permissions = excluded.permissions,
			password = excluded.password;
	`, ids, usernames, permissions, passwords)

	return err
}
