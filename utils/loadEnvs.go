package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	// Load environment variables
	fmt.Println("Load environment variables from .env file...")

	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}
}
