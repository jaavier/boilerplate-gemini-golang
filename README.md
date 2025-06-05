# Gemini Golang Boilerplate

A lightweight starting point for building Go applications powered by the Gemini LLM. It ships with helpers for history management, response parsing, and a few prompt examples so you can focus on your idea.

## Quick Start

1. Copy `.env.example` to `.env` and add your `GEMINI_API_KEY` (grab one from [Google AI Studio](https://aistudio.google.com/app/apikey)).
2. Run `go run .` and type your prompt when asked.

The boilerplate includes three ready‑to‑use prompt handlers:

- **BasicPersonalAssistant** – converse with the model as a friendly assistant.
- **PythonInterpreter** – execute Python commands and receive the result as JSON.
- **TalkFile** – ask questions about the contents of a file wrapped between `[content_filename=...]` tags.

Choose one in `main.go` or create your own.

## Creating a Custom Prompt

Add a new file in `prompts/` that formats the user input into the instructions you want to send to Gemini. A minimal handler looks like this:

```go
// prompts/my_prompt.go
package prompts

import (
    "app/gemini"
    "fmt"
)

var TEMPLATE = `Context: %s
User Prompt: %s`

func MyPrompt(userPrompt string) string {
    return fmt.Sprintf(TEMPLATE, gemini.BuildHistory(), userPrompt)
}
```

Update `main.go` to call `prompts.MyPrompt` and set `jsonMode` depending on your output format.

## Processing Responses

`processResponse` can optionally parse JSON and stores the conversation history for you. The history length defaults to 100 messages but can be changed in `gemini/history.go`.

## Why Use This Boilerplate?

- Quickly experiment with Gemini using concise Go code.
- Extend or replace the existing prompts to fit your own workflows.
- Built‑in helpers remove repetitive boilerplate so you can test ideas faster.

Start hacking and let Gemini power your next project!
