package main

import (
	"fmt"
	"os"
	"regexp"
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

func removeCodeMarks(s string) string {
	if strings.HasPrefix(s, "```") {
		lines := strings.Split(s, "\n")
		if lines[len(lines)-1] == "" {
			if len(lines) > 2 && strings.TrimSpace(lines[len(lines)-2]) == "```" {
				return strings.Join(lines[1:len(lines)-2], "\n")
			}
		} else {
			if len(lines) > 2 && strings.TrimSpace(lines[len(lines)-1]) == "```" {
				return strings.Join(lines[1:len(lines)-1], "\n")
			}
		}
	}
	return s
}

func escapeSpecialChars(str string) string {
	re := regexp.MustCompile(`[^\w\s]`)
	return re.ReplaceAllString(str, "\\\\$0")
}
