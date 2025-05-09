package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface{
	Querier
	CreateSNMPTempate(ctx context.Context, template SnmpTemplateUnion) error 
	UpdateSNMPTemplate(ctx context.Context, template SnmpTemplateUnion) error
	DeleteSNMPTemplate(ctx context.Context, id int32) error
	ListSNMPTemplates(ctx context.Context, limit int32, offset int32) ([]SnmpTemplateUnion, error)
	GetSNMPTemplate(ctx context.Context, id int32) (SnmpTemplateUnion, error)

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