package main

import (
	"fmt"
	"log"
)

const (
	PERSONAL_ASSISTANT = iota
	PYTHON_INTERPRETER
	TALK_FILE
	LIMIT = iota - 1
)

func main() {
	if err := loadEnv("./.env"); err != nil {
		log.Fatal(err)
	}
	for {
		prompt := readPrompt("Prompt: ")
		if err := processPrompt(prompt, TALK_FILE); err != nil {
			fmt.Println("Error on request:", err)
		}
	}
}
