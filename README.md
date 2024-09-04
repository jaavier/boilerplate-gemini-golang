# boilerplate-gemini-golang
This is a boilerplate for starting to create applications with the Gemini LLM (by Google).

# How to use?
Add a mode in main.go under TALK_FILE. For example, lets add "MY_CUSTOM_MODE"
```go
// main.go
const (
	PERSONAL_ASSISTANT = iota
	PYTHON_INTERPRETER
	TALK_FILE
	MY_CUSTOM_MODE
	LIMIT = iota - 1
)
```
Now you have to add your own prompt in folder **prompts**. Lets add "**my_custom_prompt.go**"
```sh
prompts
├── my_custom_prompt.go
├── personal_assistant.go
├── python_interpreter.go
└── talk_file.go
```
`prompts/my_custom_prompt.go` must contain a prompt template and a function that receives the prompt written by the user and convert it to useful instructions for the LLM.

**prompts/my_custom_prompt.go**
```go
// prompts/my_custom_prompt.go
package prompts

import (
	"app/gemini"
	"fmt"
)

var YOUR_OWN_PROMPT = `Context: %s
User Prompt: %s

[YOUR_RESPONSE_HERE]`

// var YOUR_OWN_PROMPT = `HERE YOUR OWN PROMPT (JSON, MARKDOWN, YAML, ...)`

func YourOwnHandler(userPrompt string) string {
	return fmt.Sprintf(YOUR_OWN_PROMPT, gemini.BuildHistory(), userPrompt)
}
```
Now open **process_prompt** and add a case "**MY_CUSTOM_MODE**" to the switch

```go
switch mode {
case MY_CUSTOM_MODE:
  finalPrompt = prompts.YourOwnHandler(prompt)
  // if the LLM will answer a JSON, set jsonMode = true
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
```
# It's your turn!
Start building an app with LLM now ❤️