package app

import (
	"context"

	"github.com/P04KA/auth/database"
	"github.com/P04KA/auth/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

var conn *pgxpool.Pool
var ctx = context.Background()

func Run() error {
	if err := database.Migrate("postgresql://postgres1:postgres1@postgres:5432/postgres1"); err != nil {
		return errors.Wrap(err, "migrate")
	}

	var err error
	conn, err = storage.GetConnect("postgresql://postgres1:postgres1@postgres:5432/postgres1")
	if err != nil {
		return errors.Wrap(err, "connect")
	}

	if err := conn.Ping(ctx); err != nil {
		return errors.Wrap(err, "ping db")
	}
	return nil
}
