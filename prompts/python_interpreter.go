package prompts

import (
	"app/gemini"
	"fmt"
)

var PYTHON_INTERPRETER_PROMPT = `{
	"context": "%s",
	"role": "%s",
	"task": "%s",
	"instructions": [
		"- The output_format must include just one key called 'result'.",
		"- The 'result' key contains the result of the instructions emulation"
	],
	"output_format": {
		"result": "RESULT_PYTHON_INSTRUCTIONS"
	},
}`

func PythonInterpreter(userPrompt string) string {
	context := gemini.BuildHistory()
	role := "Act as Python interpreter. You are able to return realistic syntax error" // Python || os.Args[1]?

	return fmt.Sprintf(
		PYTHON_INTERPRETER_PROMPT,
		context,
		role,
		userPrompt,
	)
}
