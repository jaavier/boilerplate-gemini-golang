package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Stream(prompt string) {
	secretKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(secretKey))
	if err != nil {
		log.Println("err.gemini.1:", err)
		fmt.Printf("\u001B[31merror creating client: %v\u001B[0m\n", err)
		return
	}
	defer client.Close()
	generateContent(client, prompt)
}

func generateContent(client *genai.Client, prompt string) {
	ctx := context.Background()
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(float32(temperature))
	iter := model.GenerateContentStream(ctx, genai.Text(prompt))
	processStream(iter)
}

// PENDING: Add possibility to add the chat to context
func processStream(iter *genai.GenerateContentResponseIterator) {
	var output = ""
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			continue
		}
		if resp == nil || resp.Candidates == nil || len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			continue
		}
		var content = fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])
		output += content
		fmt.Print(content)
	}
}
