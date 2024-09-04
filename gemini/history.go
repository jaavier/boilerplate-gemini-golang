package gemini

var history = []string{}
var MAX_MEMORY = 100

func AddHistory(content string) {
	if len(history)%MAX_MEMORY == 0 && len(history) > 1 {
		history = history[1:]
	}
	history = append(history, content)
}

func GetHistory() []string {
	return history
}
