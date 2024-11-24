package ctxvalue

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func GetLog(ctx context.Context) (log *slog.Logger) {
	var ok bool

	ctxLog := ctx.Value(ValueLog)
	if ctxLog == nil {
		fmt.Printf("error get log from context\n")
		os.Exit(1)
	}

	if log, ok = ctxLog.(*slog.Logger); !ok {
		fmt.Printf("error get log from context\n")
		os.Exit(1)
	}

	return log
}
