package dfltenv

import "os"

func Getenv(key string, defaultvalue string) string {
	if v, f := os.LookupEnv(key); f {
		return v
	} else {
		return defaultvalue
	}
}
