package dbhandlers

import (
	"context"
	"fmt"
	ctxvalue "javacode-test/internal/ctx-value"

	"github.com/google/uuid"
)

func Withdraw(ctx context.Context, id uuid.UUID, amount float64) error {
	db := ctxvalue.GetDbPostgres(ctx)
	log := ctxvalue.GetLog(ctx)

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error("begin transaction", "error", err)
		return fmt.Errorf("begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	log.Debug("started WITHDRAW transaction")

	query := `UPDATE bank.accounts
					SET balance = balance - $1 
					WHERE id = $2 AND balance >= $1`

	result, err := tx.Exec(ctx, query, amount, id)
	if err != nil {
		return fmt.Errorf("update bank accounts: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("insufficient funds or wrong uuid")
	}

	log.Debug("executed WITHDRAW update")

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	log.Debug("commited WITHDRAW transaction")

	return nil
}
