package service

import (
	"coinflow/coinflow-server/collection-service/config"
	pkgHttp "coinflow/coinflow-server/pkg/http"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	LanguageEnglish = "en"
)

type TranslateRequestBody struct {
	TargetLanguageCode string `json:"targetLanguageCode"`
	Texts []string `json:"texts"`
}

type Translation struct {
	Text string `json:"text"`
	DetectedLanguageCode string `json:"detectedLanguageCode"`
}

type TranslateResponse struct {
	Translations []Translation
}

type TranslateError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func TranslateToLanguage(cli *http.Client, text string, lang string, cfg config.ServicesConfig) (string, error) {
	const method = "TranslateToLanguage"

	reqBody := &TranslateRequestBody{
		TargetLanguageCode: lang,
		Texts: []string{text},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Millisecond)
	defer cancel()

	req := pkgHttp.NewRequest(http.MethodPost, cfg.TranslateApiAddress).
		WithApiKeyAuthorization(cfg.TranslateApiKey).
		WithBody(reqBody).
		WithContext(ctx)

	if req.Err() != nil {
		return "", req.Err()
	}

	resp, err := cli.Do(req.Http())

	if err != nil {
		return "", fmt.Errorf("%s: %w", method, err)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("%s: %w", method, err)
	}

	var tlError TranslateError
	err = json.Unmarshal(data, &tlError)

	if err == nil && len(tlError.Message) > 0 {
		return "", fmt.Errorf("%s: response from %s: %s", method, cfg.TranslateApiAddress, tlError.Message)
	}

	var tls TranslateResponse
	err = json.Unmarshal(data, &tls)

	if err != nil {
		return "", fmt.Errorf("%s: %w", method, err)
	} else if len(tls.Translations) == 0 {
		return "", fmt.Errorf("%s: bad response from %s: null response length", method, cfg.TranslateApiAddress)
	}

	return tls.Translations[0].Text, nil
}
