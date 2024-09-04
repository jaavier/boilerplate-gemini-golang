package main

import (
	"app/gemini"
	"app/prompts"
	"fmt"
)

func processPrompt(prompt string, mode uint) error {
	if mode > LIMIT {
		return fmt.Errorf("invalid mode")
	}
	finalPrompt := ""
	jsonMode := false
	switch mode {
	case PERSONAL_ASSISTANT:
		finalPrompt = prompts.BasicPersonalAssistant(prompt)
	case PYTHON_INTERPRETER:
		finalPrompt = prompts.PythonInterpreter(prompt)
		jsonMode = true
	case TALK_FILE:
		err := gemini.AppendFile("./main.go") // os.Args?
		if err != nil {
			fmt.Println(err)
			return err
		}
		jsonMode = true
		finalPrompt = prompts.TalkFile(prompt)
	}
	response, err := gemini.Request(finalPrompt)
	if err != nil {
		return err
	}
	gemini.AddMessage(prompt)
	return processResponse(response, jsonMode)
}
