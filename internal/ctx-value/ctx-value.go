package ctxvalue

type ValueKey int

const (
	ValueDbPostgres ValueKey = iota
	ValueLog
	ValueWP
)
