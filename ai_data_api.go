package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type TogetherRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type TogetherResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func getAIResponse(history []Message) (string, error) {
	apiKey := "keyhere"
	url := "https://api.together.xyz/v1/chat/completions"

	reqBody := TogetherRequest{
		Model:    "openai/gpt-oss-20b",
		Messages: history,
		Stream:   false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	fmt.Println("--- SENDING REQUEST ---")
	fmt.Println(string(jsonData))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println("Status Code:", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("--- RAW RESPONSE ---")
	fmt.Println(string(body))

	var result TogetherResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}

	return "no response from ze ai", nil

}
