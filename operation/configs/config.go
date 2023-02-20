package configs

import (
	"log"
	"operation_service/models"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	err := models.Connect(GetEnv("MYSQL_USER"), GetEnv("MYSQL_PASS"), GetEnv("MYSQL_HOST_ADDR"), GetEnv("MYSQL_DB_NAME"))
	if err != nil {
		log.Fatalf("error in connecting to mysql-db:%s", err.Error())
	}
}

func loadEnv() {
	err := godotenv.Overload()
	if err != nil {
		log.Fatalf("error in loading .env file")
	}
}

func GetEnv(key string) string {
	val, _ := syscall.Getenv(key)
	return val
}
