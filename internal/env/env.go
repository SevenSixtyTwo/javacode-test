package env

import utilenv "javacode-test/util/env"

var (
	POSTGRES_CONN string
	LOG_LEVEL     string
)

func init() {
	utilenv.LoadFileEnv("./config/config.env")

	utilenv.LoadStrVar(&POSTGRES_CONN, "POSTGRES_CONN")
	utilenv.LoadStrVar(&LOG_LEVEL, "LOG_LEVEL")
}
