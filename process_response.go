package main

import (
	"app/gemini"
	"fmt"
)

func processResponse(response string, jsonMode bool) (string, error) {
	if response == "" {
		return "", fmt.Errorf("response cannot be empty")
	}
	gemini.AddResponse(response)
	if jsonMode {
		// here we have the JSON as string
		// we can unmarshal to custom struct
		return parseResponse(response), nil
	}
	return response, nil
}
