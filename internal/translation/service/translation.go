package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
)

type TranslationService struct {
}

func NewTranslationService() TranslationService {
	return TranslationService{}
}

type getTranslationInput struct {
	Word   string `json:"q"`
	Sourse string `json:"source"`
	Target string `json:"target"`
}

type getTranslationResponse struct {
	Translation string `json:"translatedText"`
}

func (s TranslationService) GetTranslation(ctx context.Context, text string) (string, error) {
	input := &getTranslationInput{
		Word:   text,
		Sourse: "en",
		Target: "ru",
	}

	data, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", os.Getenv("LIBRETRANSLATE_URL"), bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	translation := &getTranslationResponse{}
	err = json.NewDecoder(resp.Body).Decode(translation)
	if err != nil {
		return "", err
	}

	return translation.Translation, nil
}
