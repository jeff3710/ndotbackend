package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface{
	Querier
	CreateSNMPTempate(ctx context.Context, template SnmpTemplateUnion) error 

}

type SQLStore struct{
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries: New(connPool),
	}
}