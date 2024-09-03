package main

import (
	"log"
)

func main() {
	if err := loadEnv("./.env"); err != nil {
		log.Fatal(err)
	}
}
