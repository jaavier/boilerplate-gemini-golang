package main

import (
	"app/gemini"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := loadEnv("./.env"); err != nil {
		log.Fatal(err)
	}
	var prompt string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Prompt: ")
		if !scanner.Scan() {
			break
		}
		prompt = strings.TrimSpace(scanner.Text())
		res, err := gemini.Request(prompt)
		if err != nil {
			log.Fatal("Error on request:", err)
		}
		fmt.Println("AI answer:", res)
	}
}
