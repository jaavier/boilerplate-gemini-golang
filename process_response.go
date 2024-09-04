package main

import (
	"app/gemini"
	"fmt"
)

func processResponse(response string, jsonMode bool) error {
	gemini.AddResponse(response)
	if jsonMode {
		// here we have the JSON as string
		// we can unmarshal to custom struct
		fmt.Println(parseResponse(response))
		return nil
	}
	fmt.Println(response)
	// here you can
	return nil
}
