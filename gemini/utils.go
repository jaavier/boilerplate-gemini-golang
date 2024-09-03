package gemini

import (
	"strings"
)

func BuildContext() string {
	return strings.Join(GetHistory(), "\n")
}
