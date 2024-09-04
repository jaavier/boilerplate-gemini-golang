package gemini

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func BuildHistory() string {
	return strings.Join(GetHistory(), "\n")
}

func AppendFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	template := "[content_filename=%s]\n%s\n[/content_filename=%s]"
	newContext := fmt.Sprintf(template, filename, escapeSpecialChars(string(content)), filename)
	AddHistory(newContext)
	return nil
}

func escapeSpecialChars(str string) string {
	re := regexp.MustCompile(`[^\w\s]`)
	return re.ReplaceAllString(str, "\\\\$0")
}
