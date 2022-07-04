package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connString := os.Getenv("AZ_STORAGE_ACCOUNT_CONN_STRING")
	fmt.Println(connString)
}
