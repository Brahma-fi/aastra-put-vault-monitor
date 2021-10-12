package utils

import "fmt"

func GetDbURI() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		GetEnvVar("DB_USERNAME"),
		GetEnvVar("DB_PASSWORD"),
		GetEnvVar("DB_ENDPOINT"),
		GetEnvVar("DB_PORT"),
		GetEnvVar("DB_NAME"),
	)
}