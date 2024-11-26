package dbhandlers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Deposit(ctx context.Context, db *pgxpool.Pool, log *slog.Logger, id uuid.UUID, amount float64) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error("begin transaction", "error", err)
		return fmt.Errorf("begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	log.Debug("started DEPOSIT transaction")

	query := `UPDATE bank.accounts
					SET balance = balance + $1
					WHERE id = $2`

	_, err = tx.Exec(ctx, query, amount, id)
	if err != nil {
		return fmt.Errorf("update bank accounts: %v", err)
	}

	log.Debug("executed DEPOSIT update")

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	log.Debug("commited DEPOSIT transaction")

	return nil
}
