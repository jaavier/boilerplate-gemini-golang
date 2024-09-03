package main

import (
	"app/gemini"
	"fmt"
	"log"
)

func main() {
	if err := loadEnv("./.env"); err != nil {
		log.Fatal(err)
	}
	var prompt string
	fmt.Print("Prompt: ")
	fmt.Scanln(&prompt)
	res, err := gemini.Request(prompt)
	if err != nil {
		log.Fatal("Error on request:", err)
	}
	fmt.Println("AI answer:", res)
}
