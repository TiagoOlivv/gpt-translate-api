package services

import (
	"encoding/json"
	"fmt"
	"gpt-translate-api/src/config"
	"github.com/go-resty/resty/v2"
)

func TranslateTextWithChatGPT(text, sourceLang, targetLang string) (string, error) {
	client := resty.New()
	cfg := config.AppConfig

	message := fmt.Sprintf("Translate the following text from %s to %s: %s", sourceLang, targetLang, text)

	requestBody := map[string]interface{}{
		"model": cfg.OpenAIModel,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are a helpful assistant that translates text.",
			},
			{
				"role":    "user",
				"content": message,
			},
		},
	}

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+cfg.OpenAIApiKey).
		SetHeader("OpenAI-Organization", cfg.OpenAIOrganization).
		SetHeader("OpenAI-Project", cfg.OpenAIProject).
		SetBody(requestBody).
		Post(cfg.OpenAIURL)

	if err != nil {
		return "", fmt.Errorf("error making request to OpenAI: %v", err)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(response.Body(), &result); err != nil {
		return "", fmt.Errorf("error unmarshaling OpenAI response: %v", err)
	}

	if errorField, ok := result["error"]; ok {
		return "", fmt.Errorf("OpenAI API error: %v", errorField)
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("invalid response from OpenAI: 'choices' field is missing or invalid")
	}

	messageMap, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response from OpenAI: 'message' field is missing")
	}

	translatedText, ok := messageMap["content"].(string)
	if !ok {
		return "", fmt.Errorf("invalid response from OpenAI: 'content' field is missing or invalid")
	}

	return translatedText, nil
}
