# boilerplate-gemini-golang
This is a boilerplate for starting to create applications with the Gemini LLM (by Google).

# How to use?
You have to add your own prompt in folder **prompts**. Lets add "**my_custom_prompt.go**"
```sh
prompts
├── my_custom_prompt.go
├── personal_assistant.go
├── python_interpreter.go
└── talk_file.go
```
`prompts/my_custom_prompt.go` must contain a prompt template and a function that receives the prompt written by the user and convert it to useful instructions for the LLM.

This is a basic example, so we will receive the answer just in Mark Down.

**prompts/my_custom_prompt.go**
```go
// prompts/my_custom_prompt.go
package prompts

import (
	"app/gemini"
	"fmt"
)

var YOUR_OWN_PROMPT_TEMPLATE = `Context: %s
User Prompt: %s

[YOUR_RESPONSE_HERE]`

// var YOUR_OWN_PROMPT = `HERE YOUR OWN PROMPT (JSON, MARKDOWN, YAML, ...)`

func YourOwnHandler(userPrompt string) string {
	return fmt.Sprintf(YOUR_OWN_PROMPT_TEMPLATE, gemini.BuildHistory(), userPrompt)
}
```
## Adapt main.go
Now that we have "**YourOwnHandler**", we must use it on our **main.go**, line 16
```go
// main.go
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
		prompt := prompts.YourOwnHandler(userPrompt)
		jsonMode := true
		response, err := gemini.Request(prompt)
		if err != nil {
			fmt.Println("Error gemini request:", err)
			continue
		}
		gemini.AddMessage(prompt)
		processResponse(response, jsonMode)
	}
}
```

## Explaining utilities
Gemini package includes 2 functions for manage the history (memory):
- AddMessage: If you want the LLM remember something that you wrote, you must use:
```go
gemini.AddMessage(message)
```
- AddResponse: If you want the LLM remember something that it answered in the past, you must use
```go
gemini.AddResponse(response)
```
_Note: AddMessage and AddResponse are used by default, like that the LLM has "memory". The memory is limited to 100 messages but you can modify this limit on history.go (MAX_MEMORY)_

# It's your turn!
Start building an app with LLM now ❤️