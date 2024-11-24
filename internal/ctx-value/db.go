package ctxvalue

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDbPostgres(ctx context.Context) (dbPostgres *pgxpool.Pool) {
	var ok bool

	ctxLog := ctx.Value(ValueDbPostgres)
	if ctxLog == nil {
		fmt.Printf("error get dbPostgres from context\n")
		os.Exit(1)
	}

	if dbPostgres, ok = ctxLog.(*pgxpool.Pool); !ok {
		fmt.Printf("error get dbPostgres from context\n")
		os.Exit(1)
	}

	return dbPostgres
}
