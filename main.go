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
			fmt.Println("Error on request:", err)
			continue
		}
		promptTemplate := fmt.Sprintf("[message_to_llm]%s[/message_to_llm]", prompt)
		resTemplate := fmt.Sprintf("[answer_from_llm]%s[/answer_from_llm]", res)
		gemini.AddHistory(promptTemplate)
		gemini.AddHistory(resTemplate)
		fmt.Println(res)
	}
}
}
