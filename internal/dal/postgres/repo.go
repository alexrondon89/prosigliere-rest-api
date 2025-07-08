package postgres

import (
	"context"
	"github.com/alexrondon89/prosigliere-rest-api/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(ctx context.Context, url string) *PostgresRepo {
	ctx = util.SafeCtx(ctx)
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(err)
	}
	return &PostgresRepo{
		pool: pool,
	}
}
