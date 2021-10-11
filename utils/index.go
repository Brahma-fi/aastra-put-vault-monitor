package utils

import (
	"log"
	"os"
)

func LogInfo(message string) {
	log.Println("INFO:", message)
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}

func CheckError(err error, fallbackMessage string) {
	if err != nil {
		log.Fatal("ERROR: ", fallbackMessage, err)
	}
}