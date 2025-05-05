package db

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	dsn := "postgres://postgres:postgres@localhost:5432/ndot?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// store := NewStore(pool)
	testStore = NewStore(pool)

	os.Exit(m.Run())

}
