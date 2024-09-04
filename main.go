package main

import (
	"app/gemini"
	"app/prompts"
	"fmt"
	"log"
)

func main() {
	if err := loadEnv("./.env"); err != nil {
		log.Fatal(err)
	}
	for {
		userPrompt := readPrompt("Prompt: ")
		prompt := prompts.PythonInterpreter(userPrompt)
		jsonMode := true
		response, err := gemini.Request(prompt)
		if err != nil {
			fmt.Println("Error gemini request:", err)
			continue
		}
		gemini.AddMessage(prompt)
		if newResponse, err := processResponse(response, jsonMode); err != nil {
			fmt.Println("Error processing response:", err.Error())
		} else {
			fmt.Println(newResponse)
		}
	}
}
