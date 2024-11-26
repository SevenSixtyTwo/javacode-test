package dbhandlers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetBalance(ctx context.Context, db *pgxpool.Pool, id uuid.UUID) (float64, error) {
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
