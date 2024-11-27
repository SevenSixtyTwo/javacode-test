package dbhandlers

import (
	"context"
	"fmt"
	ctxvalue "javacode-test/internal/ctx-value"

	"github.com/google/uuid"
)

func DbBalance(ctx context.Context, id uuid.UUID) (float64, error) {
	db := ctxvalue.GetDbPostgres(ctx)
	var balance float64
	query := `SELECT balance 
				FROM bank.accounts
				WHERE id = $1`

	row := db.QueryRow(ctx, query, id)
	if err := row.Scan(&balance); err != nil {
		return 0, fmt.Errorf("scan from row: %v", err)
	}

	return balance, nil
}
