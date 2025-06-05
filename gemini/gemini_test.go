package gemini

import (
	"testing"

	"github.com/google/generative-ai-go/genai"
)

func TestParseResponseNil(t *testing.T) {
	_, err := parseResponse(nil)
	if err == nil {
		t.Fatal("expected error for nil response")
	}

	resp := &genai.GenerateContentResponse{}
	_, err = parseResponse(resp)
	if err == nil {
		t.Fatal("expected error for nil candidates")
	}

	resp.Candidates = []*genai.Candidate{}
	_, err = parseResponse(resp)
	if err == nil {
		t.Fatal("expected error for empty candidates")
	}

	resp.Candidates = []*genai.Candidate{{Content: &genai.Content{}}}
	_, err = parseResponse(resp)
	if err == nil {
		t.Fatal("expected error for empty parts")
	}
}

func TestParseResponseSuccess(t *testing.T) {
	part := genai.Text("ok")
	resp := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{{
			FinishReason: genai.FinishReasonStop,
			Content:      &genai.Content{Parts: []genai.Part{part}},
		}},
	}
	s, err := parseResponse(resp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if s != "ok" {
		t.Fatalf("expected 'ok', got %s", s)
	}
}
