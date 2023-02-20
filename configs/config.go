package configs

import (
	"log"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error in loading .env file")
	}
}

func GetEnv(key string) string {
	val, _ := syscall.Getenv(key)
	return val
}
