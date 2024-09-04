package main

import "encoding/json"

func parseResponse(response string) string {
	type Result struct {
		Result string `json:"result"` // esto depende del output_format en el prompt
	}
	clean := removeCodeMarks(response)
	var r Result
	if err := json.Unmarshal([]byte(clean), &r); err != nil {
		return clean
	}
	return r.Result
}
