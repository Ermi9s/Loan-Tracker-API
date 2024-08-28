package enivronment

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func OsGet(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	return os.Getenv(key)
}