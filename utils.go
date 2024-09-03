package main

import (
	"fmt"
	"os"
	"strings"
)

func loadEnv(filename string) error {
	if filename == "" {
		return fmt.Errorf("must specify filename")
	}
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading filename '", filename, "' :", err)
		return err
	}
	lines := strings.Split(string(bytes), "\n")
	apiKeyFound := false
	for _, line := range lines {
		splitLine := strings.Split(line, "=")
		if len(splitLine) < 2 {
			continue
		}
		variable := splitLine[0]
		value := splitLine[1]
		if variable == "GEMINI_API_KEY" {
			apiKeyFound = true
		}
		os.Setenv(variable, value)
	}
	if !apiKeyFound {
		return fmt.Errorf("environment variable 'GEMINI_API_KEY' not found on %s", filename)
	}
	return nil
}
