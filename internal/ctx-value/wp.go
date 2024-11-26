package ctxvalue

import (
	"context"
	"fmt"
	"javacode-test/util/workerpool"
	"os"
)

func GetWP(ctx context.Context) (wp *workerpool.Pool) {
	var ok bool

	ctxWP := ctx.Value(ValueWP)
	if ctxWP == nil {
		fmt.Printf("error get worker pool from context\n")
		os.Exit(1)
	}

	if wp, ok = ctxWP.(*workerpool.Pool); !ok {
		fmt.Printf("error get worker pool from context\n")
		os.Exit(1)
	}

	return wp
}
