package postgres

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"one-stop/internal/config"
)

type PostgresStore struct {
	Db *bun.DB
}

func NewPostgresClient(env *config.Environment) *PostgresStore {

	db := &PostgresStore{Db: bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(
			pgdriver.WithAddr(env.PostgresAddress),
			pgdriver.WithUser(env.PostgresUser),
			pgdriver.WithPassword(env.PostgresPassword),
			pgdriver.WithDatabase(env.PostgresDatabase),
			pgdriver.WithInsecure(true),
		)), pgdialect.New(),
	)}
	return db
}
