package prompts

import (
	"app/gemini"
	"fmt"
)

var PERSONAL_ASSISTANT_PROMPT = `Context: %s
Role: You're a personal assistant.
User Prompt: %s

[HERE_YOUR_RESPONSE_PLAIN_TEXT_NO_TAGS]`

func BasicPersonalAssistant(prompt string) string {
	return fmt.Sprintf(PERSONAL_ASSISTANT_PROMPT, gemini.BuildContext(), prompt)
}
