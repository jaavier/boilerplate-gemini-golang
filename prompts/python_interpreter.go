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

func PythonInterpreter(prompt string) string {
	context := gemini.BuildContext()
	role := "Act as Python interpreter." // Python || os.Args[1]?

	return fmt.Sprintf(
		PYTHON_INTERPRETER_PROMPT,
		context,
		role,
		prompt,
	)
}
