package prompts

import (
	"app/gemini"
	"fmt"
)

var TALK_FILE_PROMPT = `{
	"context": "%s",
	"role": "%s",
	"task": "%s",
	"instructions": [
		"- The output_format must include just one key called 'result'.",
		"- The 'result' key contains the result of the instructions emulation",
		"- The 'task' key contains questions about the files in context between tags [content_filename=FILENAME]FILE_CONTENT_ESCAPED[/content_filename=FILENAME]",
	],
	"output_format": {
		"result": "RESULT_INSTRUCTIONS"
	},
}`

func TalkFile(prompt string) string {
	context := gemini.BuildContext()
	role := "Act as a software engineer especialized in refactor, performance and optimization."

	return fmt.Sprintf(
		TALK_FILE_PROMPT,
		context,
		role,
		prompt,
	)
}
