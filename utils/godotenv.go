package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}
	return os.Getenv(key)
}
