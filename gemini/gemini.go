package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Request(prompt string) (string, error) {
	secretKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(secretKey))

	if err != nil {
		log.Println("cannot stablish connection to Gemini:", err)
		return "", fmt.Errorf("\u001B[31merror creating client: %v\u001B[0m", err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(float32(temperature))

	response, err := model.GenerateContent(ctx, genai.Text(prompt))

	if err != nil {
		fmt.Println("error generating content", err)
	}
	return parseResponse(response)
}

func parseResponse(response *genai.GenerateContentResponse) (string, error) {
	if response.Candidates[0].FinishReason != genai.FinishReasonStop {
		return "", fmt.Errorf("invalid response: %s", response.Candidates[0].FinishReason)
	}
	if response == nil ||
		response.Candidates == nil ||
		len(response.Candidates) == 0 ||
		len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("invalid response")
	}
	return fmt.Sprintf("%s", response.Candidates[0].Content.Parts[0]), nil
}

func AddMessage(message string) error {
	messageTemplate := fmt.Sprintf("[message_to_llm]%s[/message_to_llm]", message)
	AddHistory(messageTemplate)
	return nil
}

func AddResponse(response string) error {
	responseTemplate := fmt.Sprintf("[answer_from_llm]%s[/answer_from_llm]", response)
	AddHistory(responseTemplate)
	return nil
}
